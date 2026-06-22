package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/verifiedpermissions"
)

var fields_batch_get_policy = []leanruntime.Field{
	{Name: "Requests", Flag: "requests", Type: "[]types.BatchGetPolicyInputItem", Required: true},
}

var fields_batch_is_authorized = []leanruntime.Field{
	{Name: "Entities", Flag: "entities", Type: "types.EntitiesDefinition", Required: false},
	{Name: "PolicyStoreId", Flag: "policy-store-id", Type: "*string", Required: true},
	{Name: "Requests", Flag: "requests", Type: "[]types.BatchIsAuthorizedInputItem", Required: true},
}

var fields_batch_is_authorized_with_token = []leanruntime.Field{
	{Name: "AccessToken", Flag: "access-token", Type: "*string", Required: false},
	{Name: "Entities", Flag: "entities", Type: "types.EntitiesDefinition", Required: false},
	{Name: "IdentityToken", Flag: "identity-token", Type: "*string", Required: false},
	{Name: "PolicyStoreId", Flag: "policy-store-id", Type: "*string", Required: true},
	{Name: "Requests", Flag: "requests", Type: "[]types.BatchIsAuthorizedWithTokenInputItem", Required: true},
}

var fields_create_identity_source = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Configuration", Flag: "configuration", Type: "types.Configuration", Required: true},
	{Name: "PolicyStoreId", Flag: "policy-store-id", Type: "*string", Required: true},
	{Name: "PrincipalEntityType", Flag: "principal-entity-type", Type: "*string", Required: false},
}

var fields_create_policy = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Definition", Flag: "definition", Type: "types.PolicyDefinition", Required: true},
	{Name: "PolicyStoreId", Flag: "policy-store-id", Type: "*string", Required: true},
}

var fields_create_policy_store = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DeletionProtection", Flag: "deletion-protection", Type: "types.DeletionProtection", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EncryptionSettings", Flag: "encryption-settings", Type: "types.EncryptionSettings", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "ValidationSettings", Flag: "validation-settings", Type: "*types.ValidationSettings", Required: true},
}

var fields_create_policy_template = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "PolicyStoreId", Flag: "policy-store-id", Type: "*string", Required: true},
	{Name: "Statement", Flag: "statement", Type: "*string", Required: true},
}

var fields_delete_identity_source = []leanruntime.Field{
	{Name: "IdentitySourceId", Flag: "identity-source-id", Type: "*string", Required: true},
	{Name: "PolicyStoreId", Flag: "policy-store-id", Type: "*string", Required: true},
}

var fields_delete_policy = []leanruntime.Field{
	{Name: "PolicyId", Flag: "policy-id", Type: "*string", Required: true},
	{Name: "PolicyStoreId", Flag: "policy-store-id", Type: "*string", Required: true},
}

var fields_delete_policy_store = []leanruntime.Field{
	{Name: "PolicyStoreId", Flag: "policy-store-id", Type: "*string", Required: true},
}

var fields_delete_policy_template = []leanruntime.Field{
	{Name: "PolicyStoreId", Flag: "policy-store-id", Type: "*string", Required: true},
	{Name: "PolicyTemplateId", Flag: "policy-template-id", Type: "*string", Required: true},
}

var fields_get_identity_source = []leanruntime.Field{
	{Name: "IdentitySourceId", Flag: "identity-source-id", Type: "*string", Required: true},
	{Name: "PolicyStoreId", Flag: "policy-store-id", Type: "*string", Required: true},
}

var fields_get_policy = []leanruntime.Field{
	{Name: "PolicyId", Flag: "policy-id", Type: "*string", Required: true},
	{Name: "PolicyStoreId", Flag: "policy-store-id", Type: "*string", Required: true},
}

var fields_get_policy_store = []leanruntime.Field{
	{Name: "PolicyStoreId", Flag: "policy-store-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "bool", Required: false},
}

var fields_get_policy_template = []leanruntime.Field{
	{Name: "PolicyStoreId", Flag: "policy-store-id", Type: "*string", Required: true},
	{Name: "PolicyTemplateId", Flag: "policy-template-id", Type: "*string", Required: true},
}

var fields_get_schema = []leanruntime.Field{
	{Name: "PolicyStoreId", Flag: "policy-store-id", Type: "*string", Required: true},
}

var fields_is_authorized = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "*types.ActionIdentifier", Required: false},
	{Name: "Context", Flag: "context", Type: "types.ContextDefinition", Required: false},
	{Name: "Entities", Flag: "entities", Type: "types.EntitiesDefinition", Required: false},
	{Name: "PolicyStoreId", Flag: "policy-store-id", Type: "*string", Required: true},
	{Name: "Principal", Flag: "principal", Type: "*types.EntityIdentifier", Required: false},
	{Name: "Resource", Flag: "resource", Type: "*types.EntityIdentifier", Required: false},
}

var fields_is_authorized_with_token = []leanruntime.Field{
	{Name: "AccessToken", Flag: "access-token", Type: "*string", Required: false},
	{Name: "Action", Flag: "action", Type: "*types.ActionIdentifier", Required: false},
	{Name: "Context", Flag: "context", Type: "types.ContextDefinition", Required: false},
	{Name: "Entities", Flag: "entities", Type: "types.EntitiesDefinition", Required: false},
	{Name: "IdentityToken", Flag: "identity-token", Type: "*string", Required: false},
	{Name: "PolicyStoreId", Flag: "policy-store-id", Type: "*string", Required: true},
	{Name: "Resource", Flag: "resource", Type: "*types.EntityIdentifier", Required: false},
}

var fields_list_identity_sources = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.IdentitySourceFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PolicyStoreId", Flag: "policy-store-id", Type: "*string", Required: true},
}

var fields_list_policies = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.PolicyFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PolicyStoreId", Flag: "policy-store-id", Type: "*string", Required: true},
}

var fields_list_policy_stores = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_policy_templates = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PolicyStoreId", Flag: "policy-store-id", Type: "*string", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_schema = []leanruntime.Field{
	{Name: "Definition", Flag: "definition", Type: "types.SchemaDefinition", Required: true},
	{Name: "PolicyStoreId", Flag: "policy-store-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_identity_source = []leanruntime.Field{
	{Name: "IdentitySourceId", Flag: "identity-source-id", Type: "*string", Required: true},
	{Name: "PolicyStoreId", Flag: "policy-store-id", Type: "*string", Required: true},
	{Name: "PrincipalEntityType", Flag: "principal-entity-type", Type: "*string", Required: false},
	{Name: "UpdateConfiguration", Flag: "update-configuration", Type: "types.UpdateConfiguration", Required: true},
}

var fields_update_policy = []leanruntime.Field{
	{Name: "Definition", Flag: "definition", Type: "types.UpdatePolicyDefinition", Required: false},
	{Name: "PolicyId", Flag: "policy-id", Type: "*string", Required: true},
	{Name: "PolicyStoreId", Flag: "policy-store-id", Type: "*string", Required: true},
}

var fields_update_policy_store = []leanruntime.Field{
	{Name: "DeletionProtection", Flag: "deletion-protection", Type: "types.DeletionProtection", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "PolicyStoreId", Flag: "policy-store-id", Type: "*string", Required: true},
	{Name: "ValidationSettings", Flag: "validation-settings", Type: "*types.ValidationSettings", Required: true},
}

var fields_update_policy_template = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "PolicyStoreId", Flag: "policy-store-id", Type: "*string", Required: true},
	{Name: "PolicyTemplateId", Flag: "policy-template-id", Type: "*string", Required: true},
	{Name: "Statement", Flag: "statement", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-get-policy": {
			Name:   "batch-get-policy",
			Fields: fields_batch_get_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetPolicy(ctx, input)
			},
		},
		"batch-is-authorized": {
			Name:   "batch-is-authorized",
			Fields: fields_batch_is_authorized,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchIsAuthorizedInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_is_authorized, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchIsAuthorized(ctx, input)
			},
		},
		"batch-is-authorized-with-token": {
			Name:   "batch-is-authorized-with-token",
			Fields: fields_batch_is_authorized_with_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchIsAuthorizedWithTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_is_authorized_with_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchIsAuthorizedWithToken(ctx, input)
			},
		},
		"create-identity-source": {
			Name:   "create-identity-source",
			Fields: fields_create_identity_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIdentitySourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_identity_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIdentitySource(ctx, input)
			},
		},
		"create-policy": {
			Name:   "create-policy",
			Fields: fields_create_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePolicy(ctx, input)
			},
		},
		"create-policy-store": {
			Name:   "create-policy-store",
			Fields: fields_create_policy_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePolicyStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_policy_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePolicyStore(ctx, input)
			},
		},
		"create-policy-template": {
			Name:   "create-policy-template",
			Fields: fields_create_policy_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePolicyTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_policy_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePolicyTemplate(ctx, input)
			},
		},
		"delete-identity-source": {
			Name:   "delete-identity-source",
			Fields: fields_delete_identity_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIdentitySourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_identity_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIdentitySource(ctx, input)
			},
		},
		"delete-policy": {
			Name:   "delete-policy",
			Fields: fields_delete_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePolicy(ctx, input)
			},
		},
		"delete-policy-store": {
			Name:   "delete-policy-store",
			Fields: fields_delete_policy_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePolicyStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_policy_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePolicyStore(ctx, input)
			},
		},
		"delete-policy-template": {
			Name:   "delete-policy-template",
			Fields: fields_delete_policy_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePolicyTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_policy_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePolicyTemplate(ctx, input)
			},
		},
		"get-identity-source": {
			Name:   "get-identity-source",
			Fields: fields_get_identity_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIdentitySourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_identity_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIdentitySource(ctx, input)
			},
		},
		"get-policy": {
			Name:   "get-policy",
			Fields: fields_get_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPolicy(ctx, input)
			},
		},
		"get-policy-store": {
			Name:   "get-policy-store",
			Fields: fields_get_policy_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPolicyStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_policy_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPolicyStore(ctx, input)
			},
		},
		"get-policy-template": {
			Name:   "get-policy-template",
			Fields: fields_get_policy_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPolicyTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_policy_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPolicyTemplate(ctx, input)
			},
		},
		"get-schema": {
			Name:   "get-schema",
			Fields: fields_get_schema,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSchemaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_schema, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSchema(ctx, input)
			},
		},
		"is-authorized": {
			Name:   "is-authorized",
			Fields: fields_is_authorized,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.IsAuthorizedInput{}
				if _, err := leanruntime.ApplyInput(input, fields_is_authorized, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.IsAuthorized(ctx, input)
			},
		},
		"is-authorized-with-token": {
			Name:   "is-authorized-with-token",
			Fields: fields_is_authorized_with_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.IsAuthorizedWithTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_is_authorized_with_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.IsAuthorizedWithToken(ctx, input)
			},
		},
		"list-identity-sources": {
			Name:   "list-identity-sources",
			Fields: fields_list_identity_sources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIdentitySourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_identity_sources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIdentitySources(ctx, input)
				}
				var results []*svc.ListIdentitySourcesOutput
				p := svc.NewListIdentitySourcesPaginator(client, input)
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
		"list-policies": {
			Name:   "list-policies",
			Fields: fields_list_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPolicies(ctx, input)
				}
				var results []*svc.ListPoliciesOutput
				p := svc.NewListPoliciesPaginator(client, input)
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
		"list-policy-stores": {
			Name:   "list-policy-stores",
			Fields: fields_list_policy_stores,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPolicyStoresInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_policy_stores, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPolicyStores(ctx, input)
				}
				var results []*svc.ListPolicyStoresOutput
				p := svc.NewListPolicyStoresPaginator(client, input)
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
		"list-policy-templates": {
			Name:   "list-policy-templates",
			Fields: fields_list_policy_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPolicyTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_policy_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPolicyTemplates(ctx, input)
				}
				var results []*svc.ListPolicyTemplatesOutput
				p := svc.NewListPolicyTemplatesPaginator(client, input)
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
		"put-schema": {
			Name:   "put-schema",
			Fields: fields_put_schema,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutSchemaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_schema, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutSchema(ctx, input)
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
		"update-identity-source": {
			Name:   "update-identity-source",
			Fields: fields_update_identity_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateIdentitySourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_identity_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateIdentitySource(ctx, input)
			},
		},
		"update-policy": {
			Name:   "update-policy",
			Fields: fields_update_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePolicy(ctx, input)
			},
		},
		"update-policy-store": {
			Name:   "update-policy-store",
			Fields: fields_update_policy_store,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePolicyStoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_policy_store, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePolicyStore(ctx, input)
			},
		},
		"update-policy-template": {
			Name:   "update-policy-template",
			Fields: fields_update_policy_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePolicyTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_policy_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePolicyTemplate(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("verifiedpermissions", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
