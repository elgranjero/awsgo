package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/applicationautoscaling/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"delete-scaling-policy", "delete-scheduled-action", "deregister-scalable-target", "describe-scalable-targets", "describe-scaling-activities", "describe-scaling-policies", "describe-scheduled-actions", "get-predictive-scaling-forecast", "list-tags-for-resource", "put-scaling-policy", "put-scheduled-action", "register-scalable-target", "tag-resource", "untag-resource"},
		OperationSet: map[string]bool{"delete-scaling-policy": true, "delete-scheduled-action": true, "deregister-scalable-target": true, "describe-scalable-targets": true, "describe-scaling-activities": true, "describe-scaling-policies": true, "describe-scheduled-actions": true, "get-predictive-scaling-forecast": true, "list-tags-for-resource": true, "put-scaling-policy": true, "put-scheduled-action": true, "register-scalable-target": true, "tag-resource": true, "untag-resource": true},
		OperationInputs: map[string][]string{
			"delete-scaling-policy":           {"PolicyName", "ResourceId", "ScalableDimension", "ServiceNamespace"},
			"delete-scheduled-action":         {"ResourceId", "ScalableDimension", "ScheduledActionName", "ServiceNamespace"},
			"deregister-scalable-target":      {"ResourceId", "ScalableDimension", "ServiceNamespace"},
			"describe-scalable-targets":       {"MaxResults", "NextToken", "ResourceIds", "ScalableDimension", "ServiceNamespace"},
			"describe-scaling-activities":     {"IncludeNotScaledActivities", "MaxResults", "NextToken", "ResourceId", "ScalableDimension", "ServiceNamespace"},
			"describe-scaling-policies":       {"MaxResults", "NextToken", "PolicyNames", "ResourceId", "ScalableDimension", "ServiceNamespace"},
			"describe-scheduled-actions":      {"MaxResults", "NextToken", "ResourceId", "ScalableDimension", "ScheduledActionNames", "ServiceNamespace"},
			"get-predictive-scaling-forecast": {"EndTime", "PolicyName", "ResourceId", "ScalableDimension", "ServiceNamespace", "StartTime"},
			"list-tags-for-resource":          {"ResourceARN"},
			"put-scaling-policy":              {"PolicyName", "PolicyType", "PredictiveScalingPolicyConfiguration", "ResourceId", "ScalableDimension", "ServiceNamespace", "StepScalingPolicyConfiguration", "TargetTrackingScalingPolicyConfiguration"},
			"put-scheduled-action":            {"EndTime", "ResourceId", "ScalableDimension", "ScalableTargetAction", "Schedule", "ScheduledActionName", "ServiceNamespace", "StartTime", "Timezone"},
			"register-scalable-target":        {"MaxCapacity", "MinCapacity", "ResourceId", "RoleARN", "ScalableDimension", "ServiceNamespace", "SuspendedState", "Tags"},
			"tag-resource":                    {"ResourceARN", "Tags"},
			"untag-resource":                  {"ResourceARN", "TagKeys"},
		},
		OperationInputTypes: map[string]map[string]string{
			"delete-scaling-policy":           {"PolicyName": "*string", "ResourceId": "*string", "ScalableDimension": "types.ScalableDimension", "ServiceNamespace": "types.ServiceNamespace"},
			"delete-scheduled-action":         {"ResourceId": "*string", "ScalableDimension": "types.ScalableDimension", "ScheduledActionName": "*string", "ServiceNamespace": "types.ServiceNamespace"},
			"deregister-scalable-target":      {"ResourceId": "*string", "ScalableDimension": "types.ScalableDimension", "ServiceNamespace": "types.ServiceNamespace"},
			"describe-scalable-targets":       {"MaxResults": "*int32", "NextToken": "*string", "ResourceIds": "[]string", "ScalableDimension": "types.ScalableDimension", "ServiceNamespace": "types.ServiceNamespace"},
			"describe-scaling-activities":     {"IncludeNotScaledActivities": "*bool", "MaxResults": "*int32", "NextToken": "*string", "ResourceId": "*string", "ScalableDimension": "types.ScalableDimension", "ServiceNamespace": "types.ServiceNamespace"},
			"describe-scaling-policies":       {"MaxResults": "*int32", "NextToken": "*string", "PolicyNames": "[]string", "ResourceId": "*string", "ScalableDimension": "types.ScalableDimension", "ServiceNamespace": "types.ServiceNamespace"},
			"describe-scheduled-actions":      {"MaxResults": "*int32", "NextToken": "*string", "ResourceId": "*string", "ScalableDimension": "types.ScalableDimension", "ScheduledActionNames": "[]string", "ServiceNamespace": "types.ServiceNamespace"},
			"get-predictive-scaling-forecast": {"EndTime": "*time.Time", "PolicyName": "*string", "ResourceId": "*string", "ScalableDimension": "types.ScalableDimension", "ServiceNamespace": "types.ServiceNamespace", "StartTime": "*time.Time"},
			"list-tags-for-resource":          {"ResourceARN": "*string"},
			"put-scaling-policy":              {"PolicyName": "*string", "PolicyType": "types.PolicyType", "PredictiveScalingPolicyConfiguration": "*types.PredictiveScalingPolicyConfiguration", "ResourceId": "*string", "ScalableDimension": "types.ScalableDimension", "ServiceNamespace": "types.ServiceNamespace", "StepScalingPolicyConfiguration": "*types.StepScalingPolicyConfiguration", "TargetTrackingScalingPolicyConfiguration": "*types.TargetTrackingScalingPolicyConfiguration"},
			"put-scheduled-action":            {"EndTime": "*time.Time", "ResourceId": "*string", "ScalableDimension": "types.ScalableDimension", "ScalableTargetAction": "*types.ScalableTargetAction", "Schedule": "*string", "ScheduledActionName": "*string", "ServiceNamespace": "types.ServiceNamespace", "StartTime": "*time.Time", "Timezone": "*string"},
			"register-scalable-target":        {"MaxCapacity": "*int32", "MinCapacity": "*int32", "ResourceId": "*string", "RoleARN": "*string", "ScalableDimension": "types.ScalableDimension", "ServiceNamespace": "types.ServiceNamespace", "SuspendedState": "*types.SuspendedState", "Tags": "map[string]string"},
			"tag-resource":                    {"ResourceARN": "*string", "Tags": "map[string]string"},
			"untag-resource":                  {"ResourceARN": "*string", "TagKeys": "[]string"},
		},
		OperationInputRequired: map[string][]string{
			"delete-scaling-policy":           {"PolicyName", "ResourceId", "ScalableDimension", "ServiceNamespace"},
			"delete-scheduled-action":         {"ResourceId", "ScalableDimension", "ScheduledActionName", "ServiceNamespace"},
			"deregister-scalable-target":      {"ResourceId", "ScalableDimension", "ServiceNamespace"},
			"describe-scalable-targets":       {"ServiceNamespace"},
			"describe-scaling-activities":     {"ServiceNamespace"},
			"describe-scaling-policies":       {"ServiceNamespace"},
			"describe-scheduled-actions":      {"ServiceNamespace"},
			"get-predictive-scaling-forecast": {"EndTime", "PolicyName", "ResourceId", "ScalableDimension", "ServiceNamespace", "StartTime"},
			"list-tags-for-resource":          {"ResourceARN"},
			"put-scaling-policy":              {"PolicyName", "ResourceId", "ScalableDimension", "ServiceNamespace"},
			"put-scheduled-action":            {"ResourceId", "ScalableDimension", "ScheduledActionName", "ServiceNamespace"},
			"register-scalable-target":        {"ResourceId", "ScalableDimension", "ServiceNamespace"},
			"tag-resource":                    {"ResourceARN", "Tags"},
			"untag-resource":                  {"ResourceARN", "TagKeys"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("applicationautoscaling", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
