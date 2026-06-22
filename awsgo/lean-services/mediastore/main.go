package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/mediastore"
)

var fields_create_container = []leanruntime.Field{
	{Name: "ContainerName", Flag: "container-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_container = []leanruntime.Field{
	{Name: "ContainerName", Flag: "container-name", Type: "*string", Required: true},
}

var fields_delete_container_policy = []leanruntime.Field{
	{Name: "ContainerName", Flag: "container-name", Type: "*string", Required: true},
}

var fields_delete_cors_policy = []leanruntime.Field{
	{Name: "ContainerName", Flag: "container-name", Type: "*string", Required: true},
}

var fields_delete_lifecycle_policy = []leanruntime.Field{
	{Name: "ContainerName", Flag: "container-name", Type: "*string", Required: true},
}

var fields_delete_metric_policy = []leanruntime.Field{
	{Name: "ContainerName", Flag: "container-name", Type: "*string", Required: true},
}

var fields_describe_container = []leanruntime.Field{
	{Name: "ContainerName", Flag: "container-name", Type: "*string", Required: false},
}

var fields_get_container_policy = []leanruntime.Field{
	{Name: "ContainerName", Flag: "container-name", Type: "*string", Required: true},
}

var fields_get_cors_policy = []leanruntime.Field{
	{Name: "ContainerName", Flag: "container-name", Type: "*string", Required: true},
}

var fields_get_lifecycle_policy = []leanruntime.Field{
	{Name: "ContainerName", Flag: "container-name", Type: "*string", Required: true},
}

var fields_get_metric_policy = []leanruntime.Field{
	{Name: "ContainerName", Flag: "container-name", Type: "*string", Required: true},
}

var fields_list_containers = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "Resource", Flag: "resource", Type: "*string", Required: true},
}

var fields_put_container_policy = []leanruntime.Field{
	{Name: "ContainerName", Flag: "container-name", Type: "*string", Required: true},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
}

var fields_put_cors_policy = []leanruntime.Field{
	{Name: "ContainerName", Flag: "container-name", Type: "*string", Required: true},
	{Name: "CorsPolicy", Flag: "cors-policy", Type: "[]types.CorsRule", Required: true},
}

var fields_put_lifecycle_policy = []leanruntime.Field{
	{Name: "ContainerName", Flag: "container-name", Type: "*string", Required: true},
	{Name: "LifecyclePolicy", Flag: "lifecycle-policy", Type: "*string", Required: true},
}

var fields_put_metric_policy = []leanruntime.Field{
	{Name: "ContainerName", Flag: "container-name", Type: "*string", Required: true},
	{Name: "MetricPolicy", Flag: "metric-policy", Type: "*types.MetricPolicy", Required: true},
}

var fields_start_access_logging = []leanruntime.Field{
	{Name: "ContainerName", Flag: "container-name", Type: "*string", Required: true},
}

var fields_stop_access_logging = []leanruntime.Field{
	{Name: "ContainerName", Flag: "container-name", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "Resource", Flag: "resource", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "Resource", Flag: "resource", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-container": {
			Name:   "create-container",
			Fields: fields_create_container,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateContainerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_container, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateContainer(ctx, input)
			},
		},
		"delete-container": {
			Name:   "delete-container",
			Fields: fields_delete_container,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteContainerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_container, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteContainer(ctx, input)
			},
		},
		"delete-container-policy": {
			Name:   "delete-container-policy",
			Fields: fields_delete_container_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteContainerPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_container_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteContainerPolicy(ctx, input)
			},
		},
		"delete-cors-policy": {
			Name:   "delete-cors-policy",
			Fields: fields_delete_cors_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCorsPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_cors_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCorsPolicy(ctx, input)
			},
		},
		"delete-lifecycle-policy": {
			Name:   "delete-lifecycle-policy",
			Fields: fields_delete_lifecycle_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLifecyclePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_lifecycle_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLifecyclePolicy(ctx, input)
			},
		},
		"delete-metric-policy": {
			Name:   "delete-metric-policy",
			Fields: fields_delete_metric_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMetricPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_metric_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMetricPolicy(ctx, input)
			},
		},
		"describe-container": {
			Name:   "describe-container",
			Fields: fields_describe_container,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeContainerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_container, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeContainer(ctx, input)
			},
		},
		"get-container-policy": {
			Name:   "get-container-policy",
			Fields: fields_get_container_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetContainerPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_container_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetContainerPolicy(ctx, input)
			},
		},
		"get-cors-policy": {
			Name:   "get-cors-policy",
			Fields: fields_get_cors_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCorsPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_cors_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCorsPolicy(ctx, input)
			},
		},
		"get-lifecycle-policy": {
			Name:   "get-lifecycle-policy",
			Fields: fields_get_lifecycle_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLifecyclePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_lifecycle_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLifecyclePolicy(ctx, input)
			},
		},
		"get-metric-policy": {
			Name:   "get-metric-policy",
			Fields: fields_get_metric_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMetricPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_metric_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMetricPolicy(ctx, input)
			},
		},
		"list-containers": {
			Name:   "list-containers",
			Fields: fields_list_containers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListContainersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_containers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListContainers(ctx, input)
				}
				var results []*svc.ListContainersOutput
				p := svc.NewListContainersPaginator(client, input)
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
		"put-container-policy": {
			Name:   "put-container-policy",
			Fields: fields_put_container_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutContainerPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_container_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutContainerPolicy(ctx, input)
			},
		},
		"put-cors-policy": {
			Name:   "put-cors-policy",
			Fields: fields_put_cors_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutCorsPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_cors_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutCorsPolicy(ctx, input)
			},
		},
		"put-lifecycle-policy": {
			Name:   "put-lifecycle-policy",
			Fields: fields_put_lifecycle_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutLifecyclePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_lifecycle_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutLifecyclePolicy(ctx, input)
			},
		},
		"put-metric-policy": {
			Name:   "put-metric-policy",
			Fields: fields_put_metric_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutMetricPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_metric_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutMetricPolicy(ctx, input)
			},
		},
		"start-access-logging": {
			Name:   "start-access-logging",
			Fields: fields_start_access_logging,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartAccessLoggingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_access_logging, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartAccessLogging(ctx, input)
			},
		},
		"stop-access-logging": {
			Name:   "stop-access-logging",
			Fields: fields_stop_access_logging,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopAccessLoggingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_access_logging, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopAccessLogging(ctx, input)
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
	}
	if err := leanruntime.Execute("mediastore", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
