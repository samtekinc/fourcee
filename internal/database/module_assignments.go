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

func (c *OrganizationsDatabaseClient) GetModuleAssignment(ctx context.Context, moduleAssignmentId string) (*models.ModuleAssignment, error) {
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

func (c OrganizationsDatabaseClient) GetModuleAssignments(ctx context.Context, limit int32, cursor string) (*models.ModuleAssignments, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	scanInput := &dynamodb.ScanInput{
		TableName:         &c.moduleAssignmentsTableName,
		Limit:             &limit,
		ExclusiveStartKey: startKey,
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

func (c OrganizationsDatabaseClient) GetModuleAssignmentsByModulePropagationId(ctx context.Context, modulePropagationId string, limit int32, cursor string) (*models.ModuleAssignments, error) {
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

func (c OrganizationsDatabaseClient) GetModuleAssignmentsByOrgAccountId(ctx context.Context, orgAccountId string, limit int32, cursor string) (*models.ModuleAssignments, error) {
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
		IndexName:                 aws.String("OrgAccountId-ModulePropagationId-index"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		Limit:                     &limit,
		ExclusiveStartKey:         startKey,
		ScanIndexForward:          aws.Bool(false),
	}

	resultItems, lastEvaluatedKey, err := helpers.QueryDynamoDBUntilLimit(ctx, c.dynamodb, queryInput, limit, []string{"ModuleAssignmentId", "OrgAccountId", "ModulePropagationId"})
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

func (c OrganizationsDatabaseClient) GetModuleAssignmentsByModuleVersionId(ctx context.Context, moduleVersionId string, limit int32, cursor string) (*models.ModuleAssignments, error) {
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

func (c OrganizationsDatabaseClient) GetModuleAssignmentsByModuleGroupId(ctx context.Context, moduleGroupId string, limit int32, cursor string) (*models.ModuleAssignments, error) {
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

func (c *OrganizationsDatabaseClient) PutModuleAssignment(ctx context.Context, input *models.ModuleAssignment) error {
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

func (c *OrganizationsDatabaseClient) UpdateModuleAssignment(ctx context.Context, moduleAssignmentId string, update *models.ModuleAssignmentUpdate) (*models.ModuleAssignment, error) {
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
