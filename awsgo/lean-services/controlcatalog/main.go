package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/controlcatalog"
)

var fields_get_control = []leanruntime.Field{
	{Name: "ControlArn", Flag: "control-arn", Type: "*string", Required: true},
}

var fields_list_common_controls = []leanruntime.Field{
	{Name: "CommonControlFilter", Flag: "common-control-filter", Type: "*types.CommonControlFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_control_mappings = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.ControlMappingFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_controls = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.ControlFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_domains = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_objectives = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ObjectiveFilter", Flag: "objective-filter", Type: "*types.ObjectiveFilter", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"get-control": {
			Name:   "get-control",
			Fields: fields_get_control,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetControlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_control, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetControl(ctx, input)
			},
		},
		"list-common-controls": {
			Name:   "list-common-controls",
			Fields: fields_list_common_controls,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCommonControlsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_common_controls, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCommonControls(ctx, input)
				}
				var results []*svc.ListCommonControlsOutput
				p := svc.NewListCommonControlsPaginator(client, input)
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
		"list-control-mappings": {
			Name:   "list-control-mappings",
			Fields: fields_list_control_mappings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListControlMappingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_control_mappings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListControlMappings(ctx, input)
				}
				var results []*svc.ListControlMappingsOutput
				p := svc.NewListControlMappingsPaginator(client, input)
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
		"list-controls": {
			Name:   "list-controls",
			Fields: fields_list_controls,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListControlsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_controls, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListControls(ctx, input)
				}
				var results []*svc.ListControlsOutput
				p := svc.NewListControlsPaginator(client, input)
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
		"list-domains": {
			Name:   "list-domains",
			Fields: fields_list_domains,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDomainsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_domains, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDomains(ctx, input)
				}
				var results []*svc.ListDomainsOutput
				p := svc.NewListDomainsPaginator(client, input)
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
		"list-objectives": {
			Name:   "list-objectives",
			Fields: fields_list_objectives,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListObjectivesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_objectives, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListObjectives(ctx, input)
				}
				var results []*svc.ListObjectivesOutput
				p := svc.NewListObjectivesPaginator(client, input)
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
	}
	if err := leanruntime.Execute("controlcatalog", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
