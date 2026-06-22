package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/iotjobsdataplane"
)

var fields_describe_job_execution = []leanruntime.Field{
	{Name: "ExecutionNumber", Flag: "execution-number", Type: "*int64", Required: false},
	{Name: "IncludeJobDocument", Flag: "include-job-document", Type: "*bool", Required: false},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: true},
}

var fields_get_pending_job_executions = []leanruntime.Field{
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: true},
}

var fields_start_command_execution = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CommandArn", Flag: "command-arn", Type: "*string", Required: true},
	{Name: "ExecutionTimeoutSeconds", Flag: "execution-timeout-seconds", Type: "*int64", Required: false},
	{Name: "Parameters", Flag: "parameters", Type: "map[string]types.CommandParameterValue", Required: false},
	{Name: "TargetArn", Flag: "target-arn", Type: "*string", Required: true},
}

var fields_start_next_pending_job_execution = []leanruntime.Field{
	{Name: "StatusDetails", Flag: "status-details", Type: "map[string]string", Required: false},
	{Name: "StepTimeoutInMinutes", Flag: "step-timeout-in-minutes", Type: "*int64", Required: false},
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: true},
}

var fields_update_job_execution = []leanruntime.Field{
	{Name: "ExecutionNumber", Flag: "execution-number", Type: "*int64", Required: false},
	{Name: "ExpectedVersion", Flag: "expected-version", Type: "*int64", Required: false},
	{Name: "IncludeJobDocument", Flag: "include-job-document", Type: "*bool", Required: false},
	{Name: "IncludeJobExecutionState", Flag: "include-job-execution-state", Type: "*bool", Required: false},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.JobExecutionStatus", Required: true},
	{Name: "StatusDetails", Flag: "status-details", Type: "map[string]string", Required: false},
	{Name: "StepTimeoutInMinutes", Flag: "step-timeout-in-minutes", Type: "*int64", Required: false},
	{Name: "ThingName", Flag: "thing-name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"describe-job-execution": {
			Name:   "describe-job-execution",
			Fields: fields_describe_job_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeJobExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_job_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeJobExecution(ctx, input)
			},
		},
		"get-pending-job-executions": {
			Name:   "get-pending-job-executions",
			Fields: fields_get_pending_job_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPendingJobExecutionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_pending_job_executions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPendingJobExecutions(ctx, input)
			},
		},
		"start-command-execution": {
			Name:   "start-command-execution",
			Fields: fields_start_command_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartCommandExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_command_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartCommandExecution(ctx, input)
			},
		},
		"start-next-pending-job-execution": {
			Name:   "start-next-pending-job-execution",
			Fields: fields_start_next_pending_job_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartNextPendingJobExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_next_pending_job_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartNextPendingJobExecution(ctx, input)
			},
		},
		"update-job-execution": {
			Name:   "update-job-execution",
			Fields: fields_update_job_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateJobExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_job_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateJobExecution(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("iotjobsdataplane", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
