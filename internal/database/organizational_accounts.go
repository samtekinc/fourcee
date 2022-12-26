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
	"github.com/sheacloud/tfom/pkg/models"
)

func (c *OrganizationsDatabaseClient) GetOrganizationalAccount(ctx context.Context, orgAccountId string) (*models.OrganizationalAccount, error) {
	response, err := c.dynamodb.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &c.accountsTableName,
		Key: map[string]types.AttributeValue{
			"OrgAccountId": &types.AttributeValueMemberS{Value: orgAccountId},
		},
	})
	if err != nil {
		return nil, err
	} else if response.Item == nil {
		return nil, helpers.NotFoundError{Message: fmt.Sprintf("Organizational Account %q not found", orgAccountId)}
	}

	orgAccount := models.OrganizationalAccount{}
	if err = attributevalue.UnmarshalMap(response.Item, &orgAccount); err != nil {
		return nil, err
	}

	return &orgAccount, nil
}

func (c OrganizationsDatabaseClient) GetOrganizationalAccounts(ctx context.Context, limit int32, cursor string) (*models.OrganizationalAccounts, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	scanInput := &dynamodb.ScanInput{
		TableName:         &c.accountsTableName,
		Limit:             &limit,
		ExclusiveStartKey: startKey,
	}

	resultItems, lastEvaluatedKey, err := helpers.ScanDynamoDBUntilLimit(ctx, c.dynamodb, scanInput, limit, []string{"OrgAccountId"})
	if err != nil {
		return nil, err
	}

	orgAccounts := []models.OrganizationalAccount{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &orgAccounts)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.OrganizationalAccounts{
		Items:      orgAccounts,
		NextCursor: nextCursor,
	}, nil
}

func (c *OrganizationsDatabaseClient) PutOrganizationalAccount(ctx context.Context, input *models.OrganizationalAccount) error {
	item, err := attributevalue.MarshalMap(input)
	if err != nil {
		return err
	}

	condition := expression.AttributeNotExists(expression.Name("OrgAccountId"))

	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return err
	}

	_, err = c.dynamodb.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:                 &c.accountsTableName,
		Item:                      item,
		ConditionExpression:       expr.Condition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})

	ccfe := &types.ConditionalCheckFailedException{}
	switch {
	case errors.As(err, &ccfe):
		return helpers.AlreadyExistsError{Message: fmt.Sprintf("Organizational Account %q already exists", input.OrgAccountId)}
	default:
		return err
	}
}

func (c *OrganizationsDatabaseClient) DeleteOrganizationalAccount(ctx context.Context, orgAccountId string) error {
	condition := expression.AttributeExists(expression.Name("OrgAccountId"))

	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return err
	}

	_, err = c.dynamodb.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: &c.accountsTableName,
		Key: map[string]types.AttributeValue{
			"OrgAccountId": &types.AttributeValueMemberS{Value: orgAccountId},
		},
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ConditionExpression:       expr.Condition(),
	})
	ccfe := &types.ConditionalCheckFailedException{}
	switch {
	case errors.As(err, &ccfe):
		return helpers.NotFoundError{Message: fmt.Sprintf("Organizational Account %q not found", orgAccountId)}
	default:
		return err
	}
}
