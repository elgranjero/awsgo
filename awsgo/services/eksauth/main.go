package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/eksauth/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"assume-role-for-pod-identity"},
		OperationSet: map[string]bool{"assume-role-for-pod-identity": true},
		OperationInputs: map[string][]string{
			"assume-role-for-pod-identity": {"ClusterName", "Token"},
		},
		OperationInputTypes: map[string]map[string]string{
			"assume-role-for-pod-identity": {"ClusterName": "*string", "Token": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"assume-role-for-pod-identity": {"ClusterName", "Token"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("eksauth", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
