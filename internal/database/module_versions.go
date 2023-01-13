package database

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/sheacloud/tfom/internal/helpers"
	"github.com/sheacloud/tfom/pkg/models"
)

func (c *DatabaseClient) GetModuleVersion(ctx context.Context, moduleGroupId string, moduleVersionId string) (*models.ModuleVersion, error) {
	response, err := c.dynamodb.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &c.versionsTableName,
		Key: map[string]types.AttributeValue{
			"ModuleGroupId":   &types.AttributeValueMemberS{Value: moduleGroupId},
			"ModuleVersionId": &types.AttributeValueMemberS{Value: moduleVersionId},
		},
	})
	if err != nil {
		return nil, err
	} else if response.Item == nil {
		return nil, helpers.NotFoundError{Message: fmt.Sprintf("Module Version %q not found", moduleVersionId)}
	}

	moduleVersion := models.ModuleVersion{}
	if err = attributevalue.UnmarshalMap(response.Item, &moduleVersion); err != nil {
		return nil, err
	}

	return &moduleVersion, nil
}

func (c *DatabaseClient) GetModuleVersionsByIds(ctx context.Context, ids []string) ([]models.ModuleVersion, error) {
	fmt.Println("getting module versions by ids", len(ids))
	var keys []map[string]types.AttributeValue

	for _, id := range ids {
		parts := strings.Split(id, ":")
		keys = append(keys, map[string]types.AttributeValue{
			"ModuleGroupId":   &types.AttributeValueMemberS{Value: parts[0]},
			"ModuleVersionId": &types.AttributeValueMemberS{Value: parts[1]},
		})
	}

	bii := dynamodb.BatchGetItemInput{
		RequestItems: map[string]types.KeysAndAttributes{
			c.versionsTableName: {
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
			items = append(items, bgo.Responses[c.versionsTableName]...)
		}
		requestItems := bgo.UnprocessedKeys
		bii = dynamodb.BatchGetItemInput{RequestItems: requestItems}
		if len(requestItems) == 0 {
			break
		}
	}

	items = SortDynamoDBBatchResponses(keys, items)

	results := []models.ModuleVersion{}
	err := attributevalue.UnmarshalListOfMaps(items, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (c DatabaseClient) GetModuleVersions(ctx context.Context, moduleGroupId string, limit int32, cursor string) (*models.ModuleVersions, error) {
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
		TableName:                 &c.versionsTableName,
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

	moduleVersions := []models.ModuleVersion{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &moduleVersions)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.ModuleVersions{
		Items:      moduleVersions,
		NextCursor: nextCursor,
	}, nil
}

func (c *DatabaseClient) PutModuleVersion(ctx context.Context, input *models.ModuleVersion) error {
	item, err := attributevalue.MarshalMap(input)
	if err != nil {
		return err
	}

	condition := expression.AttributeNotExists(expression.Name("ModuleVersionId"))

	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return err
	}

	_, err = c.dynamodb.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:                 &c.versionsTableName,
		Item:                      item,
		ConditionExpression:       expr.Condition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})

	ccfe := &types.ConditionalCheckFailedException{}
	switch {
	case errors.As(err, &ccfe):
		return helpers.AlreadyExistsError{Message: fmt.Sprintf("Module Version %q already exists", input.ModuleVersionId)}
	default:
		return err
	}
}

func (c *DatabaseClient) DeleteModuleVersion(ctx context.Context, moduleGroupId string, moduleVersionId string) error {
	condition := expression.AttributeExists(expression.Name("ModuleVersionId"))

	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return err
	}

	_, err = c.dynamodb.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: &c.versionsTableName,
		Key: map[string]types.AttributeValue{
			"ModuleGroupId":   &types.AttributeValueMemberS{Value: moduleGroupId},
			"ModuleVersionId": &types.AttributeValueMemberS{Value: moduleVersionId},
		},
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ConditionExpression:       expr.Condition(),
	})
	ccfe := &types.ConditionalCheckFailedException{}
	switch {
	case errors.As(err, &ccfe):
		return helpers.NotFoundError{Message: fmt.Sprintf("Module Version %q not found", moduleVersionId)}
	default:
		return err
	}
}
