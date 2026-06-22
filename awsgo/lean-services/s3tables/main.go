package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/s3tables"
)

var fields_create_namespace = []leanruntime.Field{
	{Name: "Namespace", Flag: "namespace", Type: "[]string", Required: true},
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: true},
}

var fields_create_table = []leanruntime.Field{
	{Name: "EncryptionConfiguration", Flag: "encryption-configuration", Type: "*types.EncryptionConfiguration", Required: false},
	{Name: "Format", Flag: "format", Type: "types.OpenTableFormat", Required: true},
	{Name: "Metadata", Flag: "metadata", Type: "types.TableMetadata", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "StorageClassConfiguration", Flag: "storage-class-configuration", Type: "*types.StorageClassConfiguration", Required: false},
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_table_bucket = []leanruntime.Field{
	{Name: "EncryptionConfiguration", Flag: "encryption-configuration", Type: "*types.EncryptionConfiguration", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "StorageClassConfiguration", Flag: "storage-class-configuration", Type: "*types.StorageClassConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_namespace = []leanruntime.Field{
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: true},
}

var fields_delete_table = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: true},
	{Name: "VersionToken", Flag: "version-token", Type: "*string", Required: false},
}

var fields_delete_table_bucket = []leanruntime.Field{
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: true},
}

var fields_delete_table_bucket_encryption = []leanruntime.Field{
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: true},
}

var fields_delete_table_bucket_metrics_configuration = []leanruntime.Field{
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: true},
}

var fields_delete_table_bucket_policy = []leanruntime.Field{
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: true},
}

var fields_delete_table_bucket_replication = []leanruntime.Field{
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: true},
	{Name: "VersionToken", Flag: "version-token", Type: "*string", Required: false},
}

var fields_delete_table_policy = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: true},
}

var fields_delete_table_replication = []leanruntime.Field{
	{Name: "TableArn", Flag: "table-arn", Type: "*string", Required: true},
	{Name: "VersionToken", Flag: "version-token", Type: "*string", Required: true},
}

var fields_get_namespace = []leanruntime.Field{
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: true},
}

var fields_get_table = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: false},
	{Name: "TableArn", Flag: "table-arn", Type: "*string", Required: false},
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: false},
}

var fields_get_table_bucket = []leanruntime.Field{
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: true},
}

var fields_get_table_bucket_encryption = []leanruntime.Field{
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: true},
}

var fields_get_table_bucket_maintenance_configuration = []leanruntime.Field{
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: true},
}

var fields_get_table_bucket_metrics_configuration = []leanruntime.Field{
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: true},
}

var fields_get_table_bucket_policy = []leanruntime.Field{
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: true},
}

var fields_get_table_bucket_replication = []leanruntime.Field{
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: true},
}

var fields_get_table_bucket_storage_class = []leanruntime.Field{
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: true},
}

var fields_get_table_encryption = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: true},
}

var fields_get_table_maintenance_configuration = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: true},
}

var fields_get_table_maintenance_job_status = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: true},
}

var fields_get_table_metadata_location = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: true},
}

var fields_get_table_policy = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: true},
}

var fields_get_table_record_expiration_configuration = []leanruntime.Field{
	{Name: "TableArn", Flag: "table-arn", Type: "*string", Required: true},
}

var fields_get_table_record_expiration_job_status = []leanruntime.Field{
	{Name: "TableArn", Flag: "table-arn", Type: "*string", Required: true},
}

var fields_get_table_replication = []leanruntime.Field{
	{Name: "TableArn", Flag: "table-arn", Type: "*string", Required: true},
}

var fields_get_table_replication_status = []leanruntime.Field{
	{Name: "TableArn", Flag: "table-arn", Type: "*string", Required: true},
}

var fields_get_table_storage_class = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: true},
}

var fields_list_namespaces = []leanruntime.Field{
	{Name: "ContinuationToken", Flag: "continuation-token", Type: "*string", Required: false},
	{Name: "MaxNamespaces", Flag: "max-namespaces", Type: "*int32", Required: false},
	{Name: "Prefix", Flag: "prefix", Type: "*string", Required: false},
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: true},
}

var fields_list_table_buckets = []leanruntime.Field{
	{Name: "ContinuationToken", Flag: "continuation-token", Type: "*string", Required: false},
	{Name: "MaxBuckets", Flag: "max-buckets", Type: "*int32", Required: false},
	{Name: "Prefix", Flag: "prefix", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.TableBucketType", Required: false},
}

var fields_list_tables = []leanruntime.Field{
	{Name: "ContinuationToken", Flag: "continuation-token", Type: "*string", Required: false},
	{Name: "MaxTables", Flag: "max-tables", Type: "*int32", Required: false},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: false},
	{Name: "Prefix", Flag: "prefix", Type: "*string", Required: false},
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_table_bucket_encryption = []leanruntime.Field{
	{Name: "EncryptionConfiguration", Flag: "encryption-configuration", Type: "*types.EncryptionConfiguration", Required: true},
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: true},
}

var fields_put_table_bucket_maintenance_configuration = []leanruntime.Field{
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.TableBucketMaintenanceType", Required: true},
	{Name: "Value", Flag: "value", Type: "*types.TableBucketMaintenanceConfigurationValue", Required: true},
}

var fields_put_table_bucket_metrics_configuration = []leanruntime.Field{
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: true},
}

var fields_put_table_bucket_policy = []leanruntime.Field{
	{Name: "ResourcePolicy", Flag: "resource-policy", Type: "*string", Required: true},
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: true},
}

var fields_put_table_bucket_replication = []leanruntime.Field{
	{Name: "Configuration", Flag: "configuration", Type: "*types.TableBucketReplicationConfiguration", Required: true},
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: true},
	{Name: "VersionToken", Flag: "version-token", Type: "*string", Required: false},
}

var fields_put_table_bucket_storage_class = []leanruntime.Field{
	{Name: "StorageClassConfiguration", Flag: "storage-class-configuration", Type: "*types.StorageClassConfiguration", Required: true},
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: true},
}

var fields_put_table_maintenance_configuration = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.TableMaintenanceType", Required: true},
	{Name: "Value", Flag: "value", Type: "*types.TableMaintenanceConfigurationValue", Required: true},
}

var fields_put_table_policy = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "ResourcePolicy", Flag: "resource-policy", Type: "*string", Required: true},
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: true},
}

var fields_put_table_record_expiration_configuration = []leanruntime.Field{
	{Name: "TableArn", Flag: "table-arn", Type: "*string", Required: true},
	{Name: "Value", Flag: "value", Type: "*types.TableRecordExpirationConfigurationValue", Required: true},
}

var fields_put_table_replication = []leanruntime.Field{
	{Name: "Configuration", Flag: "configuration", Type: "*types.TableReplicationConfiguration", Required: true},
	{Name: "TableArn", Flag: "table-arn", Type: "*string", Required: true},
	{Name: "VersionToken", Flag: "version-token", Type: "*string", Required: false},
}

var fields_rename_table = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "NewName", Flag: "new-name", Type: "*string", Required: false},
	{Name: "NewNamespaceName", Flag: "new-namespace-name", Type: "*string", Required: false},
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: true},
	{Name: "VersionToken", Flag: "version-token", Type: "*string", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_table_metadata_location = []leanruntime.Field{
	{Name: "MetadataLocation", Flag: "metadata-location", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "TableBucketARN", Flag: "table-bucket-arn", Type: "*string", Required: true},
	{Name: "VersionToken", Flag: "version-token", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-namespace": {
			Name:   "create-namespace",
			Fields: fields_create_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateNamespaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_namespace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateNamespace(ctx, input)
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
		"create-table-bucket": {
			Name:   "create-table-bucket",
			Fields: fields_create_table_bucket,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTableBucketInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_table_bucket, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTableBucket(ctx, input)
			},
		},
		"delete-namespace": {
			Name:   "delete-namespace",
			Fields: fields_delete_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteNamespaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_namespace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteNamespace(ctx, input)
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
		"delete-table-bucket": {
			Name:   "delete-table-bucket",
			Fields: fields_delete_table_bucket,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTableBucketInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_table_bucket, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTableBucket(ctx, input)
			},
		},
		"delete-table-bucket-encryption": {
			Name:   "delete-table-bucket-encryption",
			Fields: fields_delete_table_bucket_encryption,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTableBucketEncryptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_table_bucket_encryption, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTableBucketEncryption(ctx, input)
			},
		},
		"delete-table-bucket-metrics-configuration": {
			Name:   "delete-table-bucket-metrics-configuration",
			Fields: fields_delete_table_bucket_metrics_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTableBucketMetricsConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_table_bucket_metrics_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTableBucketMetricsConfiguration(ctx, input)
			},
		},
		"delete-table-bucket-policy": {
			Name:   "delete-table-bucket-policy",
			Fields: fields_delete_table_bucket_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTableBucketPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_table_bucket_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTableBucketPolicy(ctx, input)
			},
		},
		"delete-table-bucket-replication": {
			Name:   "delete-table-bucket-replication",
			Fields: fields_delete_table_bucket_replication,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTableBucketReplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_table_bucket_replication, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTableBucketReplication(ctx, input)
			},
		},
		"delete-table-policy": {
			Name:   "delete-table-policy",
			Fields: fields_delete_table_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTablePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_table_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTablePolicy(ctx, input)
			},
		},
		"delete-table-replication": {
			Name:   "delete-table-replication",
			Fields: fields_delete_table_replication,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTableReplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_table_replication, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTableReplication(ctx, input)
			},
		},
		"get-namespace": {
			Name:   "get-namespace",
			Fields: fields_get_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetNamespaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_namespace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetNamespace(ctx, input)
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
		"get-table-bucket": {
			Name:   "get-table-bucket",
			Fields: fields_get_table_bucket,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTableBucketInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_table_bucket, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTableBucket(ctx, input)
			},
		},
		"get-table-bucket-encryption": {
			Name:   "get-table-bucket-encryption",
			Fields: fields_get_table_bucket_encryption,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTableBucketEncryptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_table_bucket_encryption, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTableBucketEncryption(ctx, input)
			},
		},
		"get-table-bucket-maintenance-configuration": {
			Name:   "get-table-bucket-maintenance-configuration",
			Fields: fields_get_table_bucket_maintenance_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTableBucketMaintenanceConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_table_bucket_maintenance_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTableBucketMaintenanceConfiguration(ctx, input)
			},
		},
		"get-table-bucket-metrics-configuration": {
			Name:   "get-table-bucket-metrics-configuration",
			Fields: fields_get_table_bucket_metrics_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTableBucketMetricsConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_table_bucket_metrics_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTableBucketMetricsConfiguration(ctx, input)
			},
		},
		"get-table-bucket-policy": {
			Name:   "get-table-bucket-policy",
			Fields: fields_get_table_bucket_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTableBucketPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_table_bucket_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTableBucketPolicy(ctx, input)
			},
		},
		"get-table-bucket-replication": {
			Name:   "get-table-bucket-replication",
			Fields: fields_get_table_bucket_replication,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTableBucketReplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_table_bucket_replication, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTableBucketReplication(ctx, input)
			},
		},
		"get-table-bucket-storage-class": {
			Name:   "get-table-bucket-storage-class",
			Fields: fields_get_table_bucket_storage_class,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTableBucketStorageClassInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_table_bucket_storage_class, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTableBucketStorageClass(ctx, input)
			},
		},
		"get-table-encryption": {
			Name:   "get-table-encryption",
			Fields: fields_get_table_encryption,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTableEncryptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_table_encryption, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTableEncryption(ctx, input)
			},
		},
		"get-table-maintenance-configuration": {
			Name:   "get-table-maintenance-configuration",
			Fields: fields_get_table_maintenance_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTableMaintenanceConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_table_maintenance_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTableMaintenanceConfiguration(ctx, input)
			},
		},
		"get-table-maintenance-job-status": {
			Name:   "get-table-maintenance-job-status",
			Fields: fields_get_table_maintenance_job_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTableMaintenanceJobStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_table_maintenance_job_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTableMaintenanceJobStatus(ctx, input)
			},
		},
		"get-table-metadata-location": {
			Name:   "get-table-metadata-location",
			Fields: fields_get_table_metadata_location,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTableMetadataLocationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_table_metadata_location, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTableMetadataLocation(ctx, input)
			},
		},
		"get-table-policy": {
			Name:   "get-table-policy",
			Fields: fields_get_table_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTablePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_table_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTablePolicy(ctx, input)
			},
		},
		"get-table-record-expiration-configuration": {
			Name:   "get-table-record-expiration-configuration",
			Fields: fields_get_table_record_expiration_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTableRecordExpirationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_table_record_expiration_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTableRecordExpirationConfiguration(ctx, input)
			},
		},
		"get-table-record-expiration-job-status": {
			Name:   "get-table-record-expiration-job-status",
			Fields: fields_get_table_record_expiration_job_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTableRecordExpirationJobStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_table_record_expiration_job_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTableRecordExpirationJobStatus(ctx, input)
			},
		},
		"get-table-replication": {
			Name:   "get-table-replication",
			Fields: fields_get_table_replication,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTableReplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_table_replication, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTableReplication(ctx, input)
			},
		},
		"get-table-replication-status": {
			Name:   "get-table-replication-status",
			Fields: fields_get_table_replication_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTableReplicationStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_table_replication_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTableReplicationStatus(ctx, input)
			},
		},
		"get-table-storage-class": {
			Name:   "get-table-storage-class",
			Fields: fields_get_table_storage_class,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTableStorageClassInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_table_storage_class, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTableStorageClass(ctx, input)
			},
		},
		"list-namespaces": {
			Name:   "list-namespaces",
			Fields: fields_list_namespaces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListNamespacesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_namespaces, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListNamespaces(ctx, input)
				}
				var results []*svc.ListNamespacesOutput
				p := svc.NewListNamespacesPaginator(client, input)
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
		"list-table-buckets": {
			Name:   "list-table-buckets",
			Fields: fields_list_table_buckets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTableBucketsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_table_buckets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTableBuckets(ctx, input)
				}
				var results []*svc.ListTableBucketsOutput
				p := svc.NewListTableBucketsPaginator(client, input)
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
				if _, err := leanruntime.ApplyInput(input, fields_list_tags_for_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTagsForResource(ctx, input)
			},
		},
		"put-table-bucket-encryption": {
			Name:   "put-table-bucket-encryption",
			Fields: fields_put_table_bucket_encryption,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutTableBucketEncryptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_table_bucket_encryption, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutTableBucketEncryption(ctx, input)
			},
		},
		"put-table-bucket-maintenance-configuration": {
			Name:   "put-table-bucket-maintenance-configuration",
			Fields: fields_put_table_bucket_maintenance_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutTableBucketMaintenanceConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_table_bucket_maintenance_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutTableBucketMaintenanceConfiguration(ctx, input)
			},
		},
		"put-table-bucket-metrics-configuration": {
			Name:   "put-table-bucket-metrics-configuration",
			Fields: fields_put_table_bucket_metrics_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutTableBucketMetricsConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_table_bucket_metrics_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutTableBucketMetricsConfiguration(ctx, input)
			},
		},
		"put-table-bucket-policy": {
			Name:   "put-table-bucket-policy",
			Fields: fields_put_table_bucket_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutTableBucketPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_table_bucket_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutTableBucketPolicy(ctx, input)
			},
		},
		"put-table-bucket-replication": {
			Name:   "put-table-bucket-replication",
			Fields: fields_put_table_bucket_replication,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutTableBucketReplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_table_bucket_replication, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutTableBucketReplication(ctx, input)
			},
		},
		"put-table-bucket-storage-class": {
			Name:   "put-table-bucket-storage-class",
			Fields: fields_put_table_bucket_storage_class,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutTableBucketStorageClassInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_table_bucket_storage_class, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutTableBucketStorageClass(ctx, input)
			},
		},
		"put-table-maintenance-configuration": {
			Name:   "put-table-maintenance-configuration",
			Fields: fields_put_table_maintenance_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutTableMaintenanceConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_table_maintenance_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutTableMaintenanceConfiguration(ctx, input)
			},
		},
		"put-table-policy": {
			Name:   "put-table-policy",
			Fields: fields_put_table_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutTablePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_table_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutTablePolicy(ctx, input)
			},
		},
		"put-table-record-expiration-configuration": {
			Name:   "put-table-record-expiration-configuration",
			Fields: fields_put_table_record_expiration_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutTableRecordExpirationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_table_record_expiration_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutTableRecordExpirationConfiguration(ctx, input)
			},
		},
		"put-table-replication": {
			Name:   "put-table-replication",
			Fields: fields_put_table_replication,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutTableReplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_table_replication, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutTableReplication(ctx, input)
			},
		},
		"rename-table": {
			Name:   "rename-table",
			Fields: fields_rename_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RenameTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_rename_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RenameTable(ctx, input)
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
		"update-table-metadata-location": {
			Name:   "update-table-metadata-location",
			Fields: fields_update_table_metadata_location,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTableMetadataLocationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_table_metadata_location, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTableMetadataLocation(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("s3tables", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
