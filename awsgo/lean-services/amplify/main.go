package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/amplify"
)

var fields_create_app = []leanruntime.Field{
	{Name: "AccessToken", Flag: "access-token", Type: "*string", Required: false},
	{Name: "AutoBranchCreationConfig", Flag: "auto-branch-creation-config", Type: "*types.AutoBranchCreationConfig", Required: false},
	{Name: "AutoBranchCreationPatterns", Flag: "auto-branch-creation-patterns", Type: "[]string", Required: false},
	{Name: "BasicAuthCredentials", Flag: "basic-auth-credentials", Type: "*string", Required: false},
	{Name: "BuildSpec", Flag: "build-spec", Type: "*string", Required: false},
	{Name: "CacheConfig", Flag: "cache-config", Type: "*types.CacheConfig", Required: false},
	{Name: "ComputeRoleArn", Flag: "compute-role-arn", Type: "*string", Required: false},
	{Name: "CustomHeaders", Flag: "custom-headers", Type: "*string", Required: false},
	{Name: "CustomRules", Flag: "custom-rules", Type: "[]types.CustomRule", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EnableAutoBranchCreation", Flag: "enable-auto-branch-creation", Type: "*bool", Required: false},
	{Name: "EnableBasicAuth", Flag: "enable-basic-auth", Type: "*bool", Required: false},
	{Name: "EnableBranchAutoBuild", Flag: "enable-branch-auto-build", Type: "*bool", Required: false},
	{Name: "EnableBranchAutoDeletion", Flag: "enable-branch-auto-deletion", Type: "*bool", Required: false},
	{Name: "EnvironmentVariables", Flag: "environment-variables", Type: "map[string]string", Required: false},
	{Name: "IamServiceRoleArn", Flag: "iam-service-role-arn", Type: "*string", Required: false},
	{Name: "JobConfig", Flag: "job-config", Type: "*types.JobConfig", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OauthToken", Flag: "oauth-token", Type: "*string", Required: false},
	{Name: "Platform", Flag: "platform", Type: "types.Platform", Required: false},
	{Name: "Repository", Flag: "repository", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_backend_environment = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "DeploymentArtifacts", Flag: "deployment-artifacts", Type: "*string", Required: false},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: true},
	{Name: "StackName", Flag: "stack-name", Type: "*string", Required: false},
}

var fields_create_branch = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "Backend", Flag: "backend", Type: "*types.Backend", Required: false},
	{Name: "BackendEnvironmentArn", Flag: "backend-environment-arn", Type: "*string", Required: false},
	{Name: "BasicAuthCredentials", Flag: "basic-auth-credentials", Type: "*string", Required: false},
	{Name: "BranchName", Flag: "branch-name", Type: "*string", Required: true},
	{Name: "BuildSpec", Flag: "build-spec", Type: "*string", Required: false},
	{Name: "ComputeRoleArn", Flag: "compute-role-arn", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "EnableAutoBuild", Flag: "enable-auto-build", Type: "*bool", Required: false},
	{Name: "EnableBasicAuth", Flag: "enable-basic-auth", Type: "*bool", Required: false},
	{Name: "EnableNotification", Flag: "enable-notification", Type: "*bool", Required: false},
	{Name: "EnablePerformanceMode", Flag: "enable-performance-mode", Type: "*bool", Required: false},
	{Name: "EnablePullRequestPreview", Flag: "enable-pull-request-preview", Type: "*bool", Required: false},
	{Name: "EnableSkewProtection", Flag: "enable-skew-protection", Type: "*bool", Required: false},
	{Name: "EnvironmentVariables", Flag: "environment-variables", Type: "map[string]string", Required: false},
	{Name: "Framework", Flag: "framework", Type: "*string", Required: false},
	{Name: "PullRequestEnvironmentName", Flag: "pull-request-environment-name", Type: "*string", Required: false},
	{Name: "Stage", Flag: "stage", Type: "types.Stage", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Ttl", Flag: "ttl", Type: "*string", Required: false},
}

var fields_create_deployment = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "BranchName", Flag: "branch-name", Type: "*string", Required: true},
	{Name: "FileMap", Flag: "file-map", Type: "map[string]string", Required: false},
}

var fields_create_domain_association = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "AutoSubDomainCreationPatterns", Flag: "auto-sub-domain-creation-patterns", Type: "[]string", Required: false},
	{Name: "AutoSubDomainIAMRole", Flag: "auto-sub-domain-iam-role", Type: "*string", Required: false},
	{Name: "CertificateSettings", Flag: "certificate-settings", Type: "*types.CertificateSettings", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "EnableAutoSubDomain", Flag: "enable-auto-sub-domain", Type: "*bool", Required: false},
	{Name: "SubDomainSettings", Flag: "sub-domain-settings", Type: "[]types.SubDomainSetting", Required: true},
}

var fields_create_webhook = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "BranchName", Flag: "branch-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
}

var fields_delete_app = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
}

var fields_delete_backend_environment = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: true},
}

var fields_delete_branch = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "BranchName", Flag: "branch-name", Type: "*string", Required: true},
}

var fields_delete_domain_association = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_delete_job = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "BranchName", Flag: "branch-name", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_delete_webhook = []leanruntime.Field{
	{Name: "WebhookId", Flag: "webhook-id", Type: "*string", Required: true},
}

var fields_generate_access_logs = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
}

var fields_get_app = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
}

var fields_get_artifact_url = []leanruntime.Field{
	{Name: "ArtifactId", Flag: "artifact-id", Type: "*string", Required: true},
}

var fields_get_backend_environment = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: true},
}

var fields_get_branch = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "BranchName", Flag: "branch-name", Type: "*string", Required: true},
}

var fields_get_domain_association = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_get_job = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "BranchName", Flag: "branch-name", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_get_webhook = []leanruntime.Field{
	{Name: "WebhookId", Flag: "webhook-id", Type: "*string", Required: true},
}

var fields_list_apps = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_artifacts = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "BranchName", Flag: "branch-name", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_backend_environments = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "EnvironmentName", Flag: "environment-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_branches = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_domain_associations = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_jobs = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "BranchName", Flag: "branch-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_webhooks = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_start_deployment = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "BranchName", Flag: "branch-name", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: false},
	{Name: "SourceUrl", Flag: "source-url", Type: "*string", Required: false},
	{Name: "SourceUrlType", Flag: "source-url-type", Type: "types.SourceUrlType", Required: false},
}

var fields_start_job = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "BranchName", Flag: "branch-name", Type: "*string", Required: true},
	{Name: "CommitId", Flag: "commit-id", Type: "*string", Required: false},
	{Name: "CommitMessage", Flag: "commit-message", Type: "*string", Required: false},
	{Name: "CommitTime", Flag: "commit-time", Type: "*time.Time", Required: false},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: false},
	{Name: "JobReason", Flag: "job-reason", Type: "*string", Required: false},
	{Name: "JobType", Flag: "job-type", Type: "types.JobType", Required: true},
}

var fields_stop_job = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "BranchName", Flag: "branch-name", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_app = []leanruntime.Field{
	{Name: "AccessToken", Flag: "access-token", Type: "*string", Required: false},
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "AutoBranchCreationConfig", Flag: "auto-branch-creation-config", Type: "*types.AutoBranchCreationConfig", Required: false},
	{Name: "AutoBranchCreationPatterns", Flag: "auto-branch-creation-patterns", Type: "[]string", Required: false},
	{Name: "BasicAuthCredentials", Flag: "basic-auth-credentials", Type: "*string", Required: false},
	{Name: "BuildSpec", Flag: "build-spec", Type: "*string", Required: false},
	{Name: "CacheConfig", Flag: "cache-config", Type: "*types.CacheConfig", Required: false},
	{Name: "ComputeRoleArn", Flag: "compute-role-arn", Type: "*string", Required: false},
	{Name: "CustomHeaders", Flag: "custom-headers", Type: "*string", Required: false},
	{Name: "CustomRules", Flag: "custom-rules", Type: "[]types.CustomRule", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EnableAutoBranchCreation", Flag: "enable-auto-branch-creation", Type: "*bool", Required: false},
	{Name: "EnableBasicAuth", Flag: "enable-basic-auth", Type: "*bool", Required: false},
	{Name: "EnableBranchAutoBuild", Flag: "enable-branch-auto-build", Type: "*bool", Required: false},
	{Name: "EnableBranchAutoDeletion", Flag: "enable-branch-auto-deletion", Type: "*bool", Required: false},
	{Name: "EnvironmentVariables", Flag: "environment-variables", Type: "map[string]string", Required: false},
	{Name: "IamServiceRoleArn", Flag: "iam-service-role-arn", Type: "*string", Required: false},
	{Name: "JobConfig", Flag: "job-config", Type: "*types.JobConfig", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "OauthToken", Flag: "oauth-token", Type: "*string", Required: false},
	{Name: "Platform", Flag: "platform", Type: "types.Platform", Required: false},
	{Name: "Repository", Flag: "repository", Type: "*string", Required: false},
}

var fields_update_branch = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "Backend", Flag: "backend", Type: "*types.Backend", Required: false},
	{Name: "BackendEnvironmentArn", Flag: "backend-environment-arn", Type: "*string", Required: false},
	{Name: "BasicAuthCredentials", Flag: "basic-auth-credentials", Type: "*string", Required: false},
	{Name: "BranchName", Flag: "branch-name", Type: "*string", Required: true},
	{Name: "BuildSpec", Flag: "build-spec", Type: "*string", Required: false},
	{Name: "ComputeRoleArn", Flag: "compute-role-arn", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "EnableAutoBuild", Flag: "enable-auto-build", Type: "*bool", Required: false},
	{Name: "EnableBasicAuth", Flag: "enable-basic-auth", Type: "*bool", Required: false},
	{Name: "EnableNotification", Flag: "enable-notification", Type: "*bool", Required: false},
	{Name: "EnablePerformanceMode", Flag: "enable-performance-mode", Type: "*bool", Required: false},
	{Name: "EnablePullRequestPreview", Flag: "enable-pull-request-preview", Type: "*bool", Required: false},
	{Name: "EnableSkewProtection", Flag: "enable-skew-protection", Type: "*bool", Required: false},
	{Name: "EnvironmentVariables", Flag: "environment-variables", Type: "map[string]string", Required: false},
	{Name: "Framework", Flag: "framework", Type: "*string", Required: false},
	{Name: "PullRequestEnvironmentName", Flag: "pull-request-environment-name", Type: "*string", Required: false},
	{Name: "Stage", Flag: "stage", Type: "types.Stage", Required: false},
	{Name: "Ttl", Flag: "ttl", Type: "*string", Required: false},
}

var fields_update_domain_association = []leanruntime.Field{
	{Name: "AppId", Flag: "app-id", Type: "*string", Required: true},
	{Name: "AutoSubDomainCreationPatterns", Flag: "auto-sub-domain-creation-patterns", Type: "[]string", Required: false},
	{Name: "AutoSubDomainIAMRole", Flag: "auto-sub-domain-iam-role", Type: "*string", Required: false},
	{Name: "CertificateSettings", Flag: "certificate-settings", Type: "*types.CertificateSettings", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "EnableAutoSubDomain", Flag: "enable-auto-sub-domain", Type: "*bool", Required: false},
	{Name: "SubDomainSettings", Flag: "sub-domain-settings", Type: "[]types.SubDomainSetting", Required: false},
}

var fields_update_webhook = []leanruntime.Field{
	{Name: "BranchName", Flag: "branch-name", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "WebhookId", Flag: "webhook-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-app": {
			Name:   "create-app",
			Fields: fields_create_app,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAppInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_app, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateApp(ctx, input)
			},
		},
		"create-backend-environment": {
			Name:   "create-backend-environment",
			Fields: fields_create_backend_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBackendEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_backend_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBackendEnvironment(ctx, input)
			},
		},
		"create-branch": {
			Name:   "create-branch",
			Fields: fields_create_branch,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBranchInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_branch, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBranch(ctx, input)
			},
		},
		"create-deployment": {
			Name:   "create-deployment",
			Fields: fields_create_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDeployment(ctx, input)
			},
		},
		"create-domain-association": {
			Name:   "create-domain-association",
			Fields: fields_create_domain_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDomainAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_domain_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDomainAssociation(ctx, input)
			},
		},
		"create-webhook": {
			Name:   "create-webhook",
			Fields: fields_create_webhook,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWebhookInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_webhook, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWebhook(ctx, input)
			},
		},
		"delete-app": {
			Name:   "delete-app",
			Fields: fields_delete_app,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAppInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_app, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApp(ctx, input)
			},
		},
		"delete-backend-environment": {
			Name:   "delete-backend-environment",
			Fields: fields_delete_backend_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBackendEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_backend_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBackendEnvironment(ctx, input)
			},
		},
		"delete-branch": {
			Name:   "delete-branch",
			Fields: fields_delete_branch,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBranchInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_branch, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBranch(ctx, input)
			},
		},
		"delete-domain-association": {
			Name:   "delete-domain-association",
			Fields: fields_delete_domain_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDomainAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_domain_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDomainAssociation(ctx, input)
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
		"generate-access-logs": {
			Name:   "generate-access-logs",
			Fields: fields_generate_access_logs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GenerateAccessLogsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_generate_access_logs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GenerateAccessLogs(ctx, input)
			},
		},
		"get-app": {
			Name:   "get-app",
			Fields: fields_get_app,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAppInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_app, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApp(ctx, input)
			},
		},
		"get-artifact-url": {
			Name:   "get-artifact-url",
			Fields: fields_get_artifact_url,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetArtifactUrlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_artifact_url, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetArtifactUrl(ctx, input)
			},
		},
		"get-backend-environment": {
			Name:   "get-backend-environment",
			Fields: fields_get_backend_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBackendEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_backend_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBackendEnvironment(ctx, input)
			},
		},
		"get-branch": {
			Name:   "get-branch",
			Fields: fields_get_branch,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBranchInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_branch, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBranch(ctx, input)
			},
		},
		"get-domain-association": {
			Name:   "get-domain-association",
			Fields: fields_get_domain_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDomainAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_domain_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDomainAssociation(ctx, input)
			},
		},
		"get-job": {
			Name:   "get-job",
			Fields: fields_get_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetJob(ctx, input)
			},
		},
		"get-webhook": {
			Name:   "get-webhook",
			Fields: fields_get_webhook,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWebhookInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_webhook, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWebhook(ctx, input)
			},
		},
		"list-apps": {
			Name:   "list-apps",
			Fields: fields_list_apps,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAppsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_apps, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApps(ctx, input)
				}
				var results []*svc.ListAppsOutput
				p := svc.NewListAppsPaginator(client, input)
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
		"list-artifacts": {
			Name:   "list-artifacts",
			Fields: fields_list_artifacts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListArtifactsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_artifacts, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListArtifacts(ctx, input)
			},
		},
		"list-backend-environments": {
			Name:   "list-backend-environments",
			Fields: fields_list_backend_environments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBackendEnvironmentsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_backend_environments, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListBackendEnvironments(ctx, input)
			},
		},
		"list-branches": {
			Name:   "list-branches",
			Fields: fields_list_branches,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBranchesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_branches, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBranches(ctx, input)
				}
				var results []*svc.ListBranchesOutput
				p := svc.NewListBranchesPaginator(client, input)
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
		"list-domain-associations": {
			Name:   "list-domain-associations",
			Fields: fields_list_domain_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDomainAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_domain_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDomainAssociations(ctx, input)
				}
				var results []*svc.ListDomainAssociationsOutput
				p := svc.NewListDomainAssociationsPaginator(client, input)
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
		"list-webhooks": {
			Name:   "list-webhooks",
			Fields: fields_list_webhooks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWebhooksInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_webhooks, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListWebhooks(ctx, input)
			},
		},
		"start-deployment": {
			Name:   "start-deployment",
			Fields: fields_start_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDeployment(ctx, input)
			},
		},
		"start-job": {
			Name:   "start-job",
			Fields: fields_start_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartJob(ctx, input)
			},
		},
		"stop-job": {
			Name:   "stop-job",
			Fields: fields_stop_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopJob(ctx, input)
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
		"update-app": {
			Name:   "update-app",
			Fields: fields_update_app,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAppInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_app, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateApp(ctx, input)
			},
		},
		"update-branch": {
			Name:   "update-branch",
			Fields: fields_update_branch,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBranchInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_branch, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBranch(ctx, input)
			},
		},
		"update-domain-association": {
			Name:   "update-domain-association",
			Fields: fields_update_domain_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDomainAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_domain_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDomainAssociation(ctx, input)
			},
		},
		"update-webhook": {
			Name:   "update-webhook",
			Fields: fields_update_webhook,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWebhookInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_webhook, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWebhook(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("amplify", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
