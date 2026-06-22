package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/neptunegraph"
)

var fields_cancel_export_task = []leanruntime.Field{
	{Name: "TaskIdentifier", Flag: "task-identifier", Type: "*string", Required: true},
}

var fields_cancel_import_task = []leanruntime.Field{
	{Name: "TaskIdentifier", Flag: "task-identifier", Type: "*string", Required: true},
}

var fields_cancel_query = []leanruntime.Field{
	{Name: "GraphIdentifier", Flag: "graph-identifier", Type: "*string", Required: true},
	{Name: "QueryId", Flag: "query-id", Type: "*string", Required: true},
}

var fields_create_graph = []leanruntime.Field{
	{Name: "DeletionProtection", Flag: "deletion-protection", Type: "*bool", Required: false},
	{Name: "GraphName", Flag: "graph-name", Type: "*string", Required: true},
	{Name: "KmsKeyIdentifier", Flag: "kms-key-identifier", Type: "*string", Required: false},
	{Name: "ProvisionedMemory", Flag: "provisioned-memory", Type: "*int32", Required: true},
	{Name: "PublicConnectivity", Flag: "public-connectivity", Type: "*bool", Required: false},
	{Name: "ReplicaCount", Flag: "replica-count", Type: "*int32", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "VectorSearchConfiguration", Flag: "vector-search-configuration", Type: "*types.VectorSearchConfiguration", Required: false},
}

var fields_create_graph_snapshot = []leanruntime.Field{
	{Name: "GraphIdentifier", Flag: "graph-identifier", Type: "*string", Required: true},
	{Name: "SnapshotName", Flag: "snapshot-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_graph_using_import_task = []leanruntime.Field{
	{Name: "BlankNodeHandling", Flag: "blank-node-handling", Type: "types.BlankNodeHandling", Required: false},
	{Name: "DeletionProtection", Flag: "deletion-protection", Type: "*bool", Required: false},
	{Name: "FailOnError", Flag: "fail-on-error", Type: "*bool", Required: false},
	{Name: "Format", Flag: "format", Type: "types.Format", Required: false},
	{Name: "GraphName", Flag: "graph-name", Type: "*string", Required: true},
	{Name: "ImportOptions", Flag: "import-options", Type: "types.ImportOptions", Required: false},
	{Name: "KmsKeyIdentifier", Flag: "kms-key-identifier", Type: "*string", Required: false},
	{Name: "MaxProvisionedMemory", Flag: "max-provisioned-memory", Type: "*int32", Required: false},
	{Name: "MinProvisionedMemory", Flag: "min-provisioned-memory", Type: "*int32", Required: false},
	{Name: "ParquetType", Flag: "parquet-type", Type: "types.ParquetType", Required: false},
	{Name: "PublicConnectivity", Flag: "public-connectivity", Type: "*bool", Required: false},
	{Name: "ReplicaCount", Flag: "replica-count", Type: "*int32", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Source", Flag: "source", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "VectorSearchConfiguration", Flag: "vector-search-configuration", Type: "*types.VectorSearchConfiguration", Required: false},
}

var fields_create_private_graph_endpoint = []leanruntime.Field{
	{Name: "GraphIdentifier", Flag: "graph-identifier", Type: "*string", Required: true},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: false},
	{Name: "VpcSecurityGroupIds", Flag: "vpc-security-group-ids", Type: "[]string", Required: false},
}

var fields_delete_graph = []leanruntime.Field{
	{Name: "GraphIdentifier", Flag: "graph-identifier", Type: "*string", Required: true},
	{Name: "SkipSnapshot", Flag: "skip-snapshot", Type: "*bool", Required: true},
}

var fields_delete_graph_snapshot = []leanruntime.Field{
	{Name: "SnapshotIdentifier", Flag: "snapshot-identifier", Type: "*string", Required: true},
}

var fields_delete_private_graph_endpoint = []leanruntime.Field{
	{Name: "GraphIdentifier", Flag: "graph-identifier", Type: "*string", Required: true},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_execute_query = []leanruntime.Field{
	{Name: "ExplainMode", Flag: "explain-mode", Type: "types.ExplainMode", Required: false},
	{Name: "GraphIdentifier", Flag: "graph-identifier", Type: "*string", Required: true},
	{Name: "Language", Flag: "language", Type: "types.QueryLanguage", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "map[string]document.Interface", Required: false},
	{Name: "PlanCache", Flag: "plan-cache", Type: "types.PlanCacheType", Required: false},
	{Name: "QueryString", Flag: "query-string", Type: "*string", Required: true},
	{Name: "QueryTimeoutMilliseconds", Flag: "query-timeout-milliseconds", Type: "*int32", Required: false},
}

var fields_get_export_task = []leanruntime.Field{
	{Name: "TaskIdentifier", Flag: "task-identifier", Type: "*string", Required: true},
}

var fields_get_graph = []leanruntime.Field{
	{Name: "GraphIdentifier", Flag: "graph-identifier", Type: "*string", Required: true},
}

var fields_get_graph_snapshot = []leanruntime.Field{
	{Name: "SnapshotIdentifier", Flag: "snapshot-identifier", Type: "*string", Required: true},
}

var fields_get_graph_summary = []leanruntime.Field{
	{Name: "GraphIdentifier", Flag: "graph-identifier", Type: "*string", Required: true},
	{Name: "Mode", Flag: "mode", Type: "types.GraphSummaryMode", Required: false},
}

var fields_get_import_task = []leanruntime.Field{
	{Name: "TaskIdentifier", Flag: "task-identifier", Type: "*string", Required: true},
}

var fields_get_private_graph_endpoint = []leanruntime.Field{
	{Name: "GraphIdentifier", Flag: "graph-identifier", Type: "*string", Required: true},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_get_query = []leanruntime.Field{
	{Name: "GraphIdentifier", Flag: "graph-identifier", Type: "*string", Required: true},
	{Name: "QueryId", Flag: "query-id", Type: "*string", Required: true},
}

var fields_list_export_tasks = []leanruntime.Field{
	{Name: "GraphIdentifier", Flag: "graph-identifier", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_graph_snapshots = []leanruntime.Field{
	{Name: "GraphIdentifier", Flag: "graph-identifier", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_graphs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_import_tasks = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_private_graph_endpoints = []leanruntime.Field{
	{Name: "GraphIdentifier", Flag: "graph-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_queries = []leanruntime.Field{
	{Name: "GraphIdentifier", Flag: "graph-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: true},
	{Name: "State", Flag: "state", Type: "types.QueryStateInput", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_reset_graph = []leanruntime.Field{
	{Name: "GraphIdentifier", Flag: "graph-identifier", Type: "*string", Required: true},
	{Name: "SkipSnapshot", Flag: "skip-snapshot", Type: "*bool", Required: true},
}

var fields_restore_graph_from_snapshot = []leanruntime.Field{
	{Name: "DeletionProtection", Flag: "deletion-protection", Type: "*bool", Required: false},
	{Name: "GraphName", Flag: "graph-name", Type: "*string", Required: true},
	{Name: "ProvisionedMemory", Flag: "provisioned-memory", Type: "*int32", Required: false},
	{Name: "PublicConnectivity", Flag: "public-connectivity", Type: "*bool", Required: false},
	{Name: "ReplicaCount", Flag: "replica-count", Type: "*int32", Required: false},
	{Name: "SnapshotIdentifier", Flag: "snapshot-identifier", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_start_export_task = []leanruntime.Field{
	{Name: "Destination", Flag: "destination", Type: "*string", Required: true},
	{Name: "ExportFilter", Flag: "export-filter", Type: "*types.ExportFilter", Required: false},
	{Name: "Format", Flag: "format", Type: "types.ExportFormat", Required: true},
	{Name: "GraphIdentifier", Flag: "graph-identifier", Type: "*string", Required: true},
	{Name: "KmsKeyIdentifier", Flag: "kms-key-identifier", Type: "*string", Required: true},
	{Name: "ParquetType", Flag: "parquet-type", Type: "types.ParquetType", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_start_graph = []leanruntime.Field{
	{Name: "GraphIdentifier", Flag: "graph-identifier", Type: "*string", Required: true},
}

var fields_start_import_task = []leanruntime.Field{
	{Name: "BlankNodeHandling", Flag: "blank-node-handling", Type: "types.BlankNodeHandling", Required: false},
	{Name: "FailOnError", Flag: "fail-on-error", Type: "*bool", Required: false},
	{Name: "Format", Flag: "format", Type: "types.Format", Required: false},
	{Name: "GraphIdentifier", Flag: "graph-identifier", Type: "*string", Required: true},
	{Name: "ImportOptions", Flag: "import-options", Type: "types.ImportOptions", Required: false},
	{Name: "ParquetType", Flag: "parquet-type", Type: "types.ParquetType", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Source", Flag: "source", Type: "*string", Required: true},
}

var fields_stop_graph = []leanruntime.Field{
	{Name: "GraphIdentifier", Flag: "graph-identifier", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_graph = []leanruntime.Field{
	{Name: "DeletionProtection", Flag: "deletion-protection", Type: "*bool", Required: false},
	{Name: "GraphIdentifier", Flag: "graph-identifier", Type: "*string", Required: true},
	{Name: "ProvisionedMemory", Flag: "provisioned-memory", Type: "*int32", Required: false},
	{Name: "PublicConnectivity", Flag: "public-connectivity", Type: "*bool", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"cancel-export-task": {
			Name:   "cancel-export-task",
			Fields: fields_cancel_export_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelExportTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_export_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelExportTask(ctx, input)
			},
		},
		"cancel-import-task": {
			Name:   "cancel-import-task",
			Fields: fields_cancel_import_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelImportTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_import_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelImportTask(ctx, input)
			},
		},
		"cancel-query": {
			Name:   "cancel-query",
			Fields: fields_cancel_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelQuery(ctx, input)
			},
		},
		"create-graph": {
			Name:   "create-graph",
			Fields: fields_create_graph,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGraphInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_graph, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGraph(ctx, input)
			},
		},
		"create-graph-snapshot": {
			Name:   "create-graph-snapshot",
			Fields: fields_create_graph_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGraphSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_graph_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGraphSnapshot(ctx, input)
			},
		},
		"create-graph-using-import-task": {
			Name:   "create-graph-using-import-task",
			Fields: fields_create_graph_using_import_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGraphUsingImportTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_graph_using_import_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGraphUsingImportTask(ctx, input)
			},
		},
		"create-private-graph-endpoint": {
			Name:   "create-private-graph-endpoint",
			Fields: fields_create_private_graph_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePrivateGraphEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_private_graph_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePrivateGraphEndpoint(ctx, input)
			},
		},
		"delete-graph": {
			Name:   "delete-graph",
			Fields: fields_delete_graph,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGraphInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_graph, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGraph(ctx, input)
			},
		},
		"delete-graph-snapshot": {
			Name:   "delete-graph-snapshot",
			Fields: fields_delete_graph_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGraphSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_graph_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGraphSnapshot(ctx, input)
			},
		},
		"delete-private-graph-endpoint": {
			Name:   "delete-private-graph-endpoint",
			Fields: fields_delete_private_graph_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePrivateGraphEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_private_graph_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePrivateGraphEndpoint(ctx, input)
			},
		},
		"execute-query": {
			Name:   "execute-query",
			Fields: fields_execute_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExecuteQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_execute_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExecuteQuery(ctx, input)
			},
		},
		"get-export-task": {
			Name:   "get-export-task",
			Fields: fields_get_export_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetExportTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_export_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetExportTask(ctx, input)
			},
		},
		"get-graph": {
			Name:   "get-graph",
			Fields: fields_get_graph,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGraphInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_graph, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGraph(ctx, input)
			},
		},
		"get-graph-snapshot": {
			Name:   "get-graph-snapshot",
			Fields: fields_get_graph_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGraphSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_graph_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGraphSnapshot(ctx, input)
			},
		},
		"get-graph-summary": {
			Name:   "get-graph-summary",
			Fields: fields_get_graph_summary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGraphSummaryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_graph_summary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGraphSummary(ctx, input)
			},
		},
		"get-import-task": {
			Name:   "get-import-task",
			Fields: fields_get_import_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetImportTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_import_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetImportTask(ctx, input)
			},
		},
		"get-private-graph-endpoint": {
			Name:   "get-private-graph-endpoint",
			Fields: fields_get_private_graph_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPrivateGraphEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_private_graph_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPrivateGraphEndpoint(ctx, input)
			},
		},
		"get-query": {
			Name:   "get-query",
			Fields: fields_get_query,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetQueryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_query, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetQuery(ctx, input)
			},
		},
		"list-export-tasks": {
			Name:   "list-export-tasks",
			Fields: fields_list_export_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListExportTasksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_export_tasks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListExportTasks(ctx, input)
				}
				var results []*svc.ListExportTasksOutput
				p := svc.NewListExportTasksPaginator(client, input)
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
		"list-graph-snapshots": {
			Name:   "list-graph-snapshots",
			Fields: fields_list_graph_snapshots,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGraphSnapshotsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_graph_snapshots, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGraphSnapshots(ctx, input)
				}
				var results []*svc.ListGraphSnapshotsOutput
				p := svc.NewListGraphSnapshotsPaginator(client, input)
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
		"list-graphs": {
			Name:   "list-graphs",
			Fields: fields_list_graphs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGraphsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_graphs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGraphs(ctx, input)
				}
				var results []*svc.ListGraphsOutput
				p := svc.NewListGraphsPaginator(client, input)
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
		"list-import-tasks": {
			Name:   "list-import-tasks",
			Fields: fields_list_import_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListImportTasksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_import_tasks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListImportTasks(ctx, input)
				}
				var results []*svc.ListImportTasksOutput
				p := svc.NewListImportTasksPaginator(client, input)
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
		"list-private-graph-endpoints": {
			Name:   "list-private-graph-endpoints",
			Fields: fields_list_private_graph_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPrivateGraphEndpointsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_private_graph_endpoints, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPrivateGraphEndpoints(ctx, input)
				}
				var results []*svc.ListPrivateGraphEndpointsOutput
				p := svc.NewListPrivateGraphEndpointsPaginator(client, input)
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
		"list-queries": {
			Name:   "list-queries",
			Fields: fields_list_queries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListQueriesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_queries, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListQueries(ctx, input)
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
		"reset-graph": {
			Name:   "reset-graph",
			Fields: fields_reset_graph,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResetGraphInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reset_graph, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResetGraph(ctx, input)
			},
		},
		"restore-graph-from-snapshot": {
			Name:   "restore-graph-from-snapshot",
			Fields: fields_restore_graph_from_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestoreGraphFromSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_graph_from_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestoreGraphFromSnapshot(ctx, input)
			},
		},
		"start-export-task": {
			Name:   "start-export-task",
			Fields: fields_start_export_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartExportTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_export_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartExportTask(ctx, input)
			},
		},
		"start-graph": {
			Name:   "start-graph",
			Fields: fields_start_graph,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartGraphInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_graph, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartGraph(ctx, input)
			},
		},
		"start-import-task": {
			Name:   "start-import-task",
			Fields: fields_start_import_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartImportTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_import_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartImportTask(ctx, input)
			},
		},
		"stop-graph": {
			Name:   "stop-graph",
			Fields: fields_stop_graph,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopGraphInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_graph, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopGraph(ctx, input)
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
		"update-graph": {
			Name:   "update-graph",
			Fields: fields_update_graph,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGraphInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_graph, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGraph(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("neptunegraph", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
