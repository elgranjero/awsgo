package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/simspaceweaver"
)

var fields_create_snapshot = []leanruntime.Field{
	{Name: "Destination", Flag: "destination", Type: "*types.S3Destination", Required: true},
	{Name: "Simulation", Flag: "simulation", Type: "*string", Required: true},
}

var fields_delete_app = []leanruntime.Field{
	{Name: "App", Flag: "app", Type: "*string", Required: true},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "Simulation", Flag: "simulation", Type: "*string", Required: true},
}

var fields_delete_simulation = []leanruntime.Field{
	{Name: "Simulation", Flag: "simulation", Type: "*string", Required: true},
}

var fields_describe_app = []leanruntime.Field{
	{Name: "App", Flag: "app", Type: "*string", Required: true},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "Simulation", Flag: "simulation", Type: "*string", Required: true},
}

var fields_describe_simulation = []leanruntime.Field{
	{Name: "Simulation", Flag: "simulation", Type: "*string", Required: true},
}

var fields_list_apps = []leanruntime.Field{
	{Name: "Domain", Flag: "domain", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Simulation", Flag: "simulation", Type: "*string", Required: true},
}

var fields_list_simulations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_start_app = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "LaunchOverrides", Flag: "launch-overrides", Type: "*types.LaunchOverrides", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Simulation", Flag: "simulation", Type: "*string", Required: true},
}

var fields_start_clock = []leanruntime.Field{
	{Name: "Simulation", Flag: "simulation", Type: "*string", Required: true},
}

var fields_start_simulation = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "MaximumDuration", Flag: "maximum-duration", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "SchemaS3Location", Flag: "schema-s3-location", Type: "*types.S3Location", Required: false},
	{Name: "SnapshotS3Location", Flag: "snapshot-s3-location", Type: "*types.S3Location", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_stop_app = []leanruntime.Field{
	{Name: "App", Flag: "app", Type: "*string", Required: true},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: true},
	{Name: "Simulation", Flag: "simulation", Type: "*string", Required: true},
}

var fields_stop_clock = []leanruntime.Field{
	{Name: "Simulation", Flag: "simulation", Type: "*string", Required: true},
}

var fields_stop_simulation = []leanruntime.Field{
	{Name: "Simulation", Flag: "simulation", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-snapshot": {
			Name:   "create-snapshot",
			Fields: fields_create_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSnapshot(ctx, input)
			},
		},
		"delete-app": {
			Name:   "delete-app",
			Fields: fields_delete_app,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAppInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_app, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApp(ctx, input)
			},
		},
		"delete-simulation": {
			Name:   "delete-simulation",
			Fields: fields_delete_simulation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSimulationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_simulation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSimulation(ctx, input)
			},
		},
		"describe-app": {
			Name:   "describe-app",
			Fields: fields_describe_app,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAppInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_app, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeApp(ctx, input)
			},
		},
		"describe-simulation": {
			Name:   "describe-simulation",
			Fields: fields_describe_simulation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSimulationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_simulation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSimulation(ctx, input)
			},
		},
		"list-apps": {
			Name:   "list-apps",
			Fields: fields_list_apps,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAppsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_apps, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApps(ctx, input)
				}
				var results []*svc.ListAppsOutput
				p := svc.NewListAppsPaginator(client, input)
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
		"list-simulations": {
			Name:   "list-simulations",
			Fields: fields_list_simulations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSimulationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_simulations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSimulations(ctx, input)
				}
				var results []*svc.ListSimulationsOutput
				p := svc.NewListSimulationsPaginator(client, input)
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
		"start-app": {
			Name:   "start-app",
			Fields: fields_start_app,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartAppInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_app, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartApp(ctx, input)
			},
		},
		"start-clock": {
			Name:   "start-clock",
			Fields: fields_start_clock,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartClockInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_clock, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartClock(ctx, input)
			},
		},
		"start-simulation": {
			Name:   "start-simulation",
			Fields: fields_start_simulation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartSimulationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_simulation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartSimulation(ctx, input)
			},
		},
		"stop-app": {
			Name:   "stop-app",
			Fields: fields_stop_app,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopAppInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_app, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopApp(ctx, input)
			},
		},
		"stop-clock": {
			Name:   "stop-clock",
			Fields: fields_stop_clock,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopClockInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_clock, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopClock(ctx, input)
			},
		},
		"stop-simulation": {
			Name:   "stop-simulation",
			Fields: fields_stop_simulation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopSimulationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_simulation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopSimulation(ctx, input)
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
	if err := leanruntime.Execute("simspaceweaver", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
