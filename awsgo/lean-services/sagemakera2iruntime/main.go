package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/sagemakera2iruntime"
)

var fields_delete_human_loop = []leanruntime.Field{
	{Name: "HumanLoopName", Flag: "human-loop-name", Type: "*string", Required: true},
}

var fields_describe_human_loop = []leanruntime.Field{
	{Name: "HumanLoopName", Flag: "human-loop-name", Type: "*string", Required: true},
}

var fields_list_human_loops = []leanruntime.Field{
	{Name: "CreationTimeAfter", Flag: "creation-time-after", Type: "*time.Time", Required: false},
	{Name: "CreationTimeBefore", Flag: "creation-time-before", Type: "*time.Time", Required: false},
	{Name: "FlowDefinitionArn", Flag: "flow-definition-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_start_human_loop = []leanruntime.Field{
	{Name: "DataAttributes", Flag: "data-attributes", Type: "*types.HumanLoopDataAttributes", Required: false},
	{Name: "FlowDefinitionArn", Flag: "flow-definition-arn", Type: "*string", Required: true},
	{Name: "HumanLoopInput", Flag: "human-loop-input", Type: "*types.HumanLoopInput", Required: true},
	{Name: "HumanLoopName", Flag: "human-loop-name", Type: "*string", Required: true},
}

var fields_stop_human_loop = []leanruntime.Field{
	{Name: "HumanLoopName", Flag: "human-loop-name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"delete-human-loop": {
			Name:   "delete-human-loop",
			Fields: fields_delete_human_loop,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteHumanLoopInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_human_loop, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteHumanLoop(ctx, input)
			},
		},
		"describe-human-loop": {
			Name:   "describe-human-loop",
			Fields: fields_describe_human_loop,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeHumanLoopInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_human_loop, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeHumanLoop(ctx, input)
			},
		},
		"list-human-loops": {
			Name:   "list-human-loops",
			Fields: fields_list_human_loops,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListHumanLoopsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_human_loops, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListHumanLoops(ctx, input)
				}
				var results []*svc.ListHumanLoopsOutput
				p := svc.NewListHumanLoopsPaginator(client, input)
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
		"start-human-loop": {
			Name:   "start-human-loop",
			Fields: fields_start_human_loop,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartHumanLoopInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_human_loop, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartHumanLoop(ctx, input)
			},
		},
		"stop-human-loop": {
			Name:   "stop-human-loop",
			Fields: fields_stop_human_loop,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopHumanLoopInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_human_loop, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopHumanLoop(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("sagemakera2iruntime", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
