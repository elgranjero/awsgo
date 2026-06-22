package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/migrationhubconfig"
)

var fields_create_home_region_control = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "bool", Required: false},
	{Name: "HomeRegion", Flag: "home-region", Type: "*string", Required: true},
	{Name: "Target", Flag: "target", Type: "*types.Target", Required: true},
}

var fields_delete_home_region_control = []leanruntime.Field{
	{Name: "ControlId", Flag: "control-id", Type: "*string", Required: true},
}

var fields_describe_home_region_controls = []leanruntime.Field{
	{Name: "ControlId", Flag: "control-id", Type: "*string", Required: false},
	{Name: "HomeRegion", Flag: "home-region", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Target", Flag: "target", Type: "*types.Target", Required: false},
}

var fields_get_home_region = []leanruntime.Field{}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-home-region-control": {
			Name:   "create-home-region-control",
			Fields: fields_create_home_region_control,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateHomeRegionControlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_home_region_control, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateHomeRegionControl(ctx, input)
			},
		},
		"delete-home-region-control": {
			Name:   "delete-home-region-control",
			Fields: fields_delete_home_region_control,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteHomeRegionControlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_home_region_control, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteHomeRegionControl(ctx, input)
			},
		},
		"describe-home-region-controls": {
			Name:   "describe-home-region-controls",
			Fields: fields_describe_home_region_controls,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeHomeRegionControlsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_home_region_controls, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeHomeRegionControls(ctx, input)
				}
				var results []*svc.DescribeHomeRegionControlsOutput
				p := svc.NewDescribeHomeRegionControlsPaginator(client, input)
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
		"get-home-region": {
			Name:   "get-home-region",
			Fields: fields_get_home_region,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetHomeRegionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_home_region, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetHomeRegion(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("migrationhubconfig", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
