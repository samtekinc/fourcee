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
	"github.com/sheacloud/tfom/pkg/organizations/models"
)

type OrganizationalUnitUpdate struct {
	Name            *string
	ParentOrgUnitId *string
	Hierarchy       *string
}

func (c *OrganizationsDatabaseClient) GetOrganizationalUnit(ctx context.Context, orgUnitId string) (*models.OrganizationalUnit, error) {
	response, err := c.dynamodb.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &c.unitsTableName,
		Key: map[string]types.AttributeValue{
			"OrgUnitId": &types.AttributeValueMemberS{Value: orgUnitId},
		},
	})
	if err != nil {
		return nil, err
	} else if response.Item == nil {
		return nil, helpers.NotFoundError{Message: fmt.Sprintf("Organizational Unit %q not found", orgUnitId)}
	}

	orgUnit := models.OrganizationalUnit{}
	if err = attributevalue.UnmarshalMap(response.Item, &orgUnit); err != nil {
		return nil, err
	}

	return &orgUnit, nil
}

func (c OrganizationsDatabaseClient) GetOrganizationalUnits(ctx context.Context, limit int32, cursor string) (*models.OrganizationalUnits, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	scanInput := &dynamodb.ScanInput{
		TableName:         &c.unitsTableName,
		Limit:             &limit,
		ExclusiveStartKey: startKey,
	}

	resultItems, lastEvaluatedKey, err := helpers.ScanDynamoDBUntilLimit(ctx, c.dynamodb, scanInput, limit, []string{"OrgUnitId"})
	if err != nil {
		return nil, err
	}

	orgUnits := []models.OrganizationalUnit{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &orgUnits)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.OrganizationalUnits{
		Items:      orgUnits,
		NextCursor: nextCursor,
	}, nil
}

func (c OrganizationsDatabaseClient) GetOrganizationalUnitsByDimension(ctx context.Context, dimensionId string, limit int32, cursor string) (*models.OrganizationalUnits, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	keyCondition := expression.Key("DimensionId").Equal(expression.Value(dimensionId))
	expressionBuilder := expression.NewBuilder().WithKeyCondition(keyCondition)
	expr, err := expressionBuilder.Build()
	if err != nil {
		return nil, err
	}

	queryInput := &dynamodb.QueryInput{
		TableName:                 &c.unitsTableName,
		IndexName:                 aws.String("DimensionId-ParentOrgUnitId-index"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		Limit:                     &limit,
		ExclusiveStartKey:         startKey,
	}

	resultItems, lastEvaluatedKey, err := helpers.QueryDynamoDBUntilLimit(ctx, c.dynamodb, queryInput, limit, []string{"Id"})
	if err != nil {
		return nil, err
	}

	orgUnits := []models.OrganizationalUnit{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &orgUnits)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.OrganizationalUnits{
		Items:      orgUnits,
		NextCursor: nextCursor,
	}, nil
}

func (c OrganizationsDatabaseClient) GetOrganizationalUnitsByParent(ctx context.Context, dimensionId string, parentOrgUnitId string, limit int32, cursor string) (*models.OrganizationalUnits, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	keyCondition := expression.Key("DimensionId").Equal(expression.Value(dimensionId)).And(expression.Key("ParentOrgUnitId").Equal(expression.Value(parentOrgUnitId)))
	expressionBuilder := expression.NewBuilder().WithKeyCondition(keyCondition)
	expr, err := expressionBuilder.Build()
	if err != nil {
		return nil, err
	}

	queryInput := &dynamodb.QueryInput{
		TableName:                 &c.unitsTableName,
		IndexName:                 aws.String("DimensionId-ParentOrgUnitId-index"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		Limit:                     &limit,
		ExclusiveStartKey:         startKey,
	}

	resultItems, lastEvaluatedKey, err := helpers.QueryDynamoDBUntilLimit(ctx, c.dynamodb, queryInput, limit, []string{"Id"})
	if err != nil {
		return nil, err
	}

	orgUnits := []models.OrganizationalUnit{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &orgUnits)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.OrganizationalUnits{
		Items:      orgUnits,
		NextCursor: nextCursor,
	}, nil
}

func (c OrganizationsDatabaseClient) GetOrganizationalUnitsByHierarchy(ctx context.Context, dimensionId string, hierarchy string, limit int32, cursor string) (*models.OrganizationalUnits, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	keyCondition := expression.Key("DimensionId").Equal(expression.Value(dimensionId)).And(expression.Key("Hierarchy").BeginsWith(hierarchy))
	expressionBuilder := expression.NewBuilder().WithKeyCondition(keyCondition)
	expr, err := expressionBuilder.Build()
	if err != nil {
		return nil, err
	}

	queryInput := &dynamodb.QueryInput{
		TableName:                 &c.unitsTableName,
		IndexName:                 aws.String("DimensionId-Hierarchy-index"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		Limit:                     &limit,
		ExclusiveStartKey:         startKey,
	}

	resultItems, lastEvaluatedKey, err := helpers.QueryDynamoDBUntilLimit(ctx, c.dynamodb, queryInput, limit, []string{"Id"})
	if err != nil {
		return nil, err
	}

	orgUnits := []models.OrganizationalUnit{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &orgUnits)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.OrganizationalUnits{
		Items:      orgUnits,
		NextCursor: nextCursor,
	}, nil
}

func (c *OrganizationsDatabaseClient) PutOrganizationalUnit(ctx context.Context, input *models.OrganizationalUnit) error {
	item, err := attributevalue.MarshalMap(input)
	if err != nil {
		return err
	}

	condition := expression.AttributeNotExists(expression.Name("OrgUnitId"))

	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return err
	}

	_, err = c.dynamodb.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:                 &c.unitsTableName,
		Item:                      item,
		ConditionExpression:       expr.Condition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})

	ccfe := &types.ConditionalCheckFailedException{}
	switch {
	case errors.As(err, &ccfe):
		return helpers.AlreadyExistsError{Message: fmt.Sprintf("Organizational Unit %q already exists", input.OrgUnitId)}
	default:
		return err
	}
}

func (c *OrganizationsDatabaseClient) DeleteOrganizationalUnit(ctx context.Context, orgUnitId string) error {
	condition := expression.AttributeExists(expression.Name("OrgUnitId"))

	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return err
	}

	_, err = c.dynamodb.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: &c.unitsTableName,
		Key: map[string]types.AttributeValue{
			"OrgUnitId": &types.AttributeValueMemberS{Value: orgUnitId},
		},
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ConditionExpression:       expr.Condition(),
	})
	ccfe := &types.ConditionalCheckFailedException{}
	switch {
	case errors.As(err, &ccfe):
		return helpers.NotFoundError{Message: fmt.Sprintf("Organizational Unit %q not found", orgUnitId)}
	default:
		return err
	}
}

func (c *OrganizationsDatabaseClient) UpdateOrganizationalUnit(ctx context.Context, orgUnitId string, update *OrganizationalUnitUpdate) (*models.OrganizationalUnit, error) {
	condition := expression.AttributeExists(expression.Name("Id"))

	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return nil, err
	}

	updateBuilder := expression.UpdateBuilder{}
	if update.Name != nil {
		updateBuilder = updateBuilder.Set(expression.Name("Name"), expression.Value(*update.Name))
	}
	if update.ParentOrgUnitId != nil {
		updateBuilder = updateBuilder.Set(expression.Name("ParentOrgUnitId"), expression.Value(*update.ParentOrgUnitId))
	}
	if update.Hierarchy != nil {
		updateBuilder = updateBuilder.Set(expression.Name("Hierarchy"), expression.Value(*update.Hierarchy))
	}

	updateExpression, err := expression.NewBuilder().WithUpdate(updateBuilder).Build()
	if err != nil {
		return nil, err
	}

	updateInput := &dynamodb.UpdateItemInput{
		TableName:                 &c.unitsTableName,
		Key:                       map[string]types.AttributeValue{"Id": &types.AttributeValueMemberS{Value: orgUnitId}},
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
		return nil, helpers.NotFoundError{Message: fmt.Sprintf("Organizational Unit %q not found", orgUnitId)}
	default:
		if err != nil {
			return nil, err
		}
	}

	orgUnit := models.OrganizationalUnit{}
	err = attributevalue.UnmarshalMap(result.Attributes, &orgUnit)
	if err != nil {
		return nil, err
	}

	return &orgUnit, nil
}
