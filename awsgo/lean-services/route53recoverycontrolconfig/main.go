package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig"
)

var fields_create_cluster = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "NetworkType", Flag: "network-type", Type: "types.NetworkType", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_control_panel = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
	{Name: "ControlPanelName", Flag: "control-panel-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_routing_control = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
	{Name: "ControlPanelArn", Flag: "control-panel-arn", Type: "*string", Required: false},
	{Name: "RoutingControlName", Flag: "routing-control-name", Type: "*string", Required: true},
}

var fields_create_safety_rule = []leanruntime.Field{
	{Name: "AssertionRule", Flag: "assertion-rule", Type: "*types.NewAssertionRule", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "GatingRule", Flag: "gating-rule", Type: "*types.NewGatingRule", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_cluster = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
}

var fields_delete_control_panel = []leanruntime.Field{
	{Name: "ControlPanelArn", Flag: "control-panel-arn", Type: "*string", Required: true},
}

var fields_delete_routing_control = []leanruntime.Field{
	{Name: "RoutingControlArn", Flag: "routing-control-arn", Type: "*string", Required: true},
}

var fields_delete_safety_rule = []leanruntime.Field{
	{Name: "SafetyRuleArn", Flag: "safety-rule-arn", Type: "*string", Required: true},
}

var fields_describe_cluster = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
}

var fields_describe_control_panel = []leanruntime.Field{
	{Name: "ControlPanelArn", Flag: "control-panel-arn", Type: "*string", Required: true},
}

var fields_describe_routing_control = []leanruntime.Field{
	{Name: "RoutingControlArn", Flag: "routing-control-arn", Type: "*string", Required: true},
}

var fields_describe_safety_rule = []leanruntime.Field{
	{Name: "SafetyRuleArn", Flag: "safety-rule-arn", Type: "*string", Required: true},
}

var fields_get_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_associated_route53_health_checks = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RoutingControlArn", Flag: "routing-control-arn", Type: "*string", Required: true},
}

var fields_list_clusters = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_control_panels = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_routing_controls = []leanruntime.Field{
	{Name: "ControlPanelArn", Flag: "control-panel-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_safety_rules = []leanruntime.Field{
	{Name: "ControlPanelArn", Flag: "control-panel-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
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

var fields_update_cluster = []leanruntime.Field{
	{Name: "ClusterArn", Flag: "cluster-arn", Type: "*string", Required: true},
	{Name: "NetworkType", Flag: "network-type", Type: "types.NetworkType", Required: true},
}

var fields_update_control_panel = []leanruntime.Field{
	{Name: "ControlPanelArn", Flag: "control-panel-arn", Type: "*string", Required: true},
	{Name: "ControlPanelName", Flag: "control-panel-name", Type: "*string", Required: true},
}

var fields_update_routing_control = []leanruntime.Field{
	{Name: "RoutingControlArn", Flag: "routing-control-arn", Type: "*string", Required: true},
	{Name: "RoutingControlName", Flag: "routing-control-name", Type: "*string", Required: true},
}

var fields_update_safety_rule = []leanruntime.Field{
	{Name: "AssertionRuleUpdate", Flag: "assertion-rule-update", Type: "*types.AssertionRuleUpdate", Required: false},
	{Name: "GatingRuleUpdate", Flag: "gating-rule-update", Type: "*types.GatingRuleUpdate", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-cluster": {
			Name:   "create-cluster",
			Fields: fields_create_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCluster(ctx, input)
			},
		},
		"create-control-panel": {
			Name:   "create-control-panel",
			Fields: fields_create_control_panel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateControlPanelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_control_panel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateControlPanel(ctx, input)
			},
		},
		"create-routing-control": {
			Name:   "create-routing-control",
			Fields: fields_create_routing_control,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRoutingControlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_routing_control, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRoutingControl(ctx, input)
			},
		},
		"create-safety-rule": {
			Name:   "create-safety-rule",
			Fields: fields_create_safety_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSafetyRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_safety_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSafetyRule(ctx, input)
			},
		},
		"delete-cluster": {
			Name:   "delete-cluster",
			Fields: fields_delete_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCluster(ctx, input)
			},
		},
		"delete-control-panel": {
			Name:   "delete-control-panel",
			Fields: fields_delete_control_panel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteControlPanelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_control_panel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteControlPanel(ctx, input)
			},
		},
		"delete-routing-control": {
			Name:   "delete-routing-control",
			Fields: fields_delete_routing_control,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRoutingControlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_routing_control, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRoutingControl(ctx, input)
			},
		},
		"delete-safety-rule": {
			Name:   "delete-safety-rule",
			Fields: fields_delete_safety_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSafetyRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_safety_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSafetyRule(ctx, input)
			},
		},
		"describe-cluster": {
			Name:   "describe-cluster",
			Fields: fields_describe_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCluster(ctx, input)
			},
		},
		"describe-control-panel": {
			Name:   "describe-control-panel",
			Fields: fields_describe_control_panel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeControlPanelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_control_panel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeControlPanel(ctx, input)
			},
		},
		"describe-routing-control": {
			Name:   "describe-routing-control",
			Fields: fields_describe_routing_control,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRoutingControlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_routing_control, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRoutingControl(ctx, input)
			},
		},
		"describe-safety-rule": {
			Name:   "describe-safety-rule",
			Fields: fields_describe_safety_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSafetyRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_safety_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSafetyRule(ctx, input)
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
		"list-associated-route53-health-checks": {
			Name:   "list-associated-route53-health-checks",
			Fields: fields_list_associated_route53_health_checks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssociatedRoute53HealthChecksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_associated_route53_health_checks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssociatedRoute53HealthChecks(ctx, input)
				}
				var results []*svc.ListAssociatedRoute53HealthChecksOutput
				p := svc.NewListAssociatedRoute53HealthChecksPaginator(client, input)
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
		"list-clusters": {
			Name:   "list-clusters",
			Fields: fields_list_clusters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListClustersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_clusters, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListClusters(ctx, input)
				}
				var results []*svc.ListClustersOutput
				p := svc.NewListClustersPaginator(client, input)
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
		"list-control-panels": {
			Name:   "list-control-panels",
			Fields: fields_list_control_panels,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListControlPanelsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_control_panels, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListControlPanels(ctx, input)
				}
				var results []*svc.ListControlPanelsOutput
				p := svc.NewListControlPanelsPaginator(client, input)
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
		"list-safety-rules": {
			Name:   "list-safety-rules",
			Fields: fields_list_safety_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSafetyRulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_safety_rules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSafetyRules(ctx, input)
				}
				var results []*svc.ListSafetyRulesOutput
				p := svc.NewListSafetyRulesPaginator(client, input)
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
		"update-cluster": {
			Name:   "update-cluster",
			Fields: fields_update_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCluster(ctx, input)
			},
		},
		"update-control-panel": {
			Name:   "update-control-panel",
			Fields: fields_update_control_panel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateControlPanelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_control_panel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateControlPanel(ctx, input)
			},
		},
		"update-routing-control": {
			Name:   "update-routing-control",
			Fields: fields_update_routing_control,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRoutingControlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_routing_control, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRoutingControl(ctx, input)
			},
		},
		"update-safety-rule": {
			Name:   "update-safety-rule",
			Fields: fields_update_safety_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSafetyRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_safety_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSafetyRule(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("route53recoverycontrolconfig", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
