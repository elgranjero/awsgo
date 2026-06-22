package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/amplifyuibuilder"
)

var fields_create_component = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ComponentToCreate", Flag: "component-to-create", Type: "*types.CreateComponentData", Required: true},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: true},
}

var fields_create_form = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: true},
	{Name: "FormToCreate", Flag: "form-to-create", Type: "*types.CreateFormData", Required: true},
}

var fields_create_theme = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: true},
	{Name: "ThemeToCreate", Flag: "theme-to-create", Type: "*types.CreateThemeData", Required: true},
}

var fields_delete_component = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_form = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_theme = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_exchange_code_for_token = []leanruntime.Field{
	{Name: "Provider", Flag: "provider", Type: "types.TokenProviders", Required: true},
	{Name: "Request", Flag: "request", Type: "*types.ExchangeCodeForTokenRequestBody", Required: true},
}

var fields_export_components = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_export_forms = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_export_themes = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_codegen_job = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_component = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_form = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_metadata = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: true},
}

var fields_get_theme = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_list_codegen_jobs = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_components = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_forms = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_themes = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_put_metadata_flag = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "Body", Flag: "body", Type: "*types.PutMetadataFlagBody", Required: true},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: true},
	{Name: "FeatureName", Flag: "feature-name", Type: "*string", Required: true},
}

var fields_refresh_token = []leanruntime.Field{
	{Name: "Provider", Flag: "provider", Type: "types.TokenProviders", Required: true},
	{Name: "RefreshTokenBody", Flag: "refresh-token-body", Type: "*types.RefreshTokenRequestBody", Required: true},
}

var fields_start_codegen_job = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CodegenJobToCreate", Flag: "codegen-job-to-create", Type: "*types.StartCodegenJobData", Required: true},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_component = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "UpdatedComponent", Flag: "updated-component", Type: "*types.UpdateComponentData", Required: true},
}

var fields_update_form = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "UpdatedForm", Flag: "updated-form", Type: "*types.UpdateFormData", Required: true},
}

var fields_update_theme = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "UpdatedTheme", Flag: "updated-theme", Type: "*types.UpdateThemeData", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-component": {
			Name:   "create-component",
			Fields: fields_create_component,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateComponentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_component, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateComponent(ctx, input)
			},
		},
		"create-form": {
			Name:   "create-form",
			Fields: fields_create_form,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFormInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_form, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateForm(ctx, input)
			},
		},
		"create-theme": {
			Name:   "create-theme",
			Fields: fields_create_theme,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateThemeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_theme, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTheme(ctx, input)
			},
		},
		"delete-component": {
			Name:   "delete-component",
			Fields: fields_delete_component,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteComponentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_component, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteComponent(ctx, input)
			},
		},
		"delete-form": {
			Name:   "delete-form",
			Fields: fields_delete_form,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFormInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_form, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteForm(ctx, input)
			},
		},
		"delete-theme": {
			Name:   "delete-theme",
			Fields: fields_delete_theme,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteThemeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_theme, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTheme(ctx, input)
			},
		},
		"exchange-code-for-token": {
			Name:   "exchange-code-for-token",
			Fields: fields_exchange_code_for_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExchangeCodeForTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_exchange_code_for_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExchangeCodeForToken(ctx, input)
			},
		},
		"export-components": {
			Name:   "export-components",
			Fields: fields_export_components,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExportComponentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_export_components, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ExportComponents(ctx, input)
				}
				var results []*svc.ExportComponentsOutput
				p := svc.NewExportComponentsPaginator(client, input)
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
		"export-forms": {
			Name:   "export-forms",
			Fields: fields_export_forms,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExportFormsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_export_forms, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ExportForms(ctx, input)
				}
				var results []*svc.ExportFormsOutput
				p := svc.NewExportFormsPaginator(client, input)
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
		"export-themes": {
			Name:   "export-themes",
			Fields: fields_export_themes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExportThemesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_export_themes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ExportThemes(ctx, input)
				}
				var results []*svc.ExportThemesOutput
				p := svc.NewExportThemesPaginator(client, input)
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
		"get-codegen-job": {
			Name:   "get-codegen-job",
			Fields: fields_get_codegen_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCodegenJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_codegen_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCodegenJob(ctx, input)
			},
		},
		"get-component": {
			Name:   "get-component",
			Fields: fields_get_component,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetComponentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_component, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetComponent(ctx, input)
			},
		},
		"get-form": {
			Name:   "get-form",
			Fields: fields_get_form,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFormInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_form, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetForm(ctx, input)
			},
		},
		"get-metadata": {
			Name:   "get-metadata",
			Fields: fields_get_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMetadata(ctx, input)
			},
		},
		"get-theme": {
			Name:   "get-theme",
			Fields: fields_get_theme,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetThemeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_theme, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTheme(ctx, input)
			},
		},
		"list-codegen-jobs": {
			Name:   "list-codegen-jobs",
			Fields: fields_list_codegen_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCodegenJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_codegen_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCodegenJobs(ctx, input)
				}
				var results []*svc.ListCodegenJobsOutput
				p := svc.NewListCodegenJobsPaginator(client, input)
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
		"list-components": {
			Name:   "list-components",
			Fields: fields_list_components,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListComponentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_components, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListComponents(ctx, input)
				}
				var results []*svc.ListComponentsOutput
				p := svc.NewListComponentsPaginator(client, input)
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
		"list-forms": {
			Name:   "list-forms",
			Fields: fields_list_forms,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFormsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_forms, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListForms(ctx, input)
				}
				var results []*svc.ListFormsOutput
				p := svc.NewListFormsPaginator(client, input)
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
		"list-themes": {
			Name:   "list-themes",
			Fields: fields_list_themes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListThemesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_themes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListThemes(ctx, input)
				}
				var results []*svc.ListThemesOutput
				p := svc.NewListThemesPaginator(client, input)
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
		"put-metadata-flag": {
			Name:   "put-metadata-flag",
			Fields: fields_put_metadata_flag,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutMetadataFlagInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_metadata_flag, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutMetadataFlag(ctx, input)
			},
		},
		"refresh-token": {
			Name:   "refresh-token",
			Fields: fields_refresh_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RefreshTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_refresh_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RefreshToken(ctx, input)
			},
		},
		"start-codegen-job": {
			Name:   "start-codegen-job",
			Fields: fields_start_codegen_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartCodegenJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_codegen_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartCodegenJob(ctx, input)
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
		"update-component": {
			Name:   "update-component",
			Fields: fields_update_component,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateComponentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_component, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateComponent(ctx, input)
			},
		},
		"update-form": {
			Name:   "update-form",
			Fields: fields_update_form,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFormInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_form, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateForm(ctx, input)
			},
		},
		"update-theme": {
			Name:   "update-theme",
			Fields: fields_update_theme,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateThemeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_theme, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTheme(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("amplifyuibuilder", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
