package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/keyspacesstreams/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"get-records", "get-shard-iterator", "get-stream", "list-streams"},
		OperationSet: map[string]bool{"get-records": true, "get-shard-iterator": true, "get-stream": true, "list-streams": true},
		OperationInputs: map[string][]string{
			"get-records":        {"MaxResults", "ShardIterator"},
			"get-shard-iterator": {"SequenceNumber", "ShardId", "ShardIteratorType", "StreamArn"},
			"get-stream":         {"MaxResults", "NextToken", "ShardFilter", "StreamArn"},
			"list-streams":       {"KeyspaceName", "MaxResults", "NextToken", "TableName"},
		},
		OperationInputTypes: map[string]map[string]string{
			"get-records":        {"MaxResults": "*int32", "ShardIterator": "*string"},
			"get-shard-iterator": {"SequenceNumber": "*string", "ShardId": "*string", "ShardIteratorType": "types.ShardIteratorType", "StreamArn": "*string"},
			"get-stream":         {"MaxResults": "*int32", "NextToken": "*string", "ShardFilter": "*types.ShardFilter", "StreamArn": "*string"},
			"list-streams":       {"KeyspaceName": "*string", "MaxResults": "*int32", "NextToken": "*string", "TableName": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"get-records":        {"ShardIterator"},
			"get-shard-iterator": {"ShardId", "ShardIteratorType", "StreamArn"},
			"get-stream":         {"StreamArn"},
			"list-streams":       {},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("keyspacesstreams", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
