package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/pipes/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-pipe", "delete-pipe", "describe-pipe", "list-pipes", "list-tags-for-resource", "start-pipe", "stop-pipe", "tag-resource", "untag-resource", "update-pipe"},
		OperationSet: map[string]bool{"create-pipe": true, "delete-pipe": true, "describe-pipe": true, "list-pipes": true, "list-tags-for-resource": true, "start-pipe": true, "stop-pipe": true, "tag-resource": true, "untag-resource": true, "update-pipe": true},
		OperationInputs: map[string][]string{
			"create-pipe":            {"Description", "DesiredState", "Enrichment", "EnrichmentParameters", "KmsKeyIdentifier", "LogConfiguration", "Name", "RoleArn", "Source", "SourceParameters", "Tags", "Target", "TargetParameters"},
			"delete-pipe":            {"Name"},
			"describe-pipe":          {"Name"},
			"list-pipes":             {"CurrentState", "DesiredState", "Limit", "NamePrefix", "NextToken", "SourcePrefix", "TargetPrefix"},
			"list-tags-for-resource": {"ResourceArn"},
			"start-pipe":             {"Name"},
			"stop-pipe":              {"Name"},
			"tag-resource":           {"ResourceArn", "Tags"},
			"untag-resource":         {"ResourceArn", "TagKeys"},
			"update-pipe":            {"Description", "DesiredState", "Enrichment", "EnrichmentParameters", "KmsKeyIdentifier", "LogConfiguration", "Name", "RoleArn", "SourceParameters", "Target", "TargetParameters"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-pipe":            {"Description": "*string", "DesiredState": "types.RequestedPipeState", "Enrichment": "*string", "EnrichmentParameters": "*types.PipeEnrichmentParameters", "KmsKeyIdentifier": "*string", "LogConfiguration": "*types.PipeLogConfigurationParameters", "Name": "*string", "RoleArn": "*string", "Source": "*string", "SourceParameters": "*types.PipeSourceParameters", "Tags": "map[string]string", "Target": "*string", "TargetParameters": "*types.PipeTargetParameters"},
			"delete-pipe":            {"Name": "*string"},
			"describe-pipe":          {"Name": "*string"},
			"list-pipes":             {"CurrentState": "types.PipeState", "DesiredState": "types.RequestedPipeState", "Limit": "*int32", "NamePrefix": "*string", "NextToken": "*string", "SourcePrefix": "*string", "TargetPrefix": "*string"},
			"list-tags-for-resource": {"ResourceArn": "*string"},
			"start-pipe":             {"Name": "*string"},
			"stop-pipe":              {"Name": "*string"},
			"tag-resource":           {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":         {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-pipe":            {"Description": "*string", "DesiredState": "types.RequestedPipeState", "Enrichment": "*string", "EnrichmentParameters": "*types.PipeEnrichmentParameters", "KmsKeyIdentifier": "*string", "LogConfiguration": "*types.PipeLogConfigurationParameters", "Name": "*string", "RoleArn": "*string", "SourceParameters": "*types.UpdatePipeSourceParameters", "Target": "*string", "TargetParameters": "*types.PipeTargetParameters"},
		},
		OperationInputRequired: map[string][]string{
			"create-pipe":            {"Name", "RoleArn", "Source", "Target"},
			"delete-pipe":            {"Name"},
			"describe-pipe":          {"Name"},
			"list-pipes":             {},
			"list-tags-for-resource": {"ResourceArn"},
			"start-pipe":             {"Name"},
			"stop-pipe":              {"Name"},
			"tag-resource":           {"ResourceArn", "Tags"},
			"untag-resource":         {"ResourceArn", "TagKeys"},
			"update-pipe":            {"Name", "RoleArn"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("pipes", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
