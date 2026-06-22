package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/apprunner"
)

var fields_associate_custom_domain = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "EnableWWWSubdomain", Flag: "enable-www-subdomain", Type: "*bool", Required: false},
	{Name: "ServiceArn", Flag: "service-arn", Type: "*string", Required: true},
}

var fields_create_auto_scaling_configuration = []leanruntime.Field{
	{Name: "AutoScalingConfigurationName", Flag: "auto-scaling-configuration-name", Type: "*string", Required: true},
	{Name: "MaxConcurrency", Flag: "max-concurrency", Type: "*int32", Required: false},
	{Name: "MaxSize", Flag: "max-size", Type: "*int32", Required: false},
	{Name: "MinSize", Flag: "min-size", Type: "*int32", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_connection = []leanruntime.Field{
	{Name: "ConnectionName", Flag: "connection-name", Type: "*string", Required: true},
	{Name: "ProviderType", Flag: "provider-type", Type: "types.ProviderType", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_observability_configuration = []leanruntime.Field{
	{Name: "ObservabilityConfigurationName", Flag: "observability-configuration-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TraceConfiguration", Flag: "trace-configuration", Type: "*types.TraceConfiguration", Required: false},
}

var fields_create_service = []leanruntime.Field{
	{Name: "AutoScalingConfigurationArn", Flag: "auto-scaling-configuration-arn", Type: "*string", Required: false},
	{Name: "EncryptionConfiguration", Flag: "encryption-configuration", Type: "*types.EncryptionConfiguration", Required: false},
	{Name: "HealthCheckConfiguration", Flag: "health-check-configuration", Type: "*types.HealthCheckConfiguration", Required: false},
	{Name: "InstanceConfiguration", Flag: "instance-configuration", Type: "*types.InstanceConfiguration", Required: false},
	{Name: "NetworkConfiguration", Flag: "network-configuration", Type: "*types.NetworkConfiguration", Required: false},
	{Name: "ObservabilityConfiguration", Flag: "observability-configuration", Type: "*types.ServiceObservabilityConfiguration", Required: false},
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: true},
	{Name: "SourceConfiguration", Flag: "source-configuration", Type: "*types.SourceConfiguration", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_vpc_connector = []leanruntime.Field{
	{Name: "SecurityGroups", Flag: "security-groups", Type: "[]string", Required: false},
	{Name: "Subnets", Flag: "subnets", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VpcConnectorName", Flag: "vpc-connector-name", Type: "*string", Required: true},
}

var fields_create_vpc_ingress_connection = []leanruntime.Field{
	{Name: "IngressVpcConfiguration", Flag: "ingress-vpc-configuration", Type: "*types.IngressVpcConfiguration", Required: true},
	{Name: "ServiceArn", Flag: "service-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VpcIngressConnectionName", Flag: "vpc-ingress-connection-name", Type: "*string", Required: true},
}

var fields_delete_auto_scaling_configuration = []leanruntime.Field{
	{Name: "AutoScalingConfigurationArn", Flag: "auto-scaling-configuration-arn", Type: "*string", Required: true},
	{Name: "DeleteAllRevisions", Flag: "delete-all-revisions", Type: "bool", Required: false},
}

var fields_delete_connection = []leanruntime.Field{
	{Name: "ConnectionArn", Flag: "connection-arn", Type: "*string", Required: true},
}

var fields_delete_observability_configuration = []leanruntime.Field{
	{Name: "ObservabilityConfigurationArn", Flag: "observability-configuration-arn", Type: "*string", Required: true},
}

var fields_delete_service = []leanruntime.Field{
	{Name: "ServiceArn", Flag: "service-arn", Type: "*string", Required: true},
}

var fields_delete_vpc_connector = []leanruntime.Field{
	{Name: "VpcConnectorArn", Flag: "vpc-connector-arn", Type: "*string", Required: true},
}

var fields_delete_vpc_ingress_connection = []leanruntime.Field{
	{Name: "VpcIngressConnectionArn", Flag: "vpc-ingress-connection-arn", Type: "*string", Required: true},
}

var fields_describe_auto_scaling_configuration = []leanruntime.Field{
	{Name: "AutoScalingConfigurationArn", Flag: "auto-scaling-configuration-arn", Type: "*string", Required: true},
}

var fields_describe_custom_domains = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServiceArn", Flag: "service-arn", Type: "*string", Required: true},
}

var fields_describe_observability_configuration = []leanruntime.Field{
	{Name: "ObservabilityConfigurationArn", Flag: "observability-configuration-arn", Type: "*string", Required: true},
}

var fields_describe_service = []leanruntime.Field{
	{Name: "ServiceArn", Flag: "service-arn", Type: "*string", Required: true},
}

var fields_describe_vpc_connector = []leanruntime.Field{
	{Name: "VpcConnectorArn", Flag: "vpc-connector-arn", Type: "*string", Required: true},
}

var fields_describe_vpc_ingress_connection = []leanruntime.Field{
	{Name: "VpcIngressConnectionArn", Flag: "vpc-ingress-connection-arn", Type: "*string", Required: true},
}

var fields_disassociate_custom_domain = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "ServiceArn", Flag: "service-arn", Type: "*string", Required: true},
}

var fields_list_auto_scaling_configurations = []leanruntime.Field{
	{Name: "AutoScalingConfigurationName", Flag: "auto-scaling-configuration-name", Type: "*string", Required: false},
	{Name: "LatestOnly", Flag: "latest-only", Type: "bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_connections = []leanruntime.Field{
	{Name: "ConnectionName", Flag: "connection-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_observability_configurations = []leanruntime.Field{
	{Name: "LatestOnly", Flag: "latest-only", Type: "bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ObservabilityConfigurationName", Flag: "observability-configuration-name", Type: "*string", Required: false},
}

var fields_list_operations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServiceArn", Flag: "service-arn", Type: "*string", Required: true},
}

var fields_list_services = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_services_for_auto_scaling_configuration = []leanruntime.Field{
	{Name: "AutoScalingConfigurationArn", Flag: "auto-scaling-configuration-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_vpc_connectors = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_vpc_ingress_connections = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.ListVpcIngressConnectionsFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_pause_service = []leanruntime.Field{
	{Name: "ServiceArn", Flag: "service-arn", Type: "*string", Required: true},
}

var fields_resume_service = []leanruntime.Field{
	{Name: "ServiceArn", Flag: "service-arn", Type: "*string", Required: true},
}

var fields_start_deployment = []leanruntime.Field{
	{Name: "ServiceArn", Flag: "service-arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_default_auto_scaling_configuration = []leanruntime.Field{
	{Name: "AutoScalingConfigurationArn", Flag: "auto-scaling-configuration-arn", Type: "*string", Required: true},
}

var fields_update_service = []leanruntime.Field{
	{Name: "AutoScalingConfigurationArn", Flag: "auto-scaling-configuration-arn", Type: "*string", Required: false},
	{Name: "HealthCheckConfiguration", Flag: "health-check-configuration", Type: "*types.HealthCheckConfiguration", Required: false},
	{Name: "InstanceConfiguration", Flag: "instance-configuration", Type: "*types.InstanceConfiguration", Required: false},
	{Name: "NetworkConfiguration", Flag: "network-configuration", Type: "*types.NetworkConfiguration", Required: false},
	{Name: "ObservabilityConfiguration", Flag: "observability-configuration", Type: "*types.ServiceObservabilityConfiguration", Required: false},
	{Name: "ServiceArn", Flag: "service-arn", Type: "*string", Required: true},
	{Name: "SourceConfiguration", Flag: "source-configuration", Type: "*types.SourceConfiguration", Required: false},
}

var fields_update_vpc_ingress_connection = []leanruntime.Field{
	{Name: "IngressVpcConfiguration", Flag: "ingress-vpc-configuration", Type: "*types.IngressVpcConfiguration", Required: true},
	{Name: "VpcIngressConnectionArn", Flag: "vpc-ingress-connection-arn", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-custom-domain": {
			Name:   "associate-custom-domain",
			Fields: fields_associate_custom_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateCustomDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_custom_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateCustomDomain(ctx, input)
			},
		},
		"create-auto-scaling-configuration": {
			Name:   "create-auto-scaling-configuration",
			Fields: fields_create_auto_scaling_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAutoScalingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_auto_scaling_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAutoScalingConfiguration(ctx, input)
			},
		},
		"create-connection": {
			Name:   "create-connection",
			Fields: fields_create_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConnection(ctx, input)
			},
		},
		"create-observability-configuration": {
			Name:   "create-observability-configuration",
			Fields: fields_create_observability_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateObservabilityConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_observability_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateObservabilityConfiguration(ctx, input)
			},
		},
		"create-service": {
			Name:   "create-service",
			Fields: fields_create_service,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateServiceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_service, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateService(ctx, input)
			},
		},
		"create-vpc-connector": {
			Name:   "create-vpc-connector",
			Fields: fields_create_vpc_connector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVpcConnectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_vpc_connector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVpcConnector(ctx, input)
			},
		},
		"create-vpc-ingress-connection": {
			Name:   "create-vpc-ingress-connection",
			Fields: fields_create_vpc_ingress_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVpcIngressConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_vpc_ingress_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVpcIngressConnection(ctx, input)
			},
		},
		"delete-auto-scaling-configuration": {
			Name:   "delete-auto-scaling-configuration",
			Fields: fields_delete_auto_scaling_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAutoScalingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_auto_scaling_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAutoScalingConfiguration(ctx, input)
			},
		},
		"delete-connection": {
			Name:   "delete-connection",
			Fields: fields_delete_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConnection(ctx, input)
			},
		},
		"delete-observability-configuration": {
			Name:   "delete-observability-configuration",
			Fields: fields_delete_observability_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteObservabilityConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_observability_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteObservabilityConfiguration(ctx, input)
			},
		},
		"delete-service": {
			Name:   "delete-service",
			Fields: fields_delete_service,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteServiceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_service, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteService(ctx, input)
			},
		},
		"delete-vpc-connector": {
			Name:   "delete-vpc-connector",
			Fields: fields_delete_vpc_connector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVpcConnectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vpc_connector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVpcConnector(ctx, input)
			},
		},
		"delete-vpc-ingress-connection": {
			Name:   "delete-vpc-ingress-connection",
			Fields: fields_delete_vpc_ingress_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVpcIngressConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vpc_ingress_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVpcIngressConnection(ctx, input)
			},
		},
		"describe-auto-scaling-configuration": {
			Name:   "describe-auto-scaling-configuration",
			Fields: fields_describe_auto_scaling_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAutoScalingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_auto_scaling_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAutoScalingConfiguration(ctx, input)
			},
		},
		"describe-custom-domains": {
			Name:   "describe-custom-domains",
			Fields: fields_describe_custom_domains,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCustomDomainsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_custom_domains, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeCustomDomains(ctx, input)
				}
				var results []*svc.DescribeCustomDomainsOutput
				p := svc.NewDescribeCustomDomainsPaginator(client, input)
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
		"describe-observability-configuration": {
			Name:   "describe-observability-configuration",
			Fields: fields_describe_observability_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeObservabilityConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_observability_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeObservabilityConfiguration(ctx, input)
			},
		},
		"describe-service": {
			Name:   "describe-service",
			Fields: fields_describe_service,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeServiceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_service, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeService(ctx, input)
			},
		},
		"describe-vpc-connector": {
			Name:   "describe-vpc-connector",
			Fields: fields_describe_vpc_connector,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVpcConnectorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_vpc_connector, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeVpcConnector(ctx, input)
			},
		},
		"describe-vpc-ingress-connection": {
			Name:   "describe-vpc-ingress-connection",
			Fields: fields_describe_vpc_ingress_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVpcIngressConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_vpc_ingress_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeVpcIngressConnection(ctx, input)
			},
		},
		"disassociate-custom-domain": {
			Name:   "disassociate-custom-domain",
			Fields: fields_disassociate_custom_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateCustomDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_custom_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateCustomDomain(ctx, input)
			},
		},
		"list-auto-scaling-configurations": {
			Name:   "list-auto-scaling-configurations",
			Fields: fields_list_auto_scaling_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAutoScalingConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_auto_scaling_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAutoScalingConfigurations(ctx, input)
				}
				var results []*svc.ListAutoScalingConfigurationsOutput
				p := svc.NewListAutoScalingConfigurationsPaginator(client, input)
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
		"list-connections": {
			Name:   "list-connections",
			Fields: fields_list_connections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConnectionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_connections, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConnections(ctx, input)
				}
				var results []*svc.ListConnectionsOutput
				p := svc.NewListConnectionsPaginator(client, input)
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
		"list-observability-configurations": {
			Name:   "list-observability-configurations",
			Fields: fields_list_observability_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListObservabilityConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_observability_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListObservabilityConfigurations(ctx, input)
				}
				var results []*svc.ListObservabilityConfigurationsOutput
				p := svc.NewListObservabilityConfigurationsPaginator(client, input)
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
		"list-operations": {
			Name:   "list-operations",
			Fields: fields_list_operations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOperationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_operations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOperations(ctx, input)
				}
				var results []*svc.ListOperationsOutput
				p := svc.NewListOperationsPaginator(client, input)
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
		"list-services": {
			Name:   "list-services",
			Fields: fields_list_services,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServicesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_services, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServices(ctx, input)
				}
				var results []*svc.ListServicesOutput
				p := svc.NewListServicesPaginator(client, input)
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
		"list-services-for-auto-scaling-configuration": {
			Name:   "list-services-for-auto-scaling-configuration",
			Fields: fields_list_services_for_auto_scaling_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServicesForAutoScalingConfigurationInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_services_for_auto_scaling_configuration, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServicesForAutoScalingConfiguration(ctx, input)
				}
				var results []*svc.ListServicesForAutoScalingConfigurationOutput
				p := svc.NewListServicesForAutoScalingConfigurationPaginator(client, input)
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
		"list-vpc-connectors": {
			Name:   "list-vpc-connectors",
			Fields: fields_list_vpc_connectors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVpcConnectorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_vpc_connectors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListVpcConnectors(ctx, input)
				}
				var results []*svc.ListVpcConnectorsOutput
				p := svc.NewListVpcConnectorsPaginator(client, input)
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
		"list-vpc-ingress-connections": {
			Name:   "list-vpc-ingress-connections",
			Fields: fields_list_vpc_ingress_connections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVpcIngressConnectionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_vpc_ingress_connections, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListVpcIngressConnections(ctx, input)
				}
				var results []*svc.ListVpcIngressConnectionsOutput
				p := svc.NewListVpcIngressConnectionsPaginator(client, input)
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
		"pause-service": {
			Name:   "pause-service",
			Fields: fields_pause_service,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PauseServiceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_pause_service, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PauseService(ctx, input)
			},
		},
		"resume-service": {
			Name:   "resume-service",
			Fields: fields_resume_service,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResumeServiceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_resume_service, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResumeService(ctx, input)
			},
		},
		"start-deployment": {
			Name:   "start-deployment",
			Fields: fields_start_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDeployment(ctx, input)
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
		"update-default-auto-scaling-configuration": {
			Name:   "update-default-auto-scaling-configuration",
			Fields: fields_update_default_auto_scaling_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDefaultAutoScalingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_default_auto_scaling_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDefaultAutoScalingConfiguration(ctx, input)
			},
		},
		"update-service": {
			Name:   "update-service",
			Fields: fields_update_service,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateServiceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_service, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateService(ctx, input)
			},
		},
		"update-vpc-ingress-connection": {
			Name:   "update-vpc-ingress-connection",
			Fields: fields_update_vpc_ingress_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateVpcIngressConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_vpc_ingress_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateVpcIngressConnection(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("apprunner", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
