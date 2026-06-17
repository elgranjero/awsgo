package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/servicecatalogappregistry/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"associate-attribute-group", "associate-resource", "create-application", "create-attribute-group", "delete-application", "delete-attribute-group", "disassociate-attribute-group", "disassociate-resource", "get-application", "get-associated-resource", "get-attribute-group", "get-configuration", "list-applications", "list-associated-attribute-groups", "list-associated-resources", "list-attribute-groups", "list-attribute-groups-for-application", "list-tags-for-resource", "put-configuration", "sync-resource", "tag-resource", "untag-resource", "update-application", "update-attribute-group"},
		OperationSet: map[string]bool{"associate-attribute-group": true, "associate-resource": true, "create-application": true, "create-attribute-group": true, "delete-application": true, "delete-attribute-group": true, "disassociate-attribute-group": true, "disassociate-resource": true, "get-application": true, "get-associated-resource": true, "get-attribute-group": true, "get-configuration": true, "list-applications": true, "list-associated-attribute-groups": true, "list-associated-resources": true, "list-attribute-groups": true, "list-attribute-groups-for-application": true, "list-tags-for-resource": true, "put-configuration": true, "sync-resource": true, "tag-resource": true, "untag-resource": true, "update-application": true, "update-attribute-group": true},
		OperationInputs: map[string][]string{
			"associate-attribute-group":             {"Application", "AttributeGroup"},
			"associate-resource":                    {"Application", "Options", "Resource", "ResourceType"},
			"create-application":                    {"ClientToken", "Description", "Name", "Tags"},
			"create-attribute-group":                {"Attributes", "ClientToken", "Description", "Name", "Tags"},
			"delete-application":                    {"Application"},
			"delete-attribute-group":                {"AttributeGroup"},
			"disassociate-attribute-group":          {"Application", "AttributeGroup"},
			"disassociate-resource":                 {"Application", "Resource", "ResourceType"},
			"get-application":                       {"Application"},
			"get-associated-resource":               {"Application", "MaxResults", "NextToken", "Resource", "ResourceTagStatus", "ResourceType"},
			"get-attribute-group":                   {"AttributeGroup"},
			"get-configuration":                     {},
			"list-applications":                     {"MaxResults", "NextToken"},
			"list-associated-attribute-groups":      {"Application", "MaxResults", "NextToken"},
			"list-associated-resources":             {"Application", "MaxResults", "NextToken"},
			"list-attribute-groups":                 {"MaxResults", "NextToken"},
			"list-attribute-groups-for-application": {"Application", "MaxResults", "NextToken"},
			"list-tags-for-resource":                {"ResourceArn"},
			"put-configuration":                     {"Configuration"},
			"sync-resource":                         {"Resource", "ResourceType"},
			"tag-resource":                          {"ResourceArn", "Tags"},
			"untag-resource":                        {"ResourceArn", "TagKeys"},
			"update-application":                    {"Application", "Description", "Name"},
			"update-attribute-group":                {"AttributeGroup", "Attributes", "Description", "Name"},
		},
		OperationInputTypes: map[string]map[string]string{
			"associate-attribute-group":             {"Application": "*string", "AttributeGroup": "*string"},
			"associate-resource":                    {"Application": "*string", "Options": "[]types.AssociationOption", "Resource": "*string", "ResourceType": "types.ResourceType"},
			"create-application":                    {"ClientToken": "*string", "Description": "*string", "Name": "*string", "Tags": "map[string]string"},
			"create-attribute-group":                {"Attributes": "*string", "ClientToken": "*string", "Description": "*string", "Name": "*string", "Tags": "map[string]string"},
			"delete-application":                    {"Application": "*string"},
			"delete-attribute-group":                {"AttributeGroup": "*string"},
			"disassociate-attribute-group":          {"Application": "*string", "AttributeGroup": "*string"},
			"disassociate-resource":                 {"Application": "*string", "Resource": "*string", "ResourceType": "types.ResourceType"},
			"get-application":                       {"Application": "*string"},
			"get-associated-resource":               {"Application": "*string", "MaxResults": "*int32", "NextToken": "*string", "Resource": "*string", "ResourceTagStatus": "[]types.ResourceItemStatus", "ResourceType": "types.ResourceType"},
			"get-attribute-group":                   {"AttributeGroup": "*string"},
			"get-configuration":                     {},
			"list-applications":                     {"MaxResults": "*int32", "NextToken": "*string"},
			"list-associated-attribute-groups":      {"Application": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-associated-resources":             {"Application": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-attribute-groups":                 {"MaxResults": "*int32", "NextToken": "*string"},
			"list-attribute-groups-for-application": {"Application": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":                {"ResourceArn": "*string"},
			"put-configuration":                     {"Configuration": "*types.AppRegistryConfiguration"},
			"sync-resource":                         {"Resource": "*string", "ResourceType": "types.ResourceType"},
			"tag-resource":                          {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":                        {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-application":                    {"Application": "*string", "Description": "*string", "Name": "*string"},
			"update-attribute-group":                {"AttributeGroup": "*string", "Attributes": "*string", "Description": "*string", "Name": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"associate-attribute-group":             {"Application", "AttributeGroup"},
			"associate-resource":                    {"Application", "Resource", "ResourceType"},
			"create-application":                    {"ClientToken", "Name"},
			"create-attribute-group":                {"Attributes", "ClientToken", "Name"},
			"delete-application":                    {"Application"},
			"delete-attribute-group":                {"AttributeGroup"},
			"disassociate-attribute-group":          {"Application", "AttributeGroup"},
			"disassociate-resource":                 {"Application", "Resource", "ResourceType"},
			"get-application":                       {"Application"},
			"get-associated-resource":               {"Application", "Resource", "ResourceType"},
			"get-attribute-group":                   {"AttributeGroup"},
			"get-configuration":                     {},
			"list-applications":                     {},
			"list-associated-attribute-groups":      {"Application"},
			"list-associated-resources":             {"Application"},
			"list-attribute-groups":                 {},
			"list-attribute-groups-for-application": {"Application"},
			"list-tags-for-resource":                {"ResourceArn"},
			"put-configuration":                     {"Configuration"},
			"sync-resource":                         {"Resource", "ResourceType"},
			"tag-resource":                          {"ResourceArn", "Tags"},
			"untag-resource":                        {"ResourceArn", "TagKeys"},
			"update-application":                    {"Application"},
			"update-attribute-group":                {"AttributeGroup"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("servicecatalogappregistry", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
