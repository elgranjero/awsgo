package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/signin/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-oauth2-token"},
		OperationSet: map[string]bool{"create-oauth2-token": true},
		OperationInputs: map[string][]string{
			"create-oauth2-token": {"TokenInput"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-oauth2-token": {"TokenInput": "*types.CreateOAuth2TokenRequestBody"},
		},
		OperationInputRequired: map[string][]string{
			"create-oauth2-token": {"TokenInput"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("signin", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
