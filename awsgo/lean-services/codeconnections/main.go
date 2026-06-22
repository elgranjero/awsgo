package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/codeconnections"
)

var fields_create_connection = []leanruntime.Field{
	{Name: "ConnectionName", Flag: "connection-name", Type: "*string", Required: true},
	{Name: "HostArn", Flag: "host-arn", Type: "*string", Required: false},
	{Name: "ProviderType", Flag: "provider-type", Type: "types.ProviderType", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_host = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ProviderEndpoint", Flag: "provider-endpoint", Type: "*string", Required: true},
	{Name: "ProviderType", Flag: "provider-type", Type: "types.ProviderType", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VpcConfiguration", Flag: "vpc-configuration", Type: "*types.VpcConfiguration", Required: false},
}

var fields_create_repository_link = []leanruntime.Field{
	{Name: "ConnectionArn", Flag: "connection-arn", Type: "*string", Required: true},
	{Name: "EncryptionKeyArn", Flag: "encryption-key-arn", Type: "*string", Required: false},
	{Name: "OwnerId", Flag: "owner-id", Type: "*string", Required: true},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_sync_configuration = []leanruntime.Field{
	{Name: "Branch", Flag: "branch", Type: "*string", Required: true},
	{Name: "ConfigFile", Flag: "config-file", Type: "*string", Required: true},
	{Name: "PublishDeploymentStatus", Flag: "publish-deployment-status", Type: "types.PublishDeploymentStatus", Required: false},
	{Name: "PullRequestComment", Flag: "pull-request-comment", Type: "types.PullRequestComment", Required: false},
	{Name: "RepositoryLinkId", Flag: "repository-link-id", Type: "*string", Required: true},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "SyncType", Flag: "sync-type", Type: "types.SyncConfigurationType", Required: true},
	{Name: "TriggerResourceUpdateOn", Flag: "trigger-resource-update-on", Type: "types.TriggerResourceUpdateOn", Required: false},
}

var fields_delete_connection = []leanruntime.Field{
	{Name: "ConnectionArn", Flag: "connection-arn", Type: "*string", Required: true},
}

var fields_delete_host = []leanruntime.Field{
	{Name: "HostArn", Flag: "host-arn", Type: "*string", Required: true},
}

var fields_delete_repository_link = []leanruntime.Field{
	{Name: "RepositoryLinkId", Flag: "repository-link-id", Type: "*string", Required: true},
}

var fields_delete_sync_configuration = []leanruntime.Field{
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
	{Name: "SyncType", Flag: "sync-type", Type: "types.SyncConfigurationType", Required: true},
}

var fields_get_connection = []leanruntime.Field{
	{Name: "ConnectionArn", Flag: "connection-arn", Type: "*string", Required: true},
}

var fields_get_host = []leanruntime.Field{
	{Name: "HostArn", Flag: "host-arn", Type: "*string", Required: true},
}

var fields_get_repository_link = []leanruntime.Field{
	{Name: "RepositoryLinkId", Flag: "repository-link-id", Type: "*string", Required: true},
}

var fields_get_repository_sync_status = []leanruntime.Field{
	{Name: "Branch", Flag: "branch", Type: "*string", Required: true},
	{Name: "RepositoryLinkId", Flag: "repository-link-id", Type: "*string", Required: true},
	{Name: "SyncType", Flag: "sync-type", Type: "types.SyncConfigurationType", Required: true},
}

var fields_get_resource_sync_status = []leanruntime.Field{
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
	{Name: "SyncType", Flag: "sync-type", Type: "types.SyncConfigurationType", Required: true},
}

var fields_get_sync_blocker_summary = []leanruntime.Field{
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
	{Name: "SyncType", Flag: "sync-type", Type: "types.SyncConfigurationType", Required: true},
}

var fields_get_sync_configuration = []leanruntime.Field{
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
	{Name: "SyncType", Flag: "sync-type", Type: "types.SyncConfigurationType", Required: true},
}

var fields_list_connections = []leanruntime.Field{
	{Name: "HostArnFilter", Flag: "host-arn-filter", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProviderTypeFilter", Flag: "provider-type-filter", Type: "types.ProviderType", Required: false},
}

var fields_list_hosts = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_repository_links = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_repository_sync_definitions = []leanruntime.Field{
	{Name: "RepositoryLinkId", Flag: "repository-link-id", Type: "*string", Required: true},
	{Name: "SyncType", Flag: "sync-type", Type: "types.SyncConfigurationType", Required: true},
}

var fields_list_sync_configurations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RepositoryLinkId", Flag: "repository-link-id", Type: "*string", Required: true},
	{Name: "SyncType", Flag: "sync-type", Type: "types.SyncConfigurationType", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_host = []leanruntime.Field{
	{Name: "HostArn", Flag: "host-arn", Type: "*string", Required: true},
	{Name: "ProviderEndpoint", Flag: "provider-endpoint", Type: "*string", Required: false},
	{Name: "VpcConfiguration", Flag: "vpc-configuration", Type: "*types.VpcConfiguration", Required: false},
}

var fields_update_repository_link = []leanruntime.Field{
	{Name: "ConnectionArn", Flag: "connection-arn", Type: "*string", Required: false},
	{Name: "EncryptionKeyArn", Flag: "encryption-key-arn", Type: "*string", Required: false},
	{Name: "RepositoryLinkId", Flag: "repository-link-id", Type: "*string", Required: true},
}

var fields_update_sync_blocker = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "ResolvedReason", Flag: "resolved-reason", Type: "*string", Required: true},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
	{Name: "SyncType", Flag: "sync-type", Type: "types.SyncConfigurationType", Required: true},
}

var fields_update_sync_configuration = []leanruntime.Field{
	{Name: "Branch", Flag: "branch", Type: "*string", Required: false},
	{Name: "ConfigFile", Flag: "config-file", Type: "*string", Required: false},
	{Name: "PublishDeploymentStatus", Flag: "publish-deployment-status", Type: "types.PublishDeploymentStatus", Required: false},
	{Name: "PullRequestComment", Flag: "pull-request-comment", Type: "types.PullRequestComment", Required: false},
	{Name: "RepositoryLinkId", Flag: "repository-link-id", Type: "*string", Required: false},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "SyncType", Flag: "sync-type", Type: "types.SyncConfigurationType", Required: true},
	{Name: "TriggerResourceUpdateOn", Flag: "trigger-resource-update-on", Type: "types.TriggerResourceUpdateOn", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-connection": {
			Name:   "create-connection",
			Fields: fields_create_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConnection(ctx, input)
			},
		},
		"create-host": {
			Name:   "create-host",
			Fields: fields_create_host,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateHostInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_host, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateHost(ctx, input)
			},
		},
		"create-repository-link": {
			Name:   "create-repository-link",
			Fields: fields_create_repository_link,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRepositoryLinkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_repository_link, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRepositoryLink(ctx, input)
			},
		},
		"create-sync-configuration": {
			Name:   "create-sync-configuration",
			Fields: fields_create_sync_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSyncConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_sync_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSyncConfiguration(ctx, input)
			},
		},
		"delete-connection": {
			Name:   "delete-connection",
			Fields: fields_delete_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConnection(ctx, input)
			},
		},
		"delete-host": {
			Name:   "delete-host",
			Fields: fields_delete_host,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteHostInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_host, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteHost(ctx, input)
			},
		},
		"delete-repository-link": {
			Name:   "delete-repository-link",
			Fields: fields_delete_repository_link,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRepositoryLinkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_repository_link, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRepositoryLink(ctx, input)
			},
		},
		"delete-sync-configuration": {
			Name:   "delete-sync-configuration",
			Fields: fields_delete_sync_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSyncConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_sync_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSyncConfiguration(ctx, input)
			},
		},
		"get-connection": {
			Name:   "get-connection",
			Fields: fields_get_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConnection(ctx, input)
			},
		},
		"get-host": {
			Name:   "get-host",
			Fields: fields_get_host,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetHostInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_host, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetHost(ctx, input)
			},
		},
		"get-repository-link": {
			Name:   "get-repository-link",
			Fields: fields_get_repository_link,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRepositoryLinkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_repository_link, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRepositoryLink(ctx, input)
			},
		},
		"get-repository-sync-status": {
			Name:   "get-repository-sync-status",
			Fields: fields_get_repository_sync_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRepositorySyncStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_repository_sync_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRepositorySyncStatus(ctx, input)
			},
		},
		"get-resource-sync-status": {
			Name:   "get-resource-sync-status",
			Fields: fields_get_resource_sync_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetResourceSyncStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_resource_sync_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetResourceSyncStatus(ctx, input)
			},
		},
		"get-sync-blocker-summary": {
			Name:   "get-sync-blocker-summary",
			Fields: fields_get_sync_blocker_summary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSyncBlockerSummaryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sync_blocker_summary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSyncBlockerSummary(ctx, input)
			},
		},
		"get-sync-configuration": {
			Name:   "get-sync-configuration",
			Fields: fields_get_sync_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSyncConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sync_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSyncConfiguration(ctx, input)
			},
		},
		"list-connections": {
			Name:   "list-connections",
			Fields: fields_list_connections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConnectionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_connections, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConnections(ctx, input)
				}
				var results []*svc.ListConnectionsOutput
				p := svc.NewListConnectionsPaginator(client, input)
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
		"list-hosts": {
			Name:   "list-hosts",
			Fields: fields_list_hosts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListHostsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_hosts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListHosts(ctx, input)
				}
				var results []*svc.ListHostsOutput
				p := svc.NewListHostsPaginator(client, input)
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
		"list-repository-links": {
			Name:   "list-repository-links",
			Fields: fields_list_repository_links,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRepositoryLinksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_repository_links, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRepositoryLinks(ctx, input)
				}
				var results []*svc.ListRepositoryLinksOutput
				p := svc.NewListRepositoryLinksPaginator(client, input)
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
		"list-repository-sync-definitions": {
			Name:   "list-repository-sync-definitions",
			Fields: fields_list_repository_sync_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRepositorySyncDefinitionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_repository_sync_definitions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListRepositorySyncDefinitions(ctx, input)
			},
		},
		"list-sync-configurations": {
			Name:   "list-sync-configurations",
			Fields: fields_list_sync_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSyncConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_sync_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSyncConfigurations(ctx, input)
				}
				var results []*svc.ListSyncConfigurationsOutput
				p := svc.NewListSyncConfigurationsPaginator(client, input)
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
		"update-host": {
			Name:   "update-host",
			Fields: fields_update_host,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateHostInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_host, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateHost(ctx, input)
			},
		},
		"update-repository-link": {
			Name:   "update-repository-link",
			Fields: fields_update_repository_link,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRepositoryLinkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_repository_link, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRepositoryLink(ctx, input)
			},
		},
		"update-sync-blocker": {
			Name:   "update-sync-blocker",
			Fields: fields_update_sync_blocker,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSyncBlockerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_sync_blocker, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSyncBlocker(ctx, input)
			},
		},
		"update-sync-configuration": {
			Name:   "update-sync-configuration",
			Fields: fields_update_sync_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSyncConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_sync_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSyncConfiguration(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("codeconnections", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
