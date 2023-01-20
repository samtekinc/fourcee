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
	"go.uber.org/zap"
)

func getExpressionForModuleAssignmentFilters(filters *models.ModuleAssignmentFilters) expression.ConditionBuilder {
	expr := expression.ConditionBuilder{}

	if filters == nil {
		return expr
	}

	if filters.IsPropagated != nil {
		var filter expression.ConditionBuilder
		if *filters.IsPropagated {
			filter = expression.Name("ModulePropagationId").AttributeExists()
		} else {
			filter = expression.Name("ModulePropagationId").AttributeNotExists()
		}
		if !expr.IsSet() {
			expr = filter
		} else {
			expr = expr.And(filter)
		}
	}

	return expr
}

func (c *DatabaseClient) GetModuleAssignment(ctx context.Context, moduleAssignmentId string) (*models.ModuleAssignment, error) {
	fmt.Println("getting module assignment")
	response, err := c.dynamodb.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &c.moduleAssignmentsTableName,
		Key: map[string]types.AttributeValue{
			"ModuleAssignmentId": &types.AttributeValueMemberS{Value: moduleAssignmentId},
		},
	})
	if err != nil {
		return nil, err
	} else if response.Item == nil {
		return nil, helpers.NotFoundError{Message: fmt.Sprintf("Module Assignment %q not found", moduleAssignmentId)}
	}

	moduleAssignment := models.ModuleAssignment{}
	if err = attributevalue.UnmarshalMap(response.Item, &moduleAssignment); err != nil {
		return nil, err
	}

	return &moduleAssignment, nil
}

func (c *DatabaseClient) GetModuleAssignmentsByIds(ctx context.Context, ids []string) ([]models.ModuleAssignment, error) {
	zap.L().Sugar().Debugw("GetModuleAssignmentsByIds", "ids", ids)
	var keys []map[string]types.AttributeValue

	for _, id := range ids {
		keys = append(keys, map[string]types.AttributeValue{
			"ModuleAssignmentId": &types.AttributeValueMemberS{Value: id},
		})
	}

	bii := dynamodb.BatchGetItemInput{
		RequestItems: map[string]types.KeysAndAttributes{
			c.moduleAssignmentsTableName: {
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
			items = append(items, bgo.Responses[c.moduleAssignmentsTableName]...)
		}
		requestItems := bgo.UnprocessedKeys
		bii = dynamodb.BatchGetItemInput{RequestItems: requestItems}
		if len(requestItems) == 0 {
			break
		}
	}

	items = SortDynamoDBBatchResponses(keys, items)

	results := []models.ModuleAssignment{}
	err := attributevalue.UnmarshalListOfMaps(items, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (c DatabaseClient) GetModuleAssignments(ctx context.Context, filters *models.ModuleAssignmentFilters, limit int32, cursor string) (*models.ModuleAssignments, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	var scanInput *dynamodb.ScanInput
	expressionBuilder := expression.NewBuilder()
	filterCondition := getExpressionForModuleAssignmentFilters(filters)

	if filterCondition.IsSet() {
		expressionBuilder = expressionBuilder.WithFilter(filterCondition)

		expr, err := expressionBuilder.Build()
		if err != nil {
			return nil, err
		}

		scanInput = &dynamodb.ScanInput{
			TableName:                 &c.moduleAssignmentsTableName,
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			FilterExpression:          expr.Filter(),
			Limit:                     &limit,
			ExclusiveStartKey:         startKey,
		}
	} else {
		scanInput = &dynamodb.ScanInput{
			TableName:         &c.moduleAssignmentsTableName,
			Limit:             &limit,
			ExclusiveStartKey: startKey,
		}
	}

	resultItems, lastEvaluatedKey, err := helpers.ScanDynamoDBUntilLimit(ctx, c.dynamodb, scanInput, limit, []string{"ModuleAssignmentId"})
	if err != nil {
		return nil, err
	}

	moduleAssignments := []models.ModuleAssignment{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &moduleAssignments)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.ModuleAssignments{
		Items:      moduleAssignments,
		NextCursor: nextCursor,
	}, nil
}

func (c DatabaseClient) GetModuleAssignmentsByModulePropagationId(ctx context.Context, modulePropagationId string, limit int32, cursor string) (*models.ModuleAssignments, error) {
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
		TableName:                 &c.moduleAssignmentsTableName,
		IndexName:                 aws.String("ModulePropagationId-OrgAccountId-index"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		Limit:                     &limit,
		ExclusiveStartKey:         startKey,
		ScanIndexForward:          aws.Bool(false),
	}

	resultItems, lastEvaluatedKey, err := helpers.QueryDynamoDBUntilLimit(ctx, c.dynamodb, queryInput, limit, []string{"ModuleAssignmentId", "ModulePropagationId", "OrgAccountId"})
	if err != nil {
		return nil, err
	}

	moduleAssignments := []models.ModuleAssignment{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &moduleAssignments)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.ModuleAssignments{
		Items:      moduleAssignments,
		NextCursor: nextCursor,
	}, nil
}

func (c DatabaseClient) GetModuleAssignmentsByOrgAccountId(ctx context.Context, orgAccountId string, limit int32, cursor string) (*models.ModuleAssignments, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	keyCondition := expression.Key("OrgAccountId").Equal(expression.Value(orgAccountId))
	expressionBuilder := expression.NewBuilder().WithKeyCondition(keyCondition)
	expr, err := expressionBuilder.Build()
	if err != nil {
		return nil, err
	}

	queryInput := &dynamodb.QueryInput{
		TableName:                 &c.moduleAssignmentsTableName,
		IndexName:                 aws.String("OrgAccountId-ModuleGroupId-index"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		Limit:                     &limit,
		ExclusiveStartKey:         startKey,
		ScanIndexForward:          aws.Bool(false),
	}

	resultItems, lastEvaluatedKey, err := helpers.QueryDynamoDBUntilLimit(ctx, c.dynamodb, queryInput, limit, []string{"ModuleAssignmentId", "OrgAccountId", "ModuleGroupId"})
	if err != nil {
		return nil, err
	}

	moduleAssignments := []models.ModuleAssignment{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &moduleAssignments)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.ModuleAssignments{
		Items:      moduleAssignments,
		NextCursor: nextCursor,
	}, nil
}

func (c DatabaseClient) GetModuleAssignmentsByModuleVersionId(ctx context.Context, moduleVersionId string, limit int32, cursor string) (*models.ModuleAssignments, error) {
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
		TableName:                 &c.moduleAssignmentsTableName,
		IndexName:                 aws.String("ModuleVersionId-index"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		Limit:                     &limit,
		ExclusiveStartKey:         startKey,
		ScanIndexForward:          aws.Bool(false),
	}

	resultItems, lastEvaluatedKey, err := helpers.QueryDynamoDBUntilLimit(ctx, c.dynamodb, queryInput, limit, []string{"ModuleAssignmentId", "ModuleVersionId"})
	if err != nil {
		return nil, err
	}

	moduleAssignments := []models.ModuleAssignment{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &moduleAssignments)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.ModuleAssignments{
		Items:      moduleAssignments,
		NextCursor: nextCursor,
	}, nil
}

func (c DatabaseClient) GetModuleAssignmentsByModuleGroupId(ctx context.Context, moduleGroupId string, limit int32, cursor string) (*models.ModuleAssignments, error) {
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
		TableName:                 &c.moduleAssignmentsTableName,
		IndexName:                 aws.String("ModuleGroupId-index"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		Limit:                     &limit,
		ExclusiveStartKey:         startKey,
		ScanIndexForward:          aws.Bool(false),
	}

	resultItems, lastEvaluatedKey, err := helpers.QueryDynamoDBUntilLimit(ctx, c.dynamodb, queryInput, limit, []string{"ModuleAssignmentId", "ModuleGroupId"})
	if err != nil {
		return nil, err
	}

	moduleAssignments := []models.ModuleAssignment{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &moduleAssignments)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.ModuleAssignments{
		Items:      moduleAssignments,
		NextCursor: nextCursor,
	}, nil
}

func (c *DatabaseClient) PutModuleAssignment(ctx context.Context, input *models.ModuleAssignment) error {
	item, err := attributevalue.MarshalMap(input)
	if err != nil {
		return err
	}

	condition := expression.AttributeNotExists(expression.Name("ModuleAssignmentId"))

	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return err
	}

	_, err = c.dynamodb.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:                 &c.moduleAssignmentsTableName,
		Item:                      item,
		ConditionExpression:       expr.Condition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})

	ccfe := &types.ConditionalCheckFailedException{}
	switch {
	case errors.As(err, &ccfe):
		return helpers.AlreadyExistsError{Message: fmt.Sprintf("Module Assignment %q already exists", input.ModuleAssignmentId)}
	default:
		return err
	}
}

func (c *DatabaseClient) UpdateModuleAssignment(ctx context.Context, moduleAssignmentId string, update *models.ModuleAssignmentUpdate) (*models.ModuleAssignment, error) {
	condition := expression.AttributeExists(expression.Name("ModuleAssignmentId"))

	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return nil, err
	}

	updateBuilder := expression.UpdateBuilder{}
	if update.Arguments != nil {
		updateBuilder = updateBuilder.Set(expression.Name("Arguments"), expression.Value(update.Arguments))
	}
	if update.AwsProviderConfigurations != nil {
		updateBuilder = updateBuilder.Set(expression.Name("AwsProviderConfigurations"), expression.Value(update.AwsProviderConfigurations))
	}
	if update.GcpProviderConfigurations != nil {
		updateBuilder = updateBuilder.Set(expression.Name("GcpProviderConfigurations"), expression.Value(update.GcpProviderConfigurations))
	}
	if update.Name != nil {
		updateBuilder = updateBuilder.Set(expression.Name("Name"), expression.Value(*update.Name))
	}
	if update.Description != nil {
		updateBuilder = updateBuilder.Set(expression.Name("Description"), expression.Value(*update.Description))
	}
	if update.ModuleVersionId != nil {
		updateBuilder = updateBuilder.Set(expression.Name("ModuleVersionId"), expression.Value(*update.ModuleVersionId))
	}
	if update.Status != nil {
		updateBuilder = updateBuilder.Set(expression.Name("Status"), expression.Value(*update.Status))
	}

	updateExpression, err := expression.NewBuilder().WithUpdate(updateBuilder).Build()
	if err != nil {
		return nil, err
	}

	updateInput := &dynamodb.UpdateItemInput{
		TableName: &c.moduleAssignmentsTableName,
		Key: map[string]types.AttributeValue{
			"ModuleAssignmentId": &types.AttributeValueMemberS{Value: moduleAssignmentId},
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
		return nil, helpers.NotFoundError{Message: fmt.Sprintf("Module Assignment %q not found", moduleAssignmentId)}
	default:
		if err != nil {
			return nil, err
		}
	}

	moduleAssignment := models.ModuleAssignment{}
	err = attributevalue.UnmarshalMap(result.Attributes, &moduleAssignment)
	if err != nil {
		return nil, err
	}

	return &moduleAssignment, nil
}
