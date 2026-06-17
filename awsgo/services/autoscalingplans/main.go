package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/autoscalingplans/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-scaling-plan", "delete-scaling-plan", "describe-scaling-plan-resources", "describe-scaling-plans", "get-scaling-plan-resource-forecast-data", "update-scaling-plan"},
		OperationSet: map[string]bool{"create-scaling-plan": true, "delete-scaling-plan": true, "describe-scaling-plan-resources": true, "describe-scaling-plans": true, "get-scaling-plan-resource-forecast-data": true, "update-scaling-plan": true},
		OperationInputs: map[string][]string{
			"create-scaling-plan":                     {"ApplicationSource", "ScalingInstructions", "ScalingPlanName"},
			"delete-scaling-plan":                     {"ScalingPlanName", "ScalingPlanVersion"},
			"describe-scaling-plan-resources":         {"MaxResults", "NextToken", "ScalingPlanName", "ScalingPlanVersion"},
			"describe-scaling-plans":                  {"ApplicationSources", "MaxResults", "NextToken", "ScalingPlanNames", "ScalingPlanVersion"},
			"get-scaling-plan-resource-forecast-data": {"EndTime", "ForecastDataType", "ResourceId", "ScalableDimension", "ScalingPlanName", "ScalingPlanVersion", "ServiceNamespace", "StartTime"},
			"update-scaling-plan":                     {"ApplicationSource", "ScalingInstructions", "ScalingPlanName", "ScalingPlanVersion"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-scaling-plan":                     {"ApplicationSource": "*types.ApplicationSource", "ScalingInstructions": "[]types.ScalingInstruction", "ScalingPlanName": "*string"},
			"delete-scaling-plan":                     {"ScalingPlanName": "*string", "ScalingPlanVersion": "*int64"},
			"describe-scaling-plan-resources":         {"MaxResults": "*int32", "NextToken": "*string", "ScalingPlanName": "*string", "ScalingPlanVersion": "*int64"},
			"describe-scaling-plans":                  {"ApplicationSources": "[]types.ApplicationSource", "MaxResults": "*int32", "NextToken": "*string", "ScalingPlanNames": "[]string", "ScalingPlanVersion": "*int64"},
			"get-scaling-plan-resource-forecast-data": {"EndTime": "*time.Time", "ForecastDataType": "types.ForecastDataType", "ResourceId": "*string", "ScalableDimension": "types.ScalableDimension", "ScalingPlanName": "*string", "ScalingPlanVersion": "*int64", "ServiceNamespace": "types.ServiceNamespace", "StartTime": "*time.Time"},
			"update-scaling-plan":                     {"ApplicationSource": "*types.ApplicationSource", "ScalingInstructions": "[]types.ScalingInstruction", "ScalingPlanName": "*string", "ScalingPlanVersion": "*int64"},
		},
		OperationInputRequired: map[string][]string{
			"create-scaling-plan":                     {"ApplicationSource", "ScalingInstructions", "ScalingPlanName"},
			"delete-scaling-plan":                     {"ScalingPlanName", "ScalingPlanVersion"},
			"describe-scaling-plan-resources":         {"ScalingPlanName", "ScalingPlanVersion"},
			"describe-scaling-plans":                  {},
			"get-scaling-plan-resource-forecast-data": {"EndTime", "ForecastDataType", "ResourceId", "ScalableDimension", "ScalingPlanName", "ScalingPlanVersion", "ServiceNamespace", "StartTime"},
			"update-scaling-plan":                     {"ScalingPlanName", "ScalingPlanVersion"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("autoscalingplans", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
