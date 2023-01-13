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

func (c *DatabaseClient) GetModulePropagation(ctx context.Context, modulePropagationId string) (*models.ModulePropagation, error) {
	response, err := c.dynamodb.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &c.propagationsTableName,
		Key: map[string]types.AttributeValue{
			"ModulePropagationId": &types.AttributeValueMemberS{Value: modulePropagationId},
		},
	})
	if err != nil {
		return nil, err
	} else if response.Item == nil {
		return nil, helpers.NotFoundError{Message: fmt.Sprintf("Module Propagation %q not found", modulePropagationId)}
	}

	modulePropagation := models.ModulePropagation{}
	if err = attributevalue.UnmarshalMap(response.Item, &modulePropagation); err != nil {
		return nil, err
	}

	return &modulePropagation, nil
}

func (c *DatabaseClient) GetModulePropagationsByIds(ctx context.Context, ids []string) ([]models.ModulePropagation, error) {
	fmt.Println("getting module propagations by ids", len(ids))
	var keys []map[string]types.AttributeValue

	for _, id := range ids {
		keys = append(keys, map[string]types.AttributeValue{
			"ModulePropagationId": &types.AttributeValueMemberS{Value: id},
		})
	}

	bii := dynamodb.BatchGetItemInput{
		RequestItems: map[string]types.KeysAndAttributes{
			c.propagationsTableName: {
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
			items = append(items, bgo.Responses[c.propagationsTableName]...)
		}
		requestItems := bgo.UnprocessedKeys
		bii = dynamodb.BatchGetItemInput{RequestItems: requestItems}
		if len(requestItems) == 0 {
			break
		}
	}

	items = SortDynamoDBBatchResponses(keys, items)

	results := []models.ModulePropagation{}
	err := attributevalue.UnmarshalListOfMaps(items, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (c DatabaseClient) GetModulePropagations(ctx context.Context, limit int32, cursor string) (*models.ModulePropagations, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	scanInput := &dynamodb.ScanInput{
		TableName:         &c.propagationsTableName,
		Limit:             &limit,
		ExclusiveStartKey: startKey,
	}

	resultItems, lastEvaluatedKey, err := helpers.ScanDynamoDBUntilLimit(ctx, c.dynamodb, scanInput, limit, []string{"ModulePropagationId"})
	if err != nil {
		return nil, err
	}

	modulePropagations := []models.ModulePropagation{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &modulePropagations)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.ModulePropagations{
		Items:      modulePropagations,
		NextCursor: nextCursor,
	}, nil
}

func (c DatabaseClient) GetModulePropagationsByModuleGroupId(ctx context.Context, moduleGroupId string, limit int32, cursor string) (*models.ModulePropagations, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	keyCondition := expression.Key("ModuleGroupId").Equal(expression.Value(moduleGroupId))
	expressionBuilder := expression.NewBuilder().WithKeyCondition(keyCondition)
	expr, err := expressionBuilder.Build()
	if err != nil {
		return nil, err
	}

	queryInput := &dynamodb.QueryInput{
		TableName:                 &c.propagationsTableName,
		IndexName:                 aws.String("ModuleGroupId-ModuleVersionId-index"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		Limit:                     &limit,
		ExclusiveStartKey:         startKey,
	}

	resultItems, lastEvaluatedKey, err := helpers.QueryDynamoDBUntilLimit(ctx, c.dynamodb, queryInput, limit, []string{"ModuleGroupId", "ModuleVersionId"})
	if err != nil {
		return nil, err
	}

	modulePropagations := []models.ModulePropagation{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &modulePropagations)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.ModulePropagations{
		Items:      modulePropagations,
		NextCursor: nextCursor,
	}, nil
}

func (c DatabaseClient) GetModulePropagationsByModuleVersionId(ctx context.Context, moduleVersionId string, limit int32, cursor string) (*models.ModulePropagations, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	keyCondition := expression.Key("ModuleVersionId").Equal(expression.Value(moduleVersionId))
	expressionBuilder := expression.NewBuilder().WithKeyCondition(keyCondition)
	expr, err := expressionBuilder.Build()
	if err != nil {
		return nil, err
	}

	queryInput := &dynamodb.QueryInput{
		TableName:                 &c.propagationsTableName,
		IndexName:                 aws.String("ModuleVersionId-index"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		Limit:                     &limit,
		ExclusiveStartKey:         startKey,
	}

	resultItems, lastEvaluatedKey, err := helpers.QueryDynamoDBUntilLimit(ctx, c.dynamodb, queryInput, limit, []string{"ModuleVersionId"})
	if err != nil {
		return nil, err
	}

	modulePropagations := []models.ModulePropagation{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &modulePropagations)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.ModulePropagations{
		Items:      modulePropagations,
		NextCursor: nextCursor,
	}, nil
}

func (c DatabaseClient) GetModulePropagationsByOrgUnitId(ctx context.Context, orgUnitId string, limit int32, cursor string) (*models.ModulePropagations, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	keyCondition := expression.Key("OrgUnitId").Equal(expression.Value(orgUnitId))
	expressionBuilder := expression.NewBuilder().WithKeyCondition(keyCondition)
	expr, err := expressionBuilder.Build()
	if err != nil {
		return nil, err
	}

	queryInput := &dynamodb.QueryInput{
		TableName:                 &c.propagationsTableName,
		IndexName:                 aws.String("OrgUnitId-index"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		Limit:                     &limit,
		ExclusiveStartKey:         startKey,
	}

	resultItems, lastEvaluatedKey, err := helpers.QueryDynamoDBUntilLimit(ctx, c.dynamodb, queryInput, limit, []string{"OrgUnitId"})
	if err != nil {
		return nil, err
	}

	modulePropagations := []models.ModulePropagation{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &modulePropagations)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.ModulePropagations{
		Items:      modulePropagations,
		NextCursor: nextCursor,
	}, nil
}

func (c DatabaseClient) GetModulePropagationsByOrgDimensionId(ctx context.Context, moduleVersionId string, limit int32, cursor string) (*models.ModulePropagations, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	keyCondition := expression.Key("OrgDimensionId").Equal(expression.Value(moduleVersionId))
	expressionBuilder := expression.NewBuilder().WithKeyCondition(keyCondition)
	expr, err := expressionBuilder.Build()
	if err != nil {
		return nil, err
	}

	queryInput := &dynamodb.QueryInput{
		TableName:                 &c.propagationsTableName,
		IndexName:                 aws.String("OrgDimensionId-OrgUnitId-index"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		Limit:                     &limit,
		ExclusiveStartKey:         startKey,
	}

	resultItems, lastEvaluatedKey, err := helpers.QueryDynamoDBUntilLimit(ctx, c.dynamodb, queryInput, limit, []string{"OrgDimensionId"})
	if err != nil {
		return nil, err
	}

	modulePropagations := []models.ModulePropagation{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &modulePropagations)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.ModulePropagations{
		Items:      modulePropagations,
		NextCursor: nextCursor,
	}, nil
}

func (c *DatabaseClient) PutModulePropagation(ctx context.Context, input *models.ModulePropagation) error {
	item, err := attributevalue.MarshalMap(input)
	if err != nil {
		return err
	}

	condition := expression.AttributeNotExists(expression.Name("ModulePropagationId"))

	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return err
	}

	_, err = c.dynamodb.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:                 &c.propagationsTableName,
		Item:                      item,
		ConditionExpression:       expr.Condition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})

	ccfe := &types.ConditionalCheckFailedException{}
	switch {
	case errors.As(err, &ccfe):
		return helpers.AlreadyExistsError{Message: fmt.Sprintf("Module Propagation %q already exists", input.ModulePropagationId)}
	default:
		return err
	}
}

func (c *DatabaseClient) DeleteModulePropagation(ctx context.Context, modulePropagationId string) error {
	condition := expression.AttributeExists(expression.Name("ModulePropagationId"))

	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return err
	}

	_, err = c.dynamodb.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: &c.propagationsTableName,
		Key: map[string]types.AttributeValue{
			"ModulePropagationId": &types.AttributeValueMemberS{Value: modulePropagationId},
		},
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ConditionExpression:       expr.Condition(),
	})
	ccfe := &types.ConditionalCheckFailedException{}
	switch {
	case errors.As(err, &ccfe):
		return helpers.NotFoundError{Message: fmt.Sprintf("Module Propagation %q not found", modulePropagationId)}
	default:
		return err
	}
}

func (c *DatabaseClient) UpdateModulePropagation(ctx context.Context, modulePropagationId string, update *models.ModulePropagationUpdate) (*models.ModulePropagation, error) {
	condition := expression.AttributeExists(expression.Name("ModulePropagationId"))

	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return nil, err
	}

	updateBuilder := expression.UpdateBuilder{}
	if update.OrgDimensionId != nil {
		updateBuilder = updateBuilder.Set(expression.Name("OrgDimensionId"), expression.Value(*update.OrgDimensionId))
	}
	if update.OrgUnitId != nil {
		updateBuilder = updateBuilder.Set(expression.Name("OrgUnitId"), expression.Value(*update.OrgUnitId))
	}
	if update.ModuleVersionId != nil {
		updateBuilder = updateBuilder.Set(expression.Name("ModuleVersionId"), expression.Value(*update.ModuleVersionId))
	}
	if update.Name != nil {
		updateBuilder = updateBuilder.Set(expression.Name("Name"), expression.Value(*update.Name))
	}
	if update.Description != nil {
		updateBuilder = updateBuilder.Set(expression.Name("Description"), expression.Value(*update.Description))
	}
	if update.Arguments != nil {
		updateBuilder = updateBuilder.Set(expression.Name("Arguments"), expression.Value(update.Arguments))
	}
	if update.AwsProviderConfigurations != nil {
		updateBuilder = updateBuilder.Set(expression.Name("AwsProviderConfigurations"), expression.Value(update.AwsProviderConfigurations))
	}

	updateExpression, err := expression.NewBuilder().WithUpdate(updateBuilder).Build()
	if err != nil {
		return nil, err
	}

	updateInput := &dynamodb.UpdateItemInput{
		TableName: &c.propagationsTableName,
		Key: map[string]types.AttributeValue{
			"ModulePropagationId": &types.AttributeValueMemberS{Value: modulePropagationId},
		},
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
		return nil, helpers.NotFoundError{Message: fmt.Sprintf("Module Propagation %q not found", modulePropagationId)}
	default:
		if err != nil {
			return nil, err
		}
	}

	modulePropagation := models.ModulePropagation{}
	err = attributevalue.UnmarshalMap(result.Attributes, &modulePropagation)
	if err != nil {
		return nil, err
	}

	return &modulePropagation, nil
}
