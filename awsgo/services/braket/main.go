package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/braket/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"cancel-job", "cancel-quantum-task", "create-job", "create-quantum-task", "create-spending-limit", "delete-spending-limit", "get-device", "get-job", "get-quantum-task", "list-tags-for-resource", "search-devices", "search-jobs", "search-quantum-tasks", "search-spending-limits", "tag-resource", "untag-resource", "update-spending-limit"},
		OperationSet: map[string]bool{"cancel-job": true, "cancel-quantum-task": true, "create-job": true, "create-quantum-task": true, "create-spending-limit": true, "delete-spending-limit": true, "get-device": true, "get-job": true, "get-quantum-task": true, "list-tags-for-resource": true, "search-devices": true, "search-jobs": true, "search-quantum-tasks": true, "search-spending-limits": true, "tag-resource": true, "untag-resource": true, "update-spending-limit": true},
		OperationInputs: map[string][]string{
			"cancel-job":             {"JobArn"},
			"cancel-quantum-task":    {"ClientToken", "QuantumTaskArn"},
			"create-job":             {"AlgorithmSpecification", "Associations", "CheckpointConfig", "ClientToken", "DeviceConfig", "HyperParameters", "InputDataConfig", "InstanceConfig", "JobName", "OutputDataConfig", "RoleArn", "StoppingCondition", "Tags"},
			"create-quantum-task":    {"Action", "Associations", "ClientToken", "DeviceArn", "DeviceParameters", "ExperimentalCapabilities", "JobToken", "OutputS3Bucket", "OutputS3KeyPrefix", "Shots", "Tags"},
			"create-spending-limit":  {"ClientToken", "DeviceArn", "SpendingLimit", "Tags", "TimePeriod"},
			"delete-spending-limit":  {"SpendingLimitArn"},
			"get-device":             {"DeviceArn"},
			"get-job":                {"AdditionalAttributeNames", "JobArn"},
			"get-quantum-task":       {"AdditionalAttributeNames", "QuantumTaskArn"},
			"list-tags-for-resource": {"ResourceArn"},
			"search-devices":         {"Filters", "MaxResults", "NextToken"},
			"search-jobs":            {"Filters", "MaxResults", "NextToken"},
			"search-quantum-tasks":   {"Filters", "MaxResults", "NextToken"},
			"search-spending-limits": {"Filters", "MaxResults", "NextToken"},
			"tag-resource":           {"ResourceArn", "Tags"},
			"untag-resource":         {"ResourceArn", "TagKeys"},
			"update-spending-limit":  {"ClientToken", "SpendingLimit", "SpendingLimitArn", "TimePeriod"},
		},
		OperationInputTypes: map[string]map[string]string{
			"cancel-job":             {"JobArn": "*string"},
			"cancel-quantum-task":    {"ClientToken": "*string", "QuantumTaskArn": "*string"},
			"create-job":             {"AlgorithmSpecification": "*types.AlgorithmSpecification", "Associations": "[]types.Association", "CheckpointConfig": "*types.JobCheckpointConfig", "ClientToken": "*string", "DeviceConfig": "*types.DeviceConfig", "HyperParameters": "map[string]string", "InputDataConfig": "[]types.InputFileConfig", "InstanceConfig": "*types.InstanceConfig", "JobName": "*string", "OutputDataConfig": "*types.JobOutputDataConfig", "RoleArn": "*string", "StoppingCondition": "*types.JobStoppingCondition", "Tags": "map[string]string"},
			"create-quantum-task":    {"Action": "*string", "Associations": "[]types.Association", "ClientToken": "*string", "DeviceArn": "*string", "DeviceParameters": "*string", "ExperimentalCapabilities": "types.ExperimentalCapabilities", "JobToken": "*string", "OutputS3Bucket": "*string", "OutputS3KeyPrefix": "*string", "Shots": "*int64", "Tags": "map[string]string"},
			"create-spending-limit":  {"ClientToken": "*string", "DeviceArn": "*string", "SpendingLimit": "*string", "Tags": "map[string]string", "TimePeriod": "*types.TimePeriod"},
			"delete-spending-limit":  {"SpendingLimitArn": "*string"},
			"get-device":             {"DeviceArn": "*string"},
			"get-job":                {"AdditionalAttributeNames": "[]types.HybridJobAdditionalAttributeName", "JobArn": "*string"},
			"get-quantum-task":       {"AdditionalAttributeNames": "[]types.QuantumTaskAdditionalAttributeName", "QuantumTaskArn": "*string"},
			"list-tags-for-resource": {"ResourceArn": "*string"},
			"search-devices":         {"Filters": "[]types.SearchDevicesFilter", "MaxResults": "*int32", "NextToken": "*string"},
			"search-jobs":            {"Filters": "[]types.SearchJobsFilter", "MaxResults": "*int32", "NextToken": "*string"},
			"search-quantum-tasks":   {"Filters": "[]types.SearchQuantumTasksFilter", "MaxResults": "*int32", "NextToken": "*string"},
			"search-spending-limits": {"Filters": "[]types.SearchSpendingLimitsFilter", "MaxResults": "*int32", "NextToken": "*string"},
			"tag-resource":           {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":         {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-spending-limit":  {"ClientToken": "*string", "SpendingLimit": "*string", "SpendingLimitArn": "*string", "TimePeriod": "*types.TimePeriod"},
		},
		OperationInputRequired: map[string][]string{
			"cancel-job":             {"JobArn"},
			"cancel-quantum-task":    {"ClientToken", "QuantumTaskArn"},
			"create-job":             {"AlgorithmSpecification", "ClientToken", "DeviceConfig", "InstanceConfig", "JobName", "OutputDataConfig", "RoleArn"},
			"create-quantum-task":    {"Action", "ClientToken", "DeviceArn", "OutputS3Bucket", "OutputS3KeyPrefix", "Shots"},
			"create-spending-limit":  {"ClientToken", "DeviceArn", "SpendingLimit"},
			"delete-spending-limit":  {"SpendingLimitArn"},
			"get-device":             {"DeviceArn"},
			"get-job":                {"JobArn"},
			"get-quantum-task":       {"QuantumTaskArn"},
			"list-tags-for-resource": {"ResourceArn"},
			"search-devices":         {"Filters"},
			"search-jobs":            {"Filters"},
			"search-quantum-tasks":   {"Filters"},
			"search-spending-limits": {},
			"tag-resource":           {"ResourceArn", "Tags"},
			"untag-resource":         {"ResourceArn", "TagKeys"},
			"update-spending-limit":  {"ClientToken", "SpendingLimitArn"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("braket", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
