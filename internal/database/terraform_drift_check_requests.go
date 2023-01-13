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

func (c *DatabaseClient) GetTerraformDriftCheckRequest(ctx context.Context, terraformDriftCheckRequestId string) (*models.TerraformDriftCheckRequest, error) {
	response, err := c.dynamodb.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &c.terraformDriftCheckRequestsTableName,
		Key: map[string]types.AttributeValue{
			"TerraformDriftCheckRequestId": &types.AttributeValueMemberS{Value: terraformDriftCheckRequestId},
		},
	})
	if err != nil {
		return nil, err
	} else if response.Item == nil {
		return nil, helpers.NotFoundError{Message: fmt.Sprintf("Terraform Workflow Request %q not found", terraformDriftCheckRequestId)}
	}

	terraformDriftCheckRequest := models.TerraformDriftCheckRequest{}
	if err = attributevalue.UnmarshalMap(response.Item, &terraformDriftCheckRequest); err != nil {
		return nil, err
	}

	return &terraformDriftCheckRequest, nil
}

func (c *DatabaseClient) GetTerraformDriftCheckRequestsByIds(ctx context.Context, ids []string) ([]models.TerraformDriftCheckRequest, error) {
	fmt.Println("getting terraform drift check workflow requests by ids", len(ids))
	var keys []map[string]types.AttributeValue

	for _, id := range ids {
		keys = append(keys, map[string]types.AttributeValue{
			"TerraformDriftCheckRequestId": &types.AttributeValueMemberS{Value: id},
		})
	}

	bii := dynamodb.BatchGetItemInput{
		RequestItems: map[string]types.KeysAndAttributes{
			c.terraformDriftCheckRequestsTableName: {
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
			items = append(items, bgo.Responses[c.terraformDriftCheckRequestsTableName]...)
		}
		requestItems := bgo.UnprocessedKeys
		bii = dynamodb.BatchGetItemInput{RequestItems: requestItems}
		if len(requestItems) == 0 {
			break
		}
	}

	items = SortDynamoDBBatchResponses(keys, items)

	results := []models.TerraformDriftCheckRequest{}
	err := attributevalue.UnmarshalListOfMaps(items, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (c *DatabaseClient) GetTerraformDriftCheckRequests(ctx context.Context, limit int32, cursor string) (*models.TerraformDriftCheckRequests, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	scanInput := &dynamodb.ScanInput{
		TableName:         &c.terraformDriftCheckRequestsTableName,
		Limit:             &limit,
		ExclusiveStartKey: startKey,
	}

	resultItems, lastEvaluatedKey, err := helpers.ScanDynamoDBUntilLimit(ctx, c.dynamodb, scanInput, limit, []string{"TerraformDriftCheckRequestId"})
	if err != nil {
		return nil, err
	}

	terraformDriftCheckRequests := []models.TerraformDriftCheckRequest{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &terraformDriftCheckRequests)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.TerraformDriftCheckRequests{
		Items:      terraformDriftCheckRequests,
		NextCursor: nextCursor,
	}, nil
}

func (c DatabaseClient) GetTerraformDriftCheckRequestsByModulePropagationDriftCheckRequestId(ctx context.Context, modulePropagationDriftCheckRequestId string, limit int32, cursor string) (*models.TerraformDriftCheckRequests, error) {
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
		TableName:                 &c.terraformDriftCheckRequestsTableName,
		IndexName:                 aws.String("ModulePropagationDriftCheckRequestId-RequestTime-index"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		Limit:                     &limit,
		ExclusiveStartKey:         startKey,
		ScanIndexForward:          aws.Bool(false),
	}

	resultItems, lastEvaluatedKey, err := helpers.QueryDynamoDBUntilLimit(ctx, c.dynamodb, queryInput, limit, []string{"TerraformDriftCheckRequestId", "ModulePropagationDriftCheckRequestId", "RequestTime"})
	if err != nil {
		return nil, err
	}

	terraformDriftCheckRequests := []models.TerraformDriftCheckRequest{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &terraformDriftCheckRequests)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.TerraformDriftCheckRequests{
		Items:      terraformDriftCheckRequests,
		NextCursor: nextCursor,
	}, nil
}

func (c DatabaseClient) GetTerraformDriftCheckRequestsByModuleAssignmentId(ctx context.Context, moduleAssignmentId string, limit int32, cursor string) (*models.TerraformDriftCheckRequests, error) {
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
		TableName:                 &c.terraformDriftCheckRequestsTableName,
		IndexName:                 aws.String("ModuleAssignmentId-RequestTime-index"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		Limit:                     &limit,
		ExclusiveStartKey:         startKey,
		ScanIndexForward:          aws.Bool(false),
	}

	resultItems, lastEvaluatedKey, err := helpers.QueryDynamoDBUntilLimit(ctx, c.dynamodb, queryInput, limit, []string{"TerraformDriftCheckRequestId", "ModuleAssignmentId", "RequestTime"})
	if err != nil {
		return nil, err
	}

	terraformDriftCheckRequests := []models.TerraformDriftCheckRequest{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &terraformDriftCheckRequests)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.TerraformDriftCheckRequests{
		Items:      terraformDriftCheckRequests,
		NextCursor: nextCursor,
	}, nil
}

func (c *DatabaseClient) PutTerraformDriftCheckRequest(ctx context.Context, input *models.TerraformDriftCheckRequest) error {
	item, err := attributevalue.MarshalMap(input)
	if err != nil {
		return err
	}

	condition := expression.AttributeNotExists(expression.Name("TerraformDriftCheckRequestId"))

	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return err
	}

	_, err = c.dynamodb.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:                 &c.terraformDriftCheckRequestsTableName,
		Item:                      item,
		ConditionExpression:       expr.Condition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})

	ccfe := &types.ConditionalCheckFailedException{}
	switch {
	case errors.As(err, &ccfe):
		return helpers.AlreadyExistsError{Message: fmt.Sprintf("Terraform Workflow Request %q already exists", input.TerraformDriftCheckRequestId)}
	default:
		return err
	}
}

func (c *DatabaseClient) UpdateTerraformDriftCheckRequest(ctx context.Context, terraformDriftCheckRequestId string, update *models.TerraformDriftCheckRequestUpdate) (*models.TerraformDriftCheckRequest, error) {
	condition := expression.AttributeExists(expression.Name("TerraformDriftCheckRequestId"))

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
		TableName:                 &c.terraformDriftCheckRequestsTableName,
		Key:                       map[string]types.AttributeValue{"TerraformDriftCheckRequestId": &types.AttributeValueMemberS{Value: terraformDriftCheckRequestId}},
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
		return nil, helpers.NotFoundError{Message: fmt.Sprintf("Terraform Workflow Request %q not found", terraformDriftCheckRequestId)}
	default:
		if err != nil {
			return nil, err
		}
	}

	terraformDriftCheckRequest := models.TerraformDriftCheckRequest{}
	err = attributevalue.UnmarshalMap(result.Attributes, &terraformDriftCheckRequest)
	if err != nil {
		return nil, err
	}

	return &terraformDriftCheckRequest, nil
}
