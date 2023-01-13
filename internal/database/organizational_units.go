package database

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/sheacloud/tfom/internal/helpers"
	"github.com/sheacloud/tfom/pkg/models"
)

type OrganizationalUnitUpdate struct {
	Name            *string
	ParentOrgUnitId *string
	Hierarchy       *string
}

func (c *DatabaseClient) GetOrganizationalUnit(ctx context.Context, orgDimensionId string, orgUnitId string) (*models.OrganizationalUnit, error) {
	response, err := c.dynamodb.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &c.unitsTableName,
		Key: map[string]types.AttributeValue{
			"OrgDimensionId": &types.AttributeValueMemberS{Value: orgDimensionId},
			"OrgUnitId":      &types.AttributeValueMemberS{Value: orgUnitId},
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

func (c *DatabaseClient) GetOrganizationalUnitsByIds(ctx context.Context, ids []string) ([]models.OrganizationalUnit, error) {
	fmt.Println("getting org units by ids", len(ids))
	var keys []map[string]types.AttributeValue

	for _, id := range ids {
		parts := strings.Split(id, ":")
		keys = append(keys, map[string]types.AttributeValue{
			"OrgDimensionId": &types.AttributeValueMemberS{Value: parts[0]},
			"OrgUnitId":      &types.AttributeValueMemberS{Value: parts[1]},
		})
	}

	bii := dynamodb.BatchGetItemInput{
		RequestItems: map[string]types.KeysAndAttributes{
			c.unitsTableName: {
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
			items = append(items, bgo.Responses[c.unitsTableName]...)
		}
		requestItems := bgo.UnprocessedKeys
		bii = dynamodb.BatchGetItemInput{RequestItems: requestItems}
		if len(requestItems) == 0 {
			break
		}
	}

	items = SortDynamoDBBatchResponses(keys, items)

	results := []models.OrganizationalUnit{}
	err := attributevalue.UnmarshalListOfMaps(items, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (c DatabaseClient) GetOrganizationalUnits(ctx context.Context, limit int32, cursor string) (*models.OrganizationalUnits, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	scanInput := &dynamodb.ScanInput{
		TableName:         &c.unitsTableName,
		Limit:             &limit,
		ExclusiveStartKey: startKey,
	}

	resultItems, lastEvaluatedKey, err := helpers.ScanDynamoDBUntilLimit(ctx, c.dynamodb, scanInput, limit, []string{"OrgDimensionId", "OrgUnitId"})
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

func (c DatabaseClient) GetOrganizationalUnitsByDimension(ctx context.Context, orgDimensionId string, limit int32, cursor string) (*models.OrganizationalUnits, error) {
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
		TableName:                 &c.unitsTableName,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		Limit:                     &limit,
		ExclusiveStartKey:         startKey,
	}

	resultItems, lastEvaluatedKey, err := helpers.QueryDynamoDBUntilLimit(ctx, c.dynamodb, queryInput, limit, []string{"OrgDimensionId", "OrgUnitId"})
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

func (c DatabaseClient) GetOrganizationalUnitsByParent(ctx context.Context, orgDimensionId string, parentOrgUnitId string, limit int32, cursor string) (*models.OrganizationalUnits, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	keyCondition := expression.Key("OrgDimensionId").Equal(expression.Value(orgDimensionId)).And(expression.Key("ParentOrgUnitId").Equal(expression.Value(parentOrgUnitId)))
	expressionBuilder := expression.NewBuilder().WithKeyCondition(keyCondition)
	expr, err := expressionBuilder.Build()
	if err != nil {
		return nil, err
	}

	queryInput := &dynamodb.QueryInput{
		TableName:                 &c.unitsTableName,
		IndexName:                 aws.String("OrgDimensionId-ParentOrgUnitId-index"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		Limit:                     &limit,
		ExclusiveStartKey:         startKey,
	}

	resultItems, lastEvaluatedKey, err := helpers.QueryDynamoDBUntilLimit(ctx, c.dynamodb, queryInput, limit, []string{"OrgDimensionId", "OrgUnitId", "ParentOrgUnitId"})
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

func (c DatabaseClient) GetOrganizationalUnitsByHierarchy(ctx context.Context, orgDimensionId string, hierarchy string, limit int32, cursor string) (*models.OrganizationalUnits, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	keyCondition := expression.Key("OrgDimensionId").Equal(expression.Value(orgDimensionId)).And(expression.Key("Hierarchy").BeginsWith(hierarchy))
	expressionBuilder := expression.NewBuilder().WithKeyCondition(keyCondition)
	expr, err := expressionBuilder.Build()
	if err != nil {
		return nil, err
	}

	queryInput := &dynamodb.QueryInput{
		TableName:                 &c.unitsTableName,
		IndexName:                 aws.String("OrgDimensionId-Hierarchy-index"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		Limit:                     &limit,
		ExclusiveStartKey:         startKey,
	}

	resultItems, lastEvaluatedKey, err := helpers.QueryDynamoDBUntilLimit(ctx, c.dynamodb, queryInput, limit, []string{"OrgDimensionId", "OrgUnitId", "Hierarchy"})
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

func (c *DatabaseClient) PutOrganizationalUnit(ctx context.Context, input *models.OrganizationalUnit) error {
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

func (c *DatabaseClient) DeleteOrganizationalUnit(ctx context.Context, orgDimensionId string, orgUnitId string) error {
	condition := expression.AttributeExists(expression.Name("OrgUnitId"))

	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return err
	}

	_, err = c.dynamodb.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: &c.unitsTableName,
		Key: map[string]types.AttributeValue{
			"OrgDimensionId": &types.AttributeValueMemberS{Value: orgDimensionId},
			"OrgUnitId":      &types.AttributeValueMemberS{Value: orgUnitId},
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

func (c *DatabaseClient) UpdateOrganizationalUnit(ctx context.Context, orgDimensionId string, orgUnitId string, update *OrganizationalUnitUpdate) (*models.OrganizationalUnit, error) {
	condition := expression.AttributeExists(expression.Name("OrgUnitId"))

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
		TableName: &c.unitsTableName,
		Key: map[string]types.AttributeValue{
			"OrgDimensionId": &types.AttributeValueMemberS{Value: orgDimensionId},
			"OrgUnitId":      &types.AttributeValueMemberS{Value: orgUnitId},
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
