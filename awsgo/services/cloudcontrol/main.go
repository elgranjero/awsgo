package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/cloudcontrol/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"cancel-resource-request", "create-resource", "delete-resource", "get-resource", "get-resource-request-status", "list-resource-requests", "list-resources", "update-resource"},
		OperationSet: map[string]bool{"cancel-resource-request": true, "create-resource": true, "delete-resource": true, "get-resource": true, "get-resource-request-status": true, "list-resource-requests": true, "list-resources": true, "update-resource": true},
		OperationInputs: map[string][]string{
			"cancel-resource-request":     {"RequestToken"},
			"create-resource":             {"ClientToken", "DesiredState", "RoleArn", "TypeName", "TypeVersionId"},
			"delete-resource":             {"ClientToken", "Identifier", "RoleArn", "TypeName", "TypeVersionId"},
			"get-resource":                {"Identifier", "RoleArn", "TypeName", "TypeVersionId"},
			"get-resource-request-status": {"RequestToken"},
			"list-resource-requests":      {"MaxResults", "NextToken", "ResourceRequestStatusFilter"},
			"list-resources":              {"MaxResults", "NextToken", "ResourceModel", "RoleArn", "TypeName", "TypeVersionId"},
			"update-resource":             {"ClientToken", "Identifier", "PatchDocument", "RoleArn", "TypeName", "TypeVersionId"},
		},
		OperationInputTypes: map[string]map[string]string{
			"cancel-resource-request":     {"RequestToken": "*string"},
			"create-resource":             {"ClientToken": "*string", "DesiredState": "*string", "RoleArn": "*string", "TypeName": "*string", "TypeVersionId": "*string"},
			"delete-resource":             {"ClientToken": "*string", "Identifier": "*string", "RoleArn": "*string", "TypeName": "*string", "TypeVersionId": "*string"},
			"get-resource":                {"Identifier": "*string", "RoleArn": "*string", "TypeName": "*string", "TypeVersionId": "*string"},
			"get-resource-request-status": {"RequestToken": "*string"},
			"list-resource-requests":      {"MaxResults": "*int32", "NextToken": "*string", "ResourceRequestStatusFilter": "*types.ResourceRequestStatusFilter"},
			"list-resources":              {"MaxResults": "*int32", "NextToken": "*string", "ResourceModel": "*string", "RoleArn": "*string", "TypeName": "*string", "TypeVersionId": "*string"},
			"update-resource":             {"ClientToken": "*string", "Identifier": "*string", "PatchDocument": "*string", "RoleArn": "*string", "TypeName": "*string", "TypeVersionId": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"cancel-resource-request":     {"RequestToken"},
			"create-resource":             {"DesiredState", "TypeName"},
			"delete-resource":             {"Identifier", "TypeName"},
			"get-resource":                {"Identifier", "TypeName"},
			"get-resource-request-status": {"RequestToken"},
			"list-resource-requests":      {},
			"list-resources":              {"TypeName"},
			"update-resource":             {"Identifier", "PatchDocument", "TypeName"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("cloudcontrol", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
