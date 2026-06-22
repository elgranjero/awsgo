package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/novaact"
)

var fields_create_act = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
	{Name: "Task", Flag: "task", Type: "*string", Required: true},
	{Name: "ToolSpecs", Flag: "tool-specs", Type: "[]types.ToolSpec", Required: false},
	{Name: "WorkflowDefinitionName", Flag: "workflow-definition-name", Type: "*string", Required: true},
	{Name: "WorkflowRunId", Flag: "workflow-run-id", Type: "*string", Required: true},
}

var fields_create_session = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "WorkflowDefinitionName", Flag: "workflow-definition-name", Type: "*string", Required: true},
	{Name: "WorkflowRunId", Flag: "workflow-run-id", Type: "*string", Required: true},
}

var fields_create_workflow_definition = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ExportConfig", Flag: "export-config", Type: "*types.WorkflowExportConfig", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_workflow_run = []leanruntime.Field{
	{Name: "ClientInfo", Flag: "client-info", Type: "*types.ClientInfo", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "LogGroupName", Flag: "log-group-name", Type: "*string", Required: false},
	{Name: "ModelId", Flag: "model-id", Type: "*string", Required: true},
	{Name: "WorkflowDefinitionName", Flag: "workflow-definition-name", Type: "*string", Required: true},
}

var fields_delete_workflow_definition = []leanruntime.Field{
	{Name: "WorkflowDefinitionName", Flag: "workflow-definition-name", Type: "*string", Required: true},
}

var fields_delete_workflow_run = []leanruntime.Field{
	{Name: "WorkflowDefinitionName", Flag: "workflow-definition-name", Type: "*string", Required: true},
	{Name: "WorkflowRunId", Flag: "workflow-run-id", Type: "*string", Required: true},
}

var fields_get_workflow_definition = []leanruntime.Field{
	{Name: "WorkflowDefinitionName", Flag: "workflow-definition-name", Type: "*string", Required: true},
}

var fields_get_workflow_run = []leanruntime.Field{
	{Name: "WorkflowDefinitionName", Flag: "workflow-definition-name", Type: "*string", Required: true},
	{Name: "WorkflowRunId", Flag: "workflow-run-id", Type: "*string", Required: true},
}

var fields_invoke_act_step = []leanruntime.Field{
	{Name: "ActId", Flag: "act-id", Type: "*string", Required: true},
	{Name: "CallResults", Flag: "call-results", Type: "[]types.CallResult", Required: true},
	{Name: "PreviousStepId", Flag: "previous-step-id", Type: "*string", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
	{Name: "WorkflowDefinitionName", Flag: "workflow-definition-name", Type: "*string", Required: true},
	{Name: "WorkflowRunId", Flag: "workflow-run-id", Type: "*string", Required: true},
}

var fields_list_acts = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "WorkflowDefinitionName", Flag: "workflow-definition-name", Type: "*string", Required: true},
	{Name: "WorkflowRunId", Flag: "workflow-run-id", Type: "*string", Required: false},
}

var fields_list_models = []leanruntime.Field{
	{Name: "ClientCompatibilityVersion", Flag: "client-compatibility-version", Type: "*int32", Required: true},
}

var fields_list_sessions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "WorkflowDefinitionName", Flag: "workflow-definition-name", Type: "*string", Required: true},
	{Name: "WorkflowRunId", Flag: "workflow-run-id", Type: "*string", Required: true},
}

var fields_list_workflow_definitions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_workflow_runs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "WorkflowDefinitionName", Flag: "workflow-definition-name", Type: "*string", Required: true},
}

var fields_update_act = []leanruntime.Field{
	{Name: "ActId", Flag: "act-id", Type: "*string", Required: true},
	{Name: "Error", Flag: "error", Type: "*types.ActError", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.ActStatus", Required: true},
	{Name: "WorkflowDefinitionName", Flag: "workflow-definition-name", Type: "*string", Required: true},
	{Name: "WorkflowRunId", Flag: "workflow-run-id", Type: "*string", Required: true},
}

var fields_update_workflow_run = []leanruntime.Field{
	{Name: "Status", Flag: "status", Type: "types.WorkflowRunStatus", Required: true},
	{Name: "WorkflowDefinitionName", Flag: "workflow-definition-name", Type: "*string", Required: true},
	{Name: "WorkflowRunId", Flag: "workflow-run-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-act": {
			Name:   "create-act",
			Fields: fields_create_act,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateActInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_act, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAct(ctx, input)
			},
		},
		"create-session": {
			Name:   "create-session",
			Fields: fields_create_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSession(ctx, input)
			},
		},
		"create-workflow-definition": {
			Name:   "create-workflow-definition",
			Fields: fields_create_workflow_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWorkflowDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_workflow_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWorkflowDefinition(ctx, input)
			},
		},
		"create-workflow-run": {
			Name:   "create-workflow-run",
			Fields: fields_create_workflow_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWorkflowRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_workflow_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWorkflowRun(ctx, input)
			},
		},
		"delete-workflow-definition": {
			Name:   "delete-workflow-definition",
			Fields: fields_delete_workflow_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWorkflowDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_workflow_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWorkflowDefinition(ctx, input)
			},
		},
		"delete-workflow-run": {
			Name:   "delete-workflow-run",
			Fields: fields_delete_workflow_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWorkflowRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_workflow_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWorkflowRun(ctx, input)
			},
		},
		"get-workflow-definition": {
			Name:   "get-workflow-definition",
			Fields: fields_get_workflow_definition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWorkflowDefinitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_workflow_definition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWorkflowDefinition(ctx, input)
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
		"invoke-act-step": {
			Name:   "invoke-act-step",
			Fields: fields_invoke_act_step,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InvokeActStepInput{}
				if _, err := leanruntime.ApplyInput(input, fields_invoke_act_step, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.InvokeActStep(ctx, input)
			},
		},
		"list-acts": {
			Name:   "list-acts",
			Fields: fields_list_acts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListActsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_acts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListActs(ctx, input)
				}
				var results []*svc.ListActsOutput
				p := svc.NewListActsPaginator(client, input)
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
		"list-models": {
			Name:   "list-models",
			Fields: fields_list_models,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListModelsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_models, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListModels(ctx, input)
			},
		},
		"list-sessions": {
			Name:   "list-sessions",
			Fields: fields_list_sessions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSessionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_sessions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSessions(ctx, input)
				}
				var results []*svc.ListSessionsOutput
				p := svc.NewListSessionsPaginator(client, input)
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
		"list-workflow-definitions": {
			Name:   "list-workflow-definitions",
			Fields: fields_list_workflow_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkflowDefinitionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_workflow_definitions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkflowDefinitions(ctx, input)
				}
				var results []*svc.ListWorkflowDefinitionsOutput
				p := svc.NewListWorkflowDefinitionsPaginator(client, input)
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
		"update-act": {
			Name:   "update-act",
			Fields: fields_update_act,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateActInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_act, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAct(ctx, input)
			},
		},
		"update-workflow-run": {
			Name:   "update-workflow-run",
			Fields: fields_update_workflow_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWorkflowRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_workflow_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWorkflowRun(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("novaact", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
