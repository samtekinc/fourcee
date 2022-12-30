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

func (c *OrganizationsDatabaseClient) GetModulePropagationExecutionRequest(ctx context.Context, modulePropagationId string, modulePropagationExecutionRequestId string) (*models.ModulePropagationExecutionRequest, error) {
	response, err := c.dynamodb.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &c.modulePropagationExecutionRequestsTableName,
		Key: map[string]types.AttributeValue{
			"ModulePropagationId":                 &types.AttributeValueMemberS{Value: modulePropagationId},
			"ModulePropagationExecutionRequestId": &types.AttributeValueMemberS{Value: modulePropagationExecutionRequestId},
		},
	})
	if err != nil {
		return nil, err
	} else if response.Item == nil {
		return nil, helpers.NotFoundError{Message: fmt.Sprintf("Module Propagation Execution Request %q not found", modulePropagationExecutionRequestId)}
	}

	modulePropagationExecutionRequest := models.ModulePropagationExecutionRequest{}
	if err = attributevalue.UnmarshalMap(response.Item, &modulePropagationExecutionRequest); err != nil {
		return nil, err
	}

	return &modulePropagationExecutionRequest, nil
}

func (c *OrganizationsDatabaseClient) GetModulePropagationExecutionRequests(ctx context.Context, limit int32, cursor string) (*models.ModulePropagationExecutionRequests, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	scanInput := &dynamodb.ScanInput{
		TableName:         &c.modulePropagationExecutionRequestsTableName,
		Limit:             &limit,
		ExclusiveStartKey: startKey,
	}

	resultItems, lastEvaluatedKey, err := helpers.ScanDynamoDBUntilLimit(ctx, c.dynamodb, scanInput, limit, []string{"ModulePropagationId", "ModulePropagationExecutionRequestId"})
	if err != nil {
		return nil, err
	}

	modulePropagationExecutionRequests := []models.ModulePropagationExecutionRequest{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &modulePropagationExecutionRequests)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.ModulePropagationExecutionRequests{
		Items:      modulePropagationExecutionRequests,
		NextCursor: nextCursor,
	}, nil
}

func (c OrganizationsDatabaseClient) GetModulePropagationExecutionRequestsByModulePropagationId(ctx context.Context, modulePropagationId string, limit int32, cursor string) (*models.ModulePropagationExecutionRequests, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	keyCondition := expression.Key("ModulePropagationId").Equal(expression.Value(modulePropagationId))
	expressionBuilder := expression.NewBuilder().WithKeyCondition(keyCondition)
	expr, err := expressionBuilder.Build()
	if err != nil {
		return nil, err
	}

	queryInput := &dynamodb.QueryInput{
		TableName:                 &c.modulePropagationExecutionRequestsTableName,
		IndexName:                 aws.String("ModulePropagationId-RequestTime-index"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		Limit:                     &limit,
		ExclusiveStartKey:         startKey,
		ScanIndexForward:          aws.Bool(false),
	}

	resultItems, lastEvaluatedKey, err := helpers.QueryDynamoDBUntilLimit(ctx, c.dynamodb, queryInput, limit, []string{"ModulePropagationId", "ModulePropagationExecutionRequestId", "RequestTime"})
	if err != nil {
		return nil, err
	}

	modulePropagationExecutionRequests := []models.ModulePropagationExecutionRequest{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &modulePropagationExecutionRequests)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.ModulePropagationExecutionRequests{
		Items:      modulePropagationExecutionRequests,
		NextCursor: nextCursor,
	}, nil
}

func (c *OrganizationsDatabaseClient) PutModulePropagationExecutionRequest(ctx context.Context, input *models.ModulePropagationExecutionRequest) error {
	item, err := attributevalue.MarshalMap(input)
	if err != nil {
		return err
	}

	condition := expression.AttributeNotExists(expression.Name("ModulePropagationExecutionRequestId"))

	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return err
	}

	_, err = c.dynamodb.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:                 &c.modulePropagationExecutionRequestsTableName,
		Item:                      item,
		ConditionExpression:       expr.Condition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})

	ccfe := &types.ConditionalCheckFailedException{}
	switch {
	case errors.As(err, &ccfe):
		return helpers.AlreadyExistsError{Message: fmt.Sprintf("Module Propagation Execution Request %q already exists", input.ModulePropagationExecutionRequestId)}
	default:
		return err
	}
}

func (c *OrganizationsDatabaseClient) UpdateModulePropagationExecutionRequest(ctx context.Context, modulePropagationId string, modulePropagationExecutionRequestId string, update *models.ModulePropagationExecutionRequestUpdate) (*models.ModulePropagationExecutionRequest, error) {
	condition := expression.AttributeExists(expression.Name("ModulePropagationExecutionRequestId"))

	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return nil, err
	}

	updateBuilder := expression.UpdateBuilder{}
	if update.Status != nil {
		updateBuilder = updateBuilder.Set(expression.Name("Status"), expression.Value(*update.Status))
	}

	updateExpression, err := expression.NewBuilder().WithUpdate(updateBuilder).Build()
	if err != nil {
		return nil, err
	}

	updateInput := &dynamodb.UpdateItemInput{
		TableName:                 &c.modulePropagationExecutionRequestsTableName,
		Key:                       map[string]types.AttributeValue{"ModulePropagationExecutionRequestId": &types.AttributeValueMemberS{Value: modulePropagationExecutionRequestId}},
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
		return nil, helpers.NotFoundError{Message: fmt.Sprintf("Module Propagation Execution Request %q not found", modulePropagationExecutionRequestId)}
	default:
		if err != nil {
			return nil, err
		}
	}

	modulePropagationExecutionRequest := models.ModulePropagationExecutionRequest{}
	err = attributevalue.UnmarshalMap(result.Attributes, &modulePropagationExecutionRequest)
	if err != nil {
		return nil, err
	}

	return &modulePropagationExecutionRequest, nil
}
