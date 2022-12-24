package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/sheacloud/tfom/internal/helpers"
	"github.com/sheacloud/tfom/pkg/organizations/models"
)

func (c *OrganizationsDatabaseClient) GetModuleGroup(ctx context.Context, moduleGroupId string) (*models.ModuleGroup, error) {
	response, err := c.dynamodb.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &c.groupsTableName,
		Key: map[string]types.AttributeValue{
			"ModuleGroupId": &types.AttributeValueMemberS{Value: moduleGroupId},
		},
	})
	if err != nil {
		return nil, err
	} else if response.Item == nil {
		return nil, helpers.NotFoundError{Message: fmt.Sprintf("Module Group %q not found", moduleGroupId)}
	}

	moduleGroup := models.ModuleGroup{}
	if err = attributevalue.UnmarshalMap(response.Item, &moduleGroup); err != nil {
		return nil, err
	}

	return &moduleGroup, nil
}

func (c OrganizationsDatabaseClient) GetModuleGroups(ctx context.Context, limit int32, cursor string) (*models.ModuleGroups, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	scanInput := &dynamodb.ScanInput{
		TableName:         &c.groupsTableName,
		Limit:             &limit,
		ExclusiveStartKey: startKey,
	}

	resultItems, lastEvaluatedKey, err := helpers.ScanDynamoDBUntilLimit(ctx, c.dynamodb, scanInput, limit, []string{"ModuleGroupId"})
	if err != nil {
		return nil, err
	}

	moduleGroups := []models.ModuleGroup{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &moduleGroups)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.ModuleGroups{
		Items:      moduleGroups,
		NextCursor: nextCursor,
	}, nil
}

func (c *OrganizationsDatabaseClient) PutModuleGroup(ctx context.Context, input *models.ModuleGroup) error {
	item, err := attributevalue.MarshalMap(input)
	if err != nil {
		return err
	}

	condition := expression.AttributeNotExists(expression.Name("ModuleGroupId"))

	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return err
	}

	_, err = c.dynamodb.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:                 &c.groupsTableName,
		Item:                      item,
		ConditionExpression:       expr.Condition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})

	ccfe := &types.ConditionalCheckFailedException{}
	switch {
	case errors.As(err, &ccfe):
		return helpers.AlreadyExistsError{Message: fmt.Sprintf("Module Group %q already exists", input.ModuleGroupId)}
	default:
		return err
	}
}

func (c *OrganizationsDatabaseClient) DeleteModuleGroup(ctx context.Context, moduleGroupId string) error {
	condition := expression.AttributeExists(expression.Name("ModuleGroupId"))

	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return err
	}

	_, err = c.dynamodb.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: &c.groupsTableName,
		Key: map[string]types.AttributeValue{
			"ModuleGroupId": &types.AttributeValueMemberS{Value: moduleGroupId},
		},
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ConditionExpression:       expr.Condition(),
	})
	ccfe := &types.ConditionalCheckFailedException{}
	switch {
	case errors.As(err, &ccfe):
		return helpers.NotFoundError{Message: fmt.Sprintf("Module Group %q not found", moduleGroupId)}
	default:
		return err
	}
}
