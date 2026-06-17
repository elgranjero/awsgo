package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/cloudfrontkeyvaluestore/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"delete-key", "describe-key-value-store", "get-key", "list-keys", "put-key", "update-keys"},
		OperationSet: map[string]bool{"delete-key": true, "describe-key-value-store": true, "get-key": true, "list-keys": true, "put-key": true, "update-keys": true},
		OperationInputs: map[string][]string{
			"delete-key":               {"IfMatch", "Key", "KvsARN"},
			"describe-key-value-store": {"KvsARN"},
			"get-key":                  {"Key", "KvsARN"},
			"list-keys":                {"KvsARN", "MaxResults", "NextToken"},
			"put-key":                  {"IfMatch", "Key", "KvsARN", "Value"},
			"update-keys":              {"Deletes", "IfMatch", "KvsARN", "Puts"},
		},
		OperationInputTypes: map[string]map[string]string{
			"delete-key":               {"IfMatch": "*string", "Key": "*string", "KvsARN": "*string"},
			"describe-key-value-store": {"KvsARN": "*string"},
			"get-key":                  {"Key": "*string", "KvsARN": "*string"},
			"list-keys":                {"KvsARN": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"put-key":                  {"IfMatch": "*string", "Key": "*string", "KvsARN": "*string", "Value": "*string"},
			"update-keys":              {"Deletes": "[]types.DeleteKeyRequestListItem", "IfMatch": "*string", "KvsARN": "*string", "Puts": "[]types.PutKeyRequestListItem"},
		},
		OperationInputRequired: map[string][]string{
			"delete-key":               {"IfMatch", "Key", "KvsARN"},
			"describe-key-value-store": {"KvsARN"},
			"get-key":                  {"Key", "KvsARN"},
			"list-keys":                {"KvsARN"},
			"put-key":                  {"IfMatch", "Key", "KvsARN", "Value"},
			"update-keys":              {"IfMatch", "KvsARN"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("cloudfrontkeyvaluestore", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
