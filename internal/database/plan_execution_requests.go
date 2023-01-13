package database

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/sheacloud/tfom/internal/helpers"
	"github.com/sheacloud/tfom/pkg/models"
)

func (c *DatabaseClient) GetPlanExecutionRequest(ctx context.Context, planExecutionRequestId string) (*models.PlanExecutionRequest, error) {
	fmt.Println("getting plan execution request")
	response, err := c.dynamodb.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &c.planExecutionsTableName,
		Key: map[string]types.AttributeValue{
			"PlanExecutionRequestId": &types.AttributeValueMemberS{Value: planExecutionRequestId},
		},
	})
	if err != nil {
		return nil, err
	} else if response.Item == nil {
		return nil, helpers.NotFoundError{Message: fmt.Sprintf("Plan Execution Request %q not found", planExecutionRequestId)}
	}

	planExecutionRequest := models.PlanExecutionRequest{}
	if err = attributevalue.UnmarshalMap(response.Item, &planExecutionRequest); err != nil {
		return nil, err
	}

	return &planExecutionRequest, nil
}

func (c *DatabaseClient) GetPlanExecutionRequestsByIds(ctx context.Context, ids []string) ([]models.PlanExecutionRequest, error) {
	fmt.Println("getting plan execution requests by ids", len(ids))
	var keys []map[string]types.AttributeValue

	for _, id := range ids {
		keys = append(keys, map[string]types.AttributeValue{
			"PlanExecutionRequestId": &types.AttributeValueMemberS{Value: id},
		})
	}

	bii := dynamodb.BatchGetItemInput{
		RequestItems: map[string]types.KeysAndAttributes{
			c.planExecutionsTableName: {
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
			items = append(items, bgo.Responses[c.planExecutionsTableName]...)
		}
		requestItems := bgo.UnprocessedKeys
		bii = dynamodb.BatchGetItemInput{RequestItems: requestItems}
		if len(requestItems) == 0 {
			break
		}
	}

	items = SortDynamoDBBatchResponses(keys, items)

	results := []models.PlanExecutionRequest{}
	err := attributevalue.UnmarshalListOfMaps(items, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (c DatabaseClient) GetPlanExecutionRequests(ctx context.Context, limit int32, cursor string) (*models.PlanExecutionRequests, error) {
	startKey, err := helpers.GetKeyFromCursor(cursor)
	if err != nil {
		return nil, err
	}

	scanInput := &dynamodb.ScanInput{
		TableName:         &c.planExecutionsTableName,
		Limit:             &limit,
		ExclusiveStartKey: startKey,
	}

	resultItems, lastEvaluatedKey, err := helpers.ScanDynamoDBUntilLimit(ctx, c.dynamodb, scanInput, limit, []string{"PlanExecutionRequestId"})
	if err != nil {
		return nil, err
	}

	planExecutionRequests := []models.PlanExecutionRequest{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &planExecutionRequests)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.PlanExecutionRequests{
		Items:      planExecutionRequests,
		NextCursor: nextCursor,
	}, nil
}

func (c DatabaseClient) GetPlanExecutionRequestsByModuleAssignmentId(ctx context.Context, moduleAssignmentId string, limit int32, cursor string) (*models.PlanExecutionRequests, error) {
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
		TableName:                 &c.planExecutionsTableName,
		IndexName:                 aws.String("ModuleAssignmentId-RequestTime-index"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		Limit:                     &limit,
		ExclusiveStartKey:         startKey,
		ScanIndexForward:          aws.Bool(false),
	}

	resultItems, lastEvaluatedKey, err := helpers.QueryDynamoDBUntilLimit(ctx, c.dynamodb, queryInput, limit, []string{"PlanExecutionRequestId", "ModuleAssignmentId", "RequestTime"})
	if err != nil {
		return nil, err
	}

	planExecutionRequests := []models.PlanExecutionRequest{}
	var nextCursor string

	err = attributevalue.UnmarshalListOfMaps(resultItems, &planExecutionRequests)
	if err != nil {
		return nil, err
	}

	if lastEvaluatedKey != nil {
		nextCursor, err = helpers.GetCursorFromKey(lastEvaluatedKey)
		if err != nil {
			return nil, err
		}
	}

	return &models.PlanExecutionRequests{
		Items:      planExecutionRequests,
		NextCursor: nextCursor,
	}, nil
}

func (c *DatabaseClient) PutPlanExecutionRequest(ctx context.Context, input *models.PlanExecutionRequest) error {
	item, err := attributevalue.MarshalMap(input)
	if err != nil {
		return err
	}

	condition := expression.AttributeNotExists(expression.Name("PlanExecutionRequestId"))

	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return err
	}

	_, err = c.dynamodb.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:                 &c.planExecutionsTableName,
		Item:                      item,
		ConditionExpression:       expr.Condition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})

	ccfe := &types.ConditionalCheckFailedException{}
	switch {
	case errors.As(err, &ccfe):
		return helpers.AlreadyExistsError{Message: fmt.Sprintf("Plan Execution Request %q already exists", input.PlanExecutionRequestId)}
	default:
		return err
	}
}

func (c *DatabaseClient) UpdatePlanExecutionRequest(ctx context.Context, planExecutionRequestId string, update *models.PlanExecutionRequestUpdate) (*models.PlanExecutionRequest, error) {
	condition := expression.AttributeExists(expression.Name("PlanExecutionRequestId"))

	expr, err := expression.NewBuilder().WithCondition(condition).Build()
	if err != nil {
		return nil, err
	}

	updateBuilder := expression.UpdateBuilder{}
	if update.InitOutputKey != nil {
		updateBuilder = updateBuilder.Set(expression.Name("InitOutputKey"), expression.Value(*update.InitOutputKey))
	}
	if update.PlanOutputKey != nil {
		updateBuilder = updateBuilder.Set(expression.Name("PlanOutputKey"), expression.Value(*update.PlanOutputKey))
	}
	if update.Status != nil {
		updateBuilder = updateBuilder.Set(expression.Name("Status"), expression.Value(*update.Status))
	}

	updateExpression, err := expression.NewBuilder().WithUpdate(updateBuilder).Build()
	if err != nil {
		return nil, err
	}

	updateInput := &dynamodb.UpdateItemInput{
		TableName:                 &c.planExecutionsTableName,
		Key:                       map[string]types.AttributeValue{"PlanExecutionRequestId": &types.AttributeValueMemberS{Value: planExecutionRequestId}},
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
		return nil, helpers.NotFoundError{Message: fmt.Sprintf("Plan Execution Request %q not found", planExecutionRequestId)}
	default:
		if err != nil {
			return nil, err
		}
	}

	planExecutionRequest := models.PlanExecutionRequest{}
	err = attributevalue.UnmarshalMap(result.Attributes, &planExecutionRequest)
	if err != nil {
		return nil, err
	}

	return &planExecutionRequest, nil
}

func (c *DatabaseClient) UploadTerraformPlanInitResults(ctx context.Context, planExecutionRequestId string, initResults *models.TerraformInitOutput) (string, error) {
	outputKey := fmt.Sprintf("plans/%s/init-results.json", planExecutionRequestId)

	initResultsBytes, err := json.Marshal(initResults)
	if err != nil {
		return "", err
	}

	_, err = c.s3.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &c.resultsBucketName,
		Key:    &outputKey,
		Body:   bytes.NewReader(initResultsBytes),
	})

	return outputKey, err
}

func (c *DatabaseClient) UploadTerraformPlanResults(ctx context.Context, planExecutionRequestId string, planResults *models.TerraformPlanOutput) (string, error) {
	outputKey := fmt.Sprintf("plans/%s/plan-results.json", planExecutionRequestId)

	planResultsBytes, err := json.Marshal(planResults)
	if err != nil {
		return "", err
	}

	_, err = c.s3.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &c.resultsBucketName,
		Key:    &outputKey,
		Body:   bytes.NewReader(planResultsBytes),
	})

	return outputKey, err
}

func (c *DatabaseClient) DownloadTerraformPlanInitResults(ctx context.Context, initResultsObjectKey string) (*models.TerraformInitOutput, error) {
	result, err := c.s3.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &c.resultsBucketName,
		Key:    &initResultsObjectKey,
	})
	if err != nil {
		return nil, err
	}

	initResults := models.TerraformInitOutput{}
	err = json.NewDecoder(result.Body).Decode(&initResults)
	if err != nil {
		return nil, err
	}

	return &initResults, nil
}

func (c *DatabaseClient) DownloadTerraformPlanResults(ctx context.Context, planResultsObjectKey string) (*models.TerraformPlanOutput, error) {
	result, err := c.s3.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &c.resultsBucketName,
		Key:    &planResultsObjectKey,
	})
	if err != nil {
		return nil, err
	}

	planResults := models.TerraformPlanOutput{}
	err = json.NewDecoder(result.Body).Decode(&planResults)
	if err != nil {
		return nil, err
	}

	return &planResults, nil
}
