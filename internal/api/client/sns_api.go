package client

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func (c *APIClient) SendAlert(ctx context.Context, subject string, message string) error {
	_, err := c.snsClient.Publish(ctx, &sns.PublishInput{
		Message:  &message,
		Subject:  &subject,
		TopicArn: &c.alertsTopic,
	})

	return err
}
