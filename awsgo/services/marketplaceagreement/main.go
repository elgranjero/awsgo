package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/marketplaceagreement/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"describe-agreement", "get-agreement-terms", "search-agreements"},
		OperationSet: map[string]bool{"describe-agreement": true, "get-agreement-terms": true, "search-agreements": true},
		OperationInputs: map[string][]string{
			"describe-agreement":  {"AgreementId"},
			"get-agreement-terms": {"AgreementId", "MaxResults", "NextToken"},
			"search-agreements":   {"Catalog", "Filters", "MaxResults", "NextToken", "Sort"},
		},
		OperationInputTypes: map[string]map[string]string{
			"describe-agreement":  {"AgreementId": "*string"},
			"get-agreement-terms": {"AgreementId": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"search-agreements":   {"Catalog": "*string", "Filters": "[]types.Filter", "MaxResults": "*int32", "NextToken": "*string", "Sort": "*types.Sort"},
		},
		OperationInputRequired: map[string][]string{
			"describe-agreement":  {"AgreementId"},
			"get-agreement-terms": {"AgreementId"},
			"search-agreements":   {},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("marketplaceagreement", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
