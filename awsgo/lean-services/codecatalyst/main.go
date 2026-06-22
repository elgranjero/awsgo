package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/codecatalyst"
)

var fields_create_access_token = []leanruntime.Field{
	{Name: "ExpiresTime", Flag: "expires-time", Type: "*time.Time", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_dev_environment = []leanruntime.Field{
	{Name: "Alias", Flag: "alias", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Ides", Flag: "ides", Type: "[]types.IdeConfiguration", Required: false},
	{Name: "InactivityTimeoutMinutes", Flag: "inactivity-timeout-minutes", Type: "int32", Required: false},
	{Name: "InstanceType", Flag: "instance-type", Type: "types.InstanceType", Required: true},
	{Name: "PersistentStorage", Flag: "persistent-storage", Type: "*types.PersistentStorageConfiguration", Required: true},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
	{Name: "Repositories", Flag: "repositories", Type: "[]types.RepositoryInput", Required: false},
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: true},
	{Name: "VpcConnectionName", Flag: "vpc-connection-name", Type: "*string", Required: false},
}

var fields_create_project = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: true},
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: true},
}

var fields_create_source_repository = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: true},
}

var fields_create_source_repository_branch = []leanruntime.Field{
	{Name: "HeadCommitId", Flag: "head-commit-id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
	{Name: "SourceRepositoryName", Flag: "source-repository-name", Type: "*string", Required: true},
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: true},
}

var fields_delete_access_token = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_dev_environment = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: true},
}

var fields_delete_project = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: true},
}

var fields_delete_source_repository = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: true},
}

var fields_delete_space = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_dev_environment = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: true},
}

var fields_get_project = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: true},
}

var fields_get_source_repository = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: true},
}

var fields_get_source_repository_clone_urls = []leanruntime.Field{
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
	{Name: "SourceRepositoryName", Flag: "source-repository-name", Type: "*string", Required: true},
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: true},
}

var fields_get_space = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_subscription = []leanruntime.Field{
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: true},
}

var fields_get_user_details = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: false},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: false},
}

var fields_get_workflow = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: true},
}

var fields_get_workflow_run = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: true},
}

var fields_list_access_tokens = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_dev_environment_sessions = []leanruntime.Field{
	{Name: "DevEnvironmentId", Flag: "dev-environment-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: true},
}

var fields_list_dev_environments = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: false},
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: true},
}

var fields_list_event_logs = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "EventName", Flag: "event-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_list_projects = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.ProjectListFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: true},
}

var fields_list_source_repositories = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: true},
}

var fields_list_source_repository_branches = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
	{Name: "SourceRepositoryName", Flag: "source-repository-name", Type: "*string", Required: true},
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: true},
}

var fields_list_spaces = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_workflow_runs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
	{Name: "SortBy", Flag: "sort-by", Type: "[]types.WorkflowRunSortCriteria", Required: false},
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: true},
	{Name: "WorkflowId", Flag: "workflow-id", Type: "*string", Required: false},
}

var fields_list_workflows = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
	{Name: "SortBy", Flag: "sort-by", Type: "[]types.WorkflowSortCriteria", Required: false},
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: true},
}

var fields_start_dev_environment = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Ides", Flag: "ides", Type: "[]types.IdeConfiguration", Required: false},
	{Name: "InactivityTimeoutMinutes", Flag: "inactivity-timeout-minutes", Type: "int32", Required: false},
	{Name: "InstanceType", Flag: "instance-type", Type: "types.InstanceType", Required: false},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: true},
}

var fields_start_dev_environment_session = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
	{Name: "SessionConfiguration", Flag: "session-configuration", Type: "*types.DevEnvironmentSessionConfiguration", Required: true},
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: true},
}

var fields_start_workflow_run = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: true},
	{Name: "WorkflowId", Flag: "workflow-id", Type: "*string", Required: true},
}

var fields_stop_dev_environment = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: true},
}

var fields_stop_dev_environment_session = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: true},
}

var fields_update_dev_environment = []leanruntime.Field{
	{Name: "Alias", Flag: "alias", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Ides", Flag: "ides", Type: "[]types.IdeConfiguration", Required: false},
	{Name: "InactivityTimeoutMinutes", Flag: "inactivity-timeout-minutes", Type: "int32", Required: false},
	{Name: "InstanceType", Flag: "instance-type", Type: "types.InstanceType", Required: false},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: true},
}

var fields_update_project = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SpaceName", Flag: "space-name", Type: "*string", Required: true},
}

var fields_update_space = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_verify_session = []leanruntime.Field{}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-access-token": {
			Name:   "create-access-token",
			Fields: fields_create_access_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAccessTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_access_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAccessToken(ctx, input)
			},
		},
		"create-dev-environment": {
			Name:   "create-dev-environment",
			Fields: fields_create_dev_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDevEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_dev_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDevEnvironment(ctx, input)
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
		"create-source-repository": {
			Name:   "create-source-repository",
			Fields: fields_create_source_repository,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSourceRepositoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_source_repository, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSourceRepository(ctx, input)
			},
		},
		"create-source-repository-branch": {
			Name:   "create-source-repository-branch",
			Fields: fields_create_source_repository_branch,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSourceRepositoryBranchInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_source_repository_branch, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSourceRepositoryBranch(ctx, input)
			},
		},
		"delete-access-token": {
			Name:   "delete-access-token",
			Fields: fields_delete_access_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccessTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_access_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccessToken(ctx, input)
			},
		},
		"delete-dev-environment": {
			Name:   "delete-dev-environment",
			Fields: fields_delete_dev_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDevEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_dev_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDevEnvironment(ctx, input)
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
		"delete-source-repository": {
			Name:   "delete-source-repository",
			Fields: fields_delete_source_repository,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSourceRepositoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_source_repository, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSourceRepository(ctx, input)
			},
		},
		"delete-space": {
			Name:   "delete-space",
			Fields: fields_delete_space,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSpaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_space, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSpace(ctx, input)
			},
		},
		"get-dev-environment": {
			Name:   "get-dev-environment",
			Fields: fields_get_dev_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDevEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_dev_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDevEnvironment(ctx, input)
			},
		},
		"get-project": {
			Name:   "get-project",
			Fields: fields_get_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetProject(ctx, input)
			},
		},
		"get-source-repository": {
			Name:   "get-source-repository",
			Fields: fields_get_source_repository,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSourceRepositoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_source_repository, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSourceRepository(ctx, input)
			},
		},
		"get-source-repository-clone-urls": {
			Name:   "get-source-repository-clone-urls",
			Fields: fields_get_source_repository_clone_urls,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSourceRepositoryCloneUrlsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_source_repository_clone_urls, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSourceRepositoryCloneUrls(ctx, input)
			},
		},
		"get-space": {
			Name:   "get-space",
			Fields: fields_get_space,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSpaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_space, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSpace(ctx, input)
			},
		},
		"get-subscription": {
			Name:   "get-subscription",
			Fields: fields_get_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSubscription(ctx, input)
			},
		},
		"get-user-details": {
			Name:   "get-user-details",
			Fields: fields_get_user_details,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUserDetailsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_user_details, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetUserDetails(ctx, input)
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
		"list-access-tokens": {
			Name:   "list-access-tokens",
			Fields: fields_list_access_tokens,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccessTokensInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_access_tokens, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccessTokens(ctx, input)
				}
				var results []*svc.ListAccessTokensOutput
				p := svc.NewListAccessTokensPaginator(client, input)
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
		"list-dev-environment-sessions": {
			Name:   "list-dev-environment-sessions",
			Fields: fields_list_dev_environment_sessions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDevEnvironmentSessionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_dev_environment_sessions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDevEnvironmentSessions(ctx, input)
				}
				var results []*svc.ListDevEnvironmentSessionsOutput
				p := svc.NewListDevEnvironmentSessionsPaginator(client, input)
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
		"list-dev-environments": {
			Name:   "list-dev-environments",
			Fields: fields_list_dev_environments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDevEnvironmentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_dev_environments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDevEnvironments(ctx, input)
				}
				var results []*svc.ListDevEnvironmentsOutput
				p := svc.NewListDevEnvironmentsPaginator(client, input)
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
		"list-event-logs": {
			Name:   "list-event-logs",
			Fields: fields_list_event_logs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEventLogsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_event_logs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEventLogs(ctx, input)
				}
				var results []*svc.ListEventLogsOutput
				p := svc.NewListEventLogsPaginator(client, input)
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
		"list-source-repositories": {
			Name:   "list-source-repositories",
			Fields: fields_list_source_repositories,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSourceRepositoriesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_source_repositories, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSourceRepositories(ctx, input)
				}
				var results []*svc.ListSourceRepositoriesOutput
				p := svc.NewListSourceRepositoriesPaginator(client, input)
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
		"list-source-repository-branches": {
			Name:   "list-source-repository-branches",
			Fields: fields_list_source_repository_branches,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSourceRepositoryBranchesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_source_repository_branches, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSourceRepositoryBranches(ctx, input)
				}
				var results []*svc.ListSourceRepositoryBranchesOutput
				p := svc.NewListSourceRepositoryBranchesPaginator(client, input)
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
		"list-spaces": {
			Name:   "list-spaces",
			Fields: fields_list_spaces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSpacesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_spaces, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSpaces(ctx, input)
				}
				var results []*svc.ListSpacesOutput
				p := svc.NewListSpacesPaginator(client, input)
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
		"start-dev-environment": {
			Name:   "start-dev-environment",
			Fields: fields_start_dev_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDevEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_dev_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDevEnvironment(ctx, input)
			},
		},
		"start-dev-environment-session": {
			Name:   "start-dev-environment-session",
			Fields: fields_start_dev_environment_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDevEnvironmentSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_dev_environment_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDevEnvironmentSession(ctx, input)
			},
		},
		"start-workflow-run": {
			Name:   "start-workflow-run",
			Fields: fields_start_workflow_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartWorkflowRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_workflow_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartWorkflowRun(ctx, input)
			},
		},
		"stop-dev-environment": {
			Name:   "stop-dev-environment",
			Fields: fields_stop_dev_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopDevEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_dev_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopDevEnvironment(ctx, input)
			},
		},
		"stop-dev-environment-session": {
			Name:   "stop-dev-environment-session",
			Fields: fields_stop_dev_environment_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopDevEnvironmentSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_dev_environment_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopDevEnvironmentSession(ctx, input)
			},
		},
		"update-dev-environment": {
			Name:   "update-dev-environment",
			Fields: fields_update_dev_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDevEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_dev_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDevEnvironment(ctx, input)
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
		"update-space": {
			Name:   "update-space",
			Fields: fields_update_space,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSpaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_space, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSpace(ctx, input)
			},
		},
		"verify-session": {
			Name:   "verify-session",
			Fields: fields_verify_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.VerifySessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_verify_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.VerifySession(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("codecatalyst", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
