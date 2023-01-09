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

func (c *OrganizationsDatabaseClient) GetModulePropagationAssignment(ctx context.Context, modulePropagationId string, orgAccountId string) (*models.ModulePropagationAssignment, error) {
	response, err := c.dynamodb.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &c.modulePropagationAssignmentsTableName,
		Key: map[string]types.AttributeValue{
			"ModulePropagationId": &types.AttributeValueMemberS{Value: modulePropagationId},
			"OrgAccountId":        &types.AttributeValueMemberS{Value: orgAccountId},
		},
	})
	if err != nil {
		return nil, err
	} else if response.Item == nil {
		return nil, helpers.NotFoundError{Message: fmt.Sprintf("Module Propagation Assignment %q / %q not found", modulePropagationId, orgAccountId)}
	}

	modulePropagationAssignment := models.ModulePropagationAssignment{}
	if err = attributevalue.UnmarshalMap(response.Item, &modulePropagationAssignment); err != nil {
		return nil, err
	}

	return &modulePropagationAssignment, nil
}

func (c OrganizationsDatabaseClient) GetModulePropagationAssignments(ctx context.Context, limit int32, cursor string) (*models.ModulePropagationAssignments, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	scanInput := &dynamodb.ScanInput{
		TableName:         &c.modulePropagationAssignmentsTableName,
		Limit:             &limit,
		ExclusiveStartKey: startKey,
	}

	resultItems, lastEvaluatedKey, err := helpers.ScanDynamoDBUntilLimit(ctx, c.dynamodb, scanInput, limit, []string{"ModulePropagationId", "OrgAccountId"})
	if err != nil {
		return nil, err
	}

	modulePropagationAssignments := []models.ModulePropagationAssignment{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &modulePropagationAssignments)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.ModulePropagationAssignments{
		Items:      modulePropagationAssignments,
		NextCursor: nextCursor,
	}, nil
}

func (c OrganizationsDatabaseClient) GetModulePropagationAssignmentsByModulePropagationId(ctx context.Context, modulePropagationId string, limit int32, cursor string) (*models.ModulePropagationAssignments, error) {
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
		TableName:                 &c.modulePropagationAssignmentsTableName,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		Limit:                     &limit,
		ExclusiveStartKey:         startKey,
		ScanIndexForward:          aws.Bool(false),
	}

	resultItems, lastEvaluatedKey, err := helpers.QueryDynamoDBUntilLimit(ctx, c.dynamodb, queryInput, limit, []string{"ModulePropagationId", "OrgAccountId"})
	if err != nil {
		return nil, err
	}

	modulePropagationAssignments := []models.ModulePropagationAssignment{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &modulePropagationAssignments)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.ModulePropagationAssignments{
		Items:      modulePropagationAssignments,
		NextCursor: nextCursor,
	}, nil
}

func (c OrganizationsDatabaseClient) GetModulePropagationAssignmentsByOrgAccountId(ctx context.Context, modulePropagationId string, limit int32, cursor string) (*models.ModulePropagationAssignments, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	keyCondition := expression.Key("OrgAccountId").Equal(expression.Value(modulePropagationId))
	expressionBuilder := expression.NewBuilder().WithKeyCondition(keyCondition)
	expr, err := expressionBuilder.Build()
	if err != nil {
		return nil, err
	}

	queryInput := &dynamodb.QueryInput{
		TableName:                 &c.modulePropagationAssignmentsTableName,
		IndexName:                 aws.String("OrgAccountId-ModulePropagationId-index"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		Limit:                     &limit,
		ExclusiveStartKey:         startKey,
		ScanIndexForward:          aws.Bool(false),
	}

	resultItems, lastEvaluatedKey, err := helpers.QueryDynamoDBUntilLimit(ctx, c.dynamodb, queryInput, limit, []string{"ModulePropagationId", "OrgAccountId"})
	if err != nil {
		return nil, err
	}

	modulePropagationAssignments := []models.ModulePropagationAssignment{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &modulePropagationAssignments)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.ModulePropagationAssignments{
		Items:      modulePropagationAssignments,
		NextCursor: nextCursor,
	}, nil
}

func (c *OrganizationsDatabaseClient) PutModulePropagationAssignment(ctx context.Context, input *models.ModuleAssignment) (*models.ModulePropagationAssignment, *models.ModuleAssignment, error) {
	if input.ModulePropagationId == nil {
		return nil, nil, errors.New("modulePropagationId is required")
	}

	modulePropagationAssignment := models.ModulePropagationAssignment{
		ModulePropagationId: *input.ModulePropagationId,
		OrgAccountId:        input.OrgAccountId,
		ModuleAssignmentId:  input.ModuleAssignmentId,
	}
	modulePropagationAssignmentItem, err := attributevalue.MarshalMap(modulePropagationAssignment)
	if err != nil {
		return nil, nil, err
	}

	modulePropagationAssignmentCondition := expression.AttributeNotExists(expression.Name("ModulePropagationId"))

	modulePropagationAssignmentExpr, err := expression.NewBuilder().WithCondition(modulePropagationAssignmentCondition).Build()
	if err != nil {
		return nil, nil, err
	}

	modulePropagationAssignmentTransaction := types.TransactWriteItem{
		Put: &types.Put{
			TableName:                 &c.modulePropagationAssignmentsTableName,
			Item:                      modulePropagationAssignmentItem,
			ConditionExpression:       modulePropagationAssignmentExpr.Condition(),
			ExpressionAttributeNames:  modulePropagationAssignmentExpr.Names(),
			ExpressionAttributeValues: modulePropagationAssignmentExpr.Values(),
		},
	}

	moduleAssignmentItem, err := attributevalue.MarshalMap(input)
	if err != nil {
		return nil, nil, err
	}

	moduleAssignmentCondition := expression.AttributeNotExists(expression.Name("ModuleAssignmentId"))

	moduleAssignmentExpr, err := expression.NewBuilder().WithCondition(moduleAssignmentCondition).Build()
	if err != nil {
		return nil, nil, err
	}

	moduleAssignmentTransaction := types.TransactWriteItem{
		Put: &types.Put{
			TableName:                 &c.moduleAssignmentsTableName,
			Item:                      moduleAssignmentItem,
			ConditionExpression:       moduleAssignmentExpr.Condition(),
			ExpressionAttributeNames:  moduleAssignmentExpr.Names(),
			ExpressionAttributeValues: moduleAssignmentExpr.Values(),
		},
	}

	_, err = c.dynamodb.TransactWriteItems(ctx, &dynamodb.TransactWriteItemsInput{
		TransactItems: []types.TransactWriteItem{moduleAssignmentTransaction, modulePropagationAssignmentTransaction},
	})

	ccfe := &types.ConditionalCheckFailedException{}
	switch {
	case errors.As(err, &ccfe):
		return nil, nil, helpers.AlreadyExistsError{Message: fmt.Sprintf("Module Assignment %q already exists", input.ModuleAssignmentId)}
	default:
		return &modulePropagationAssignment, input, err
	}
}
