package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/migrationhub"
)

var fields_associate_created_artifact = []leanruntime.Field{
	{Name: "CreatedArtifact", Flag: "created-artifact", Type: "*types.CreatedArtifact", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "bool", Required: false},
	{Name: "MigrationTaskName", Flag: "migration-task-name", Type: "*string", Required: true},
	{Name: "ProgressUpdateStream", Flag: "progress-update-stream", Type: "*string", Required: true},
}

var fields_associate_discovered_resource = []leanruntime.Field{
	{Name: "DiscoveredResource", Flag: "discovered-resource", Type: "*types.DiscoveredResource", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "bool", Required: false},
	{Name: "MigrationTaskName", Flag: "migration-task-name", Type: "*string", Required: true},
	{Name: "ProgressUpdateStream", Flag: "progress-update-stream", Type: "*string", Required: true},
}

var fields_associate_source_resource = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "bool", Required: false},
	{Name: "MigrationTaskName", Flag: "migration-task-name", Type: "*string", Required: true},
	{Name: "ProgressUpdateStream", Flag: "progress-update-stream", Type: "*string", Required: true},
	{Name: "SourceResource", Flag: "source-resource", Type: "*types.SourceResource", Required: true},
}

var fields_create_progress_update_stream = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "bool", Required: false},
	{Name: "ProgressUpdateStreamName", Flag: "progress-update-stream-name", Type: "*string", Required: true},
}

var fields_delete_progress_update_stream = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "bool", Required: false},
	{Name: "ProgressUpdateStreamName", Flag: "progress-update-stream-name", Type: "*string", Required: true},
}

var fields_describe_application_state = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
}

var fields_describe_migration_task = []leanruntime.Field{
	{Name: "MigrationTaskName", Flag: "migration-task-name", Type: "*string", Required: true},
	{Name: "ProgressUpdateStream", Flag: "progress-update-stream", Type: "*string", Required: true},
}

var fields_disassociate_created_artifact = []leanruntime.Field{
	{Name: "CreatedArtifactName", Flag: "created-artifact-name", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "bool", Required: false},
	{Name: "MigrationTaskName", Flag: "migration-task-name", Type: "*string", Required: true},
	{Name: "ProgressUpdateStream", Flag: "progress-update-stream", Type: "*string", Required: true},
}

var fields_disassociate_discovered_resource = []leanruntime.Field{
	{Name: "ConfigurationId", Flag: "configuration-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "bool", Required: false},
	{Name: "MigrationTaskName", Flag: "migration-task-name", Type: "*string", Required: true},
	{Name: "ProgressUpdateStream", Flag: "progress-update-stream", Type: "*string", Required: true},
}

var fields_disassociate_source_resource = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "bool", Required: false},
	{Name: "MigrationTaskName", Flag: "migration-task-name", Type: "*string", Required: true},
	{Name: "ProgressUpdateStream", Flag: "progress-update-stream", Type: "*string", Required: true},
	{Name: "SourceResourceName", Flag: "source-resource-name", Type: "*string", Required: true},
}

var fields_import_migration_task = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "bool", Required: false},
	{Name: "MigrationTaskName", Flag: "migration-task-name", Type: "*string", Required: true},
	{Name: "ProgressUpdateStream", Flag: "progress-update-stream", Type: "*string", Required: true},
}

var fields_list_application_states = []leanruntime.Field{
	{Name: "ApplicationIds", Flag: "application-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_created_artifacts = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MigrationTaskName", Flag: "migration-task-name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProgressUpdateStream", Flag: "progress-update-stream", Type: "*string", Required: true},
}

var fields_list_discovered_resources = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MigrationTaskName", Flag: "migration-task-name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProgressUpdateStream", Flag: "progress-update-stream", Type: "*string", Required: true},
}

var fields_list_migration_task_updates = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MigrationTaskName", Flag: "migration-task-name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProgressUpdateStream", Flag: "progress-update-stream", Type: "*string", Required: true},
}

var fields_list_migration_tasks = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: false},
}

var fields_list_progress_update_streams = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_source_resources = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MigrationTaskName", Flag: "migration-task-name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProgressUpdateStream", Flag: "progress-update-stream", Type: "*string", Required: true},
}

var fields_notify_application_state = []leanruntime.Field{
	{Name: "ApplicationId", Flag: "application-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "bool", Required: false},
	{Name: "Status", Flag: "status", Type: "types.ApplicationStatus", Required: true},
	{Name: "UpdateDateTime", Flag: "update-date-time", Type: "*time.Time", Required: false},
}

var fields_notify_migration_task_state = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "bool", Required: false},
	{Name: "MigrationTaskName", Flag: "migration-task-name", Type: "*string", Required: true},
	{Name: "NextUpdateSeconds", Flag: "next-update-seconds", Type: "int32", Required: true},
	{Name: "ProgressUpdateStream", Flag: "progress-update-stream", Type: "*string", Required: true},
	{Name: "Task", Flag: "task", Type: "*types.Task", Required: true},
	{Name: "UpdateDateTime", Flag: "update-date-time", Type: "*time.Time", Required: true},
}

var fields_put_resource_attributes = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "bool", Required: false},
	{Name: "MigrationTaskName", Flag: "migration-task-name", Type: "*string", Required: true},
	{Name: "ProgressUpdateStream", Flag: "progress-update-stream", Type: "*string", Required: true},
	{Name: "ResourceAttributeList", Flag: "resource-attribute-list", Type: "[]types.ResourceAttribute", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-created-artifact": {
			Name:   "associate-created-artifact",
			Fields: fields_associate_created_artifact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateCreatedArtifactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_created_artifact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateCreatedArtifact(ctx, input)
			},
		},
		"associate-discovered-resource": {
			Name:   "associate-discovered-resource",
			Fields: fields_associate_discovered_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateDiscoveredResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_discovered_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateDiscoveredResource(ctx, input)
			},
		},
		"associate-source-resource": {
			Name:   "associate-source-resource",
			Fields: fields_associate_source_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateSourceResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_source_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateSourceResource(ctx, input)
			},
		},
		"create-progress-update-stream": {
			Name:   "create-progress-update-stream",
			Fields: fields_create_progress_update_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProgressUpdateStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_progress_update_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProgressUpdateStream(ctx, input)
			},
		},
		"delete-progress-update-stream": {
			Name:   "delete-progress-update-stream",
			Fields: fields_delete_progress_update_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProgressUpdateStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_progress_update_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProgressUpdateStream(ctx, input)
			},
		},
		"describe-application-state": {
			Name:   "describe-application-state",
			Fields: fields_describe_application_state,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeApplicationStateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_application_state, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeApplicationState(ctx, input)
			},
		},
		"describe-migration-task": {
			Name:   "describe-migration-task",
			Fields: fields_describe_migration_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMigrationTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_migration_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeMigrationTask(ctx, input)
			},
		},
		"disassociate-created-artifact": {
			Name:   "disassociate-created-artifact",
			Fields: fields_disassociate_created_artifact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateCreatedArtifactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_created_artifact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateCreatedArtifact(ctx, input)
			},
		},
		"disassociate-discovered-resource": {
			Name:   "disassociate-discovered-resource",
			Fields: fields_disassociate_discovered_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateDiscoveredResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_discovered_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateDiscoveredResource(ctx, input)
			},
		},
		"disassociate-source-resource": {
			Name:   "disassociate-source-resource",
			Fields: fields_disassociate_source_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateSourceResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_source_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateSourceResource(ctx, input)
			},
		},
		"import-migration-task": {
			Name:   "import-migration-task",
			Fields: fields_import_migration_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportMigrationTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_migration_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportMigrationTask(ctx, input)
			},
		},
		"list-application-states": {
			Name:   "list-application-states",
			Fields: fields_list_application_states,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListApplicationStatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_application_states, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApplicationStates(ctx, input)
				}
				var results []*svc.ListApplicationStatesOutput
				p := svc.NewListApplicationStatesPaginator(client, input)
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
		"list-created-artifacts": {
			Name:   "list-created-artifacts",
			Fields: fields_list_created_artifacts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCreatedArtifactsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_created_artifacts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCreatedArtifacts(ctx, input)
				}
				var results []*svc.ListCreatedArtifactsOutput
				p := svc.NewListCreatedArtifactsPaginator(client, input)
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
		"list-discovered-resources": {
			Name:   "list-discovered-resources",
			Fields: fields_list_discovered_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDiscoveredResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_discovered_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDiscoveredResources(ctx, input)
				}
				var results []*svc.ListDiscoveredResourcesOutput
				p := svc.NewListDiscoveredResourcesPaginator(client, input)
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
		"list-migration-task-updates": {
			Name:   "list-migration-task-updates",
			Fields: fields_list_migration_task_updates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMigrationTaskUpdatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_migration_task_updates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMigrationTaskUpdates(ctx, input)
				}
				var results []*svc.ListMigrationTaskUpdatesOutput
				p := svc.NewListMigrationTaskUpdatesPaginator(client, input)
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
		"list-migration-tasks": {
			Name:   "list-migration-tasks",
			Fields: fields_list_migration_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMigrationTasksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_migration_tasks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMigrationTasks(ctx, input)
				}
				var results []*svc.ListMigrationTasksOutput
				p := svc.NewListMigrationTasksPaginator(client, input)
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
		"list-progress-update-streams": {
			Name:   "list-progress-update-streams",
			Fields: fields_list_progress_update_streams,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProgressUpdateStreamsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_progress_update_streams, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProgressUpdateStreams(ctx, input)
				}
				var results []*svc.ListProgressUpdateStreamsOutput
				p := svc.NewListProgressUpdateStreamsPaginator(client, input)
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
		"list-source-resources": {
			Name:   "list-source-resources",
			Fields: fields_list_source_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSourceResourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_source_resources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSourceResources(ctx, input)
				}
				var results []*svc.ListSourceResourcesOutput
				p := svc.NewListSourceResourcesPaginator(client, input)
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
		"notify-application-state": {
			Name:   "notify-application-state",
			Fields: fields_notify_application_state,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.NotifyApplicationStateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_notify_application_state, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.NotifyApplicationState(ctx, input)
			},
		},
		"notify-migration-task-state": {
			Name:   "notify-migration-task-state",
			Fields: fields_notify_migration_task_state,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.NotifyMigrationTaskStateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_notify_migration_task_state, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.NotifyMigrationTaskState(ctx, input)
			},
		},
		"put-resource-attributes": {
			Name:   "put-resource-attributes",
			Fields: fields_put_resource_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutResourceAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_resource_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutResourceAttributes(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("migrationhub", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
