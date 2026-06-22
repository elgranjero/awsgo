package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/datapipeline"
)

var fields_activate_pipeline = []leanruntime.Field{
	{Name: "ParameterValues", Flag: "parameter-values", Type: "[]types.ParameterValue", Required: false},
	{Name: "PipelineId", Flag: "pipeline-id", Type: "*string", Required: true},
	{Name: "StartTimestamp", Flag: "start-timestamp", Type: "*time.Time", Required: false},
}

var fields_add_tags = []leanruntime.Field{
	{Name: "PipelineId", Flag: "pipeline-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_create_pipeline = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "UniqueId", Flag: "unique-id", Type: "*string", Required: true},
}

var fields_deactivate_pipeline = []leanruntime.Field{
	{Name: "CancelActive", Flag: "cancel-active", Type: "*bool", Required: false},
	{Name: "PipelineId", Flag: "pipeline-id", Type: "*string", Required: true},
}

var fields_delete_pipeline = []leanruntime.Field{
	{Name: "PipelineId", Flag: "pipeline-id", Type: "*string", Required: true},
}

var fields_describe_objects = []leanruntime.Field{
	{Name: "EvaluateExpressions", Flag: "evaluate-expressions", Type: "bool", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "ObjectIds", Flag: "object-ids", Type: "[]string", Required: true},
	{Name: "PipelineId", Flag: "pipeline-id", Type: "*string", Required: true},
}

var fields_describe_pipelines = []leanruntime.Field{
	{Name: "PipelineIds", Flag: "pipeline-ids", Type: "[]string", Required: true},
}

var fields_evaluate_expression = []leanruntime.Field{
	{Name: "Expression", Flag: "expression", Type: "*string", Required: true},
	{Name: "ObjectId", Flag: "object-id", Type: "*string", Required: true},
	{Name: "PipelineId", Flag: "pipeline-id", Type: "*string", Required: true},
}

var fields_get_pipeline_definition = []leanruntime.Field{
	{Name: "PipelineId", Flag: "pipeline-id", Type: "*string", Required: true},
	{Name: "Version", Flag: "version", Type: "*string", Required: false},
}

var fields_list_pipelines = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
}

var fields_poll_for_task = []leanruntime.Field{
	{Name: "Hostname", Flag: "hostname", Type: "*string", Required: false},
	{Name: "InstanceIdentity", Flag: "instance-identity", Type: "*types.InstanceIdentity", Required: false},
	{Name: "WorkerGroup", Flag: "worker-group", Type: "*string", Required: true},
}

var fields_put_pipeline_definition = []leanruntime.Field{
	{Name: "ParameterObjects", Flag: "parameter-objects", Type: "[]types.ParameterObject", Required: false},
	{Name: "ParameterValues", Flag: "parameter-values", Type: "[]types.ParameterValue", Required: false},
	{Name: "PipelineId", Flag: "pipeline-id", Type: "*string", Required: true},
	{Name: "PipelineObjects", Flag: "pipeline-objects", Type: "[]types.PipelineObject", Required: true},
}

var fields_query_objects = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "PipelineId", Flag: "pipeline-id", Type: "*string", Required: true},
	{Name: "Query", Flag: "query", Type: "*types.Query", Required: false},
	{Name: "Sphere", Flag: "sphere", Type: "*string", Required: true},
}

var fields_remove_tags = []leanruntime.Field{
	{Name: "PipelineId", Flag: "pipeline-id", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_report_task_progress = []leanruntime.Field{
	{Name: "Fields", Flag: "fields", Type: "[]types.Field", Required: false},
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: true},
}

var fields_report_task_runner_heartbeat = []leanruntime.Field{
	{Name: "Hostname", Flag: "hostname", Type: "*string", Required: false},
	{Name: "TaskrunnerId", Flag: "taskrunner-id", Type: "*string", Required: true},
	{Name: "WorkerGroup", Flag: "worker-group", Type: "*string", Required: false},
}

var fields_set_status = []leanruntime.Field{
	{Name: "ObjectIds", Flag: "object-ids", Type: "[]string", Required: true},
	{Name: "PipelineId", Flag: "pipeline-id", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "*string", Required: true},
}

var fields_set_task_status = []leanruntime.Field{
	{Name: "ErrorId", Flag: "error-id", Type: "*string", Required: false},
	{Name: "ErrorMessage", Flag: "error-message", Type: "*string", Required: false},
	{Name: "ErrorStackTrace", Flag: "error-stack-trace", Type: "*string", Required: false},
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: true},
	{Name: "TaskStatus", Flag: "task-status", Type: "types.TaskStatus", Required: true},
}

var fields_validate_pipeline_definition = []leanruntime.Field{
	{Name: "ParameterObjects", Flag: "parameter-objects", Type: "[]types.ParameterObject", Required: false},
	{Name: "ParameterValues", Flag: "parameter-values", Type: "[]types.ParameterValue", Required: false},
	{Name: "PipelineId", Flag: "pipeline-id", Type: "*string", Required: true},
	{Name: "PipelineObjects", Flag: "pipeline-objects", Type: "[]types.PipelineObject", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"activate-pipeline": {
			Name:   "activate-pipeline",
			Fields: fields_activate_pipeline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ActivatePipelineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_activate_pipeline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ActivatePipeline(ctx, input)
			},
		},
		"add-tags": {
			Name:   "add-tags",
			Fields: fields_add_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddTags(ctx, input)
			},
		},
		"create-pipeline": {
			Name:   "create-pipeline",
			Fields: fields_create_pipeline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePipelineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_pipeline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePipeline(ctx, input)
			},
		},
		"deactivate-pipeline": {
			Name:   "deactivate-pipeline",
			Fields: fields_deactivate_pipeline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeactivatePipelineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deactivate_pipeline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeactivatePipeline(ctx, input)
			},
		},
		"delete-pipeline": {
			Name:   "delete-pipeline",
			Fields: fields_delete_pipeline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePipelineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_pipeline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePipeline(ctx, input)
			},
		},
		"describe-objects": {
			Name:   "describe-objects",
			Fields: fields_describe_objects,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeObjectsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_objects, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeObjects(ctx, input)
				}
				var results []*svc.DescribeObjectsOutput
				p := svc.NewDescribeObjectsPaginator(client, input)
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
		"describe-pipelines": {
			Name:   "describe-pipelines",
			Fields: fields_describe_pipelines,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePipelinesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_pipelines, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePipelines(ctx, input)
			},
		},
		"evaluate-expression": {
			Name:   "evaluate-expression",
			Fields: fields_evaluate_expression,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EvaluateExpressionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_evaluate_expression, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EvaluateExpression(ctx, input)
			},
		},
		"get-pipeline-definition": {
			Name:   "get-pipeline-definition",
			Fields: fields_get_pipeline_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPipelineDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_pipeline_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPipelineDefinition(ctx, input)
			},
		},
		"list-pipelines": {
			Name:   "list-pipelines",
			Fields: fields_list_pipelines,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPipelinesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_pipelines, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPipelines(ctx, input)
				}
				var results []*svc.ListPipelinesOutput
				p := svc.NewListPipelinesPaginator(client, input)
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
		"poll-for-task": {
			Name:   "poll-for-task",
			Fields: fields_poll_for_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PollForTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_poll_for_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PollForTask(ctx, input)
			},
		},
		"put-pipeline-definition": {
			Name:   "put-pipeline-definition",
			Fields: fields_put_pipeline_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutPipelineDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_pipeline_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutPipelineDefinition(ctx, input)
			},
		},
		"query-objects": {
			Name:   "query-objects",
			Fields: fields_query_objects,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.QueryObjectsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_query_objects, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.QueryObjects(ctx, input)
				}
				var results []*svc.QueryObjectsOutput
				p := svc.NewQueryObjectsPaginator(client, input)
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
		"remove-tags": {
			Name:   "remove-tags",
			Fields: fields_remove_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveTags(ctx, input)
			},
		},
		"report-task-progress": {
			Name:   "report-task-progress",
			Fields: fields_report_task_progress,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ReportTaskProgressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_report_task_progress, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ReportTaskProgress(ctx, input)
			},
		},
		"report-task-runner-heartbeat": {
			Name:   "report-task-runner-heartbeat",
			Fields: fields_report_task_runner_heartbeat,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ReportTaskRunnerHeartbeatInput{}
				if _, err := leanruntime.ApplyInput(input, fields_report_task_runner_heartbeat, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ReportTaskRunnerHeartbeat(ctx, input)
			},
		},
		"set-status": {
			Name:   "set-status",
			Fields: fields_set_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetStatus(ctx, input)
			},
		},
		"set-task-status": {
			Name:   "set-task-status",
			Fields: fields_set_task_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetTaskStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_task_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetTaskStatus(ctx, input)
			},
		},
		"validate-pipeline-definition": {
			Name:   "validate-pipeline-definition",
			Fields: fields_validate_pipeline_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ValidatePipelineDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_validate_pipeline_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ValidatePipelineDefinition(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("datapipeline", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
