package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/marketplaceentitlementservice/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"get-entitlements"},
		OperationSet: map[string]bool{"get-entitlements": true},
		OperationInputs: map[string][]string{
			"get-entitlements": {"Filter", "MaxResults", "NextToken", "ProductCode"},
		},
		OperationInputTypes: map[string]map[string]string{
			"get-entitlements": {"Filter": "map[string][]string", "MaxResults": "*int32", "NextToken": "*string", "ProductCode": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"get-entitlements": {"ProductCode"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("marketplaceentitlementservice", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
