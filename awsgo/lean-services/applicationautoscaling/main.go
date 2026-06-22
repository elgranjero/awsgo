package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
)

var fields_delete_scaling_policy = []leanruntime.Field{
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "ScalableDimension", Flag: "scalable-dimension", Type: "types.ScalableDimension", Required: true},
	{Name: "ServiceNamespace", Flag: "service-namespace", Type: "types.ServiceNamespace", Required: true},
}

var fields_delete_scheduled_action = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "ScalableDimension", Flag: "scalable-dimension", Type: "types.ScalableDimension", Required: true},
	{Name: "ScheduledActionName", Flag: "scheduled-action-name", Type: "*string", Required: true},
	{Name: "ServiceNamespace", Flag: "service-namespace", Type: "types.ServiceNamespace", Required: true},
}

var fields_deregister_scalable_target = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "ScalableDimension", Flag: "scalable-dimension", Type: "types.ScalableDimension", Required: true},
	{Name: "ServiceNamespace", Flag: "service-namespace", Type: "types.ServiceNamespace", Required: true},
}

var fields_describe_scalable_targets = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceIds", Flag: "resource-ids", Type: "[]string", Required: false},
	{Name: "ScalableDimension", Flag: "scalable-dimension", Type: "types.ScalableDimension", Required: false},
	{Name: "ServiceNamespace", Flag: "service-namespace", Type: "types.ServiceNamespace", Required: true},
}

var fields_describe_scaling_activities = []leanruntime.Field{
	{Name: "IncludeNotScaledActivities", Flag: "include-not-scaled-activities", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: false},
	{Name: "ScalableDimension", Flag: "scalable-dimension", Type: "types.ScalableDimension", Required: false},
	{Name: "ServiceNamespace", Flag: "service-namespace", Type: "types.ServiceNamespace", Required: true},
}

var fields_describe_scaling_policies = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PolicyNames", Flag: "policy-names", Type: "[]string", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: false},
	{Name: "ScalableDimension", Flag: "scalable-dimension", Type: "types.ScalableDimension", Required: false},
	{Name: "ServiceNamespace", Flag: "service-namespace", Type: "types.ServiceNamespace", Required: true},
}

var fields_describe_scheduled_actions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: false},
	{Name: "ScalableDimension", Flag: "scalable-dimension", Type: "types.ScalableDimension", Required: false},
	{Name: "ScheduledActionNames", Flag: "scheduled-action-names", Type: "[]string", Required: false},
	{Name: "ServiceNamespace", Flag: "service-namespace", Type: "types.ServiceNamespace", Required: true},
}

var fields_get_predictive_scaling_forecast = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "ScalableDimension", Flag: "scalable-dimension", Type: "types.ScalableDimension", Required: true},
	{Name: "ServiceNamespace", Flag: "service-namespace", Type: "types.ServiceNamespace", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_scaling_policy = []leanruntime.Field{
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
	{Name: "PolicyType", Flag: "policy-type", Type: "types.PolicyType", Required: false},
	{Name: "PredictiveScalingPolicyConfiguration", Flag: "predictive-scaling-policy-configuration", Type: "*types.PredictiveScalingPolicyConfiguration", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "ScalableDimension", Flag: "scalable-dimension", Type: "types.ScalableDimension", Required: true},
	{Name: "ServiceNamespace", Flag: "service-namespace", Type: "types.ServiceNamespace", Required: true},
	{Name: "StepScalingPolicyConfiguration", Flag: "step-scaling-policy-configuration", Type: "*types.StepScalingPolicyConfiguration", Required: false},
	{Name: "TargetTrackingScalingPolicyConfiguration", Flag: "target-tracking-scaling-policy-configuration", Type: "*types.TargetTrackingScalingPolicyConfiguration", Required: false},
}

var fields_put_scheduled_action = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "ScalableDimension", Flag: "scalable-dimension", Type: "types.ScalableDimension", Required: true},
	{Name: "ScalableTargetAction", Flag: "scalable-target-action", Type: "*types.ScalableTargetAction", Required: false},
	{Name: "Schedule", Flag: "schedule", Type: "*string", Required: false},
	{Name: "ScheduledActionName", Flag: "scheduled-action-name", Type: "*string", Required: true},
	{Name: "ServiceNamespace", Flag: "service-namespace", Type: "types.ServiceNamespace", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
	{Name: "Timezone", Flag: "timezone", Type: "*string", Required: false},
}

var fields_register_scalable_target = []leanruntime.Field{
	{Name: "MaxCapacity", Flag: "max-capacity", Type: "*int32", Required: false},
	{Name: "MinCapacity", Flag: "min-capacity", Type: "*int32", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "RoleARN", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "ScalableDimension", Flag: "scalable-dimension", Type: "types.ScalableDimension", Required: true},
	{Name: "ServiceNamespace", Flag: "service-namespace", Type: "types.ServiceNamespace", Required: true},
	{Name: "SuspendedState", Flag: "suspended-state", Type: "*types.SuspendedState", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"delete-scaling-policy": {
			Name:   "delete-scaling-policy",
			Fields: fields_delete_scaling_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteScalingPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_scaling_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteScalingPolicy(ctx, input)
			},
		},
		"delete-scheduled-action": {
			Name:   "delete-scheduled-action",
			Fields: fields_delete_scheduled_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteScheduledActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_scheduled_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteScheduledAction(ctx, input)
			},
		},
		"deregister-scalable-target": {
			Name:   "deregister-scalable-target",
			Fields: fields_deregister_scalable_target,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterScalableTargetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_scalable_target, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterScalableTarget(ctx, input)
			},
		},
		"describe-scalable-targets": {
			Name:   "describe-scalable-targets",
			Fields: fields_describe_scalable_targets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeScalableTargetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_scalable_targets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeScalableTargets(ctx, input)
				}
				var results []*svc.DescribeScalableTargetsOutput
				p := svc.NewDescribeScalableTargetsPaginator(client, input)
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
		"describe-scaling-activities": {
			Name:   "describe-scaling-activities",
			Fields: fields_describe_scaling_activities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeScalingActivitiesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_scaling_activities, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeScalingActivities(ctx, input)
				}
				var results []*svc.DescribeScalingActivitiesOutput
				p := svc.NewDescribeScalingActivitiesPaginator(client, input)
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
		"describe-scaling-policies": {
			Name:   "describe-scaling-policies",
			Fields: fields_describe_scaling_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeScalingPoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_scaling_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeScalingPolicies(ctx, input)
				}
				var results []*svc.DescribeScalingPoliciesOutput
				p := svc.NewDescribeScalingPoliciesPaginator(client, input)
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
		"describe-scheduled-actions": {
			Name:   "describe-scheduled-actions",
			Fields: fields_describe_scheduled_actions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeScheduledActionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_scheduled_actions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeScheduledActions(ctx, input)
				}
				var results []*svc.DescribeScheduledActionsOutput
				p := svc.NewDescribeScheduledActionsPaginator(client, input)
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
		"get-predictive-scaling-forecast": {
			Name:   "get-predictive-scaling-forecast",
			Fields: fields_get_predictive_scaling_forecast,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPredictiveScalingForecastInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_predictive_scaling_forecast, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPredictiveScalingForecast(ctx, input)
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
		"put-scaling-policy": {
			Name:   "put-scaling-policy",
			Fields: fields_put_scaling_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutScalingPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_scaling_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutScalingPolicy(ctx, input)
			},
		},
		"put-scheduled-action": {
			Name:   "put-scheduled-action",
			Fields: fields_put_scheduled_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutScheduledActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_scheduled_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutScheduledAction(ctx, input)
			},
		},
		"register-scalable-target": {
			Name:   "register-scalable-target",
			Fields: fields_register_scalable_target,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterScalableTargetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_scalable_target, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterScalableTarget(ctx, input)
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
	if err := leanruntime.Execute("applicationautoscaling", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
