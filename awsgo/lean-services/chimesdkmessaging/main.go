package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/chimesdkmessaging"
)

var fields_associate_channel_flow = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ChannelFlowArn", Flag: "channel-flow-arn", Type: "*string", Required: true},
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: true},
}

var fields_batch_create_channel_membership = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: true},
	{Name: "MemberArns", Flag: "member-arns", Type: "[]string", Required: true},
	{Name: "SubChannelId", Flag: "sub-channel-id", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.ChannelMembershipType", Required: false},
}

var fields_channel_flow_callback = []leanruntime.Field{
	{Name: "CallbackId", Flag: "callback-id", Type: "*string", Required: true},
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ChannelMessage", Flag: "channel-message", Type: "*types.ChannelMessageCallback", Required: true},
	{Name: "DeleteResource", Flag: "delete-resource", Type: "bool", Required: false},
}

var fields_create_channel = []leanruntime.Field{
	{Name: "AppInstanceArn", Flag: "app-instance-arn", Type: "*string", Required: true},
	{Name: "ChannelId", Flag: "channel-id", Type: "*string", Required: false},
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "ElasticChannelConfiguration", Flag: "elastic-channel-configuration", Type: "*types.ElasticChannelConfiguration", Required: false},
	{Name: "ExpirationSettings", Flag: "expiration-settings", Type: "*types.ExpirationSettings", Required: false},
	{Name: "MemberArns", Flag: "member-arns", Type: "[]string", Required: false},
	{Name: "Metadata", Flag: "metadata", Type: "*string", Required: false},
	{Name: "Mode", Flag: "mode", Type: "types.ChannelMode", Required: false},
	{Name: "ModeratorArns", Flag: "moderator-arns", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Privacy", Flag: "privacy", Type: "types.ChannelPrivacy", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_channel_ban = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: true},
	{Name: "MemberArn", Flag: "member-arn", Type: "*string", Required: true},
}

var fields_create_channel_flow = []leanruntime.Field{
	{Name: "AppInstanceArn", Flag: "app-instance-arn", Type: "*string", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Processors", Flag: "processors", Type: "[]types.Processor", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_channel_membership = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: true},
	{Name: "MemberArn", Flag: "member-arn", Type: "*string", Required: true},
	{Name: "SubChannelId", Flag: "sub-channel-id", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.ChannelMembershipType", Required: true},
}

var fields_create_channel_moderator = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ChannelModeratorArn", Flag: "channel-moderator-arn", Type: "*string", Required: true},
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: true},
}

var fields_delete_channel = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: true},
}

var fields_delete_channel_ban = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: true},
	{Name: "MemberArn", Flag: "member-arn", Type: "*string", Required: true},
}

var fields_delete_channel_flow = []leanruntime.Field{
	{Name: "ChannelFlowArn", Flag: "channel-flow-arn", Type: "*string", Required: true},
}

var fields_delete_channel_membership = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: true},
	{Name: "MemberArn", Flag: "member-arn", Type: "*string", Required: true},
	{Name: "SubChannelId", Flag: "sub-channel-id", Type: "*string", Required: false},
}

var fields_delete_channel_message = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: true},
	{Name: "MessageId", Flag: "message-id", Type: "*string", Required: true},
	{Name: "SubChannelId", Flag: "sub-channel-id", Type: "*string", Required: false},
}

var fields_delete_channel_moderator = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ChannelModeratorArn", Flag: "channel-moderator-arn", Type: "*string", Required: true},
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: true},
}

var fields_delete_messaging_streaming_configurations = []leanruntime.Field{
	{Name: "AppInstanceArn", Flag: "app-instance-arn", Type: "*string", Required: true},
}

var fields_describe_channel = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: true},
}

var fields_describe_channel_ban = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: true},
	{Name: "MemberArn", Flag: "member-arn", Type: "*string", Required: true},
}

var fields_describe_channel_flow = []leanruntime.Field{
	{Name: "ChannelFlowArn", Flag: "channel-flow-arn", Type: "*string", Required: true},
}

var fields_describe_channel_membership = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: true},
	{Name: "MemberArn", Flag: "member-arn", Type: "*string", Required: true},
	{Name: "SubChannelId", Flag: "sub-channel-id", Type: "*string", Required: false},
}

var fields_describe_channel_membership_for_app_instance_user = []leanruntime.Field{
	{Name: "AppInstanceUserArn", Flag: "app-instance-user-arn", Type: "*string", Required: true},
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: true},
}

var fields_describe_channel_moderated_by_app_instance_user = []leanruntime.Field{
	{Name: "AppInstanceUserArn", Flag: "app-instance-user-arn", Type: "*string", Required: true},
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: true},
}

var fields_describe_channel_moderator = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ChannelModeratorArn", Flag: "channel-moderator-arn", Type: "*string", Required: true},
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: true},
}

var fields_disassociate_channel_flow = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ChannelFlowArn", Flag: "channel-flow-arn", Type: "*string", Required: true},
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: true},
}

var fields_get_channel_membership_preferences = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: true},
	{Name: "MemberArn", Flag: "member-arn", Type: "*string", Required: true},
}

var fields_get_channel_message = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: true},
	{Name: "MessageId", Flag: "message-id", Type: "*string", Required: true},
	{Name: "SubChannelId", Flag: "sub-channel-id", Type: "*string", Required: false},
}

var fields_get_channel_message_status = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: true},
	{Name: "MessageId", Flag: "message-id", Type: "*string", Required: true},
	{Name: "SubChannelId", Flag: "sub-channel-id", Type: "*string", Required: false},
}

var fields_get_messaging_session_endpoint = []leanruntime.Field{
	{Name: "NetworkType", Flag: "network-type", Type: "types.NetworkType", Required: false},
}

var fields_get_messaging_streaming_configurations = []leanruntime.Field{
	{Name: "AppInstanceArn", Flag: "app-instance-arn", Type: "*string", Required: true},
}

var fields_list_channel_bans = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_channel_flows = []leanruntime.Field{
	{Name: "AppInstanceArn", Flag: "app-instance-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_channel_memberships = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SubChannelId", Flag: "sub-channel-id", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.ChannelMembershipType", Required: false},
}

var fields_list_channel_memberships_for_app_instance_user = []leanruntime.Field{
	{Name: "AppInstanceUserArn", Flag: "app-instance-user-arn", Type: "*string", Required: false},
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_channel_messages = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "NotAfter", Flag: "not-after", Type: "*time.Time", Required: false},
	{Name: "NotBefore", Flag: "not-before", Type: "*time.Time", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "SubChannelId", Flag: "sub-channel-id", Type: "*string", Required: false},
}

var fields_list_channel_moderators = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_channels = []leanruntime.Field{
	{Name: "AppInstanceArn", Flag: "app-instance-arn", Type: "*string", Required: true},
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Privacy", Flag: "privacy", Type: "types.ChannelPrivacy", Required: false},
}

var fields_list_channels_associated_with_channel_flow = []leanruntime.Field{
	{Name: "ChannelFlowArn", Flag: "channel-flow-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_channels_moderated_by_app_instance_user = []leanruntime.Field{
	{Name: "AppInstanceUserArn", Flag: "app-instance-user-arn", Type: "*string", Required: false},
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_sub_channels = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_channel_expiration_settings = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: false},
	{Name: "ExpirationSettings", Flag: "expiration-settings", Type: "*types.ExpirationSettings", Required: false},
}

var fields_put_channel_membership_preferences = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: true},
	{Name: "MemberArn", Flag: "member-arn", Type: "*string", Required: true},
	{Name: "Preferences", Flag: "preferences", Type: "*types.ChannelMembershipPreferences", Required: true},
}

var fields_put_messaging_streaming_configurations = []leanruntime.Field{
	{Name: "AppInstanceArn", Flag: "app-instance-arn", Type: "*string", Required: true},
	{Name: "StreamingConfigurations", Flag: "streaming-configurations", Type: "[]types.StreamingConfiguration", Required: true},
}

var fields_redact_channel_message = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: true},
	{Name: "MessageId", Flag: "message-id", Type: "*string", Required: true},
	{Name: "SubChannelId", Flag: "sub-channel-id", Type: "*string", Required: false},
}

var fields_search_channels = []leanruntime.Field{
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: false},
	{Name: "Fields", Flag: "fields", Type: "[]types.SearchField", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_send_channel_message = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "Content", Flag: "content", Type: "*string", Required: true},
	{Name: "ContentType", Flag: "content-type", Type: "*string", Required: false},
	{Name: "MessageAttributes", Flag: "message-attributes", Type: "map[string]types.MessageAttributeValue", Required: false},
	{Name: "Metadata", Flag: "metadata", Type: "*string", Required: false},
	{Name: "Persistence", Flag: "persistence", Type: "types.ChannelMessagePersistenceType", Required: true},
	{Name: "PushNotification", Flag: "push-notification", Type: "*types.PushNotificationConfiguration", Required: false},
	{Name: "SubChannelId", Flag: "sub-channel-id", Type: "*string", Required: false},
	{Name: "Target", Flag: "target", Type: "[]types.Target", Required: false},
	{Name: "Type", Flag: "type", Type: "types.ChannelMessageType", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_channel = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: true},
	{Name: "Metadata", Flag: "metadata", Type: "*string", Required: false},
	{Name: "Mode", Flag: "mode", Type: "types.ChannelMode", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_channel_flow = []leanruntime.Field{
	{Name: "ChannelFlowArn", Flag: "channel-flow-arn", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Processors", Flag: "processors", Type: "[]types.Processor", Required: true},
}

var fields_update_channel_message = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: true},
	{Name: "Content", Flag: "content", Type: "*string", Required: true},
	{Name: "ContentType", Flag: "content-type", Type: "*string", Required: false},
	{Name: "MessageId", Flag: "message-id", Type: "*string", Required: true},
	{Name: "Metadata", Flag: "metadata", Type: "*string", Required: false},
	{Name: "SubChannelId", Flag: "sub-channel-id", Type: "*string", Required: false},
}

var fields_update_channel_read_marker = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ChimeBearer", Flag: "chime-bearer", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-channel-flow": {
			Name:   "associate-channel-flow",
			Fields: fields_associate_channel_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateChannelFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_channel_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateChannelFlow(ctx, input)
			},
		},
		"batch-create-channel-membership": {
			Name:   "batch-create-channel-membership",
			Fields: fields_batch_create_channel_membership,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchCreateChannelMembershipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_create_channel_membership, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchCreateChannelMembership(ctx, input)
			},
		},
		"channel-flow-callback": {
			Name:   "channel-flow-callback",
			Fields: fields_channel_flow_callback,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ChannelFlowCallbackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_channel_flow_callback, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ChannelFlowCallback(ctx, input)
			},
		},
		"create-channel": {
			Name:   "create-channel",
			Fields: fields_create_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateChannel(ctx, input)
			},
		},
		"create-channel-ban": {
			Name:   "create-channel-ban",
			Fields: fields_create_channel_ban,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateChannelBanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_channel_ban, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateChannelBan(ctx, input)
			},
		},
		"create-channel-flow": {
			Name:   "create-channel-flow",
			Fields: fields_create_channel_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateChannelFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_channel_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateChannelFlow(ctx, input)
			},
		},
		"create-channel-membership": {
			Name:   "create-channel-membership",
			Fields: fields_create_channel_membership,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateChannelMembershipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_channel_membership, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateChannelMembership(ctx, input)
			},
		},
		"create-channel-moderator": {
			Name:   "create-channel-moderator",
			Fields: fields_create_channel_moderator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateChannelModeratorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_channel_moderator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateChannelModerator(ctx, input)
			},
		},
		"delete-channel": {
			Name:   "delete-channel",
			Fields: fields_delete_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteChannel(ctx, input)
			},
		},
		"delete-channel-ban": {
			Name:   "delete-channel-ban",
			Fields: fields_delete_channel_ban,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteChannelBanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_channel_ban, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteChannelBan(ctx, input)
			},
		},
		"delete-channel-flow": {
			Name:   "delete-channel-flow",
			Fields: fields_delete_channel_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteChannelFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_channel_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteChannelFlow(ctx, input)
			},
		},
		"delete-channel-membership": {
			Name:   "delete-channel-membership",
			Fields: fields_delete_channel_membership,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteChannelMembershipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_channel_membership, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteChannelMembership(ctx, input)
			},
		},
		"delete-channel-message": {
			Name:   "delete-channel-message",
			Fields: fields_delete_channel_message,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteChannelMessageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_channel_message, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteChannelMessage(ctx, input)
			},
		},
		"delete-channel-moderator": {
			Name:   "delete-channel-moderator",
			Fields: fields_delete_channel_moderator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteChannelModeratorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_channel_moderator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteChannelModerator(ctx, input)
			},
		},
		"delete-messaging-streaming-configurations": {
			Name:   "delete-messaging-streaming-configurations",
			Fields: fields_delete_messaging_streaming_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMessagingStreamingConfigurationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_messaging_streaming_configurations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMessagingStreamingConfigurations(ctx, input)
			},
		},
		"describe-channel": {
			Name:   "describe-channel",
			Fields: fields_describe_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeChannel(ctx, input)
			},
		},
		"describe-channel-ban": {
			Name:   "describe-channel-ban",
			Fields: fields_describe_channel_ban,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeChannelBanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_channel_ban, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeChannelBan(ctx, input)
			},
		},
		"describe-channel-flow": {
			Name:   "describe-channel-flow",
			Fields: fields_describe_channel_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeChannelFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_channel_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeChannelFlow(ctx, input)
			},
		},
		"describe-channel-membership": {
			Name:   "describe-channel-membership",
			Fields: fields_describe_channel_membership,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeChannelMembershipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_channel_membership, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeChannelMembership(ctx, input)
			},
		},
		"describe-channel-membership-for-app-instance-user": {
			Name:   "describe-channel-membership-for-app-instance-user",
			Fields: fields_describe_channel_membership_for_app_instance_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeChannelMembershipForAppInstanceUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_channel_membership_for_app_instance_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeChannelMembershipForAppInstanceUser(ctx, input)
			},
		},
		"describe-channel-moderated-by-app-instance-user": {
			Name:   "describe-channel-moderated-by-app-instance-user",
			Fields: fields_describe_channel_moderated_by_app_instance_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeChannelModeratedByAppInstanceUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_channel_moderated_by_app_instance_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeChannelModeratedByAppInstanceUser(ctx, input)
			},
		},
		"describe-channel-moderator": {
			Name:   "describe-channel-moderator",
			Fields: fields_describe_channel_moderator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeChannelModeratorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_channel_moderator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeChannelModerator(ctx, input)
			},
		},
		"disassociate-channel-flow": {
			Name:   "disassociate-channel-flow",
			Fields: fields_disassociate_channel_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateChannelFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_channel_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateChannelFlow(ctx, input)
			},
		},
		"get-channel-membership-preferences": {
			Name:   "get-channel-membership-preferences",
			Fields: fields_get_channel_membership_preferences,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetChannelMembershipPreferencesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_channel_membership_preferences, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetChannelMembershipPreferences(ctx, input)
			},
		},
		"get-channel-message": {
			Name:   "get-channel-message",
			Fields: fields_get_channel_message,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetChannelMessageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_channel_message, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetChannelMessage(ctx, input)
			},
		},
		"get-channel-message-status": {
			Name:   "get-channel-message-status",
			Fields: fields_get_channel_message_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetChannelMessageStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_channel_message_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetChannelMessageStatus(ctx, input)
			},
		},
		"get-messaging-session-endpoint": {
			Name:   "get-messaging-session-endpoint",
			Fields: fields_get_messaging_session_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMessagingSessionEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_messaging_session_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMessagingSessionEndpoint(ctx, input)
			},
		},
		"get-messaging-streaming-configurations": {
			Name:   "get-messaging-streaming-configurations",
			Fields: fields_get_messaging_streaming_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMessagingStreamingConfigurationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_messaging_streaming_configurations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMessagingStreamingConfigurations(ctx, input)
			},
		},
		"list-channel-bans": {
			Name:   "list-channel-bans",
			Fields: fields_list_channel_bans,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListChannelBansInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_channel_bans, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListChannelBans(ctx, input)
				}
				var results []*svc.ListChannelBansOutput
				p := svc.NewListChannelBansPaginator(client, input)
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
		"list-channel-flows": {
			Name:   "list-channel-flows",
			Fields: fields_list_channel_flows,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListChannelFlowsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_channel_flows, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListChannelFlows(ctx, input)
				}
				var results []*svc.ListChannelFlowsOutput
				p := svc.NewListChannelFlowsPaginator(client, input)
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
		"list-channel-memberships": {
			Name:   "list-channel-memberships",
			Fields: fields_list_channel_memberships,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListChannelMembershipsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_channel_memberships, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListChannelMemberships(ctx, input)
				}
				var results []*svc.ListChannelMembershipsOutput
				p := svc.NewListChannelMembershipsPaginator(client, input)
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
		"list-channel-memberships-for-app-instance-user": {
			Name:   "list-channel-memberships-for-app-instance-user",
			Fields: fields_list_channel_memberships_for_app_instance_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListChannelMembershipsForAppInstanceUserInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_channel_memberships_for_app_instance_user, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListChannelMembershipsForAppInstanceUser(ctx, input)
				}
				var results []*svc.ListChannelMembershipsForAppInstanceUserOutput
				p := svc.NewListChannelMembershipsForAppInstanceUserPaginator(client, input)
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
		"list-channel-messages": {
			Name:   "list-channel-messages",
			Fields: fields_list_channel_messages,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListChannelMessagesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_channel_messages, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListChannelMessages(ctx, input)
				}
				var results []*svc.ListChannelMessagesOutput
				p := svc.NewListChannelMessagesPaginator(client, input)
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
		"list-channel-moderators": {
			Name:   "list-channel-moderators",
			Fields: fields_list_channel_moderators,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListChannelModeratorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_channel_moderators, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListChannelModerators(ctx, input)
				}
				var results []*svc.ListChannelModeratorsOutput
				p := svc.NewListChannelModeratorsPaginator(client, input)
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
		"list-channels": {
			Name:   "list-channels",
			Fields: fields_list_channels,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListChannelsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_channels, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListChannels(ctx, input)
				}
				var results []*svc.ListChannelsOutput
				p := svc.NewListChannelsPaginator(client, input)
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
		"list-channels-associated-with-channel-flow": {
			Name:   "list-channels-associated-with-channel-flow",
			Fields: fields_list_channels_associated_with_channel_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListChannelsAssociatedWithChannelFlowInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_channels_associated_with_channel_flow, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListChannelsAssociatedWithChannelFlow(ctx, input)
				}
				var results []*svc.ListChannelsAssociatedWithChannelFlowOutput
				p := svc.NewListChannelsAssociatedWithChannelFlowPaginator(client, input)
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
		"list-channels-moderated-by-app-instance-user": {
			Name:   "list-channels-moderated-by-app-instance-user",
			Fields: fields_list_channels_moderated_by_app_instance_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListChannelsModeratedByAppInstanceUserInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_channels_moderated_by_app_instance_user, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListChannelsModeratedByAppInstanceUser(ctx, input)
				}
				var results []*svc.ListChannelsModeratedByAppInstanceUserOutput
				p := svc.NewListChannelsModeratedByAppInstanceUserPaginator(client, input)
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
		"list-sub-channels": {
			Name:   "list-sub-channels",
			Fields: fields_list_sub_channels,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSubChannelsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_sub_channels, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSubChannels(ctx, input)
				}
				var results []*svc.ListSubChannelsOutput
				p := svc.NewListSubChannelsPaginator(client, input)
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
		"put-channel-expiration-settings": {
			Name:   "put-channel-expiration-settings",
			Fields: fields_put_channel_expiration_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutChannelExpirationSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_channel_expiration_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutChannelExpirationSettings(ctx, input)
			},
		},
		"put-channel-membership-preferences": {
			Name:   "put-channel-membership-preferences",
			Fields: fields_put_channel_membership_preferences,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutChannelMembershipPreferencesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_channel_membership_preferences, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutChannelMembershipPreferences(ctx, input)
			},
		},
		"put-messaging-streaming-configurations": {
			Name:   "put-messaging-streaming-configurations",
			Fields: fields_put_messaging_streaming_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutMessagingStreamingConfigurationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_messaging_streaming_configurations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutMessagingStreamingConfigurations(ctx, input)
			},
		},
		"redact-channel-message": {
			Name:   "redact-channel-message",
			Fields: fields_redact_channel_message,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RedactChannelMessageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_redact_channel_message, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RedactChannelMessage(ctx, input)
			},
		},
		"search-channels": {
			Name:   "search-channels",
			Fields: fields_search_channels,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchChannelsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_channels, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchChannels(ctx, input)
				}
				var results []*svc.SearchChannelsOutput
				p := svc.NewSearchChannelsPaginator(client, input)
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
		"send-channel-message": {
			Name:   "send-channel-message",
			Fields: fields_send_channel_message,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendChannelMessageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_channel_message, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendChannelMessage(ctx, input)
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
		"update-channel": {
			Name:   "update-channel",
			Fields: fields_update_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateChannel(ctx, input)
			},
		},
		"update-channel-flow": {
			Name:   "update-channel-flow",
			Fields: fields_update_channel_flow,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateChannelFlowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_channel_flow, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateChannelFlow(ctx, input)
			},
		},
		"update-channel-message": {
			Name:   "update-channel-message",
			Fields: fields_update_channel_message,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateChannelMessageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_channel_message, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateChannelMessage(ctx, input)
			},
		},
		"update-channel-read-marker": {
			Name:   "update-channel-read-marker",
			Fields: fields_update_channel_read_marker,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateChannelReadMarkerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_channel_read_marker, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateChannelReadMarker(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("chimesdkmessaging", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
