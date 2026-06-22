package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/bedrockdataautomation"
)

var fields_copy_blueprint_stage = []leanruntime.Field{
	{Name: "BlueprintArn", Flag: "blueprint-arn", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "SourceStage", Flag: "source-stage", Type: "types.BlueprintStage", Required: true},
	{Name: "TargetStage", Flag: "target-stage", Type: "types.BlueprintStage", Required: true},
}

var fields_create_blueprint = []leanruntime.Field{
	{Name: "BlueprintName", Flag: "blueprint-name", Type: "*string", Required: true},
	{Name: "BlueprintStage", Flag: "blueprint-stage", Type: "types.BlueprintStage", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "EncryptionConfiguration", Flag: "encryption-configuration", Type: "*types.EncryptionConfiguration", Required: false},
	{Name: "Schema", Flag: "schema", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Type", Flag: "type", Type: "types.Type", Required: true},
}

var fields_create_blueprint_version = []leanruntime.Field{
	{Name: "BlueprintArn", Flag: "blueprint-arn", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
}

var fields_create_data_automation_project = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CustomOutputConfiguration", Flag: "custom-output-configuration", Type: "*types.CustomOutputConfiguration", Required: false},
	{Name: "EncryptionConfiguration", Flag: "encryption-configuration", Type: "*types.EncryptionConfiguration", Required: false},
	{Name: "OverrideConfiguration", Flag: "override-configuration", Type: "*types.OverrideConfiguration", Required: false},
	{Name: "ProjectDescription", Flag: "project-description", Type: "*string", Required: false},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
	{Name: "ProjectStage", Flag: "project-stage", Type: "types.DataAutomationProjectStage", Required: false},
	{Name: "ProjectType", Flag: "project-type", Type: "types.DataAutomationProjectType", Required: false},
	{Name: "StandardOutputConfiguration", Flag: "standard-output-configuration", Type: "*types.StandardOutputConfiguration", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_blueprint = []leanruntime.Field{
	{Name: "BlueprintArn", Flag: "blueprint-arn", Type: "*string", Required: true},
	{Name: "BlueprintVersion", Flag: "blueprint-version", Type: "*string", Required: false},
}

var fields_delete_data_automation_project = []leanruntime.Field{
	{Name: "ProjectArn", Flag: "project-arn", Type: "*string", Required: true},
}

var fields_get_blueprint = []leanruntime.Field{
	{Name: "BlueprintArn", Flag: "blueprint-arn", Type: "*string", Required: true},
	{Name: "BlueprintStage", Flag: "blueprint-stage", Type: "types.BlueprintStage", Required: false},
	{Name: "BlueprintVersion", Flag: "blueprint-version", Type: "*string", Required: false},
}

var fields_get_blueprint_optimization_status = []leanruntime.Field{
	{Name: "InvocationArn", Flag: "invocation-arn", Type: "*string", Required: true},
}

var fields_get_data_automation_project = []leanruntime.Field{
	{Name: "ProjectArn", Flag: "project-arn", Type: "*string", Required: true},
	{Name: "ProjectStage", Flag: "project-stage", Type: "types.DataAutomationProjectStage", Required: false},
}

var fields_invoke_blueprint_optimization_async = []leanruntime.Field{
	{Name: "Blueprint", Flag: "blueprint", Type: "*types.BlueprintOptimizationObject", Required: true},
	{Name: "DataAutomationProfileArn", Flag: "data-automation-profile-arn", Type: "*string", Required: true},
	{Name: "EncryptionConfiguration", Flag: "encryption-configuration", Type: "*types.EncryptionConfiguration", Required: false},
	{Name: "OutputConfiguration", Flag: "output-configuration", Type: "*types.BlueprintOptimizationOutputConfiguration", Required: true},
	{Name: "Samples", Flag: "samples", Type: "[]types.BlueprintOptimizationSample", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_list_blueprints = []leanruntime.Field{
	{Name: "BlueprintArn", Flag: "blueprint-arn", Type: "*string", Required: false},
	{Name: "BlueprintStageFilter", Flag: "blueprint-stage-filter", Type: "types.BlueprintStageFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProjectFilter", Flag: "project-filter", Type: "*types.DataAutomationProjectFilter", Required: false},
	{Name: "ResourceOwner", Flag: "resource-owner", Type: "types.ResourceOwner", Required: false},
}

var fields_list_data_automation_projects = []leanruntime.Field{
	{Name: "BlueprintFilter", Flag: "blueprint-filter", Type: "*types.BlueprintFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProjectStageFilter", Flag: "project-stage-filter", Type: "types.DataAutomationProjectStageFilter", Required: false},
	{Name: "ResourceOwner", Flag: "resource-owner", Type: "types.ResourceOwner", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_blueprint = []leanruntime.Field{
	{Name: "BlueprintArn", Flag: "blueprint-arn", Type: "*string", Required: true},
	{Name: "BlueprintStage", Flag: "blueprint-stage", Type: "types.BlueprintStage", Required: false},
	{Name: "EncryptionConfiguration", Flag: "encryption-configuration", Type: "*types.EncryptionConfiguration", Required: false},
	{Name: "Schema", Flag: "schema", Type: "*string", Required: true},
}

var fields_update_data_automation_project = []leanruntime.Field{
	{Name: "CustomOutputConfiguration", Flag: "custom-output-configuration", Type: "*types.CustomOutputConfiguration", Required: false},
	{Name: "EncryptionConfiguration", Flag: "encryption-configuration", Type: "*types.EncryptionConfiguration", Required: false},
	{Name: "OverrideConfiguration", Flag: "override-configuration", Type: "*types.OverrideConfiguration", Required: false},
	{Name: "ProjectArn", Flag: "project-arn", Type: "*string", Required: true},
	{Name: "ProjectDescription", Flag: "project-description", Type: "*string", Required: false},
	{Name: "ProjectStage", Flag: "project-stage", Type: "types.DataAutomationProjectStage", Required: false},
	{Name: "StandardOutputConfiguration", Flag: "standard-output-configuration", Type: "*types.StandardOutputConfiguration", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"copy-blueprint-stage": {
			Name:   "copy-blueprint-stage",
			Fields: fields_copy_blueprint_stage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CopyBlueprintStageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_copy_blueprint_stage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CopyBlueprintStage(ctx, input)
			},
		},
		"create-blueprint": {
			Name:   "create-blueprint",
			Fields: fields_create_blueprint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBlueprintInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_blueprint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBlueprint(ctx, input)
			},
		},
		"create-blueprint-version": {
			Name:   "create-blueprint-version",
			Fields: fields_create_blueprint_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBlueprintVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_blueprint_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBlueprintVersion(ctx, input)
			},
		},
		"create-data-automation-project": {
			Name:   "create-data-automation-project",
			Fields: fields_create_data_automation_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataAutomationProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_automation_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataAutomationProject(ctx, input)
			},
		},
		"delete-blueprint": {
			Name:   "delete-blueprint",
			Fields: fields_delete_blueprint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBlueprintInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_blueprint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBlueprint(ctx, input)
			},
		},
		"delete-data-automation-project": {
			Name:   "delete-data-automation-project",
			Fields: fields_delete_data_automation_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDataAutomationProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_data_automation_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataAutomationProject(ctx, input)
			},
		},
		"get-blueprint": {
			Name:   "get-blueprint",
			Fields: fields_get_blueprint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBlueprintInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_blueprint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBlueprint(ctx, input)
			},
		},
		"get-blueprint-optimization-status": {
			Name:   "get-blueprint-optimization-status",
			Fields: fields_get_blueprint_optimization_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBlueprintOptimizationStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_blueprint_optimization_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBlueprintOptimizationStatus(ctx, input)
			},
		},
		"get-data-automation-project": {
			Name:   "get-data-automation-project",
			Fields: fields_get_data_automation_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataAutomationProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_automation_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataAutomationProject(ctx, input)
			},
		},
		"invoke-blueprint-optimization-async": {
			Name:   "invoke-blueprint-optimization-async",
			Fields: fields_invoke_blueprint_optimization_async,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InvokeBlueprintOptimizationAsyncInput{}
				if _, err := leanruntime.ApplyInput(input, fields_invoke_blueprint_optimization_async, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.InvokeBlueprintOptimizationAsync(ctx, input)
			},
		},
		"list-blueprints": {
			Name:   "list-blueprints",
			Fields: fields_list_blueprints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBlueprintsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_blueprints, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBlueprints(ctx, input)
				}
				var results []*svc.ListBlueprintsOutput
				p := svc.NewListBlueprintsPaginator(client, input)
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
		"list-data-automation-projects": {
			Name:   "list-data-automation-projects",
			Fields: fields_list_data_automation_projects,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataAutomationProjectsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_automation_projects, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataAutomationProjects(ctx, input)
				}
				var results []*svc.ListDataAutomationProjectsOutput
				p := svc.NewListDataAutomationProjectsPaginator(client, input)
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
		"update-blueprint": {
			Name:   "update-blueprint",
			Fields: fields_update_blueprint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBlueprintInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_blueprint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBlueprint(ctx, input)
			},
		},
		"update-data-automation-project": {
			Name:   "update-data-automation-project",
			Fields: fields_update_data_automation_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDataAutomationProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_data_automation_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDataAutomationProject(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("bedrockdataautomation", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
