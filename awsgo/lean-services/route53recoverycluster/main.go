package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/route53recoverycluster"
)

var fields_get_routing_control_state = []leanruntime.Field{
	{Name: "RoutingControlArn", Flag: "routing-control-arn", Type: "*string", Required: true},
}

var fields_list_routing_controls = []leanruntime.Field{
	{Name: "ControlPanelArn", Flag: "control-panel-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_update_routing_control_state = []leanruntime.Field{
	{Name: "RoutingControlArn", Flag: "routing-control-arn", Type: "*string", Required: true},
	{Name: "RoutingControlState", Flag: "routing-control-state", Type: "types.RoutingControlState", Required: true},
	{Name: "SafetyRulesToOverride", Flag: "safety-rules-to-override", Type: "[]string", Required: false},
}

var fields_update_routing_control_states = []leanruntime.Field{
	{Name: "SafetyRulesToOverride", Flag: "safety-rules-to-override", Type: "[]string", Required: false},
	{Name: "UpdateRoutingControlStateEntries", Flag: "update-routing-control-state-entries", Type: "[]types.UpdateRoutingControlStateEntry", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"get-routing-control-state": {
			Name:   "get-routing-control-state",
			Fields: fields_get_routing_control_state,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRoutingControlStateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_routing_control_state, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRoutingControlState(ctx, input)
			},
		},
		"list-routing-controls": {
			Name:   "list-routing-controls",
			Fields: fields_list_routing_controls,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRoutingControlsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_routing_controls, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRoutingControls(ctx, input)
				}
				var results []*svc.ListRoutingControlsOutput
				p := svc.NewListRoutingControlsPaginator(client, input)
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
		"update-routing-control-state": {
			Name:   "update-routing-control-state",
			Fields: fields_update_routing_control_state,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRoutingControlStateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_routing_control_state, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRoutingControlState(ctx, input)
			},
		},
		"update-routing-control-states": {
			Name:   "update-routing-control-states",
			Fields: fields_update_routing_control_states,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRoutingControlStatesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_routing_control_states, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRoutingControlStates(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("route53recoverycluster", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
