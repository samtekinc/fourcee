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

func (c *OrganizationsDatabaseClient) GetModulePropagationDriftCheckRequest(ctx context.Context, modulePropagationId string, modulePropagationDriftCheckRequestId string) (*models.ModulePropagationDriftCheckRequest, error) {
	response, err := c.dynamodb.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &c.modulePropagationDriftCheckRequestsTableName,
		Key: map[string]types.AttributeValue{
			"ModulePropagationId":                  &types.AttributeValueMemberS{Value: modulePropagationId},
			"ModulePropagationDriftCheckRequestId": &types.AttributeValueMemberS{Value: modulePropagationDriftCheckRequestId},
		},
	})
	if err != nil {
		return nil, err
	} else if response.Item == nil {
		return nil, helpers.NotFoundError{Message: fmt.Sprintf("Module Propagation Sync Request %q not found", modulePropagationDriftCheckRequestId)}
	}

	modulePropagationDriftCheckRequest := models.ModulePropagationDriftCheckRequest{}
	if err = attributevalue.UnmarshalMap(response.Item, &modulePropagationDriftCheckRequest); err != nil {
		return nil, err
	}

	return &modulePropagationDriftCheckRequest, nil
}

func (c *OrganizationsDatabaseClient) GetModulePropagationDriftCheckRequests(ctx context.Context, limit int32, cursor string) (*models.ModulePropagationDriftCheckRequests, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	scanInput := &dynamodb.ScanInput{
		TableName:         &c.modulePropagationDriftCheckRequestsTableName,
		Limit:             &limit,
		ExclusiveStartKey: startKey,
	}

	resultItems, lastEvaluatedKey, err := helpers.ScanDynamoDBUntilLimit(ctx, c.dynamodb, scanInput, limit, []string{"ModulePropagationId", "ModulePropagationDriftCheckRequestId"})
	if err != nil {
		return nil, err
	}

	modulePropagationDriftCheckRequests := []models.ModulePropagationDriftCheckRequest{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &modulePropagationDriftCheckRequests)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.ModulePropagationDriftCheckRequests{
		Items:      modulePropagationDriftCheckRequests,
		NextCursor: nextCursor,
	}, nil
}

func (c OrganizationsDatabaseClient) GetModulePropagationDriftCheckRequestsByModulePropagationId(ctx context.Context, modulePropagationId string, limit int32, cursor string) (*models.ModulePropagationDriftCheckRequests, error) {
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
		TableName:                 &c.modulePropagationDriftCheckRequestsTableName,
		IndexName:                 aws.String("ModulePropagationId-RequestTime-index"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		Limit:                     &limit,
		ExclusiveStartKey:         startKey,
		ScanIndexForward:          aws.Bool(false),
	}

	resultItems, lastEvaluatedKey, err := helpers.QueryDynamoDBUntilLimit(ctx, c.dynamodb, queryInput, limit, []string{"ModulePropagationId", "ModulePropagationDriftCheckRequestId", "RequestTime"})
	if err != nil {
		return nil, err
	}

	modulePropagationDriftCheckRequests := []models.ModulePropagationDriftCheckRequest{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &modulePropagationDriftCheckRequests)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.ModulePropagationDriftCheckRequests{
		Items:      modulePropagationDriftCheckRequests,
		NextCursor: nextCursor,
	}, nil
}

func (c *OrganizationsDatabaseClient) PutModulePropagationDriftCheckRequest(ctx context.Context, input *models.ModulePropagationDriftCheckRequest) error {
	item, err := attributevalue.MarshalMap(input)
	if err != nil {
		return err
	}

	condition := expression.AttributeNotExists(expression.Name("ModulePropagationDriftCheckRequestId"))

	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return err
	}

	_, err = c.dynamodb.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:                 &c.modulePropagationDriftCheckRequestsTableName,
		Item:                      item,
		ConditionExpression:       expr.Condition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})

	ccfe := &types.ConditionalCheckFailedException{}
	switch {
	case errors.As(err, &ccfe):
		return helpers.AlreadyExistsError{Message: fmt.Sprintf("Module Propagation Sync Request %q already exists", input.ModulePropagationDriftCheckRequestId)}
	default:
		return err
	}
}

func (c *OrganizationsDatabaseClient) UpdateModulePropagationDriftCheckRequest(ctx context.Context, modulePropagationId string, modulePropagationDriftCheckRequestId string, update *models.ModulePropagationDriftCheckRequestUpdate) (*models.ModulePropagationDriftCheckRequest, error) {
	condition := expression.AttributeExists(expression.Name("ModulePropagationDriftCheckRequestId"))

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
		TableName:                 &c.modulePropagationDriftCheckRequestsTableName,
		Key:                       map[string]types.AttributeValue{"ModulePropagationDriftCheckRequestId": &types.AttributeValueMemberS{Value: modulePropagationDriftCheckRequestId}},
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
		return nil, helpers.NotFoundError{Message: fmt.Sprintf("Module Propagation Sync Request %q not found", modulePropagationDriftCheckRequestId)}
	default:
		if err != nil {
			return nil, err
		}
	}

	modulePropagationDriftCheckRequest := models.ModulePropagationDriftCheckRequest{}
	err = attributevalue.UnmarshalMap(result.Attributes, &modulePropagationDriftCheckRequest)
	if err != nil {
		return nil, err
	}

	return &modulePropagationDriftCheckRequest, nil
}
