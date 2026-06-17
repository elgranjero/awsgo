package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/grafana/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"associate-license", "create-workspace", "create-workspace-api-key", "create-workspace-service-account", "create-workspace-service-account-token", "delete-workspace", "delete-workspace-api-key", "delete-workspace-service-account", "delete-workspace-service-account-token", "describe-workspace", "describe-workspace-authentication", "describe-workspace-configuration", "disassociate-license", "list-permissions", "list-tags-for-resource", "list-versions", "list-workspace-service-account-tokens", "list-workspace-service-accounts", "list-workspaces", "tag-resource", "untag-resource", "update-permissions", "update-workspace", "update-workspace-authentication", "update-workspace-configuration"},
		OperationSet: map[string]bool{"associate-license": true, "create-workspace": true, "create-workspace-api-key": true, "create-workspace-service-account": true, "create-workspace-service-account-token": true, "delete-workspace": true, "delete-workspace-api-key": true, "delete-workspace-service-account": true, "delete-workspace-service-account-token": true, "describe-workspace": true, "describe-workspace-authentication": true, "describe-workspace-configuration": true, "disassociate-license": true, "list-permissions": true, "list-tags-for-resource": true, "list-versions": true, "list-workspace-service-account-tokens": true, "list-workspace-service-accounts": true, "list-workspaces": true, "tag-resource": true, "untag-resource": true, "update-permissions": true, "update-workspace": true, "update-workspace-authentication": true, "update-workspace-configuration": true},
		OperationInputs: map[string][]string{
			"associate-license":                      {"GrafanaToken", "LicenseType", "WorkspaceId"},
			"create-workspace":                       {"AccountAccessType", "AuthenticationProviders", "ClientToken", "Configuration", "GrafanaVersion", "KmsKeyId", "NetworkAccessControl", "OrganizationRoleName", "PermissionType", "StackSetName", "Tags", "VpcConfiguration", "WorkspaceDataSources", "WorkspaceDescription", "WorkspaceName", "WorkspaceNotificationDestinations", "WorkspaceOrganizationalUnits", "WorkspaceRoleArn"},
			"create-workspace-api-key":               {"KeyName", "KeyRole", "SecondsToLive", "WorkspaceId"},
			"create-workspace-service-account":       {"GrafanaRole", "Name", "WorkspaceId"},
			"create-workspace-service-account-token": {"Name", "SecondsToLive", "ServiceAccountId", "WorkspaceId"},
			"delete-workspace":                       {"WorkspaceId"},
			"delete-workspace-api-key":               {"KeyName", "WorkspaceId"},
			"delete-workspace-service-account":       {"ServiceAccountId", "WorkspaceId"},
			"delete-workspace-service-account-token": {"ServiceAccountId", "TokenId", "WorkspaceId"},
			"describe-workspace":                     {"WorkspaceId"},
			"describe-workspace-authentication":      {"WorkspaceId"},
			"describe-workspace-configuration":       {"WorkspaceId"},
			"disassociate-license":                   {"LicenseType", "WorkspaceId"},
			"list-permissions":                       {"GroupId", "MaxResults", "NextToken", "UserId", "UserType", "WorkspaceId"},
			"list-tags-for-resource":                 {"ResourceArn"},
			"list-versions":                          {"MaxResults", "NextToken", "WorkspaceId"},
			"list-workspace-service-account-tokens":  {"MaxResults", "NextToken", "ServiceAccountId", "WorkspaceId"},
			"list-workspace-service-accounts":        {"MaxResults", "NextToken", "WorkspaceId"},
			"list-workspaces":                        {"MaxResults", "NextToken"},
			"tag-resource":                           {"ResourceArn", "Tags"},
			"untag-resource":                         {"ResourceArn", "TagKeys"},
			"update-permissions":                     {"UpdateInstructionBatch", "WorkspaceId"},
			"update-workspace":                       {"AccountAccessType", "NetworkAccessControl", "OrganizationRoleName", "PermissionType", "RemoveNetworkAccessConfiguration", "RemoveVpcConfiguration", "StackSetName", "VpcConfiguration", "WorkspaceDataSources", "WorkspaceDescription", "WorkspaceId", "WorkspaceName", "WorkspaceNotificationDestinations", "WorkspaceOrganizationalUnits", "WorkspaceRoleArn"},
			"update-workspace-authentication":        {"AuthenticationProviders", "SamlConfiguration", "WorkspaceId"},
			"update-workspace-configuration":         {"Configuration", "GrafanaVersion", "WorkspaceId"},
		},
		OperationInputTypes: map[string]map[string]string{
			"associate-license":                      {"GrafanaToken": "*string", "LicenseType": "types.LicenseType", "WorkspaceId": "*string"},
			"create-workspace":                       {"AccountAccessType": "types.AccountAccessType", "AuthenticationProviders": "[]types.AuthenticationProviderTypes", "ClientToken": "*string", "Configuration": "*string", "GrafanaVersion": "*string", "KmsKeyId": "*string", "NetworkAccessControl": "*types.NetworkAccessConfiguration", "OrganizationRoleName": "*string", "PermissionType": "types.PermissionType", "StackSetName": "*string", "Tags": "map[string]string", "VpcConfiguration": "*types.VpcConfiguration", "WorkspaceDataSources": "[]types.DataSourceType", "WorkspaceDescription": "*string", "WorkspaceName": "*string", "WorkspaceNotificationDestinations": "[]types.NotificationDestinationType", "WorkspaceOrganizationalUnits": "[]string", "WorkspaceRoleArn": "*string"},
			"create-workspace-api-key":               {"KeyName": "*string", "KeyRole": "*string", "SecondsToLive": "*int32", "WorkspaceId": "*string"},
			"create-workspace-service-account":       {"GrafanaRole": "types.Role", "Name": "*string", "WorkspaceId": "*string"},
			"create-workspace-service-account-token": {"Name": "*string", "SecondsToLive": "*int32", "ServiceAccountId": "*string", "WorkspaceId": "*string"},
			"delete-workspace":                       {"WorkspaceId": "*string"},
			"delete-workspace-api-key":               {"KeyName": "*string", "WorkspaceId": "*string"},
			"delete-workspace-service-account":       {"ServiceAccountId": "*string", "WorkspaceId": "*string"},
			"delete-workspace-service-account-token": {"ServiceAccountId": "*string", "TokenId": "*string", "WorkspaceId": "*string"},
			"describe-workspace":                     {"WorkspaceId": "*string"},
			"describe-workspace-authentication":      {"WorkspaceId": "*string"},
			"describe-workspace-configuration":       {"WorkspaceId": "*string"},
			"disassociate-license":                   {"LicenseType": "types.LicenseType", "WorkspaceId": "*string"},
			"list-permissions":                       {"GroupId": "*string", "MaxResults": "*int32", "NextToken": "*string", "UserId": "*string", "UserType": "types.UserType", "WorkspaceId": "*string"},
			"list-tags-for-resource":                 {"ResourceArn": "*string"},
			"list-versions":                          {"MaxResults": "*int32", "NextToken": "*string", "WorkspaceId": "*string"},
			"list-workspace-service-account-tokens":  {"MaxResults": "*int32", "NextToken": "*string", "ServiceAccountId": "*string", "WorkspaceId": "*string"},
			"list-workspace-service-accounts":        {"MaxResults": "*int32", "NextToken": "*string", "WorkspaceId": "*string"},
			"list-workspaces":                        {"MaxResults": "*int32", "NextToken": "*string"},
			"tag-resource":                           {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":                         {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-permissions":                     {"UpdateInstructionBatch": "[]types.UpdateInstruction", "WorkspaceId": "*string"},
			"update-workspace":                       {"AccountAccessType": "types.AccountAccessType", "NetworkAccessControl": "*types.NetworkAccessConfiguration", "OrganizationRoleName": "*string", "PermissionType": "types.PermissionType", "RemoveNetworkAccessConfiguration": "*bool", "RemoveVpcConfiguration": "*bool", "StackSetName": "*string", "VpcConfiguration": "*types.VpcConfiguration", "WorkspaceDataSources": "[]types.DataSourceType", "WorkspaceDescription": "*string", "WorkspaceId": "*string", "WorkspaceName": "*string", "WorkspaceNotificationDestinations": "[]types.NotificationDestinationType", "WorkspaceOrganizationalUnits": "[]string", "WorkspaceRoleArn": "*string"},
			"update-workspace-authentication":        {"AuthenticationProviders": "[]types.AuthenticationProviderTypes", "SamlConfiguration": "*types.SamlConfiguration", "WorkspaceId": "*string"},
			"update-workspace-configuration":         {"Configuration": "*string", "GrafanaVersion": "*string", "WorkspaceId": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"associate-license":                      {"LicenseType", "WorkspaceId"},
			"create-workspace":                       {"AccountAccessType", "AuthenticationProviders", "PermissionType"},
			"create-workspace-api-key":               {"KeyName", "KeyRole", "SecondsToLive", "WorkspaceId"},
			"create-workspace-service-account":       {"GrafanaRole", "Name", "WorkspaceId"},
			"create-workspace-service-account-token": {"Name", "SecondsToLive", "ServiceAccountId", "WorkspaceId"},
			"delete-workspace":                       {"WorkspaceId"},
			"delete-workspace-api-key":               {"KeyName", "WorkspaceId"},
			"delete-workspace-service-account":       {"ServiceAccountId", "WorkspaceId"},
			"delete-workspace-service-account-token": {"ServiceAccountId", "TokenId", "WorkspaceId"},
			"describe-workspace":                     {"WorkspaceId"},
			"describe-workspace-authentication":      {"WorkspaceId"},
			"describe-workspace-configuration":       {"WorkspaceId"},
			"disassociate-license":                   {"LicenseType", "WorkspaceId"},
			"list-permissions":                       {"WorkspaceId"},
			"list-tags-for-resource":                 {"ResourceArn"},
			"list-versions":                          {},
			"list-workspace-service-account-tokens":  {"ServiceAccountId", "WorkspaceId"},
			"list-workspace-service-accounts":        {"WorkspaceId"},
			"list-workspaces":                        {},
			"tag-resource":                           {"ResourceArn", "Tags"},
			"untag-resource":                         {"ResourceArn", "TagKeys"},
			"update-permissions":                     {"UpdateInstructionBatch", "WorkspaceId"},
			"update-workspace":                       {"WorkspaceId"},
			"update-workspace-authentication":        {"AuthenticationProviders", "WorkspaceId"},
			"update-workspace-configuration":         {"Configuration", "WorkspaceId"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("grafana", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
