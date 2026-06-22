package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/lakeformation"
)

var fields_add_lf_tags_to_resource = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "LFTags", Flag: "lf-tags", Type: "[]types.LFTagPair", Required: true},
	{Name: "Resource", Flag: "resource", Type: "*types.Resource", Required: true},
}

var fields_assume_decorated_role_with_saml = []leanruntime.Field{
	{Name: "DurationSeconds", Flag: "duration-seconds", Type: "*int32", Required: false},
	{Name: "PrincipalArn", Flag: "principal-arn", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "SAMLAssertion", Flag: "saml-assertion", Type: "*string", Required: true},
}

var fields_batch_grant_permissions = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "Entries", Flag: "entries", Type: "[]types.BatchPermissionsRequestEntry", Required: true},
}

var fields_batch_revoke_permissions = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "Entries", Flag: "entries", Type: "[]types.BatchPermissionsRequestEntry", Required: true},
}

var fields_cancel_transaction = []leanruntime.Field{
	{Name: "TransactionId", Flag: "transaction-id", Type: "*string", Required: true},
}

var fields_commit_transaction = []leanruntime.Field{
	{Name: "TransactionId", Flag: "transaction-id", Type: "*string", Required: true},
}

var fields_create_data_cells_filter = []leanruntime.Field{
	{Name: "TableData", Flag: "table-data", Type: "*types.DataCellsFilter", Required: true},
}

var fields_create_lake_formation_identity_center_configuration = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "ExternalFiltering", Flag: "external-filtering", Type: "*types.ExternalFilteringConfiguration", Required: false},
	{Name: "InstanceArn", Flag: "instance-arn", Type: "*string", Required: false},
	{Name: "ServiceIntegrations", Flag: "service-integrations", Type: "[]types.ServiceIntegrationUnion", Required: false},
	{Name: "ShareRecipients", Flag: "share-recipients", Type: "[]types.DataLakePrincipal", Required: false},
}

var fields_create_lake_formation_opt_in = []leanruntime.Field{
	{Name: "Condition", Flag: "condition", Type: "*types.Condition", Required: false},
	{Name: "Principal", Flag: "principal", Type: "*types.DataLakePrincipal", Required: true},
	{Name: "Resource", Flag: "resource", Type: "*types.Resource", Required: true},
}

var fields_create_lf_tag = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "TagKey", Flag: "tag-key", Type: "*string", Required: true},
	{Name: "TagValues", Flag: "tag-values", Type: "[]string", Required: true},
}

var fields_create_lf_tag_expression = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Expression", Flag: "expression", Type: "[]types.LFTag", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_data_cells_filter = []leanruntime.Field{
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "TableCatalogId", Flag: "table-catalog-id", Type: "*string", Required: false},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: false},
}

var fields_delete_lake_formation_identity_center_configuration = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
}

var fields_delete_lake_formation_opt_in = []leanruntime.Field{
	{Name: "Condition", Flag: "condition", Type: "*types.Condition", Required: false},
	{Name: "Principal", Flag: "principal", Type: "*types.DataLakePrincipal", Required: true},
	{Name: "Resource", Flag: "resource", Type: "*types.Resource", Required: true},
}

var fields_delete_lf_tag = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "TagKey", Flag: "tag-key", Type: "*string", Required: true},
}

var fields_delete_lf_tag_expression = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_objects_on_cancel = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "Objects", Flag: "objects", Type: "[]types.VirtualObject", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
	{Name: "TransactionId", Flag: "transaction-id", Type: "*string", Required: true},
}

var fields_deregister_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_describe_lake_formation_identity_center_configuration = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
}

var fields_describe_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_describe_transaction = []leanruntime.Field{
	{Name: "TransactionId", Flag: "transaction-id", Type: "*string", Required: true},
}

var fields_extend_transaction = []leanruntime.Field{
	{Name: "TransactionId", Flag: "transaction-id", Type: "*string", Required: false},
}

var fields_get_data_cells_filter = []leanruntime.Field{
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "TableCatalogId", Flag: "table-catalog-id", Type: "*string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_get_data_lake_principal = []leanruntime.Field{}

var fields_get_data_lake_settings = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
}

var fields_get_effective_permissions_for_path = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_get_lf_tag = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "TagKey", Flag: "tag-key", Type: "*string", Required: true},
}

var fields_get_lf_tag_expression = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_query_state = []leanruntime.Field{
	{Name: "QueryId", Flag: "query-id", Type: "*string", Required: true},
}

var fields_get_query_statistics = []leanruntime.Field{
	{Name: "QueryId", Flag: "query-id", Type: "*string", Required: true},
}

var fields_get_resource_lf_tags = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "Resource", Flag: "resource", Type: "*types.Resource", Required: true},
	{Name: "ShowAssignedLFTags", Flag: "show-assigned-lf-tags", Type: "*bool", Required: false},
}

var fields_get_table_objects = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PartitionPredicate", Flag: "partition-predicate", Type: "*string", Required: false},
	{Name: "QueryAsOfTime", Flag: "query-as-of-time", Type: "*time.Time", Required: false},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
	{Name: "TransactionId", Flag: "transaction-id", Type: "*string", Required: false},
}

var fields_get_temporary_data_location_credentials = []leanruntime.Field{
	{Name: "AuditContext", Flag: "audit-context", Type: "*types.AuditContext", Required: false},
	{Name: "CredentialsScope", Flag: "credentials-scope", Type: "types.CredentialsScope", Required: false},
	{Name: "DataLocations", Flag: "data-locations", Type: "[]string", Required: false},
	{Name: "DurationSeconds", Flag: "duration-seconds", Type: "*int32", Required: false},
}

var fields_get_temporary_glue_partition_credentials = []leanruntime.Field{
	{Name: "AuditContext", Flag: "audit-context", Type: "*types.AuditContext", Required: false},
	{Name: "DurationSeconds", Flag: "duration-seconds", Type: "*int32", Required: false},
	{Name: "Partition", Flag: "partition", Type: "*types.PartitionValueList", Required: true},
	{Name: "Permissions", Flag: "permissions", Type: "[]types.Permission", Required: false},
	{Name: "SupportedPermissionTypes", Flag: "supported-permission-types", Type: "[]types.PermissionType", Required: false},
	{Name: "TableArn", Flag: "table-arn", Type: "*string", Required: true},
}

var fields_get_temporary_glue_table_credentials = []leanruntime.Field{
	{Name: "AuditContext", Flag: "audit-context", Type: "*types.AuditContext", Required: false},
	{Name: "DurationSeconds", Flag: "duration-seconds", Type: "*int32", Required: false},
	{Name: "Permissions", Flag: "permissions", Type: "[]types.Permission", Required: false},
	{Name: "QuerySessionContext", Flag: "query-session-context", Type: "*types.QuerySessionContext", Required: false},
	{Name: "S3Path", Flag: "s3-path", Type: "*string", Required: false},
	{Name: "SupportedPermissionTypes", Flag: "supported-permission-types", Type: "[]types.PermissionType", Required: false},
	{Name: "TableArn", Flag: "table-arn", Type: "*string", Required: true},
}

var fields_get_work_unit_results = []leanruntime.Field{
	{Name: "QueryId", Flag: "query-id", Type: "*string", Required: true},
	{Name: "WorkUnitId", Flag: "work-unit-id", Type: "int64", Required: true},
	{Name: "WorkUnitToken", Flag: "work-unit-token", Type: "*string", Required: true},
}

var fields_get_work_units = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
	{Name: "QueryId", Flag: "query-id", Type: "*string", Required: true},
}

var fields_grant_permissions = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "Condition", Flag: "condition", Type: "*types.Condition", Required: false},
	{Name: "Permissions", Flag: "permissions", Type: "[]types.Permission", Required: true},
	{Name: "PermissionsWithGrantOption", Flag: "permissions-with-grant-option", Type: "[]types.Permission", Required: false},
	{Name: "Principal", Flag: "principal", Type: "*types.DataLakePrincipal", Required: true},
	{Name: "Resource", Flag: "resource", Type: "*types.Resource", Required: true},
}

var fields_list_data_cells_filter = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Table", Flag: "table", Type: "*types.TableResource", Required: false},
}

var fields_list_lake_formation_opt_ins = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Principal", Flag: "principal", Type: "*types.DataLakePrincipal", Required: false},
	{Name: "Resource", Flag: "resource", Type: "*types.Resource", Required: false},
}

var fields_list_lf_tag_expressions = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_lf_tags = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceShareType", Flag: "resource-share-type", Type: "types.ResourceShareType", Required: false},
}

var fields_list_permissions = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "IncludeRelated", Flag: "include-related", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Principal", Flag: "principal", Type: "*types.DataLakePrincipal", Required: false},
	{Name: "Resource", Flag: "resource", Type: "*types.Resource", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.DataLakeResourceType", Required: false},
}

var fields_list_resources = []leanruntime.Field{
	{Name: "FilterConditionList", Flag: "filter-condition-list", Type: "[]types.FilterCondition", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_table_storage_optimizers = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StorageOptimizerType", Flag: "storage-optimizer-type", Type: "types.OptimizerType", Required: false},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

var fields_list_transactions = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StatusFilter", Flag: "status-filter", Type: "types.TransactionStatusFilter", Required: false},
}

var fields_put_data_lake_settings = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "DataLakeSettings", Flag: "data-lake-settings", Type: "*types.DataLakeSettings", Required: true},
}

var fields_register_resource = []leanruntime.Field{
	{Name: "ExpectedResourceOwnerAccount", Flag: "expected-resource-owner-account", Type: "*string", Required: false},
	{Name: "HybridAccessEnabled", Flag: "hybrid-access-enabled", Type: "*bool", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "UseServiceLinkedRole", Flag: "use-service-linked-role", Type: "*bool", Required: false},
	{Name: "WithFederation", Flag: "with-federation", Type: "*bool", Required: false},
	{Name: "WithPrivilegedAccess", Flag: "with-privileged-access", Type: "bool", Required: false},
}

var fields_remove_lf_tags_from_resource = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "LFTags", Flag: "lf-tags", Type: "[]types.LFTagPair", Required: true},
	{Name: "Resource", Flag: "resource", Type: "*types.Resource", Required: true},
}

var fields_revoke_permissions = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "Condition", Flag: "condition", Type: "*types.Condition", Required: false},
	{Name: "Permissions", Flag: "permissions", Type: "[]types.Permission", Required: true},
	{Name: "PermissionsWithGrantOption", Flag: "permissions-with-grant-option", Type: "[]types.Permission", Required: false},
	{Name: "Principal", Flag: "principal", Type: "*types.DataLakePrincipal", Required: true},
	{Name: "Resource", Flag: "resource", Type: "*types.Resource", Required: true},
}

var fields_search_databases_by_lf_tags = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "Expression", Flag: "expression", Type: "[]types.LFTag", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_search_tables_by_lf_tags = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "Expression", Flag: "expression", Type: "[]types.LFTag", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_start_query_planning = []leanruntime.Field{
	{Name: "QueryPlanningContext", Flag: "query-planning-context", Type: "*types.QueryPlanningContext", Required: true},
	{Name: "QueryString", Flag: "query-string", Type: "*string", Required: true},
}

var fields_start_transaction = []leanruntime.Field{
	{Name: "TransactionType", Flag: "transaction-type", Type: "types.TransactionType", Required: false},
}

var fields_update_data_cells_filter = []leanruntime.Field{
	{Name: "TableData", Flag: "table-data", Type: "*types.DataCellsFilter", Required: true},
}

var fields_update_lake_formation_identity_center_configuration = []leanruntime.Field{
	{Name: "ApplicationStatus", Flag: "application-status", Type: "types.ApplicationStatus", Required: false},
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "ExternalFiltering", Flag: "external-filtering", Type: "*types.ExternalFilteringConfiguration", Required: false},
	{Name: "ServiceIntegrations", Flag: "service-integrations", Type: "[]types.ServiceIntegrationUnion", Required: false},
	{Name: "ShareRecipients", Flag: "share-recipients", Type: "[]types.DataLakePrincipal", Required: false},
}

var fields_update_lf_tag = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "TagKey", Flag: "tag-key", Type: "*string", Required: true},
	{Name: "TagValuesToAdd", Flag: "tag-values-to-add", Type: "[]string", Required: false},
	{Name: "TagValuesToDelete", Flag: "tag-values-to-delete", Type: "[]string", Required: false},
}

var fields_update_lf_tag_expression = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Expression", Flag: "expression", Type: "[]types.LFTag", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_update_resource = []leanruntime.Field{
	{Name: "ExpectedResourceOwnerAccount", Flag: "expected-resource-owner-account", Type: "*string", Required: false},
	{Name: "HybridAccessEnabled", Flag: "hybrid-access-enabled", Type: "*bool", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "WithFederation", Flag: "with-federation", Type: "*bool", Required: false},
}

var fields_update_table_objects = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
	{Name: "TransactionId", Flag: "transaction-id", Type: "*string", Required: false},
	{Name: "WriteOperations", Flag: "write-operations", Type: "[]types.WriteOperation", Required: true},
}

var fields_update_table_storage_optimizer = []leanruntime.Field{
	{Name: "CatalogId", Flag: "catalog-id", Type: "*string", Required: false},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "StorageOptimizerConfig", Flag: "storage-optimizer-config", Type: "map[string]map[string]string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"add-lf-tags-to-resource": {
			Name:   "add-lf-tags-to-resource",
			Fields: fields_add_lf_tags_to_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddLFTagsToResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_lf_tags_to_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddLFTagsToResource(ctx, input)
			},
		},
		"assume-decorated-role-with-saml": {
			Name:   "assume-decorated-role-with-saml",
			Fields: fields_assume_decorated_role_with_saml,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssumeDecoratedRoleWithSAMLInput{}
				if _, err := leanruntime.ApplyInput(input, fields_assume_decorated_role_with_saml, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssumeDecoratedRoleWithSAML(ctx, input)
			},
		},
		"batch-grant-permissions": {
			Name:   "batch-grant-permissions",
			Fields: fields_batch_grant_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGrantPermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_grant_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGrantPermissions(ctx, input)
			},
		},
		"batch-revoke-permissions": {
			Name:   "batch-revoke-permissions",
			Fields: fields_batch_revoke_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchRevokePermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_revoke_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchRevokePermissions(ctx, input)
			},
		},
		"cancel-transaction": {
			Name:   "cancel-transaction",
			Fields: fields_cancel_transaction,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelTransactionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_transaction, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelTransaction(ctx, input)
			},
		},
		"commit-transaction": {
			Name:   "commit-transaction",
			Fields: fields_commit_transaction,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CommitTransactionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_commit_transaction, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CommitTransaction(ctx, input)
			},
		},
		"create-data-cells-filter": {
			Name:   "create-data-cells-filter",
			Fields: fields_create_data_cells_filter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataCellsFilterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_cells_filter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataCellsFilter(ctx, input)
			},
		},
		"create-lake-formation-identity-center-configuration": {
			Name:   "create-lake-formation-identity-center-configuration",
			Fields: fields_create_lake_formation_identity_center_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLakeFormationIdentityCenterConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_lake_formation_identity_center_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLakeFormationIdentityCenterConfiguration(ctx, input)
			},
		},
		"create-lake-formation-opt-in": {
			Name:   "create-lake-formation-opt-in",
			Fields: fields_create_lake_formation_opt_in,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLakeFormationOptInInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_lake_formation_opt_in, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLakeFormationOptIn(ctx, input)
			},
		},
		"create-lf-tag": {
			Name:   "create-lf-tag",
			Fields: fields_create_lf_tag,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLFTagInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_lf_tag, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLFTag(ctx, input)
			},
		},
		"create-lf-tag-expression": {
			Name:   "create-lf-tag-expression",
			Fields: fields_create_lf_tag_expression,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLFTagExpressionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_lf_tag_expression, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLFTagExpression(ctx, input)
			},
		},
		"delete-data-cells-filter": {
			Name:   "delete-data-cells-filter",
			Fields: fields_delete_data_cells_filter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDataCellsFilterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_data_cells_filter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataCellsFilter(ctx, input)
			},
		},
		"delete-lake-formation-identity-center-configuration": {
			Name:   "delete-lake-formation-identity-center-configuration",
			Fields: fields_delete_lake_formation_identity_center_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLakeFormationIdentityCenterConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_lake_formation_identity_center_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLakeFormationIdentityCenterConfiguration(ctx, input)
			},
		},
		"delete-lake-formation-opt-in": {
			Name:   "delete-lake-formation-opt-in",
			Fields: fields_delete_lake_formation_opt_in,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLakeFormationOptInInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_lake_formation_opt_in, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLakeFormationOptIn(ctx, input)
			},
		},
		"delete-lf-tag": {
			Name:   "delete-lf-tag",
			Fields: fields_delete_lf_tag,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLFTagInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_lf_tag, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLFTag(ctx, input)
			},
		},
		"delete-lf-tag-expression": {
			Name:   "delete-lf-tag-expression",
			Fields: fields_delete_lf_tag_expression,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLFTagExpressionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_lf_tag_expression, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLFTagExpression(ctx, input)
			},
		},
		"delete-objects-on-cancel": {
			Name:   "delete-objects-on-cancel",
			Fields: fields_delete_objects_on_cancel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteObjectsOnCancelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_objects_on_cancel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteObjectsOnCancel(ctx, input)
			},
		},
		"deregister-resource": {
			Name:   "deregister-resource",
			Fields: fields_deregister_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterResource(ctx, input)
			},
		},
		"describe-lake-formation-identity-center-configuration": {
			Name:   "describe-lake-formation-identity-center-configuration",
			Fields: fields_describe_lake_formation_identity_center_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLakeFormationIdentityCenterConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_lake_formation_identity_center_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLakeFormationIdentityCenterConfiguration(ctx, input)
			},
		},
		"describe-resource": {
			Name:   "describe-resource",
			Fields: fields_describe_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeResource(ctx, input)
			},
		},
		"describe-transaction": {
			Name:   "describe-transaction",
			Fields: fields_describe_transaction,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTransactionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_transaction, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTransaction(ctx, input)
			},
		},
		"extend-transaction": {
			Name:   "extend-transaction",
			Fields: fields_extend_transaction,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExtendTransactionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_extend_transaction, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExtendTransaction(ctx, input)
			},
		},
		"get-data-cells-filter": {
			Name:   "get-data-cells-filter",
			Fields: fields_get_data_cells_filter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataCellsFilterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_cells_filter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataCellsFilter(ctx, input)
			},
		},
		"get-data-lake-principal": {
			Name:   "get-data-lake-principal",
			Fields: fields_get_data_lake_principal,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataLakePrincipalInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_lake_principal, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataLakePrincipal(ctx, input)
			},
		},
		"get-data-lake-settings": {
			Name:   "get-data-lake-settings",
			Fields: fields_get_data_lake_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataLakeSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_lake_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataLakeSettings(ctx, input)
			},
		},
		"get-effective-permissions-for-path": {
			Name:   "get-effective-permissions-for-path",
			Fields: fields_get_effective_permissions_for_path,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEffectivePermissionsForPathInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_effective_permissions_for_path, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetEffectivePermissionsForPath(ctx, input)
				}
				var results []*svc.GetEffectivePermissionsForPathOutput
				p := svc.NewGetEffectivePermissionsForPathPaginator(client, input)
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
		"get-lf-tag": {
			Name:   "get-lf-tag",
			Fields: fields_get_lf_tag,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLFTagInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_lf_tag, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLFTag(ctx, input)
			},
		},
		"get-lf-tag-expression": {
			Name:   "get-lf-tag-expression",
			Fields: fields_get_lf_tag_expression,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLFTagExpressionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_lf_tag_expression, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLFTagExpression(ctx, input)
			},
		},
		"get-query-state": {
			Name:   "get-query-state",
			Fields: fields_get_query_state,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetQueryStateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_query_state, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetQueryState(ctx, input)
			},
		},
		"get-query-statistics": {
			Name:   "get-query-statistics",
			Fields: fields_get_query_statistics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetQueryStatisticsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_query_statistics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetQueryStatistics(ctx, input)
			},
		},
		"get-resource-lf-tags": {
			Name:   "get-resource-lf-tags",
			Fields: fields_get_resource_lf_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourceLFTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resource_lf_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResourceLFTags(ctx, input)
			},
		},
		"get-table-objects": {
			Name:   "get-table-objects",
			Fields: fields_get_table_objects,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTableObjectsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_table_objects, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetTableObjects(ctx, input)
				}
				var results []*svc.GetTableObjectsOutput
				p := svc.NewGetTableObjectsPaginator(client, input)
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
		"get-temporary-data-location-credentials": {
			Name:   "get-temporary-data-location-credentials",
			Fields: fields_get_temporary_data_location_credentials,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTemporaryDataLocationCredentialsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_temporary_data_location_credentials, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTemporaryDataLocationCredentials(ctx, input)
			},
		},
		"get-temporary-glue-partition-credentials": {
			Name:   "get-temporary-glue-partition-credentials",
			Fields: fields_get_temporary_glue_partition_credentials,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTemporaryGluePartitionCredentialsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_temporary_glue_partition_credentials, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTemporaryGluePartitionCredentials(ctx, input)
			},
		},
		"get-temporary-glue-table-credentials": {
			Name:   "get-temporary-glue-table-credentials",
			Fields: fields_get_temporary_glue_table_credentials,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTemporaryGlueTableCredentialsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_temporary_glue_table_credentials, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTemporaryGlueTableCredentials(ctx, input)
			},
		},
		"get-work-unit-results": {
			Name:   "get-work-unit-results",
			Fields: fields_get_work_unit_results,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWorkUnitResultsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_work_unit_results, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWorkUnitResults(ctx, input)
			},
		},
		"get-work-units": {
			Name:   "get-work-units",
			Fields: fields_get_work_units,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWorkUnitsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_work_units, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetWorkUnits(ctx, input)
				}
				var results []*svc.GetWorkUnitsOutput
				p := svc.NewGetWorkUnitsPaginator(client, input)
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
		"grant-permissions": {
			Name:   "grant-permissions",
			Fields: fields_grant_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GrantPermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_grant_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GrantPermissions(ctx, input)
			},
		},
		"list-data-cells-filter": {
			Name:   "list-data-cells-filter",
			Fields: fields_list_data_cells_filter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataCellsFilterInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_cells_filter, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataCellsFilter(ctx, input)
				}
				var results []*svc.ListDataCellsFilterOutput
				p := svc.NewListDataCellsFilterPaginator(client, input)
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
		"list-lake-formation-opt-ins": {
			Name:   "list-lake-formation-opt-ins",
			Fields: fields_list_lake_formation_opt_ins,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLakeFormationOptInsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_lake_formation_opt_ins, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLakeFormationOptIns(ctx, input)
				}
				var results []*svc.ListLakeFormationOptInsOutput
				p := svc.NewListLakeFormationOptInsPaginator(client, input)
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
		"list-lf-tag-expressions": {
			Name:   "list-lf-tag-expressions",
			Fields: fields_list_lf_tag_expressions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLFTagExpressionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_lf_tag_expressions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLFTagExpressions(ctx, input)
				}
				var results []*svc.ListLFTagExpressionsOutput
				p := svc.NewListLFTagExpressionsPaginator(client, input)
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
		"list-lf-tags": {
			Name:   "list-lf-tags",
			Fields: fields_list_lf_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLFTagsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_lf_tags, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLFTags(ctx, input)
				}
				var results []*svc.ListLFTagsOutput
				p := svc.NewListLFTagsPaginator(client, input)
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
		"list-permissions": {
			Name:   "list-permissions",
			Fields: fields_list_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPermissionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_permissions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPermissions(ctx, input)
				}
				var results []*svc.ListPermissionsOutput
				p := svc.NewListPermissionsPaginator(client, input)
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
		"list-resources": {
			Name:   "list-resources",
			Fields: fields_list_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResources(ctx, input)
				}
				var results []*svc.ListResourcesOutput
				p := svc.NewListResourcesPaginator(client, input)
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
		"list-table-storage-optimizers": {
			Name:   "list-table-storage-optimizers",
			Fields: fields_list_table_storage_optimizers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTableStorageOptimizersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_table_storage_optimizers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTableStorageOptimizers(ctx, input)
				}
				var results []*svc.ListTableStorageOptimizersOutput
				p := svc.NewListTableStorageOptimizersPaginator(client, input)
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
		"list-transactions": {
			Name:   "list-transactions",
			Fields: fields_list_transactions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTransactionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_transactions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTransactions(ctx, input)
				}
				var results []*svc.ListTransactionsOutput
				p := svc.NewListTransactionsPaginator(client, input)
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
		"put-data-lake-settings": {
			Name:   "put-data-lake-settings",
			Fields: fields_put_data_lake_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutDataLakeSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_data_lake_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutDataLakeSettings(ctx, input)
			},
		},
		"register-resource": {
			Name:   "register-resource",
			Fields: fields_register_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterResource(ctx, input)
			},
		},
		"remove-lf-tags-from-resource": {
			Name:   "remove-lf-tags-from-resource",
			Fields: fields_remove_lf_tags_from_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveLFTagsFromResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_lf_tags_from_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveLFTagsFromResource(ctx, input)
			},
		},
		"revoke-permissions": {
			Name:   "revoke-permissions",
			Fields: fields_revoke_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RevokePermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_revoke_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RevokePermissions(ctx, input)
			},
		},
		"search-databases-by-lf-tags": {
			Name:   "search-databases-by-lf-tags",
			Fields: fields_search_databases_by_lf_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchDatabasesByLFTagsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_databases_by_lf_tags, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchDatabasesByLFTags(ctx, input)
				}
				var results []*svc.SearchDatabasesByLFTagsOutput
				p := svc.NewSearchDatabasesByLFTagsPaginator(client, input)
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
		"search-tables-by-lf-tags": {
			Name:   "search-tables-by-lf-tags",
			Fields: fields_search_tables_by_lf_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchTablesByLFTagsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_tables_by_lf_tags, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchTablesByLFTags(ctx, input)
				}
				var results []*svc.SearchTablesByLFTagsOutput
				p := svc.NewSearchTablesByLFTagsPaginator(client, input)
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
		"start-query-planning": {
			Name:   "start-query-planning",
			Fields: fields_start_query_planning,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartQueryPlanningInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_query_planning, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartQueryPlanning(ctx, input)
			},
		},
		"start-transaction": {
			Name:   "start-transaction",
			Fields: fields_start_transaction,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartTransactionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_transaction, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartTransaction(ctx, input)
			},
		},
		"update-data-cells-filter": {
			Name:   "update-data-cells-filter",
			Fields: fields_update_data_cells_filter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDataCellsFilterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_data_cells_filter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDataCellsFilter(ctx, input)
			},
		},
		"update-lake-formation-identity-center-configuration": {
			Name:   "update-lake-formation-identity-center-configuration",
			Fields: fields_update_lake_formation_identity_center_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLakeFormationIdentityCenterConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_lake_formation_identity_center_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLakeFormationIdentityCenterConfiguration(ctx, input)
			},
		},
		"update-lf-tag": {
			Name:   "update-lf-tag",
			Fields: fields_update_lf_tag,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLFTagInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_lf_tag, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLFTag(ctx, input)
			},
		},
		"update-lf-tag-expression": {
			Name:   "update-lf-tag-expression",
			Fields: fields_update_lf_tag_expression,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLFTagExpressionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_lf_tag_expression, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLFTagExpression(ctx, input)
			},
		},
		"update-resource": {
			Name:   "update-resource",
			Fields: fields_update_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateResource(ctx, input)
			},
		},
		"update-table-objects": {
			Name:   "update-table-objects",
			Fields: fields_update_table_objects,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTableObjectsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_table_objects, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTableObjects(ctx, input)
			},
		},
		"update-table-storage-optimizer": {
			Name:   "update-table-storage-optimizer",
			Fields: fields_update_table_storage_optimizer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTableStorageOptimizerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_table_storage_optimizer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTableStorageOptimizer(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("lakeformation", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
