package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/marketplacereporting/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"get-buyer-dashboard"},
		OperationSet: map[string]bool{"get-buyer-dashboard": true},
		OperationInputs: map[string][]string{
			"get-buyer-dashboard": {"DashboardIdentifier", "EmbeddingDomains"},
		},
		OperationInputTypes: map[string]map[string]string{
			"get-buyer-dashboard": {"DashboardIdentifier": "*string", "EmbeddingDomains": "[]string"},
		},
		OperationInputRequired: map[string][]string{
			"get-buyer-dashboard": {"DashboardIdentifier", "EmbeddingDomains"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("marketplacereporting", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
