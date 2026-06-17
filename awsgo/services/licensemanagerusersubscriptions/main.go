package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/licensemanagerusersubscriptions/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"associate-user", "create-license-server-endpoint", "delete-license-server-endpoint", "deregister-identity-provider", "disassociate-user", "list-identity-providers", "list-instances", "list-license-server-endpoints", "list-product-subscriptions", "list-tags-for-resource", "list-user-associations", "register-identity-provider", "start-product-subscription", "stop-product-subscription", "tag-resource", "untag-resource", "update-identity-provider-settings"},
		OperationSet: map[string]bool{"associate-user": true, "create-license-server-endpoint": true, "delete-license-server-endpoint": true, "deregister-identity-provider": true, "disassociate-user": true, "list-identity-providers": true, "list-instances": true, "list-license-server-endpoints": true, "list-product-subscriptions": true, "list-tags-for-resource": true, "list-user-associations": true, "register-identity-provider": true, "start-product-subscription": true, "stop-product-subscription": true, "tag-resource": true, "untag-resource": true, "update-identity-provider-settings": true},
		OperationInputs: map[string][]string{
			"associate-user":                    {"Domain", "IdentityProvider", "InstanceId", "Tags", "Username"},
			"create-license-server-endpoint":    {"IdentityProviderArn", "LicenseServerSettings", "Tags"},
			"delete-license-server-endpoint":    {"LicenseServerEndpointArn", "ServerType"},
			"deregister-identity-provider":      {"IdentityProvider", "IdentityProviderArn", "Product"},
			"disassociate-user":                 {"Domain", "IdentityProvider", "InstanceId", "InstanceUserArn", "Username"},
			"list-identity-providers":           {"Filters", "MaxResults", "NextToken"},
			"list-instances":                    {"Filters", "MaxResults", "NextToken"},
			"list-license-server-endpoints":     {"Filters", "MaxResults", "NextToken"},
			"list-product-subscriptions":        {"Filters", "IdentityProvider", "MaxResults", "NextToken", "Product"},
			"list-tags-for-resource":            {"ResourceArn"},
			"list-user-associations":            {"Filters", "IdentityProvider", "InstanceId", "MaxResults", "NextToken"},
			"register-identity-provider":        {"IdentityProvider", "Product", "Settings", "Tags"},
			"start-product-subscription":        {"Domain", "IdentityProvider", "Product", "Tags", "Username"},
			"stop-product-subscription":         {"Domain", "IdentityProvider", "Product", "ProductUserArn", "Username"},
			"tag-resource":                      {"ResourceArn", "Tags"},
			"untag-resource":                    {"ResourceArn", "TagKeys"},
			"update-identity-provider-settings": {"IdentityProvider", "IdentityProviderArn", "Product", "UpdateSettings"},
		},
		OperationInputTypes: map[string]map[string]string{
			"associate-user":                    {"Domain": "*string", "IdentityProvider": "types.IdentityProvider", "InstanceId": "*string", "Tags": "map[string]string", "Username": "*string"},
			"create-license-server-endpoint":    {"IdentityProviderArn": "*string", "LicenseServerSettings": "*types.LicenseServerSettings", "Tags": "map[string]string"},
			"delete-license-server-endpoint":    {"LicenseServerEndpointArn": "*string", "ServerType": "types.ServerType"},
			"deregister-identity-provider":      {"IdentityProvider": "types.IdentityProvider", "IdentityProviderArn": "*string", "Product": "*string"},
			"disassociate-user":                 {"Domain": "*string", "IdentityProvider": "types.IdentityProvider", "InstanceId": "*string", "InstanceUserArn": "*string", "Username": "*string"},
			"list-identity-providers":           {"Filters": "[]types.Filter", "MaxResults": "*int32", "NextToken": "*string"},
			"list-instances":                    {"Filters": "[]types.Filter", "MaxResults": "*int32", "NextToken": "*string"},
			"list-license-server-endpoints":     {"Filters": "[]types.Filter", "MaxResults": "*int32", "NextToken": "*string"},
			"list-product-subscriptions":        {"Filters": "[]types.Filter", "IdentityProvider": "types.IdentityProvider", "MaxResults": "*int32", "NextToken": "*string", "Product": "*string"},
			"list-tags-for-resource":            {"ResourceArn": "*string"},
			"list-user-associations":            {"Filters": "[]types.Filter", "IdentityProvider": "types.IdentityProvider", "InstanceId": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"register-identity-provider":        {"IdentityProvider": "types.IdentityProvider", "Product": "*string", "Settings": "*types.Settings", "Tags": "map[string]string"},
			"start-product-subscription":        {"Domain": "*string", "IdentityProvider": "types.IdentityProvider", "Product": "*string", "Tags": "map[string]string", "Username": "*string"},
			"stop-product-subscription":         {"Domain": "*string", "IdentityProvider": "types.IdentityProvider", "Product": "*string", "ProductUserArn": "*string", "Username": "*string"},
			"tag-resource":                      {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":                    {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-identity-provider-settings": {"IdentityProvider": "types.IdentityProvider", "IdentityProviderArn": "*string", "Product": "*string", "UpdateSettings": "*types.UpdateSettings"},
		},
		OperationInputRequired: map[string][]string{
			"associate-user":                    {"IdentityProvider", "InstanceId", "Username"},
			"create-license-server-endpoint":    {"IdentityProviderArn", "LicenseServerSettings"},
			"delete-license-server-endpoint":    {"LicenseServerEndpointArn", "ServerType"},
			"deregister-identity-provider":      {},
			"disassociate-user":                 {},
			"list-identity-providers":           {},
			"list-instances":                    {},
			"list-license-server-endpoints":     {},
			"list-product-subscriptions":        {"IdentityProvider"},
			"list-tags-for-resource":            {"ResourceArn"},
			"list-user-associations":            {"IdentityProvider", "InstanceId"},
			"register-identity-provider":        {"IdentityProvider", "Product"},
			"start-product-subscription":        {"IdentityProvider", "Product", "Username"},
			"stop-product-subscription":         {},
			"tag-resource":                      {"ResourceArn", "Tags"},
			"untag-resource":                    {"ResourceArn", "TagKeys"},
			"update-identity-provider-settings": {"UpdateSettings"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("licensemanagerusersubscriptions", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
