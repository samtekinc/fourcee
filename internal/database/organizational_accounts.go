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

func (c *DatabaseClient) GetOrganizationalAccount(ctx context.Context, orgAccountId string) (*models.OrganizationalAccount, error) {
	fmt.Println("getting org account")
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

func (c *DatabaseClient) GetOrganizationalAccountsByIds(ctx context.Context, ids []string) ([]models.OrganizationalAccount, error) {
	fmt.Println("getting org accounts by ids", len(ids))
	var keys []map[string]types.AttributeValue

	for _, id := range ids {
		keys = append(keys, map[string]types.AttributeValue{
			"OrgAccountId": &types.AttributeValueMemberS{Value: id},
		})
	}

	bii := dynamodb.BatchGetItemInput{
		RequestItems: map[string]types.KeysAndAttributes{
			c.accountsTableName: {
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
			items = append(items, bgo.Responses[c.accountsTableName]...)
		}
		requestItems := bgo.UnprocessedKeys
		bii = dynamodb.BatchGetItemInput{RequestItems: requestItems}
		if len(requestItems) == 0 {
			break
		}
	}

	items = SortDynamoDBBatchResponses(keys, items)

	results := []models.OrganizationalAccount{}
	err := attributevalue.UnmarshalListOfMaps(items, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (c DatabaseClient) GetOrganizationalAccounts(ctx context.Context, limit int32, cursor string) (*models.OrganizationalAccounts, error) {
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

func (c *DatabaseClient) PutOrganizationalAccount(ctx context.Context, input *models.OrganizationalAccount) error {
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

func (c *DatabaseClient) DeleteOrganizationalAccount(ctx context.Context, orgAccountId string) error {
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

func (c *DatabaseClient) UpdateOrganizationalAccount(ctx context.Context, orgAccountId string, update *models.OrganizationalAccountUpdate) (*models.OrganizationalAccount, error) {
	condition := expression.AttributeExists(expression.Name("OrgAccountId"))

	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return nil, err
	}

	updateBuilder := expression.UpdateBuilder{}
	if update.Metadata != nil {
		updateBuilder = updateBuilder.Set(expression.Name("Metadata"), expression.Value(update.Metadata))
	}

	updateExpression, err := expression.NewBuilder().WithUpdate(updateBuilder).Build()
	if err != nil {
		return nil, err
	}

	updateInput := &dynamodb.UpdateItemInput{
		TableName: &c.accountsTableName,
		Key: map[string]types.AttributeValue{
			"OrgAccountId": &types.AttributeValueMemberS{Value: orgAccountId},
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
		return nil, helpers.NotFoundError{Message: fmt.Sprintf("Org Account %q not found", orgAccountId)}
	default:
		if err != nil {
			return nil, err
		}
	}

	orgAccount := models.OrganizationalAccount{}
	err = attributevalue.UnmarshalMap(result.Attributes, &orgAccount)
	if err != nil {
		return nil, err
	}

	return &orgAccount, nil
}
