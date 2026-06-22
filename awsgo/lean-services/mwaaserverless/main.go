package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/mwaaserverless"
)

var fields_create_workflow = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DefinitionS3Location", Flag: "definition-s3-location", Type: "*types.DefinitionS3Location", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EncryptionConfiguration", Flag: "encryption-configuration", Type: "*types.EncryptionConfiguration", Required: false},
	{Name: "EngineVersion", Flag: "engine-version", Type: "types.EngineVersion", Required: false},
	{Name: "LoggingConfiguration", Flag: "logging-configuration", Type: "*types.LoggingConfiguration", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NetworkConfiguration", Flag: "network-configuration", Type: "*types.NetworkConfiguration", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TriggerMode", Flag: "trigger-mode", Type: "*string", Required: false},
}

var fields_delete_workflow = []leanruntime.Field{
	{Name: "WorkflowArn", Flag: "workflow-arn", Type: "*string", Required: true},
	{Name: "WorkflowVersion", Flag: "workflow-version", Type: "*string", Required: false},
}

var fields_get_task_instance = []leanruntime.Field{
	{Name: "RunId", Flag: "run-id", Type: "*string", Required: true},
	{Name: "TaskInstanceId", Flag: "task-instance-id", Type: "*string", Required: true},
	{Name: "WorkflowArn", Flag: "workflow-arn", Type: "*string", Required: true},
}

var fields_get_workflow = []leanruntime.Field{
	{Name: "WorkflowArn", Flag: "workflow-arn", Type: "*string", Required: true},
	{Name: "WorkflowVersion", Flag: "workflow-version", Type: "*string", Required: false},
}

var fields_get_workflow_run = []leanruntime.Field{
	{Name: "RunId", Flag: "run-id", Type: "*string", Required: true},
	{Name: "WorkflowArn", Flag: "workflow-arn", Type: "*string", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_task_instances = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RunId", Flag: "run-id", Type: "*string", Required: true},
	{Name: "WorkflowArn", Flag: "workflow-arn", Type: "*string", Required: true},
}

var fields_list_workflow_runs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkflowArn", Flag: "workflow-arn", Type: "*string", Required: true},
	{Name: "WorkflowVersion", Flag: "workflow-version", Type: "*string", Required: false},
}

var fields_list_workflow_versions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkflowArn", Flag: "workflow-arn", Type: "*string", Required: true},
}

var fields_list_workflows = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_start_workflow_run = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "OverrideParameters", Flag: "override-parameters", Type: "map[string]document.Interface", Required: false},
	{Name: "WorkflowArn", Flag: "workflow-arn", Type: "*string", Required: true},
	{Name: "WorkflowVersion", Flag: "workflow-version", Type: "*string", Required: false},
}

var fields_stop_workflow_run = []leanruntime.Field{
	{Name: "RunId", Flag: "run-id", Type: "*string", Required: true},
	{Name: "WorkflowArn", Flag: "workflow-arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_workflow = []leanruntime.Field{
	{Name: "DefinitionS3Location", Flag: "definition-s3-location", Type: "*types.DefinitionS3Location", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EngineVersion", Flag: "engine-version", Type: "types.EngineVersion", Required: false},
	{Name: "LoggingConfiguration", Flag: "logging-configuration", Type: "*types.LoggingConfiguration", Required: false},
	{Name: "NetworkConfiguration", Flag: "network-configuration", Type: "*types.NetworkConfiguration", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "TriggerMode", Flag: "trigger-mode", Type: "*string", Required: false},
	{Name: "WorkflowArn", Flag: "workflow-arn", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-workflow": {
			Name:   "create-workflow",
			Fields: fields_create_workflow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWorkflowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_workflow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWorkflow(ctx, input)
			},
		},
		"delete-workflow": {
			Name:   "delete-workflow",
			Fields: fields_delete_workflow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWorkflowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_workflow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWorkflow(ctx, input)
			},
		},
		"get-task-instance": {
			Name:   "get-task-instance",
			Fields: fields_get_task_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTaskInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_task_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTaskInstance(ctx, input)
			},
		},
		"get-workflow": {
			Name:   "get-workflow",
			Fields: fields_get_workflow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWorkflowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_workflow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWorkflow(ctx, input)
			},
		},
		"get-workflow-run": {
			Name:   "get-workflow-run",
			Fields: fields_get_workflow_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWorkflowRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_workflow_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWorkflowRun(ctx, input)
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
		"list-task-instances": {
			Name:   "list-task-instances",
			Fields: fields_list_task_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTaskInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_task_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTaskInstances(ctx, input)
				}
				var results []*svc.ListTaskInstancesOutput
				p := svc.NewListTaskInstancesPaginator(client, input)
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
		"list-workflow-runs": {
			Name:   "list-workflow-runs",
			Fields: fields_list_workflow_runs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkflowRunsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_workflow_runs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkflowRuns(ctx, input)
				}
				var results []*svc.ListWorkflowRunsOutput
				p := svc.NewListWorkflowRunsPaginator(client, input)
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
		"list-workflow-versions": {
			Name:   "list-workflow-versions",
			Fields: fields_list_workflow_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkflowVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_workflow_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkflowVersions(ctx, input)
				}
				var results []*svc.ListWorkflowVersionsOutput
				p := svc.NewListWorkflowVersionsPaginator(client, input)
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
		"list-workflows": {
			Name:   "list-workflows",
			Fields: fields_list_workflows,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkflowsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_workflows, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkflows(ctx, input)
				}
				var results []*svc.ListWorkflowsOutput
				p := svc.NewListWorkflowsPaginator(client, input)
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
		"start-workflow-run": {
			Name:   "start-workflow-run",
			Fields: fields_start_workflow_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartWorkflowRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_workflow_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartWorkflowRun(ctx, input)
			},
		},
		"stop-workflow-run": {
			Name:   "stop-workflow-run",
			Fields: fields_stop_workflow_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopWorkflowRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_workflow_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopWorkflowRun(ctx, input)
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
		"update-workflow": {
			Name:   "update-workflow",
			Fields: fields_update_workflow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWorkflowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_workflow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWorkflow(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("mwaaserverless", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
