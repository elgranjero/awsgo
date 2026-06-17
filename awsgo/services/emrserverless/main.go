package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/emrserverless/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"cancel-job-run", "create-application", "delete-application", "get-application", "get-dashboard-for-job-run", "get-job-run", "list-applications", "list-job-run-attempts", "list-job-runs", "list-tags-for-resource", "start-application", "start-job-run", "stop-application", "tag-resource", "untag-resource", "update-application"},
		OperationSet: map[string]bool{"cancel-job-run": true, "create-application": true, "delete-application": true, "get-application": true, "get-dashboard-for-job-run": true, "get-job-run": true, "list-applications": true, "list-job-run-attempts": true, "list-job-runs": true, "list-tags-for-resource": true, "start-application": true, "start-job-run": true, "stop-application": true, "tag-resource": true, "untag-resource": true, "update-application": true},
		OperationInputs: map[string][]string{
			"cancel-job-run":            {"ApplicationId", "JobRunId", "ShutdownGracePeriodInSeconds"},
			"create-application":        {"Architecture", "AutoStartConfiguration", "AutoStopConfiguration", "ClientToken", "DiskEncryptionConfiguration", "IdentityCenterConfiguration", "ImageConfiguration", "InitialCapacity", "InteractiveConfiguration", "JobLevelCostAllocationConfiguration", "MaximumCapacity", "MonitoringConfiguration", "Name", "NetworkConfiguration", "ReleaseLabel", "RuntimeConfiguration", "SchedulerConfiguration", "Tags", "Type", "WorkerTypeSpecifications"},
			"delete-application":        {"ApplicationId"},
			"get-application":           {"ApplicationId"},
			"get-dashboard-for-job-run": {"AccessSystemProfileLogs", "ApplicationId", "Attempt", "JobRunId"},
			"get-job-run":               {"ApplicationId", "Attempt", "JobRunId"},
			"list-applications":         {"MaxResults", "NextToken", "States"},
			"list-job-run-attempts":     {"ApplicationId", "JobRunId", "MaxResults", "NextToken"},
			"list-job-runs":             {"ApplicationId", "CreatedAtAfter", "CreatedAtBefore", "MaxResults", "Mode", "NextToken", "States"},
			"list-tags-for-resource":    {"ResourceArn"},
			"start-application":         {"ApplicationId"},
			"start-job-run":             {"ApplicationId", "ClientToken", "ConfigurationOverrides", "ExecutionIamPolicy", "ExecutionRoleArn", "ExecutionTimeoutMinutes", "JobDriver", "Mode", "Name", "RetryPolicy", "Tags"},
			"stop-application":          {"ApplicationId"},
			"tag-resource":              {"ResourceArn", "Tags"},
			"untag-resource":            {"ResourceArn", "TagKeys"},
			"update-application":        {"ApplicationId", "Architecture", "AutoStartConfiguration", "AutoStopConfiguration", "ClientToken", "DiskEncryptionConfiguration", "IdentityCenterConfiguration", "ImageConfiguration", "InitialCapacity", "InteractiveConfiguration", "JobLevelCostAllocationConfiguration", "MaximumCapacity", "MonitoringConfiguration", "NetworkConfiguration", "ReleaseLabel", "RuntimeConfiguration", "SchedulerConfiguration", "WorkerTypeSpecifications"},
		},
		OperationInputTypes: map[string]map[string]string{
			"cancel-job-run":            {"ApplicationId": "*string", "JobRunId": "*string", "ShutdownGracePeriodInSeconds": "*int32"},
			"create-application":        {"Architecture": "types.Architecture", "AutoStartConfiguration": "*types.AutoStartConfig", "AutoStopConfiguration": "*types.AutoStopConfig", "ClientToken": "*string", "DiskEncryptionConfiguration": "*types.DiskEncryptionConfiguration", "IdentityCenterConfiguration": "*types.IdentityCenterConfigurationInput", "ImageConfiguration": "*types.ImageConfigurationInput", "InitialCapacity": "map[string]types.InitialCapacityConfig", "InteractiveConfiguration": "*types.InteractiveConfiguration", "JobLevelCostAllocationConfiguration": "*types.JobLevelCostAllocationConfiguration", "MaximumCapacity": "*types.MaximumAllowedResources", "MonitoringConfiguration": "*types.MonitoringConfiguration", "Name": "*string", "NetworkConfiguration": "*types.NetworkConfiguration", "ReleaseLabel": "*string", "RuntimeConfiguration": "[]types.Configuration", "SchedulerConfiguration": "*types.SchedulerConfiguration", "Tags": "map[string]string", "Type": "*string", "WorkerTypeSpecifications": "map[string]types.WorkerTypeSpecificationInput"},
			"delete-application":        {"ApplicationId": "*string"},
			"get-application":           {"ApplicationId": "*string"},
			"get-dashboard-for-job-run": {"AccessSystemProfileLogs": "*bool", "ApplicationId": "*string", "Attempt": "*int32", "JobRunId": "*string"},
			"get-job-run":               {"ApplicationId": "*string", "Attempt": "*int32", "JobRunId": "*string"},
			"list-applications":         {"MaxResults": "*int32", "NextToken": "*string", "States": "[]types.ApplicationState"},
			"list-job-run-attempts":     {"ApplicationId": "*string", "JobRunId": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-job-runs":             {"ApplicationId": "*string", "CreatedAtAfter": "*time.Time", "CreatedAtBefore": "*time.Time", "MaxResults": "*int32", "Mode": "types.JobRunMode", "NextToken": "*string", "States": "[]types.JobRunState"},
			"list-tags-for-resource":    {"ResourceArn": "*string"},
			"start-application":         {"ApplicationId": "*string"},
			"start-job-run":             {"ApplicationId": "*string", "ClientToken": "*string", "ConfigurationOverrides": "*types.ConfigurationOverrides", "ExecutionIamPolicy": "*types.JobRunExecutionIamPolicy", "ExecutionRoleArn": "*string", "ExecutionTimeoutMinutes": "*int64", "JobDriver": "types.JobDriver", "Mode": "types.JobRunMode", "Name": "*string", "RetryPolicy": "*types.RetryPolicy", "Tags": "map[string]string"},
			"stop-application":          {"ApplicationId": "*string"},
			"tag-resource":              {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":            {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-application":        {"ApplicationId": "*string", "Architecture": "types.Architecture", "AutoStartConfiguration": "*types.AutoStartConfig", "AutoStopConfiguration": "*types.AutoStopConfig", "ClientToken": "*string", "DiskEncryptionConfiguration": "*types.DiskEncryptionConfiguration", "IdentityCenterConfiguration": "*types.IdentityCenterConfigurationInput", "ImageConfiguration": "*types.ImageConfigurationInput", "InitialCapacity": "map[string]types.InitialCapacityConfig", "InteractiveConfiguration": "*types.InteractiveConfiguration", "JobLevelCostAllocationConfiguration": "*types.JobLevelCostAllocationConfiguration", "MaximumCapacity": "*types.MaximumAllowedResources", "MonitoringConfiguration": "*types.MonitoringConfiguration", "NetworkConfiguration": "*types.NetworkConfiguration", "ReleaseLabel": "*string", "RuntimeConfiguration": "[]types.Configuration", "SchedulerConfiguration": "*types.SchedulerConfiguration", "WorkerTypeSpecifications": "map[string]types.WorkerTypeSpecificationInput"},
		},
		OperationInputRequired: map[string][]string{
			"cancel-job-run":            {"ApplicationId", "JobRunId"},
			"create-application":        {"ClientToken", "ReleaseLabel", "Type"},
			"delete-application":        {"ApplicationId"},
			"get-application":           {"ApplicationId"},
			"get-dashboard-for-job-run": {"ApplicationId", "JobRunId"},
			"get-job-run":               {"ApplicationId", "JobRunId"},
			"list-applications":         {},
			"list-job-run-attempts":     {"ApplicationId", "JobRunId"},
			"list-job-runs":             {"ApplicationId"},
			"list-tags-for-resource":    {"ResourceArn"},
			"start-application":         {"ApplicationId"},
			"start-job-run":             {"ApplicationId", "ClientToken", "ExecutionRoleArn"},
			"stop-application":          {"ApplicationId"},
			"tag-resource":              {"ResourceArn", "Tags"},
			"untag-resource":            {"ResourceArn", "TagKeys"},
			"update-application":        {"ApplicationId", "ClientToken"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("emrserverless", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
