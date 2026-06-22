package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/batch"
)

var fields_cancel_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "Reason", Flag: "reason", Type: "*string", Required: true},
}

var fields_create_compute_environment = []leanruntime.Field{
	{Name: "ComputeEnvironmentName", Flag: "compute-environment-name", Type: "*string", Required: true},
	{Name: "ComputeResources", Flag: "compute-resources", Type: "*types.ComputeResource", Required: false},
	{Name: "Context", Flag: "context", Type: "*string", Required: false},
	{Name: "EksConfiguration", Flag: "eks-configuration", Type: "*types.EksConfiguration", Required: false},
	{Name: "ServiceRole", Flag: "service-role", Type: "*string", Required: false},
	{Name: "State", Flag: "state", Type: "types.CEState", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.CEType", Required: true},
	{Name: "UnmanagedvCpus", Flag: "unmanagedv-cpus", Type: "*int32", Required: false},
}

var fields_create_consumable_resource = []leanruntime.Field{
	{Name: "ConsumableResourceName", Flag: "consumable-resource-name", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TotalQuantity", Flag: "total-quantity", Type: "*int64", Required: false},
}

var fields_create_job_queue = []leanruntime.Field{
	{Name: "ComputeEnvironmentOrder", Flag: "compute-environment-order", Type: "[]types.ComputeEnvironmentOrder", Required: false},
	{Name: "JobQueueName", Flag: "job-queue-name", Type: "*string", Required: true},
	{Name: "JobQueueType", Flag: "job-queue-type", Type: "types.JobQueueType", Required: false},
	{Name: "JobStateTimeLimitActions", Flag: "job-state-time-limit-actions", Type: "[]types.JobStateTimeLimitAction", Required: false},
	{Name: "Priority", Flag: "priority", Type: "*int32", Required: true},
	{Name: "SchedulingPolicyArn", Flag: "scheduling-policy-arn", Type: "*string", Required: false},
	{Name: "ServiceEnvironmentOrder", Flag: "service-environment-order", Type: "[]types.ServiceEnvironmentOrder", Required: false},
	{Name: "State", Flag: "state", Type: "types.JQState", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_scheduling_policy = []leanruntime.Field{
	{Name: "FairsharePolicy", Flag: "fairshare-policy", Type: "*types.FairsharePolicy", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_service_environment = []leanruntime.Field{
	{Name: "CapacityLimits", Flag: "capacity-limits", Type: "[]types.CapacityLimit", Required: true},
	{Name: "ServiceEnvironmentName", Flag: "service-environment-name", Type: "*string", Required: true},
	{Name: "ServiceEnvironmentType", Flag: "service-environment-type", Type: "types.ServiceEnvironmentType", Required: true},
	{Name: "State", Flag: "state", Type: "types.ServiceEnvironmentState", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_compute_environment = []leanruntime.Field{
	{Name: "ComputeEnvironment", Flag: "compute-environment", Type: "*string", Required: true},
}

var fields_delete_consumable_resource = []leanruntime.Field{
	{Name: "ConsumableResource", Flag: "consumable-resource", Type: "*string", Required: true},
}

var fields_delete_job_queue = []leanruntime.Field{
	{Name: "JobQueue", Flag: "job-queue", Type: "*string", Required: true},
}

var fields_delete_scheduling_policy = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_delete_service_environment = []leanruntime.Field{
	{Name: "ServiceEnvironment", Flag: "service-environment", Type: "*string", Required: true},
}

var fields_deregister_job_definition = []leanruntime.Field{
	{Name: "JobDefinition", Flag: "job-definition", Type: "*string", Required: true},
}

var fields_describe_compute_environments = []leanruntime.Field{
	{Name: "ComputeEnvironments", Flag: "compute-environments", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_consumable_resource = []leanruntime.Field{
	{Name: "ConsumableResource", Flag: "consumable-resource", Type: "*string", Required: true},
}

var fields_describe_job_definitions = []leanruntime.Field{
	{Name: "JobDefinitionName", Flag: "job-definition-name", Type: "*string", Required: false},
	{Name: "JobDefinitions", Flag: "job-definitions", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "*string", Required: false},
}

var fields_describe_job_queues = []leanruntime.Field{
	{Name: "JobQueues", Flag: "job-queues", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_jobs = []leanruntime.Field{
	{Name: "Jobs", Flag: "jobs", Type: "[]string", Required: true},
}

var fields_describe_scheduling_policies = []leanruntime.Field{
	{Name: "Arns", Flag: "arns", Type: "[]string", Required: true},
}

var fields_describe_service_environments = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServiceEnvironments", Flag: "service-environments", Type: "[]string", Required: false},
}

var fields_describe_service_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_get_job_queue_snapshot = []leanruntime.Field{
	{Name: "JobQueue", Flag: "job-queue", Type: "*string", Required: true},
}

var fields_list_consumable_resources = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.KeyValuesPair", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_jobs = []leanruntime.Field{
	{Name: "ArrayJobId", Flag: "array-job-id", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.KeyValuesPair", Required: false},
	{Name: "JobQueue", Flag: "job-queue", Type: "*string", Required: false},
	{Name: "JobStatus", Flag: "job-status", Type: "types.JobStatus", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MultiNodeJobId", Flag: "multi-node-job-id", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_jobs_by_consumable_resource = []leanruntime.Field{
	{Name: "ConsumableResource", Flag: "consumable-resource", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.KeyValuesPair", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_scheduling_policies = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_service_jobs = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.KeyValuesPair", Required: false},
	{Name: "JobQueue", Flag: "job-queue", Type: "*string", Required: false},
	{Name: "JobStatus", Flag: "job-status", Type: "types.ServiceJobStatus", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_register_job_definition = []leanruntime.Field{
	{Name: "ConsumableResourceProperties", Flag: "consumable-resource-properties", Type: "*types.ConsumableResourceProperties", Required: false},
	{Name: "ContainerProperties", Flag: "container-properties", Type: "*types.ContainerProperties", Required: false},
	{Name: "EcsProperties", Flag: "ecs-properties", Type: "*types.EcsProperties", Required: false},
	{Name: "EksProperties", Flag: "eks-properties", Type: "*types.EksProperties", Required: false},
	{Name: "JobDefinitionName", Flag: "job-definition-name", Type: "*string", Required: true},
	{Name: "NodeProperties", Flag: "node-properties", Type: "*types.NodeProperties", Required: false},
	{Name: "Parameters", Flag: "parameters", Type: "map[string]string", Required: false},
	{Name: "PlatformCapabilities", Flag: "platform-capabilities", Type: "[]types.PlatformCapability", Required: false},
	{Name: "PropagateTags", Flag: "propagate-tags", Type: "*bool", Required: false},
	{Name: "RetryStrategy", Flag: "retry-strategy", Type: "*types.RetryStrategy", Required: false},
	{Name: "SchedulingPriority", Flag: "scheduling-priority", Type: "*int32", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Timeout", Flag: "timeout", Type: "*types.JobTimeout", Required: false},
	{Name: "Type", Flag: "type", Type: "types.JobDefinitionType", Required: true},
}

var fields_submit_job = []leanruntime.Field{
	{Name: "ArrayProperties", Flag: "array-properties", Type: "*types.ArrayProperties", Required: false},
	{Name: "ConsumableResourcePropertiesOverride", Flag: "consumable-resource-properties-override", Type: "*types.ConsumableResourceProperties", Required: false},
	{Name: "ContainerOverrides", Flag: "container-overrides", Type: "*types.ContainerOverrides", Required: false},
	{Name: "DependsOn", Flag: "depends-on", Type: "[]types.JobDependency", Required: false},
	{Name: "EcsPropertiesOverride", Flag: "ecs-properties-override", Type: "*types.EcsPropertiesOverride", Required: false},
	{Name: "EksPropertiesOverride", Flag: "eks-properties-override", Type: "*types.EksPropertiesOverride", Required: false},
	{Name: "JobDefinition", Flag: "job-definition", Type: "*string", Required: true},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: true},
	{Name: "JobQueue", Flag: "job-queue", Type: "*string", Required: true},
	{Name: "NodeOverrides", Flag: "node-overrides", Type: "*types.NodeOverrides", Required: false},
	{Name: "Parameters", Flag: "parameters", Type: "map[string]string", Required: false},
	{Name: "PropagateTags", Flag: "propagate-tags", Type: "*bool", Required: false},
	{Name: "RetryStrategy", Flag: "retry-strategy", Type: "*types.RetryStrategy", Required: false},
	{Name: "SchedulingPriorityOverride", Flag: "scheduling-priority-override", Type: "*int32", Required: false},
	{Name: "ShareIdentifier", Flag: "share-identifier", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Timeout", Flag: "timeout", Type: "*types.JobTimeout", Required: false},
}

var fields_submit_service_job = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: true},
	{Name: "JobQueue", Flag: "job-queue", Type: "*string", Required: true},
	{Name: "RetryStrategy", Flag: "retry-strategy", Type: "*types.ServiceJobRetryStrategy", Required: false},
	{Name: "SchedulingPriority", Flag: "scheduling-priority", Type: "*int32", Required: false},
	{Name: "ServiceJobType", Flag: "service-job-type", Type: "types.ServiceJobType", Required: true},
	{Name: "ServiceRequestPayload", Flag: "service-request-payload", Type: "*string", Required: true},
	{Name: "ShareIdentifier", Flag: "share-identifier", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TimeoutConfig", Flag: "timeout-config", Type: "*types.ServiceJobTimeout", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_terminate_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "Reason", Flag: "reason", Type: "*string", Required: true},
}

var fields_terminate_service_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "Reason", Flag: "reason", Type: "*string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_compute_environment = []leanruntime.Field{
	{Name: "ComputeEnvironment", Flag: "compute-environment", Type: "*string", Required: true},
	{Name: "ComputeResources", Flag: "compute-resources", Type: "*types.ComputeResourceUpdate", Required: false},
	{Name: "Context", Flag: "context", Type: "*string", Required: false},
	{Name: "ServiceRole", Flag: "service-role", Type: "*string", Required: false},
	{Name: "State", Flag: "state", Type: "types.CEState", Required: false},
	{Name: "UnmanagedvCpus", Flag: "unmanagedv-cpus", Type: "*int32", Required: false},
	{Name: "UpdatePolicy", Flag: "update-policy", Type: "*types.UpdatePolicy", Required: false},
}

var fields_update_consumable_resource = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ConsumableResource", Flag: "consumable-resource", Type: "*string", Required: true},
	{Name: "Operation", Flag: "operation", Type: "*string", Required: false},
	{Name: "Quantity", Flag: "quantity", Type: "*int64", Required: false},
}

var fields_update_job_queue = []leanruntime.Field{
	{Name: "ComputeEnvironmentOrder", Flag: "compute-environment-order", Type: "[]types.ComputeEnvironmentOrder", Required: false},
	{Name: "JobQueue", Flag: "job-queue", Type: "*string", Required: true},
	{Name: "JobStateTimeLimitActions", Flag: "job-state-time-limit-actions", Type: "[]types.JobStateTimeLimitAction", Required: false},
	{Name: "Priority", Flag: "priority", Type: "*int32", Required: false},
	{Name: "SchedulingPolicyArn", Flag: "scheduling-policy-arn", Type: "*string", Required: false},
	{Name: "ServiceEnvironmentOrder", Flag: "service-environment-order", Type: "[]types.ServiceEnvironmentOrder", Required: false},
	{Name: "State", Flag: "state", Type: "types.JQState", Required: false},
}

var fields_update_scheduling_policy = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "FairsharePolicy", Flag: "fairshare-policy", Type: "*types.FairsharePolicy", Required: false},
}

var fields_update_service_environment = []leanruntime.Field{
	{Name: "CapacityLimits", Flag: "capacity-limits", Type: "[]types.CapacityLimit", Required: false},
	{Name: "ServiceEnvironment", Flag: "service-environment", Type: "*string", Required: true},
	{Name: "State", Flag: "state", Type: "types.ServiceEnvironmentState", Required: false},
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
		"create-compute-environment": {
			Name:   "create-compute-environment",
			Fields: fields_create_compute_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateComputeEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_compute_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateComputeEnvironment(ctx, input)
			},
		},
		"create-consumable-resource": {
			Name:   "create-consumable-resource",
			Fields: fields_create_consumable_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConsumableResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_consumable_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConsumableResource(ctx, input)
			},
		},
		"create-job-queue": {
			Name:   "create-job-queue",
			Fields: fields_create_job_queue,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateJobQueueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_job_queue, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateJobQueue(ctx, input)
			},
		},
		"create-scheduling-policy": {
			Name:   "create-scheduling-policy",
			Fields: fields_create_scheduling_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSchedulingPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_scheduling_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSchedulingPolicy(ctx, input)
			},
		},
		"create-service-environment": {
			Name:   "create-service-environment",
			Fields: fields_create_service_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateServiceEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_service_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateServiceEnvironment(ctx, input)
			},
		},
		"delete-compute-environment": {
			Name:   "delete-compute-environment",
			Fields: fields_delete_compute_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteComputeEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_compute_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteComputeEnvironment(ctx, input)
			},
		},
		"delete-consumable-resource": {
			Name:   "delete-consumable-resource",
			Fields: fields_delete_consumable_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConsumableResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_consumable_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConsumableResource(ctx, input)
			},
		},
		"delete-job-queue": {
			Name:   "delete-job-queue",
			Fields: fields_delete_job_queue,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteJobQueueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_job_queue, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteJobQueue(ctx, input)
			},
		},
		"delete-scheduling-policy": {
			Name:   "delete-scheduling-policy",
			Fields: fields_delete_scheduling_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSchedulingPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_scheduling_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSchedulingPolicy(ctx, input)
			},
		},
		"delete-service-environment": {
			Name:   "delete-service-environment",
			Fields: fields_delete_service_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteServiceEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_service_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteServiceEnvironment(ctx, input)
			},
		},
		"deregister-job-definition": {
			Name:   "deregister-job-definition",
			Fields: fields_deregister_job_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterJobDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_job_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterJobDefinition(ctx, input)
			},
		},
		"describe-compute-environments": {
			Name:   "describe-compute-environments",
			Fields: fields_describe_compute_environments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeComputeEnvironmentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_compute_environments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeComputeEnvironments(ctx, input)
				}
				var results []*svc.DescribeComputeEnvironmentsOutput
				p := svc.NewDescribeComputeEnvironmentsPaginator(client, input)
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
		"describe-consumable-resource": {
			Name:   "describe-consumable-resource",
			Fields: fields_describe_consumable_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConsumableResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_consumable_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeConsumableResource(ctx, input)
			},
		},
		"describe-job-definitions": {
			Name:   "describe-job-definitions",
			Fields: fields_describe_job_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeJobDefinitionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_job_definitions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeJobDefinitions(ctx, input)
				}
				var results []*svc.DescribeJobDefinitionsOutput
				p := svc.NewDescribeJobDefinitionsPaginator(client, input)
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
		"describe-job-queues": {
			Name:   "describe-job-queues",
			Fields: fields_describe_job_queues,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeJobQueuesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_job_queues, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeJobQueues(ctx, input)
				}
				var results []*svc.DescribeJobQueuesOutput
				p := svc.NewDescribeJobQueuesPaginator(client, input)
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
		"describe-jobs": {
			Name:   "describe-jobs",
			Fields: fields_describe_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeJobsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_jobs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeJobs(ctx, input)
			},
		},
		"describe-scheduling-policies": {
			Name:   "describe-scheduling-policies",
			Fields: fields_describe_scheduling_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSchedulingPoliciesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_scheduling_policies, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSchedulingPolicies(ctx, input)
			},
		},
		"describe-service-environments": {
			Name:   "describe-service-environments",
			Fields: fields_describe_service_environments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeServiceEnvironmentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_service_environments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeServiceEnvironments(ctx, input)
				}
				var results []*svc.DescribeServiceEnvironmentsOutput
				p := svc.NewDescribeServiceEnvironmentsPaginator(client, input)
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
		"describe-service-job": {
			Name:   "describe-service-job",
			Fields: fields_describe_service_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeServiceJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_service_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeServiceJob(ctx, input)
			},
		},
		"get-job-queue-snapshot": {
			Name:   "get-job-queue-snapshot",
			Fields: fields_get_job_queue_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetJobQueueSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_job_queue_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetJobQueueSnapshot(ctx, input)
			},
		},
		"list-consumable-resources": {
			Name:   "list-consumable-resources",
			Fields: fields_list_consumable_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConsumableResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_consumable_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConsumableResources(ctx, input)
				}
				var results []*svc.ListConsumableResourcesOutput
				p := svc.NewListConsumableResourcesPaginator(client, input)
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
		"list-jobs": {
			Name:   "list-jobs",
			Fields: fields_list_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListJobs(ctx, input)
				}
				var results []*svc.ListJobsOutput
				p := svc.NewListJobsPaginator(client, input)
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
		"list-jobs-by-consumable-resource": {
			Name:   "list-jobs-by-consumable-resource",
			Fields: fields_list_jobs_by_consumable_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListJobsByConsumableResourceInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_jobs_by_consumable_resource, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListJobsByConsumableResource(ctx, input)
				}
				var results []*svc.ListJobsByConsumableResourceOutput
				p := svc.NewListJobsByConsumableResourcePaginator(client, input)
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
		"list-scheduling-policies": {
			Name:   "list-scheduling-policies",
			Fields: fields_list_scheduling_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSchedulingPoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_scheduling_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSchedulingPolicies(ctx, input)
				}
				var results []*svc.ListSchedulingPoliciesOutput
				p := svc.NewListSchedulingPoliciesPaginator(client, input)
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
		"list-service-jobs": {
			Name:   "list-service-jobs",
			Fields: fields_list_service_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServiceJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_service_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServiceJobs(ctx, input)
				}
				var results []*svc.ListServiceJobsOutput
				p := svc.NewListServiceJobsPaginator(client, input)
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
		"register-job-definition": {
			Name:   "register-job-definition",
			Fields: fields_register_job_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterJobDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_job_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterJobDefinition(ctx, input)
			},
		},
		"submit-job": {
			Name:   "submit-job",
			Fields: fields_submit_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SubmitJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_submit_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SubmitJob(ctx, input)
			},
		},
		"submit-service-job": {
			Name:   "submit-service-job",
			Fields: fields_submit_service_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SubmitServiceJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_submit_service_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SubmitServiceJob(ctx, input)
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
		"terminate-job": {
			Name:   "terminate-job",
			Fields: fields_terminate_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TerminateJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_terminate_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TerminateJob(ctx, input)
			},
		},
		"terminate-service-job": {
			Name:   "terminate-service-job",
			Fields: fields_terminate_service_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TerminateServiceJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_terminate_service_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TerminateServiceJob(ctx, input)
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
		"update-compute-environment": {
			Name:   "update-compute-environment",
			Fields: fields_update_compute_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateComputeEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_compute_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateComputeEnvironment(ctx, input)
			},
		},
		"update-consumable-resource": {
			Name:   "update-consumable-resource",
			Fields: fields_update_consumable_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConsumableResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_consumable_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConsumableResource(ctx, input)
			},
		},
		"update-job-queue": {
			Name:   "update-job-queue",
			Fields: fields_update_job_queue,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateJobQueueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_job_queue, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateJobQueue(ctx, input)
			},
		},
		"update-scheduling-policy": {
			Name:   "update-scheduling-policy",
			Fields: fields_update_scheduling_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSchedulingPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_scheduling_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSchedulingPolicy(ctx, input)
			},
		},
		"update-service-environment": {
			Name:   "update-service-environment",
			Fields: fields_update_service_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateServiceEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_service_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateServiceEnvironment(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("batch", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
