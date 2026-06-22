package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository"
)

var fields_create_application = []leanruntime.Field{
	{Name: "Author", Flag: "author", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "HomePageUrl", Flag: "home-page-url", Type: "*string", Required: false},
	{Name: "Labels", Flag: "labels", Type: "[]string", Required: false},
	{Name: "LicenseBody", Flag: "license-body", Type: "*string", Required: false},
	{Name: "LicenseUrl", Flag: "license-url", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ReadmeBody", Flag: "readme-body", Type: "*string", Required: false},
	{Name: "ReadmeUrl", Flag: "readme-url", Type: "*string", Required: false},
	{Name: "SemanticVersion", Flag: "semantic-version", Type: "*string", Required: false},
	{Name: "SourceCodeArchiveUrl", Flag: "source-code-archive-url", Type: "*string", Required: false},
	{Name: "SourceCodeUrl", Flag: "source-code-url", Type: "*string", Required: false},
	{Name: "SpdxLicenseId", Flag: "spdx-license-id", Type: "*string", Required: false},
	{Name: "TemplateBody", Flag: "template-body", Type: "*string", Required: false},
	{Name: "TemplateUrl", Flag: "template-url", Type: "*string", Required: false},
}

var fields_create_application_version = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "SemanticVersion", Flag: "semantic-version", Type: "*string", Required: true},
	{Name: "SourceCodeArchiveUrl", Flag: "source-code-archive-url", Type: "*string", Required: false},
	{Name: "SourceCodeUrl", Flag: "source-code-url", Type: "*string", Required: false},
	{Name: "TemplateBody", Flag: "template-body", Type: "*string", Required: false},
	{Name: "TemplateUrl", Flag: "template-url", Type: "*string", Required: false},
}

var fields_create_cloud_formation_change_set = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "Capabilities", Flag: "capabilities", Type: "[]string", Required: false},
	{Name: "ChangeSetName", Flag: "change-set-name", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "NotificationArns", Flag: "notification-arns", Type: "[]string", Required: false},
	{Name: "ParameterOverrides", Flag: "parameter-overrides", Type: "[]types.ParameterValue", Required: false},
	{Name: "ResourceTypes", Flag: "resource-types", Type: "[]string", Required: false},
	{Name: "RollbackConfiguration", Flag: "rollback-configuration", Type: "*types.RollbackConfiguration", Required: false},
	{Name: "SemanticVersion", Flag: "semantic-version", Type: "*string", Required: false},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TemplateId", Flag: "template-id", Type: "*string", Required: false},
}

var fields_create_cloud_formation_template = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "SemanticVersion", Flag: "semantic-version", Type: "*string", Required: false},
}

var fields_delete_application = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_get_application = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "SemanticVersion", Flag: "semantic-version", Type: "*string", Required: false},
}

var fields_get_application_policy = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_get_cloud_formation_template = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "TemplateId", Flag: "template-id", Type: "*string", Required: true},
}

var fields_list_application_dependencies = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SemanticVersion", Flag: "semantic-version", Type: "*string", Required: false},
}

var fields_list_application_versions = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_applications = []leanruntime.Field{
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_put_application_policy = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "Statements", Flag: "statements", Type: "[]types.ApplicationPolicyStatement", Required: true},
}

var fields_unshare_application = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "OrganizationId", Flag: "organization-id", Type: "*string", Required: true},
}

var fields_update_application = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "Author", Flag: "author", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "HomePageUrl", Flag: "home-page-url", Type: "*string", Required: false},
	{Name: "Labels", Flag: "labels", Type: "[]string", Required: false},
	{Name: "ReadmeBody", Flag: "readme-body", Type: "*string", Required: false},
	{Name: "ReadmeUrl", Flag: "readme-url", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-application": {
			Name:   "create-application",
			Fields: fields_create_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateApplication(ctx, input)
			},
		},
		"create-application-version": {
			Name:   "create-application-version",
			Fields: fields_create_application_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateApplicationVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_application_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateApplicationVersion(ctx, input)
			},
		},
		"create-cloud-formation-change-set": {
			Name:   "create-cloud-formation-change-set",
			Fields: fields_create_cloud_formation_change_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCloudFormationChangeSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_cloud_formation_change_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCloudFormationChangeSet(ctx, input)
			},
		},
		"create-cloud-formation-template": {
			Name:   "create-cloud-formation-template",
			Fields: fields_create_cloud_formation_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCloudFormationTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_cloud_formation_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCloudFormationTemplate(ctx, input)
			},
		},
		"delete-application": {
			Name:   "delete-application",
			Fields: fields_delete_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApplication(ctx, input)
			},
		},
		"get-application": {
			Name:   "get-application",
			Fields: fields_get_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApplication(ctx, input)
			},
		},
		"get-application-policy": {
			Name:   "get-application-policy",
			Fields: fields_get_application_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetApplicationPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_application_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApplicationPolicy(ctx, input)
			},
		},
		"get-cloud-formation-template": {
			Name:   "get-cloud-formation-template",
			Fields: fields_get_cloud_formation_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCloudFormationTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_cloud_formation_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCloudFormationTemplate(ctx, input)
			},
		},
		"list-application-dependencies": {
			Name:   "list-application-dependencies",
			Fields: fields_list_application_dependencies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListApplicationDependenciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_application_dependencies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApplicationDependencies(ctx, input)
				}
				var results []*svc.ListApplicationDependenciesOutput
				p := svc.NewListApplicationDependenciesPaginator(client, input)
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
		"list-application-versions": {
			Name:   "list-application-versions",
			Fields: fields_list_application_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListApplicationVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_application_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApplicationVersions(ctx, input)
				}
				var results []*svc.ListApplicationVersionsOutput
				p := svc.NewListApplicationVersionsPaginator(client, input)
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
		"list-applications": {
			Name:   "list-applications",
			Fields: fields_list_applications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListApplicationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_applications, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApplications(ctx, input)
				}
				var results []*svc.ListApplicationsOutput
				p := svc.NewListApplicationsPaginator(client, input)
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
		"put-application-policy": {
			Name:   "put-application-policy",
			Fields: fields_put_application_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutApplicationPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_application_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutApplicationPolicy(ctx, input)
			},
		},
		"unshare-application": {
			Name:   "unshare-application",
			Fields: fields_unshare_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UnshareApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_unshare_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UnshareApplication(ctx, input)
			},
		},
		"update-application": {
			Name:   "update-application",
			Fields: fields_update_application,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateApplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_application, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateApplication(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("serverlessapplicationrepository", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
