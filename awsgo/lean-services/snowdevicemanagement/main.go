package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/snowdevicemanagement"
)

var fields_cancel_task = []leanruntime.Field{
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: true},
}

var fields_create_task = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Command", Flag: "command", Type: "types.Command", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Targets", Flag: "targets", Type: "[]string", Required: true},
}

var fields_describe_device = []leanruntime.Field{
	{Name: "ManagedDeviceId", Flag: "managed-device-id", Type: "*string", Required: true},
}

var fields_describe_device_ec2_instances = []leanruntime.Field{
	{Name: "InstanceIds", Flag: "instance-ids", Type: "[]string", Required: true},
	{Name: "ManagedDeviceId", Flag: "managed-device-id", Type: "*string", Required: true},
}

var fields_describe_execution = []leanruntime.Field{
	{Name: "ManagedDeviceId", Flag: "managed-device-id", Type: "*string", Required: true},
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: true},
}

var fields_describe_task = []leanruntime.Field{
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: true},
}

var fields_list_device_resources = []leanruntime.Field{
	{Name: "ManagedDeviceId", Flag: "managed-device-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "*string", Required: false},
}

var fields_list_devices = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_executions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "State", Flag: "state", Type: "types.ExecutionState", Required: false},
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_tasks = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "State", Flag: "state", Type: "types.TaskState", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"cancel-task": {
			Name:   "cancel-task",
			Fields: fields_cancel_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelTask(ctx, input)
			},
		},
		"create-task": {
			Name:   "create-task",
			Fields: fields_create_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTask(ctx, input)
			},
		},
		"describe-device": {
			Name:   "describe-device",
			Fields: fields_describe_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDevice(ctx, input)
			},
		},
		"describe-device-ec2-instances": {
			Name:   "describe-device-ec2-instances",
			Fields: fields_describe_device_ec2_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDeviceEc2InstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_device_ec2_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDeviceEc2Instances(ctx, input)
			},
		},
		"describe-execution": {
			Name:   "describe-execution",
			Fields: fields_describe_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeExecution(ctx, input)
			},
		},
		"describe-task": {
			Name:   "describe-task",
			Fields: fields_describe_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTask(ctx, input)
			},
		},
		"list-device-resources": {
			Name:   "list-device-resources",
			Fields: fields_list_device_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDeviceResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_device_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDeviceResources(ctx, input)
				}
				var results []*svc.ListDeviceResourcesOutput
				p := svc.NewListDeviceResourcesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-devices": {
			Name:   "list-devices",
			Fields: fields_list_devices,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDevicesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_devices, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDevices(ctx, input)
				}
				var results []*svc.ListDevicesOutput
				p := svc.NewListDevicesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-executions": {
			Name:   "list-executions",
			Fields: fields_list_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListExecutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_executions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListExecutions(ctx, input)
				}
				var results []*svc.ListExecutionsOutput
				p := svc.NewListExecutionsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-tags-for-resource": {
			Name:   "list-tags-for-resource",
			Fields: fields_list_tags_for_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsForResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_tags_for_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTagsForResource(ctx, input)
			},
		},
		"list-tasks": {
			Name:   "list-tasks",
			Fields: fields_list_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTasksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tasks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTasks(ctx, input)
				}
				var results []*svc.ListTasksOutput
				p := svc.NewListTasksPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"tag-resource": {
			Name:   "tag-resource",
			Fields: fields_tag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_tag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TagResource(ctx, input)
			},
		},
		"untag-resource": {
			Name:   "untag-resource",
			Fields: fields_untag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UntagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_untag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UntagResource(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("snowdevicemanagement", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
