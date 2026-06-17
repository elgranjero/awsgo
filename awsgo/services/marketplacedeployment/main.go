package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/marketplacedeployment/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"list-tags-for-resource", "put-deployment-parameter", "tag-resource", "untag-resource"},
		OperationSet: map[string]bool{"list-tags-for-resource": true, "put-deployment-parameter": true, "tag-resource": true, "untag-resource": true},
		OperationInputs: map[string][]string{
			"list-tags-for-resource":   {"ResourceArn"},
			"put-deployment-parameter": {"AgreementId", "Catalog", "ClientToken", "DeploymentParameter", "ExpirationDate", "ProductId", "Tags"},
			"tag-resource":             {"ResourceArn", "Tags"},
			"untag-resource":           {"ResourceArn", "TagKeys"},
		},
		OperationInputTypes: map[string]map[string]string{
			"list-tags-for-resource":   {"ResourceArn": "*string"},
			"put-deployment-parameter": {"AgreementId": "*string", "Catalog": "*string", "ClientToken": "*string", "DeploymentParameter": "*types.DeploymentParameterInput", "ExpirationDate": "*time.Time", "ProductId": "*string", "Tags": "map[string]string"},
			"tag-resource":             {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":           {"ResourceArn": "*string", "TagKeys": "[]string"},
		},
		OperationInputRequired: map[string][]string{
			"list-tags-for-resource":   {"ResourceArn"},
			"put-deployment-parameter": {"AgreementId", "Catalog", "DeploymentParameter", "ProductId"},
			"tag-resource":             {"ResourceArn"},
			"untag-resource":           {"ResourceArn", "TagKeys"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("marketplacedeployment", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
