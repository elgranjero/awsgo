package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/migrationhubconfig/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-home-region-control", "delete-home-region-control", "describe-home-region-controls", "get-home-region"},
		OperationSet: map[string]bool{"create-home-region-control": true, "delete-home-region-control": true, "describe-home-region-controls": true, "get-home-region": true},
		OperationInputs: map[string][]string{
			"create-home-region-control":    {"DryRun", "HomeRegion", "Target"},
			"delete-home-region-control":    {"ControlId"},
			"describe-home-region-controls": {"ControlId", "HomeRegion", "MaxResults", "NextToken", "Target"},
			"get-home-region":               {},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-home-region-control":    {"DryRun": "bool", "HomeRegion": "*string", "Target": "*types.Target"},
			"delete-home-region-control":    {"ControlId": "*string"},
			"describe-home-region-controls": {"ControlId": "*string", "HomeRegion": "*string", "MaxResults": "*int32", "NextToken": "*string", "Target": "*types.Target"},
			"get-home-region":               {},
		},
		OperationInputRequired: map[string][]string{
			"create-home-region-control":    {"HomeRegion", "Target"},
			"delete-home-region-control":    {"ControlId"},
			"describe-home-region-controls": {},
			"get-home-region":               {},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("migrationhubconfig", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
