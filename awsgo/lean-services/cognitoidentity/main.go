package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
)

var fields_create_identity_pool = []leanruntime.Field{
	{Name: "AllowClassicFlow", Flag: "allow-classic-flow", Type: "*bool", Required: false},
	{Name: "AllowUnauthenticatedIdentities", Flag: "allow-unauthenticated-identities", Type: "bool", Required: true},
	{Name: "CognitoIdentityProviders", Flag: "cognito-identity-providers", Type: "[]types.CognitoIdentityProvider", Required: false},
	{Name: "DeveloperProviderName", Flag: "developer-provider-name", Type: "*string", Required: false},
	{Name: "IdentityPoolName", Flag: "identity-pool-name", Type: "*string", Required: true},
	{Name: "IdentityPoolTags", Flag: "identity-pool-tags", Type: "map[string]string", Required: false},
	{Name: "OpenIdConnectProviderARNs", Flag: "open-id-connect-provider-arns", Type: "[]string", Required: false},
	{Name: "SamlProviderARNs", Flag: "saml-provider-arns", Type: "[]string", Required: false},
	{Name: "SupportedLoginProviders", Flag: "supported-login-providers", Type: "map[string]string", Required: false},
}

var fields_delete_identities = []leanruntime.Field{
	{Name: "IdentityIdsToDelete", Flag: "identity-ids-to-delete", Type: "[]string", Required: true},
}

var fields_delete_identity_pool = []leanruntime.Field{
	{Name: "IdentityPoolId", Flag: "identity-pool-id", Type: "*string", Required: true},
}

var fields_describe_identity = []leanruntime.Field{
	{Name: "IdentityId", Flag: "identity-id", Type: "*string", Required: true},
}

var fields_describe_identity_pool = []leanruntime.Field{
	{Name: "IdentityPoolId", Flag: "identity-pool-id", Type: "*string", Required: true},
}

var fields_get_credentials_for_identity = []leanruntime.Field{
	{Name: "CustomRoleArn", Flag: "custom-role-arn", Type: "*string", Required: false},
	{Name: "IdentityId", Flag: "identity-id", Type: "*string", Required: true},
	{Name: "Logins", Flag: "logins", Type: "map[string]string", Required: false},
}

var fields_get_id = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "IdentityPoolId", Flag: "identity-pool-id", Type: "*string", Required: true},
	{Name: "Logins", Flag: "logins", Type: "map[string]string", Required: false},
}

var fields_get_identity_pool_roles = []leanruntime.Field{
	{Name: "IdentityPoolId", Flag: "identity-pool-id", Type: "*string", Required: true},
}

var fields_get_open_id_token = []leanruntime.Field{
	{Name: "IdentityId", Flag: "identity-id", Type: "*string", Required: true},
	{Name: "Logins", Flag: "logins", Type: "map[string]string", Required: false},
}

var fields_get_open_id_token_for_developer_identity = []leanruntime.Field{
	{Name: "IdentityId", Flag: "identity-id", Type: "*string", Required: false},
	{Name: "IdentityPoolId", Flag: "identity-pool-id", Type: "*string", Required: true},
	{Name: "Logins", Flag: "logins", Type: "map[string]string", Required: true},
	{Name: "PrincipalTags", Flag: "principal-tags", Type: "map[string]string", Required: false},
	{Name: "TokenDuration", Flag: "token-duration", Type: "*int64", Required: false},
}

var fields_get_principal_tag_attribute_map = []leanruntime.Field{
	{Name: "IdentityPoolId", Flag: "identity-pool-id", Type: "*string", Required: true},
	{Name: "IdentityProviderName", Flag: "identity-provider-name", Type: "*string", Required: true},
}

var fields_list_identities = []leanruntime.Field{
	{Name: "HideDisabled", Flag: "hide-disabled", Type: "bool", Required: false},
	{Name: "IdentityPoolId", Flag: "identity-pool-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_identity_pools = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_lookup_developer_identity = []leanruntime.Field{
	{Name: "DeveloperUserIdentifier", Flag: "developer-user-identifier", Type: "*string", Required: false},
	{Name: "IdentityId", Flag: "identity-id", Type: "*string", Required: false},
	{Name: "IdentityPoolId", Flag: "identity-pool-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_merge_developer_identities = []leanruntime.Field{
	{Name: "DestinationUserIdentifier", Flag: "destination-user-identifier", Type: "*string", Required: true},
	{Name: "DeveloperProviderName", Flag: "developer-provider-name", Type: "*string", Required: true},
	{Name: "IdentityPoolId", Flag: "identity-pool-id", Type: "*string", Required: true},
	{Name: "SourceUserIdentifier", Flag: "source-user-identifier", Type: "*string", Required: true},
}

var fields_set_identity_pool_roles = []leanruntime.Field{
	{Name: "IdentityPoolId", Flag: "identity-pool-id", Type: "*string", Required: true},
	{Name: "RoleMappings", Flag: "role-mappings", Type: "map[string]types.RoleMapping", Required: false},
	{Name: "Roles", Flag: "roles", Type: "map[string]string", Required: true},
}

var fields_set_principal_tag_attribute_map = []leanruntime.Field{
	{Name: "IdentityPoolId", Flag: "identity-pool-id", Type: "*string", Required: true},
	{Name: "IdentityProviderName", Flag: "identity-provider-name", Type: "*string", Required: true},
	{Name: "PrincipalTags", Flag: "principal-tags", Type: "map[string]string", Required: false},
	{Name: "UseDefaults", Flag: "use-defaults", Type: "*bool", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_unlink_developer_identity = []leanruntime.Field{
	{Name: "DeveloperProviderName", Flag: "developer-provider-name", Type: "*string", Required: true},
	{Name: "DeveloperUserIdentifier", Flag: "developer-user-identifier", Type: "*string", Required: true},
	{Name: "IdentityId", Flag: "identity-id", Type: "*string", Required: true},
	{Name: "IdentityPoolId", Flag: "identity-pool-id", Type: "*string", Required: true},
}

var fields_unlink_identity = []leanruntime.Field{
	{Name: "IdentityId", Flag: "identity-id", Type: "*string", Required: true},
	{Name: "Logins", Flag: "logins", Type: "map[string]string", Required: true},
	{Name: "LoginsToRemove", Flag: "logins-to-remove", Type: "[]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_identity_pool = []leanruntime.Field{
	{Name: "AllowClassicFlow", Flag: "allow-classic-flow", Type: "*bool", Required: false},
	{Name: "AllowUnauthenticatedIdentities", Flag: "allow-unauthenticated-identities", Type: "bool", Required: true},
	{Name: "CognitoIdentityProviders", Flag: "cognito-identity-providers", Type: "[]types.CognitoIdentityProvider", Required: false},
	{Name: "DeveloperProviderName", Flag: "developer-provider-name", Type: "*string", Required: false},
	{Name: "IdentityPoolId", Flag: "identity-pool-id", Type: "*string", Required: true},
	{Name: "IdentityPoolName", Flag: "identity-pool-name", Type: "*string", Required: true},
	{Name: "IdentityPoolTags", Flag: "identity-pool-tags", Type: "map[string]string", Required: false},
	{Name: "OpenIdConnectProviderARNs", Flag: "open-id-connect-provider-arns", Type: "[]string", Required: false},
	{Name: "SamlProviderARNs", Flag: "saml-provider-arns", Type: "[]string", Required: false},
	{Name: "SupportedLoginProviders", Flag: "supported-login-providers", Type: "map[string]string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-identity-pool": {
			Name:   "create-identity-pool",
			Fields: fields_create_identity_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIdentityPoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_identity_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIdentityPool(ctx, input)
			},
		},
		"delete-identities": {
			Name:   "delete-identities",
			Fields: fields_delete_identities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIdentitiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_identities, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIdentities(ctx, input)
			},
		},
		"delete-identity-pool": {
			Name:   "delete-identity-pool",
			Fields: fields_delete_identity_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIdentityPoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_identity_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIdentityPool(ctx, input)
			},
		},
		"describe-identity": {
			Name:   "describe-identity",
			Fields: fields_describe_identity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeIdentityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_identity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeIdentity(ctx, input)
			},
		},
		"describe-identity-pool": {
			Name:   "describe-identity-pool",
			Fields: fields_describe_identity_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeIdentityPoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_identity_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeIdentityPool(ctx, input)
			},
		},
		"get-credentials-for-identity": {
			Name:   "get-credentials-for-identity",
			Fields: fields_get_credentials_for_identity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCredentialsForIdentityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_credentials_for_identity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCredentialsForIdentity(ctx, input)
			},
		},
		"get-id": {
			Name:   "get-id",
			Fields: fields_get_id,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIdInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_id, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetId(ctx, input)
			},
		},
		"get-identity-pool-roles": {
			Name:   "get-identity-pool-roles",
			Fields: fields_get_identity_pool_roles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIdentityPoolRolesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_identity_pool_roles, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIdentityPoolRoles(ctx, input)
			},
		},
		"get-open-id-token": {
			Name:   "get-open-id-token",
			Fields: fields_get_open_id_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOpenIdTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_open_id_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOpenIdToken(ctx, input)
			},
		},
		"get-open-id-token-for-developer-identity": {
			Name:   "get-open-id-token-for-developer-identity",
			Fields: fields_get_open_id_token_for_developer_identity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOpenIdTokenForDeveloperIdentityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_open_id_token_for_developer_identity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOpenIdTokenForDeveloperIdentity(ctx, input)
			},
		},
		"get-principal-tag-attribute-map": {
			Name:   "get-principal-tag-attribute-map",
			Fields: fields_get_principal_tag_attribute_map,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPrincipalTagAttributeMapInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_principal_tag_attribute_map, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPrincipalTagAttributeMap(ctx, input)
			},
		},
		"list-identities": {
			Name:   "list-identities",
			Fields: fields_list_identities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIdentitiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_identities, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListIdentities(ctx, input)
			},
		},
		"list-identity-pools": {
			Name:   "list-identity-pools",
			Fields: fields_list_identity_pools,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIdentityPoolsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_identity_pools, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIdentityPools(ctx, input)
				}
				var results []*svc.ListIdentityPoolsOutput
				p := svc.NewListIdentityPoolsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-tags-for-resource": {
			Name:   "list-tags-for-resource",
			Fields: fields_list_tags_for_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsForResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_tags_for_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTagsForResource(ctx, input)
			},
		},
		"lookup-developer-identity": {
			Name:   "lookup-developer-identity",
			Fields: fields_lookup_developer_identity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.LookupDeveloperIdentityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_lookup_developer_identity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.LookupDeveloperIdentity(ctx, input)
			},
		},
		"merge-developer-identities": {
			Name:   "merge-developer-identities",
			Fields: fields_merge_developer_identities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.MergeDeveloperIdentitiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_merge_developer_identities, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.MergeDeveloperIdentities(ctx, input)
			},
		},
		"set-identity-pool-roles": {
			Name:   "set-identity-pool-roles",
			Fields: fields_set_identity_pool_roles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetIdentityPoolRolesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_identity_pool_roles, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetIdentityPoolRoles(ctx, input)
			},
		},
		"set-principal-tag-attribute-map": {
			Name:   "set-principal-tag-attribute-map",
			Fields: fields_set_principal_tag_attribute_map,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetPrincipalTagAttributeMapInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_principal_tag_attribute_map, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetPrincipalTagAttributeMap(ctx, input)
			},
		},
		"tag-resource": {
			Name:   "tag-resource",
			Fields: fields_tag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_tag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TagResource(ctx, input)
			},
		},
		"unlink-developer-identity": {
			Name:   "unlink-developer-identity",
			Fields: fields_unlink_developer_identity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UnlinkDeveloperIdentityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_unlink_developer_identity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UnlinkDeveloperIdentity(ctx, input)
			},
		},
		"unlink-identity": {
			Name:   "unlink-identity",
			Fields: fields_unlink_identity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UnlinkIdentityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_unlink_identity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UnlinkIdentity(ctx, input)
			},
		},
		"untag-resource": {
			Name:   "untag-resource",
			Fields: fields_untag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UntagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_untag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UntagResource(ctx, input)
			},
		},
		"update-identity-pool": {
			Name:   "update-identity-pool",
			Fields: fields_update_identity_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateIdentityPoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_identity_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateIdentityPool(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("cognitoidentity", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
