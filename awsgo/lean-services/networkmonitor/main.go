package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/networkmonitor"
)

var fields_create_monitor = []leanruntime.Field{
	{Name: "AggregationPeriod", Flag: "aggregation-period", Type: "*int64", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "MonitorName", Flag: "monitor-name", Type: "*string", Required: true},
	{Name: "Probes", Flag: "probes", Type: "[]types.CreateMonitorProbeInput", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_probe = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "MonitorName", Flag: "monitor-name", Type: "*string", Required: true},
	{Name: "Probe", Flag: "probe", Type: "*types.ProbeInput", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_monitor = []leanruntime.Field{
	{Name: "MonitorName", Flag: "monitor-name", Type: "*string", Required: true},
}

var fields_delete_probe = []leanruntime.Field{
	{Name: "MonitorName", Flag: "monitor-name", Type: "*string", Required: true},
	{Name: "ProbeId", Flag: "probe-id", Type: "*string", Required: true},
}

var fields_get_monitor = []leanruntime.Field{
	{Name: "MonitorName", Flag: "monitor-name", Type: "*string", Required: true},
}

var fields_get_probe = []leanruntime.Field{
	{Name: "MonitorName", Flag: "monitor-name", Type: "*string", Required: true},
	{Name: "ProbeId", Flag: "probe-id", Type: "*string", Required: true},
}

var fields_list_monitors = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "State", Flag: "state", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
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

var fields_update_monitor = []leanruntime.Field{
	{Name: "AggregationPeriod", Flag: "aggregation-period", Type: "*int64", Required: true},
	{Name: "MonitorName", Flag: "monitor-name", Type: "*string", Required: true},
}

var fields_update_probe = []leanruntime.Field{
	{Name: "Destination", Flag: "destination", Type: "*string", Required: false},
	{Name: "DestinationPort", Flag: "destination-port", Type: "*int32", Required: false},
	{Name: "MonitorName", Flag: "monitor-name", Type: "*string", Required: true},
	{Name: "PacketSize", Flag: "packet-size", Type: "*int32", Required: false},
	{Name: "ProbeId", Flag: "probe-id", Type: "*string", Required: true},
	{Name: "Protocol", Flag: "protocol", Type: "types.Protocol", Required: false},
	{Name: "State", Flag: "state", Type: "types.ProbeState", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-monitor": {
			Name:   "create-monitor",
			Fields: fields_create_monitor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMonitorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_monitor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMonitor(ctx, input)
			},
		},
		"create-probe": {
			Name:   "create-probe",
			Fields: fields_create_probe,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProbeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_probe, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProbe(ctx, input)
			},
		},
		"delete-monitor": {
			Name:   "delete-monitor",
			Fields: fields_delete_monitor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMonitorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_monitor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMonitor(ctx, input)
			},
		},
		"delete-probe": {
			Name:   "delete-probe",
			Fields: fields_delete_probe,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProbeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_probe, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProbe(ctx, input)
			},
		},
		"get-monitor": {
			Name:   "get-monitor",
			Fields: fields_get_monitor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMonitorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_monitor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMonitor(ctx, input)
			},
		},
		"get-probe": {
			Name:   "get-probe",
			Fields: fields_get_probe,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetProbeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_probe, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetProbe(ctx, input)
			},
		},
		"list-monitors": {
			Name:   "list-monitors",
			Fields: fields_list_monitors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMonitorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_monitors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMonitors(ctx, input)
				}
				var results []*svc.ListMonitorsOutput
				p := svc.NewListMonitorsPaginator(client, input)
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
		"update-monitor": {
			Name:   "update-monitor",
			Fields: fields_update_monitor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMonitorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_monitor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMonitor(ctx, input)
			},
		},
		"update-probe": {
			Name:   "update-probe",
			Fields: fields_update_probe,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProbeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_probe, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProbe(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("networkmonitor", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
