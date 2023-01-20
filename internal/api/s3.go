package api

import (
	"context"
	"io"
)

func (c *APIClient) DownloadResultObject(ctx context.Context, objectKey string) ([]byte, error) {
	return c.dbClient.DownloadResultObject(ctx, objectKey)
}

func (c *APIClient) GetResultObjectWriter(ctx context.Context, objectKey string, withLiveUploads bool) (io.WriteCloser, error) {
	return c.dbClient.GetResultObjectWriter(ctx, objectKey, withLiveUploads)
}
