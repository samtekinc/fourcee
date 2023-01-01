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

func (c *OrganizationsDatabaseClient) GetTerraformWorkflowRequest(ctx context.Context, terraformWorkflowRequestId string) (*models.TerraformWorkflowRequest, error) {
	response, err := c.dynamodb.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &c.terraformWorkflowRequestsTableName,
		Key: map[string]types.AttributeValue{
			"TerraformWorkflowRequestId": &types.AttributeValueMemberS{Value: terraformWorkflowRequestId},
		},
	})
	if err != nil {
		return nil, err
	} else if response.Item == nil {
		return nil, helpers.NotFoundError{Message: fmt.Sprintf("Terraform Workflow Request %q not found", terraformWorkflowRequestId)}
	}

	terraformWorkflowRequest := models.TerraformWorkflowRequest{}
	if err = attributevalue.UnmarshalMap(response.Item, &terraformWorkflowRequest); err != nil {
		return nil, err
	}

	return &terraformWorkflowRequest, nil
}

func (c *OrganizationsDatabaseClient) GetTerraformWorkflowRequests(ctx context.Context, limit int32, cursor string) (*models.TerraformWorkflowRequests, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	scanInput := &dynamodb.ScanInput{
		TableName:         &c.terraformWorkflowRequestsTableName,
		Limit:             &limit,
		ExclusiveStartKey: startKey,
	}

	resultItems, lastEvaluatedKey, err := helpers.ScanDynamoDBUntilLimit(ctx, c.dynamodb, scanInput, limit, []string{"TerraformWorkflowRequestId"})
	if err != nil {
		return nil, err
	}

	terraformWorkflowRequests := []models.TerraformWorkflowRequest{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &terraformWorkflowRequests)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.TerraformWorkflowRequests{
		Items:      terraformWorkflowRequests,
		NextCursor: nextCursor,
	}, nil
}

func (c OrganizationsDatabaseClient) GetTerraformWorkflowRequestsByModulePropagationExecutionRequestId(ctx context.Context, modulePropagationExecutionRequestId string, limit int32, cursor string) (*models.TerraformWorkflowRequests, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	keyCondition := expression.Key("ModulePropagationExecutionRequestId").Equal(expression.Value(modulePropagationExecutionRequestId))
	expressionBuilder := expression.NewBuilder().WithKeyCondition(keyCondition)
	expr, err := expressionBuilder.Build()
	if err != nil {
		return nil, err
	}

	queryInput := &dynamodb.QueryInput{
		TableName:                 &c.terraformWorkflowRequestsTableName,
		IndexName:                 aws.String("ModulePropagationExecutionRequestId-RequestTime-index"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		Limit:                     &limit,
		ExclusiveStartKey:         startKey,
		ScanIndexForward:          aws.Bool(false),
	}

	resultItems, lastEvaluatedKey, err := helpers.QueryDynamoDBUntilLimit(ctx, c.dynamodb, queryInput, limit, []string{"TerraformWorkflowRequestId", "ModulePropagationExecutionRequestId", "RequestTime"})
	if err != nil {
		return nil, err
	}

	terraformWorkflowRequests := []models.TerraformWorkflowRequest{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &terraformWorkflowRequests)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.TerraformWorkflowRequests{
		Items:      terraformWorkflowRequests,
		NextCursor: nextCursor,
	}, nil
}

func (c OrganizationsDatabaseClient) GetTerraformWorkflowRequestsByModuleAccountAssociationKey(ctx context.Context, moduleAccountAssociationKey string, limit int32, cursor string) (*models.TerraformWorkflowRequests, error) {
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
		TableName:                 &c.terraformWorkflowRequestsTableName,
		IndexName:                 aws.String("ModuleAccountAssociationKey-RequestTime-index"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		Limit:                     &limit,
		ExclusiveStartKey:         startKey,
		ScanIndexForward:          aws.Bool(false),
	}

	resultItems, lastEvaluatedKey, err := helpers.QueryDynamoDBUntilLimit(ctx, c.dynamodb, queryInput, limit, []string{"TerraformWorkflowRequestId", "ModuleAccountAssociationKey", "RequestTime"})
	if err != nil {
		return nil, err
	}

	terraformWorkflowRequests := []models.TerraformWorkflowRequest{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &terraformWorkflowRequests)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.TerraformWorkflowRequests{
		Items:      terraformWorkflowRequests,
		NextCursor: nextCursor,
	}, nil
}

func (c *OrganizationsDatabaseClient) PutTerraformWorkflowRequest(ctx context.Context, input *models.TerraformWorkflowRequest) error {
	item, err := attributevalue.MarshalMap(input)
	if err != nil {
		return err
	}

	condition := expression.AttributeNotExists(expression.Name("TerraformWorkflowRequestId"))

	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return err
	}

	_, err = c.dynamodb.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:                 &c.terraformWorkflowRequestsTableName,
		Item:                      item,
		ConditionExpression:       expr.Condition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})

	ccfe := &types.ConditionalCheckFailedException{}
	switch {
	case errors.As(err, &ccfe):
		return helpers.AlreadyExistsError{Message: fmt.Sprintf("Terraform Workflow Request %q already exists", input.TerraformWorkflowRequestId)}
	default:
		return err
	}
}

func (c *OrganizationsDatabaseClient) UpdateTerraformWorkflowRequest(ctx context.Context, terraformWorkflowRequestId string, update *models.TerraformWorkflowRequestUpdate) (*models.TerraformWorkflowRequest, error) {
	condition := expression.AttributeExists(expression.Name("TerraformWorkflowRequestId"))

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
	if update.ApplyExecutionRequestId != nil {
		updateBuilder = updateBuilder.Set(expression.Name("ApplyExecutionRequestId"), expression.Value(*update.ApplyExecutionRequestId))
	}

	updateExpression, err := expression.NewBuilder().WithUpdate(updateBuilder).Build()
	if err != nil {
		return nil, err
	}

	updateInput := &dynamodb.UpdateItemInput{
		TableName:                 &c.terraformWorkflowRequestsTableName,
		Key:                       map[string]types.AttributeValue{"TerraformWorkflowRequestId": &types.AttributeValueMemberS{Value: terraformWorkflowRequestId}},
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
		return nil, helpers.NotFoundError{Message: fmt.Sprintf("Terraform Workflow Request %q not found", terraformWorkflowRequestId)}
	default:
		if err != nil {
			return nil, err
		}
	}

	terraformWorkflowRequest := models.TerraformWorkflowRequest{}
	err = attributevalue.UnmarshalMap(result.Attributes, &terraformWorkflowRequest)
	if err != nil {
		return nil, err
	}

	return &terraformWorkflowRequest, nil
}
