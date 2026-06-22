package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/supplychain"
)

var fields_create_bill_of_materials_import_job = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "S3uri", Flag: "s3uri", Type: "*string", Required: true},
}

var fields_create_data_integration_flow = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Sources", Flag: "sources", Type: "[]types.DataIntegrationFlowSource", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Target", Flag: "target", Type: "*types.DataIntegrationFlowTarget", Required: true},
	{Name: "Transformation", Flag: "transformation", Type: "*types.DataIntegrationFlowTransformation", Required: true},
}

var fields_create_data_lake_dataset = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "PartitionSpec", Flag: "partition-spec", Type: "*types.DataLakeDatasetPartitionSpec", Required: false},
	{Name: "Schema", Flag: "schema", Type: "*types.DataLakeDatasetSchema", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_data_lake_namespace = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_instance = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "InstanceDescription", Flag: "instance-description", Type: "*string", Required: false},
	{Name: "InstanceName", Flag: "instance-name", Type: "*string", Required: false},
	{Name: "KmsKeyArn", Flag: "kms-key-arn", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "WebAppDnsDomain", Flag: "web-app-dns-domain", Type: "*string", Required: false},
}

var fields_delete_data_integration_flow = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_data_lake_dataset = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
}

var fields_delete_data_lake_namespace = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_instance = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_get_bill_of_materials_import_job = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_get_data_integration_event = []leanruntime.Field{
	{Name: "EventId", Flag: "event-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_get_data_integration_flow = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_data_integration_flow_execution = []leanruntime.Field{
	{Name: "ExecutionId", Flag: "execution-id", Type: "*string", Required: true},
	{Name: "FlowName", Flag: "flow-name", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_get_data_lake_dataset = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
}

var fields_get_data_lake_namespace = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_instance = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_list_data_integration_events = []leanruntime.Field{
	{Name: "EventType", Flag: "event-type", Type: "types.DataIntegrationEventType", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_data_integration_flow_executions = []leanruntime.Field{
	{Name: "FlowName", Flag: "flow-name", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_data_integration_flows = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_data_lake_datasets = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_data_lake_namespaces = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_instances = []leanruntime.Field{
	{Name: "InstanceNameFilter", Flag: "instance-name-filter", Type: "[]string", Required: false},
	{Name: "InstanceStateFilter", Flag: "instance-state-filter", Type: "[]types.InstanceState", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_send_data_integration_event = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Data", Flag: "data", Type: "*string", Required: true},
	{Name: "DatasetTarget", Flag: "dataset-target", Type: "*types.DataIntegrationEventDatasetTargetConfiguration", Required: false},
	{Name: "EventGroupId", Flag: "event-group-id", Type: "*string", Required: true},
	{Name: "EventTimestamp", Flag: "event-timestamp", Type: "*time.Time", Required: false},
	{Name: "EventType", Flag: "event-type", Type: "types.DataIntegrationEventType", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_data_integration_flow = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Sources", Flag: "sources", Type: "[]types.DataIntegrationFlowSource", Required: false},
	{Name: "Target", Flag: "target", Type: "*types.DataIntegrationFlowTarget", Required: false},
	{Name: "Transformation", Flag: "transformation", Type: "*types.DataIntegrationFlowTransformation", Required: false},
}

var fields_update_data_lake_dataset = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
}

var fields_update_data_lake_namespace = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_update_instance = []leanruntime.Field{
	{Name: "InstanceDescription", Flag: "instance-description", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "InstanceName", Flag: "instance-name", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-bill-of-materials-import-job": {
			Name:   "create-bill-of-materials-import-job",
			Fields: fields_create_bill_of_materials_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBillOfMaterialsImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_bill_of_materials_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBillOfMaterialsImportJob(ctx, input)
			},
		},
		"create-data-integration-flow": {
			Name:   "create-data-integration-flow",
			Fields: fields_create_data_integration_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataIntegrationFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_integration_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataIntegrationFlow(ctx, input)
			},
		},
		"create-data-lake-dataset": {
			Name:   "create-data-lake-dataset",
			Fields: fields_create_data_lake_dataset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataLakeDatasetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_lake_dataset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataLakeDataset(ctx, input)
			},
		},
		"create-data-lake-namespace": {
			Name:   "create-data-lake-namespace",
			Fields: fields_create_data_lake_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataLakeNamespaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_lake_namespace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataLakeNamespace(ctx, input)
			},
		},
		"create-instance": {
			Name:   "create-instance",
			Fields: fields_create_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateInstance(ctx, input)
			},
		},
		"delete-data-integration-flow": {
			Name:   "delete-data-integration-flow",
			Fields: fields_delete_data_integration_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDataIntegrationFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_data_integration_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataIntegrationFlow(ctx, input)
			},
		},
		"delete-data-lake-dataset": {
			Name:   "delete-data-lake-dataset",
			Fields: fields_delete_data_lake_dataset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDataLakeDatasetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_data_lake_dataset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataLakeDataset(ctx, input)
			},
		},
		"delete-data-lake-namespace": {
			Name:   "delete-data-lake-namespace",
			Fields: fields_delete_data_lake_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDataLakeNamespaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_data_lake_namespace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataLakeNamespace(ctx, input)
			},
		},
		"delete-instance": {
			Name:   "delete-instance",
			Fields: fields_delete_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteInstance(ctx, input)
			},
		},
		"get-bill-of-materials-import-job": {
			Name:   "get-bill-of-materials-import-job",
			Fields: fields_get_bill_of_materials_import_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBillOfMaterialsImportJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bill_of_materials_import_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBillOfMaterialsImportJob(ctx, input)
			},
		},
		"get-data-integration-event": {
			Name:   "get-data-integration-event",
			Fields: fields_get_data_integration_event,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataIntegrationEventInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_integration_event, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataIntegrationEvent(ctx, input)
			},
		},
		"get-data-integration-flow": {
			Name:   "get-data-integration-flow",
			Fields: fields_get_data_integration_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataIntegrationFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_integration_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataIntegrationFlow(ctx, input)
			},
		},
		"get-data-integration-flow-execution": {
			Name:   "get-data-integration-flow-execution",
			Fields: fields_get_data_integration_flow_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataIntegrationFlowExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_integration_flow_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataIntegrationFlowExecution(ctx, input)
			},
		},
		"get-data-lake-dataset": {
			Name:   "get-data-lake-dataset",
			Fields: fields_get_data_lake_dataset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataLakeDatasetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_lake_dataset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataLakeDataset(ctx, input)
			},
		},
		"get-data-lake-namespace": {
			Name:   "get-data-lake-namespace",
			Fields: fields_get_data_lake_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataLakeNamespaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_lake_namespace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataLakeNamespace(ctx, input)
			},
		},
		"get-instance": {
			Name:   "get-instance",
			Fields: fields_get_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInstance(ctx, input)
			},
		},
		"list-data-integration-events": {
			Name:   "list-data-integration-events",
			Fields: fields_list_data_integration_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataIntegrationEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_integration_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataIntegrationEvents(ctx, input)
				}
				var results []*svc.ListDataIntegrationEventsOutput
				p := svc.NewListDataIntegrationEventsPaginator(client, input)
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
		"list-data-integration-flow-executions": {
			Name:   "list-data-integration-flow-executions",
			Fields: fields_list_data_integration_flow_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataIntegrationFlowExecutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_integration_flow_executions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataIntegrationFlowExecutions(ctx, input)
				}
				var results []*svc.ListDataIntegrationFlowExecutionsOutput
				p := svc.NewListDataIntegrationFlowExecutionsPaginator(client, input)
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
		"list-data-integration-flows": {
			Name:   "list-data-integration-flows",
			Fields: fields_list_data_integration_flows,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataIntegrationFlowsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_integration_flows, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataIntegrationFlows(ctx, input)
				}
				var results []*svc.ListDataIntegrationFlowsOutput
				p := svc.NewListDataIntegrationFlowsPaginator(client, input)
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
		"list-data-lake-datasets": {
			Name:   "list-data-lake-datasets",
			Fields: fields_list_data_lake_datasets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataLakeDatasetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_lake_datasets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataLakeDatasets(ctx, input)
				}
				var results []*svc.ListDataLakeDatasetsOutput
				p := svc.NewListDataLakeDatasetsPaginator(client, input)
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
		"list-data-lake-namespaces": {
			Name:   "list-data-lake-namespaces",
			Fields: fields_list_data_lake_namespaces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataLakeNamespacesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_lake_namespaces, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataLakeNamespaces(ctx, input)
				}
				var results []*svc.ListDataLakeNamespacesOutput
				p := svc.NewListDataLakeNamespacesPaginator(client, input)
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
		"list-instances": {
			Name:   "list-instances",
			Fields: fields_list_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInstances(ctx, input)
				}
				var results []*svc.ListInstancesOutput
				p := svc.NewListInstancesPaginator(client, input)
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
		"send-data-integration-event": {
			Name:   "send-data-integration-event",
			Fields: fields_send_data_integration_event,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendDataIntegrationEventInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_data_integration_event, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendDataIntegrationEvent(ctx, input)
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
		"update-data-integration-flow": {
			Name:   "update-data-integration-flow",
			Fields: fields_update_data_integration_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDataIntegrationFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_data_integration_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDataIntegrationFlow(ctx, input)
			},
		},
		"update-data-lake-dataset": {
			Name:   "update-data-lake-dataset",
			Fields: fields_update_data_lake_dataset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDataLakeDatasetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_data_lake_dataset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDataLakeDataset(ctx, input)
			},
		},
		"update-data-lake-namespace": {
			Name:   "update-data-lake-namespace",
			Fields: fields_update_data_lake_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDataLakeNamespaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_data_lake_namespace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDataLakeNamespace(ctx, input)
			},
		},
		"update-instance": {
			Name:   "update-instance",
			Fields: fields_update_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateInstance(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("supplychain", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
