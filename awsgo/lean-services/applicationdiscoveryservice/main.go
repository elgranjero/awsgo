package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice"
)

var fields_associate_configuration_items_to_application = []leanruntime.Field{
	{Name: "ApplicationConfigurationId", Flag: "application-configuration-id", Type: "*string", Required: true},
	{Name: "ConfigurationIds", Flag: "configuration-ids", Type: "[]string", Required: true},
}

var fields_batch_delete_agents = []leanruntime.Field{
	{Name: "DeleteAgents", Flag: "delete-agents", Type: "[]types.DeleteAgent", Required: true},
}

var fields_batch_delete_import_data = []leanruntime.Field{
	{Name: "DeleteHistory", Flag: "delete-history", Type: "bool", Required: false},
	{Name: "ImportTaskIds", Flag: "import-task-ids", Type: "[]string", Required: true},
}

var fields_create_application = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Wave", Flag: "wave", Type: "*string", Required: false},
}

var fields_create_tags = []leanruntime.Field{
	{Name: "ConfigurationIds", Flag: "configuration-ids", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_delete_applications = []leanruntime.Field{
	{Name: "ConfigurationIds", Flag: "configuration-ids", Type: "[]string", Required: true},
}

var fields_delete_tags = []leanruntime.Field{
	{Name: "ConfigurationIds", Flag: "configuration-ids", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_describe_agents = []leanruntime.Field{
	{Name: "AgentIds", Flag: "agent-ids", Type: "[]string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_batch_delete_configuration_task = []leanruntime.Field{
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: true},
}

var fields_describe_configurations = []leanruntime.Field{
	{Name: "ConfigurationIds", Flag: "configuration-ids", Type: "[]string", Required: true},
}

var fields_describe_continuous_exports = []leanruntime.Field{
	{Name: "ExportIds", Flag: "export-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_export_configurations = []leanruntime.Field{
	{Name: "ExportIds", Flag: "export-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_export_tasks = []leanruntime.Field{
	{Name: "ExportIds", Flag: "export-ids", Type: "[]string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.ExportFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_import_tasks = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.ImportTaskFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_tags = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.TagFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_disassociate_configuration_items_from_application = []leanruntime.Field{
	{Name: "ApplicationConfigurationId", Flag: "application-configuration-id", Type: "*string", Required: true},
	{Name: "ConfigurationIds", Flag: "configuration-ids", Type: "[]string", Required: true},
}

var fields_export_configurations = []leanruntime.Field{}

var fields_get_discovery_summary = []leanruntime.Field{}

var fields_list_configurations = []leanruntime.Field{
	{Name: "ConfigurationType", Flag: "configuration-type", Type: "types.ConfigurationItemType", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OrderBy", Flag: "order-by", Type: "[]types.OrderByElement", Required: false},
}

var fields_list_server_neighbors = []leanruntime.Field{
	{Name: "ConfigurationId", Flag: "configuration-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NeighborConfigurationIds", Flag: "neighbor-configuration-ids", Type: "[]string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PortInformationNeeded", Flag: "port-information-needed", Type: "bool", Required: false},
}

var fields_start_batch_delete_configuration_task = []leanruntime.Field{
	{Name: "ConfigurationIds", Flag: "configuration-ids", Type: "[]string", Required: true},
	{Name: "ConfigurationType", Flag: "configuration-type", Type: "types.DeletionConfigurationItemType", Required: true},
}

var fields_start_continuous_export = []leanruntime.Field{}

var fields_start_data_collection_by_agent_ids = []leanruntime.Field{
	{Name: "AgentIds", Flag: "agent-ids", Type: "[]string", Required: true},
}

var fields_start_export_task = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "ExportDataFormat", Flag: "export-data-format", Type: "[]types.ExportDataFormat", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.ExportFilter", Required: false},
	{Name: "Preferences", Flag: "preferences", Type: "types.ExportPreferences", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
}

var fields_start_import_task = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "ImportUrl", Flag: "import-url", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_stop_continuous_export = []leanruntime.Field{
	{Name: "ExportId", Flag: "export-id", Type: "*string", Required: true},
}

var fields_stop_data_collection_by_agent_ids = []leanruntime.Field{
	{Name: "AgentIds", Flag: "agent-ids", Type: "[]string", Required: true},
}

var fields_update_application = []leanruntime.Field{
	{Name: "ConfigurationId", Flag: "configuration-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Wave", Flag: "wave", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-configuration-items-to-application": {
			Name:   "associate-configuration-items-to-application",
			Fields: fields_associate_configuration_items_to_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateConfigurationItemsToApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_configuration_items_to_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateConfigurationItemsToApplication(ctx, input)
			},
		},
		"batch-delete-agents": {
			Name:   "batch-delete-agents",
			Fields: fields_batch_delete_agents,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeleteAgentsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete_agents, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDeleteAgents(ctx, input)
			},
		},
		"batch-delete-import-data": {
			Name:   "batch-delete-import-data",
			Fields: fields_batch_delete_import_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeleteImportDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete_import_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDeleteImportData(ctx, input)
			},
		},
		"create-application": {
			Name:   "create-application",
			Fields: fields_create_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateApplication(ctx, input)
			},
		},
		"create-tags": {
			Name:   "create-tags",
			Fields: fields_create_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTags(ctx, input)
			},
		},
		"delete-applications": {
			Name:   "delete-applications",
			Fields: fields_delete_applications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteApplicationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_applications, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApplications(ctx, input)
			},
		},
		"delete-tags": {
			Name:   "delete-tags",
			Fields: fields_delete_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTags(ctx, input)
			},
		},
		"describe-agents": {
			Name:   "describe-agents",
			Fields: fields_describe_agents,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAgentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_agents, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeAgents(ctx, input)
				}
				var results []*svc.DescribeAgentsOutput
				p := svc.NewDescribeAgentsPaginator(client, input)
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
		"describe-batch-delete-configuration-task": {
			Name:   "describe-batch-delete-configuration-task",
			Fields: fields_describe_batch_delete_configuration_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBatchDeleteConfigurationTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_batch_delete_configuration_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeBatchDeleteConfigurationTask(ctx, input)
			},
		},
		"describe-configurations": {
			Name:   "describe-configurations",
			Fields: fields_describe_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConfigurationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_configurations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeConfigurations(ctx, input)
			},
		},
		"describe-continuous-exports": {
			Name:   "describe-continuous-exports",
			Fields: fields_describe_continuous_exports,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeContinuousExportsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_continuous_exports, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeContinuousExports(ctx, input)
				}
				var results []*svc.DescribeContinuousExportsOutput
				p := svc.NewDescribeContinuousExportsPaginator(client, input)
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
		"describe-export-configurations": {
			Name:   "describe-export-configurations",
			Fields: fields_describe_export_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeExportConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_export_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeExportConfigurations(ctx, input)
				}
				var results []*svc.DescribeExportConfigurationsOutput
				p := svc.NewDescribeExportConfigurationsPaginator(client, input)
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
		"describe-export-tasks": {
			Name:   "describe-export-tasks",
			Fields: fields_describe_export_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeExportTasksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_export_tasks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeExportTasks(ctx, input)
				}
				var results []*svc.DescribeExportTasksOutput
				p := svc.NewDescribeExportTasksPaginator(client, input)
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
		"describe-import-tasks": {
			Name:   "describe-import-tasks",
			Fields: fields_describe_import_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeImportTasksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_import_tasks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeImportTasks(ctx, input)
				}
				var results []*svc.DescribeImportTasksOutput
				p := svc.NewDescribeImportTasksPaginator(client, input)
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
		"describe-tags": {
			Name:   "describe-tags",
			Fields: fields_describe_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTagsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_tags, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeTags(ctx, input)
				}
				var results []*svc.DescribeTagsOutput
				p := svc.NewDescribeTagsPaginator(client, input)
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
		"disassociate-configuration-items-from-application": {
			Name:   "disassociate-configuration-items-from-application",
			Fields: fields_disassociate_configuration_items_from_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateConfigurationItemsFromApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_configuration_items_from_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateConfigurationItemsFromApplication(ctx, input)
			},
		},
		"export-configurations": {
			Name:   "export-configurations",
			Fields: fields_export_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExportConfigurationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_export_configurations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExportConfigurations(ctx, input)
			},
		},
		"get-discovery-summary": {
			Name:   "get-discovery-summary",
			Fields: fields_get_discovery_summary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDiscoverySummaryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_discovery_summary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDiscoverySummary(ctx, input)
			},
		},
		"list-configurations": {
			Name:   "list-configurations",
			Fields: fields_list_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConfigurations(ctx, input)
				}
				var results []*svc.ListConfigurationsOutput
				p := svc.NewListConfigurationsPaginator(client, input)
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
		"list-server-neighbors": {
			Name:   "list-server-neighbors",
			Fields: fields_list_server_neighbors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServerNeighborsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_server_neighbors, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListServerNeighbors(ctx, input)
			},
		},
		"start-batch-delete-configuration-task": {
			Name:   "start-batch-delete-configuration-task",
			Fields: fields_start_batch_delete_configuration_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartBatchDeleteConfigurationTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_batch_delete_configuration_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartBatchDeleteConfigurationTask(ctx, input)
			},
		},
		"start-continuous-export": {
			Name:   "start-continuous-export",
			Fields: fields_start_continuous_export,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartContinuousExportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_continuous_export, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartContinuousExport(ctx, input)
			},
		},
		"start-data-collection-by-agent-ids": {
			Name:   "start-data-collection-by-agent-ids",
			Fields: fields_start_data_collection_by_agent_ids,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDataCollectionByAgentIdsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_data_collection_by_agent_ids, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDataCollectionByAgentIds(ctx, input)
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
		"stop-continuous-export": {
			Name:   "stop-continuous-export",
			Fields: fields_stop_continuous_export,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopContinuousExportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_continuous_export, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopContinuousExport(ctx, input)
			},
		},
		"stop-data-collection-by-agent-ids": {
			Name:   "stop-data-collection-by-agent-ids",
			Fields: fields_stop_data_collection_by_agent_ids,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopDataCollectionByAgentIdsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_data_collection_by_agent_ids, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopDataCollectionByAgentIds(ctx, input)
			},
		},
		"update-application": {
			Name:   "update-application",
			Fields: fields_update_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateApplication(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("applicationdiscoveryservice", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
