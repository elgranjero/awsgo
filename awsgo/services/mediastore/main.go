package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/mediastore/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-container", "delete-container", "delete-container-policy", "delete-cors-policy", "delete-lifecycle-policy", "delete-metric-policy", "describe-container", "get-container-policy", "get-cors-policy", "get-lifecycle-policy", "get-metric-policy", "list-containers", "list-tags-for-resource", "put-container-policy", "put-cors-policy", "put-lifecycle-policy", "put-metric-policy", "start-access-logging", "stop-access-logging", "tag-resource", "untag-resource"},
		OperationSet: map[string]bool{"create-container": true, "delete-container": true, "delete-container-policy": true, "delete-cors-policy": true, "delete-lifecycle-policy": true, "delete-metric-policy": true, "describe-container": true, "get-container-policy": true, "get-cors-policy": true, "get-lifecycle-policy": true, "get-metric-policy": true, "list-containers": true, "list-tags-for-resource": true, "put-container-policy": true, "put-cors-policy": true, "put-lifecycle-policy": true, "put-metric-policy": true, "start-access-logging": true, "stop-access-logging": true, "tag-resource": true, "untag-resource": true},
		OperationInputs: map[string][]string{
			"create-container":        {"ContainerName", "Tags"},
			"delete-container":        {"ContainerName"},
			"delete-container-policy": {"ContainerName"},
			"delete-cors-policy":      {"ContainerName"},
			"delete-lifecycle-policy": {"ContainerName"},
			"delete-metric-policy":    {"ContainerName"},
			"describe-container":      {"ContainerName"},
			"get-container-policy":    {"ContainerName"},
			"get-cors-policy":         {"ContainerName"},
			"get-lifecycle-policy":    {"ContainerName"},
			"get-metric-policy":       {"ContainerName"},
			"list-containers":         {"MaxResults", "NextToken"},
			"list-tags-for-resource":  {"Resource"},
			"put-container-policy":    {"ContainerName", "Policy"},
			"put-cors-policy":         {"ContainerName", "CorsPolicy"},
			"put-lifecycle-policy":    {"ContainerName", "LifecyclePolicy"},
			"put-metric-policy":       {"ContainerName", "MetricPolicy"},
			"start-access-logging":    {"ContainerName"},
			"stop-access-logging":     {"ContainerName"},
			"tag-resource":            {"Resource", "Tags"},
			"untag-resource":          {"Resource", "TagKeys"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-container":        {"ContainerName": "*string", "Tags": "[]types.Tag"},
			"delete-container":        {"ContainerName": "*string"},
			"delete-container-policy": {"ContainerName": "*string"},
			"delete-cors-policy":      {"ContainerName": "*string"},
			"delete-lifecycle-policy": {"ContainerName": "*string"},
			"delete-metric-policy":    {"ContainerName": "*string"},
			"describe-container":      {"ContainerName": "*string"},
			"get-container-policy":    {"ContainerName": "*string"},
			"get-cors-policy":         {"ContainerName": "*string"},
			"get-lifecycle-policy":    {"ContainerName": "*string"},
			"get-metric-policy":       {"ContainerName": "*string"},
			"list-containers":         {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":  {"Resource": "*string"},
			"put-container-policy":    {"ContainerName": "*string", "Policy": "*string"},
			"put-cors-policy":         {"ContainerName": "*string", "CorsPolicy": "[]types.CorsRule"},
			"put-lifecycle-policy":    {"ContainerName": "*string", "LifecyclePolicy": "*string"},
			"put-metric-policy":       {"ContainerName": "*string", "MetricPolicy": "*types.MetricPolicy"},
			"start-access-logging":    {"ContainerName": "*string"},
			"stop-access-logging":     {"ContainerName": "*string"},
			"tag-resource":            {"Resource": "*string", "Tags": "[]types.Tag"},
			"untag-resource":          {"Resource": "*string", "TagKeys": "[]string"},
		},
		OperationInputRequired: map[string][]string{
			"create-container":        {"ContainerName"},
			"delete-container":        {"ContainerName"},
			"delete-container-policy": {"ContainerName"},
			"delete-cors-policy":      {"ContainerName"},
			"delete-lifecycle-policy": {"ContainerName"},
			"delete-metric-policy":    {"ContainerName"},
			"describe-container":      {},
			"get-container-policy":    {"ContainerName"},
			"get-cors-policy":         {"ContainerName"},
			"get-lifecycle-policy":    {"ContainerName"},
			"get-metric-policy":       {"ContainerName"},
			"list-containers":         {},
			"list-tags-for-resource":  {"Resource"},
			"put-container-policy":    {"ContainerName", "Policy"},
			"put-cors-policy":         {"ContainerName", "CorsPolicy"},
			"put-lifecycle-policy":    {"ContainerName", "LifecyclePolicy"},
			"put-metric-policy":       {"ContainerName", "MetricPolicy"},
			"start-access-logging":    {"ContainerName"},
			"stop-access-logging":     {"ContainerName"},
			"tag-resource":            {"Resource", "Tags"},
			"untag-resource":          {"Resource", "TagKeys"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("mediastore", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
