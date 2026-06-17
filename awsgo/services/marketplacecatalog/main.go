package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/marketplacecatalog/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"batch-describe-entities", "cancel-change-set", "delete-resource-policy", "describe-change-set", "describe-entity", "get-resource-policy", "list-change-sets", "list-entities", "list-tags-for-resource", "put-resource-policy", "start-change-set", "tag-resource", "untag-resource"},
		OperationSet: map[string]bool{"batch-describe-entities": true, "cancel-change-set": true, "delete-resource-policy": true, "describe-change-set": true, "describe-entity": true, "get-resource-policy": true, "list-change-sets": true, "list-entities": true, "list-tags-for-resource": true, "put-resource-policy": true, "start-change-set": true, "tag-resource": true, "untag-resource": true},
		OperationInputs: map[string][]string{
			"batch-describe-entities": {"EntityRequestList"},
			"cancel-change-set":       {"Catalog", "ChangeSetId"},
			"delete-resource-policy":  {"ResourceArn"},
			"describe-change-set":     {"Catalog", "ChangeSetId"},
			"describe-entity":         {"Catalog", "EntityId"},
			"get-resource-policy":     {"ResourceArn"},
			"list-change-sets":        {"Catalog", "FilterList", "MaxResults", "NextToken", "Sort"},
			"list-entities":           {"Catalog", "EntityType", "EntityTypeFilters", "EntityTypeSort", "FilterList", "MaxResults", "NextToken", "OwnershipType", "Sort"},
			"list-tags-for-resource":  {"ResourceArn"},
			"put-resource-policy":     {"Policy", "ResourceArn"},
			"start-change-set":        {"Catalog", "ChangeSet", "ChangeSetName", "ChangeSetTags", "ClientRequestToken", "Intent"},
			"tag-resource":            {"ResourceArn", "Tags"},
			"untag-resource":          {"ResourceArn", "TagKeys"},
		},
		OperationInputTypes: map[string]map[string]string{
			"batch-describe-entities": {"EntityRequestList": "[]types.EntityRequest"},
			"cancel-change-set":       {"Catalog": "*string", "ChangeSetId": "*string"},
			"delete-resource-policy":  {"ResourceArn": "*string"},
			"describe-change-set":     {"Catalog": "*string", "ChangeSetId": "*string"},
			"describe-entity":         {"Catalog": "*string", "EntityId": "*string"},
			"get-resource-policy":     {"ResourceArn": "*string"},
			"list-change-sets":        {"Catalog": "*string", "FilterList": "[]types.Filter", "MaxResults": "*int32", "NextToken": "*string", "Sort": "*types.Sort"},
			"list-entities":           {"Catalog": "*string", "EntityType": "*string", "EntityTypeFilters": "types.EntityTypeFilters", "EntityTypeSort": "types.EntityTypeSort", "FilterList": "[]types.Filter", "MaxResults": "*int32", "NextToken": "*string", "OwnershipType": "types.OwnershipType", "Sort": "*types.Sort"},
			"list-tags-for-resource":  {"ResourceArn": "*string"},
			"put-resource-policy":     {"Policy": "*string", "ResourceArn": "*string"},
			"start-change-set":        {"Catalog": "*string", "ChangeSet": "[]types.Change", "ChangeSetName": "*string", "ChangeSetTags": "[]types.Tag", "ClientRequestToken": "*string", "Intent": "types.Intent"},
			"tag-resource":            {"ResourceArn": "*string", "Tags": "[]types.Tag"},
			"untag-resource":          {"ResourceArn": "*string", "TagKeys": "[]string"},
		},
		OperationInputRequired: map[string][]string{
			"batch-describe-entities": {"EntityRequestList"},
			"cancel-change-set":       {"Catalog", "ChangeSetId"},
			"delete-resource-policy":  {"ResourceArn"},
			"describe-change-set":     {"Catalog", "ChangeSetId"},
			"describe-entity":         {"Catalog", "EntityId"},
			"get-resource-policy":     {"ResourceArn"},
			"list-change-sets":        {"Catalog"},
			"list-entities":           {"Catalog", "EntityType"},
			"list-tags-for-resource":  {"ResourceArn"},
			"put-resource-policy":     {"Policy", "ResourceArn"},
			"start-change-set":        {"Catalog", "ChangeSet"},
			"tag-resource":            {"ResourceArn", "Tags"},
			"untag-resource":          {"ResourceArn", "TagKeys"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("marketplacecatalog", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
