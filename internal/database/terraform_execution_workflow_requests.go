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

func (c *OrganizationsDatabaseClient) GetTerraformExecutionWorkflowRequest(ctx context.Context, terraformExecutionWorkflowRequestId string) (*models.TerraformExecutionWorkflowRequest, error) {
	response, err := c.dynamodb.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &c.terraformExecutionWorkflowRequestsTableName,
		Key: map[string]types.AttributeValue{
			"TerraformExecutionWorkflowRequestId": &types.AttributeValueMemberS{Value: terraformExecutionWorkflowRequestId},
		},
	})
	if err != nil {
		return nil, err
	} else if response.Item == nil {
		return nil, helpers.NotFoundError{Message: fmt.Sprintf("Terraform Workflow Request %q not found", terraformExecutionWorkflowRequestId)}
	}

	terraformExecutionWorkflowRequest := models.TerraformExecutionWorkflowRequest{}
	if err = attributevalue.UnmarshalMap(response.Item, &terraformExecutionWorkflowRequest); err != nil {
		return nil, err
	}

	return &terraformExecutionWorkflowRequest, nil
}

func (c *OrganizationsDatabaseClient) GetTerraformExecutionWorkflowRequests(ctx context.Context, limit int32, cursor string) (*models.TerraformExecutionWorkflowRequests, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	scanInput := &dynamodb.ScanInput{
		TableName:         &c.terraformExecutionWorkflowRequestsTableName,
		Limit:             &limit,
		ExclusiveStartKey: startKey,
	}

	resultItems, lastEvaluatedKey, err := helpers.ScanDynamoDBUntilLimit(ctx, c.dynamodb, scanInput, limit, []string{"TerraformExecutionWorkflowRequestId"})
	if err != nil {
		return nil, err
	}

	terraformExecutionWorkflowRequests := []models.TerraformExecutionWorkflowRequest{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &terraformExecutionWorkflowRequests)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.TerraformExecutionWorkflowRequests{
		Items:      terraformExecutionWorkflowRequests,
		NextCursor: nextCursor,
	}, nil
}

func (c OrganizationsDatabaseClient) GetTerraformExecutionWorkflowRequestsByModulePropagationExecutionRequestId(ctx context.Context, modulePropagationExecutionRequestId string, limit int32, cursor string) (*models.TerraformExecutionWorkflowRequests, error) {
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
		TableName:                 &c.terraformExecutionWorkflowRequestsTableName,
		IndexName:                 aws.String("ModulePropagationExecutionRequestId-RequestTime-index"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		Limit:                     &limit,
		ExclusiveStartKey:         startKey,
		ScanIndexForward:          aws.Bool(false),
	}

	resultItems, lastEvaluatedKey, err := helpers.QueryDynamoDBUntilLimit(ctx, c.dynamodb, queryInput, limit, []string{"TerraformExecutionWorkflowRequestId", "ModulePropagationExecutionRequestId", "RequestTime"})
	if err != nil {
		return nil, err
	}

	terraformExecutionWorkflowRequests := []models.TerraformExecutionWorkflowRequest{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &terraformExecutionWorkflowRequests)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.TerraformExecutionWorkflowRequests{
		Items:      terraformExecutionWorkflowRequests,
		NextCursor: nextCursor,
	}, nil
}

func (c OrganizationsDatabaseClient) GetTerraformExecutionWorkflowRequestsByModuleAssignmentId(ctx context.Context, moduleAssignmentId string, limit int32, cursor string) (*models.TerraformExecutionWorkflowRequests, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	keyCondition := expression.Key("ModuleAssignmentId").Equal(expression.Value(moduleAssignmentId))
	expressionBuilder := expression.NewBuilder().WithKeyCondition(keyCondition)
	expr, err := expressionBuilder.Build()
	if err != nil {
		return nil, err
	}

	queryInput := &dynamodb.QueryInput{
		TableName:                 &c.terraformExecutionWorkflowRequestsTableName,
		IndexName:                 aws.String("ModuleAssignmentId-RequestTime-index"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		Limit:                     &limit,
		ExclusiveStartKey:         startKey,
		ScanIndexForward:          aws.Bool(false),
	}

	resultItems, lastEvaluatedKey, err := helpers.QueryDynamoDBUntilLimit(ctx, c.dynamodb, queryInput, limit, []string{"TerraformExecutionWorkflowRequestId", "ModuleAssignmentId", "RequestTime"})
	if err != nil {
		return nil, err
	}

	terraformExecutionWorkflowRequests := []models.TerraformExecutionWorkflowRequest{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &terraformExecutionWorkflowRequests)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.TerraformExecutionWorkflowRequests{
		Items:      terraformExecutionWorkflowRequests,
		NextCursor: nextCursor,
	}, nil
}

func (c *OrganizationsDatabaseClient) PutTerraformExecutionWorkflowRequest(ctx context.Context, input *models.TerraformExecutionWorkflowRequest) error {
	item, err := attributevalue.MarshalMap(input)
	if err != nil {
		return err
	}

	condition := expression.AttributeNotExists(expression.Name("TerraformExecutionWorkflowRequestId"))

	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return err
	}

	_, err = c.dynamodb.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:                 &c.terraformExecutionWorkflowRequestsTableName,
		Item:                      item,
		ConditionExpression:       expr.Condition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})

	ccfe := &types.ConditionalCheckFailedException{}
	switch {
	case errors.As(err, &ccfe):
		return helpers.AlreadyExistsError{Message: fmt.Sprintf("Terraform Workflow Request %q already exists", input.TerraformExecutionWorkflowRequestId)}
	default:
		return err
	}
}

func (c *OrganizationsDatabaseClient) UpdateTerraformExecutionWorkflowRequest(ctx context.Context, terraformExecutionWorkflowRequestId string, update *models.TerraformExecutionWorkflowRequestUpdate) (*models.TerraformExecutionWorkflowRequest, error) {
	condition := expression.AttributeExists(expression.Name("TerraformExecutionWorkflowRequestId"))

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
		TableName:                 &c.terraformExecutionWorkflowRequestsTableName,
		Key:                       map[string]types.AttributeValue{"TerraformExecutionWorkflowRequestId": &types.AttributeValueMemberS{Value: terraformExecutionWorkflowRequestId}},
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
		return nil, helpers.NotFoundError{Message: fmt.Sprintf("Terraform Workflow Request %q not found", terraformExecutionWorkflowRequestId)}
	default:
		if err != nil {
			return nil, err
		}
	}

	terraformExecutionWorkflowRequest := models.TerraformExecutionWorkflowRequest{}
	err = attributevalue.UnmarshalMap(result.Attributes, &terraformExecutionWorkflowRequest)
	if err != nil {
		return nil, err
	}

	return &terraformExecutionWorkflowRequest, nil
}
