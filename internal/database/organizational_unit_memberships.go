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

func (c OrganizationsDatabaseClient) GetOrganizationalUnitMembershipsByAccount(ctx context.Context, accountId string, limit int32, cursor string) (*models.OrganizationalUnitMemberships, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	keyCondition := expression.Key("OrgAccountId").Equal(expression.Value(accountId))
	expressionBuilder := expression.NewBuilder().WithKeyCondition(keyCondition)
	expr, err := expressionBuilder.Build()
	if err != nil {
		return nil, err
	}

	queryInput := &dynamodb.QueryInput{
		TableName:                 &c.membershipsTableName,
		IndexName:                 aws.String("OrgAccountId-OrgUnitId-index"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		Limit:                     &limit,
		ExclusiveStartKey:         startKey,
	}

	resultItems, lastEvaluatedKey, err := helpers.QueryDynamoDBUntilLimit(ctx, c.dynamodb, queryInput, limit, []string{"OrgAccountId", "OrgUnitId"})
	if err != nil {
		return nil, err
	}

	orgUnitMemberships := []models.OrganizationalUnitMembership{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &orgUnitMemberships)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.OrganizationalUnitMemberships{
		Items:      orgUnitMemberships,
		NextCursor: nextCursor,
	}, nil
}

func (c OrganizationsDatabaseClient) GetOrganizationalUnitMembershipsByOrgUnit(ctx context.Context, orgUnitId string, limit int32, cursor string) (*models.OrganizationalUnitMemberships, error) {
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
		TableName:                 &c.membershipsTableName,
		IndexName:                 aws.String("OrgUnitId-OrgAccountId-index"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		Limit:                     &limit,
		ExclusiveStartKey:         startKey,
	}

	resultItems, lastEvaluatedKey, err := helpers.QueryDynamoDBUntilLimit(ctx, c.dynamodb, queryInput, limit, []string{"OrgUnitId", "OrgAccountId"})
	if err != nil {
		return nil, err
	}

	orgUnitMemberships := []models.OrganizationalUnitMembership{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &orgUnitMemberships)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.OrganizationalUnitMemberships{
		Items:      orgUnitMemberships,
		NextCursor: nextCursor,
	}, nil
}

func (c OrganizationsDatabaseClient) GetOrganizationalUnitMembershipsByDimension(ctx context.Context, orgDimensionId string, limit int32, cursor string) (*models.OrganizationalUnitMemberships, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	keyCondition := expression.Key("OrgDimensionId").Equal(expression.Value(orgDimensionId))
	expressionBuilder := expression.NewBuilder().WithKeyCondition(keyCondition)
	expr, err := expressionBuilder.Build()
	if err != nil {
		return nil, err
	}

	queryInput := &dynamodb.QueryInput{
		TableName:                 &c.membershipsTableName,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		Limit:                     &limit,
		ExclusiveStartKey:         startKey,
	}

	resultItems, lastEvaluatedKey, err := helpers.QueryDynamoDBUntilLimit(ctx, c.dynamodb, queryInput, limit, []string{"OrgDimensionId", "OrgAccountId"})
	if err != nil {
		return nil, err
	}

	orgUnitMemberships := []models.OrganizationalUnitMembership{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &orgUnitMemberships)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.OrganizationalUnitMemberships{
		Items:      orgUnitMemberships,
		NextCursor: nextCursor,
	}, nil
}

func (c *OrganizationsDatabaseClient) PutOrganizationalUnitMembership(ctx context.Context, input *models.OrganizationalUnitMembership) error {
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
		TableName:                 &c.membershipsTableName,
		Item:                      item,
		ConditionExpression:       expr.Condition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})

	ccfe := &types.ConditionalCheckFailedException{}
	switch {
	case errors.As(err, &ccfe):
		return helpers.AlreadyExistsError{Message: fmt.Sprintf("Organizational UnitMembership between %q and %q already exists", input.OrgDimensionId, input.OrgAccountId)}
	default:
		return err
	}
}

func (c *OrganizationsDatabaseClient) DeleteOrganizationalUnitMembership(ctx context.Context, orgDimensionId string, accountId string) error {
	condition := expression.AttributeExists(expression.Name("OrgDimensionId"))

	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return err
	}

	_, err = c.dynamodb.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: &c.membershipsTableName,
		Key: map[string]types.AttributeValue{
			"OrgDimensionId": &types.AttributeValueMemberS{Value: orgDimensionId},
			"OrgAccountId":   &types.AttributeValueMemberS{Value: accountId},
		},
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ConditionExpression:       expr.Condition(),
	})
	ccfe := &types.ConditionalCheckFailedException{}
	switch {
	case errors.As(err, &ccfe):
		return helpers.NotFoundError{Message: fmt.Sprintf("Organizational UnitMembership between %q and %q not found", orgDimensionId, accountId)}
	default:
		return err
	}
}
