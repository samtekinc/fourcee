package helpers

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/sheacloud/tfom/internal/awsclients"
)

func GetCursorFromKey(key map[string]types.AttributeValue) (string, error) {
	cursor := map[string]interface{}{}
	err := attributevalue.UnmarshalMap(key, &cursor)
	if err != nil {
		err = fmt.Errorf("got error unmarshaling key: %w", err)
		log.Print(err)
		return "", err
	}
	cursorBytes, err := json.Marshal(cursor)
	if err != nil {
		err = fmt.Errorf("got error marshaling cursor: %w", err)
		log.Print(err)
		return "", err
	}
	cursorString := base64.URLEncoding.EncodeToString(cursorBytes)
	return cursorString, nil
}

func GetKeyFromCursor(cursor string) (map[string]types.AttributeValue, error) {
	if cursor == "" {
		return nil, nil // nolint:nilnil // Returning nil is valid for the DAOs.
	}
	cursorBytes, err := base64.URLEncoding.DecodeString(cursor)
	if err != nil {
		err = fmt.Errorf("got error decoding cursor %s: %w", cursor, err)
		log.Print(err)
		return nil, err
	}
	cursorMap := map[string]interface{}{}
	err = json.Unmarshal(cursorBytes, &cursorMap)
	if err != nil {
		err = fmt.Errorf("got error unmarshaling cursor %s: %w", cursor, err)
		log.Print(err)
		return nil, err
	}
	key, err := attributevalue.MarshalMap(cursorMap)
	if err != nil {
		err = fmt.Errorf("got error marshaling cursor %s: %w", cursor, err)
		log.Print(err)
		return nil, err
	}
	return key, nil
}

func ScanDynamoDBUntilLimit(ctx context.Context, client awsclients.ScanInterface, scanInput *dynamodb.ScanInput, limit int32, keyNames []string) ([]map[string]types.AttributeValue, map[string]types.AttributeValue, error) {
	resultItems := []map[string]types.AttributeValue{}
	var lastEvaluatedKey map[string]types.AttributeValue

	for {
		result, err := client.Scan(ctx, scanInput)
		if err != nil {
			return nil, nil, err
		}
		resultItems = append(resultItems, result.Items...)
		if len(resultItems) > int(limit) {
			fmt.Println(result.LastEvaluatedKey)
			resultItems = resultItems[:limit]
			lastEvaluatedItem := resultItems[len(resultItems)-1]
			newLastEvaluatedKey := map[string]types.AttributeValue{}
			for _, key := range keyNames {
				newLastEvaluatedKey[key] = lastEvaluatedItem[key]
			}
			fmt.Println(newLastEvaluatedKey)
			result.LastEvaluatedKey = newLastEvaluatedKey
		}

		if result.LastEvaluatedKey == nil || len(resultItems) >= int(limit) {
			lastEvaluatedKey = result.LastEvaluatedKey
			break
		}
		scanInput.ExclusiveStartKey = result.LastEvaluatedKey
	}

	return resultItems, lastEvaluatedKey, nil
}

func QueryDynamoDBUntilLimit(ctx context.Context, client awsclients.QueryInterface, queryInput *dynamodb.QueryInput, limit int32, keyNames []string) ([]map[string]types.AttributeValue, map[string]types.AttributeValue, error) {
	resultItems := []map[string]types.AttributeValue{}
	var lastEvaluatedKey map[string]types.AttributeValue

	for {
		result, err := client.Query(ctx, queryInput)
		if err != nil {
			return nil, nil, err
		}
		resultItems = append(resultItems, result.Items...)
		if len(resultItems) > int(limit) {
			fmt.Println(result.LastEvaluatedKey)
			resultItems = resultItems[:limit]
			lastEvaluatedItem := resultItems[len(resultItems)-1]
			newLastEvaluatedKey := map[string]types.AttributeValue{}
			for _, key := range keyNames {
				newLastEvaluatedKey[key] = lastEvaluatedItem[key]
			}
			fmt.Println(newLastEvaluatedKey)
			result.LastEvaluatedKey = newLastEvaluatedKey
		}

		if result.LastEvaluatedKey == nil || len(resultItems) >= int(limit) {
			lastEvaluatedKey = result.LastEvaluatedKey
			break
		}
		queryInput.ExclusiveStartKey = result.LastEvaluatedKey
	}

	return resultItems, lastEvaluatedKey, nil
}
