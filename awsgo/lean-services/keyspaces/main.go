package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/keyspaces"
)

var fields_create_keyspace = []leanruntime.Field{
	{Name: "KeyspaceName", Flag: "keyspace-name", Type: "*string", Required: true},
	{Name: "ReplicationSpecification", Flag: "replication-specification", Type: "*types.ReplicationSpecification", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_table = []leanruntime.Field{
	{Name: "AutoScalingSpecification", Flag: "auto-scaling-specification", Type: "*types.AutoScalingSpecification", Required: false},
	{Name: "CapacitySpecification", Flag: "capacity-specification", Type: "*types.CapacitySpecification", Required: false},
	{Name: "CdcSpecification", Flag: "cdc-specification", Type: "*types.CdcSpecification", Required: false},
	{Name: "ClientSideTimestamps", Flag: "client-side-timestamps", Type: "*types.ClientSideTimestamps", Required: false},
	{Name: "Comment", Flag: "comment", Type: "*types.Comment", Required: false},
	{Name: "DefaultTimeToLive", Flag: "default-time-to-live", Type: "*int32", Required: false},
	{Name: "EncryptionSpecification", Flag: "encryption-specification", Type: "*types.EncryptionSpecification", Required: false},
	{Name: "KeyspaceName", Flag: "keyspace-name", Type: "*string", Required: true},
	{Name: "PointInTimeRecovery", Flag: "point-in-time-recovery", Type: "*types.PointInTimeRecovery", Required: false},
	{Name: "ReplicaSpecifications", Flag: "replica-specifications", Type: "[]types.ReplicaSpecification", Required: false},
	{Name: "SchemaDefinition", Flag: "schema-definition", Type: "*types.SchemaDefinition", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Ttl", Flag: "ttl", Type: "*types.TimeToLive", Required: false},
	{Name: "WarmThroughputSpecification", Flag: "warm-throughput-specification", Type: "*types.WarmThroughputSpecification", Required: false},
}

var fields_create_type = []leanruntime.Field{
	{Name: "FieldDefinitions", Flag: "field-definitions", Type: "[]types.FieldDefinition", Required: true},
	{Name: "KeyspaceName", Flag: "keyspace-name", Type: "*string", Required: true},
	{Name: "TypeName", Flag: "type-name", Type: "*string", Required: true},
}

var fields_delete_keyspace = []leanruntime.Field{
	{Name: "KeyspaceName", Flag: "keyspace-name", Type: "*string", Required: true},
}

var fields_delete_table = []leanruntime.Field{
	{Name: "KeyspaceName", Flag: "keyspace-name", Type: "*string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_delete_type = []leanruntime.Field{
	{Name: "KeyspaceName", Flag: "keyspace-name", Type: "*string", Required: true},
	{Name: "TypeName", Flag: "type-name", Type: "*string", Required: true},
}

var fields_get_keyspace = []leanruntime.Field{
	{Name: "KeyspaceName", Flag: "keyspace-name", Type: "*string", Required: true},
}

var fields_get_table = []leanruntime.Field{
	{Name: "KeyspaceName", Flag: "keyspace-name", Type: "*string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_get_table_auto_scaling_settings = []leanruntime.Field{
	{Name: "KeyspaceName", Flag: "keyspace-name", Type: "*string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_get_type = []leanruntime.Field{
	{Name: "KeyspaceName", Flag: "keyspace-name", Type: "*string", Required: true},
	{Name: "TypeName", Flag: "type-name", Type: "*string", Required: true},
}

var fields_list_keyspaces = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tables = []leanruntime.Field{
	{Name: "KeyspaceName", Flag: "keyspace-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_types = []leanruntime.Field{
	{Name: "KeyspaceName", Flag: "keyspace-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_restore_table = []leanruntime.Field{
	{Name: "AutoScalingSpecification", Flag: "auto-scaling-specification", Type: "*types.AutoScalingSpecification", Required: false},
	{Name: "CapacitySpecificationOverride", Flag: "capacity-specification-override", Type: "*types.CapacitySpecification", Required: false},
	{Name: "EncryptionSpecificationOverride", Flag: "encryption-specification-override", Type: "*types.EncryptionSpecification", Required: false},
	{Name: "PointInTimeRecoveryOverride", Flag: "point-in-time-recovery-override", Type: "*types.PointInTimeRecovery", Required: false},
	{Name: "ReplicaSpecifications", Flag: "replica-specifications", Type: "[]types.ReplicaSpecification", Required: false},
	{Name: "RestoreTimestamp", Flag: "restore-timestamp", Type: "*time.Time", Required: false},
	{Name: "SourceKeyspaceName", Flag: "source-keyspace-name", Type: "*string", Required: true},
	{Name: "SourceTableName", Flag: "source-table-name", Type: "*string", Required: true},
	{Name: "TagsOverride", Flag: "tags-override", Type: "[]types.Tag", Required: false},
	{Name: "TargetKeyspaceName", Flag: "target-keyspace-name", Type: "*string", Required: true},
	{Name: "TargetTableName", Flag: "target-table-name", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_update_keyspace = []leanruntime.Field{
	{Name: "ClientSideTimestamps", Flag: "client-side-timestamps", Type: "*types.ClientSideTimestamps", Required: false},
	{Name: "KeyspaceName", Flag: "keyspace-name", Type: "*string", Required: true},
	{Name: "ReplicationSpecification", Flag: "replication-specification", Type: "*types.ReplicationSpecification", Required: true},
}

var fields_update_table = []leanruntime.Field{
	{Name: "AddColumns", Flag: "add-columns", Type: "[]types.ColumnDefinition", Required: false},
	{Name: "AutoScalingSpecification", Flag: "auto-scaling-specification", Type: "*types.AutoScalingSpecification", Required: false},
	{Name: "CapacitySpecification", Flag: "capacity-specification", Type: "*types.CapacitySpecification", Required: false},
	{Name: "CdcSpecification", Flag: "cdc-specification", Type: "*types.CdcSpecification", Required: false},
	{Name: "ClientSideTimestamps", Flag: "client-side-timestamps", Type: "*types.ClientSideTimestamps", Required: false},
	{Name: "DefaultTimeToLive", Flag: "default-time-to-live", Type: "*int32", Required: false},
	{Name: "EncryptionSpecification", Flag: "encryption-specification", Type: "*types.EncryptionSpecification", Required: false},
	{Name: "KeyspaceName", Flag: "keyspace-name", Type: "*string", Required: true},
	{Name: "PointInTimeRecovery", Flag: "point-in-time-recovery", Type: "*types.PointInTimeRecovery", Required: false},
	{Name: "ReplicaSpecifications", Flag: "replica-specifications", Type: "[]types.ReplicaSpecification", Required: false},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
	{Name: "Ttl", Flag: "ttl", Type: "*types.TimeToLive", Required: false},
	{Name: "WarmThroughputSpecification", Flag: "warm-throughput-specification", Type: "*types.WarmThroughputSpecification", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-keyspace": {
			Name:   "create-keyspace",
			Fields: fields_create_keyspace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateKeyspaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_keyspace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateKeyspace(ctx, input)
			},
		},
		"create-table": {
			Name:   "create-table",
			Fields: fields_create_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTable(ctx, input)
			},
		},
		"create-type": {
			Name:   "create-type",
			Fields: fields_create_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateType(ctx, input)
			},
		},
		"delete-keyspace": {
			Name:   "delete-keyspace",
			Fields: fields_delete_keyspace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteKeyspaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_keyspace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteKeyspace(ctx, input)
			},
		},
		"delete-table": {
			Name:   "delete-table",
			Fields: fields_delete_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTable(ctx, input)
			},
		},
		"delete-type": {
			Name:   "delete-type",
			Fields: fields_delete_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteType(ctx, input)
			},
		},
		"get-keyspace": {
			Name:   "get-keyspace",
			Fields: fields_get_keyspace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetKeyspaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_keyspace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetKeyspace(ctx, input)
			},
		},
		"get-table": {
			Name:   "get-table",
			Fields: fields_get_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTable(ctx, input)
			},
		},
		"get-table-auto-scaling-settings": {
			Name:   "get-table-auto-scaling-settings",
			Fields: fields_get_table_auto_scaling_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTableAutoScalingSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_table_auto_scaling_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTableAutoScalingSettings(ctx, input)
			},
		},
		"get-type": {
			Name:   "get-type",
			Fields: fields_get_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetType(ctx, input)
			},
		},
		"list-keyspaces": {
			Name:   "list-keyspaces",
			Fields: fields_list_keyspaces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListKeyspacesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_keyspaces, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListKeyspaces(ctx, input)
				}
				var results []*svc.ListKeyspacesOutput
				p := svc.NewListKeyspacesPaginator(client, input)
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
		"list-tables": {
			Name:   "list-tables",
			Fields: fields_list_tables,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTablesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tables, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTables(ctx, input)
				}
				var results []*svc.ListTablesOutput
				p := svc.NewListTablesPaginator(client, input)
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
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tags_for_resource, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTagsForResource(ctx, input)
				}
				var results []*svc.ListTagsForResourceOutput
				p := svc.NewListTagsForResourcePaginator(client, input)
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
		"list-types": {
			Name:   "list-types",
			Fields: fields_list_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTypesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_types, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTypes(ctx, input)
				}
				var results []*svc.ListTypesOutput
				p := svc.NewListTypesPaginator(client, input)
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
		"restore-table": {
			Name:   "restore-table",
			Fields: fields_restore_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestoreTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestoreTable(ctx, input)
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
		"update-keyspace": {
			Name:   "update-keyspace",
			Fields: fields_update_keyspace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateKeyspaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_keyspace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateKeyspace(ctx, input)
			},
		},
		"update-table": {
			Name:   "update-table",
			Fields: fields_update_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTable(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("keyspaces", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
