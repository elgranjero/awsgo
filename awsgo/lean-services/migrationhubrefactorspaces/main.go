package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/migrationhubrefactorspaces"
)

var fields_create_application = []leanruntime.Field{
	{Name: "ApiGatewayProxy", Flag: "api-gateway-proxy", Type: "*types.ApiGatewayProxyInput", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "EnvironmentIdentifier", Flag: "environment-identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ProxyType", Flag: "proxy-type", Type: "types.ProxyType", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_create_environment = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NetworkFabricType", Flag: "network-fabric-type", Type: "types.NetworkFabricType", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_route = []leanruntime.Field{
	{Name: "ApplicationIdentifier", Flag: "application-identifier", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DefaultRoute", Flag: "default-route", Type: "*types.DefaultRouteInput", Required: false},
	{Name: "EnvironmentIdentifier", Flag: "environment-identifier", Type: "*string", Required: true},
	{Name: "RouteType", Flag: "route-type", Type: "types.RouteType", Required: true},
	{Name: "ServiceIdentifier", Flag: "service-identifier", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "UriPathRoute", Flag: "uri-path-route", Type: "*types.UriPathRouteInput", Required: false},
}

var fields_create_service = []leanruntime.Field{
	{Name: "ApplicationIdentifier", Flag: "application-identifier", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EndpointType", Flag: "endpoint-type", Type: "types.ServiceEndpointType", Required: true},
	{Name: "EnvironmentIdentifier", Flag: "environment-identifier", Type: "*string", Required: true},
	{Name: "LambdaEndpoint", Flag: "lambda-endpoint", Type: "*types.LambdaEndpointInput", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "UrlEndpoint", Flag: "url-endpoint", Type: "*types.UrlEndpointInput", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: false},
}

var fields_delete_application = []leanruntime.Field{
	{Name: "ApplicationIdentifier", Flag: "application-identifier", Type: "*string", Required: true},
	{Name: "EnvironmentIdentifier", Flag: "environment-identifier", Type: "*string", Required: true},
}

var fields_delete_environment = []leanruntime.Field{
	{Name: "EnvironmentIdentifier", Flag: "environment-identifier", Type: "*string", Required: true},
}

var fields_delete_resource_policy = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_route = []leanruntime.Field{
	{Name: "ApplicationIdentifier", Flag: "application-identifier", Type: "*string", Required: true},
	{Name: "EnvironmentIdentifier", Flag: "environment-identifier", Type: "*string", Required: true},
	{Name: "RouteIdentifier", Flag: "route-identifier", Type: "*string", Required: true},
}

var fields_delete_service = []leanruntime.Field{
	{Name: "ApplicationIdentifier", Flag: "application-identifier", Type: "*string", Required: true},
	{Name: "EnvironmentIdentifier", Flag: "environment-identifier", Type: "*string", Required: true},
	{Name: "ServiceIdentifier", Flag: "service-identifier", Type: "*string", Required: true},
}

var fields_get_application = []leanruntime.Field{
	{Name: "ApplicationIdentifier", Flag: "application-identifier", Type: "*string", Required: true},
	{Name: "EnvironmentIdentifier", Flag: "environment-identifier", Type: "*string", Required: true},
}

var fields_get_environment = []leanruntime.Field{
	{Name: "EnvironmentIdentifier", Flag: "environment-identifier", Type: "*string", Required: true},
}

var fields_get_resource_policy = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_route = []leanruntime.Field{
	{Name: "ApplicationIdentifier", Flag: "application-identifier", Type: "*string", Required: true},
	{Name: "EnvironmentIdentifier", Flag: "environment-identifier", Type: "*string", Required: true},
	{Name: "RouteIdentifier", Flag: "route-identifier", Type: "*string", Required: true},
}

var fields_get_service = []leanruntime.Field{
	{Name: "ApplicationIdentifier", Flag: "application-identifier", Type: "*string", Required: true},
	{Name: "EnvironmentIdentifier", Flag: "environment-identifier", Type: "*string", Required: true},
	{Name: "ServiceIdentifier", Flag: "service-identifier", Type: "*string", Required: true},
}

var fields_list_applications = []leanruntime.Field{
	{Name: "EnvironmentIdentifier", Flag: "environment-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_environment_vpcs = []leanruntime.Field{
	{Name: "EnvironmentIdentifier", Flag: "environment-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_environments = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_routes = []leanruntime.Field{
	{Name: "ApplicationIdentifier", Flag: "application-identifier", Type: "*string", Required: true},
	{Name: "EnvironmentIdentifier", Flag: "environment-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_services = []leanruntime.Field{
	{Name: "ApplicationIdentifier", Flag: "application-identifier", Type: "*string", Required: true},
	{Name: "EnvironmentIdentifier", Flag: "environment-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_resource_policy = []leanruntime.Field{
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_route = []leanruntime.Field{
	{Name: "ActivationState", Flag: "activation-state", Type: "types.RouteActivationState", Required: true},
	{Name: "ApplicationIdentifier", Flag: "application-identifier", Type: "*string", Required: true},
	{Name: "EnvironmentIdentifier", Flag: "environment-identifier", Type: "*string", Required: true},
	{Name: "RouteIdentifier", Flag: "route-identifier", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
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
		"create-environment": {
			Name:   "create-environment",
			Fields: fields_create_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEnvironment(ctx, input)
			},
		},
		"create-route": {
			Name:   "create-route",
			Fields: fields_create_route,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRouteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_route, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRoute(ctx, input)
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
		"delete-application": {
			Name:   "delete-application",
			Fields: fields_delete_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApplication(ctx, input)
			},
		},
		"delete-environment": {
			Name:   "delete-environment",
			Fields: fields_delete_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEnvironment(ctx, input)
			},
		},
		"delete-resource-policy": {
			Name:   "delete-resource-policy",
			Fields: fields_delete_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResourcePolicy(ctx, input)
			},
		},
		"delete-route": {
			Name:   "delete-route",
			Fields: fields_delete_route,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRouteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_route, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRoute(ctx, input)
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
		"get-application": {
			Name:   "get-application",
			Fields: fields_get_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApplication(ctx, input)
			},
		},
		"get-environment": {
			Name:   "get-environment",
			Fields: fields_get_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEnvironment(ctx, input)
			},
		},
		"get-resource-policy": {
			Name:   "get-resource-policy",
			Fields: fields_get_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResourcePolicy(ctx, input)
			},
		},
		"get-route": {
			Name:   "get-route",
			Fields: fields_get_route,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRouteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_route, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRoute(ctx, input)
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
		"list-applications": {
			Name:   "list-applications",
			Fields: fields_list_applications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListApplicationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_applications, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApplications(ctx, input)
				}
				var results []*svc.ListApplicationsOutput
				p := svc.NewListApplicationsPaginator(client, input)
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
		"list-environment-vpcs": {
			Name:   "list-environment-vpcs",
			Fields: fields_list_environment_vpcs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEnvironmentVpcsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_environment_vpcs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEnvironmentVpcs(ctx, input)
				}
				var results []*svc.ListEnvironmentVpcsOutput
				p := svc.NewListEnvironmentVpcsPaginator(client, input)
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
		"list-environments": {
			Name:   "list-environments",
			Fields: fields_list_environments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEnvironmentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_environments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEnvironments(ctx, input)
				}
				var results []*svc.ListEnvironmentsOutput
				p := svc.NewListEnvironmentsPaginator(client, input)
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
		"list-routes": {
			Name:   "list-routes",
			Fields: fields_list_routes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRoutesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_routes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRoutes(ctx, input)
				}
				var results []*svc.ListRoutesOutput
				p := svc.NewListRoutesPaginator(client, input)
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
		"put-resource-policy": {
			Name:   "put-resource-policy",
			Fields: fields_put_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutResourcePolicy(ctx, input)
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
		"update-route": {
			Name:   "update-route",
			Fields: fields_update_route,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRouteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_route, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRoute(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("migrationhubrefactorspaces", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
