package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/bcmrecommendedactions/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"list-recommended-actions"},
		OperationSet: map[string]bool{"list-recommended-actions": true},
		OperationInputs: map[string][]string{
			"list-recommended-actions": {"Filter", "MaxResults", "NextToken"},
		},
		OperationInputTypes: map[string]map[string]string{
			"list-recommended-actions": {"Filter": "*types.RequestFilter", "MaxResults": "*int32", "NextToken": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"list-recommended-actions": {},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("bcmrecommendedactions", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
