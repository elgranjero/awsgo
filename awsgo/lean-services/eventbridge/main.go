package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/eventbridge"
)

var fields_activate_event_source = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_cancel_replay = []leanruntime.Field{
	{Name: "ReplayName", Flag: "replay-name", Type: "*string", Required: true},
}

var fields_create_api_destination = []leanruntime.Field{
	{Name: "ConnectionArn", Flag: "connection-arn", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "HttpMethod", Flag: "http-method", Type: "types.ApiDestinationHttpMethod", Required: true},
	{Name: "InvocationEndpoint", Flag: "invocation-endpoint", Type: "*string", Required: true},
	{Name: "InvocationRateLimitPerSecond", Flag: "invocation-rate-limit-per-second", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_archive = []leanruntime.Field{
	{Name: "ArchiveName", Flag: "archive-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EventPattern", Flag: "event-pattern", Type: "*string", Required: false},
	{Name: "EventSourceArn", Flag: "event-source-arn", Type: "*string", Required: true},
	{Name: "KmsKeyIdentifier", Flag: "kms-key-identifier", Type: "*string", Required: false},
	{Name: "RetentionDays", Flag: "retention-days", Type: "*int32", Required: false},
}

var fields_create_connection = []leanruntime.Field{
	{Name: "AuthParameters", Flag: "auth-parameters", Type: "*types.CreateConnectionAuthRequestParameters", Required: true},
	{Name: "AuthorizationType", Flag: "authorization-type", Type: "types.ConnectionAuthorizationType", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InvocationConnectivityParameters", Flag: "invocation-connectivity-parameters", Type: "*types.ConnectivityResourceParameters", Required: false},
	{Name: "KmsKeyIdentifier", Flag: "kms-key-identifier", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_endpoint = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EventBuses", Flag: "event-buses", Type: "[]types.EndpointEventBus", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ReplicationConfig", Flag: "replication-config", Type: "*types.ReplicationConfig", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "RoutingConfig", Flag: "routing-config", Type: "*types.RoutingConfig", Required: true},
}

var fields_create_event_bus = []leanruntime.Field{
	{Name: "DeadLetterConfig", Flag: "dead-letter-config", Type: "*types.DeadLetterConfig", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EventSourceName", Flag: "event-source-name", Type: "*string", Required: false},
	{Name: "KmsKeyIdentifier", Flag: "kms-key-identifier", Type: "*string", Required: false},
	{Name: "LogConfig", Flag: "log-config", Type: "*types.LogConfig", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_partner_event_source = []leanruntime.Field{
	{Name: "Account", Flag: "account", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_deactivate_event_source = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_deauthorize_connection = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_api_destination = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_archive = []leanruntime.Field{
	{Name: "ArchiveName", Flag: "archive-name", Type: "*string", Required: true},
}

var fields_delete_connection = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_endpoint = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_event_bus = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_partner_event_source = []leanruntime.Field{
	{Name: "Account", Flag: "account", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_rule = []leanruntime.Field{
	{Name: "EventBusName", Flag: "event-bus-name", Type: "*string", Required: false},
	{Name: "Force", Flag: "force", Type: "bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_describe_api_destination = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_describe_archive = []leanruntime.Field{
	{Name: "ArchiveName", Flag: "archive-name", Type: "*string", Required: true},
}

var fields_describe_connection = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_describe_endpoint = []leanruntime.Field{
	{Name: "HomeRegion", Flag: "home-region", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_describe_event_bus = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_describe_event_source = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_describe_partner_event_source = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_describe_replay = []leanruntime.Field{
	{Name: "ReplayName", Flag: "replay-name", Type: "*string", Required: true},
}

var fields_describe_rule = []leanruntime.Field{
	{Name: "EventBusName", Flag: "event-bus-name", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_disable_rule = []leanruntime.Field{
	{Name: "EventBusName", Flag: "event-bus-name", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_enable_rule = []leanruntime.Field{
	{Name: "EventBusName", Flag: "event-bus-name", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_list_api_destinations = []leanruntime.Field{
	{Name: "ConnectionArn", Flag: "connection-arn", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NamePrefix", Flag: "name-prefix", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_archives = []leanruntime.Field{
	{Name: "EventSourceArn", Flag: "event-source-arn", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NamePrefix", Flag: "name-prefix", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "State", Flag: "state", Type: "types.ArchiveState", Required: false},
}

var fields_list_connections = []leanruntime.Field{
	{Name: "ConnectionState", Flag: "connection-state", Type: "types.ConnectionState", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NamePrefix", Flag: "name-prefix", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_endpoints = []leanruntime.Field{
	{Name: "HomeRegion", Flag: "home-region", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NamePrefix", Flag: "name-prefix", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_event_buses = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NamePrefix", Flag: "name-prefix", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_event_sources = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NamePrefix", Flag: "name-prefix", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_partner_event_source_accounts = []leanruntime.Field{
	{Name: "EventSourceName", Flag: "event-source-name", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_partner_event_sources = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NamePrefix", Flag: "name-prefix", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_replays = []leanruntime.Field{
	{Name: "EventSourceArn", Flag: "event-source-arn", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NamePrefix", Flag: "name-prefix", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "State", Flag: "state", Type: "types.ReplayState", Required: false},
}

var fields_list_rule_names_by_target = []leanruntime.Field{
	{Name: "EventBusName", Flag: "event-bus-name", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TargetArn", Flag: "target-arn", Type: "*string", Required: true},
}

var fields_list_rules = []leanruntime.Field{
	{Name: "EventBusName", Flag: "event-bus-name", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NamePrefix", Flag: "name-prefix", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_targets_by_rule = []leanruntime.Field{
	{Name: "EventBusName", Flag: "event-bus-name", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Rule", Flag: "rule", Type: "*string", Required: true},
}

var fields_put_events = []leanruntime.Field{
	{Name: "EndpointId", Flag: "endpoint-id", Type: "*string", Required: false},
	{Name: "Entries", Flag: "entries", Type: "[]types.PutEventsRequestEntry", Required: true},
}

var fields_put_partner_events = []leanruntime.Field{
	{Name: "Entries", Flag: "entries", Type: "[]types.PutPartnerEventsRequestEntry", Required: true},
}

var fields_put_permission = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "*string", Required: false},
	{Name: "Condition", Flag: "condition", Type: "*types.Condition", Required: false},
	{Name: "EventBusName", Flag: "event-bus-name", Type: "*string", Required: false},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: false},
	{Name: "Principal", Flag: "principal", Type: "*string", Required: false},
	{Name: "StatementId", Flag: "statement-id", Type: "*string", Required: false},
}

var fields_put_rule = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EventBusName", Flag: "event-bus-name", Type: "*string", Required: false},
	{Name: "EventPattern", Flag: "event-pattern", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "ScheduleExpression", Flag: "schedule-expression", Type: "*string", Required: false},
	{Name: "State", Flag: "state", Type: "types.RuleState", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_put_targets = []leanruntime.Field{
	{Name: "EventBusName", Flag: "event-bus-name", Type: "*string", Required: false},
	{Name: "Rule", Flag: "rule", Type: "*string", Required: true},
	{Name: "Targets", Flag: "targets", Type: "[]types.Target", Required: true},
}

var fields_remove_permission = []leanruntime.Field{
	{Name: "EventBusName", Flag: "event-bus-name", Type: "*string", Required: false},
	{Name: "RemoveAllPermissions", Flag: "remove-all-permissions", Type: "bool", Required: false},
	{Name: "StatementId", Flag: "statement-id", Type: "*string", Required: false},
}

var fields_remove_targets = []leanruntime.Field{
	{Name: "EventBusName", Flag: "event-bus-name", Type: "*string", Required: false},
	{Name: "Force", Flag: "force", Type: "bool", Required: false},
	{Name: "Ids", Flag: "ids", Type: "[]string", Required: true},
	{Name: "Rule", Flag: "rule", Type: "*string", Required: true},
}

var fields_start_replay = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Destination", Flag: "destination", Type: "*types.ReplayDestination", Required: true},
	{Name: "EventEndTime", Flag: "event-end-time", Type: "*time.Time", Required: true},
	{Name: "EventSourceArn", Flag: "event-source-arn", Type: "*string", Required: true},
	{Name: "EventStartTime", Flag: "event-start-time", Type: "*time.Time", Required: true},
	{Name: "ReplayName", Flag: "replay-name", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_test_event_pattern = []leanruntime.Field{
	{Name: "Event", Flag: "event", Type: "*string", Required: true},
	{Name: "EventPattern", Flag: "event-pattern", Type: "*string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_api_destination = []leanruntime.Field{
	{Name: "ConnectionArn", Flag: "connection-arn", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "HttpMethod", Flag: "http-method", Type: "types.ApiDestinationHttpMethod", Required: false},
	{Name: "InvocationEndpoint", Flag: "invocation-endpoint", Type: "*string", Required: false},
	{Name: "InvocationRateLimitPerSecond", Flag: "invocation-rate-limit-per-second", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_update_archive = []leanruntime.Field{
	{Name: "ArchiveName", Flag: "archive-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EventPattern", Flag: "event-pattern", Type: "*string", Required: false},
	{Name: "KmsKeyIdentifier", Flag: "kms-key-identifier", Type: "*string", Required: false},
	{Name: "RetentionDays", Flag: "retention-days", Type: "*int32", Required: false},
}

var fields_update_connection = []leanruntime.Field{
	{Name: "AuthParameters", Flag: "auth-parameters", Type: "*types.UpdateConnectionAuthRequestParameters", Required: false},
	{Name: "AuthorizationType", Flag: "authorization-type", Type: "types.ConnectionAuthorizationType", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "InvocationConnectivityParameters", Flag: "invocation-connectivity-parameters", Type: "*types.ConnectivityResourceParameters", Required: false},
	{Name: "KmsKeyIdentifier", Flag: "kms-key-identifier", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_update_endpoint = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EventBuses", Flag: "event-buses", Type: "[]types.EndpointEventBus", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ReplicationConfig", Flag: "replication-config", Type: "*types.ReplicationConfig", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "RoutingConfig", Flag: "routing-config", Type: "*types.RoutingConfig", Required: false},
}

var fields_update_event_bus = []leanruntime.Field{
	{Name: "DeadLetterConfig", Flag: "dead-letter-config", Type: "*types.DeadLetterConfig", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "KmsKeyIdentifier", Flag: "kms-key-identifier", Type: "*string", Required: false},
	{Name: "LogConfig", Flag: "log-config", Type: "*types.LogConfig", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"activate-event-source": {
			Name:   "activate-event-source",
			Fields: fields_activate_event_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ActivateEventSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_activate_event_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ActivateEventSource(ctx, input)
			},
		},
		"cancel-replay": {
			Name:   "cancel-replay",
			Fields: fields_cancel_replay,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelReplayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_replay, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelReplay(ctx, input)
			},
		},
		"create-api-destination": {
			Name:   "create-api-destination",
			Fields: fields_create_api_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateApiDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_api_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateApiDestination(ctx, input)
			},
		},
		"create-archive": {
			Name:   "create-archive",
			Fields: fields_create_archive,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateArchiveInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_archive, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateArchive(ctx, input)
			},
		},
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
		"create-endpoint": {
			Name:   "create-endpoint",
			Fields: fields_create_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEndpoint(ctx, input)
			},
		},
		"create-event-bus": {
			Name:   "create-event-bus",
			Fields: fields_create_event_bus,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEventBusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_event_bus, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEventBus(ctx, input)
			},
		},
		"create-partner-event-source": {
			Name:   "create-partner-event-source",
			Fields: fields_create_partner_event_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePartnerEventSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_partner_event_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePartnerEventSource(ctx, input)
			},
		},
		"deactivate-event-source": {
			Name:   "deactivate-event-source",
			Fields: fields_deactivate_event_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeactivateEventSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deactivate_event_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeactivateEventSource(ctx, input)
			},
		},
		"deauthorize-connection": {
			Name:   "deauthorize-connection",
			Fields: fields_deauthorize_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeauthorizeConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deauthorize_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeauthorizeConnection(ctx, input)
			},
		},
		"delete-api-destination": {
			Name:   "delete-api-destination",
			Fields: fields_delete_api_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteApiDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_api_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApiDestination(ctx, input)
			},
		},
		"delete-archive": {
			Name:   "delete-archive",
			Fields: fields_delete_archive,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteArchiveInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_archive, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteArchive(ctx, input)
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
		"delete-endpoint": {
			Name:   "delete-endpoint",
			Fields: fields_delete_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEndpoint(ctx, input)
			},
		},
		"delete-event-bus": {
			Name:   "delete-event-bus",
			Fields: fields_delete_event_bus,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEventBusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_event_bus, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEventBus(ctx, input)
			},
		},
		"delete-partner-event-source": {
			Name:   "delete-partner-event-source",
			Fields: fields_delete_partner_event_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePartnerEventSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_partner_event_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePartnerEventSource(ctx, input)
			},
		},
		"delete-rule": {
			Name:   "delete-rule",
			Fields: fields_delete_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRule(ctx, input)
			},
		},
		"describe-api-destination": {
			Name:   "describe-api-destination",
			Fields: fields_describe_api_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeApiDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_api_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeApiDestination(ctx, input)
			},
		},
		"describe-archive": {
			Name:   "describe-archive",
			Fields: fields_describe_archive,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeArchiveInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_archive, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeArchive(ctx, input)
			},
		},
		"describe-connection": {
			Name:   "describe-connection",
			Fields: fields_describe_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeConnection(ctx, input)
			},
		},
		"describe-endpoint": {
			Name:   "describe-endpoint",
			Fields: fields_describe_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEndpoint(ctx, input)
			},
		},
		"describe-event-bus": {
			Name:   "describe-event-bus",
			Fields: fields_describe_event_bus,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEventBusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_event_bus, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEventBus(ctx, input)
			},
		},
		"describe-event-source": {
			Name:   "describe-event-source",
			Fields: fields_describe_event_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEventSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_event_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEventSource(ctx, input)
			},
		},
		"describe-partner-event-source": {
			Name:   "describe-partner-event-source",
			Fields: fields_describe_partner_event_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePartnerEventSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_partner_event_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePartnerEventSource(ctx, input)
			},
		},
		"describe-replay": {
			Name:   "describe-replay",
			Fields: fields_describe_replay,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReplayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_replay, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeReplay(ctx, input)
			},
		},
		"describe-rule": {
			Name:   "describe-rule",
			Fields: fields_describe_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRule(ctx, input)
			},
		},
		"disable-rule": {
			Name:   "disable-rule",
			Fields: fields_disable_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableRule(ctx, input)
			},
		},
		"enable-rule": {
			Name:   "enable-rule",
			Fields: fields_enable_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableRule(ctx, input)
			},
		},
		"list-api-destinations": {
			Name:   "list-api-destinations",
			Fields: fields_list_api_destinations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListApiDestinationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_api_destinations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListApiDestinations(ctx, input)
			},
		},
		"list-archives": {
			Name:   "list-archives",
			Fields: fields_list_archives,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListArchivesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_archives, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListArchives(ctx, input)
			},
		},
		"list-connections": {
			Name:   "list-connections",
			Fields: fields_list_connections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConnectionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_connections, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListConnections(ctx, input)
			},
		},
		"list-endpoints": {
			Name:   "list-endpoints",
			Fields: fields_list_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEndpointsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_endpoints, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListEndpoints(ctx, input)
			},
		},
		"list-event-buses": {
			Name:   "list-event-buses",
			Fields: fields_list_event_buses,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEventBusesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_event_buses, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListEventBuses(ctx, input)
			},
		},
		"list-event-sources": {
			Name:   "list-event-sources",
			Fields: fields_list_event_sources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEventSourcesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_event_sources, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListEventSources(ctx, input)
			},
		},
		"list-partner-event-source-accounts": {
			Name:   "list-partner-event-source-accounts",
			Fields: fields_list_partner_event_source_accounts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPartnerEventSourceAccountsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_partner_event_source_accounts, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListPartnerEventSourceAccounts(ctx, input)
			},
		},
		"list-partner-event-sources": {
			Name:   "list-partner-event-sources",
			Fields: fields_list_partner_event_sources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPartnerEventSourcesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_partner_event_sources, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListPartnerEventSources(ctx, input)
			},
		},
		"list-replays": {
			Name:   "list-replays",
			Fields: fields_list_replays,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReplaysInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_replays, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListReplays(ctx, input)
			},
		},
		"list-rule-names-by-target": {
			Name:   "list-rule-names-by-target",
			Fields: fields_list_rule_names_by_target,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRuleNamesByTargetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_rule_names_by_target, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListRuleNamesByTarget(ctx, input)
			},
		},
		"list-rules": {
			Name:   "list-rules",
			Fields: fields_list_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRulesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_rules, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListRules(ctx, input)
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
		"list-targets-by-rule": {
			Name:   "list-targets-by-rule",
			Fields: fields_list_targets_by_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTargetsByRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_targets_by_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTargetsByRule(ctx, input)
			},
		},
		"put-events": {
			Name:   "put-events",
			Fields: fields_put_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutEventsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_events, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutEvents(ctx, input)
			},
		},
		"put-partner-events": {
			Name:   "put-partner-events",
			Fields: fields_put_partner_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutPartnerEventsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_partner_events, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutPartnerEvents(ctx, input)
			},
		},
		"put-permission": {
			Name:   "put-permission",
			Fields: fields_put_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutPermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutPermission(ctx, input)
			},
		},
		"put-rule": {
			Name:   "put-rule",
			Fields: fields_put_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutRule(ctx, input)
			},
		},
		"put-targets": {
			Name:   "put-targets",
			Fields: fields_put_targets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutTargetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_targets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutTargets(ctx, input)
			},
		},
		"remove-permission": {
			Name:   "remove-permission",
			Fields: fields_remove_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemovePermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemovePermission(ctx, input)
			},
		},
		"remove-targets": {
			Name:   "remove-targets",
			Fields: fields_remove_targets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveTargetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_targets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveTargets(ctx, input)
			},
		},
		"start-replay": {
			Name:   "start-replay",
			Fields: fields_start_replay,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartReplayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_replay, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartReplay(ctx, input)
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
		"test-event-pattern": {
			Name:   "test-event-pattern",
			Fields: fields_test_event_pattern,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TestEventPatternInput{}
				if _, err := leanruntime.ApplyInput(input, fields_test_event_pattern, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TestEventPattern(ctx, input)
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
		"update-api-destination": {
			Name:   "update-api-destination",
			Fields: fields_update_api_destination,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateApiDestinationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_api_destination, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateApiDestination(ctx, input)
			},
		},
		"update-archive": {
			Name:   "update-archive",
			Fields: fields_update_archive,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateArchiveInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_archive, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateArchive(ctx, input)
			},
		},
		"update-connection": {
			Name:   "update-connection",
			Fields: fields_update_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConnection(ctx, input)
			},
		},
		"update-endpoint": {
			Name:   "update-endpoint",
			Fields: fields_update_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEndpoint(ctx, input)
			},
		},
		"update-event-bus": {
			Name:   "update-event-bus",
			Fields: fields_update_event_bus,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEventBusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_event_bus, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEventBus(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("eventbridge", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
