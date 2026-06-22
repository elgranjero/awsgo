package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/braket"
)

var fields_cancel_job = []leanruntime.Field{
	{Name: "JobArn", Flag: "job-arn", Type: "*string", Required: true},
}

var fields_cancel_quantum_task = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "QuantumTaskArn", Flag: "quantum-task-arn", Type: "*string", Required: true},
}

var fields_create_job = []leanruntime.Field{
	{Name: "AlgorithmSpecification", Flag: "algorithm-specification", Type: "*types.AlgorithmSpecification", Required: true},
	{Name: "Associations", Flag: "associations", Type: "[]types.Association", Required: false},
	{Name: "CheckpointConfig", Flag: "checkpoint-config", Type: "*types.JobCheckpointConfig", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "DeviceConfig", Flag: "device-config", Type: "*types.DeviceConfig", Required: true},
	{Name: "HyperParameters", Flag: "hyper-parameters", Type: "map[string]string", Required: false},
	{Name: "InputDataConfig", Flag: "input-data-config", Type: "[]types.InputFileConfig", Required: false},
	{Name: "InstanceConfig", Flag: "instance-config", Type: "*types.InstanceConfig", Required: true},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: true},
	{Name: "OutputDataConfig", Flag: "output-data-config", Type: "*types.JobOutputDataConfig", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "StoppingCondition", Flag: "stopping-condition", Type: "*types.JobStoppingCondition", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_quantum_task = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "*string", Required: true},
	{Name: "Associations", Flag: "associations", Type: "[]types.Association", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "DeviceArn", Flag: "device-arn", Type: "*string", Required: true},
	{Name: "DeviceParameters", Flag: "device-parameters", Type: "*string", Required: false},
	{Name: "ExperimentalCapabilities", Flag: "experimental-capabilities", Type: "types.ExperimentalCapabilities", Required: false},
	{Name: "JobToken", Flag: "job-token", Type: "*string", Required: false},
	{Name: "OutputS3Bucket", Flag: "output-s3-bucket", Type: "*string", Required: true},
	{Name: "OutputS3KeyPrefix", Flag: "output-s3-key-prefix", Type: "*string", Required: true},
	{Name: "Shots", Flag: "shots", Type: "*int64", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_spending_limit = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "DeviceArn", Flag: "device-arn", Type: "*string", Required: true},
	{Name: "SpendingLimit", Flag: "spending-limit", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TimePeriod", Flag: "time-period", Type: "*types.TimePeriod", Required: false},
}

var fields_delete_spending_limit = []leanruntime.Field{
	{Name: "SpendingLimitArn", Flag: "spending-limit-arn", Type: "*string", Required: true},
}

var fields_get_device = []leanruntime.Field{
	{Name: "DeviceArn", Flag: "device-arn", Type: "*string", Required: true},
}

var fields_get_job = []leanruntime.Field{
	{Name: "AdditionalAttributeNames", Flag: "additional-attribute-names", Type: "[]types.HybridJobAdditionalAttributeName", Required: false},
	{Name: "JobArn", Flag: "job-arn", Type: "*string", Required: true},
}

var fields_get_quantum_task = []leanruntime.Field{
	{Name: "AdditionalAttributeNames", Flag: "additional-attribute-names", Type: "[]types.QuantumTaskAdditionalAttributeName", Required: false},
	{Name: "QuantumTaskArn", Flag: "quantum-task-arn", Type: "*string", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_search_devices = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.SearchDevicesFilter", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_search_jobs = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.SearchJobsFilter", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_search_quantum_tasks = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.SearchQuantumTasksFilter", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_search_spending_limits = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.SearchSpendingLimitsFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_spending_limit = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "SpendingLimit", Flag: "spending-limit", Type: "*string", Required: false},
	{Name: "SpendingLimitArn", Flag: "spending-limit-arn", Type: "*string", Required: true},
	{Name: "TimePeriod", Flag: "time-period", Type: "*types.TimePeriod", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"cancel-job": {
			Name:   "cancel-job",
			Fields: fields_cancel_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelJob(ctx, input)
			},
		},
		"cancel-quantum-task": {
			Name:   "cancel-quantum-task",
			Fields: fields_cancel_quantum_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelQuantumTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_quantum_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelQuantumTask(ctx, input)
			},
		},
		"create-job": {
			Name:   "create-job",
			Fields: fields_create_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateJob(ctx, input)
			},
		},
		"create-quantum-task": {
			Name:   "create-quantum-task",
			Fields: fields_create_quantum_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateQuantumTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_quantum_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateQuantumTask(ctx, input)
			},
		},
		"create-spending-limit": {
			Name:   "create-spending-limit",
			Fields: fields_create_spending_limit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSpendingLimitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_spending_limit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSpendingLimit(ctx, input)
			},
		},
		"delete-spending-limit": {
			Name:   "delete-spending-limit",
			Fields: fields_delete_spending_limit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSpendingLimitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_spending_limit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSpendingLimit(ctx, input)
			},
		},
		"get-device": {
			Name:   "get-device",
			Fields: fields_get_device,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDeviceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_device, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDevice(ctx, input)
			},
		},
		"get-job": {
			Name:   "get-job",
			Fields: fields_get_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetJob(ctx, input)
			},
		},
		"get-quantum-task": {
			Name:   "get-quantum-task",
			Fields: fields_get_quantum_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetQuantumTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_quantum_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetQuantumTask(ctx, input)
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
		"search-devices": {
			Name:   "search-devices",
			Fields: fields_search_devices,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchDevicesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_devices, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchDevices(ctx, input)
				}
				var results []*svc.SearchDevicesOutput
				p := svc.NewSearchDevicesPaginator(client, input)
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
		"search-jobs": {
			Name:   "search-jobs",
			Fields: fields_search_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchJobs(ctx, input)
				}
				var results []*svc.SearchJobsOutput
				p := svc.NewSearchJobsPaginator(client, input)
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
		"search-quantum-tasks": {
			Name:   "search-quantum-tasks",
			Fields: fields_search_quantum_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchQuantumTasksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_quantum_tasks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchQuantumTasks(ctx, input)
				}
				var results []*svc.SearchQuantumTasksOutput
				p := svc.NewSearchQuantumTasksPaginator(client, input)
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
		"search-spending-limits": {
			Name:   "search-spending-limits",
			Fields: fields_search_spending_limits,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchSpendingLimitsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_spending_limits, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchSpendingLimits(ctx, input)
				}
				var results []*svc.SearchSpendingLimitsOutput
				p := svc.NewSearchSpendingLimitsPaginator(client, input)
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
		"update-spending-limit": {
			Name:   "update-spending-limit",
			Fields: fields_update_spending_limit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSpendingLimitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_spending_limit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSpendingLimit(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("braket", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
