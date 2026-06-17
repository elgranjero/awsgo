package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/mediapackagevod/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"configure-logs", "create-asset", "create-packaging-configuration", "create-packaging-group", "delete-asset", "delete-packaging-configuration", "delete-packaging-group", "describe-asset", "describe-packaging-configuration", "describe-packaging-group", "list-assets", "list-packaging-configurations", "list-packaging-groups", "list-tags-for-resource", "tag-resource", "untag-resource", "update-packaging-group"},
		OperationSet: map[string]bool{"configure-logs": true, "create-asset": true, "create-packaging-configuration": true, "create-packaging-group": true, "delete-asset": true, "delete-packaging-configuration": true, "delete-packaging-group": true, "describe-asset": true, "describe-packaging-configuration": true, "describe-packaging-group": true, "list-assets": true, "list-packaging-configurations": true, "list-packaging-groups": true, "list-tags-for-resource": true, "tag-resource": true, "untag-resource": true, "update-packaging-group": true},
		OperationInputs: map[string][]string{
			"configure-logs":                   {"EgressAccessLogs", "Id"},
			"create-asset":                     {"Id", "PackagingGroupId", "ResourceId", "SourceArn", "SourceRoleArn", "Tags"},
			"create-packaging-configuration":   {"CmafPackage", "DashPackage", "HlsPackage", "Id", "MssPackage", "PackagingGroupId", "Tags"},
			"create-packaging-group":           {"Authorization", "EgressAccessLogs", "Id", "Tags"},
			"delete-asset":                     {"Id"},
			"delete-packaging-configuration":   {"Id"},
			"delete-packaging-group":           {"Id"},
			"describe-asset":                   {"Id"},
			"describe-packaging-configuration": {"Id"},
			"describe-packaging-group":         {"Id"},
			"list-assets":                      {"MaxResults", "NextToken", "PackagingGroupId"},
			"list-packaging-configurations":    {"MaxResults", "NextToken", "PackagingGroupId"},
			"list-packaging-groups":            {"MaxResults", "NextToken"},
			"list-tags-for-resource":           {"ResourceArn"},
			"tag-resource":                     {"ResourceArn", "Tags"},
			"untag-resource":                   {"ResourceArn", "TagKeys"},
			"update-packaging-group":           {"Authorization", "Id"},
		},
		OperationInputTypes: map[string]map[string]string{
			"configure-logs":                   {"EgressAccessLogs": "*types.EgressAccessLogs", "Id": "*string"},
			"create-asset":                     {"Id": "*string", "PackagingGroupId": "*string", "ResourceId": "*string", "SourceArn": "*string", "SourceRoleArn": "*string", "Tags": "map[string]string"},
			"create-packaging-configuration":   {"CmafPackage": "*types.CmafPackage", "DashPackage": "*types.DashPackage", "HlsPackage": "*types.HlsPackage", "Id": "*string", "MssPackage": "*types.MssPackage", "PackagingGroupId": "*string", "Tags": "map[string]string"},
			"create-packaging-group":           {"Authorization": "*types.Authorization", "EgressAccessLogs": "*types.EgressAccessLogs", "Id": "*string", "Tags": "map[string]string"},
			"delete-asset":                     {"Id": "*string"},
			"delete-packaging-configuration":   {"Id": "*string"},
			"delete-packaging-group":           {"Id": "*string"},
			"describe-asset":                   {"Id": "*string"},
			"describe-packaging-configuration": {"Id": "*string"},
			"describe-packaging-group":         {"Id": "*string"},
			"list-assets":                      {"MaxResults": "*int32", "NextToken": "*string", "PackagingGroupId": "*string"},
			"list-packaging-configurations":    {"MaxResults": "*int32", "NextToken": "*string", "PackagingGroupId": "*string"},
			"list-packaging-groups":            {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":           {"ResourceArn": "*string"},
			"tag-resource":                     {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":                   {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-packaging-group":           {"Authorization": "*types.Authorization", "Id": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"configure-logs":                   {"Id"},
			"create-asset":                     {"Id", "PackagingGroupId", "SourceArn", "SourceRoleArn"},
			"create-packaging-configuration":   {"Id", "PackagingGroupId"},
			"create-packaging-group":           {"Id"},
			"delete-asset":                     {"Id"},
			"delete-packaging-configuration":   {"Id"},
			"delete-packaging-group":           {"Id"},
			"describe-asset":                   {"Id"},
			"describe-packaging-configuration": {"Id"},
			"describe-packaging-group":         {"Id"},
			"list-assets":                      {},
			"list-packaging-configurations":    {},
			"list-packaging-groups":            {},
			"list-tags-for-resource":           {"ResourceArn"},
			"tag-resource":                     {"ResourceArn", "Tags"},
			"untag-resource":                   {"ResourceArn", "TagKeys"},
			"update-packaging-group":           {"Id"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("mediapackagevod", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
