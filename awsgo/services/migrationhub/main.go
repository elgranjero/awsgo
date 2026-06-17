package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/migrationhub/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"associate-created-artifact", "associate-discovered-resource", "associate-source-resource", "create-progress-update-stream", "delete-progress-update-stream", "describe-application-state", "describe-migration-task", "disassociate-created-artifact", "disassociate-discovered-resource", "disassociate-source-resource", "import-migration-task", "list-application-states", "list-created-artifacts", "list-discovered-resources", "list-migration-task-updates", "list-migration-tasks", "list-progress-update-streams", "list-source-resources", "notify-application-state", "notify-migration-task-state", "put-resource-attributes"},
		OperationSet: map[string]bool{"associate-created-artifact": true, "associate-discovered-resource": true, "associate-source-resource": true, "create-progress-update-stream": true, "delete-progress-update-stream": true, "describe-application-state": true, "describe-migration-task": true, "disassociate-created-artifact": true, "disassociate-discovered-resource": true, "disassociate-source-resource": true, "import-migration-task": true, "list-application-states": true, "list-created-artifacts": true, "list-discovered-resources": true, "list-migration-task-updates": true, "list-migration-tasks": true, "list-progress-update-streams": true, "list-source-resources": true, "notify-application-state": true, "notify-migration-task-state": true, "put-resource-attributes": true},
		OperationInputs: map[string][]string{
			"associate-created-artifact":       {"CreatedArtifact", "DryRun", "MigrationTaskName", "ProgressUpdateStream"},
			"associate-discovered-resource":    {"DiscoveredResource", "DryRun", "MigrationTaskName", "ProgressUpdateStream"},
			"associate-source-resource":        {"DryRun", "MigrationTaskName", "ProgressUpdateStream", "SourceResource"},
			"create-progress-update-stream":    {"DryRun", "ProgressUpdateStreamName"},
			"delete-progress-update-stream":    {"DryRun", "ProgressUpdateStreamName"},
			"describe-application-state":       {"ApplicationId"},
			"describe-migration-task":          {"MigrationTaskName", "ProgressUpdateStream"},
			"disassociate-created-artifact":    {"CreatedArtifactName", "DryRun", "MigrationTaskName", "ProgressUpdateStream"},
			"disassociate-discovered-resource": {"ConfigurationId", "DryRun", "MigrationTaskName", "ProgressUpdateStream"},
			"disassociate-source-resource":     {"DryRun", "MigrationTaskName", "ProgressUpdateStream", "SourceResourceName"},
			"import-migration-task":            {"DryRun", "MigrationTaskName", "ProgressUpdateStream"},
			"list-application-states":          {"ApplicationIds", "MaxResults", "NextToken"},
			"list-created-artifacts":           {"MaxResults", "MigrationTaskName", "NextToken", "ProgressUpdateStream"},
			"list-discovered-resources":        {"MaxResults", "MigrationTaskName", "NextToken", "ProgressUpdateStream"},
			"list-migration-task-updates":      {"MaxResults", "MigrationTaskName", "NextToken", "ProgressUpdateStream"},
			"list-migration-tasks":             {"MaxResults", "NextToken", "ResourceName"},
			"list-progress-update-streams":     {"MaxResults", "NextToken"},
			"list-source-resources":            {"MaxResults", "MigrationTaskName", "NextToken", "ProgressUpdateStream"},
			"notify-application-state":         {"ApplicationId", "DryRun", "Status", "UpdateDateTime"},
			"notify-migration-task-state":      {"DryRun", "MigrationTaskName", "NextUpdateSeconds", "ProgressUpdateStream", "Task", "UpdateDateTime"},
			"put-resource-attributes":          {"DryRun", "MigrationTaskName", "ProgressUpdateStream", "ResourceAttributeList"},
		},
		OperationInputTypes: map[string]map[string]string{
			"associate-created-artifact":       {"CreatedArtifact": "*types.CreatedArtifact", "DryRun": "bool", "MigrationTaskName": "*string", "ProgressUpdateStream": "*string"},
			"associate-discovered-resource":    {"DiscoveredResource": "*types.DiscoveredResource", "DryRun": "bool", "MigrationTaskName": "*string", "ProgressUpdateStream": "*string"},
			"associate-source-resource":        {"DryRun": "bool", "MigrationTaskName": "*string", "ProgressUpdateStream": "*string", "SourceResource": "*types.SourceResource"},
			"create-progress-update-stream":    {"DryRun": "bool", "ProgressUpdateStreamName": "*string"},
			"delete-progress-update-stream":    {"DryRun": "bool", "ProgressUpdateStreamName": "*string"},
			"describe-application-state":       {"ApplicationId": "*string"},
			"describe-migration-task":          {"MigrationTaskName": "*string", "ProgressUpdateStream": "*string"},
			"disassociate-created-artifact":    {"CreatedArtifactName": "*string", "DryRun": "bool", "MigrationTaskName": "*string", "ProgressUpdateStream": "*string"},
			"disassociate-discovered-resource": {"ConfigurationId": "*string", "DryRun": "bool", "MigrationTaskName": "*string", "ProgressUpdateStream": "*string"},
			"disassociate-source-resource":     {"DryRun": "bool", "MigrationTaskName": "*string", "ProgressUpdateStream": "*string", "SourceResourceName": "*string"},
			"import-migration-task":            {"DryRun": "bool", "MigrationTaskName": "*string", "ProgressUpdateStream": "*string"},
			"list-application-states":          {"ApplicationIds": "[]string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-created-artifacts":           {"MaxResults": "*int32", "MigrationTaskName": "*string", "NextToken": "*string", "ProgressUpdateStream": "*string"},
			"list-discovered-resources":        {"MaxResults": "*int32", "MigrationTaskName": "*string", "NextToken": "*string", "ProgressUpdateStream": "*string"},
			"list-migration-task-updates":      {"MaxResults": "*int32", "MigrationTaskName": "*string", "NextToken": "*string", "ProgressUpdateStream": "*string"},
			"list-migration-tasks":             {"MaxResults": "*int32", "NextToken": "*string", "ResourceName": "*string"},
			"list-progress-update-streams":     {"MaxResults": "*int32", "NextToken": "*string"},
			"list-source-resources":            {"MaxResults": "*int32", "MigrationTaskName": "*string", "NextToken": "*string", "ProgressUpdateStream": "*string"},
			"notify-application-state":         {"ApplicationId": "*string", "DryRun": "bool", "Status": "types.ApplicationStatus", "UpdateDateTime": "*time.Time"},
			"notify-migration-task-state":      {"DryRun": "bool", "MigrationTaskName": "*string", "NextUpdateSeconds": "int32", "ProgressUpdateStream": "*string", "Task": "*types.Task", "UpdateDateTime": "*time.Time"},
			"put-resource-attributes":          {"DryRun": "bool", "MigrationTaskName": "*string", "ProgressUpdateStream": "*string", "ResourceAttributeList": "[]types.ResourceAttribute"},
		},
		OperationInputRequired: map[string][]string{
			"associate-created-artifact":       {"CreatedArtifact", "MigrationTaskName", "ProgressUpdateStream"},
			"associate-discovered-resource":    {"DiscoveredResource", "MigrationTaskName", "ProgressUpdateStream"},
			"associate-source-resource":        {"MigrationTaskName", "ProgressUpdateStream", "SourceResource"},
			"create-progress-update-stream":    {"ProgressUpdateStreamName"},
			"delete-progress-update-stream":    {"ProgressUpdateStreamName"},
			"describe-application-state":       {"ApplicationId"},
			"describe-migration-task":          {"MigrationTaskName", "ProgressUpdateStream"},
			"disassociate-created-artifact":    {"CreatedArtifactName", "MigrationTaskName", "ProgressUpdateStream"},
			"disassociate-discovered-resource": {"ConfigurationId", "MigrationTaskName", "ProgressUpdateStream"},
			"disassociate-source-resource":     {"MigrationTaskName", "ProgressUpdateStream", "SourceResourceName"},
			"import-migration-task":            {"MigrationTaskName", "ProgressUpdateStream"},
			"list-application-states":          {},
			"list-created-artifacts":           {"MigrationTaskName", "ProgressUpdateStream"},
			"list-discovered-resources":        {"MigrationTaskName", "ProgressUpdateStream"},
			"list-migration-task-updates":      {"MigrationTaskName", "ProgressUpdateStream"},
			"list-migration-tasks":             {},
			"list-progress-update-streams":     {},
			"list-source-resources":            {"MigrationTaskName", "ProgressUpdateStream"},
			"notify-application-state":         {"ApplicationId", "Status"},
			"notify-migration-task-state":      {"MigrationTaskName", "NextUpdateSeconds", "ProgressUpdateStream", "Task", "UpdateDateTime"},
			"put-resource-attributes":          {"MigrationTaskName", "ProgressUpdateStream", "ResourceAttributeList"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("migrationhub", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
