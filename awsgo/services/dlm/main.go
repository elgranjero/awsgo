package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/dlm/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-lifecycle-policy", "delete-lifecycle-policy", "get-lifecycle-policies", "get-lifecycle-policy", "list-tags-for-resource", "tag-resource", "untag-resource", "update-lifecycle-policy"},
		OperationSet: map[string]bool{"create-lifecycle-policy": true, "delete-lifecycle-policy": true, "get-lifecycle-policies": true, "get-lifecycle-policy": true, "list-tags-for-resource": true, "tag-resource": true, "untag-resource": true, "update-lifecycle-policy": true},
		OperationInputs: map[string][]string{
			"create-lifecycle-policy": {"CopyTags", "CreateInterval", "CrossRegionCopyTargets", "DefaultPolicy", "Description", "Exclusions", "ExecutionRoleArn", "ExtendDeletion", "PolicyDetails", "RetainInterval", "State", "Tags"},
			"delete-lifecycle-policy": {"PolicyId"},
			"get-lifecycle-policies":  {"DefaultPolicyType", "PolicyIds", "ResourceTypes", "State", "TagsToAdd", "TargetTags"},
			"get-lifecycle-policy":    {"PolicyId"},
			"list-tags-for-resource":  {"ResourceArn"},
			"tag-resource":            {"ResourceArn", "Tags"},
			"untag-resource":          {"ResourceArn", "TagKeys"},
			"update-lifecycle-policy": {"CopyTags", "CreateInterval", "CrossRegionCopyTargets", "Description", "Exclusions", "ExecutionRoleArn", "ExtendDeletion", "PolicyDetails", "PolicyId", "RetainInterval", "State"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-lifecycle-policy": {"CopyTags": "*bool", "CreateInterval": "*int32", "CrossRegionCopyTargets": "[]types.CrossRegionCopyTarget", "DefaultPolicy": "types.DefaultPolicyTypeValues", "Description": "*string", "Exclusions": "*types.Exclusions", "ExecutionRoleArn": "*string", "ExtendDeletion": "*bool", "PolicyDetails": "*types.PolicyDetails", "RetainInterval": "*int32", "State": "types.SettablePolicyStateValues", "Tags": "map[string]string"},
			"delete-lifecycle-policy": {"PolicyId": "*string"},
			"get-lifecycle-policies":  {"DefaultPolicyType": "types.DefaultPoliciesTypeValues", "PolicyIds": "[]string", "ResourceTypes": "[]types.ResourceTypeValues", "State": "types.GettablePolicyStateValues", "TagsToAdd": "[]string", "TargetTags": "[]string"},
			"get-lifecycle-policy":    {"PolicyId": "*string"},
			"list-tags-for-resource":  {"ResourceArn": "*string"},
			"tag-resource":            {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":          {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-lifecycle-policy": {"CopyTags": "*bool", "CreateInterval": "*int32", "CrossRegionCopyTargets": "[]types.CrossRegionCopyTarget", "Description": "*string", "Exclusions": "*types.Exclusions", "ExecutionRoleArn": "*string", "ExtendDeletion": "*bool", "PolicyDetails": "*types.PolicyDetails", "PolicyId": "*string", "RetainInterval": "*int32", "State": "types.SettablePolicyStateValues"},
		},
		OperationInputRequired: map[string][]string{
			"create-lifecycle-policy": {"Description", "ExecutionRoleArn", "State"},
			"delete-lifecycle-policy": {"PolicyId"},
			"get-lifecycle-policies":  {},
			"get-lifecycle-policy":    {"PolicyId"},
			"list-tags-for-resource":  {"ResourceArn"},
			"tag-resource":            {"ResourceArn", "Tags"},
			"untag-resource":          {"ResourceArn", "TagKeys"},
			"update-lifecycle-policy": {"PolicyId"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("dlm", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
