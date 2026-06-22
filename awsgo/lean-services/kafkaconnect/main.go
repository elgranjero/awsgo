package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/kafkaconnect"
)

var fields_create_connector = []leanruntime.Field{
	{Name: "Capacity", Flag: "capacity", Type: "*types.Capacity", Required: true},
	{Name: "ConnectorConfiguration", Flag: "connector-configuration", Type: "map[string]string", Required: true},
	{Name: "ConnectorDescription", Flag: "connector-description", Type: "*string", Required: false},
	{Name: "ConnectorName", Flag: "connector-name", Type: "*string", Required: true},
	{Name: "KafkaCluster", Flag: "kafka-cluster", Type: "*types.KafkaCluster", Required: true},
	{Name: "KafkaClusterClientAuthentication", Flag: "kafka-cluster-client-authentication", Type: "*types.KafkaClusterClientAuthentication", Required: true},
	{Name: "KafkaClusterEncryptionInTransit", Flag: "kafka-cluster-encryption-in-transit", Type: "*types.KafkaClusterEncryptionInTransit", Required: true},
	{Name: "KafkaConnectVersion", Flag: "kafka-connect-version", Type: "*string", Required: true},
	{Name: "LogDelivery", Flag: "log-delivery", Type: "*types.LogDelivery", Required: false},
	{Name: "NetworkType", Flag: "network-type", Type: "types.NetworkType", Required: false},
	{Name: "Plugins", Flag: "plugins", Type: "[]types.Plugin", Required: true},
	{Name: "ServiceExecutionRoleArn", Flag: "service-execution-role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "WorkerConfiguration", Flag: "worker-configuration", Type: "*types.WorkerConfiguration", Required: false},
}

var fields_create_custom_plugin = []leanruntime.Field{
	{Name: "ContentType", Flag: "content-type", Type: "types.CustomPluginContentType", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Location", Flag: "location", Type: "*types.CustomPluginLocation", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_worker_configuration = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PropertiesFileContent", Flag: "properties-file-content", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_connector = []leanruntime.Field{
	{Name: "ConnectorArn", Flag: "connector-arn", Type: "*string", Required: true},
	{Name: "CurrentVersion", Flag: "current-version", Type: "*string", Required: false},
}

var fields_delete_custom_plugin = []leanruntime.Field{
	{Name: "CustomPluginArn", Flag: "custom-plugin-arn", Type: "*string", Required: true},
}

var fields_delete_worker_configuration = []leanruntime.Field{
	{Name: "WorkerConfigurationArn", Flag: "worker-configuration-arn", Type: "*string", Required: true},
}

var fields_describe_connector = []leanruntime.Field{
	{Name: "ConnectorArn", Flag: "connector-arn", Type: "*string", Required: true},
}

var fields_describe_connector_operation = []leanruntime.Field{
	{Name: "ConnectorOperationArn", Flag: "connector-operation-arn", Type: "*string", Required: true},
}

var fields_describe_custom_plugin = []leanruntime.Field{
	{Name: "CustomPluginArn", Flag: "custom-plugin-arn", Type: "*string", Required: true},
}

var fields_describe_worker_configuration = []leanruntime.Field{
	{Name: "WorkerConfigurationArn", Flag: "worker-configuration-arn", Type: "*string", Required: true},
}

var fields_list_connector_operations = []leanruntime.Field{
	{Name: "ConnectorArn", Flag: "connector-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_connectors = []leanruntime.Field{
	{Name: "ConnectorNamePrefix", Flag: "connector-name-prefix", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_custom_plugins = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NamePrefix", Flag: "name-prefix", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_worker_configurations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NamePrefix", Flag: "name-prefix", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_connector = []leanruntime.Field{
	{Name: "Capacity", Flag: "capacity", Type: "*types.CapacityUpdate", Required: false},
	{Name: "ConnectorArn", Flag: "connector-arn", Type: "*string", Required: true},
	{Name: "ConnectorConfiguration", Flag: "connector-configuration", Type: "map[string]string", Required: false},
	{Name: "CurrentVersion", Flag: "current-version", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-connector": {
			Name:   "create-connector",
			Fields: fields_create_connector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConnectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_connector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConnector(ctx, input)
			},
		},
		"create-custom-plugin": {
			Name:   "create-custom-plugin",
			Fields: fields_create_custom_plugin,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCustomPluginInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_custom_plugin, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCustomPlugin(ctx, input)
			},
		},
		"create-worker-configuration": {
			Name:   "create-worker-configuration",
			Fields: fields_create_worker_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWorkerConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_worker_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWorkerConfiguration(ctx, input)
			},
		},
		"delete-connector": {
			Name:   "delete-connector",
			Fields: fields_delete_connector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConnectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_connector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConnector(ctx, input)
			},
		},
		"delete-custom-plugin": {
			Name:   "delete-custom-plugin",
			Fields: fields_delete_custom_plugin,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCustomPluginInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_custom_plugin, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCustomPlugin(ctx, input)
			},
		},
		"delete-worker-configuration": {
			Name:   "delete-worker-configuration",
			Fields: fields_delete_worker_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWorkerConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_worker_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWorkerConfiguration(ctx, input)
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
		"describe-connector-operation": {
			Name:   "describe-connector-operation",
			Fields: fields_describe_connector_operation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConnectorOperationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_connector_operation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeConnectorOperation(ctx, input)
			},
		},
		"describe-custom-plugin": {
			Name:   "describe-custom-plugin",
			Fields: fields_describe_custom_plugin,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCustomPluginInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_custom_plugin, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCustomPlugin(ctx, input)
			},
		},
		"describe-worker-configuration": {
			Name:   "describe-worker-configuration",
			Fields: fields_describe_worker_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeWorkerConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_worker_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeWorkerConfiguration(ctx, input)
			},
		},
		"list-connector-operations": {
			Name:   "list-connector-operations",
			Fields: fields_list_connector_operations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConnectorOperationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_connector_operations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConnectorOperations(ctx, input)
				}
				var results []*svc.ListConnectorOperationsOutput
				p := svc.NewListConnectorOperationsPaginator(client, input)
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
		"list-custom-plugins": {
			Name:   "list-custom-plugins",
			Fields: fields_list_custom_plugins,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCustomPluginsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_custom_plugins, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCustomPlugins(ctx, input)
				}
				var results []*svc.ListCustomPluginsOutput
				p := svc.NewListCustomPluginsPaginator(client, input)
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
		"list-worker-configurations": {
			Name:   "list-worker-configurations",
			Fields: fields_list_worker_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkerConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_worker_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkerConfigurations(ctx, input)
				}
				var results []*svc.ListWorkerConfigurationsOutput
				p := svc.NewListWorkerConfigurationsPaginator(client, input)
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
		"update-connector": {
			Name:   "update-connector",
			Fields: fields_update_connector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConnectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_connector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConnector(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("kafkaconnect", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
