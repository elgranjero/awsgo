package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/sso/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"get-role-credentials", "list-account-roles", "list-accounts", "logout"},
		OperationSet: map[string]bool{"get-role-credentials": true, "list-account-roles": true, "list-accounts": true, "logout": true},
		OperationInputs: map[string][]string{
			"get-role-credentials": {"AccessToken", "AccountId", "RoleName"},
			"list-account-roles":   {"AccessToken", "AccountId", "MaxResults", "NextToken"},
			"list-accounts":        {"AccessToken", "MaxResults", "NextToken"},
			"logout":               {"AccessToken"},
		},
		OperationInputTypes: map[string]map[string]string{
			"get-role-credentials": {"AccessToken": "*string", "AccountId": "*string", "RoleName": "*string"},
			"list-account-roles":   {"AccessToken": "*string", "AccountId": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-accounts":        {"AccessToken": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"logout":               {"AccessToken": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"get-role-credentials": {"AccessToken", "AccountId", "RoleName"},
			"list-account-roles":   {"AccessToken", "AccountId"},
			"list-accounts":        {"AccessToken"},
			"logout":               {"AccessToken"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("sso", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
