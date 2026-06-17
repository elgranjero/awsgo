package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/osis/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-pipeline", "create-pipeline-endpoint", "delete-pipeline", "delete-pipeline-endpoint", "delete-resource-policy", "get-pipeline", "get-pipeline-blueprint", "get-pipeline-change-progress", "get-resource-policy", "list-pipeline-blueprints", "list-pipeline-endpoint-connections", "list-pipeline-endpoints", "list-pipelines", "list-tags-for-resource", "put-resource-policy", "revoke-pipeline-endpoint-connections", "start-pipeline", "stop-pipeline", "tag-resource", "untag-resource", "update-pipeline", "validate-pipeline"},
		OperationSet: map[string]bool{"create-pipeline": true, "create-pipeline-endpoint": true, "delete-pipeline": true, "delete-pipeline-endpoint": true, "delete-resource-policy": true, "get-pipeline": true, "get-pipeline-blueprint": true, "get-pipeline-change-progress": true, "get-resource-policy": true, "list-pipeline-blueprints": true, "list-pipeline-endpoint-connections": true, "list-pipeline-endpoints": true, "list-pipelines": true, "list-tags-for-resource": true, "put-resource-policy": true, "revoke-pipeline-endpoint-connections": true, "start-pipeline": true, "stop-pipeline": true, "tag-resource": true, "untag-resource": true, "update-pipeline": true, "validate-pipeline": true},
		OperationInputs: map[string][]string{
			"create-pipeline":                      {"BufferOptions", "EncryptionAtRestOptions", "LogPublishingOptions", "MaxUnits", "MinUnits", "PipelineConfigurationBody", "PipelineName", "PipelineRoleArn", "Tags", "VpcOptions"},
			"create-pipeline-endpoint":             {"PipelineArn", "VpcOptions"},
			"delete-pipeline":                      {"PipelineName"},
			"delete-pipeline-endpoint":             {"EndpointId"},
			"delete-resource-policy":               {"ResourceArn"},
			"get-pipeline":                         {"PipelineName"},
			"get-pipeline-blueprint":               {"BlueprintName", "Format"},
			"get-pipeline-change-progress":         {"PipelineName"},
			"get-resource-policy":                  {"ResourceArn"},
			"list-pipeline-blueprints":             {},
			"list-pipeline-endpoint-connections":   {"MaxResults", "NextToken"},
			"list-pipeline-endpoints":              {"MaxResults", "NextToken"},
			"list-pipelines":                       {"MaxResults", "NextToken"},
			"list-tags-for-resource":               {"Arn"},
			"put-resource-policy":                  {"Policy", "ResourceArn"},
			"revoke-pipeline-endpoint-connections": {"EndpointIds", "PipelineArn"},
			"start-pipeline":                       {"PipelineName"},
			"stop-pipeline":                        {"PipelineName"},
			"tag-resource":                         {"Arn", "Tags"},
			"untag-resource":                       {"Arn", "TagKeys"},
			"update-pipeline":                      {"BufferOptions", "EncryptionAtRestOptions", "LogPublishingOptions", "MaxUnits", "MinUnits", "PipelineConfigurationBody", "PipelineName", "PipelineRoleArn"},
			"validate-pipeline":                    {"PipelineConfigurationBody"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-pipeline":                      {"BufferOptions": "*types.BufferOptions", "EncryptionAtRestOptions": "*types.EncryptionAtRestOptions", "LogPublishingOptions": "*types.LogPublishingOptions", "MaxUnits": "*int32", "MinUnits": "*int32", "PipelineConfigurationBody": "*string", "PipelineName": "*string", "PipelineRoleArn": "*string", "Tags": "[]types.Tag", "VpcOptions": "*types.VpcOptions"},
			"create-pipeline-endpoint":             {"PipelineArn": "*string", "VpcOptions": "*types.PipelineEndpointVpcOptions"},
			"delete-pipeline":                      {"PipelineName": "*string"},
			"delete-pipeline-endpoint":             {"EndpointId": "*string"},
			"delete-resource-policy":               {"ResourceArn": "*string"},
			"get-pipeline":                         {"PipelineName": "*string"},
			"get-pipeline-blueprint":               {"BlueprintName": "*string", "Format": "*string"},
			"get-pipeline-change-progress":         {"PipelineName": "*string"},
			"get-resource-policy":                  {"ResourceArn": "*string"},
			"list-pipeline-blueprints":             {},
			"list-pipeline-endpoint-connections":   {"MaxResults": "*int32", "NextToken": "*string"},
			"list-pipeline-endpoints":              {"MaxResults": "*int32", "NextToken": "*string"},
			"list-pipelines":                       {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":               {"Arn": "*string"},
			"put-resource-policy":                  {"Policy": "*string", "ResourceArn": "*string"},
			"revoke-pipeline-endpoint-connections": {"EndpointIds": "[]string", "PipelineArn": "*string"},
			"start-pipeline":                       {"PipelineName": "*string"},
			"stop-pipeline":                        {"PipelineName": "*string"},
			"tag-resource":                         {"Arn": "*string", "Tags": "[]types.Tag"},
			"untag-resource":                       {"Arn": "*string", "TagKeys": "[]string"},
			"update-pipeline":                      {"BufferOptions": "*types.BufferOptions", "EncryptionAtRestOptions": "*types.EncryptionAtRestOptions", "LogPublishingOptions": "*types.LogPublishingOptions", "MaxUnits": "*int32", "MinUnits": "*int32", "PipelineConfigurationBody": "*string", "PipelineName": "*string", "PipelineRoleArn": "*string"},
			"validate-pipeline":                    {"PipelineConfigurationBody": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"create-pipeline":                      {"MaxUnits", "MinUnits", "PipelineConfigurationBody", "PipelineName"},
			"create-pipeline-endpoint":             {"PipelineArn", "VpcOptions"},
			"delete-pipeline":                      {"PipelineName"},
			"delete-pipeline-endpoint":             {"EndpointId"},
			"delete-resource-policy":               {"ResourceArn"},
			"get-pipeline":                         {"PipelineName"},
			"get-pipeline-blueprint":               {"BlueprintName"},
			"get-pipeline-change-progress":         {"PipelineName"},
			"get-resource-policy":                  {"ResourceArn"},
			"list-pipeline-blueprints":             {},
			"list-pipeline-endpoint-connections":   {},
			"list-pipeline-endpoints":              {},
			"list-pipelines":                       {},
			"list-tags-for-resource":               {"Arn"},
			"put-resource-policy":                  {"Policy", "ResourceArn"},
			"revoke-pipeline-endpoint-connections": {"EndpointIds", "PipelineArn"},
			"start-pipeline":                       {"PipelineName"},
			"stop-pipeline":                        {"PipelineName"},
			"tag-resource":                         {"Arn", "Tags"},
			"untag-resource":                       {"Arn", "TagKeys"},
			"update-pipeline":                      {"PipelineName"},
			"validate-pipeline":                    {"PipelineConfigurationBody"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("osis", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
