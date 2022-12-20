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

func (c *OrganizationsDatabaseClient) GetOrganizationalDimension(ctx context.Context, dimensionId string) (*models.OrganizationalDimension, error) {
	response, err := c.dynamodb.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &c.dimensionsTableName,
		Key: map[string]types.AttributeValue{
			"DimensionId": &types.AttributeValueMemberS{Value: dimensionId},
		},
	})
	if err != nil {
		return nil, err
	} else if response.Item == nil {
		return nil, helpers.NotFoundError{Message: fmt.Sprintf("Organizational Dimension %q not found", dimensionId)}
	}

	orgDimension := models.OrganizationalDimension{}
	if err = attributevalue.UnmarshalMap(response.Item, &orgDimension); err != nil {
		return nil, err
	}

	return &orgDimension, nil
}

func (c OrganizationsDatabaseClient) GetOrganizationalDimensions(ctx context.Context, limit int32, cursor string) (*models.OrganizationalDimensions, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	scanInput := &dynamodb.ScanInput{
		TableName:         &c.dimensionsTableName,
		Limit:             &limit,
		ExclusiveStartKey: startKey,
	}

	resultItems, lastEvaluatedKey, err := helpers.ScanDynamoDBUntilLimit(ctx, c.dynamodb, scanInput, limit, []string{"DimensionId"})
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

func (c *OrganizationsDatabaseClient) PutOrganizationalDimension(ctx context.Context, input *models.OrganizationalDimension) error {
	item, err := attributevalue.MarshalMap(input)
	if err != nil {
		return err
	}

	condition := expression.AttributeNotExists(expression.Name("DimensionId"))

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
		return helpers.AlreadyExistsError{Message: fmt.Sprintf("Organizational Dimension %q already exists", input.DimensionId)}
	default:
		return err
	}
}

func (c *OrganizationsDatabaseClient) DeleteOrganizationalDimension(ctx context.Context, dimensionId string) error {
	condition := expression.AttributeExists(expression.Name("DimensionId"))

	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return err
	}

	_, err = c.dynamodb.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: &c.dimensionsTableName,
		Key: map[string]types.AttributeValue{
			"DimensionId": &types.AttributeValueMemberS{Value: dimensionId},
		},
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ConditionExpression:       expr.Condition(),
	})
	ccfe := &types.ConditionalCheckFailedException{}
	switch {
	case errors.As(err, &ccfe):
		return helpers.NotFoundError{Message: fmt.Sprintf("Organizational Dimension %q not found", dimensionId)}
	default:
		return err
	}
}
