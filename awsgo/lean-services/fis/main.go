package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/fis"
)

var fields_create_experiment_template = []leanruntime.Field{
	{Name: "Actions", Flag: "actions", Type: "map[string]types.CreateExperimentTemplateActionInput", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "ExperimentOptions", Flag: "experiment-options", Type: "*types.CreateExperimentTemplateExperimentOptionsInput", Required: false},
	{Name: "ExperimentReportConfiguration", Flag: "experiment-report-configuration", Type: "*types.CreateExperimentTemplateReportConfigurationInput", Required: false},
	{Name: "LogConfiguration", Flag: "log-configuration", Type: "*types.CreateExperimentTemplateLogConfigurationInput", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "StopConditions", Flag: "stop-conditions", Type: "[]types.CreateExperimentTemplateStopConditionInput", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Targets", Flag: "targets", Type: "map[string]types.CreateExperimentTemplateTargetInput", Required: false},
}

var fields_create_target_account_configuration = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ExperimentTemplateId", Flag: "experiment-template-id", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
}

var fields_delete_experiment_template = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_target_account_configuration = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ExperimentTemplateId", Flag: "experiment-template-id", Type: "*string", Required: true},
}

var fields_get_action = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_experiment = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_experiment_target_account_configuration = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ExperimentId", Flag: "experiment-id", Type: "*string", Required: true},
}

var fields_get_experiment_template = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_safety_lever = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_target_account_configuration = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ExperimentTemplateId", Flag: "experiment-template-id", Type: "*string", Required: true},
}

var fields_get_target_resource_type = []leanruntime.Field{
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: true},
}

var fields_list_actions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_experiment_resolved_targets = []leanruntime.Field{
	{Name: "ExperimentId", Flag: "experiment-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TargetName", Flag: "target-name", Type: "*string", Required: false},
}

var fields_list_experiment_target_account_configurations = []leanruntime.Field{
	{Name: "ExperimentId", Flag: "experiment-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_experiment_templates = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_experiments = []leanruntime.Field{
	{Name: "ExperimentTemplateId", Flag: "experiment-template-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_target_account_configurations = []leanruntime.Field{
	{Name: "ExperimentTemplateId", Flag: "experiment-template-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_target_resource_types = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_start_experiment = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "ExperimentOptions", Flag: "experiment-options", Type: "*types.StartExperimentExperimentOptionsInput", Required: false},
	{Name: "ExperimentTemplateId", Flag: "experiment-template-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_stop_experiment = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: false},
}

var fields_update_experiment_template = []leanruntime.Field{
	{Name: "Actions", Flag: "actions", Type: "map[string]types.UpdateExperimentTemplateActionInputItem", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ExperimentOptions", Flag: "experiment-options", Type: "*types.UpdateExperimentTemplateExperimentOptionsInput", Required: false},
	{Name: "ExperimentReportConfiguration", Flag: "experiment-report-configuration", Type: "*types.UpdateExperimentTemplateReportConfigurationInput", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "LogConfiguration", Flag: "log-configuration", Type: "*types.UpdateExperimentTemplateLogConfigurationInput", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "StopConditions", Flag: "stop-conditions", Type: "[]types.UpdateExperimentTemplateStopConditionInput", Required: false},
	{Name: "Targets", Flag: "targets", Type: "map[string]types.UpdateExperimentTemplateTargetInput", Required: false},
}

var fields_update_safety_lever_state = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "State", Flag: "state", Type: "*types.UpdateSafetyLeverStateInput", Required: true},
}

var fields_update_target_account_configuration = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ExperimentTemplateId", Flag: "experiment-template-id", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-experiment-template": {
			Name:   "create-experiment-template",
			Fields: fields_create_experiment_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateExperimentTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_experiment_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateExperimentTemplate(ctx, input)
			},
		},
		"create-target-account-configuration": {
			Name:   "create-target-account-configuration",
			Fields: fields_create_target_account_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTargetAccountConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_target_account_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTargetAccountConfiguration(ctx, input)
			},
		},
		"delete-experiment-template": {
			Name:   "delete-experiment-template",
			Fields: fields_delete_experiment_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteExperimentTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_experiment_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteExperimentTemplate(ctx, input)
			},
		},
		"delete-target-account-configuration": {
			Name:   "delete-target-account-configuration",
			Fields: fields_delete_target_account_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTargetAccountConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_target_account_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTargetAccountConfiguration(ctx, input)
			},
		},
		"get-action": {
			Name:   "get-action",
			Fields: fields_get_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAction(ctx, input)
			},
		},
		"get-experiment": {
			Name:   "get-experiment",
			Fields: fields_get_experiment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetExperimentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_experiment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetExperiment(ctx, input)
			},
		},
		"get-experiment-target-account-configuration": {
			Name:   "get-experiment-target-account-configuration",
			Fields: fields_get_experiment_target_account_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetExperimentTargetAccountConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_experiment_target_account_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetExperimentTargetAccountConfiguration(ctx, input)
			},
		},
		"get-experiment-template": {
			Name:   "get-experiment-template",
			Fields: fields_get_experiment_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetExperimentTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_experiment_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetExperimentTemplate(ctx, input)
			},
		},
		"get-safety-lever": {
			Name:   "get-safety-lever",
			Fields: fields_get_safety_lever,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSafetyLeverInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_safety_lever, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSafetyLever(ctx, input)
			},
		},
		"get-target-account-configuration": {
			Name:   "get-target-account-configuration",
			Fields: fields_get_target_account_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTargetAccountConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_target_account_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTargetAccountConfiguration(ctx, input)
			},
		},
		"get-target-resource-type": {
			Name:   "get-target-resource-type",
			Fields: fields_get_target_resource_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTargetResourceTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_target_resource_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTargetResourceType(ctx, input)
			},
		},
		"list-actions": {
			Name:   "list-actions",
			Fields: fields_list_actions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListActionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_actions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListActions(ctx, input)
				}
				var results []*svc.ListActionsOutput
				p := svc.NewListActionsPaginator(client, input)
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
		"list-experiment-resolved-targets": {
			Name:   "list-experiment-resolved-targets",
			Fields: fields_list_experiment_resolved_targets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListExperimentResolvedTargetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_experiment_resolved_targets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListExperimentResolvedTargets(ctx, input)
				}
				var results []*svc.ListExperimentResolvedTargetsOutput
				p := svc.NewListExperimentResolvedTargetsPaginator(client, input)
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
		"list-experiment-target-account-configurations": {
			Name:   "list-experiment-target-account-configurations",
			Fields: fields_list_experiment_target_account_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListExperimentTargetAccountConfigurationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_experiment_target_account_configurations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListExperimentTargetAccountConfigurations(ctx, input)
			},
		},
		"list-experiment-templates": {
			Name:   "list-experiment-templates",
			Fields: fields_list_experiment_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListExperimentTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_experiment_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListExperimentTemplates(ctx, input)
				}
				var results []*svc.ListExperimentTemplatesOutput
				p := svc.NewListExperimentTemplatesPaginator(client, input)
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
		"list-experiments": {
			Name:   "list-experiments",
			Fields: fields_list_experiments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListExperimentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_experiments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListExperiments(ctx, input)
				}
				var results []*svc.ListExperimentsOutput
				p := svc.NewListExperimentsPaginator(client, input)
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
		"list-target-account-configurations": {
			Name:   "list-target-account-configurations",
			Fields: fields_list_target_account_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTargetAccountConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_target_account_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTargetAccountConfigurations(ctx, input)
				}
				var results []*svc.ListTargetAccountConfigurationsOutput
				p := svc.NewListTargetAccountConfigurationsPaginator(client, input)
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
		"list-target-resource-types": {
			Name:   "list-target-resource-types",
			Fields: fields_list_target_resource_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTargetResourceTypesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_target_resource_types, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTargetResourceTypes(ctx, input)
				}
				var results []*svc.ListTargetResourceTypesOutput
				p := svc.NewListTargetResourceTypesPaginator(client, input)
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
		"start-experiment": {
			Name:   "start-experiment",
			Fields: fields_start_experiment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartExperimentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_experiment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartExperiment(ctx, input)
			},
		},
		"stop-experiment": {
			Name:   "stop-experiment",
			Fields: fields_stop_experiment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopExperimentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_experiment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopExperiment(ctx, input)
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
		"update-experiment-template": {
			Name:   "update-experiment-template",
			Fields: fields_update_experiment_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateExperimentTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_experiment_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateExperimentTemplate(ctx, input)
			},
		},
		"update-safety-lever-state": {
			Name:   "update-safety-lever-state",
			Fields: fields_update_safety_lever_state,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSafetyLeverStateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_safety_lever_state, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSafetyLeverState(ctx, input)
			},
		},
		"update-target-account-configuration": {
			Name:   "update-target-account-configuration",
			Fields: fields_update_target_account_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTargetAccountConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_target_account_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTargetAccountConfiguration(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("fis", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
