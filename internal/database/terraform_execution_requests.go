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

func (c *DatabaseClient) GetTerraformExecutionRequest(ctx context.Context, terraformExecutionRequestId string) (*models.TerraformExecutionRequest, error) {
	response, err := c.dynamodb.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &c.terraformExecutionRequestsTableName,
		Key: map[string]types.AttributeValue{
			"TerraformExecutionRequestId": &types.AttributeValueMemberS{Value: terraformExecutionRequestId},
		},
	})
	if err != nil {
		return nil, err
	} else if response.Item == nil {
		return nil, helpers.NotFoundError{Message: fmt.Sprintf("Terraform Workflow Request %q not found", terraformExecutionRequestId)}
	}

	terraformExecutionRequest := models.TerraformExecutionRequest{}
	if err = attributevalue.UnmarshalMap(response.Item, &terraformExecutionRequest); err != nil {
		return nil, err
	}

	return &terraformExecutionRequest, nil
}

func (c *DatabaseClient) GetTerraformExecutionRequestsByIds(ctx context.Context, ids []string) ([]models.TerraformExecutionRequest, error) {
	fmt.Println("getting terraform execution workflow requests by ids", len(ids))
	var keys []map[string]types.AttributeValue

	for _, id := range ids {
		keys = append(keys, map[string]types.AttributeValue{
			"TerraformExecutionRequestId": &types.AttributeValueMemberS{Value: id},
		})
	}

	bii := dynamodb.BatchGetItemInput{
		RequestItems: map[string]types.KeysAndAttributes{
			c.terraformExecutionRequestsTableName: {
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
			items = append(items, bgo.Responses[c.terraformExecutionRequestsTableName]...)
		}
		requestItems := bgo.UnprocessedKeys
		bii = dynamodb.BatchGetItemInput{RequestItems: requestItems}
		if len(requestItems) == 0 {
			break
		}
	}

	items = SortDynamoDBBatchResponses(keys, items)

	results := []models.TerraformExecutionRequest{}
	err := attributevalue.UnmarshalListOfMaps(items, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (c *DatabaseClient) GetTerraformExecutionRequests(ctx context.Context, limit int32, cursor string) (*models.TerraformExecutionRequests, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	scanInput := &dynamodb.ScanInput{
		TableName:         &c.terraformExecutionRequestsTableName,
		Limit:             &limit,
		ExclusiveStartKey: startKey,
	}

	resultItems, lastEvaluatedKey, err := helpers.ScanDynamoDBUntilLimit(ctx, c.dynamodb, scanInput, limit, []string{"TerraformExecutionRequestId"})
	if err != nil {
		return nil, err
	}

	terraformExecutionRequests := []models.TerraformExecutionRequest{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &terraformExecutionRequests)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.TerraformExecutionRequests{
		Items:      terraformExecutionRequests,
		NextCursor: nextCursor,
	}, nil
}

func (c DatabaseClient) GetTerraformExecutionRequestsByModulePropagationExecutionRequestId(ctx context.Context, modulePropagationExecutionRequestId string, limit int32, cursor string) (*models.TerraformExecutionRequests, error) {
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
		TableName:                 &c.terraformExecutionRequestsTableName,
		IndexName:                 aws.String("ModulePropagationExecutionRequestId-RequestTime-index"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		Limit:                     &limit,
		ExclusiveStartKey:         startKey,
		ScanIndexForward:          aws.Bool(false),
	}

	resultItems, lastEvaluatedKey, err := helpers.QueryDynamoDBUntilLimit(ctx, c.dynamodb, queryInput, limit, []string{"TerraformExecutionRequestId", "ModulePropagationExecutionRequestId", "RequestTime"})
	if err != nil {
		return nil, err
	}

	terraformExecutionRequests := []models.TerraformExecutionRequest{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &terraformExecutionRequests)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.TerraformExecutionRequests{
		Items:      terraformExecutionRequests,
		NextCursor: nextCursor,
	}, nil
}

func (c DatabaseClient) GetTerraformExecutionRequestsByModuleAssignmentId(ctx context.Context, moduleAssignmentId string, limit int32, cursor string) (*models.TerraformExecutionRequests, error) {
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
		TableName:                 &c.terraformExecutionRequestsTableName,
		IndexName:                 aws.String("ModuleAssignmentId-RequestTime-index"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		Limit:                     &limit,
		ExclusiveStartKey:         startKey,
		ScanIndexForward:          aws.Bool(false),
	}

	resultItems, lastEvaluatedKey, err := helpers.QueryDynamoDBUntilLimit(ctx, c.dynamodb, queryInput, limit, []string{"TerraformExecutionRequestId", "ModuleAssignmentId", "RequestTime"})
	if err != nil {
		return nil, err
	}

	terraformExecutionRequests := []models.TerraformExecutionRequest{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &terraformExecutionRequests)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.TerraformExecutionRequests{
		Items:      terraformExecutionRequests,
		NextCursor: nextCursor,
	}, nil
}

func (c *DatabaseClient) PutTerraformExecutionRequest(ctx context.Context, input *models.TerraformExecutionRequest) error {
	item, err := attributevalue.MarshalMap(input)
	if err != nil {
		return err
	}

	condition := expression.AttributeNotExists(expression.Name("TerraformExecutionRequestId"))

	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return err
	}

	_, err = c.dynamodb.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:                 &c.terraformExecutionRequestsTableName,
		Item:                      item,
		ConditionExpression:       expr.Condition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})

	ccfe := &types.ConditionalCheckFailedException{}
	switch {
	case errors.As(err, &ccfe):
		return helpers.AlreadyExistsError{Message: fmt.Sprintf("Terraform Workflow Request %q already exists", input.TerraformExecutionRequestId)}
	default:
		return err
	}
}

func (c *DatabaseClient) UpdateTerraformExecutionRequest(ctx context.Context, terraformExecutionRequestId string, update *models.TerraformExecutionRequestUpdate) (*models.TerraformExecutionRequest, error) {
	condition := expression.AttributeExists(expression.Name("TerraformExecutionRequestId"))

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
		TableName:                 &c.terraformExecutionRequestsTableName,
		Key:                       map[string]types.AttributeValue{"TerraformExecutionRequestId": &types.AttributeValueMemberS{Value: terraformExecutionRequestId}},
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
		return nil, helpers.NotFoundError{Message: fmt.Sprintf("Terraform Workflow Request %q not found", terraformExecutionRequestId)}
	default:
		if err != nil {
			return nil, err
		}
	}

	terraformExecutionRequest := models.TerraformExecutionRequest{}
	err = attributevalue.UnmarshalMap(result.Attributes, &terraformExecutionRequest)
	if err != nil {
		return nil, err
	}

	return &terraformExecutionRequest, nil
}
