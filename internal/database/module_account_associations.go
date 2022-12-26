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

func (c *OrganizationsDatabaseClient) GetModuleAccountAssociation(ctx context.Context, modulePropagationId string, orgAccountId string) (*models.ModuleAccountAssociation, error) {
	response, err := c.dynamodb.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &c.moduleAccountAssociationsTableName,
		Key: map[string]types.AttributeValue{
			"ModulePropagationId": &types.AttributeValueMemberS{Value: modulePropagationId},
			"OrgAccountId":        &types.AttributeValueMemberS{Value: orgAccountId},
		},
	})
	if err != nil {
		return nil, err
	} else if response.Item == nil {
		return nil, helpers.NotFoundError{Message: fmt.Sprintf("Module Account Association %q %q not found", modulePropagationId, orgAccountId)}
	}

	moduleAccountAssociation := models.ModuleAccountAssociation{}
	if err = attributevalue.UnmarshalMap(response.Item, &moduleAccountAssociation); err != nil {
		return nil, err
	}

	return &moduleAccountAssociation, nil
}

func (c *OrganizationsDatabaseClient) GetModuleAccountAssociations(ctx context.Context, limit int32, cursor string) (*models.ModuleAccountAssociations, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	scanInput := &dynamodb.ScanInput{
		TableName:         &c.moduleAccountAssociationsTableName,
		Limit:             &limit,
		ExclusiveStartKey: startKey,
	}

	resultItems, lastEvaluatedKey, err := helpers.ScanDynamoDBUntilLimit(ctx, c.dynamodb, scanInput, limit, []string{"ModulePropagationId", "OrgAccountId"})
	if err != nil {
		return nil, err
	}

	moduleAccountAssociations := []models.ModuleAccountAssociation{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &moduleAccountAssociations)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.ModuleAccountAssociations{
		Items:      moduleAccountAssociations,
		NextCursor: nextCursor,
	}, nil
}

func (c OrganizationsDatabaseClient) GetModuleAccountAssociationsByModulePropagationId(ctx context.Context, modulePropagationId string, limit int32, cursor string) (*models.ModuleAccountAssociations, error) {
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
		TableName:                 &c.moduleAccountAssociationsTableName,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		Limit:                     &limit,
		ExclusiveStartKey:         startKey,
	}

	resultItems, lastEvaluatedKey, err := helpers.QueryDynamoDBUntilLimit(ctx, c.dynamodb, queryInput, limit, []string{"ModulePropagationId", "OrgAccountId"})
	if err != nil {
		return nil, err
	}

	moduleAccountAssociations := []models.ModuleAccountAssociation{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &moduleAccountAssociations)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.ModuleAccountAssociations{
		Items:      moduleAccountAssociations,
		NextCursor: nextCursor,
	}, nil
}

func (c *OrganizationsDatabaseClient) PutModuleAccountAssociation(ctx context.Context, input *models.ModuleAccountAssociation) error {
	item, err := attributevalue.MarshalMap(input)
	if err != nil {
		return err
	}

	condition := expression.AttributeNotExists(expression.Name("ModuleAccountAssociationId"))

	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return err
	}

	_, err = c.dynamodb.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:                 &c.moduleAccountAssociationsTableName,
		Item:                      item,
		ConditionExpression:       expr.Condition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})

	ccfe := &types.ConditionalCheckFailedException{}
	switch {
	case errors.As(err, &ccfe):
		return helpers.AlreadyExistsError{Message: fmt.Sprintf("Module Account Association %q %q already exists", input.ModulePropagationId, input.OrgAccountId)}
	default:
		return err
	}
}

func (c *OrganizationsDatabaseClient) UpdateModuleAccountAssociation(ctx context.Context, modulePropagationId string, orgAccountId string, update *models.ModuleAccountAssociationUpdate) (*models.ModuleAccountAssociation, error) {
	condition := expression.AttributeExists(expression.Name("OrgAccountId"))

	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return nil, err
	}

	updateBuilder := expression.UpdateBuilder{}
	if update.Status != nil {
		updateBuilder = updateBuilder.Set(expression.Name("Status"), expression.Value(*update.Status))
	}
	if update.RemoteStateBucket != nil {
		updateBuilder = updateBuilder.Set(expression.Name("RemoteStateBucket"), expression.Value(*update.RemoteStateBucket))
	}
	if update.RemoteStateKey != nil {
		updateBuilder = updateBuilder.Set(expression.Name("RemoteStateKey"), expression.Value(*update.RemoteStateKey))
	}

	updateExpression, err := expression.NewBuilder().WithUpdate(updateBuilder).Build()
	if err != nil {
		return nil, err
	}

	updateInput := &dynamodb.UpdateItemInput{
		TableName: &c.moduleAccountAssociationsTableName,
		Key: map[string]types.AttributeValue{
			"ModulePropagationId": &types.AttributeValueMemberS{Value: modulePropagationId},
			"OrgAccountId":        &types.AttributeValueMemberS{Value: orgAccountId},
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
		return nil, helpers.NotFoundError{Message: fmt.Sprintf("Module Account Association %q not found", orgAccountId)}
	default:
		if err != nil {
			return nil, err
		}
	}

	moduleAccountAssociation := models.ModuleAccountAssociation{}
	err = attributevalue.UnmarshalMap(result.Attributes, &moduleAccountAssociation)
	if err != nil {
		return nil, err
	}

	return &moduleAccountAssociation, nil
}
