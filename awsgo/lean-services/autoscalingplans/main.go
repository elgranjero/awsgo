package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/autoscalingplans"
)

var fields_create_scaling_plan = []leanruntime.Field{
	{Name: "ApplicationSource", Flag: "application-source", Type: "*types.ApplicationSource", Required: true},
	{Name: "ScalingInstructions", Flag: "scaling-instructions", Type: "[]types.ScalingInstruction", Required: true},
	{Name: "ScalingPlanName", Flag: "scaling-plan-name", Type: "*string", Required: true},
}

var fields_delete_scaling_plan = []leanruntime.Field{
	{Name: "ScalingPlanName", Flag: "scaling-plan-name", Type: "*string", Required: true},
	{Name: "ScalingPlanVersion", Flag: "scaling-plan-version", Type: "*int64", Required: true},
}

var fields_describe_scaling_plan_resources = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ScalingPlanName", Flag: "scaling-plan-name", Type: "*string", Required: true},
	{Name: "ScalingPlanVersion", Flag: "scaling-plan-version", Type: "*int64", Required: true},
}

var fields_describe_scaling_plans = []leanruntime.Field{
	{Name: "ApplicationSources", Flag: "application-sources", Type: "[]types.ApplicationSource", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ScalingPlanNames", Flag: "scaling-plan-names", Type: "[]string", Required: false},
	{Name: "ScalingPlanVersion", Flag: "scaling-plan-version", Type: "*int64", Required: false},
}

var fields_get_scaling_plan_resource_forecast_data = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "ForecastDataType", Flag: "forecast-data-type", Type: "types.ForecastDataType", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "ScalableDimension", Flag: "scalable-dimension", Type: "types.ScalableDimension", Required: true},
	{Name: "ScalingPlanName", Flag: "scaling-plan-name", Type: "*string", Required: true},
	{Name: "ScalingPlanVersion", Flag: "scaling-plan-version", Type: "*int64", Required: true},
	{Name: "ServiceNamespace", Flag: "service-namespace", Type: "types.ServiceNamespace", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_update_scaling_plan = []leanruntime.Field{
	{Name: "ApplicationSource", Flag: "application-source", Type: "*types.ApplicationSource", Required: false},
	{Name: "ScalingInstructions", Flag: "scaling-instructions", Type: "[]types.ScalingInstruction", Required: false},
	{Name: "ScalingPlanName", Flag: "scaling-plan-name", Type: "*string", Required: true},
	{Name: "ScalingPlanVersion", Flag: "scaling-plan-version", Type: "*int64", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-scaling-plan": {
			Name:   "create-scaling-plan",
			Fields: fields_create_scaling_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateScalingPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_scaling_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateScalingPlan(ctx, input)
			},
		},
		"delete-scaling-plan": {
			Name:   "delete-scaling-plan",
			Fields: fields_delete_scaling_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteScalingPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_scaling_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteScalingPlan(ctx, input)
			},
		},
		"describe-scaling-plan-resources": {
			Name:   "describe-scaling-plan-resources",
			Fields: fields_describe_scaling_plan_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeScalingPlanResourcesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_scaling_plan_resources, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeScalingPlanResources(ctx, input)
			},
		},
		"describe-scaling-plans": {
			Name:   "describe-scaling-plans",
			Fields: fields_describe_scaling_plans,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeScalingPlansInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_scaling_plans, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeScalingPlans(ctx, input)
			},
		},
		"get-scaling-plan-resource-forecast-data": {
			Name:   "get-scaling-plan-resource-forecast-data",
			Fields: fields_get_scaling_plan_resource_forecast_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetScalingPlanResourceForecastDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_scaling_plan_resource_forecast_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetScalingPlanResourceForecastData(ctx, input)
			},
		},
		"update-scaling-plan": {
			Name:   "update-scaling-plan",
			Fields: fields_update_scaling_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateScalingPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_scaling_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateScalingPlan(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("autoscalingplans", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
