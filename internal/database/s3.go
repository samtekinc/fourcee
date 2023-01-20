package database

import (
	"bytes"
	"context"
	"io"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/sheacloud/tfom/internal/awsclients"
)

type ResultObjectWriter struct {
	s3              awsclients.PutObjectInterface
	bucket          string
	objectKey       string
	ctx             context.Context
	currentBuffer   *bytes.Buffer
	bufferMutex     sync.Mutex
	doneChan        chan bool
	wg              sync.WaitGroup
	ticker          *time.Ticker
	bytesWritten    int
	withLiveUploads bool
}

func (w *ResultObjectWriter) Write(p []byte) (n int, err error) {
	w.bufferMutex.Lock()
	defer w.bufferMutex.Unlock()
	n, err = w.currentBuffer.Write(p)
	w.bytesWritten += n
	return n, err
}

func (w *ResultObjectWriter) Close() error {
	if w.withLiveUploads {
		w.ticker.Stop()
		w.doneChan <- true
		w.wg.Wait()
	}

	return w.flush()
}

func (w *ResultObjectWriter) flush() error {
	w.bufferMutex.Lock()
	defer w.bufferMutex.Unlock()
	if w.bytesWritten > 0 {
		_, err := w.s3.PutObject(w.ctx, &s3.PutObjectInput{
			Bucket: &w.bucket,
			Key:    &w.objectKey,
			Body:   bytes.NewReader(w.currentBuffer.Bytes()),
		})
		if err != nil {
			return err
		}
		w.bytesWritten = 0
	}

	return nil
}

func (c *DatabaseClient) DownloadResultObject(ctx context.Context, objectKey string) ([]byte, error) {
	result, err := c.s3.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &c.resultsBucketName,
		Key:    &objectKey,
	})
	if err != nil {
		return nil, err
	}

	return io.ReadAll(result.Body)
}

func (c *DatabaseClient) GetResultObjectWriter(ctx context.Context, objectKey string, withLiveUploads bool) (io.WriteCloser, error) {
	writer := &ResultObjectWriter{
		s3:              c.s3,
		objectKey:       objectKey,
		bucket:          c.resultsBucketName,
		ctx:             ctx,
		currentBuffer:   bytes.NewBuffer([]byte{}),
		withLiveUploads: withLiveUploads,
		bytesWritten:    1, // so that the first flush will actually write something
	}
	err := writer.flush() // create an empty file in S3
	if err != nil {
		return nil, err
	}

	if withLiveUploads {
		writer.ticker = time.NewTicker(1 * time.Second)
		writer.doneChan = make(chan bool)

		writer.wg.Add(1)

		go func() {
			for {
				select {
				case <-writer.doneChan:
					writer.wg.Done()
					return
				case <-writer.ticker.C:
					writer.flush()
				}
			}
		}()
	}

	return writer, nil
}
