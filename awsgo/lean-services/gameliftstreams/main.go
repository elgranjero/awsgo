package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/gameliftstreams"
)

var fields_add_stream_group_locations = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "LocationConfigurations", Flag: "location-configurations", Type: "[]types.LocationConfiguration", Required: true},
}

var fields_associate_applications = []leanruntime.Field{
	{Name: "ApplicationIdentifiers", Flag: "application-identifiers", Type: "[]string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_create_application = []leanruntime.Field{
	{Name: "ApplicationLogOutputUri", Flag: "application-log-output-uri", Type: "*string", Required: false},
	{Name: "ApplicationLogPaths", Flag: "application-log-paths", Type: "[]string", Required: false},
	{Name: "ApplicationSourceUri", Flag: "application-source-uri", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "ExecutablePath", Flag: "executable-path", Type: "*string", Required: true},
	{Name: "RuntimeEnvironment", Flag: "runtime-environment", Type: "*types.RuntimeEnvironment", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_stream_group = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DefaultApplicationIdentifier", Flag: "default-application-identifier", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "LocationConfigurations", Flag: "location-configurations", Type: "[]types.LocationConfiguration", Required: false},
	{Name: "StreamClass", Flag: "stream-class", Type: "types.StreamClass", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_stream_session_connection = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "SignalRequest", Flag: "signal-request", Type: "*string", Required: true},
	{Name: "StreamSessionIdentifier", Flag: "stream-session-identifier", Type: "*string", Required: true},
}

var fields_delete_application = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_stream_group = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_disassociate_applications = []leanruntime.Field{
	{Name: "ApplicationIdentifiers", Flag: "application-identifiers", Type: "[]string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_export_stream_session_files = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "OutputUri", Flag: "output-uri", Type: "*string", Required: true},
	{Name: "StreamSessionIdentifier", Flag: "stream-session-identifier", Type: "*string", Required: true},
}

var fields_get_application = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_stream_group = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_stream_session = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "StreamSessionIdentifier", Flag: "stream-session-identifier", Type: "*string", Required: true},
}

var fields_list_applications = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_stream_groups = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_stream_sessions = []leanruntime.Field{
	{Name: "ExportFilesStatus", Flag: "export-files-status", Type: "types.ExportFilesStatus", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.StreamSessionStatus", Required: false},
}

var fields_list_stream_sessions_by_account = []leanruntime.Field{
	{Name: "ExportFilesStatus", Flag: "export-files-status", Type: "types.ExportFilesStatus", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.StreamSessionStatus", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_remove_stream_group_locations = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Locations", Flag: "locations", Type: "[]string", Required: true},
}

var fields_start_stream_session = []leanruntime.Field{
	{Name: "AdditionalEnvironmentVariables", Flag: "additional-environment-variables", Type: "map[string]string", Required: false},
	{Name: "AdditionalLaunchArgs", Flag: "additional-launch-args", Type: "[]string", Required: false},
	{Name: "ApplicationIdentifier", Flag: "application-identifier", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ConnectionTimeoutSeconds", Flag: "connection-timeout-seconds", Type: "*int32", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Locations", Flag: "locations", Type: "[]string", Required: false},
	{Name: "PerformanceStatsConfiguration", Flag: "performance-stats-configuration", Type: "*types.PerformanceStatsConfiguration", Required: false},
	{Name: "Protocol", Flag: "protocol", Type: "types.Protocol", Required: true},
	{Name: "SessionLengthSeconds", Flag: "session-length-seconds", Type: "*int32", Required: false},
	{Name: "SignalRequest", Flag: "signal-request", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_terminate_stream_session = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "StreamSessionIdentifier", Flag: "stream-session-identifier", Type: "*string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_application = []leanruntime.Field{
	{Name: "ApplicationLogOutputUri", Flag: "application-log-output-uri", Type: "*string", Required: false},
	{Name: "ApplicationLogPaths", Flag: "application-log-paths", Type: "[]string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_update_stream_group = []leanruntime.Field{
	{Name: "DefaultApplicationIdentifier", Flag: "default-application-identifier", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "LocationConfigurations", Flag: "location-configurations", Type: "[]types.LocationConfiguration", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"add-stream-group-locations": {
			Name:   "add-stream-group-locations",
			Fields: fields_add_stream_group_locations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddStreamGroupLocationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_stream_group_locations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddStreamGroupLocations(ctx, input)
			},
		},
		"associate-applications": {
			Name:   "associate-applications",
			Fields: fields_associate_applications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateApplicationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_applications, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateApplications(ctx, input)
			},
		},
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
		"create-stream-group": {
			Name:   "create-stream-group",
			Fields: fields_create_stream_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateStreamGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_stream_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateStreamGroup(ctx, input)
			},
		},
		"create-stream-session-connection": {
			Name:   "create-stream-session-connection",
			Fields: fields_create_stream_session_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateStreamSessionConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_stream_session_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateStreamSessionConnection(ctx, input)
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
		"delete-stream-group": {
			Name:   "delete-stream-group",
			Fields: fields_delete_stream_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteStreamGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_stream_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteStreamGroup(ctx, input)
			},
		},
		"disassociate-applications": {
			Name:   "disassociate-applications",
			Fields: fields_disassociate_applications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateApplicationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_applications, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateApplications(ctx, input)
			},
		},
		"export-stream-session-files": {
			Name:   "export-stream-session-files",
			Fields: fields_export_stream_session_files,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExportStreamSessionFilesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_export_stream_session_files, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExportStreamSessionFiles(ctx, input)
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
		"get-stream-group": {
			Name:   "get-stream-group",
			Fields: fields_get_stream_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetStreamGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_stream_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetStreamGroup(ctx, input)
			},
		},
		"get-stream-session": {
			Name:   "get-stream-session",
			Fields: fields_get_stream_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetStreamSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_stream_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetStreamSession(ctx, input)
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
		"list-stream-groups": {
			Name:   "list-stream-groups",
			Fields: fields_list_stream_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStreamGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_stream_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStreamGroups(ctx, input)
				}
				var results []*svc.ListStreamGroupsOutput
				p := svc.NewListStreamGroupsPaginator(client, input)
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
		"list-stream-sessions": {
			Name:   "list-stream-sessions",
			Fields: fields_list_stream_sessions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStreamSessionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_stream_sessions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStreamSessions(ctx, input)
				}
				var results []*svc.ListStreamSessionsOutput
				p := svc.NewListStreamSessionsPaginator(client, input)
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
		"list-stream-sessions-by-account": {
			Name:   "list-stream-sessions-by-account",
			Fields: fields_list_stream_sessions_by_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStreamSessionsByAccountInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_stream_sessions_by_account, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStreamSessionsByAccount(ctx, input)
				}
				var results []*svc.ListStreamSessionsByAccountOutput
				p := svc.NewListStreamSessionsByAccountPaginator(client, input)
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
		"remove-stream-group-locations": {
			Name:   "remove-stream-group-locations",
			Fields: fields_remove_stream_group_locations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveStreamGroupLocationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_stream_group_locations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveStreamGroupLocations(ctx, input)
			},
		},
		"start-stream-session": {
			Name:   "start-stream-session",
			Fields: fields_start_stream_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartStreamSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_stream_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartStreamSession(ctx, input)
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
		"terminate-stream-session": {
			Name:   "terminate-stream-session",
			Fields: fields_terminate_stream_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TerminateStreamSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_terminate_stream_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TerminateStreamSession(ctx, input)
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
		"update-stream-group": {
			Name:   "update-stream-group",
			Fields: fields_update_stream_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateStreamGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_stream_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateStreamGroup(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("gameliftstreams", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
