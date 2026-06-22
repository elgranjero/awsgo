package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/appflow"
)

var fields_cancel_flow_executions = []leanruntime.Field{
	{Name: "ExecutionIds", Flag: "execution-ids", Type: "[]string", Required: false},
	{Name: "FlowName", Flag: "flow-name", Type: "*string", Required: true},
}

var fields_create_connector_profile = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ConnectionMode", Flag: "connection-mode", Type: "types.ConnectionMode", Required: true},
	{Name: "ConnectorLabel", Flag: "connector-label", Type: "*string", Required: false},
	{Name: "ConnectorProfileConfig", Flag: "connector-profile-config", Type: "*types.ConnectorProfileConfig", Required: true},
	{Name: "ConnectorProfileName", Flag: "connector-profile-name", Type: "*string", Required: true},
	{Name: "ConnectorType", Flag: "connector-type", Type: "types.ConnectorType", Required: true},
	{Name: "KmsArn", Flag: "kms-arn", Type: "*string", Required: false},
}

var fields_create_flow = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DestinationFlowConfigList", Flag: "destination-flow-config-list", Type: "[]types.DestinationFlowConfig", Required: true},
	{Name: "FlowName", Flag: "flow-name", Type: "*string", Required: true},
	{Name: "KmsArn", Flag: "kms-arn", Type: "*string", Required: false},
	{Name: "MetadataCatalogConfig", Flag: "metadata-catalog-config", Type: "*types.MetadataCatalogConfig", Required: false},
	{Name: "SourceFlowConfig", Flag: "source-flow-config", Type: "*types.SourceFlowConfig", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Tasks", Flag: "tasks", Type: "[]types.Task", Required: true},
	{Name: "TriggerConfig", Flag: "trigger-config", Type: "*types.TriggerConfig", Required: true},
}

var fields_delete_connector_profile = []leanruntime.Field{
	{Name: "ConnectorProfileName", Flag: "connector-profile-name", Type: "*string", Required: true},
	{Name: "ForceDelete", Flag: "force-delete", Type: "bool", Required: false},
}

var fields_delete_flow = []leanruntime.Field{
	{Name: "FlowName", Flag: "flow-name", Type: "*string", Required: true},
	{Name: "ForceDelete", Flag: "force-delete", Type: "bool", Required: false},
}

var fields_describe_connector = []leanruntime.Field{
	{Name: "ConnectorLabel", Flag: "connector-label", Type: "*string", Required: false},
	{Name: "ConnectorType", Flag: "connector-type", Type: "types.ConnectorType", Required: true},
}

var fields_describe_connector_entity = []leanruntime.Field{
	{Name: "ApiVersion", Flag: "api-version", Type: "*string", Required: false},
	{Name: "ConnectorEntityName", Flag: "connector-entity-name", Type: "*string", Required: true},
	{Name: "ConnectorProfileName", Flag: "connector-profile-name", Type: "*string", Required: false},
	{Name: "ConnectorType", Flag: "connector-type", Type: "types.ConnectorType", Required: false},
}

var fields_describe_connector_profiles = []leanruntime.Field{
	{Name: "ConnectorLabel", Flag: "connector-label", Type: "*string", Required: false},
	{Name: "ConnectorProfileNames", Flag: "connector-profile-names", Type: "[]string", Required: false},
	{Name: "ConnectorType", Flag: "connector-type", Type: "types.ConnectorType", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_connectors = []leanruntime.Field{
	{Name: "ConnectorTypes", Flag: "connector-types", Type: "[]types.ConnectorType", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_flow = []leanruntime.Field{
	{Name: "FlowName", Flag: "flow-name", Type: "*string", Required: true},
}

var fields_describe_flow_execution_records = []leanruntime.Field{
	{Name: "FlowName", Flag: "flow-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_connector_entities = []leanruntime.Field{
	{Name: "ApiVersion", Flag: "api-version", Type: "*string", Required: false},
	{Name: "ConnectorProfileName", Flag: "connector-profile-name", Type: "*string", Required: false},
	{Name: "ConnectorType", Flag: "connector-type", Type: "types.ConnectorType", Required: false},
	{Name: "EntitiesPath", Flag: "entities-path", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_connectors = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_flows = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_register_connector = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ConnectorLabel", Flag: "connector-label", Type: "*string", Required: false},
	{Name: "ConnectorProvisioningConfig", Flag: "connector-provisioning-config", Type: "*types.ConnectorProvisioningConfig", Required: false},
	{Name: "ConnectorProvisioningType", Flag: "connector-provisioning-type", Type: "types.ConnectorProvisioningType", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
}

var fields_reset_connector_metadata_cache = []leanruntime.Field{
	{Name: "ApiVersion", Flag: "api-version", Type: "*string", Required: false},
	{Name: "ConnectorEntityName", Flag: "connector-entity-name", Type: "*string", Required: false},
	{Name: "ConnectorProfileName", Flag: "connector-profile-name", Type: "*string", Required: false},
	{Name: "ConnectorType", Flag: "connector-type", Type: "types.ConnectorType", Required: false},
	{Name: "EntitiesPath", Flag: "entities-path", Type: "*string", Required: false},
}

var fields_start_flow = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "FlowName", Flag: "flow-name", Type: "*string", Required: true},
}

var fields_stop_flow = []leanruntime.Field{
	{Name: "FlowName", Flag: "flow-name", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_unregister_connector = []leanruntime.Field{
	{Name: "ConnectorLabel", Flag: "connector-label", Type: "*string", Required: true},
	{Name: "ForceDelete", Flag: "force-delete", Type: "bool", Required: false},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_connector_profile = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ConnectionMode", Flag: "connection-mode", Type: "types.ConnectionMode", Required: true},
	{Name: "ConnectorProfileConfig", Flag: "connector-profile-config", Type: "*types.ConnectorProfileConfig", Required: true},
	{Name: "ConnectorProfileName", Flag: "connector-profile-name", Type: "*string", Required: true},
}

var fields_update_connector_registration = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ConnectorLabel", Flag: "connector-label", Type: "*string", Required: true},
	{Name: "ConnectorProvisioningConfig", Flag: "connector-provisioning-config", Type: "*types.ConnectorProvisioningConfig", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
}

var fields_update_flow = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DestinationFlowConfigList", Flag: "destination-flow-config-list", Type: "[]types.DestinationFlowConfig", Required: true},
	{Name: "FlowName", Flag: "flow-name", Type: "*string", Required: true},
	{Name: "MetadataCatalogConfig", Flag: "metadata-catalog-config", Type: "*types.MetadataCatalogConfig", Required: false},
	{Name: "SourceFlowConfig", Flag: "source-flow-config", Type: "*types.SourceFlowConfig", Required: true},
	{Name: "Tasks", Flag: "tasks", Type: "[]types.Task", Required: true},
	{Name: "TriggerConfig", Flag: "trigger-config", Type: "*types.TriggerConfig", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"cancel-flow-executions": {
			Name:   "cancel-flow-executions",
			Fields: fields_cancel_flow_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelFlowExecutionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_flow_executions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelFlowExecutions(ctx, input)
			},
		},
		"create-connector-profile": {
			Name:   "create-connector-profile",
			Fields: fields_create_connector_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConnectorProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_connector_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConnectorProfile(ctx, input)
			},
		},
		"create-flow": {
			Name:   "create-flow",
			Fields: fields_create_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFlow(ctx, input)
			},
		},
		"delete-connector-profile": {
			Name:   "delete-connector-profile",
			Fields: fields_delete_connector_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConnectorProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_connector_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConnectorProfile(ctx, input)
			},
		},
		"delete-flow": {
			Name:   "delete-flow",
			Fields: fields_delete_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFlow(ctx, input)
			},
		},
		"describe-connector": {
			Name:   "describe-connector",
			Fields: fields_describe_connector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConnectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_connector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeConnector(ctx, input)
			},
		},
		"describe-connector-entity": {
			Name:   "describe-connector-entity",
			Fields: fields_describe_connector_entity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConnectorEntityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_connector_entity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeConnectorEntity(ctx, input)
			},
		},
		"describe-connector-profiles": {
			Name:   "describe-connector-profiles",
			Fields: fields_describe_connector_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConnectorProfilesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_connector_profiles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeConnectorProfiles(ctx, input)
				}
				var results []*svc.DescribeConnectorProfilesOutput
				p := svc.NewDescribeConnectorProfilesPaginator(client, input)
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
		"describe-connectors": {
			Name:   "describe-connectors",
			Fields: fields_describe_connectors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConnectorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_connectors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeConnectors(ctx, input)
				}
				var results []*svc.DescribeConnectorsOutput
				p := svc.NewDescribeConnectorsPaginator(client, input)
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
		"describe-flow": {
			Name:   "describe-flow",
			Fields: fields_describe_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFlow(ctx, input)
			},
		},
		"describe-flow-execution-records": {
			Name:   "describe-flow-execution-records",
			Fields: fields_describe_flow_execution_records,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFlowExecutionRecordsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_flow_execution_records, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeFlowExecutionRecords(ctx, input)
				}
				var results []*svc.DescribeFlowExecutionRecordsOutput
				p := svc.NewDescribeFlowExecutionRecordsPaginator(client, input)
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
		"list-connector-entities": {
			Name:   "list-connector-entities",
			Fields: fields_list_connector_entities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConnectorEntitiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_connector_entities, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListConnectorEntities(ctx, input)
			},
		},
		"list-connectors": {
			Name:   "list-connectors",
			Fields: fields_list_connectors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConnectorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_connectors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConnectors(ctx, input)
				}
				var results []*svc.ListConnectorsOutput
				p := svc.NewListConnectorsPaginator(client, input)
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
		"list-flows": {
			Name:   "list-flows",
			Fields: fields_list_flows,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFlowsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_flows, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFlows(ctx, input)
				}
				var results []*svc.ListFlowsOutput
				p := svc.NewListFlowsPaginator(client, input)
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
		"register-connector": {
			Name:   "register-connector",
			Fields: fields_register_connector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterConnectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_connector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterConnector(ctx, input)
			},
		},
		"reset-connector-metadata-cache": {
			Name:   "reset-connector-metadata-cache",
			Fields: fields_reset_connector_metadata_cache,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResetConnectorMetadataCacheInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reset_connector_metadata_cache, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResetConnectorMetadataCache(ctx, input)
			},
		},
		"start-flow": {
			Name:   "start-flow",
			Fields: fields_start_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartFlow(ctx, input)
			},
		},
		"stop-flow": {
			Name:   "stop-flow",
			Fields: fields_stop_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopFlow(ctx, input)
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
		"unregister-connector": {
			Name:   "unregister-connector",
			Fields: fields_unregister_connector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UnregisterConnectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_unregister_connector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UnregisterConnector(ctx, input)
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
		"update-connector-profile": {
			Name:   "update-connector-profile",
			Fields: fields_update_connector_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConnectorProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_connector_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConnectorProfile(ctx, input)
			},
		},
		"update-connector-registration": {
			Name:   "update-connector-registration",
			Fields: fields_update_connector_registration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConnectorRegistrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_connector_registration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConnectorRegistration(ctx, input)
			},
		},
		"update-flow": {
			Name:   "update-flow",
			Fields: fields_update_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFlow(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("appflow", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
