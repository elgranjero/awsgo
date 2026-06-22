package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/opensearchserverless"
)

var fields_batch_get_collection = []leanruntime.Field{
	{Name: "Ids", Flag: "ids", Type: "[]string", Required: false},
	{Name: "Names", Flag: "names", Type: "[]string", Required: false},
}

var fields_batch_get_collection_group = []leanruntime.Field{
	{Name: "Ids", Flag: "ids", Type: "[]string", Required: false},
	{Name: "Names", Flag: "names", Type: "[]string", Required: false},
}

var fields_batch_get_effective_lifecycle_policy = []leanruntime.Field{
	{Name: "ResourceIdentifiers", Flag: "resource-identifiers", Type: "[]types.LifecyclePolicyResourceIdentifier", Required: true},
}

var fields_batch_get_lifecycle_policy = []leanruntime.Field{
	{Name: "Identifiers", Flag: "identifiers", Type: "[]types.LifecyclePolicyIdentifier", Required: true},
}

var fields_batch_get_vpc_endpoint = []leanruntime.Field{
	{Name: "Ids", Flag: "ids", Type: "[]string", Required: true},
}

var fields_create_access_policy = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.AccessPolicyType", Required: true},
}

var fields_create_collection = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CollectionGroupName", Flag: "collection-group-name", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EncryptionConfig", Flag: "encryption-config", Type: "*types.EncryptionConfig", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "StandbyReplicas", Flag: "standby-replicas", Type: "types.StandbyReplicas", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Type", Flag: "type", Type: "types.CollectionType", Required: false},
	{Name: "VectorOptions", Flag: "vector-options", Type: "*types.VectorOptions", Required: false},
}

var fields_create_collection_group = []leanruntime.Field{
	{Name: "CapacityLimits", Flag: "capacity-limits", Type: "*types.CollectionGroupCapacityLimits", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "StandbyReplicas", Flag: "standby-replicas", Type: "types.StandbyReplicas", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_index = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: true},
	{Name: "IndexSchema", Flag: "index-schema", Type: "document.Interface", Required: false},
}

var fields_create_lifecycle_policy = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.LifecyclePolicyType", Required: true},
}

var fields_create_security_config = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IamFederationOptions", Flag: "iam-federation-options", Type: "*types.IamFederationConfigOptions", Required: false},
	{Name: "IamIdentityCenterOptions", Flag: "iam-identity-center-options", Type: "*types.CreateIamIdentityCenterConfigOptions", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SamlOptions", Flag: "saml-options", Type: "*types.SamlConfigOptions", Required: false},
	{Name: "Type", Flag: "type", Type: "types.SecurityConfigType", Required: true},
}

var fields_create_security_policy = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.SecurityPolicyType", Required: true},
}

var fields_create_vpc_endpoint = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: false},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: true},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_delete_access_policy = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.AccessPolicyType", Required: true},
}

var fields_delete_collection = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_collection_group = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_index = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: true},
}

var fields_delete_lifecycle_policy = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.LifecyclePolicyType", Required: true},
}

var fields_delete_security_config = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_security_policy = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.SecurityPolicyType", Required: true},
}

var fields_delete_vpc_endpoint = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_access_policy = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.AccessPolicyType", Required: true},
}

var fields_get_account_settings = []leanruntime.Field{}

var fields_get_index = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: true},
}

var fields_get_policies_stats = []leanruntime.Field{}

var fields_get_security_config = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_security_policy = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.SecurityPolicyType", Required: true},
}

var fields_list_access_policies = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Resource", Flag: "resource", Type: "[]string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.AccessPolicyType", Required: true},
}

var fields_list_collection_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_collections = []leanruntime.Field{
	{Name: "CollectionFilters", Flag: "collection-filters", Type: "*types.CollectionFilters", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_lifecycle_policies = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Resources", Flag: "resources", Type: "[]string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.LifecyclePolicyType", Required: true},
}

var fields_list_security_configs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.SecurityConfigType", Required: true},
}

var fields_list_security_policies = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Resource", Flag: "resource", Type: "[]string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.SecurityPolicyType", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_vpc_endpoints = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VpcEndpointFilters", Flag: "vpc-endpoint-filters", Type: "*types.VpcEndpointFilters", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_access_policy = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: false},
	{Name: "PolicyVersion", Flag: "policy-version", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.AccessPolicyType", Required: true},
}

var fields_update_account_settings = []leanruntime.Field{
	{Name: "CapacityLimits", Flag: "capacity-limits", Type: "*types.CapacityLimits", Required: false},
}

var fields_update_collection = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_update_collection_group = []leanruntime.Field{
	{Name: "CapacityLimits", Flag: "capacity-limits", Type: "*types.CollectionGroupCapacityLimits", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_update_index = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: true},
	{Name: "IndexSchema", Flag: "index-schema", Type: "document.Interface", Required: false},
}

var fields_update_lifecycle_policy = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: false},
	{Name: "PolicyVersion", Flag: "policy-version", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.LifecyclePolicyType", Required: true},
}

var fields_update_security_config = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ConfigVersion", Flag: "config-version", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IamFederationOptions", Flag: "iam-federation-options", Type: "*types.IamFederationConfigOptions", Required: false},
	{Name: "IamIdentityCenterOptionsUpdates", Flag: "iam-identity-center-options-updates", Type: "*types.UpdateIamIdentityCenterConfigOptions", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "SamlOptions", Flag: "saml-options", Type: "*types.SamlConfigOptions", Required: false},
}

var fields_update_security_policy = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: false},
	{Name: "PolicyVersion", Flag: "policy-version", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.SecurityPolicyType", Required: true},
}

var fields_update_vpc_endpoint = []leanruntime.Field{
	{Name: "AddSecurityGroupIds", Flag: "add-security-group-ids", Type: "[]string", Required: false},
	{Name: "AddSubnetIds", Flag: "add-subnet-ids", Type: "[]string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "RemoveSecurityGroupIds", Flag: "remove-security-group-ids", Type: "[]string", Required: false},
	{Name: "RemoveSubnetIds", Flag: "remove-subnet-ids", Type: "[]string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-get-collection": {
			Name:   "batch-get-collection",
			Fields: fields_batch_get_collection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetCollectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_collection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetCollection(ctx, input)
			},
		},
		"batch-get-collection-group": {
			Name:   "batch-get-collection-group",
			Fields: fields_batch_get_collection_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetCollectionGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_collection_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetCollectionGroup(ctx, input)
			},
		},
		"batch-get-effective-lifecycle-policy": {
			Name:   "batch-get-effective-lifecycle-policy",
			Fields: fields_batch_get_effective_lifecycle_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetEffectiveLifecyclePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_effective_lifecycle_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetEffectiveLifecyclePolicy(ctx, input)
			},
		},
		"batch-get-lifecycle-policy": {
			Name:   "batch-get-lifecycle-policy",
			Fields: fields_batch_get_lifecycle_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetLifecyclePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_lifecycle_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetLifecyclePolicy(ctx, input)
			},
		},
		"batch-get-vpc-endpoint": {
			Name:   "batch-get-vpc-endpoint",
			Fields: fields_batch_get_vpc_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetVpcEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_vpc_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetVpcEndpoint(ctx, input)
			},
		},
		"create-access-policy": {
			Name:   "create-access-policy",
			Fields: fields_create_access_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAccessPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_access_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAccessPolicy(ctx, input)
			},
		},
		"create-collection": {
			Name:   "create-collection",
			Fields: fields_create_collection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCollectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_collection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCollection(ctx, input)
			},
		},
		"create-collection-group": {
			Name:   "create-collection-group",
			Fields: fields_create_collection_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCollectionGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_collection_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCollectionGroup(ctx, input)
			},
		},
		"create-index": {
			Name:   "create-index",
			Fields: fields_create_index,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIndexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_index, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIndex(ctx, input)
			},
		},
		"create-lifecycle-policy": {
			Name:   "create-lifecycle-policy",
			Fields: fields_create_lifecycle_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLifecyclePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_lifecycle_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLifecyclePolicy(ctx, input)
			},
		},
		"create-security-config": {
			Name:   "create-security-config",
			Fields: fields_create_security_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSecurityConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_security_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSecurityConfig(ctx, input)
			},
		},
		"create-security-policy": {
			Name:   "create-security-policy",
			Fields: fields_create_security_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSecurityPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_security_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSecurityPolicy(ctx, input)
			},
		},
		"create-vpc-endpoint": {
			Name:   "create-vpc-endpoint",
			Fields: fields_create_vpc_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVpcEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_vpc_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVpcEndpoint(ctx, input)
			},
		},
		"delete-access-policy": {
			Name:   "delete-access-policy",
			Fields: fields_delete_access_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccessPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_access_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccessPolicy(ctx, input)
			},
		},
		"delete-collection": {
			Name:   "delete-collection",
			Fields: fields_delete_collection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCollectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_collection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCollection(ctx, input)
			},
		},
		"delete-collection-group": {
			Name:   "delete-collection-group",
			Fields: fields_delete_collection_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCollectionGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_collection_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCollectionGroup(ctx, input)
			},
		},
		"delete-index": {
			Name:   "delete-index",
			Fields: fields_delete_index,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIndexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_index, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIndex(ctx, input)
			},
		},
		"delete-lifecycle-policy": {
			Name:   "delete-lifecycle-policy",
			Fields: fields_delete_lifecycle_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLifecyclePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_lifecycle_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLifecyclePolicy(ctx, input)
			},
		},
		"delete-security-config": {
			Name:   "delete-security-config",
			Fields: fields_delete_security_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSecurityConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_security_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSecurityConfig(ctx, input)
			},
		},
		"delete-security-policy": {
			Name:   "delete-security-policy",
			Fields: fields_delete_security_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSecurityPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_security_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSecurityPolicy(ctx, input)
			},
		},
		"delete-vpc-endpoint": {
			Name:   "delete-vpc-endpoint",
			Fields: fields_delete_vpc_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVpcEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vpc_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVpcEndpoint(ctx, input)
			},
		},
		"get-access-policy": {
			Name:   "get-access-policy",
			Fields: fields_get_access_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccessPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_access_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccessPolicy(ctx, input)
			},
		},
		"get-account-settings": {
			Name:   "get-account-settings",
			Fields: fields_get_account_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccountSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_account_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccountSettings(ctx, input)
			},
		},
		"get-index": {
			Name:   "get-index",
			Fields: fields_get_index,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIndexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_index, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIndex(ctx, input)
			},
		},
		"get-policies-stats": {
			Name:   "get-policies-stats",
			Fields: fields_get_policies_stats,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPoliciesStatsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_policies_stats, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPoliciesStats(ctx, input)
			},
		},
		"get-security-config": {
			Name:   "get-security-config",
			Fields: fields_get_security_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSecurityConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_security_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSecurityConfig(ctx, input)
			},
		},
		"get-security-policy": {
			Name:   "get-security-policy",
			Fields: fields_get_security_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSecurityPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_security_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSecurityPolicy(ctx, input)
			},
		},
		"list-access-policies": {
			Name:   "list-access-policies",
			Fields: fields_list_access_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccessPoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_access_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccessPolicies(ctx, input)
				}
				var results []*svc.ListAccessPoliciesOutput
				p := svc.NewListAccessPoliciesPaginator(client, input)
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
		"list-collection-groups": {
			Name:   "list-collection-groups",
			Fields: fields_list_collection_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCollectionGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_collection_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCollectionGroups(ctx, input)
				}
				var results []*svc.ListCollectionGroupsOutput
				p := svc.NewListCollectionGroupsPaginator(client, input)
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
		"list-collections": {
			Name:   "list-collections",
			Fields: fields_list_collections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCollectionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_collections, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCollections(ctx, input)
				}
				var results []*svc.ListCollectionsOutput
				p := svc.NewListCollectionsPaginator(client, input)
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
		"list-lifecycle-policies": {
			Name:   "list-lifecycle-policies",
			Fields: fields_list_lifecycle_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLifecyclePoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_lifecycle_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLifecyclePolicies(ctx, input)
				}
				var results []*svc.ListLifecyclePoliciesOutput
				p := svc.NewListLifecyclePoliciesPaginator(client, input)
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
		"list-security-configs": {
			Name:   "list-security-configs",
			Fields: fields_list_security_configs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSecurityConfigsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_security_configs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSecurityConfigs(ctx, input)
				}
				var results []*svc.ListSecurityConfigsOutput
				p := svc.NewListSecurityConfigsPaginator(client, input)
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
		"list-security-policies": {
			Name:   "list-security-policies",
			Fields: fields_list_security_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSecurityPoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_security_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSecurityPolicies(ctx, input)
				}
				var results []*svc.ListSecurityPoliciesOutput
				p := svc.NewListSecurityPoliciesPaginator(client, input)
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
		"list-vpc-endpoints": {
			Name:   "list-vpc-endpoints",
			Fields: fields_list_vpc_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVpcEndpointsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_vpc_endpoints, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListVpcEndpoints(ctx, input)
				}
				var results []*svc.ListVpcEndpointsOutput
				p := svc.NewListVpcEndpointsPaginator(client, input)
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
		"update-access-policy": {
			Name:   "update-access-policy",
			Fields: fields_update_access_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAccessPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_access_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAccessPolicy(ctx, input)
			},
		},
		"update-account-settings": {
			Name:   "update-account-settings",
			Fields: fields_update_account_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAccountSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_account_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAccountSettings(ctx, input)
			},
		},
		"update-collection": {
			Name:   "update-collection",
			Fields: fields_update_collection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCollectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_collection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCollection(ctx, input)
			},
		},
		"update-collection-group": {
			Name:   "update-collection-group",
			Fields: fields_update_collection_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCollectionGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_collection_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCollectionGroup(ctx, input)
			},
		},
		"update-index": {
			Name:   "update-index",
			Fields: fields_update_index,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateIndexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_index, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateIndex(ctx, input)
			},
		},
		"update-lifecycle-policy": {
			Name:   "update-lifecycle-policy",
			Fields: fields_update_lifecycle_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLifecyclePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_lifecycle_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLifecyclePolicy(ctx, input)
			},
		},
		"update-security-config": {
			Name:   "update-security-config",
			Fields: fields_update_security_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSecurityConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_security_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSecurityConfig(ctx, input)
			},
		},
		"update-security-policy": {
			Name:   "update-security-policy",
			Fields: fields_update_security_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSecurityPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_security_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSecurityPolicy(ctx, input)
			},
		},
		"update-vpc-endpoint": {
			Name:   "update-vpc-endpoint",
			Fields: fields_update_vpc_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateVpcEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_vpc_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateVpcEndpoint(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("opensearchserverless", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
