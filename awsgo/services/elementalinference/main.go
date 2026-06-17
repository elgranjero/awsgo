package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/elementalinference/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"associate-feed", "create-feed", "delete-feed", "disassociate-feed", "get-feed", "list-feeds", "list-tags-for-resource", "tag-resource", "untag-resource", "update-feed"},
		OperationSet: map[string]bool{"associate-feed": true, "create-feed": true, "delete-feed": true, "disassociate-feed": true, "get-feed": true, "list-feeds": true, "list-tags-for-resource": true, "tag-resource": true, "untag-resource": true, "update-feed": true},
		OperationInputs: map[string][]string{
			"associate-feed":         {"AssociatedResourceName", "DryRun", "Id", "Outputs"},
			"create-feed":            {"Name", "Outputs", "Tags"},
			"delete-feed":            {"Id"},
			"disassociate-feed":      {"AssociatedResourceName", "DryRun", "Id"},
			"get-feed":               {"Id"},
			"list-feeds":             {"MaxResults", "NextToken"},
			"list-tags-for-resource": {"ResourceArn"},
			"tag-resource":           {"ResourceArn", "Tags"},
			"untag-resource":         {"ResourceArn", "TagKeys"},
			"update-feed":            {"Id", "Name", "Outputs"},
		},
		OperationInputTypes: map[string]map[string]string{
			"associate-feed":         {"AssociatedResourceName": "*string", "DryRun": "bool", "Id": "*string", "Outputs": "[]types.CreateOutput"},
			"create-feed":            {"Name": "*string", "Outputs": "[]types.CreateOutput", "Tags": "map[string]string"},
			"delete-feed":            {"Id": "*string"},
			"disassociate-feed":      {"AssociatedResourceName": "*string", "DryRun": "bool", "Id": "*string"},
			"get-feed":               {"Id": "*string"},
			"list-feeds":             {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource": {"ResourceArn": "*string"},
			"tag-resource":           {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":         {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-feed":            {"Id": "*string", "Name": "*string", "Outputs": "[]types.UpdateOutput"},
		},
		OperationInputRequired: map[string][]string{
			"associate-feed":         {"AssociatedResourceName", "Id", "Outputs"},
			"create-feed":            {"Name", "Outputs"},
			"delete-feed":            {"Id"},
			"disassociate-feed":      {"AssociatedResourceName", "Id"},
			"get-feed":               {"Id"},
			"list-feeds":             {},
			"list-tags-for-resource": {"ResourceArn"},
			"tag-resource":           {"ResourceArn", "Tags"},
			"untag-resource":         {"ResourceArn", "TagKeys"},
			"update-feed":            {"Id", "Name", "Outputs"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("elementalinference", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
