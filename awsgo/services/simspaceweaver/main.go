package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/simspaceweaver/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-snapshot", "delete-app", "delete-simulation", "describe-app", "describe-simulation", "list-apps", "list-simulations", "list-tags-for-resource", "start-app", "start-clock", "start-simulation", "stop-app", "stop-clock", "stop-simulation", "tag-resource", "untag-resource"},
		OperationSet: map[string]bool{"create-snapshot": true, "delete-app": true, "delete-simulation": true, "describe-app": true, "describe-simulation": true, "list-apps": true, "list-simulations": true, "list-tags-for-resource": true, "start-app": true, "start-clock": true, "start-simulation": true, "stop-app": true, "stop-clock": true, "stop-simulation": true, "tag-resource": true, "untag-resource": true},
		OperationInputs: map[string][]string{
			"create-snapshot":        {"Destination", "Simulation"},
			"delete-app":             {"App", "Domain", "Simulation"},
			"delete-simulation":      {"Simulation"},
			"describe-app":           {"App", "Domain", "Simulation"},
			"describe-simulation":    {"Simulation"},
			"list-apps":              {"Domain", "MaxResults", "NextToken", "Simulation"},
			"list-simulations":       {"MaxResults", "NextToken"},
			"list-tags-for-resource": {"ResourceArn"},
			"start-app":              {"ClientToken", "Description", "Domain", "LaunchOverrides", "Name", "Simulation"},
			"start-clock":            {"Simulation"},
			"start-simulation":       {"ClientToken", "Description", "MaximumDuration", "Name", "RoleArn", "SchemaS3Location", "SnapshotS3Location", "Tags"},
			"stop-app":               {"App", "Domain", "Simulation"},
			"stop-clock":             {"Simulation"},
			"stop-simulation":        {"Simulation"},
			"tag-resource":           {"ResourceArn", "Tags"},
			"untag-resource":         {"ResourceArn", "TagKeys"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-snapshot":        {"Destination": "*types.S3Destination", "Simulation": "*string"},
			"delete-app":             {"App": "*string", "Domain": "*string", "Simulation": "*string"},
			"delete-simulation":      {"Simulation": "*string"},
			"describe-app":           {"App": "*string", "Domain": "*string", "Simulation": "*string"},
			"describe-simulation":    {"Simulation": "*string"},
			"list-apps":              {"Domain": "*string", "MaxResults": "*int32", "NextToken": "*string", "Simulation": "*string"},
			"list-simulations":       {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource": {"ResourceArn": "*string"},
			"start-app":              {"ClientToken": "*string", "Description": "*string", "Domain": "*string", "LaunchOverrides": "*types.LaunchOverrides", "Name": "*string", "Simulation": "*string"},
			"start-clock":            {"Simulation": "*string"},
			"start-simulation":       {"ClientToken": "*string", "Description": "*string", "MaximumDuration": "*string", "Name": "*string", "RoleArn": "*string", "SchemaS3Location": "*types.S3Location", "SnapshotS3Location": "*types.S3Location", "Tags": "map[string]string"},
			"stop-app":               {"App": "*string", "Domain": "*string", "Simulation": "*string"},
			"stop-clock":             {"Simulation": "*string"},
			"stop-simulation":        {"Simulation": "*string"},
			"tag-resource":           {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":         {"ResourceArn": "*string", "TagKeys": "[]string"},
		},
		OperationInputRequired: map[string][]string{
			"create-snapshot":        {"Destination", "Simulation"},
			"delete-app":             {"App", "Domain", "Simulation"},
			"delete-simulation":      {"Simulation"},
			"describe-app":           {"App", "Domain", "Simulation"},
			"describe-simulation":    {"Simulation"},
			"list-apps":              {"Simulation"},
			"list-simulations":       {},
			"list-tags-for-resource": {"ResourceArn"},
			"start-app":              {"Domain", "Name", "Simulation"},
			"start-clock":            {"Simulation"},
			"start-simulation":       {"Name", "RoleArn"},
			"stop-app":               {"App", "Domain", "Simulation"},
			"stop-clock":             {"Simulation"},
			"stop-simulation":        {"Simulation"},
			"tag-resource":           {"ResourceArn", "Tags"},
			"untag-resource":         {"ResourceArn", "TagKeys"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("simspaceweaver", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
