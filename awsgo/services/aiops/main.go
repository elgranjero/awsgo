package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/aiops/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-investigation-group", "delete-investigation-group", "delete-investigation-group-policy", "get-investigation-group", "get-investigation-group-policy", "list-investigation-groups", "list-tags-for-resource", "put-investigation-group-policy", "tag-resource", "untag-resource", "update-investigation-group"},
		OperationSet: map[string]bool{"create-investigation-group": true, "delete-investigation-group": true, "delete-investigation-group-policy": true, "get-investigation-group": true, "get-investigation-group-policy": true, "list-investigation-groups": true, "list-tags-for-resource": true, "put-investigation-group-policy": true, "tag-resource": true, "untag-resource": true, "update-investigation-group": true},
		OperationInputs: map[string][]string{
			"create-investigation-group":        {"ChatbotNotificationChannel", "CrossAccountConfigurations", "EncryptionConfiguration", "IsCloudTrailEventHistoryEnabled", "Name", "RetentionInDays", "RoleArn", "TagKeyBoundaries", "Tags"},
			"delete-investigation-group":        {"Identifier"},
			"delete-investigation-group-policy": {"Identifier"},
			"get-investigation-group":           {"Identifier"},
			"get-investigation-group-policy":    {"Identifier"},
			"list-investigation-groups":         {"MaxResults", "NextToken"},
			"list-tags-for-resource":            {"ResourceArn"},
			"put-investigation-group-policy":    {"Identifier", "Policy"},
			"tag-resource":                      {"ResourceArn", "Tags"},
			"untag-resource":                    {"ResourceArn", "TagKeys"},
			"update-investigation-group":        {"ChatbotNotificationChannel", "CrossAccountConfigurations", "EncryptionConfiguration", "Identifier", "IsCloudTrailEventHistoryEnabled", "RoleArn", "TagKeyBoundaries"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-investigation-group":        {"ChatbotNotificationChannel": "map[string][]string", "CrossAccountConfigurations": "[]types.CrossAccountConfiguration", "EncryptionConfiguration": "*types.EncryptionConfiguration", "IsCloudTrailEventHistoryEnabled": "*bool", "Name": "*string", "RetentionInDays": "*int64", "RoleArn": "*string", "TagKeyBoundaries": "[]string", "Tags": "map[string]string"},
			"delete-investigation-group":        {"Identifier": "*string"},
			"delete-investigation-group-policy": {"Identifier": "*string"},
			"get-investigation-group":           {"Identifier": "*string"},
			"get-investigation-group-policy":    {"Identifier": "*string"},
			"list-investigation-groups":         {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":            {"ResourceArn": "*string"},
			"put-investigation-group-policy":    {"Identifier": "*string", "Policy": "*string"},
			"tag-resource":                      {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":                    {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-investigation-group":        {"ChatbotNotificationChannel": "map[string][]string", "CrossAccountConfigurations": "[]types.CrossAccountConfiguration", "EncryptionConfiguration": "*types.EncryptionConfiguration", "Identifier": "*string", "IsCloudTrailEventHistoryEnabled": "*bool", "RoleArn": "*string", "TagKeyBoundaries": "[]string"},
		},
		OperationInputRequired: map[string][]string{
			"create-investigation-group":        {"Name", "RoleArn"},
			"delete-investigation-group":        {"Identifier"},
			"delete-investigation-group-policy": {"Identifier"},
			"get-investigation-group":           {"Identifier"},
			"get-investigation-group-policy":    {"Identifier"},
			"list-investigation-groups":         {},
			"list-tags-for-resource":            {"ResourceArn"},
			"put-investigation-group-policy":    {"Identifier", "Policy"},
			"tag-resource":                      {"ResourceArn", "Tags"},
			"untag-resource":                    {"ResourceArn", "TagKeys"},
			"update-investigation-group":        {"Identifier"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("aiops", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
