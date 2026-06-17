package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/cognitoidentity/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-identity-pool", "delete-identities", "delete-identity-pool", "describe-identity", "describe-identity-pool", "get-credentials-for-identity", "get-id", "get-identity-pool-roles", "get-open-id-token", "get-open-id-token-for-developer-identity", "get-principal-tag-attribute-map", "list-identities", "list-identity-pools", "list-tags-for-resource", "lookup-developer-identity", "merge-developer-identities", "set-identity-pool-roles", "set-principal-tag-attribute-map", "tag-resource", "unlink-developer-identity", "unlink-identity", "untag-resource", "update-identity-pool"},
		OperationSet: map[string]bool{"create-identity-pool": true, "delete-identities": true, "delete-identity-pool": true, "describe-identity": true, "describe-identity-pool": true, "get-credentials-for-identity": true, "get-id": true, "get-identity-pool-roles": true, "get-open-id-token": true, "get-open-id-token-for-developer-identity": true, "get-principal-tag-attribute-map": true, "list-identities": true, "list-identity-pools": true, "list-tags-for-resource": true, "lookup-developer-identity": true, "merge-developer-identities": true, "set-identity-pool-roles": true, "set-principal-tag-attribute-map": true, "tag-resource": true, "unlink-developer-identity": true, "unlink-identity": true, "untag-resource": true, "update-identity-pool": true},
		OperationInputs: map[string][]string{
			"create-identity-pool":                     {"AllowClassicFlow", "AllowUnauthenticatedIdentities", "CognitoIdentityProviders", "DeveloperProviderName", "IdentityPoolName", "IdentityPoolTags", "OpenIdConnectProviderARNs", "SamlProviderARNs", "SupportedLoginProviders"},
			"delete-identities":                        {"IdentityIdsToDelete"},
			"delete-identity-pool":                     {"IdentityPoolId"},
			"describe-identity":                        {"IdentityId"},
			"describe-identity-pool":                   {"IdentityPoolId"},
			"get-credentials-for-identity":             {"CustomRoleArn", "IdentityId", "Logins"},
			"get-id":                                   {"AccountId", "IdentityPoolId", "Logins"},
			"get-identity-pool-roles":                  {"IdentityPoolId"},
			"get-open-id-token":                        {"IdentityId", "Logins"},
			"get-open-id-token-for-developer-identity": {"IdentityId", "IdentityPoolId", "Logins", "PrincipalTags", "TokenDuration"},
			"get-principal-tag-attribute-map":          {"IdentityPoolId", "IdentityProviderName"},
			"list-identities":                          {"HideDisabled", "IdentityPoolId", "MaxResults", "NextToken"},
			"list-identity-pools":                      {"MaxResults", "NextToken"},
			"list-tags-for-resource":                   {"ResourceArn"},
			"lookup-developer-identity":                {"DeveloperUserIdentifier", "IdentityId", "IdentityPoolId", "MaxResults", "NextToken"},
			"merge-developer-identities":               {"DestinationUserIdentifier", "DeveloperProviderName", "IdentityPoolId", "SourceUserIdentifier"},
			"set-identity-pool-roles":                  {"IdentityPoolId", "RoleMappings", "Roles"},
			"set-principal-tag-attribute-map":          {"IdentityPoolId", "IdentityProviderName", "PrincipalTags", "UseDefaults"},
			"tag-resource":                             {"ResourceArn", "Tags"},
			"unlink-developer-identity":                {"DeveloperProviderName", "DeveloperUserIdentifier", "IdentityId", "IdentityPoolId"},
			"unlink-identity":                          {"IdentityId", "Logins", "LoginsToRemove"},
			"untag-resource":                           {"ResourceArn", "TagKeys"},
			"update-identity-pool":                     {"AllowClassicFlow", "AllowUnauthenticatedIdentities", "CognitoIdentityProviders", "DeveloperProviderName", "IdentityPoolId", "IdentityPoolName", "IdentityPoolTags", "OpenIdConnectProviderARNs", "SamlProviderARNs", "SupportedLoginProviders"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-identity-pool":                     {"AllowClassicFlow": "*bool", "AllowUnauthenticatedIdentities": "bool", "CognitoIdentityProviders": "[]types.CognitoIdentityProvider", "DeveloperProviderName": "*string", "IdentityPoolName": "*string", "IdentityPoolTags": "map[string]string", "OpenIdConnectProviderARNs": "[]string", "SamlProviderARNs": "[]string", "SupportedLoginProviders": "map[string]string"},
			"delete-identities":                        {"IdentityIdsToDelete": "[]string"},
			"delete-identity-pool":                     {"IdentityPoolId": "*string"},
			"describe-identity":                        {"IdentityId": "*string"},
			"describe-identity-pool":                   {"IdentityPoolId": "*string"},
			"get-credentials-for-identity":             {"CustomRoleArn": "*string", "IdentityId": "*string", "Logins": "map[string]string"},
			"get-id":                                   {"AccountId": "*string", "IdentityPoolId": "*string", "Logins": "map[string]string"},
			"get-identity-pool-roles":                  {"IdentityPoolId": "*string"},
			"get-open-id-token":                        {"IdentityId": "*string", "Logins": "map[string]string"},
			"get-open-id-token-for-developer-identity": {"IdentityId": "*string", "IdentityPoolId": "*string", "Logins": "map[string]string", "PrincipalTags": "map[string]string", "TokenDuration": "*int64"},
			"get-principal-tag-attribute-map":          {"IdentityPoolId": "*string", "IdentityProviderName": "*string"},
			"list-identities":                          {"HideDisabled": "bool", "IdentityPoolId": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-identity-pools":                      {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":                   {"ResourceArn": "*string"},
			"lookup-developer-identity":                {"DeveloperUserIdentifier": "*string", "IdentityId": "*string", "IdentityPoolId": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"merge-developer-identities":               {"DestinationUserIdentifier": "*string", "DeveloperProviderName": "*string", "IdentityPoolId": "*string", "SourceUserIdentifier": "*string"},
			"set-identity-pool-roles":                  {"IdentityPoolId": "*string", "RoleMappings": "map[string]types.RoleMapping", "Roles": "map[string]string"},
			"set-principal-tag-attribute-map":          {"IdentityPoolId": "*string", "IdentityProviderName": "*string", "PrincipalTags": "map[string]string", "UseDefaults": "*bool"},
			"tag-resource":                             {"ResourceArn": "*string", "Tags": "map[string]string"},
			"unlink-developer-identity":                {"DeveloperProviderName": "*string", "DeveloperUserIdentifier": "*string", "IdentityId": "*string", "IdentityPoolId": "*string"},
			"unlink-identity":                          {"IdentityId": "*string", "Logins": "map[string]string", "LoginsToRemove": "[]string"},
			"untag-resource":                           {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-identity-pool":                     {"AllowClassicFlow": "*bool", "AllowUnauthenticatedIdentities": "bool", "CognitoIdentityProviders": "[]types.CognitoIdentityProvider", "DeveloperProviderName": "*string", "IdentityPoolId": "*string", "IdentityPoolName": "*string", "IdentityPoolTags": "map[string]string", "OpenIdConnectProviderARNs": "[]string", "SamlProviderARNs": "[]string", "SupportedLoginProviders": "map[string]string"},
		},
		OperationInputRequired: map[string][]string{
			"create-identity-pool":                     {"AllowUnauthenticatedIdentities", "IdentityPoolName"},
			"delete-identities":                        {"IdentityIdsToDelete"},
			"delete-identity-pool":                     {"IdentityPoolId"},
			"describe-identity":                        {"IdentityId"},
			"describe-identity-pool":                   {"IdentityPoolId"},
			"get-credentials-for-identity":             {"IdentityId"},
			"get-id":                                   {"IdentityPoolId"},
			"get-identity-pool-roles":                  {"IdentityPoolId"},
			"get-open-id-token":                        {"IdentityId"},
			"get-open-id-token-for-developer-identity": {"IdentityPoolId", "Logins"},
			"get-principal-tag-attribute-map":          {"IdentityPoolId", "IdentityProviderName"},
			"list-identities":                          {"IdentityPoolId", "MaxResults"},
			"list-identity-pools":                      {"MaxResults"},
			"list-tags-for-resource":                   {"ResourceArn"},
			"lookup-developer-identity":                {"IdentityPoolId"},
			"merge-developer-identities":               {"DestinationUserIdentifier", "DeveloperProviderName", "IdentityPoolId", "SourceUserIdentifier"},
			"set-identity-pool-roles":                  {"IdentityPoolId", "Roles"},
			"set-principal-tag-attribute-map":          {"IdentityPoolId", "IdentityProviderName"},
			"tag-resource":                             {"ResourceArn", "Tags"},
			"unlink-developer-identity":                {"DeveloperProviderName", "DeveloperUserIdentifier", "IdentityId", "IdentityPoolId"},
			"unlink-identity":                          {"IdentityId", "Logins", "LoginsToRemove"},
			"untag-resource":                           {"ResourceArn", "TagKeys"},
			"update-identity-pool":                     {"AllowUnauthenticatedIdentities", "IdentityPoolId", "IdentityPoolName"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("cognitoidentity", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
