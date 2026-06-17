package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/iotjobsdataplane/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"describe-job-execution", "get-pending-job-executions", "start-command-execution", "start-next-pending-job-execution", "update-job-execution"},
		OperationSet: map[string]bool{"describe-job-execution": true, "get-pending-job-executions": true, "start-command-execution": true, "start-next-pending-job-execution": true, "update-job-execution": true},
		OperationInputs: map[string][]string{
			"describe-job-execution":           {"ExecutionNumber", "IncludeJobDocument", "JobId", "ThingName"},
			"get-pending-job-executions":       {"ThingName"},
			"start-command-execution":          {"ClientToken", "CommandArn", "ExecutionTimeoutSeconds", "Parameters", "TargetArn"},
			"start-next-pending-job-execution": {"StatusDetails", "StepTimeoutInMinutes", "ThingName"},
			"update-job-execution":             {"ExecutionNumber", "ExpectedVersion", "IncludeJobDocument", "IncludeJobExecutionState", "JobId", "Status", "StatusDetails", "StepTimeoutInMinutes", "ThingName"},
		},
		OperationInputTypes: map[string]map[string]string{
			"describe-job-execution":           {"ExecutionNumber": "*int64", "IncludeJobDocument": "*bool", "JobId": "*string", "ThingName": "*string"},
			"get-pending-job-executions":       {"ThingName": "*string"},
			"start-command-execution":          {"ClientToken": "*string", "CommandArn": "*string", "ExecutionTimeoutSeconds": "*int64", "Parameters": "map[string]types.CommandParameterValue", "TargetArn": "*string"},
			"start-next-pending-job-execution": {"StatusDetails": "map[string]string", "StepTimeoutInMinutes": "*int64", "ThingName": "*string"},
			"update-job-execution":             {"ExecutionNumber": "*int64", "ExpectedVersion": "*int64", "IncludeJobDocument": "*bool", "IncludeJobExecutionState": "*bool", "JobId": "*string", "Status": "types.JobExecutionStatus", "StatusDetails": "map[string]string", "StepTimeoutInMinutes": "*int64", "ThingName": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"describe-job-execution":           {"JobId", "ThingName"},
			"get-pending-job-executions":       {"ThingName"},
			"start-command-execution":          {"CommandArn", "TargetArn"},
			"start-next-pending-job-execution": {"ThingName"},
			"update-job-execution":             {"JobId", "Status", "ThingName"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("iotjobsdataplane", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
