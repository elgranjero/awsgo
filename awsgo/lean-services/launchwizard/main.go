package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/launchwizard"
)

var fields_create_deployment = []leanruntime.Field{
	{Name: "DeploymentPatternName", Flag: "deployment-pattern-name", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Specifications", Flag: "specifications", Type: "map[string]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "WorkloadName", Flag: "workload-name", Type: "*string", Required: true},
}

var fields_delete_deployment = []leanruntime.Field{
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: true},
}

var fields_get_deployment = []leanruntime.Field{
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: true},
}

var fields_get_deployment_pattern_version = []leanruntime.Field{
	{Name: "DeploymentPatternName", Flag: "deployment-pattern-name", Type: "*string", Required: true},
	{Name: "DeploymentPatternVersionName", Flag: "deployment-pattern-version-name", Type: "*string", Required: true},
	{Name: "WorkloadName", Flag: "workload-name", Type: "*string", Required: true},
}

var fields_get_workload = []leanruntime.Field{
	{Name: "WorkloadName", Flag: "workload-name", Type: "*string", Required: true},
}

var fields_get_workload_deployment_pattern = []leanruntime.Field{
	{Name: "DeploymentPatternName", Flag: "deployment-pattern-name", Type: "*string", Required: true},
	{Name: "WorkloadName", Flag: "workload-name", Type: "*string", Required: true},
}

var fields_list_deployment_events = []leanruntime.Field{
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_deployment_pattern_versions = []leanruntime.Field{
	{Name: "DeploymentPatternName", Flag: "deployment-pattern-name", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.DeploymentPatternVersionFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkloadName", Flag: "workload-name", Type: "*string", Required: true},
}

var fields_list_deployments = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.DeploymentFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_workload_deployment_patterns = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkloadName", Flag: "workload-name", Type: "*string", Required: true},
}

var fields_list_workloads = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
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

var fields_update_deployment = []leanruntime.Field{
	{Name: "DeploymentId", Flag: "deployment-id", Type: "*string", Required: true},
	{Name: "DeploymentPatternVersionName", Flag: "deployment-pattern-version-name", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "bool", Required: false},
	{Name: "Force", Flag: "force", Type: "bool", Required: false},
	{Name: "Specifications", Flag: "specifications", Type: "map[string]string", Required: true},
	{Name: "WorkloadVersionName", Flag: "workload-version-name", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-deployment": {
			Name:   "create-deployment",
			Fields: fields_create_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDeployment(ctx, input)
			},
		},
		"delete-deployment": {
			Name:   "delete-deployment",
			Fields: fields_delete_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDeployment(ctx, input)
			},
		},
		"get-deployment": {
			Name:   "get-deployment",
			Fields: fields_get_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDeployment(ctx, input)
			},
		},
		"get-deployment-pattern-version": {
			Name:   "get-deployment-pattern-version",
			Fields: fields_get_deployment_pattern_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDeploymentPatternVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_deployment_pattern_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDeploymentPatternVersion(ctx, input)
			},
		},
		"get-workload": {
			Name:   "get-workload",
			Fields: fields_get_workload,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWorkloadInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_workload, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWorkload(ctx, input)
			},
		},
		"get-workload-deployment-pattern": {
			Name:   "get-workload-deployment-pattern",
			Fields: fields_get_workload_deployment_pattern,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWorkloadDeploymentPatternInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_workload_deployment_pattern, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWorkloadDeploymentPattern(ctx, input)
			},
		},
		"list-deployment-events": {
			Name:   "list-deployment-events",
			Fields: fields_list_deployment_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDeploymentEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_deployment_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDeploymentEvents(ctx, input)
				}
				var results []*svc.ListDeploymentEventsOutput
				p := svc.NewListDeploymentEventsPaginator(client, input)
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
		"list-deployment-pattern-versions": {
			Name:   "list-deployment-pattern-versions",
			Fields: fields_list_deployment_pattern_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDeploymentPatternVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_deployment_pattern_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDeploymentPatternVersions(ctx, input)
				}
				var results []*svc.ListDeploymentPatternVersionsOutput
				p := svc.NewListDeploymentPatternVersionsPaginator(client, input)
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
		"list-deployments": {
			Name:   "list-deployments",
			Fields: fields_list_deployments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDeploymentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_deployments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDeployments(ctx, input)
				}
				var results []*svc.ListDeploymentsOutput
				p := svc.NewListDeploymentsPaginator(client, input)
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
		"list-workload-deployment-patterns": {
			Name:   "list-workload-deployment-patterns",
			Fields: fields_list_workload_deployment_patterns,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkloadDeploymentPatternsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_workload_deployment_patterns, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkloadDeploymentPatterns(ctx, input)
				}
				var results []*svc.ListWorkloadDeploymentPatternsOutput
				p := svc.NewListWorkloadDeploymentPatternsPaginator(client, input)
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
		"list-workloads": {
			Name:   "list-workloads",
			Fields: fields_list_workloads,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkloadsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_workloads, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkloads(ctx, input)
				}
				var results []*svc.ListWorkloadsOutput
				p := svc.NewListWorkloadsPaginator(client, input)
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
		"update-deployment": {
			Name:   "update-deployment",
			Fields: fields_update_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDeployment(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("launchwizard", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
