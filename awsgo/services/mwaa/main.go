package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/mwaa/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-cli-token", "create-environment", "create-web-login-token", "delete-environment", "get-environment", "invoke-rest-api", "list-environments", "list-tags-for-resource", "publish-metrics", "tag-resource", "untag-resource", "update-environment"},
		OperationSet: map[string]bool{"create-cli-token": true, "create-environment": true, "create-web-login-token": true, "delete-environment": true, "get-environment": true, "invoke-rest-api": true, "list-environments": true, "list-tags-for-resource": true, "publish-metrics": true, "tag-resource": true, "untag-resource": true, "update-environment": true},
		OperationInputs: map[string][]string{
			"create-cli-token":       {"Name"},
			"create-environment":     {"AirflowConfigurationOptions", "AirflowVersion", "DagS3Path", "EndpointManagement", "EnvironmentClass", "ExecutionRoleArn", "KmsKey", "LoggingConfiguration", "MaxWebservers", "MaxWorkers", "MinWebservers", "MinWorkers", "Name", "NetworkConfiguration", "PluginsS3ObjectVersion", "PluginsS3Path", "RequirementsS3ObjectVersion", "RequirementsS3Path", "Schedulers", "SourceBucketArn", "StartupScriptS3ObjectVersion", "StartupScriptS3Path", "Tags", "WebserverAccessMode", "WeeklyMaintenanceWindowStart"},
			"create-web-login-token": {"Name"},
			"delete-environment":     {"Name"},
			"get-environment":        {"Name"},
			"invoke-rest-api":        {"Body", "Method", "Name", "Path", "QueryParameters"},
			"list-environments":      {"MaxResults", "NextToken"},
			"list-tags-for-resource": {"ResourceArn"},
			"publish-metrics":        {"EnvironmentName", "MetricData"},
			"tag-resource":           {"ResourceArn", "Tags"},
			"untag-resource":         {"ResourceArn", "TagKeys"},
			"update-environment":     {"AirflowConfigurationOptions", "AirflowVersion", "DagS3Path", "EnvironmentClass", "ExecutionRoleArn", "LoggingConfiguration", "MaxWebservers", "MaxWorkers", "MinWebservers", "MinWorkers", "Name", "NetworkConfiguration", "PluginsS3ObjectVersion", "PluginsS3Path", "RequirementsS3ObjectVersion", "RequirementsS3Path", "Schedulers", "SourceBucketArn", "StartupScriptS3ObjectVersion", "StartupScriptS3Path", "WebserverAccessMode", "WeeklyMaintenanceWindowStart", "WorkerReplacementStrategy"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-cli-token":       {"Name": "*string"},
			"create-environment":     {"AirflowConfigurationOptions": "map[string]string", "AirflowVersion": "*string", "DagS3Path": "*string", "EndpointManagement": "types.EndpointManagement", "EnvironmentClass": "*string", "ExecutionRoleArn": "*string", "KmsKey": "*string", "LoggingConfiguration": "*types.LoggingConfigurationInput", "MaxWebservers": "*int32", "MaxWorkers": "*int32", "MinWebservers": "*int32", "MinWorkers": "*int32", "Name": "*string", "NetworkConfiguration": "*types.NetworkConfiguration", "PluginsS3ObjectVersion": "*string", "PluginsS3Path": "*string", "RequirementsS3ObjectVersion": "*string", "RequirementsS3Path": "*string", "Schedulers": "*int32", "SourceBucketArn": "*string", "StartupScriptS3ObjectVersion": "*string", "StartupScriptS3Path": "*string", "Tags": "map[string]string", "WebserverAccessMode": "types.WebserverAccessMode", "WeeklyMaintenanceWindowStart": "*string"},
			"create-web-login-token": {"Name": "*string"},
			"delete-environment":     {"Name": "*string"},
			"get-environment":        {"Name": "*string"},
			"invoke-rest-api":        {"Body": "document.Interface", "Method": "types.RestApiMethod", "Name": "*string", "Path": "*string", "QueryParameters": "document.Interface"},
			"list-environments":      {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource": {"ResourceArn": "*string"},
			"publish-metrics":        {"EnvironmentName": "*string", "MetricData": "[]types.MetricDatum"},
			"tag-resource":           {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":         {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-environment":     {"AirflowConfigurationOptions": "map[string]string", "AirflowVersion": "*string", "DagS3Path": "*string", "EnvironmentClass": "*string", "ExecutionRoleArn": "*string", "LoggingConfiguration": "*types.LoggingConfigurationInput", "MaxWebservers": "*int32", "MaxWorkers": "*int32", "MinWebservers": "*int32", "MinWorkers": "*int32", "Name": "*string", "NetworkConfiguration": "*types.UpdateNetworkConfigurationInput", "PluginsS3ObjectVersion": "*string", "PluginsS3Path": "*string", "RequirementsS3ObjectVersion": "*string", "RequirementsS3Path": "*string", "Schedulers": "*int32", "SourceBucketArn": "*string", "StartupScriptS3ObjectVersion": "*string", "StartupScriptS3Path": "*string", "WebserverAccessMode": "types.WebserverAccessMode", "WeeklyMaintenanceWindowStart": "*string", "WorkerReplacementStrategy": "types.WorkerReplacementStrategy"},
		},
		OperationInputRequired: map[string][]string{
			"create-cli-token":       {"Name"},
			"create-environment":     {"DagS3Path", "ExecutionRoleArn", "Name", "NetworkConfiguration", "SourceBucketArn"},
			"create-web-login-token": {"Name"},
			"delete-environment":     {"Name"},
			"get-environment":        {"Name"},
			"invoke-rest-api":        {"Method", "Name", "Path"},
			"list-environments":      {},
			"list-tags-for-resource": {"ResourceArn"},
			"publish-metrics":        {"EnvironmentName", "MetricData"},
			"tag-resource":           {"ResourceArn", "Tags"},
			"untag-resource":         {"ResourceArn", "TagKeys"},
			"update-environment":     {"Name"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("mwaa", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
