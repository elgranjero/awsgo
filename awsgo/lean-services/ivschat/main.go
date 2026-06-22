package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/ivschat"
)

var fields_create_chat_token = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "map[string]string", Required: false},
	{Name: "Capabilities", Flag: "capabilities", Type: "[]types.ChatTokenCapability", Required: false},
	{Name: "RoomIdentifier", Flag: "room-identifier", Type: "*string", Required: true},
	{Name: "SessionDurationInMinutes", Flag: "session-duration-in-minutes", Type: "*int32", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_create_logging_configuration = []leanruntime.Field{
	{Name: "DestinationConfiguration", Flag: "destination-configuration", Type: "types.DestinationConfiguration", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_room = []leanruntime.Field{
	{Name: "LoggingConfigurationIdentifiers", Flag: "logging-configuration-identifiers", Type: "[]string", Required: false},
	{Name: "MaximumMessageLength", Flag: "maximum-message-length", Type: "*int32", Required: false},
	{Name: "MaximumMessageRatePerSecond", Flag: "maximum-message-rate-per-second", Type: "*int32", Required: false},
	{Name: "MessageReviewHandler", Flag: "message-review-handler", Type: "*types.MessageReviewHandler", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_logging_configuration = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_message = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Reason", Flag: "reason", Type: "*string", Required: false},
	{Name: "RoomIdentifier", Flag: "room-identifier", Type: "*string", Required: true},
}

var fields_delete_room = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_disconnect_user = []leanruntime.Field{
	{Name: "Reason", Flag: "reason", Type: "*string", Required: false},
	{Name: "RoomIdentifier", Flag: "room-identifier", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_get_logging_configuration = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_room = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_list_logging_configurations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_rooms = []leanruntime.Field{
	{Name: "LoggingConfigurationIdentifier", Flag: "logging-configuration-identifier", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MessageReviewHandlerUri", Flag: "message-review-handler-uri", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_send_event = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "map[string]string", Required: false},
	{Name: "EventName", Flag: "event-name", Type: "*string", Required: true},
	{Name: "RoomIdentifier", Flag: "room-identifier", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_logging_configuration = []leanruntime.Field{
	{Name: "DestinationConfiguration", Flag: "destination-configuration", Type: "types.DestinationConfiguration", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_room = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "LoggingConfigurationIdentifiers", Flag: "logging-configuration-identifiers", Type: "[]string", Required: false},
	{Name: "MaximumMessageLength", Flag: "maximum-message-length", Type: "*int32", Required: false},
	{Name: "MaximumMessageRatePerSecond", Flag: "maximum-message-rate-per-second", Type: "*int32", Required: false},
	{Name: "MessageReviewHandler", Flag: "message-review-handler", Type: "*types.MessageReviewHandler", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-chat-token": {
			Name:   "create-chat-token",
			Fields: fields_create_chat_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateChatTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_chat_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateChatToken(ctx, input)
			},
		},
		"create-logging-configuration": {
			Name:   "create-logging-configuration",
			Fields: fields_create_logging_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLoggingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_logging_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLoggingConfiguration(ctx, input)
			},
		},
		"create-room": {
			Name:   "create-room",
			Fields: fields_create_room,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRoomInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_room, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRoom(ctx, input)
			},
		},
		"delete-logging-configuration": {
			Name:   "delete-logging-configuration",
			Fields: fields_delete_logging_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLoggingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_logging_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLoggingConfiguration(ctx, input)
			},
		},
		"delete-message": {
			Name:   "delete-message",
			Fields: fields_delete_message,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMessageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_message, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMessage(ctx, input)
			},
		},
		"delete-room": {
			Name:   "delete-room",
			Fields: fields_delete_room,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRoomInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_room, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRoom(ctx, input)
			},
		},
		"disconnect-user": {
			Name:   "disconnect-user",
			Fields: fields_disconnect_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisconnectUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disconnect_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisconnectUser(ctx, input)
			},
		},
		"get-logging-configuration": {
			Name:   "get-logging-configuration",
			Fields: fields_get_logging_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLoggingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_logging_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLoggingConfiguration(ctx, input)
			},
		},
		"get-room": {
			Name:   "get-room",
			Fields: fields_get_room,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRoomInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_room, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRoom(ctx, input)
			},
		},
		"list-logging-configurations": {
			Name:   "list-logging-configurations",
			Fields: fields_list_logging_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLoggingConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_logging_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLoggingConfigurations(ctx, input)
				}
				var results []*svc.ListLoggingConfigurationsOutput
				p := svc.NewListLoggingConfigurationsPaginator(client, input)
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
		"list-rooms": {
			Name:   "list-rooms",
			Fields: fields_list_rooms,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRoomsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_rooms, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRooms(ctx, input)
				}
				var results []*svc.ListRoomsOutput
				p := svc.NewListRoomsPaginator(client, input)
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
		"send-event": {
			Name:   "send-event",
			Fields: fields_send_event,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendEventInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_event, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendEvent(ctx, input)
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
		"update-logging-configuration": {
			Name:   "update-logging-configuration",
			Fields: fields_update_logging_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLoggingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_logging_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLoggingConfiguration(ctx, input)
			},
		},
		"update-room": {
			Name:   "update-room",
			Fields: fields_update_room,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRoomInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_room, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRoom(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("ivschat", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
