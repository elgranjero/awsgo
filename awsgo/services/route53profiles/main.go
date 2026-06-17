package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/route53profiles/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"associate-profile", "associate-resource-to-profile", "create-profile", "delete-profile", "disassociate-profile", "disassociate-resource-from-profile", "get-profile", "get-profile-association", "get-profile-resource-association", "list-profile-associations", "list-profile-resource-associations", "list-profiles", "list-tags-for-resource", "tag-resource", "untag-resource", "update-profile-resource-association"},
		OperationSet: map[string]bool{"associate-profile": true, "associate-resource-to-profile": true, "create-profile": true, "delete-profile": true, "disassociate-profile": true, "disassociate-resource-from-profile": true, "get-profile": true, "get-profile-association": true, "get-profile-resource-association": true, "list-profile-associations": true, "list-profile-resource-associations": true, "list-profiles": true, "list-tags-for-resource": true, "tag-resource": true, "untag-resource": true, "update-profile-resource-association": true},
		OperationInputs: map[string][]string{
			"associate-profile":                   {"Name", "ProfileId", "ResourceId", "Tags"},
			"associate-resource-to-profile":       {"Name", "ProfileId", "ResourceArn", "ResourceProperties"},
			"create-profile":                      {"ClientToken", "Name", "Tags"},
			"delete-profile":                      {"ProfileId"},
			"disassociate-profile":                {"ProfileId", "ResourceId"},
			"disassociate-resource-from-profile":  {"ProfileId", "ResourceArn"},
			"get-profile":                         {"ProfileId"},
			"get-profile-association":             {"ProfileAssociationId"},
			"get-profile-resource-association":    {"ProfileResourceAssociationId"},
			"list-profile-associations":           {"MaxResults", "NextToken", "ProfileId", "ResourceId"},
			"list-profile-resource-associations":  {"MaxResults", "NextToken", "ProfileId", "ResourceType"},
			"list-profiles":                       {"MaxResults", "NextToken"},
			"list-tags-for-resource":              {"ResourceArn"},
			"tag-resource":                        {"ResourceArn", "Tags"},
			"untag-resource":                      {"ResourceArn", "TagKeys"},
			"update-profile-resource-association": {"Name", "ProfileResourceAssociationId", "ResourceProperties"},
		},
		OperationInputTypes: map[string]map[string]string{
			"associate-profile":                   {"Name": "*string", "ProfileId": "*string", "ResourceId": "*string", "Tags": "[]types.Tag"},
			"associate-resource-to-profile":       {"Name": "*string", "ProfileId": "*string", "ResourceArn": "*string", "ResourceProperties": "*string"},
			"create-profile":                      {"ClientToken": "*string", "Name": "*string", "Tags": "[]types.Tag"},
			"delete-profile":                      {"ProfileId": "*string"},
			"disassociate-profile":                {"ProfileId": "*string", "ResourceId": "*string"},
			"disassociate-resource-from-profile":  {"ProfileId": "*string", "ResourceArn": "*string"},
			"get-profile":                         {"ProfileId": "*string"},
			"get-profile-association":             {"ProfileAssociationId": "*string"},
			"get-profile-resource-association":    {"ProfileResourceAssociationId": "*string"},
			"list-profile-associations":           {"MaxResults": "*int32", "NextToken": "*string", "ProfileId": "*string", "ResourceId": "*string"},
			"list-profile-resource-associations":  {"MaxResults": "*int32", "NextToken": "*string", "ProfileId": "*string", "ResourceType": "*string"},
			"list-profiles":                       {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":              {"ResourceArn": "*string"},
			"tag-resource":                        {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":                      {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-profile-resource-association": {"Name": "*string", "ProfileResourceAssociationId": "*string", "ResourceProperties": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"associate-profile":                   {"Name", "ProfileId", "ResourceId"},
			"associate-resource-to-profile":       {"Name", "ProfileId", "ResourceArn"},
			"create-profile":                      {"ClientToken", "Name"},
			"delete-profile":                      {"ProfileId"},
			"disassociate-profile":                {"ProfileId", "ResourceId"},
			"disassociate-resource-from-profile":  {"ProfileId", "ResourceArn"},
			"get-profile":                         {"ProfileId"},
			"get-profile-association":             {"ProfileAssociationId"},
			"get-profile-resource-association":    {"ProfileResourceAssociationId"},
			"list-profile-associations":           {},
			"list-profile-resource-associations":  {"ProfileId"},
			"list-profiles":                       {},
			"list-tags-for-resource":              {"ResourceArn"},
			"tag-resource":                        {"ResourceArn", "Tags"},
			"untag-resource":                      {"ResourceArn", "TagKeys"},
			"update-profile-resource-association": {"ProfileResourceAssociationId"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("route53profiles", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
