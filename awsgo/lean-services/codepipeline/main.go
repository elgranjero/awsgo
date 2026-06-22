package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/codepipeline"
)

var fields_acknowledge_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "Nonce", Flag: "nonce", Type: "*string", Required: true},
}

var fields_acknowledge_third_party_job = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "Nonce", Flag: "nonce", Type: "*string", Required: true},
}

var fields_create_custom_action_type = []leanruntime.Field{
	{Name: "Category", Flag: "category", Type: "types.ActionCategory", Required: true},
	{Name: "ConfigurationProperties", Flag: "configuration-properties", Type: "[]types.ActionConfigurationProperty", Required: false},
	{Name: "InputArtifactDetails", Flag: "input-artifact-details", Type: "*types.ArtifactDetails", Required: true},
	{Name: "OutputArtifactDetails", Flag: "output-artifact-details", Type: "*types.ArtifactDetails", Required: true},
	{Name: "Provider", Flag: "provider", Type: "*string", Required: true},
	{Name: "Settings", Flag: "settings", Type: "*types.ActionTypeSettings", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Version", Flag: "version", Type: "*string", Required: true},
}

var fields_create_pipeline = []leanruntime.Field{
	{Name: "Pipeline", Flag: "pipeline", Type: "*types.PipelineDeclaration", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_custom_action_type = []leanruntime.Field{
	{Name: "Category", Flag: "category", Type: "types.ActionCategory", Required: true},
	{Name: "Provider", Flag: "provider", Type: "*string", Required: true},
	{Name: "Version", Flag: "version", Type: "*string", Required: true},
}

var fields_delete_pipeline = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_webhook = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_deregister_webhook_with_third_party = []leanruntime.Field{
	{Name: "WebhookName", Flag: "webhook-name", Type: "*string", Required: false},
}

var fields_disable_stage_transition = []leanruntime.Field{
	{Name: "PipelineName", Flag: "pipeline-name", Type: "*string", Required: true},
	{Name: "Reason", Flag: "reason", Type: "*string", Required: true},
	{Name: "StageName", Flag: "stage-name", Type: "*string", Required: true},
	{Name: "TransitionType", Flag: "transition-type", Type: "types.StageTransitionType", Required: true},
}

var fields_enable_stage_transition = []leanruntime.Field{
	{Name: "PipelineName", Flag: "pipeline-name", Type: "*string", Required: true},
	{Name: "StageName", Flag: "stage-name", Type: "*string", Required: true},
	{Name: "TransitionType", Flag: "transition-type", Type: "types.StageTransitionType", Required: true},
}

var fields_get_action_type = []leanruntime.Field{
	{Name: "Category", Flag: "category", Type: "types.ActionCategory", Required: true},
	{Name: "Owner", Flag: "owner", Type: "*string", Required: true},
	{Name: "Provider", Flag: "provider", Type: "*string", Required: true},
	{Name: "Version", Flag: "version", Type: "*string", Required: true},
}

var fields_get_job_details = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_get_pipeline = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Version", Flag: "version", Type: "*int32", Required: false},
}

var fields_get_pipeline_execution = []leanruntime.Field{
	{Name: "PipelineExecutionId", Flag: "pipeline-execution-id", Type: "*string", Required: true},
	{Name: "PipelineName", Flag: "pipeline-name", Type: "*string", Required: true},
}

var fields_get_pipeline_state = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_third_party_job_details = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_list_action_executions = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.ActionExecutionFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PipelineName", Flag: "pipeline-name", Type: "*string", Required: true},
}

var fields_list_action_types = []leanruntime.Field{
	{Name: "ActionOwnerFilter", Flag: "action-owner-filter", Type: "types.ActionOwner", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RegionFilter", Flag: "region-filter", Type: "*string", Required: false},
}

var fields_list_deploy_action_execution_targets = []leanruntime.Field{
	{Name: "ActionExecutionId", Flag: "action-execution-id", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "[]types.TargetFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PipelineName", Flag: "pipeline-name", Type: "*string", Required: false},
}

var fields_list_pipeline_executions = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.PipelineExecutionFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PipelineName", Flag: "pipeline-name", Type: "*string", Required: true},
}

var fields_list_pipelines = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_rule_executions = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.RuleExecutionFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PipelineName", Flag: "pipeline-name", Type: "*string", Required: true},
}

var fields_list_rule_types = []leanruntime.Field{
	{Name: "RegionFilter", Flag: "region-filter", Type: "*string", Required: false},
	{Name: "RuleOwnerFilter", Flag: "rule-owner-filter", Type: "types.RuleOwner", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_webhooks = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_override_stage_condition = []leanruntime.Field{
	{Name: "ConditionType", Flag: "condition-type", Type: "types.ConditionType", Required: true},
	{Name: "PipelineExecutionId", Flag: "pipeline-execution-id", Type: "*string", Required: true},
	{Name: "PipelineName", Flag: "pipeline-name", Type: "*string", Required: true},
	{Name: "StageName", Flag: "stage-name", Type: "*string", Required: true},
}

var fields_poll_for_jobs = []leanruntime.Field{
	{Name: "ActionTypeId", Flag: "action-type-id", Type: "*types.ActionTypeId", Required: true},
	{Name: "MaxBatchSize", Flag: "max-batch-size", Type: "*int32", Required: false},
	{Name: "QueryParam", Flag: "query-param", Type: "map[string]string", Required: false},
}

var fields_poll_for_third_party_jobs = []leanruntime.Field{
	{Name: "ActionTypeId", Flag: "action-type-id", Type: "*types.ActionTypeId", Required: true},
	{Name: "MaxBatchSize", Flag: "max-batch-size", Type: "*int32", Required: false},
}

var fields_put_action_revision = []leanruntime.Field{
	{Name: "ActionName", Flag: "action-name", Type: "*string", Required: true},
	{Name: "ActionRevision", Flag: "action-revision", Type: "*types.ActionRevision", Required: true},
	{Name: "PipelineName", Flag: "pipeline-name", Type: "*string", Required: true},
	{Name: "StageName", Flag: "stage-name", Type: "*string", Required: true},
}

var fields_put_approval_result = []leanruntime.Field{
	{Name: "ActionName", Flag: "action-name", Type: "*string", Required: true},
	{Name: "PipelineName", Flag: "pipeline-name", Type: "*string", Required: true},
	{Name: "Result", Flag: "result", Type: "*types.ApprovalResult", Required: true},
	{Name: "StageName", Flag: "stage-name", Type: "*string", Required: true},
	{Name: "Token", Flag: "token", Type: "*string", Required: true},
}

var fields_put_job_failure_result = []leanruntime.Field{
	{Name: "FailureDetails", Flag: "failure-details", Type: "*types.FailureDetails", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_put_job_success_result = []leanruntime.Field{
	{Name: "ContinuationToken", Flag: "continuation-token", Type: "*string", Required: false},
	{Name: "CurrentRevision", Flag: "current-revision", Type: "*types.CurrentRevision", Required: false},
	{Name: "ExecutionDetails", Flag: "execution-details", Type: "*types.ExecutionDetails", Required: false},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "OutputVariables", Flag: "output-variables", Type: "map[string]string", Required: false},
}

var fields_put_third_party_job_failure_result = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "FailureDetails", Flag: "failure-details", Type: "*types.FailureDetails", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_put_third_party_job_success_result = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "ContinuationToken", Flag: "continuation-token", Type: "*string", Required: false},
	{Name: "CurrentRevision", Flag: "current-revision", Type: "*types.CurrentRevision", Required: false},
	{Name: "ExecutionDetails", Flag: "execution-details", Type: "*types.ExecutionDetails", Required: false},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_put_webhook = []leanruntime.Field{
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Webhook", Flag: "webhook", Type: "*types.WebhookDefinition", Required: true},
}

var fields_register_webhook_with_third_party = []leanruntime.Field{
	{Name: "WebhookName", Flag: "webhook-name", Type: "*string", Required: false},
}

var fields_retry_stage_execution = []leanruntime.Field{
	{Name: "PipelineExecutionId", Flag: "pipeline-execution-id", Type: "*string", Required: true},
	{Name: "PipelineName", Flag: "pipeline-name", Type: "*string", Required: true},
	{Name: "RetryMode", Flag: "retry-mode", Type: "types.StageRetryMode", Required: true},
	{Name: "StageName", Flag: "stage-name", Type: "*string", Required: true},
}

var fields_rollback_stage = []leanruntime.Field{
	{Name: "PipelineName", Flag: "pipeline-name", Type: "*string", Required: true},
	{Name: "StageName", Flag: "stage-name", Type: "*string", Required: true},
	{Name: "TargetPipelineExecutionId", Flag: "target-pipeline-execution-id", Type: "*string", Required: true},
}

var fields_start_pipeline_execution = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SourceRevisions", Flag: "source-revisions", Type: "[]types.SourceRevisionOverride", Required: false},
	{Name: "Variables", Flag: "variables", Type: "[]types.PipelineVariable", Required: false},
}

var fields_stop_pipeline_execution = []leanruntime.Field{
	{Name: "Abandon", Flag: "abandon", Type: "bool", Required: false},
	{Name: "PipelineExecutionId", Flag: "pipeline-execution-id", Type: "*string", Required: true},
	{Name: "PipelineName", Flag: "pipeline-name", Type: "*string", Required: true},
	{Name: "Reason", Flag: "reason", Type: "*string", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_action_type = []leanruntime.Field{
	{Name: "ActionType", Flag: "action-type", Type: "*types.ActionTypeDeclaration", Required: true},
}

var fields_update_pipeline = []leanruntime.Field{
	{Name: "Pipeline", Flag: "pipeline", Type: "*types.PipelineDeclaration", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"acknowledge-job": {
			Name:   "acknowledge-job",
			Fields: fields_acknowledge_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcknowledgeJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_acknowledge_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcknowledgeJob(ctx, input)
			},
		},
		"acknowledge-third-party-job": {
			Name:   "acknowledge-third-party-job",
			Fields: fields_acknowledge_third_party_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcknowledgeThirdPartyJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_acknowledge_third_party_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcknowledgeThirdPartyJob(ctx, input)
			},
		},
		"create-custom-action-type": {
			Name:   "create-custom-action-type",
			Fields: fields_create_custom_action_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCustomActionTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_custom_action_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCustomActionType(ctx, input)
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
		"delete-custom-action-type": {
			Name:   "delete-custom-action-type",
			Fields: fields_delete_custom_action_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCustomActionTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_custom_action_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCustomActionType(ctx, input)
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
		"delete-webhook": {
			Name:   "delete-webhook",
			Fields: fields_delete_webhook,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWebhookInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_webhook, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWebhook(ctx, input)
			},
		},
		"deregister-webhook-with-third-party": {
			Name:   "deregister-webhook-with-third-party",
			Fields: fields_deregister_webhook_with_third_party,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterWebhookWithThirdPartyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_webhook_with_third_party, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterWebhookWithThirdParty(ctx, input)
			},
		},
		"disable-stage-transition": {
			Name:   "disable-stage-transition",
			Fields: fields_disable_stage_transition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableStageTransitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_stage_transition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableStageTransition(ctx, input)
			},
		},
		"enable-stage-transition": {
			Name:   "enable-stage-transition",
			Fields: fields_enable_stage_transition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableStageTransitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_stage_transition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableStageTransition(ctx, input)
			},
		},
		"get-action-type": {
			Name:   "get-action-type",
			Fields: fields_get_action_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetActionTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_action_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetActionType(ctx, input)
			},
		},
		"get-job-details": {
			Name:   "get-job-details",
			Fields: fields_get_job_details,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetJobDetailsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_job_details, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetJobDetails(ctx, input)
			},
		},
		"get-pipeline": {
			Name:   "get-pipeline",
			Fields: fields_get_pipeline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPipelineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_pipeline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPipeline(ctx, input)
			},
		},
		"get-pipeline-execution": {
			Name:   "get-pipeline-execution",
			Fields: fields_get_pipeline_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPipelineExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_pipeline_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPipelineExecution(ctx, input)
			},
		},
		"get-pipeline-state": {
			Name:   "get-pipeline-state",
			Fields: fields_get_pipeline_state,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPipelineStateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_pipeline_state, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPipelineState(ctx, input)
			},
		},
		"get-third-party-job-details": {
			Name:   "get-third-party-job-details",
			Fields: fields_get_third_party_job_details,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetThirdPartyJobDetailsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_third_party_job_details, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetThirdPartyJobDetails(ctx, input)
			},
		},
		"list-action-executions": {
			Name:   "list-action-executions",
			Fields: fields_list_action_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListActionExecutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_action_executions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListActionExecutions(ctx, input)
				}
				var results []*svc.ListActionExecutionsOutput
				p := svc.NewListActionExecutionsPaginator(client, input)
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
		"list-action-types": {
			Name:   "list-action-types",
			Fields: fields_list_action_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListActionTypesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_action_types, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListActionTypes(ctx, input)
				}
				var results []*svc.ListActionTypesOutput
				p := svc.NewListActionTypesPaginator(client, input)
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
		"list-deploy-action-execution-targets": {
			Name:   "list-deploy-action-execution-targets",
			Fields: fields_list_deploy_action_execution_targets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDeployActionExecutionTargetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_deploy_action_execution_targets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDeployActionExecutionTargets(ctx, input)
				}
				var results []*svc.ListDeployActionExecutionTargetsOutput
				p := svc.NewListDeployActionExecutionTargetsPaginator(client, input)
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
		"list-pipeline-executions": {
			Name:   "list-pipeline-executions",
			Fields: fields_list_pipeline_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPipelineExecutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_pipeline_executions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPipelineExecutions(ctx, input)
				}
				var results []*svc.ListPipelineExecutionsOutput
				p := svc.NewListPipelineExecutionsPaginator(client, input)
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
		"list-rule-executions": {
			Name:   "list-rule-executions",
			Fields: fields_list_rule_executions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRuleExecutionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_rule_executions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRuleExecutions(ctx, input)
				}
				var results []*svc.ListRuleExecutionsOutput
				p := svc.NewListRuleExecutionsPaginator(client, input)
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
		"list-rule-types": {
			Name:   "list-rule-types",
			Fields: fields_list_rule_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRuleTypesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_rule_types, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListRuleTypes(ctx, input)
			},
		},
		"list-tags-for-resource": {
			Name:   "list-tags-for-resource",
			Fields: fields_list_tags_for_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsForResourceInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tags_for_resource, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTagsForResource(ctx, input)
				}
				var results []*svc.ListTagsForResourceOutput
				p := svc.NewListTagsForResourcePaginator(client, input)
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
		"list-webhooks": {
			Name:   "list-webhooks",
			Fields: fields_list_webhooks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWebhooksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_webhooks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWebhooks(ctx, input)
				}
				var results []*svc.ListWebhooksOutput
				p := svc.NewListWebhooksPaginator(client, input)
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
		"override-stage-condition": {
			Name:   "override-stage-condition",
			Fields: fields_override_stage_condition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.OverrideStageConditionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_override_stage_condition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.OverrideStageCondition(ctx, input)
			},
		},
		"poll-for-jobs": {
			Name:   "poll-for-jobs",
			Fields: fields_poll_for_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PollForJobsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_poll_for_jobs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PollForJobs(ctx, input)
			},
		},
		"poll-for-third-party-jobs": {
			Name:   "poll-for-third-party-jobs",
			Fields: fields_poll_for_third_party_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PollForThirdPartyJobsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_poll_for_third_party_jobs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PollForThirdPartyJobs(ctx, input)
			},
		},
		"put-action-revision": {
			Name:   "put-action-revision",
			Fields: fields_put_action_revision,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutActionRevisionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_action_revision, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutActionRevision(ctx, input)
			},
		},
		"put-approval-result": {
			Name:   "put-approval-result",
			Fields: fields_put_approval_result,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutApprovalResultInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_approval_result, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutApprovalResult(ctx, input)
			},
		},
		"put-job-failure-result": {
			Name:   "put-job-failure-result",
			Fields: fields_put_job_failure_result,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutJobFailureResultInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_job_failure_result, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutJobFailureResult(ctx, input)
			},
		},
		"put-job-success-result": {
			Name:   "put-job-success-result",
			Fields: fields_put_job_success_result,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutJobSuccessResultInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_job_success_result, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutJobSuccessResult(ctx, input)
			},
		},
		"put-third-party-job-failure-result": {
			Name:   "put-third-party-job-failure-result",
			Fields: fields_put_third_party_job_failure_result,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutThirdPartyJobFailureResultInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_third_party_job_failure_result, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutThirdPartyJobFailureResult(ctx, input)
			},
		},
		"put-third-party-job-success-result": {
			Name:   "put-third-party-job-success-result",
			Fields: fields_put_third_party_job_success_result,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutThirdPartyJobSuccessResultInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_third_party_job_success_result, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutThirdPartyJobSuccessResult(ctx, input)
			},
		},
		"put-webhook": {
			Name:   "put-webhook",
			Fields: fields_put_webhook,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutWebhookInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_webhook, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutWebhook(ctx, input)
			},
		},
		"register-webhook-with-third-party": {
			Name:   "register-webhook-with-third-party",
			Fields: fields_register_webhook_with_third_party,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterWebhookWithThirdPartyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_webhook_with_third_party, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterWebhookWithThirdParty(ctx, input)
			},
		},
		"retry-stage-execution": {
			Name:   "retry-stage-execution",
			Fields: fields_retry_stage_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RetryStageExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_retry_stage_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RetryStageExecution(ctx, input)
			},
		},
		"rollback-stage": {
			Name:   "rollback-stage",
			Fields: fields_rollback_stage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RollbackStageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_rollback_stage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RollbackStage(ctx, input)
			},
		},
		"start-pipeline-execution": {
			Name:   "start-pipeline-execution",
			Fields: fields_start_pipeline_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartPipelineExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_pipeline_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartPipelineExecution(ctx, input)
			},
		},
		"stop-pipeline-execution": {
			Name:   "stop-pipeline-execution",
			Fields: fields_stop_pipeline_execution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopPipelineExecutionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_pipeline_execution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopPipelineExecution(ctx, input)
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
		"update-action-type": {
			Name:   "update-action-type",
			Fields: fields_update_action_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateActionTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_action_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateActionType(ctx, input)
			},
		},
		"update-pipeline": {
			Name:   "update-pipeline",
			Fields: fields_update_pipeline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePipelineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_pipeline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePipeline(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("codepipeline", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
