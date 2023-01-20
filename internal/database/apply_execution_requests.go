package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/sheacloud/tfom/internal/helpers"
	"github.com/sheacloud/tfom/pkg/models"
)

func (c *DatabaseClient) GetApplyExecutionRequest(ctx context.Context, applyExecutionRequestId string) (*models.ApplyExecutionRequest, error) {
	fmt.Println("getting apply execution request")
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

func (c *DatabaseClient) GetApplyExecutionRequestsByIds(ctx context.Context, ids []string) ([]models.ApplyExecutionRequest, error) {
	fmt.Println("getting apply execution requests by ids", len(ids))
	var keys []map[string]types.AttributeValue

	for _, id := range ids {
		keys = append(keys, map[string]types.AttributeValue{
			"ApplyExecutionRequestId": &types.AttributeValueMemberS{Value: id},
		})
	}

	bii := dynamodb.BatchGetItemInput{
		RequestItems: map[string]types.KeysAndAttributes{
			c.applyExecutionsTableName: {
				Keys: keys,
			},
		},
	}
	items := []map[string]types.AttributeValue{}

	for {
		bgo, err := c.dynamodb.BatchGetItem(ctx, &bii)
		if err != nil {
			return nil, err
		}
		if bgo.Responses != nil {
			items = append(items, bgo.Responses[c.applyExecutionsTableName]...)
		}
		requestItems := bgo.UnprocessedKeys
		bii = dynamodb.BatchGetItemInput{RequestItems: requestItems}
		if len(requestItems) == 0 {
			break
		}
	}

	items = SortDynamoDBBatchResponses(keys, items)

	results := []models.ApplyExecutionRequest{}
	err := attributevalue.UnmarshalListOfMaps(items, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (c DatabaseClient) GetApplyExecutionRequests(ctx context.Context, limit int32, cursor string) (*models.ApplyExecutionRequests, error) {
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

func (c DatabaseClient) GetApplyExecutionRequestsByModuleAssignmentId(ctx context.Context, moduleAssignmentId string, limit int32, cursor string) (*models.ApplyExecutionRequests, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	keyCondition := expression.Key("ModuleAssignmentId").Equal(expression.Value(moduleAssignmentId))
	expressionBuilder := expression.NewBuilder().WithKeyCondition(keyCondition)
	expr, err := expressionBuilder.Build()
	if err != nil {
		return nil, err
	}

	queryInput := &dynamodb.QueryInput{
		TableName:                 &c.applyExecutionsTableName,
		IndexName:                 aws.String("ModuleAssignmentId-RequestTime-index"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		Limit:                     &limit,
		ExclusiveStartKey:         startKey,
		ScanIndexForward:          aws.Bool(false),
	}

	resultItems, lastEvaluatedKey, err := helpers.QueryDynamoDBUntilLimit(ctx, c.dynamodb, queryInput, limit, []string{"PlanExecutionRequestId", "ModuleAssignmentId", "RequestTime"})
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

func (c *DatabaseClient) PutApplyExecutionRequest(ctx context.Context, input *models.ApplyExecutionRequest) error {
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

func (c *DatabaseClient) DeleteApplyExecutionRequest(ctx context.Context, applyExecutionRequestId string) error {
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

func (c *DatabaseClient) UpdateApplyExecutionRequest(ctx context.Context, applyExecutionRequestId string, update *models.ApplyExecutionRequestUpdate) (*models.ApplyExecutionRequest, error) {
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
