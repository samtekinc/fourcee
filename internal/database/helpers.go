package database

import "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

func SortDynamoDBBatchResponses(keys []map[string]types.AttributeValue, results []map[string]types.AttributeValue) []map[string]types.AttributeValue {
	if len(keys) == 0 {
		return []map[string]types.AttributeValue{}
	}
	// determine the primary key names
	key := keys[0]
	primaryKeys := make([]string, 0, len(key))
	for k := range key {
		primaryKeys = append(primaryKeys, k)
	}

	// build a mapping of keys to results
	mapping := make(map[string]map[string]types.AttributeValue, len(results))
	for i, result := range results {
		key := ""
		for _, k := range primaryKeys {
			key += result[k].(*types.AttributeValueMemberS).Value + "|"
		}
		mapping[key] = results[i]
	}

	// build the output by looking up the keys in the mapping and inserting them in the original order
	output := make([]map[string]types.AttributeValue, 0, len(keys))
	for _, key := range keys {
		keyString := ""
		for _, k := range primaryKeys {
			keyString += key[k].(*types.AttributeValueMemberS).Value + "|"
		}
		if result, ok := mapping[keyString]; ok {
			output = append(output, result)
		} else {
			output = append(output, nil)
		}
	}

	return output
}
