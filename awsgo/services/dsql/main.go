package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/dsql/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-cluster", "delete-cluster", "delete-cluster-policy", "get-cluster", "get-cluster-policy", "get-vpc-endpoint-service-name", "list-clusters", "list-tags-for-resource", "put-cluster-policy", "tag-resource", "untag-resource", "update-cluster"},
		OperationSet: map[string]bool{"create-cluster": true, "delete-cluster": true, "delete-cluster-policy": true, "get-cluster": true, "get-cluster-policy": true, "get-vpc-endpoint-service-name": true, "list-clusters": true, "list-tags-for-resource": true, "put-cluster-policy": true, "tag-resource": true, "untag-resource": true, "update-cluster": true},
		OperationInputs: map[string][]string{
			"create-cluster":                {"BypassPolicyLockoutSafetyCheck", "ClientToken", "DeletionProtectionEnabled", "KmsEncryptionKey", "MultiRegionProperties", "Policy", "Tags"},
			"delete-cluster":                {"ClientToken", "Identifier"},
			"delete-cluster-policy":         {"ClientToken", "ExpectedPolicyVersion", "Identifier"},
			"get-cluster":                   {"Identifier"},
			"get-cluster-policy":            {"Identifier"},
			"get-vpc-endpoint-service-name": {"Identifier"},
			"list-clusters":                 {"MaxResults", "NextToken"},
			"list-tags-for-resource":        {"ResourceArn"},
			"put-cluster-policy":            {"BypassPolicyLockoutSafetyCheck", "ClientToken", "ExpectedPolicyVersion", "Identifier", "Policy"},
			"tag-resource":                  {"ResourceArn", "Tags"},
			"untag-resource":                {"ResourceArn", "TagKeys"},
			"update-cluster":                {"ClientToken", "DeletionProtectionEnabled", "Identifier", "KmsEncryptionKey", "MultiRegionProperties"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-cluster":                {"BypassPolicyLockoutSafetyCheck": "bool", "ClientToken": "*string", "DeletionProtectionEnabled": "*bool", "KmsEncryptionKey": "*string", "MultiRegionProperties": "*types.MultiRegionProperties", "Policy": "*string", "Tags": "map[string]string"},
			"delete-cluster":                {"ClientToken": "*string", "Identifier": "*string"},
			"delete-cluster-policy":         {"ClientToken": "*string", "ExpectedPolicyVersion": "*string", "Identifier": "*string"},
			"get-cluster":                   {"Identifier": "*string"},
			"get-cluster-policy":            {"Identifier": "*string"},
			"get-vpc-endpoint-service-name": {"Identifier": "*string"},
			"list-clusters":                 {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":        {"ResourceArn": "*string"},
			"put-cluster-policy":            {"BypassPolicyLockoutSafetyCheck": "bool", "ClientToken": "*string", "ExpectedPolicyVersion": "*string", "Identifier": "*string", "Policy": "*string"},
			"tag-resource":                  {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":                {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-cluster":                {"ClientToken": "*string", "DeletionProtectionEnabled": "*bool", "Identifier": "*string", "KmsEncryptionKey": "*string", "MultiRegionProperties": "*types.MultiRegionProperties"},
		},
		OperationInputRequired: map[string][]string{
			"create-cluster":                {},
			"delete-cluster":                {"Identifier"},
			"delete-cluster-policy":         {"Identifier"},
			"get-cluster":                   {"Identifier"},
			"get-cluster-policy":            {"Identifier"},
			"get-vpc-endpoint-service-name": {"Identifier"},
			"list-clusters":                 {},
			"list-tags-for-resource":        {"ResourceArn"},
			"put-cluster-policy":            {"Identifier", "Policy"},
			"tag-resource":                  {"ResourceArn", "Tags"},
			"untag-resource":                {"ResourceArn", "TagKeys"},
			"update-cluster":                {"Identifier"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("dsql", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
