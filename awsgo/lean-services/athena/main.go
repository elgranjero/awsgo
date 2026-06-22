package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/athena"
)

var fields_batch_get_named_query = []leanruntime.Field{
	{Name: "NamedQueryIds", Flag: "named-query-ids", Type: "[]string", Required: true},
}

var fields_batch_get_prepared_statement = []leanruntime.Field{
	{Name: "PreparedStatementNames", Flag: "prepared-statement-names", Type: "[]string", Required: true},
	{Name: "WorkGroup", Flag: "work-group", Type: "*string", Required: true},
}

var fields_batch_get_query_execution = []leanruntime.Field{
	{Name: "QueryExecutionIds", Flag: "query-execution-ids", Type: "[]string", Required: true},
}

var fields_cancel_capacity_reservation = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_capacity_reservation = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TargetDpus", Flag: "target-dpus", Type: "*int32", Required: true},
}

var fields_create_data_catalog = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "map[string]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Type", Flag: "type", Type: "types.DataCatalogType", Required: true},
}

var fields_create_named_query = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Database", Flag: "database", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "QueryString", Flag: "query-string", Type: "*string", Required: true},
	{Name: "WorkGroup", Flag: "work-group", Type: "*string", Required: false},
}

var fields_create_notebook = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "WorkGroup", Flag: "work-group", Type: "*string", Required: true},
}

var fields_create_prepared_statement = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "QueryStatement", Flag: "query-statement", Type: "*string", Required: true},
	{Name: "StatementName", Flag: "statement-name", Type: "*string", Required: true},
	{Name: "WorkGroup", Flag: "work-group", Type: "*string", Required: true},
}

var fields_create_presigned_notebook_url = []leanruntime.Field{
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_create_work_group = []leanruntime.Field{
	{Name: "Configuration", Flag: "configuration", Type: "*types.WorkGroupConfiguration", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_capacity_reservation = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_data_catalog = []leanruntime.Field{
	{Name: "DeleteCatalogOnly", Flag: "delete-catalog-only", Type: "bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_named_query = []leanruntime.Field{
	{Name: "NamedQueryId", Flag: "named-query-id", Type: "*string", Required: true},
}

var fields_delete_notebook = []leanruntime.Field{
	{Name: "NotebookId", Flag: "notebook-id", Type: "*string", Required: true},
}

var fields_delete_prepared_statement = []leanruntime.Field{
	{Name: "StatementName", Flag: "statement-name", Type: "*string", Required: true},
	{Name: "WorkGroup", Flag: "work-group", Type: "*string", Required: true},
}

var fields_delete_work_group = []leanruntime.Field{
	{Name: "RecursiveDeleteOption", Flag: "recursive-delete-option", Type: "*bool", Required: false},
	{Name: "WorkGroup", Flag: "work-group", Type: "*string", Required: true},
}

var fields_export_notebook = []leanruntime.Field{
	{Name: "NotebookId", Flag: "notebook-id", Type: "*string", Required: true},
}

var fields_get_calculation_execution = []leanruntime.Field{
	{Name: "CalculationExecutionId", Flag: "calculation-execution-id", Type: "*string", Required: true},
}

var fields_get_calculation_execution_code = []leanruntime.Field{
	{Name: "CalculationExecutionId", Flag: "calculation-execution-id", Type: "*string", Required: true},
}

var fields_get_calculation_execution_status = []leanruntime.Field{
	{Name: "CalculationExecutionId", Flag: "calculation-execution-id", Type: "*string", Required: true},
}

var fields_get_capacity_assignment_configuration = []leanruntime.Field{
	{Name: "CapacityReservationName", Flag: "capacity-reservation-name", Type: "*string", Required: true},
}

var fields_get_capacity_reservation = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_data_catalog = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "WorkGroup", Flag: "work-group", Type: "*string", Required: false},
}

var fields_get_database = []leanruntime.Field{
	{Name: "CatalogName", Flag: "catalog-name", Type: "*string", Required: true},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "WorkGroup", Flag: "work-group", Type: "*string", Required: false},
}

var fields_get_named_query = []leanruntime.Field{
	{Name: "NamedQueryId", Flag: "named-query-id", Type: "*string", Required: true},
}

var fields_get_notebook_metadata = []leanruntime.Field{
	{Name: "NotebookId", Flag: "notebook-id", Type: "*string", Required: true},
}

var fields_get_prepared_statement = []leanruntime.Field{
	{Name: "StatementName", Flag: "statement-name", Type: "*string", Required: true},
	{Name: "WorkGroup", Flag: "work-group", Type: "*string", Required: true},
}

var fields_get_query_execution = []leanruntime.Field{
	{Name: "QueryExecutionId", Flag: "query-execution-id", Type: "*string", Required: true},
}

var fields_get_query_results = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueryExecutionId", Flag: "query-execution-id", Type: "*string", Required: true},
	{Name: "QueryResultType", Flag: "query-result-type", Type: "types.QueryResultType", Required: false},
}

var fields_get_query_runtime_statistics = []leanruntime.Field{
	{Name: "QueryExecutionId", Flag: "query-execution-id", Type: "*string", Required: true},
}

var fields_get_resource_dashboard = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_get_session = []leanruntime.Field{
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_get_session_endpoint = []leanruntime.Field{
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_get_session_status = []leanruntime.Field{
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_get_table_metadata = []leanruntime.Field{
	{Name: "CatalogName", Flag: "catalog-name", Type: "*string", Required: true},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "TableName", Flag: "table-name", Type: "*string", Required: true},
	{Name: "WorkGroup", Flag: "work-group", Type: "*string", Required: false},
}

var fields_get_work_group = []leanruntime.Field{
	{Name: "WorkGroup", Flag: "work-group", Type: "*string", Required: true},
}

var fields_import_notebook = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NotebookS3LocationUri", Flag: "notebook-s3-location-uri", Type: "*string", Required: false},
	{Name: "Payload", Flag: "payload", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.NotebookType", Required: true},
	{Name: "WorkGroup", Flag: "work-group", Type: "*string", Required: true},
}

var fields_list_application_dpu_sizes = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_calculation_executions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
	{Name: "StateFilter", Flag: "state-filter", Type: "types.CalculationExecutionState", Required: false},
}

var fields_list_capacity_reservations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_data_catalogs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkGroup", Flag: "work-group", Type: "*string", Required: false},
}

var fields_list_databases = []leanruntime.Field{
	{Name: "CatalogName", Flag: "catalog-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkGroup", Flag: "work-group", Type: "*string", Required: false},
}

var fields_list_engine_versions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_executors = []leanruntime.Field{
	{Name: "ExecutorStateFilter", Flag: "executor-state-filter", Type: "types.ExecutorState", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_list_named_queries = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkGroup", Flag: "work-group", Type: "*string", Required: false},
}

var fields_list_notebook_metadata = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "*types.FilterDefinition", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkGroup", Flag: "work-group", Type: "*string", Required: true},
}

var fields_list_notebook_sessions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "NotebookId", Flag: "notebook-id", Type: "*string", Required: true},
}

var fields_list_prepared_statements = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkGroup", Flag: "work-group", Type: "*string", Required: true},
}

var fields_list_query_executions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkGroup", Flag: "work-group", Type: "*string", Required: false},
}

var fields_list_sessions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StateFilter", Flag: "state-filter", Type: "types.SessionState", Required: false},
	{Name: "WorkGroup", Flag: "work-group", Type: "*string", Required: true},
}

var fields_list_table_metadata = []leanruntime.Field{
	{Name: "CatalogName", Flag: "catalog-name", Type: "*string", Required: true},
	{Name: "DatabaseName", Flag: "database-name", Type: "*string", Required: true},
	{Name: "Expression", Flag: "expression", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkGroup", Flag: "work-group", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_work_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_put_capacity_assignment_configuration = []leanruntime.Field{
	{Name: "CapacityAssignments", Flag: "capacity-assignments", Type: "[]types.CapacityAssignment", Required: true},
	{Name: "CapacityReservationName", Flag: "capacity-reservation-name", Type: "*string", Required: true},
}

var fields_start_calculation_execution = []leanruntime.Field{
	{Name: "CalculationConfiguration", Flag: "calculation-configuration", Type: "*types.CalculationConfiguration", Required: false},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "CodeBlock", Flag: "code-block", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_start_query_execution = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "EngineConfiguration", Flag: "engine-configuration", Type: "*types.EngineConfiguration", Required: false},
	{Name: "ExecutionParameters", Flag: "execution-parameters", Type: "[]string", Required: false},
	{Name: "QueryExecutionContext", Flag: "query-execution-context", Type: "*types.QueryExecutionContext", Required: false},
	{Name: "QueryString", Flag: "query-string", Type: "*string", Required: true},
	{Name: "ResultConfiguration", Flag: "result-configuration", Type: "*types.ResultConfiguration", Required: false},
	{Name: "ResultReuseConfiguration", Flag: "result-reuse-configuration", Type: "*types.ResultReuseConfiguration", Required: false},
	{Name: "WorkGroup", Flag: "work-group", Type: "*string", Required: false},
}

var fields_start_session = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "CopyWorkGroupTags", Flag: "copy-work-group-tags", Type: "*bool", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EngineConfiguration", Flag: "engine-configuration", Type: "*types.EngineConfiguration", Required: true},
	{Name: "ExecutionRole", Flag: "execution-role", Type: "*string", Required: false},
	{Name: "MonitoringConfiguration", Flag: "monitoring-configuration", Type: "*types.MonitoringConfiguration", Required: false},
	{Name: "NotebookVersion", Flag: "notebook-version", Type: "*string", Required: false},
	{Name: "SessionIdleTimeoutInMinutes", Flag: "session-idle-timeout-in-minutes", Type: "*int32", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "WorkGroup", Flag: "work-group", Type: "*string", Required: true},
}

var fields_stop_calculation_execution = []leanruntime.Field{
	{Name: "CalculationExecutionId", Flag: "calculation-execution-id", Type: "*string", Required: true},
}

var fields_stop_query_execution = []leanruntime.Field{
	{Name: "QueryExecutionId", Flag: "query-execution-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_terminate_session = []leanruntime.Field{
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_capacity_reservation = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "TargetDpus", Flag: "target-dpus", Type: "*int32", Required: true},
}

var fields_update_data_catalog = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "map[string]string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.DataCatalogType", Required: true},
}

var fields_update_named_query = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NamedQueryId", Flag: "named-query-id", Type: "*string", Required: true},
	{Name: "QueryString", Flag: "query-string", Type: "*string", Required: true},
}

var fields_update_notebook = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "NotebookId", Flag: "notebook-id", Type: "*string", Required: true},
	{Name: "Payload", Flag: "payload", Type: "*string", Required: true},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.NotebookType", Required: true},
}

var fields_update_notebook_metadata = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NotebookId", Flag: "notebook-id", Type: "*string", Required: true},
}

var fields_update_prepared_statement = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "QueryStatement", Flag: "query-statement", Type: "*string", Required: true},
	{Name: "StatementName", Flag: "statement-name", Type: "*string", Required: true},
	{Name: "WorkGroup", Flag: "work-group", Type: "*string", Required: true},
}

var fields_update_work_group = []leanruntime.Field{
	{Name: "ConfigurationUpdates", Flag: "configuration-updates", Type: "*types.WorkGroupConfigurationUpdates", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "State", Flag: "state", Type: "types.WorkGroupState", Required: false},
	{Name: "WorkGroup", Flag: "work-group", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-get-named-query": {
			Name:   "batch-get-named-query",
			Fields: fields_batch_get_named_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetNamedQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_named_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetNamedQuery(ctx, input)
			},
		},
		"batch-get-prepared-statement": {
			Name:   "batch-get-prepared-statement",
			Fields: fields_batch_get_prepared_statement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetPreparedStatementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_prepared_statement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetPreparedStatement(ctx, input)
			},
		},
		"batch-get-query-execution": {
			Name:   "batch-get-query-execution",
			Fields: fields_batch_get_query_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetQueryExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_query_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetQueryExecution(ctx, input)
			},
		},
		"cancel-capacity-reservation": {
			Name:   "cancel-capacity-reservation",
			Fields: fields_cancel_capacity_reservation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelCapacityReservationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_capacity_reservation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelCapacityReservation(ctx, input)
			},
		},
		"create-capacity-reservation": {
			Name:   "create-capacity-reservation",
			Fields: fields_create_capacity_reservation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCapacityReservationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_capacity_reservation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCapacityReservation(ctx, input)
			},
		},
		"create-data-catalog": {
			Name:   "create-data-catalog",
			Fields: fields_create_data_catalog,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataCatalogInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_catalog, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataCatalog(ctx, input)
			},
		},
		"create-named-query": {
			Name:   "create-named-query",
			Fields: fields_create_named_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateNamedQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_named_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateNamedQuery(ctx, input)
			},
		},
		"create-notebook": {
			Name:   "create-notebook",
			Fields: fields_create_notebook,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateNotebookInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_notebook, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateNotebook(ctx, input)
			},
		},
		"create-prepared-statement": {
			Name:   "create-prepared-statement",
			Fields: fields_create_prepared_statement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePreparedStatementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_prepared_statement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePreparedStatement(ctx, input)
			},
		},
		"create-presigned-notebook-url": {
			Name:   "create-presigned-notebook-url",
			Fields: fields_create_presigned_notebook_url,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePresignedNotebookUrlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_presigned_notebook_url, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePresignedNotebookUrl(ctx, input)
			},
		},
		"create-work-group": {
			Name:   "create-work-group",
			Fields: fields_create_work_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWorkGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_work_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWorkGroup(ctx, input)
			},
		},
		"delete-capacity-reservation": {
			Name:   "delete-capacity-reservation",
			Fields: fields_delete_capacity_reservation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCapacityReservationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_capacity_reservation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCapacityReservation(ctx, input)
			},
		},
		"delete-data-catalog": {
			Name:   "delete-data-catalog",
			Fields: fields_delete_data_catalog,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDataCatalogInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_data_catalog, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataCatalog(ctx, input)
			},
		},
		"delete-named-query": {
			Name:   "delete-named-query",
			Fields: fields_delete_named_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteNamedQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_named_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteNamedQuery(ctx, input)
			},
		},
		"delete-notebook": {
			Name:   "delete-notebook",
			Fields: fields_delete_notebook,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteNotebookInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_notebook, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteNotebook(ctx, input)
			},
		},
		"delete-prepared-statement": {
			Name:   "delete-prepared-statement",
			Fields: fields_delete_prepared_statement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePreparedStatementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_prepared_statement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePreparedStatement(ctx, input)
			},
		},
		"delete-work-group": {
			Name:   "delete-work-group",
			Fields: fields_delete_work_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWorkGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_work_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWorkGroup(ctx, input)
			},
		},
		"export-notebook": {
			Name:   "export-notebook",
			Fields: fields_export_notebook,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExportNotebookInput{}
				if _, err := leanruntime.ApplyInput(input, fields_export_notebook, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExportNotebook(ctx, input)
			},
		},
		"get-calculation-execution": {
			Name:   "get-calculation-execution",
			Fields: fields_get_calculation_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCalculationExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_calculation_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCalculationExecution(ctx, input)
			},
		},
		"get-calculation-execution-code": {
			Name:   "get-calculation-execution-code",
			Fields: fields_get_calculation_execution_code,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCalculationExecutionCodeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_calculation_execution_code, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCalculationExecutionCode(ctx, input)
			},
		},
		"get-calculation-execution-status": {
			Name:   "get-calculation-execution-status",
			Fields: fields_get_calculation_execution_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCalculationExecutionStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_calculation_execution_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCalculationExecutionStatus(ctx, input)
			},
		},
		"get-capacity-assignment-configuration": {
			Name:   "get-capacity-assignment-configuration",
			Fields: fields_get_capacity_assignment_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCapacityAssignmentConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_capacity_assignment_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCapacityAssignmentConfiguration(ctx, input)
			},
		},
		"get-capacity-reservation": {
			Name:   "get-capacity-reservation",
			Fields: fields_get_capacity_reservation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCapacityReservationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_capacity_reservation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCapacityReservation(ctx, input)
			},
		},
		"get-data-catalog": {
			Name:   "get-data-catalog",
			Fields: fields_get_data_catalog,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataCatalogInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_catalog, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataCatalog(ctx, input)
			},
		},
		"get-database": {
			Name:   "get-database",
			Fields: fields_get_database,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDatabaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_database, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDatabase(ctx, input)
			},
		},
		"get-named-query": {
			Name:   "get-named-query",
			Fields: fields_get_named_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetNamedQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_named_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetNamedQuery(ctx, input)
			},
		},
		"get-notebook-metadata": {
			Name:   "get-notebook-metadata",
			Fields: fields_get_notebook_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetNotebookMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_notebook_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetNotebookMetadata(ctx, input)
			},
		},
		"get-prepared-statement": {
			Name:   "get-prepared-statement",
			Fields: fields_get_prepared_statement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPreparedStatementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_prepared_statement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPreparedStatement(ctx, input)
			},
		},
		"get-query-execution": {
			Name:   "get-query-execution",
			Fields: fields_get_query_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetQueryExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_query_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetQueryExecution(ctx, input)
			},
		},
		"get-query-results": {
			Name:   "get-query-results",
			Fields: fields_get_query_results,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetQueryResultsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_query_results, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetQueryResults(ctx, input)
				}
				var results []*svc.GetQueryResultsOutput
				p := svc.NewGetQueryResultsPaginator(client, input)
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
		"get-query-runtime-statistics": {
			Name:   "get-query-runtime-statistics",
			Fields: fields_get_query_runtime_statistics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetQueryRuntimeStatisticsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_query_runtime_statistics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetQueryRuntimeStatistics(ctx, input)
			},
		},
		"get-resource-dashboard": {
			Name:   "get-resource-dashboard",
			Fields: fields_get_resource_dashboard,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourceDashboardInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resource_dashboard, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResourceDashboard(ctx, input)
			},
		},
		"get-session": {
			Name:   "get-session",
			Fields: fields_get_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSession(ctx, input)
			},
		},
		"get-session-endpoint": {
			Name:   "get-session-endpoint",
			Fields: fields_get_session_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSessionEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_session_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSessionEndpoint(ctx, input)
			},
		},
		"get-session-status": {
			Name:   "get-session-status",
			Fields: fields_get_session_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSessionStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_session_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSessionStatus(ctx, input)
			},
		},
		"get-table-metadata": {
			Name:   "get-table-metadata",
			Fields: fields_get_table_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTableMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_table_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTableMetadata(ctx, input)
			},
		},
		"get-work-group": {
			Name:   "get-work-group",
			Fields: fields_get_work_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWorkGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_work_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWorkGroup(ctx, input)
			},
		},
		"import-notebook": {
			Name:   "import-notebook",
			Fields: fields_import_notebook,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportNotebookInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_notebook, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportNotebook(ctx, input)
			},
		},
		"list-application-dpu-sizes": {
			Name:   "list-application-dpu-sizes",
			Fields: fields_list_application_dpu_sizes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListApplicationDPUSizesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_application_dpu_sizes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApplicationDPUSizes(ctx, input)
				}
				var results []*svc.ListApplicationDPUSizesOutput
				p := svc.NewListApplicationDPUSizesPaginator(client, input)
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
		"list-calculation-executions": {
			Name:   "list-calculation-executions",
			Fields: fields_list_calculation_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCalculationExecutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_calculation_executions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCalculationExecutions(ctx, input)
				}
				var results []*svc.ListCalculationExecutionsOutput
				p := svc.NewListCalculationExecutionsPaginator(client, input)
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
		"list-capacity-reservations": {
			Name:   "list-capacity-reservations",
			Fields: fields_list_capacity_reservations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCapacityReservationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_capacity_reservations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCapacityReservations(ctx, input)
				}
				var results []*svc.ListCapacityReservationsOutput
				p := svc.NewListCapacityReservationsPaginator(client, input)
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
		"list-data-catalogs": {
			Name:   "list-data-catalogs",
			Fields: fields_list_data_catalogs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataCatalogsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_catalogs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataCatalogs(ctx, input)
				}
				var results []*svc.ListDataCatalogsOutput
				p := svc.NewListDataCatalogsPaginator(client, input)
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
		"list-databases": {
			Name:   "list-databases",
			Fields: fields_list_databases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDatabasesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_databases, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDatabases(ctx, input)
				}
				var results []*svc.ListDatabasesOutput
				p := svc.NewListDatabasesPaginator(client, input)
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
		"list-engine-versions": {
			Name:   "list-engine-versions",
			Fields: fields_list_engine_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEngineVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_engine_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEngineVersions(ctx, input)
				}
				var results []*svc.ListEngineVersionsOutput
				p := svc.NewListEngineVersionsPaginator(client, input)
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
		"list-executors": {
			Name:   "list-executors",
			Fields: fields_list_executors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListExecutorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_executors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListExecutors(ctx, input)
				}
				var results []*svc.ListExecutorsOutput
				p := svc.NewListExecutorsPaginator(client, input)
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
		"list-named-queries": {
			Name:   "list-named-queries",
			Fields: fields_list_named_queries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListNamedQueriesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_named_queries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListNamedQueries(ctx, input)
				}
				var results []*svc.ListNamedQueriesOutput
				p := svc.NewListNamedQueriesPaginator(client, input)
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
		"list-notebook-metadata": {
			Name:   "list-notebook-metadata",
			Fields: fields_list_notebook_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListNotebookMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_notebook_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListNotebookMetadata(ctx, input)
			},
		},
		"list-notebook-sessions": {
			Name:   "list-notebook-sessions",
			Fields: fields_list_notebook_sessions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListNotebookSessionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_notebook_sessions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListNotebookSessions(ctx, input)
			},
		},
		"list-prepared-statements": {
			Name:   "list-prepared-statements",
			Fields: fields_list_prepared_statements,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPreparedStatementsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_prepared_statements, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPreparedStatements(ctx, input)
				}
				var results []*svc.ListPreparedStatementsOutput
				p := svc.NewListPreparedStatementsPaginator(client, input)
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
		"list-query-executions": {
			Name:   "list-query-executions",
			Fields: fields_list_query_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListQueryExecutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_query_executions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListQueryExecutions(ctx, input)
				}
				var results []*svc.ListQueryExecutionsOutput
				p := svc.NewListQueryExecutionsPaginator(client, input)
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
		"list-sessions": {
			Name:   "list-sessions",
			Fields: fields_list_sessions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSessionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_sessions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSessions(ctx, input)
				}
				var results []*svc.ListSessionsOutput
				p := svc.NewListSessionsPaginator(client, input)
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
		"list-table-metadata": {
			Name:   "list-table-metadata",
			Fields: fields_list_table_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTableMetadataInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_table_metadata, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTableMetadata(ctx, input)
				}
				var results []*svc.ListTableMetadataOutput
				p := svc.NewListTableMetadataPaginator(client, input)
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
		"list-work-groups": {
			Name:   "list-work-groups",
			Fields: fields_list_work_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_work_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkGroups(ctx, input)
				}
				var results []*svc.ListWorkGroupsOutput
				p := svc.NewListWorkGroupsPaginator(client, input)
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
		"put-capacity-assignment-configuration": {
			Name:   "put-capacity-assignment-configuration",
			Fields: fields_put_capacity_assignment_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutCapacityAssignmentConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_capacity_assignment_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutCapacityAssignmentConfiguration(ctx, input)
			},
		},
		"start-calculation-execution": {
			Name:   "start-calculation-execution",
			Fields: fields_start_calculation_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartCalculationExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_calculation_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartCalculationExecution(ctx, input)
			},
		},
		"start-query-execution": {
			Name:   "start-query-execution",
			Fields: fields_start_query_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartQueryExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_query_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartQueryExecution(ctx, input)
			},
		},
		"start-session": {
			Name:   "start-session",
			Fields: fields_start_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartSession(ctx, input)
			},
		},
		"stop-calculation-execution": {
			Name:   "stop-calculation-execution",
			Fields: fields_stop_calculation_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopCalculationExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_calculation_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopCalculationExecution(ctx, input)
			},
		},
		"stop-query-execution": {
			Name:   "stop-query-execution",
			Fields: fields_stop_query_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopQueryExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_query_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopQueryExecution(ctx, input)
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
		"terminate-session": {
			Name:   "terminate-session",
			Fields: fields_terminate_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TerminateSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_terminate_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TerminateSession(ctx, input)
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
		"update-capacity-reservation": {
			Name:   "update-capacity-reservation",
			Fields: fields_update_capacity_reservation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCapacityReservationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_capacity_reservation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCapacityReservation(ctx, input)
			},
		},
		"update-data-catalog": {
			Name:   "update-data-catalog",
			Fields: fields_update_data_catalog,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDataCatalogInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_data_catalog, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDataCatalog(ctx, input)
			},
		},
		"update-named-query": {
			Name:   "update-named-query",
			Fields: fields_update_named_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateNamedQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_named_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateNamedQuery(ctx, input)
			},
		},
		"update-notebook": {
			Name:   "update-notebook",
			Fields: fields_update_notebook,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateNotebookInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_notebook, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateNotebook(ctx, input)
			},
		},
		"update-notebook-metadata": {
			Name:   "update-notebook-metadata",
			Fields: fields_update_notebook_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateNotebookMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_notebook_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateNotebookMetadata(ctx, input)
			},
		},
		"update-prepared-statement": {
			Name:   "update-prepared-statement",
			Fields: fields_update_prepared_statement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePreparedStatementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_prepared_statement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePreparedStatement(ctx, input)
			},
		},
		"update-work-group": {
			Name:   "update-work-group",
			Fields: fields_update_work_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWorkGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_work_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWorkGroup(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("athena", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
