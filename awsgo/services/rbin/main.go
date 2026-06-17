package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/rbin/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-rule", "delete-rule", "get-rule", "list-rules", "list-tags-for-resource", "lock-rule", "tag-resource", "unlock-rule", "untag-resource", "update-rule"},
		OperationSet: map[string]bool{"create-rule": true, "delete-rule": true, "get-rule": true, "list-rules": true, "list-tags-for-resource": true, "lock-rule": true, "tag-resource": true, "unlock-rule": true, "untag-resource": true, "update-rule": true},
		OperationInputs: map[string][]string{
			"create-rule":            {"Description", "ExcludeResourceTags", "LockConfiguration", "ResourceTags", "ResourceType", "RetentionPeriod", "Tags"},
			"delete-rule":            {"Identifier"},
			"get-rule":               {"Identifier"},
			"list-rules":             {"ExcludeResourceTags", "LockState", "MaxResults", "NextToken", "ResourceTags", "ResourceType"},
			"list-tags-for-resource": {"ResourceArn"},
			"lock-rule":              {"Identifier", "LockConfiguration"},
			"tag-resource":           {"ResourceArn", "Tags"},
			"unlock-rule":            {"Identifier"},
			"untag-resource":         {"ResourceArn", "TagKeys"},
			"update-rule":            {"Description", "ExcludeResourceTags", "Identifier", "ResourceTags", "ResourceType", "RetentionPeriod"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-rule":            {"Description": "*string", "ExcludeResourceTags": "[]types.ResourceTag", "LockConfiguration": "*types.LockConfiguration", "ResourceTags": "[]types.ResourceTag", "ResourceType": "types.ResourceType", "RetentionPeriod": "*types.RetentionPeriod", "Tags": "[]types.Tag"},
			"delete-rule":            {"Identifier": "*string"},
			"get-rule":               {"Identifier": "*string"},
			"list-rules":             {"ExcludeResourceTags": "[]types.ResourceTag", "LockState": "types.LockState", "MaxResults": "*int32", "NextToken": "*string", "ResourceTags": "[]types.ResourceTag", "ResourceType": "types.ResourceType"},
			"list-tags-for-resource": {"ResourceArn": "*string"},
			"lock-rule":              {"Identifier": "*string", "LockConfiguration": "*types.LockConfiguration"},
			"tag-resource":           {"ResourceArn": "*string", "Tags": "[]types.Tag"},
			"unlock-rule":            {"Identifier": "*string"},
			"untag-resource":         {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-rule":            {"Description": "*string", "ExcludeResourceTags": "[]types.ResourceTag", "Identifier": "*string", "ResourceTags": "[]types.ResourceTag", "ResourceType": "types.ResourceType", "RetentionPeriod": "*types.RetentionPeriod"},
		},
		OperationInputRequired: map[string][]string{
			"create-rule":            {"ResourceType", "RetentionPeriod"},
			"delete-rule":            {"Identifier"},
			"get-rule":               {"Identifier"},
			"list-rules":             {"ResourceType"},
			"list-tags-for-resource": {"ResourceArn"},
			"lock-rule":              {"Identifier", "LockConfiguration"},
			"tag-resource":           {"ResourceArn", "Tags"},
			"unlock-rule":            {"Identifier"},
			"untag-resource":         {"ResourceArn", "TagKeys"},
			"update-rule":            {"Identifier"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("rbin", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
