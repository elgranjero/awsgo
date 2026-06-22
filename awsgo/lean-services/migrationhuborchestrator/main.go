package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/migrationhuborchestrator"
)

var fields_create_template = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TemplateDescription", Flag: "template-description", Type: "*string", Required: false},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: true},
	{Name: "TemplateSource", Flag: "template-source", Type: "types.TemplateSource", Required: true},
}

var fields_create_workflow = []leanruntime.Field{
	{Name: "ApplicationConfigurationId", Flag: "application-configuration-id", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InputParameters", Flag: "input-parameters", Type: "map[string]types.StepInput", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "StepTargets", Flag: "step-targets", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TemplateId", Flag: "template-id", Type: "*string", Required: true},
}

var fields_create_workflow_step = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Next", Flag: "next", Type: "[]string", Required: false},
	{Name: "Outputs", Flag: "outputs", Type: "[]types.WorkflowStepOutput", Required: false},
	{Name: "Previous", Flag: "previous", Type: "[]string", Required: false},
	{Name: "StepActionType", Flag: "step-action-type", Type: "types.StepActionType", Required: true},
	{Name: "StepGroupId", Flag: "step-group-id", Type: "*string", Required: true},
	{Name: "StepTarget", Flag: "step-target", Type: "[]string", Required: false},
	{Name: "WorkflowId", Flag: "workflow-id", Type: "*string", Required: true},
	{Name: "WorkflowStepAutomationConfiguration", Flag: "workflow-step-automation-configuration", Type: "*types.WorkflowStepAutomationConfiguration", Required: false},
}

var fields_create_workflow_step_group = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Next", Flag: "next", Type: "[]string", Required: false},
	{Name: "Previous", Flag: "previous", Type: "[]string", Required: false},
	{Name: "WorkflowId", Flag: "workflow-id", Type: "*string", Required: true},
}

var fields_delete_template = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_workflow = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_workflow_step = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "StepGroupId", Flag: "step-group-id", Type: "*string", Required: true},
	{Name: "WorkflowId", Flag: "workflow-id", Type: "*string", Required: true},
}

var fields_delete_workflow_step_group = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "WorkflowId", Flag: "workflow-id", Type: "*string", Required: true},
}

var fields_get_template = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_template_step = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "StepGroupId", Flag: "step-group-id", Type: "*string", Required: true},
	{Name: "TemplateId", Flag: "template-id", Type: "*string", Required: true},
}

var fields_get_template_step_group = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "TemplateId", Flag: "template-id", Type: "*string", Required: true},
}

var fields_get_workflow = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_workflow_step = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "StepGroupId", Flag: "step-group-id", Type: "*string", Required: true},
	{Name: "WorkflowId", Flag: "workflow-id", Type: "*string", Required: true},
}

var fields_get_workflow_step_group = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "WorkflowId", Flag: "workflow-id", Type: "*string", Required: true},
}

var fields_list_plugins = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_template_step_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TemplateId", Flag: "template-id", Type: "*string", Required: true},
}

var fields_list_template_steps = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StepGroupId", Flag: "step-group-id", Type: "*string", Required: true},
	{Name: "TemplateId", Flag: "template-id", Type: "*string", Required: true},
}

var fields_list_templates = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_workflow_step_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkflowId", Flag: "workflow-id", Type: "*string", Required: true},
}

var fields_list_workflow_steps = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StepGroupId", Flag: "step-group-id", Type: "*string", Required: true},
	{Name: "WorkflowId", Flag: "workflow-id", Type: "*string", Required: true},
}

var fields_list_workflows = []leanruntime.Field{
	{Name: "AdsApplicationConfigurationName", Flag: "ads-application-configuration-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.MigrationWorkflowStatusEnum", Required: false},
	{Name: "TemplateId", Flag: "template-id", Type: "*string", Required: false},
}

var fields_retry_workflow_step = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "StepGroupId", Flag: "step-group-id", Type: "*string", Required: true},
	{Name: "WorkflowId", Flag: "workflow-id", Type: "*string", Required: true},
}

var fields_start_workflow = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_stop_workflow = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_template = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "TemplateDescription", Flag: "template-description", Type: "*string", Required: false},
	{Name: "TemplateName", Flag: "template-name", Type: "*string", Required: false},
}

var fields_update_workflow = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "InputParameters", Flag: "input-parameters", Type: "map[string]types.StepInput", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "StepTargets", Flag: "step-targets", Type: "[]string", Required: false},
}

var fields_update_workflow_step = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Next", Flag: "next", Type: "[]string", Required: false},
	{Name: "Outputs", Flag: "outputs", Type: "[]types.WorkflowStepOutput", Required: false},
	{Name: "Previous", Flag: "previous", Type: "[]string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.StepStatus", Required: false},
	{Name: "StepActionType", Flag: "step-action-type", Type: "types.StepActionType", Required: false},
	{Name: "StepGroupId", Flag: "step-group-id", Type: "*string", Required: true},
	{Name: "StepTarget", Flag: "step-target", Type: "[]string", Required: false},
	{Name: "WorkflowId", Flag: "workflow-id", Type: "*string", Required: true},
	{Name: "WorkflowStepAutomationConfiguration", Flag: "workflow-step-automation-configuration", Type: "*types.WorkflowStepAutomationConfiguration", Required: false},
}

var fields_update_workflow_step_group = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Next", Flag: "next", Type: "[]string", Required: false},
	{Name: "Previous", Flag: "previous", Type: "[]string", Required: false},
	{Name: "WorkflowId", Flag: "workflow-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-template": {
			Name:   "create-template",
			Fields: fields_create_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTemplate(ctx, input)
			},
		},
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
		"create-workflow-step": {
			Name:   "create-workflow-step",
			Fields: fields_create_workflow_step,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWorkflowStepInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_workflow_step, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWorkflowStep(ctx, input)
			},
		},
		"create-workflow-step-group": {
			Name:   "create-workflow-step-group",
			Fields: fields_create_workflow_step_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWorkflowStepGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_workflow_step_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWorkflowStepGroup(ctx, input)
			},
		},
		"delete-template": {
			Name:   "delete-template",
			Fields: fields_delete_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTemplate(ctx, input)
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
		"delete-workflow-step": {
			Name:   "delete-workflow-step",
			Fields: fields_delete_workflow_step,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWorkflowStepInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_workflow_step, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWorkflowStep(ctx, input)
			},
		},
		"delete-workflow-step-group": {
			Name:   "delete-workflow-step-group",
			Fields: fields_delete_workflow_step_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWorkflowStepGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_workflow_step_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWorkflowStepGroup(ctx, input)
			},
		},
		"get-template": {
			Name:   "get-template",
			Fields: fields_get_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTemplate(ctx, input)
			},
		},
		"get-template-step": {
			Name:   "get-template-step",
			Fields: fields_get_template_step,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTemplateStepInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_template_step, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTemplateStep(ctx, input)
			},
		},
		"get-template-step-group": {
			Name:   "get-template-step-group",
			Fields: fields_get_template_step_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTemplateStepGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_template_step_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTemplateStepGroup(ctx, input)
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
		"get-workflow-step": {
			Name:   "get-workflow-step",
			Fields: fields_get_workflow_step,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWorkflowStepInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_workflow_step, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWorkflowStep(ctx, input)
			},
		},
		"get-workflow-step-group": {
			Name:   "get-workflow-step-group",
			Fields: fields_get_workflow_step_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWorkflowStepGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_workflow_step_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWorkflowStepGroup(ctx, input)
			},
		},
		"list-plugins": {
			Name:   "list-plugins",
			Fields: fields_list_plugins,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPluginsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_plugins, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPlugins(ctx, input)
				}
				var results []*svc.ListPluginsOutput
				p := svc.NewListPluginsPaginator(client, input)
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
		"list-template-step-groups": {
			Name:   "list-template-step-groups",
			Fields: fields_list_template_step_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTemplateStepGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_template_step_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTemplateStepGroups(ctx, input)
				}
				var results []*svc.ListTemplateStepGroupsOutput
				p := svc.NewListTemplateStepGroupsPaginator(client, input)
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
		"list-template-steps": {
			Name:   "list-template-steps",
			Fields: fields_list_template_steps,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTemplateStepsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_template_steps, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTemplateSteps(ctx, input)
				}
				var results []*svc.ListTemplateStepsOutput
				p := svc.NewListTemplateStepsPaginator(client, input)
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
		"list-templates": {
			Name:   "list-templates",
			Fields: fields_list_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTemplates(ctx, input)
				}
				var results []*svc.ListTemplatesOutput
				p := svc.NewListTemplatesPaginator(client, input)
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
		"list-workflow-step-groups": {
			Name:   "list-workflow-step-groups",
			Fields: fields_list_workflow_step_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkflowStepGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_workflow_step_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkflowStepGroups(ctx, input)
				}
				var results []*svc.ListWorkflowStepGroupsOutput
				p := svc.NewListWorkflowStepGroupsPaginator(client, input)
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
		"list-workflow-steps": {
			Name:   "list-workflow-steps",
			Fields: fields_list_workflow_steps,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkflowStepsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_workflow_steps, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkflowSteps(ctx, input)
				}
				var results []*svc.ListWorkflowStepsOutput
				p := svc.NewListWorkflowStepsPaginator(client, input)
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
		"retry-workflow-step": {
			Name:   "retry-workflow-step",
			Fields: fields_retry_workflow_step,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RetryWorkflowStepInput{}
				if _, err := leanruntime.ApplyInput(input, fields_retry_workflow_step, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RetryWorkflowStep(ctx, input)
			},
		},
		"start-workflow": {
			Name:   "start-workflow",
			Fields: fields_start_workflow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartWorkflowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_workflow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartWorkflow(ctx, input)
			},
		},
		"stop-workflow": {
			Name:   "stop-workflow",
			Fields: fields_stop_workflow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopWorkflowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_workflow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopWorkflow(ctx, input)
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
		"update-template": {
			Name:   "update-template",
			Fields: fields_update_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTemplate(ctx, input)
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
		"update-workflow-step": {
			Name:   "update-workflow-step",
			Fields: fields_update_workflow_step,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWorkflowStepInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_workflow_step, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWorkflowStep(ctx, input)
			},
		},
		"update-workflow-step-group": {
			Name:   "update-workflow-step-group",
			Fields: fields_update_workflow_step_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWorkflowStepGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_workflow_step_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWorkflowStepGroup(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("migrationhuborchestrator", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
