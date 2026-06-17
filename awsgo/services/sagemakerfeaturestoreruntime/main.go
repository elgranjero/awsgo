package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/sagemakerfeaturestoreruntime/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"batch-get-record", "delete-record", "get-record", "put-record"},
		OperationSet: map[string]bool{"batch-get-record": true, "delete-record": true, "get-record": true, "put-record": true},
		OperationInputs: map[string][]string{
			"batch-get-record": {"ExpirationTimeResponse", "Identifiers"},
			"delete-record":    {"DeletionMode", "EventTime", "FeatureGroupName", "RecordIdentifierValueAsString", "TargetStores"},
			"get-record":       {"ExpirationTimeResponse", "FeatureGroupName", "FeatureNames", "RecordIdentifierValueAsString"},
			"put-record":       {"FeatureGroupName", "Record", "TargetStores", "TtlDuration"},
		},
		OperationInputTypes: map[string]map[string]string{
			"batch-get-record": {"ExpirationTimeResponse": "types.ExpirationTimeResponse", "Identifiers": "[]types.BatchGetRecordIdentifier"},
			"delete-record":    {"DeletionMode": "types.DeletionMode", "EventTime": "*string", "FeatureGroupName": "*string", "RecordIdentifierValueAsString": "*string", "TargetStores": "[]types.TargetStore"},
			"get-record":       {"ExpirationTimeResponse": "types.ExpirationTimeResponse", "FeatureGroupName": "*string", "FeatureNames": "[]string", "RecordIdentifierValueAsString": "*string"},
			"put-record":       {"FeatureGroupName": "*string", "Record": "[]types.FeatureValue", "TargetStores": "[]types.TargetStore", "TtlDuration": "*types.TtlDuration"},
		},
		OperationInputRequired: map[string][]string{
			"batch-get-record": {"Identifiers"},
			"delete-record":    {"EventTime", "FeatureGroupName", "RecordIdentifierValueAsString"},
			"get-record":       {"FeatureGroupName", "RecordIdentifierValueAsString"},
			"put-record":       {"FeatureGroupName", "Record"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("sagemakerfeaturestoreruntime", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
