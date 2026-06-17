package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/dynamodbstreams/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"describe-stream", "get-records", "get-shard-iterator", "list-streams"},
		OperationSet: map[string]bool{"describe-stream": true, "get-records": true, "get-shard-iterator": true, "list-streams": true},
		OperationInputs: map[string][]string{
			"describe-stream":    {"ExclusiveStartShardId", "Limit", "ShardFilter", "StreamArn"},
			"get-records":        {"Limit", "ShardIterator"},
			"get-shard-iterator": {"SequenceNumber", "ShardId", "ShardIteratorType", "StreamArn"},
			"list-streams":       {"ExclusiveStartStreamArn", "Limit", "TableName"},
		},
		OperationInputTypes: map[string]map[string]string{
			"describe-stream":    {"ExclusiveStartShardId": "*string", "Limit": "*int32", "ShardFilter": "*types.ShardFilter", "StreamArn": "*string"},
			"get-records":        {"Limit": "*int32", "ShardIterator": "*string"},
			"get-shard-iterator": {"SequenceNumber": "*string", "ShardId": "*string", "ShardIteratorType": "types.ShardIteratorType", "StreamArn": "*string"},
			"list-streams":       {"ExclusiveStartStreamArn": "*string", "Limit": "*int32", "TableName": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"describe-stream":    {"StreamArn"},
			"get-records":        {"ShardIterator"},
			"get-shard-iterator": {"ShardId", "ShardIteratorType", "StreamArn"},
			"list-streams":       {},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("dynamodbstreams", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
