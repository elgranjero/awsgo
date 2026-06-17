package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/bcmdashboards/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-dashboard", "delete-dashboard", "get-dashboard", "get-resource-policy", "list-dashboards", "list-tags-for-resource", "tag-resource", "untag-resource", "update-dashboard"},
		OperationSet: map[string]bool{"create-dashboard": true, "delete-dashboard": true, "get-dashboard": true, "get-resource-policy": true, "list-dashboards": true, "list-tags-for-resource": true, "tag-resource": true, "untag-resource": true, "update-dashboard": true},
		OperationInputs: map[string][]string{
			"create-dashboard":       {"Description", "Name", "ResourceTags", "Widgets"},
			"delete-dashboard":       {"Arn"},
			"get-dashboard":          {"Arn"},
			"get-resource-policy":    {"ResourceArn"},
			"list-dashboards":        {"MaxResults", "NextToken"},
			"list-tags-for-resource": {"ResourceArn"},
			"tag-resource":           {"ResourceArn", "ResourceTags"},
			"untag-resource":         {"ResourceArn", "ResourceTagKeys"},
			"update-dashboard":       {"Arn", "Description", "Name", "Widgets"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-dashboard":       {"Description": "*string", "Name": "*string", "ResourceTags": "[]types.ResourceTag", "Widgets": "[]types.Widget"},
			"delete-dashboard":       {"Arn": "*string"},
			"get-dashboard":          {"Arn": "*string"},
			"get-resource-policy":    {"ResourceArn": "*string"},
			"list-dashboards":        {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource": {"ResourceArn": "*string"},
			"tag-resource":           {"ResourceArn": "*string", "ResourceTags": "[]types.ResourceTag"},
			"untag-resource":         {"ResourceArn": "*string", "ResourceTagKeys": "[]string"},
			"update-dashboard":       {"Arn": "*string", "Description": "*string", "Name": "*string", "Widgets": "[]types.Widget"},
		},
		OperationInputRequired: map[string][]string{
			"create-dashboard":       {"Name", "Widgets"},
			"delete-dashboard":       {"Arn"},
			"get-dashboard":          {"Arn"},
			"get-resource-policy":    {"ResourceArn"},
			"list-dashboards":        {},
			"list-tags-for-resource": {"ResourceArn"},
			"tag-resource":           {"ResourceArn", "ResourceTags"},
			"untag-resource":         {"ResourceArn", "ResourceTagKeys"},
			"update-dashboard":       {"Arn"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("bcmdashboards", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
