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
	"go.uber.org/zap"
)

func (c *DatabaseClient) GetOrganizationalDimension(ctx context.Context, orgDimensionId string) (*models.OrganizationalDimension, error) {
	zap.L().Sugar().Debugw("GetOrganizationalDimension", "orgDimensionId", orgDimensionId)
	response, err := c.dynamodb.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &c.dimensionsTableName,
		Key: map[string]types.AttributeValue{
			"OrgDimensionId": &types.AttributeValueMemberS{Value: orgDimensionId},
		},
	})
	if err != nil {
		return nil, err
	} else if response.Item == nil {
		return nil, helpers.NotFoundError{Message: fmt.Sprintf("Organizational Dimension %q not found", orgDimensionId)}
	}

	orgDimension := models.OrganizationalDimension{}
	if err = attributevalue.UnmarshalMap(response.Item, &orgDimension); err != nil {
		return nil, err
	}

	return &orgDimension, nil
}

func (c *DatabaseClient) GetOrganizationalDimensionsByIds(ctx context.Context, orgDimensionIds []string) ([]models.OrganizationalDimension, error) {
	zap.L().Sugar().Debugw("GetOrganizationalDimensionsByIds", "orgDimensionIds", orgDimensionIds)
	var keys []map[string]types.AttributeValue

	for _, id := range orgDimensionIds {
		keys = append(keys, map[string]types.AttributeValue{
			"OrgDimensionId": &types.AttributeValueMemberS{Value: id},
		})
	}

	bii := dynamodb.BatchGetItemInput{
		RequestItems: map[string]types.KeysAndAttributes{
			c.dimensionsTableName: {
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
			items = append(items, bgo.Responses[c.dimensionsTableName]...)
		}
		requestItems := bgo.UnprocessedKeys
		bii = dynamodb.BatchGetItemInput{RequestItems: requestItems}
		if len(requestItems) == 0 {
			break
		}
	}

	items = SortDynamoDBBatchResponses(keys, items)

	results := []models.OrganizationalDimension{}
	err := attributevalue.UnmarshalListOfMaps(items, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (c DatabaseClient) GetOrganizationalDimensions(ctx context.Context, limit int32, cursor string) (*models.OrganizationalDimensions, error) {
	zap.L().Sugar().Debugw("GetOrganizationalDimensions", "limit", limit, "cursor", cursor)
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	scanInput := &dynamodb.ScanInput{
		TableName:         &c.dimensionsTableName,
		Limit:             &limit,
		ExclusiveStartKey: startKey,
	}

	resultItems, lastEvaluatedKey, err := helpers.ScanDynamoDBUntilLimit(ctx, c.dynamodb, scanInput, limit, []string{"OrgDimensionId"})
	if err != nil {
		return nil, err
	}

	orgDimensions := []models.OrganizationalDimension{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &orgDimensions)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.OrganizationalDimensions{
		Items:      orgDimensions,
		NextCursor: nextCursor,
	}, nil
}

func (c *DatabaseClient) PutOrganizationalDimension(ctx context.Context, input *models.OrganizationalDimension) error {
	zap.L().Sugar().Debugw("PutOrganizationalDimension", "input", input)
	item, err := attributevalue.MarshalMap(input)
	if err != nil {
		return err
	}

	condition := expression.AttributeNotExists(expression.Name("OrgDimensionId"))

	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return err
	}

	_, err = c.dynamodb.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:                 &c.dimensionsTableName,
		Item:                      item,
		ConditionExpression:       expr.Condition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})

	ccfe := &types.ConditionalCheckFailedException{}
	switch {
	case errors.As(err, &ccfe):
		return helpers.AlreadyExistsError{Message: fmt.Sprintf("Organizational Dimension %q already exists", input.OrgDimensionId)}
	default:
		return err
	}
}

func (c *DatabaseClient) DeleteOrganizationalDimension(ctx context.Context, orgDimensionId string) error {
	zap.L().Sugar().Debugw("DeleteOrganizationalDimension", "orgDimensionId", orgDimensionId)
	condition := expression.AttributeExists(expression.Name("OrgDimensionId"))

	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return err
	}

	_, err = c.dynamodb.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: &c.dimensionsTableName,
		Key: map[string]types.AttributeValue{
			"OrgDimensionId": &types.AttributeValueMemberS{Value: orgDimensionId},
		},
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ConditionExpression:       expr.Condition(),
	})
	ccfe := &types.ConditionalCheckFailedException{}
	switch {
	case errors.As(err, &ccfe):
		return helpers.NotFoundError{Message: fmt.Sprintf("Organizational Dimension %q not found", orgDimensionId)}
	default:
		return err
	}
}
