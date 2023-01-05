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

func (c *OrganizationsDatabaseClient) GetTerraformDriftCheckWorkflowRequest(ctx context.Context, terraformDriftCheckWorkflowRequestId string) (*models.TerraformDriftCheckWorkflowRequest, error) {
	response, err := c.dynamodb.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &c.terraformDriftCheckWorkflowRequestsTableName,
		Key: map[string]types.AttributeValue{
			"TerraformDriftCheckWorkflowRequestId": &types.AttributeValueMemberS{Value: terraformDriftCheckWorkflowRequestId},
		},
	})
	if err != nil {
		return nil, err
	} else if response.Item == nil {
		return nil, helpers.NotFoundError{Message: fmt.Sprintf("Terraform Workflow Request %q not found", terraformDriftCheckWorkflowRequestId)}
	}

	terraformDriftCheckWorkflowRequest := models.TerraformDriftCheckWorkflowRequest{}
	if err = attributevalue.UnmarshalMap(response.Item, &terraformDriftCheckWorkflowRequest); err != nil {
		return nil, err
	}

	return &terraformDriftCheckWorkflowRequest, nil
}

func (c *OrganizationsDatabaseClient) GetTerraformDriftCheckWorkflowRequests(ctx context.Context, limit int32, cursor string) (*models.TerraformDriftCheckWorkflowRequests, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	scanInput := &dynamodb.ScanInput{
		TableName:         &c.terraformDriftCheckWorkflowRequestsTableName,
		Limit:             &limit,
		ExclusiveStartKey: startKey,
	}

	resultItems, lastEvaluatedKey, err := helpers.ScanDynamoDBUntilLimit(ctx, c.dynamodb, scanInput, limit, []string{"TerraformDriftCheckWorkflowRequestId"})
	if err != nil {
		return nil, err
	}

	terraformDriftCheckWorkflowRequests := []models.TerraformDriftCheckWorkflowRequest{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &terraformDriftCheckWorkflowRequests)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.TerraformDriftCheckWorkflowRequests{
		Items:      terraformDriftCheckWorkflowRequests,
		NextCursor: nextCursor,
	}, nil
}

func (c OrganizationsDatabaseClient) GetTerraformDriftCheckWorkflowRequestsByModulePropagationDriftCheckRequestId(ctx context.Context, modulePropagationDriftCheckRequestId string, limit int32, cursor string) (*models.TerraformDriftCheckWorkflowRequests, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	keyCondition := expression.Key("ModulePropagationDriftCheckRequestId").Equal(expression.Value(modulePropagationDriftCheckRequestId))
	expressionBuilder := expression.NewBuilder().WithKeyCondition(keyCondition)
	expr, err := expressionBuilder.Build()
	if err != nil {
		return nil, err
	}

	queryInput := &dynamodb.QueryInput{
		TableName:                 &c.terraformDriftCheckWorkflowRequestsTableName,
		IndexName:                 aws.String("ModulePropagationDriftCheckRequestId-RequestTime-index"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		Limit:                     &limit,
		ExclusiveStartKey:         startKey,
		ScanIndexForward:          aws.Bool(false),
	}

	resultItems, lastEvaluatedKey, err := helpers.QueryDynamoDBUntilLimit(ctx, c.dynamodb, queryInput, limit, []string{"TerraformDriftCheckWorkflowRequestId", "ModulePropagationDriftCheckRequestId", "RequestTime"})
	if err != nil {
		return nil, err
	}

	terraformDriftCheckWorkflowRequests := []models.TerraformDriftCheckWorkflowRequest{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &terraformDriftCheckWorkflowRequests)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.TerraformDriftCheckWorkflowRequests{
		Items:      terraformDriftCheckWorkflowRequests,
		NextCursor: nextCursor,
	}, nil
}

func (c OrganizationsDatabaseClient) GetTerraformDriftCheckWorkflowRequestsByModuleAccountAssociationKey(ctx context.Context, moduleAccountAssociationKey string, limit int32, cursor string) (*models.TerraformDriftCheckWorkflowRequests, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	keyCondition := expression.Key("ModuleAccountAssociationKey").Equal(expression.Value(moduleAccountAssociationKey))
	expressionBuilder := expression.NewBuilder().WithKeyCondition(keyCondition)
	expr, err := expressionBuilder.Build()
	if err != nil {
		return nil, err
	}

	queryInput := &dynamodb.QueryInput{
		TableName:                 &c.terraformDriftCheckWorkflowRequestsTableName,
		IndexName:                 aws.String("ModuleAccountAssociationKey-RequestTime-index"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		Limit:                     &limit,
		ExclusiveStartKey:         startKey,
		ScanIndexForward:          aws.Bool(false),
	}

	resultItems, lastEvaluatedKey, err := helpers.QueryDynamoDBUntilLimit(ctx, c.dynamodb, queryInput, limit, []string{"TerraformDriftCheckWorkflowRequestId", "ModuleAccountAssociationKey", "RequestTime"})
	if err != nil {
		return nil, err
	}

	terraformDriftCheckWorkflowRequests := []models.TerraformDriftCheckWorkflowRequest{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &terraformDriftCheckWorkflowRequests)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.TerraformDriftCheckWorkflowRequests{
		Items:      terraformDriftCheckWorkflowRequests,
		NextCursor: nextCursor,
	}, nil
}

func (c *OrganizationsDatabaseClient) PutTerraformDriftCheckWorkflowRequest(ctx context.Context, input *models.TerraformDriftCheckWorkflowRequest) error {
	item, err := attributevalue.MarshalMap(input)
	if err != nil {
		return err
	}

	condition := expression.AttributeNotExists(expression.Name("TerraformDriftCheckWorkflowRequestId"))

	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return err
	}

	_, err = c.dynamodb.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:                 &c.terraformDriftCheckWorkflowRequestsTableName,
		Item:                      item,
		ConditionExpression:       expr.Condition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})

	ccfe := &types.ConditionalCheckFailedException{}
	switch {
	case errors.As(err, &ccfe):
		return helpers.AlreadyExistsError{Message: fmt.Sprintf("Terraform Workflow Request %q already exists", input.TerraformDriftCheckWorkflowRequestId)}
	default:
		return err
	}
}

func (c *OrganizationsDatabaseClient) UpdateTerraformDriftCheckWorkflowRequest(ctx context.Context, terraformDriftCheckWorkflowRequestId string, update *models.TerraformDriftCheckWorkflowRequestUpdate) (*models.TerraformDriftCheckWorkflowRequest, error) {
	condition := expression.AttributeExists(expression.Name("TerraformDriftCheckWorkflowRequestId"))

	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return nil, err
	}

	updateBuilder := expression.UpdateBuilder{}
	if update.Status != nil {
		updateBuilder = updateBuilder.Set(expression.Name("Status"), expression.Value(*update.Status))
	}
	if update.PlanExecutionRequestId != nil {
		updateBuilder = updateBuilder.Set(expression.Name("PlanExecutionRequestId"), expression.Value(*update.PlanExecutionRequestId))
	}
	if update.SyncStatus != nil {
		updateBuilder = updateBuilder.Set(expression.Name("SyncStatus"), expression.Value(*update.SyncStatus))
	}

	updateExpression, err := expression.NewBuilder().WithUpdate(updateBuilder).Build()
	if err != nil {
		return nil, err
	}

	updateInput := &dynamodb.UpdateItemInput{
		TableName:                 &c.terraformDriftCheckWorkflowRequestsTableName,
		Key:                       map[string]types.AttributeValue{"TerraformDriftCheckWorkflowRequestId": &types.AttributeValueMemberS{Value: terraformDriftCheckWorkflowRequestId}},
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
		return nil, helpers.NotFoundError{Message: fmt.Sprintf("Terraform Workflow Request %q not found", terraformDriftCheckWorkflowRequestId)}
	default:
		if err != nil {
			return nil, err
		}
	}

	terraformDriftCheckWorkflowRequest := models.TerraformDriftCheckWorkflowRequest{}
	err = attributevalue.UnmarshalMap(result.Attributes, &terraformDriftCheckWorkflowRequest)
	if err != nil {
		return nil, err
	}

	return &terraformDriftCheckWorkflowRequest, nil
}
