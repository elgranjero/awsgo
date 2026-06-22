package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/servicediscovery"
)

var fields_create_http_namespace = []leanruntime.Field{
	{Name: "CreatorRequestId", Flag: "creator-request-id", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_private_dns_namespace = []leanruntime.Field{
	{Name: "CreatorRequestId", Flag: "creator-request-id", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Properties", Flag: "properties", Type: "*types.PrivateDnsNamespaceProperties", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Vpc", Flag: "vpc", Type: "*string", Required: true},
}

var fields_create_public_dns_namespace = []leanruntime.Field{
	{Name: "CreatorRequestId", Flag: "creator-request-id", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Properties", Flag: "properties", Type: "*types.PublicDnsNamespaceProperties", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_service = []leanruntime.Field{
	{Name: "CreatorRequestId", Flag: "creator-request-id", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DnsConfig", Flag: "dns-config", Type: "*types.DnsConfig", Required: false},
	{Name: "HealthCheckConfig", Flag: "health-check-config", Type: "*types.HealthCheckConfig", Required: false},
	{Name: "HealthCheckCustomConfig", Flag: "health-check-custom-config", Type: "*types.HealthCheckCustomConfig", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NamespaceId", Flag: "namespace-id", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Type", Flag: "type", Type: "types.ServiceTypeOption", Required: false},
}

var fields_delete_namespace = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_service = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_service_attributes = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "[]string", Required: true},
	{Name: "ServiceId", Flag: "service-id", Type: "*string", Required: true},
}

var fields_deregister_instance = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "ServiceId", Flag: "service-id", Type: "*string", Required: true},
}

var fields_discover_instances = []leanruntime.Field{
	{Name: "HealthStatus", Flag: "health-status", Type: "types.HealthStatusFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NamespaceName", Flag: "namespace-name", Type: "*string", Required: true},
	{Name: "OptionalParameters", Flag: "optional-parameters", Type: "map[string]string", Required: false},
	{Name: "OwnerAccount", Flag: "owner-account", Type: "*string", Required: false},
	{Name: "QueryParameters", Flag: "query-parameters", Type: "map[string]string", Required: false},
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: true},
}

var fields_discover_instances_revision = []leanruntime.Field{
	{Name: "NamespaceName", Flag: "namespace-name", Type: "*string", Required: true},
	{Name: "OwnerAccount", Flag: "owner-account", Type: "*string", Required: false},
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: true},
}

var fields_get_instance = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "ServiceId", Flag: "service-id", Type: "*string", Required: true},
}

var fields_get_instances_health_status = []leanruntime.Field{
	{Name: "Instances", Flag: "instances", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServiceId", Flag: "service-id", Type: "*string", Required: true},
}

var fields_get_namespace = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_operation = []leanruntime.Field{
	{Name: "OperationId", Flag: "operation-id", Type: "*string", Required: true},
	{Name: "OwnerAccount", Flag: "owner-account", Type: "*string", Required: false},
}

var fields_get_service = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_service_attributes = []leanruntime.Field{
	{Name: "ServiceId", Flag: "service-id", Type: "*string", Required: true},
}

var fields_list_instances = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServiceId", Flag: "service-id", Type: "*string", Required: true},
}

var fields_list_namespaces = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.NamespaceFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_operations = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.OperationFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_services = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.ServiceFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_register_instance = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "map[string]string", Required: true},
	{Name: "CreatorRequestId", Flag: "creator-request-id", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "ServiceId", Flag: "service-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_http_namespace = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*types.HttpNamespaceChange", Required: true},
	{Name: "UpdaterRequestId", Flag: "updater-request-id", Type: "*string", Required: false},
}

var fields_update_instance_custom_health_status = []leanruntime.Field{
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "ServiceId", Flag: "service-id", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.CustomHealthStatus", Required: true},
}

var fields_update_private_dns_namespace = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*types.PrivateDnsNamespaceChange", Required: true},
	{Name: "UpdaterRequestId", Flag: "updater-request-id", Type: "*string", Required: false},
}

var fields_update_public_dns_namespace = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Namespace", Flag: "namespace", Type: "*types.PublicDnsNamespaceChange", Required: true},
	{Name: "UpdaterRequestId", Flag: "updater-request-id", Type: "*string", Required: false},
}

var fields_update_service = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Service", Flag: "service", Type: "*types.ServiceChange", Required: true},
}

var fields_update_service_attributes = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "map[string]string", Required: true},
	{Name: "ServiceId", Flag: "service-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-http-namespace": {
			Name:   "create-http-namespace",
			Fields: fields_create_http_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateHttpNamespaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_http_namespace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateHttpNamespace(ctx, input)
			},
		},
		"create-private-dns-namespace": {
			Name:   "create-private-dns-namespace",
			Fields: fields_create_private_dns_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePrivateDnsNamespaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_private_dns_namespace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePrivateDnsNamespace(ctx, input)
			},
		},
		"create-public-dns-namespace": {
			Name:   "create-public-dns-namespace",
			Fields: fields_create_public_dns_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePublicDnsNamespaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_public_dns_namespace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePublicDnsNamespace(ctx, input)
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
		"delete-service-attributes": {
			Name:   "delete-service-attributes",
			Fields: fields_delete_service_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteServiceAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_service_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteServiceAttributes(ctx, input)
			},
		},
		"deregister-instance": {
			Name:   "deregister-instance",
			Fields: fields_deregister_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterInstance(ctx, input)
			},
		},
		"discover-instances": {
			Name:   "discover-instances",
			Fields: fields_discover_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DiscoverInstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_discover_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DiscoverInstances(ctx, input)
			},
		},
		"discover-instances-revision": {
			Name:   "discover-instances-revision",
			Fields: fields_discover_instances_revision,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DiscoverInstancesRevisionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_discover_instances_revision, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DiscoverInstancesRevision(ctx, input)
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
		"get-instances-health-status": {
			Name:   "get-instances-health-status",
			Fields: fields_get_instances_health_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInstancesHealthStatusInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_instances_health_status, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetInstancesHealthStatus(ctx, input)
				}
				var results []*svc.GetInstancesHealthStatusOutput
				p := svc.NewGetInstancesHealthStatusPaginator(client, input)
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
		"get-operation": {
			Name:   "get-operation",
			Fields: fields_get_operation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOperationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_operation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOperation(ctx, input)
			},
		},
		"get-service": {
			Name:   "get-service",
			Fields: fields_get_service,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetServiceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_service, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetService(ctx, input)
			},
		},
		"get-service-attributes": {
			Name:   "get-service-attributes",
			Fields: fields_get_service_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetServiceAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_service_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetServiceAttributes(ctx, input)
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
		"register-instance": {
			Name:   "register-instance",
			Fields: fields_register_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterInstance(ctx, input)
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
		"update-http-namespace": {
			Name:   "update-http-namespace",
			Fields: fields_update_http_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateHttpNamespaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_http_namespace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateHttpNamespace(ctx, input)
			},
		},
		"update-instance-custom-health-status": {
			Name:   "update-instance-custom-health-status",
			Fields: fields_update_instance_custom_health_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateInstanceCustomHealthStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_instance_custom_health_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateInstanceCustomHealthStatus(ctx, input)
			},
		},
		"update-private-dns-namespace": {
			Name:   "update-private-dns-namespace",
			Fields: fields_update_private_dns_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePrivateDnsNamespaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_private_dns_namespace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePrivateDnsNamespace(ctx, input)
			},
		},
		"update-public-dns-namespace": {
			Name:   "update-public-dns-namespace",
			Fields: fields_update_public_dns_namespace,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePublicDnsNamespaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_public_dns_namespace, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePublicDnsNamespace(ctx, input)
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
		"update-service-attributes": {
			Name:   "update-service-attributes",
			Fields: fields_update_service_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateServiceAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_service_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateServiceAttributes(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("servicediscovery", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
