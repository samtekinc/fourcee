package database

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/sheacloud/tfom/internal/helpers"
	"github.com/sheacloud/tfom/pkg/execution/models"
)

func (c *ExecutionDatabaseClient) GetApplyExecutionRequest(ctx context.Context, applyExecutionRequestId string) (*models.ApplyExecutionRequest, error) {
	response, err := c.dynamodb.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &c.applyExecutionsTableName,
		Key: map[string]types.AttributeValue{
			"ApplyExecutionRequestId": &types.AttributeValueMemberS{Value: applyExecutionRequestId},
		},
	})
	if err != nil {
		return nil, err
	} else if response.Item == nil {
		return nil, helpers.NotFoundError{Message: fmt.Sprintf("Apply Execution Request %q not found", applyExecutionRequestId)}
	}

	applyExecutionRequest := models.ApplyExecutionRequest{}
	if err = attributevalue.UnmarshalMap(response.Item, &applyExecutionRequest); err != nil {
		return nil, err
	}

	return &applyExecutionRequest, nil
}

func (c ExecutionDatabaseClient) GetApplyExecutionRequests(ctx context.Context, limit int32, cursor string) (*models.ApplyExecutionRequests, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	scanInput := &dynamodb.ScanInput{
		TableName:         &c.applyExecutionsTableName,
		Limit:             &limit,
		ExclusiveStartKey: startKey,
	}

	resultItems, lastEvaluatedKey, err := helpers.ScanDynamoDBUntilLimit(ctx, c.dynamodb, scanInput, limit, []string{"ApplyExecutionRequestId"})
	if err != nil {
		return nil, err
	}

	applyExecutionRequests := []models.ApplyExecutionRequest{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &applyExecutionRequests)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.ApplyExecutionRequests{
		Items:      applyExecutionRequests,
		NextCursor: nextCursor,
	}, nil
}

func (c ExecutionDatabaseClient) GetApplyExecutionRequestsByStateKey(ctx context.Context, stateKey string, limit int32, cursor string) (*models.ApplyExecutionRequests, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	keyCondition := expression.Key("StateKey").Equal(expression.Value(stateKey))
	expressionBuilder := expression.NewBuilder().WithKeyCondition(keyCondition)
	expr, err := expressionBuilder.Build()
	if err != nil {
		return nil, err
	}

	queryInput := &dynamodb.QueryInput{
		TableName:                 &c.applyExecutionsTableName,
		IndexName:                 aws.String("StateKey-RequestTime-index"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		Limit:                     &limit,
		ExclusiveStartKey:         startKey,
	}

	resultItems, lastEvaluatedKey, err := helpers.QueryDynamoDBUntilLimit(ctx, c.dynamodb, queryInput, limit, []string{"StateKey", "RequestTime"})
	if err != nil {
		return nil, err
	}

	applyExecutionRequests := []models.ApplyExecutionRequest{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &applyExecutionRequests)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.ApplyExecutionRequests{
		Items:      applyExecutionRequests,
		NextCursor: nextCursor,
	}, nil
}

func (c *ExecutionDatabaseClient) PutApplyExecutionRequest(ctx context.Context, input *models.ApplyExecutionRequest) error {
	item, err := attributevalue.MarshalMap(input)
	if err != nil {
		return err
	}

	condition := expression.AttributeNotExists(expression.Name("ApplyExecutionRequestId"))

	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return err
	}

	_, err = c.dynamodb.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:                 &c.applyExecutionsTableName,
		Item:                      item,
		ConditionExpression:       expr.Condition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})

	ccfe := &types.ConditionalCheckFailedException{}
	switch {
	case errors.As(err, &ccfe):
		return helpers.AlreadyExistsError{Message: fmt.Sprintf("Apply Execution Request %q already exists", input.ApplyExecutionRequestId)}
	default:
		return err
	}
}

func (c *ExecutionDatabaseClient) DeleteApplyExecutionRequest(ctx context.Context, applyExecutionRequestId string) error {
	condition := expression.AttributeExists(expression.Name("ApplyExecutionRequestId"))

	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return err
	}

	_, err = c.dynamodb.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: &c.applyExecutionsTableName,
		Key: map[string]types.AttributeValue{
			"ApplyExecutionRequestId": &types.AttributeValueMemberS{Value: applyExecutionRequestId},
		},
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ConditionExpression:       expr.Condition(),
	})
	ccfe := &types.ConditionalCheckFailedException{}
	switch {
	case errors.As(err, &ccfe):
		return helpers.NotFoundError{Message: fmt.Sprintf("Apply Execution Request %q not found", applyExecutionRequestId)}
	default:
		return err
	}
}

func (c *ExecutionDatabaseClient) UpdateApplyExecutionRequest(ctx context.Context, applyExecutionRequestId string, update *models.ApplyExecutionRequestUpdate) (*models.ApplyExecutionRequest, error) {
	condition := expression.AttributeExists(expression.Name("ApplyExecutionRequestId"))

	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return nil, err
	}

	updateBuilder := expression.UpdateBuilder{}
	if update.InitOutputKey != nil {
		updateBuilder = updateBuilder.Set(expression.Name("InitOutputKey"), expression.Value(*update.InitOutputKey))
	}
	if update.ApplyOutputKey != nil {
		updateBuilder = updateBuilder.Set(expression.Name("ApplyOutputKey"), expression.Value(*update.ApplyOutputKey))
	}
	if update.Status != nil {
		updateBuilder = updateBuilder.Set(expression.Name("Status"), expression.Value(*update.Status))
	}

	updateExpression, err := expression.NewBuilder().WithUpdate(updateBuilder).Build()
	if err != nil {
		return nil, err
	}

	updateInput := &dynamodb.UpdateItemInput{
		TableName:                 &c.applyExecutionsTableName,
		Key:                       map[string]types.AttributeValue{"ApplyExecutionRequestId": &types.AttributeValueMemberS{Value: applyExecutionRequestId}},
		ExpressionAttributeNames:  updateExpression.Names(),
		ExpressionAttributeValues: updateExpression.Values(),
		UpdateExpression:          updateExpression.Update(),
		ConditionExpression:       expr.Condition(),
		ReturnValues:              types.ReturnValueAllNew,
	}

	result, err := c.dynamodb.UpdateItem(ctx, updateInput)
	ccfe := &types.ConditionalCheckFailedException{}
	switch {
	case errors.As(err, &ccfe):
		return nil, helpers.NotFoundError{Message: fmt.Sprintf("Apply Execution Request %q not found", applyExecutionRequestId)}
	default:
		if err != nil {
			return nil, err
		}
	}

	applyExecutionRequest := models.ApplyExecutionRequest{}
	err = attributevalue.UnmarshalMap(result.Attributes, &applyExecutionRequest)
	if err != nil {
		return nil, err
	}

	return &applyExecutionRequest, nil
}

func (c *ExecutionDatabaseClient) UploadTerraformApplyInitResults(ctx context.Context, applyExecutionRequestId string, initResults *models.TerraformInitOutput) (string, error) {
	outputKey := fmt.Sprintf("applies/%s/init-results.json", applyExecutionRequestId)

	initResultsBytes, err := json.Marshal(initResults)
	if err != nil {
		return "", err
	}

	_, err = c.s3.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &c.resultsBucketName,
		Key:    &outputKey,
		Body:   bytes.NewReader(initResultsBytes),
	})

	return outputKey, err
}

func (c *ExecutionDatabaseClient) UploadTerraformApplyResults(ctx context.Context, applyExecutionRequestId string, applyResults *models.TerraformApplyOutput) (string, error) {
	outputKey := fmt.Sprintf("applies/%s/apply-results.json", applyExecutionRequestId)

	applyResultsBytes, err := json.Marshal(applyResults)
	if err != nil {
		return "", err
	}

	_, err = c.s3.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &c.resultsBucketName,
		Key:    &outputKey,
		Body:   bytes.NewReader(applyResultsBytes),
	})

	return outputKey, err
}

func (c *ExecutionDatabaseClient) DownloadTerraformApplyInitResults(ctx context.Context, initResultsObjectKey string) (*models.TerraformInitOutput, error) {
	result, err := c.s3.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &c.resultsBucketName,
		Key:    &initResultsObjectKey,
	})
	if err != nil {
		return nil, err
	}

	initResults := models.TerraformInitOutput{}
	err = json.NewDecoder(result.Body).Decode(&initResults)
	if err != nil {
		return nil, err
	}

	return &initResults, nil
}

func (c *ExecutionDatabaseClient) DownloadTerraformApplyResults(ctx context.Context, applyResultsObjectKey string) (*models.TerraformApplyOutput, error) {
	result, err := c.s3.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &c.resultsBucketName,
		Key:    &applyResultsObjectKey,
	})
	if err != nil {
		return nil, err
	}

	applyResults := models.TerraformApplyOutput{}
	err = json.NewDecoder(result.Body).Decode(&applyResults)
	if err != nil {
		return nil, err
	}

	return &applyResults, nil
}
