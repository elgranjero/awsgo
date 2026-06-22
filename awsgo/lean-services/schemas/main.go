package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/schemas"
)

var fields_create_discoverer = []leanruntime.Field{
	{Name: "CrossAccount", Flag: "cross-account", Type: "*bool", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "SourceArn", Flag: "source-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_registry = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "RegistryName", Flag: "registry-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_schema = []leanruntime.Field{
	{Name: "Content", Flag: "content", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "RegistryName", Flag: "registry-name", Type: "*string", Required: true},
	{Name: "SchemaName", Flag: "schema-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.Type", Required: true},
}

var fields_delete_discoverer = []leanruntime.Field{
	{Name: "DiscovererId", Flag: "discoverer-id", Type: "*string", Required: true},
}

var fields_delete_registry = []leanruntime.Field{
	{Name: "RegistryName", Flag: "registry-name", Type: "*string", Required: true},
}

var fields_delete_resource_policy = []leanruntime.Field{
	{Name: "RegistryName", Flag: "registry-name", Type: "*string", Required: false},
}

var fields_delete_schema = []leanruntime.Field{
	{Name: "RegistryName", Flag: "registry-name", Type: "*string", Required: true},
	{Name: "SchemaName", Flag: "schema-name", Type: "*string", Required: true},
}

var fields_delete_schema_version = []leanruntime.Field{
	{Name: "RegistryName", Flag: "registry-name", Type: "*string", Required: true},
	{Name: "SchemaName", Flag: "schema-name", Type: "*string", Required: true},
	{Name: "SchemaVersion", Flag: "schema-version", Type: "*string", Required: true},
}

var fields_describe_code_binding = []leanruntime.Field{
	{Name: "Language", Flag: "language", Type: "*string", Required: true},
	{Name: "RegistryName", Flag: "registry-name", Type: "*string", Required: true},
	{Name: "SchemaName", Flag: "schema-name", Type: "*string", Required: true},
	{Name: "SchemaVersion", Flag: "schema-version", Type: "*string", Required: false},
}

var fields_describe_discoverer = []leanruntime.Field{
	{Name: "DiscovererId", Flag: "discoverer-id", Type: "*string", Required: true},
}

var fields_describe_registry = []leanruntime.Field{
	{Name: "RegistryName", Flag: "registry-name", Type: "*string", Required: true},
}

var fields_describe_schema = []leanruntime.Field{
	{Name: "RegistryName", Flag: "registry-name", Type: "*string", Required: true},
	{Name: "SchemaName", Flag: "schema-name", Type: "*string", Required: true},
	{Name: "SchemaVersion", Flag: "schema-version", Type: "*string", Required: false},
}

var fields_export_schema = []leanruntime.Field{
	{Name: "RegistryName", Flag: "registry-name", Type: "*string", Required: true},
	{Name: "SchemaName", Flag: "schema-name", Type: "*string", Required: true},
	{Name: "SchemaVersion", Flag: "schema-version", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "*string", Required: true},
}

var fields_get_code_binding_source = []leanruntime.Field{
	{Name: "Language", Flag: "language", Type: "*string", Required: true},
	{Name: "RegistryName", Flag: "registry-name", Type: "*string", Required: true},
	{Name: "SchemaName", Flag: "schema-name", Type: "*string", Required: true},
	{Name: "SchemaVersion", Flag: "schema-version", Type: "*string", Required: false},
}

var fields_get_discovered_schema = []leanruntime.Field{
	{Name: "Events", Flag: "events", Type: "[]string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.Type", Required: true},
}

var fields_get_resource_policy = []leanruntime.Field{
	{Name: "RegistryName", Flag: "registry-name", Type: "*string", Required: false},
}

var fields_list_discoverers = []leanruntime.Field{
	{Name: "DiscovererIdPrefix", Flag: "discoverer-id-prefix", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SourceArnPrefix", Flag: "source-arn-prefix", Type: "*string", Required: false},
}

var fields_list_registries = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RegistryNamePrefix", Flag: "registry-name-prefix", Type: "*string", Required: false},
	{Name: "Scope", Flag: "scope", Type: "*string", Required: false},
}

var fields_list_schema_versions = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RegistryName", Flag: "registry-name", Type: "*string", Required: true},
	{Name: "SchemaName", Flag: "schema-name", Type: "*string", Required: true},
}

var fields_list_schemas = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RegistryName", Flag: "registry-name", Type: "*string", Required: true},
	{Name: "SchemaNamePrefix", Flag: "schema-name-prefix", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_code_binding = []leanruntime.Field{
	{Name: "Language", Flag: "language", Type: "*string", Required: true},
	{Name: "RegistryName", Flag: "registry-name", Type: "*string", Required: true},
	{Name: "SchemaName", Flag: "schema-name", Type: "*string", Required: true},
	{Name: "SchemaVersion", Flag: "schema-version", Type: "*string", Required: false},
}

var fields_put_resource_policy = []leanruntime.Field{
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
	{Name: "RegistryName", Flag: "registry-name", Type: "*string", Required: false},
	{Name: "RevisionId", Flag: "revision-id", Type: "*string", Required: false},
}

var fields_search_schemas = []leanruntime.Field{
	{Name: "Keywords", Flag: "keywords", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RegistryName", Flag: "registry-name", Type: "*string", Required: true},
}

var fields_start_discoverer = []leanruntime.Field{
	{Name: "DiscovererId", Flag: "discoverer-id", Type: "*string", Required: true},
}

var fields_stop_discoverer = []leanruntime.Field{
	{Name: "DiscovererId", Flag: "discoverer-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_discoverer = []leanruntime.Field{
	{Name: "CrossAccount", Flag: "cross-account", Type: "*bool", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DiscovererId", Flag: "discoverer-id", Type: "*string", Required: true},
}

var fields_update_registry = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "RegistryName", Flag: "registry-name", Type: "*string", Required: true},
}

var fields_update_schema = []leanruntime.Field{
	{Name: "ClientTokenId", Flag: "client-token-id", Type: "*string", Required: false},
	{Name: "Content", Flag: "content", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "RegistryName", Flag: "registry-name", Type: "*string", Required: true},
	{Name: "SchemaName", Flag: "schema-name", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.Type", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-discoverer": {
			Name:   "create-discoverer",
			Fields: fields_create_discoverer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDiscovererInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_discoverer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDiscoverer(ctx, input)
			},
		},
		"create-registry": {
			Name:   "create-registry",
			Fields: fields_create_registry,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRegistryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_registry, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRegistry(ctx, input)
			},
		},
		"create-schema": {
			Name:   "create-schema",
			Fields: fields_create_schema,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSchemaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_schema, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSchema(ctx, input)
			},
		},
		"delete-discoverer": {
			Name:   "delete-discoverer",
			Fields: fields_delete_discoverer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDiscovererInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_discoverer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDiscoverer(ctx, input)
			},
		},
		"delete-registry": {
			Name:   "delete-registry",
			Fields: fields_delete_registry,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRegistryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_registry, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRegistry(ctx, input)
			},
		},
		"delete-resource-policy": {
			Name:   "delete-resource-policy",
			Fields: fields_delete_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResourcePolicy(ctx, input)
			},
		},
		"delete-schema": {
			Name:   "delete-schema",
			Fields: fields_delete_schema,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSchemaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_schema, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSchema(ctx, input)
			},
		},
		"delete-schema-version": {
			Name:   "delete-schema-version",
			Fields: fields_delete_schema_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSchemaVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_schema_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSchemaVersion(ctx, input)
			},
		},
		"describe-code-binding": {
			Name:   "describe-code-binding",
			Fields: fields_describe_code_binding,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCodeBindingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_code_binding, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCodeBinding(ctx, input)
			},
		},
		"describe-discoverer": {
			Name:   "describe-discoverer",
			Fields: fields_describe_discoverer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDiscovererInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_discoverer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDiscoverer(ctx, input)
			},
		},
		"describe-registry": {
			Name:   "describe-registry",
			Fields: fields_describe_registry,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRegistryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_registry, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRegistry(ctx, input)
			},
		},
		"describe-schema": {
			Name:   "describe-schema",
			Fields: fields_describe_schema,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSchemaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_schema, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSchema(ctx, input)
			},
		},
		"export-schema": {
			Name:   "export-schema",
			Fields: fields_export_schema,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExportSchemaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_export_schema, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExportSchema(ctx, input)
			},
		},
		"get-code-binding-source": {
			Name:   "get-code-binding-source",
			Fields: fields_get_code_binding_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCodeBindingSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_code_binding_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCodeBindingSource(ctx, input)
			},
		},
		"get-discovered-schema": {
			Name:   "get-discovered-schema",
			Fields: fields_get_discovered_schema,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDiscoveredSchemaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_discovered_schema, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDiscoveredSchema(ctx, input)
			},
		},
		"get-resource-policy": {
			Name:   "get-resource-policy",
			Fields: fields_get_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResourcePolicy(ctx, input)
			},
		},
		"list-discoverers": {
			Name:   "list-discoverers",
			Fields: fields_list_discoverers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDiscoverersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_discoverers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDiscoverers(ctx, input)
				}
				var results []*svc.ListDiscoverersOutput
				p := svc.NewListDiscoverersPaginator(client, input)
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
		"list-registries": {
			Name:   "list-registries",
			Fields: fields_list_registries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRegistriesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_registries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRegistries(ctx, input)
				}
				var results []*svc.ListRegistriesOutput
				p := svc.NewListRegistriesPaginator(client, input)
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
		"list-schema-versions": {
			Name:   "list-schema-versions",
			Fields: fields_list_schema_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSchemaVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_schema_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSchemaVersions(ctx, input)
				}
				var results []*svc.ListSchemaVersionsOutput
				p := svc.NewListSchemaVersionsPaginator(client, input)
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
		"list-schemas": {
			Name:   "list-schemas",
			Fields: fields_list_schemas,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSchemasInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_schemas, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSchemas(ctx, input)
				}
				var results []*svc.ListSchemasOutput
				p := svc.NewListSchemasPaginator(client, input)
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
		"put-code-binding": {
			Name:   "put-code-binding",
			Fields: fields_put_code_binding,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutCodeBindingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_code_binding, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutCodeBinding(ctx, input)
			},
		},
		"put-resource-policy": {
			Name:   "put-resource-policy",
			Fields: fields_put_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutResourcePolicy(ctx, input)
			},
		},
		"search-schemas": {
			Name:   "search-schemas",
			Fields: fields_search_schemas,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchSchemasInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_schemas, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchSchemas(ctx, input)
				}
				var results []*svc.SearchSchemasOutput
				p := svc.NewSearchSchemasPaginator(client, input)
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
		"start-discoverer": {
			Name:   "start-discoverer",
			Fields: fields_start_discoverer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDiscovererInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_discoverer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDiscoverer(ctx, input)
			},
		},
		"stop-discoverer": {
			Name:   "stop-discoverer",
			Fields: fields_stop_discoverer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopDiscovererInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_discoverer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopDiscoverer(ctx, input)
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
		"update-discoverer": {
			Name:   "update-discoverer",
			Fields: fields_update_discoverer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDiscovererInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_discoverer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDiscoverer(ctx, input)
			},
		},
		"update-registry": {
			Name:   "update-registry",
			Fields: fields_update_registry,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRegistryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_registry, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRegistry(ctx, input)
			},
		},
		"update-schema": {
			Name:   "update-schema",
			Fields: fields_update_schema,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSchemaInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_schema, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSchema(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("schemas", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
