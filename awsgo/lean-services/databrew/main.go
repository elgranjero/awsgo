package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/databrew"
)

var fields_batch_delete_recipe_version = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RecipeVersions", Flag: "recipe-versions", Type: "[]string", Required: true},
}

var fields_create_dataset = []leanruntime.Field{
	{Name: "Format", Flag: "format", Type: "types.InputFormat", Required: false},
	{Name: "FormatOptions", Flag: "format-options", Type: "*types.FormatOptions", Required: false},
	{Name: "Input", Flag: "input", Type: "*types.Input", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PathOptions", Flag: "path-options", Type: "*types.PathOptions", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_profile_job = []leanruntime.Field{
	{Name: "Configuration", Flag: "configuration", Type: "*types.ProfileConfiguration", Required: false},
	{Name: "DatasetName", Flag: "dataset-name", Type: "*string", Required: true},
	{Name: "EncryptionKeyArn", Flag: "encryption-key-arn", Type: "*string", Required: false},
	{Name: "EncryptionMode", Flag: "encryption-mode", Type: "types.EncryptionMode", Required: false},
	{Name: "JobSample", Flag: "job-sample", Type: "*types.JobSample", Required: false},
	{Name: "LogSubscription", Flag: "log-subscription", Type: "types.LogSubscription", Required: false},
	{Name: "MaxCapacity", Flag: "max-capacity", Type: "int32", Required: false},
	{Name: "MaxRetries", Flag: "max-retries", Type: "int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OutputLocation", Flag: "output-location", Type: "*types.S3Location", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Timeout", Flag: "timeout", Type: "int32", Required: false},
	{Name: "ValidationConfigurations", Flag: "validation-configurations", Type: "[]types.ValidationConfiguration", Required: false},
}

var fields_create_project = []leanruntime.Field{
	{Name: "DatasetName", Flag: "dataset-name", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RecipeName", Flag: "recipe-name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Sample", Flag: "sample", Type: "*types.Sample", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_recipe = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Steps", Flag: "steps", Type: "[]types.RecipeStep", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_recipe_job = []leanruntime.Field{
	{Name: "DataCatalogOutputs", Flag: "data-catalog-outputs", Type: "[]types.DataCatalogOutput", Required: false},
	{Name: "DatabaseOutputs", Flag: "database-outputs", Type: "[]types.DatabaseOutput", Required: false},
	{Name: "DatasetName", Flag: "dataset-name", Type: "*string", Required: false},
	{Name: "EncryptionKeyArn", Flag: "encryption-key-arn", Type: "*string", Required: false},
	{Name: "EncryptionMode", Flag: "encryption-mode", Type: "types.EncryptionMode", Required: false},
	{Name: "LogSubscription", Flag: "log-subscription", Type: "types.LogSubscription", Required: false},
	{Name: "MaxCapacity", Flag: "max-capacity", Type: "int32", Required: false},
	{Name: "MaxRetries", Flag: "max-retries", Type: "int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Outputs", Flag: "outputs", Type: "[]types.Output", Required: false},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: false},
	{Name: "RecipeReference", Flag: "recipe-reference", Type: "*types.RecipeReference", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Timeout", Flag: "timeout", Type: "int32", Required: false},
}

var fields_create_ruleset = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Rules", Flag: "rules", Type: "[]types.Rule", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TargetArn", Flag: "target-arn", Type: "*string", Required: true},
}

var fields_create_schedule = []leanruntime.Field{
	{Name: "CronExpression", Flag: "cron-expression", Type: "*string", Required: true},
	{Name: "JobNames", Flag: "job-names", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_dataset = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_job = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_project = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_recipe_version = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RecipeVersion", Flag: "recipe-version", Type: "*string", Required: true},
}

var fields_delete_ruleset = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_schedule = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_describe_dataset = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_describe_job = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_describe_job_run = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RunId", Flag: "run-id", Type: "*string", Required: true},
}

var fields_describe_project = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_describe_recipe = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RecipeVersion", Flag: "recipe-version", Type: "*string", Required: false},
}

var fields_describe_ruleset = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_describe_schedule = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_list_datasets = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_job_runs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_jobs = []leanruntime.Field{
	{Name: "DatasetName", Flag: "dataset-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: false},
}

var fields_list_projects = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_recipe_versions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_recipes = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RecipeVersion", Flag: "recipe-version", Type: "*string", Required: false},
}

var fields_list_rulesets = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TargetArn", Flag: "target-arn", Type: "*string", Required: false},
}

var fields_list_schedules = []leanruntime.Field{
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_publish_recipe = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_send_project_session_action = []leanruntime.Field{
	{Name: "ClientSessionId", Flag: "client-session-id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Preview", Flag: "preview", Type: "bool", Required: false},
	{Name: "RecipeStep", Flag: "recipe-step", Type: "*types.RecipeStep", Required: false},
	{Name: "StepIndex", Flag: "step-index", Type: "*int32", Required: false},
	{Name: "ViewFrame", Flag: "view-frame", Type: "*types.ViewFrame", Required: false},
}

var fields_start_job_run = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_start_project_session = []leanruntime.Field{
	{Name: "AssumeControl", Flag: "assume-control", Type: "bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_stop_job_run = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RunId", Flag: "run-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_dataset = []leanruntime.Field{
	{Name: "Format", Flag: "format", Type: "types.InputFormat", Required: false},
	{Name: "FormatOptions", Flag: "format-options", Type: "*types.FormatOptions", Required: false},
	{Name: "Input", Flag: "input", Type: "*types.Input", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PathOptions", Flag: "path-options", Type: "*types.PathOptions", Required: false},
}

var fields_update_profile_job = []leanruntime.Field{
	{Name: "Configuration", Flag: "configuration", Type: "*types.ProfileConfiguration", Required: false},
	{Name: "EncryptionKeyArn", Flag: "encryption-key-arn", Type: "*string", Required: false},
	{Name: "EncryptionMode", Flag: "encryption-mode", Type: "types.EncryptionMode", Required: false},
	{Name: "JobSample", Flag: "job-sample", Type: "*types.JobSample", Required: false},
	{Name: "LogSubscription", Flag: "log-subscription", Type: "types.LogSubscription", Required: false},
	{Name: "MaxCapacity", Flag: "max-capacity", Type: "int32", Required: false},
	{Name: "MaxRetries", Flag: "max-retries", Type: "int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OutputLocation", Flag: "output-location", Type: "*types.S3Location", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Timeout", Flag: "timeout", Type: "int32", Required: false},
	{Name: "ValidationConfigurations", Flag: "validation-configurations", Type: "[]types.ValidationConfiguration", Required: false},
}

var fields_update_project = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Sample", Flag: "sample", Type: "*types.Sample", Required: false},
}

var fields_update_recipe = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Steps", Flag: "steps", Type: "[]types.RecipeStep", Required: false},
}

var fields_update_recipe_job = []leanruntime.Field{
	{Name: "DataCatalogOutputs", Flag: "data-catalog-outputs", Type: "[]types.DataCatalogOutput", Required: false},
	{Name: "DatabaseOutputs", Flag: "database-outputs", Type: "[]types.DatabaseOutput", Required: false},
	{Name: "EncryptionKeyArn", Flag: "encryption-key-arn", Type: "*string", Required: false},
	{Name: "EncryptionMode", Flag: "encryption-mode", Type: "types.EncryptionMode", Required: false},
	{Name: "LogSubscription", Flag: "log-subscription", Type: "types.LogSubscription", Required: false},
	{Name: "MaxCapacity", Flag: "max-capacity", Type: "int32", Required: false},
	{Name: "MaxRetries", Flag: "max-retries", Type: "int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Outputs", Flag: "outputs", Type: "[]types.Output", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Timeout", Flag: "timeout", Type: "int32", Required: false},
}

var fields_update_ruleset = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Rules", Flag: "rules", Type: "[]types.Rule", Required: true},
}

var fields_update_schedule = []leanruntime.Field{
	{Name: "CronExpression", Flag: "cron-expression", Type: "*string", Required: true},
	{Name: "JobNames", Flag: "job-names", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-delete-recipe-version": {
			Name:   "batch-delete-recipe-version",
			Fields: fields_batch_delete_recipe_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeleteRecipeVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete_recipe_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDeleteRecipeVersion(ctx, input)
			},
		},
		"create-dataset": {
			Name:   "create-dataset",
			Fields: fields_create_dataset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDatasetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_dataset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataset(ctx, input)
			},
		},
		"create-profile-job": {
			Name:   "create-profile-job",
			Fields: fields_create_profile_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProfileJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_profile_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProfileJob(ctx, input)
			},
		},
		"create-project": {
			Name:   "create-project",
			Fields: fields_create_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProject(ctx, input)
			},
		},
		"create-recipe": {
			Name:   "create-recipe",
			Fields: fields_create_recipe,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRecipeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_recipe, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRecipe(ctx, input)
			},
		},
		"create-recipe-job": {
			Name:   "create-recipe-job",
			Fields: fields_create_recipe_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRecipeJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_recipe_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRecipeJob(ctx, input)
			},
		},
		"create-ruleset": {
			Name:   "create-ruleset",
			Fields: fields_create_ruleset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRulesetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_ruleset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRuleset(ctx, input)
			},
		},
		"create-schedule": {
			Name:   "create-schedule",
			Fields: fields_create_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSchedule(ctx, input)
			},
		},
		"delete-dataset": {
			Name:   "delete-dataset",
			Fields: fields_delete_dataset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDatasetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_dataset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataset(ctx, input)
			},
		},
		"delete-job": {
			Name:   "delete-job",
			Fields: fields_delete_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteJob(ctx, input)
			},
		},
		"delete-project": {
			Name:   "delete-project",
			Fields: fields_delete_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProject(ctx, input)
			},
		},
		"delete-recipe-version": {
			Name:   "delete-recipe-version",
			Fields: fields_delete_recipe_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRecipeVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_recipe_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRecipeVersion(ctx, input)
			},
		},
		"delete-ruleset": {
			Name:   "delete-ruleset",
			Fields: fields_delete_ruleset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRulesetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ruleset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRuleset(ctx, input)
			},
		},
		"delete-schedule": {
			Name:   "delete-schedule",
			Fields: fields_delete_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSchedule(ctx, input)
			},
		},
		"describe-dataset": {
			Name:   "describe-dataset",
			Fields: fields_describe_dataset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDatasetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_dataset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDataset(ctx, input)
			},
		},
		"describe-job": {
			Name:   "describe-job",
			Fields: fields_describe_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeJob(ctx, input)
			},
		},
		"describe-job-run": {
			Name:   "describe-job-run",
			Fields: fields_describe_job_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeJobRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_job_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeJobRun(ctx, input)
			},
		},
		"describe-project": {
			Name:   "describe-project",
			Fields: fields_describe_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeProject(ctx, input)
			},
		},
		"describe-recipe": {
			Name:   "describe-recipe",
			Fields: fields_describe_recipe,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRecipeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_recipe, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRecipe(ctx, input)
			},
		},
		"describe-ruleset": {
			Name:   "describe-ruleset",
			Fields: fields_describe_ruleset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRulesetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_ruleset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRuleset(ctx, input)
			},
		},
		"describe-schedule": {
			Name:   "describe-schedule",
			Fields: fields_describe_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSchedule(ctx, input)
			},
		},
		"list-datasets": {
			Name:   "list-datasets",
			Fields: fields_list_datasets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDatasetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_datasets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDatasets(ctx, input)
				}
				var results []*svc.ListDatasetsOutput
				p := svc.NewListDatasetsPaginator(client, input)
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
		"list-job-runs": {
			Name:   "list-job-runs",
			Fields: fields_list_job_runs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListJobRunsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_job_runs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListJobRuns(ctx, input)
				}
				var results []*svc.ListJobRunsOutput
				p := svc.NewListJobRunsPaginator(client, input)
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
		"list-projects": {
			Name:   "list-projects",
			Fields: fields_list_projects,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProjectsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_projects, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProjects(ctx, input)
				}
				var results []*svc.ListProjectsOutput
				p := svc.NewListProjectsPaginator(client, input)
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
		"list-recipe-versions": {
			Name:   "list-recipe-versions",
			Fields: fields_list_recipe_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRecipeVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_recipe_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRecipeVersions(ctx, input)
				}
				var results []*svc.ListRecipeVersionsOutput
				p := svc.NewListRecipeVersionsPaginator(client, input)
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
		"list-recipes": {
			Name:   "list-recipes",
			Fields: fields_list_recipes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRecipesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_recipes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRecipes(ctx, input)
				}
				var results []*svc.ListRecipesOutput
				p := svc.NewListRecipesPaginator(client, input)
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
		"list-rulesets": {
			Name:   "list-rulesets",
			Fields: fields_list_rulesets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRulesetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_rulesets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRulesets(ctx, input)
				}
				var results []*svc.ListRulesetsOutput
				p := svc.NewListRulesetsPaginator(client, input)
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
		"list-schedules": {
			Name:   "list-schedules",
			Fields: fields_list_schedules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSchedulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_schedules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSchedules(ctx, input)
				}
				var results []*svc.ListSchedulesOutput
				p := svc.NewListSchedulesPaginator(client, input)
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
		"publish-recipe": {
			Name:   "publish-recipe",
			Fields: fields_publish_recipe,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PublishRecipeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_publish_recipe, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PublishRecipe(ctx, input)
			},
		},
		"send-project-session-action": {
			Name:   "send-project-session-action",
			Fields: fields_send_project_session_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendProjectSessionActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_project_session_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendProjectSessionAction(ctx, input)
			},
		},
		"start-job-run": {
			Name:   "start-job-run",
			Fields: fields_start_job_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartJobRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_job_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartJobRun(ctx, input)
			},
		},
		"start-project-session": {
			Name:   "start-project-session",
			Fields: fields_start_project_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartProjectSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_project_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartProjectSession(ctx, input)
			},
		},
		"stop-job-run": {
			Name:   "stop-job-run",
			Fields: fields_stop_job_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopJobRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_job_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopJobRun(ctx, input)
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
		"update-dataset": {
			Name:   "update-dataset",
			Fields: fields_update_dataset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDatasetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_dataset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDataset(ctx, input)
			},
		},
		"update-profile-job": {
			Name:   "update-profile-job",
			Fields: fields_update_profile_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProfileJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_profile_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProfileJob(ctx, input)
			},
		},
		"update-project": {
			Name:   "update-project",
			Fields: fields_update_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProject(ctx, input)
			},
		},
		"update-recipe": {
			Name:   "update-recipe",
			Fields: fields_update_recipe,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRecipeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_recipe, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRecipe(ctx, input)
			},
		},
		"update-recipe-job": {
			Name:   "update-recipe-job",
			Fields: fields_update_recipe_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRecipeJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_recipe_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRecipeJob(ctx, input)
			},
		},
		"update-ruleset": {
			Name:   "update-ruleset",
			Fields: fields_update_ruleset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRulesetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_ruleset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRuleset(ctx, input)
			},
		},
		"update-schedule": {
			Name:   "update-schedule",
			Fields: fields_update_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSchedule(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("databrew", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
