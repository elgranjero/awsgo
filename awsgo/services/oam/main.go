package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/oam/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-link", "create-sink", "delete-link", "delete-sink", "get-link", "get-sink", "get-sink-policy", "list-attached-links", "list-links", "list-sinks", "list-tags-for-resource", "put-sink-policy", "tag-resource", "untag-resource", "update-link"},
		OperationSet: map[string]bool{"create-link": true, "create-sink": true, "delete-link": true, "delete-sink": true, "get-link": true, "get-sink": true, "get-sink-policy": true, "list-attached-links": true, "list-links": true, "list-sinks": true, "list-tags-for-resource": true, "put-sink-policy": true, "tag-resource": true, "untag-resource": true, "update-link": true},
		OperationInputs: map[string][]string{
			"create-link":            {"LabelTemplate", "LinkConfiguration", "ResourceTypes", "SinkIdentifier", "Tags"},
			"create-sink":            {"Name", "Tags"},
			"delete-link":            {"Identifier"},
			"delete-sink":            {"Identifier"},
			"get-link":               {"Identifier", "IncludeTags"},
			"get-sink":               {"Identifier", "IncludeTags"},
			"get-sink-policy":        {"SinkIdentifier"},
			"list-attached-links":    {"MaxResults", "NextToken", "SinkIdentifier"},
			"list-links":             {"MaxResults", "NextToken"},
			"list-sinks":             {"MaxResults", "NextToken"},
			"list-tags-for-resource": {"ResourceArn"},
			"put-sink-policy":        {"Policy", "SinkIdentifier"},
			"tag-resource":           {"ResourceArn", "Tags"},
			"untag-resource":         {"ResourceArn", "TagKeys"},
			"update-link":            {"Identifier", "IncludeTags", "LinkConfiguration", "ResourceTypes"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-link":            {"LabelTemplate": "*string", "LinkConfiguration": "*types.LinkConfiguration", "ResourceTypes": "[]types.ResourceType", "SinkIdentifier": "*string", "Tags": "map[string]string"},
			"create-sink":            {"Name": "*string", "Tags": "map[string]string"},
			"delete-link":            {"Identifier": "*string"},
			"delete-sink":            {"Identifier": "*string"},
			"get-link":               {"Identifier": "*string", "IncludeTags": "*bool"},
			"get-sink":               {"Identifier": "*string", "IncludeTags": "*bool"},
			"get-sink-policy":        {"SinkIdentifier": "*string"},
			"list-attached-links":    {"MaxResults": "*int32", "NextToken": "*string", "SinkIdentifier": "*string"},
			"list-links":             {"MaxResults": "*int32", "NextToken": "*string"},
			"list-sinks":             {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource": {"ResourceArn": "*string"},
			"put-sink-policy":        {"Policy": "*string", "SinkIdentifier": "*string"},
			"tag-resource":           {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":         {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-link":            {"Identifier": "*string", "IncludeTags": "*bool", "LinkConfiguration": "*types.LinkConfiguration", "ResourceTypes": "[]types.ResourceType"},
		},
		OperationInputRequired: map[string][]string{
			"create-link":            {"LabelTemplate", "ResourceTypes", "SinkIdentifier"},
			"create-sink":            {"Name"},
			"delete-link":            {"Identifier"},
			"delete-sink":            {"Identifier"},
			"get-link":               {"Identifier"},
			"get-sink":               {"Identifier"},
			"get-sink-policy":        {"SinkIdentifier"},
			"list-attached-links":    {"SinkIdentifier"},
			"list-links":             {},
			"list-sinks":             {},
			"list-tags-for-resource": {"ResourceArn"},
			"put-sink-policy":        {"Policy", "SinkIdentifier"},
			"tag-resource":           {"ResourceArn", "Tags"},
			"untag-resource":         {"ResourceArn", "TagKeys"},
			"update-link":            {"Identifier", "ResourceTypes"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("oam", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
