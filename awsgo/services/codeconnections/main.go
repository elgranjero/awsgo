package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/codeconnections/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-connection", "create-host", "create-repository-link", "create-sync-configuration", "delete-connection", "delete-host", "delete-repository-link", "delete-sync-configuration", "get-connection", "get-host", "get-repository-link", "get-repository-sync-status", "get-resource-sync-status", "get-sync-blocker-summary", "get-sync-configuration", "list-connections", "list-hosts", "list-repository-links", "list-repository-sync-definitions", "list-sync-configurations", "list-tags-for-resource", "tag-resource", "untag-resource", "update-host", "update-repository-link", "update-sync-blocker", "update-sync-configuration"},
		OperationSet: map[string]bool{"create-connection": true, "create-host": true, "create-repository-link": true, "create-sync-configuration": true, "delete-connection": true, "delete-host": true, "delete-repository-link": true, "delete-sync-configuration": true, "get-connection": true, "get-host": true, "get-repository-link": true, "get-repository-sync-status": true, "get-resource-sync-status": true, "get-sync-blocker-summary": true, "get-sync-configuration": true, "list-connections": true, "list-hosts": true, "list-repository-links": true, "list-repository-sync-definitions": true, "list-sync-configurations": true, "list-tags-for-resource": true, "tag-resource": true, "untag-resource": true, "update-host": true, "update-repository-link": true, "update-sync-blocker": true, "update-sync-configuration": true},
		OperationInputs: map[string][]string{
			"create-connection":                {"ConnectionName", "HostArn", "ProviderType", "Tags"},
			"create-host":                      {"Name", "ProviderEndpoint", "ProviderType", "Tags", "VpcConfiguration"},
			"create-repository-link":           {"ConnectionArn", "EncryptionKeyArn", "OwnerId", "RepositoryName", "Tags"},
			"create-sync-configuration":        {"Branch", "ConfigFile", "PublishDeploymentStatus", "PullRequestComment", "RepositoryLinkId", "ResourceName", "RoleArn", "SyncType", "TriggerResourceUpdateOn"},
			"delete-connection":                {"ConnectionArn"},
			"delete-host":                      {"HostArn"},
			"delete-repository-link":           {"RepositoryLinkId"},
			"delete-sync-configuration":        {"ResourceName", "SyncType"},
			"get-connection":                   {"ConnectionArn"},
			"get-host":                         {"HostArn"},
			"get-repository-link":              {"RepositoryLinkId"},
			"get-repository-sync-status":       {"Branch", "RepositoryLinkId", "SyncType"},
			"get-resource-sync-status":         {"ResourceName", "SyncType"},
			"get-sync-blocker-summary":         {"ResourceName", "SyncType"},
			"get-sync-configuration":           {"ResourceName", "SyncType"},
			"list-connections":                 {"HostArnFilter", "MaxResults", "NextToken", "ProviderTypeFilter"},
			"list-hosts":                       {"MaxResults", "NextToken"},
			"list-repository-links":            {"MaxResults", "NextToken"},
			"list-repository-sync-definitions": {"RepositoryLinkId", "SyncType"},
			"list-sync-configurations":         {"MaxResults", "NextToken", "RepositoryLinkId", "SyncType"},
			"list-tags-for-resource":           {"ResourceArn"},
			"tag-resource":                     {"ResourceArn", "Tags"},
			"untag-resource":                   {"ResourceArn", "TagKeys"},
			"update-host":                      {"HostArn", "ProviderEndpoint", "VpcConfiguration"},
			"update-repository-link":           {"ConnectionArn", "EncryptionKeyArn", "RepositoryLinkId"},
			"update-sync-blocker":              {"Id", "ResolvedReason", "ResourceName", "SyncType"},
			"update-sync-configuration":        {"Branch", "ConfigFile", "PublishDeploymentStatus", "PullRequestComment", "RepositoryLinkId", "ResourceName", "RoleArn", "SyncType", "TriggerResourceUpdateOn"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-connection":                {"ConnectionName": "*string", "HostArn": "*string", "ProviderType": "types.ProviderType", "Tags": "[]types.Tag"},
			"create-host":                      {"Name": "*string", "ProviderEndpoint": "*string", "ProviderType": "types.ProviderType", "Tags": "[]types.Tag", "VpcConfiguration": "*types.VpcConfiguration"},
			"create-repository-link":           {"ConnectionArn": "*string", "EncryptionKeyArn": "*string", "OwnerId": "*string", "RepositoryName": "*string", "Tags": "[]types.Tag"},
			"create-sync-configuration":        {"Branch": "*string", "ConfigFile": "*string", "PublishDeploymentStatus": "types.PublishDeploymentStatus", "PullRequestComment": "types.PullRequestComment", "RepositoryLinkId": "*string", "ResourceName": "*string", "RoleArn": "*string", "SyncType": "types.SyncConfigurationType", "TriggerResourceUpdateOn": "types.TriggerResourceUpdateOn"},
			"delete-connection":                {"ConnectionArn": "*string"},
			"delete-host":                      {"HostArn": "*string"},
			"delete-repository-link":           {"RepositoryLinkId": "*string"},
			"delete-sync-configuration":        {"ResourceName": "*string", "SyncType": "types.SyncConfigurationType"},
			"get-connection":                   {"ConnectionArn": "*string"},
			"get-host":                         {"HostArn": "*string"},
			"get-repository-link":              {"RepositoryLinkId": "*string"},
			"get-repository-sync-status":       {"Branch": "*string", "RepositoryLinkId": "*string", "SyncType": "types.SyncConfigurationType"},
			"get-resource-sync-status":         {"ResourceName": "*string", "SyncType": "types.SyncConfigurationType"},
			"get-sync-blocker-summary":         {"ResourceName": "*string", "SyncType": "types.SyncConfigurationType"},
			"get-sync-configuration":           {"ResourceName": "*string", "SyncType": "types.SyncConfigurationType"},
			"list-connections":                 {"HostArnFilter": "*string", "MaxResults": "int32", "NextToken": "*string", "ProviderTypeFilter": "types.ProviderType"},
			"list-hosts":                       {"MaxResults": "int32", "NextToken": "*string"},
			"list-repository-links":            {"MaxResults": "int32", "NextToken": "*string"},
			"list-repository-sync-definitions": {"RepositoryLinkId": "*string", "SyncType": "types.SyncConfigurationType"},
			"list-sync-configurations":         {"MaxResults": "int32", "NextToken": "*string", "RepositoryLinkId": "*string", "SyncType": "types.SyncConfigurationType"},
			"list-tags-for-resource":           {"ResourceArn": "*string"},
			"tag-resource":                     {"ResourceArn": "*string", "Tags": "[]types.Tag"},
			"untag-resource":                   {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-host":                      {"HostArn": "*string", "ProviderEndpoint": "*string", "VpcConfiguration": "*types.VpcConfiguration"},
			"update-repository-link":           {"ConnectionArn": "*string", "EncryptionKeyArn": "*string", "RepositoryLinkId": "*string"},
			"update-sync-blocker":              {"Id": "*string", "ResolvedReason": "*string", "ResourceName": "*string", "SyncType": "types.SyncConfigurationType"},
			"update-sync-configuration":        {"Branch": "*string", "ConfigFile": "*string", "PublishDeploymentStatus": "types.PublishDeploymentStatus", "PullRequestComment": "types.PullRequestComment", "RepositoryLinkId": "*string", "ResourceName": "*string", "RoleArn": "*string", "SyncType": "types.SyncConfigurationType", "TriggerResourceUpdateOn": "types.TriggerResourceUpdateOn"},
		},
		OperationInputRequired: map[string][]string{
			"create-connection":                {"ConnectionName"},
			"create-host":                      {"Name", "ProviderEndpoint", "ProviderType"},
			"create-repository-link":           {"ConnectionArn", "OwnerId", "RepositoryName"},
			"create-sync-configuration":        {"Branch", "ConfigFile", "RepositoryLinkId", "ResourceName", "RoleArn", "SyncType"},
			"delete-connection":                {"ConnectionArn"},
			"delete-host":                      {"HostArn"},
			"delete-repository-link":           {"RepositoryLinkId"},
			"delete-sync-configuration":        {"ResourceName", "SyncType"},
			"get-connection":                   {"ConnectionArn"},
			"get-host":                         {"HostArn"},
			"get-repository-link":              {"RepositoryLinkId"},
			"get-repository-sync-status":       {"Branch", "RepositoryLinkId", "SyncType"},
			"get-resource-sync-status":         {"ResourceName", "SyncType"},
			"get-sync-blocker-summary":         {"ResourceName", "SyncType"},
			"get-sync-configuration":           {"ResourceName", "SyncType"},
			"list-connections":                 {},
			"list-hosts":                       {},
			"list-repository-links":            {},
			"list-repository-sync-definitions": {"RepositoryLinkId", "SyncType"},
			"list-sync-configurations":         {"RepositoryLinkId", "SyncType"},
			"list-tags-for-resource":           {"ResourceArn"},
			"tag-resource":                     {"ResourceArn", "Tags"},
			"untag-resource":                   {"ResourceArn", "TagKeys"},
			"update-host":                      {"HostArn"},
			"update-repository-link":           {"RepositoryLinkId"},
			"update-sync-blocker":              {"Id", "ResolvedReason", "ResourceName", "SyncType"},
			"update-sync-configuration":        {"ResourceName", "SyncType"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("codeconnections", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
