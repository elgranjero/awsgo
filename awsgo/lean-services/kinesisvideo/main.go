package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/kinesisvideo"
)

var fields_create_signaling_channel = []leanruntime.Field{
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
	{Name: "ChannelType", Flag: "channel-type", Type: "types.ChannelType", Required: false},
	{Name: "SingleMasterConfiguration", Flag: "single-master-configuration", Type: "*types.SingleMasterConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_stream = []leanruntime.Field{
	{Name: "DataRetentionInHours", Flag: "data-retention-in-hours", Type: "*int32", Required: false},
	{Name: "DeviceName", Flag: "device-name", Type: "*string", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "MediaType", Flag: "media-type", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: true},
	{Name: "StreamStorageConfiguration", Flag: "stream-storage-configuration", Type: "*types.StreamStorageConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_edge_configuration = []leanruntime.Field{
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
}

var fields_delete_signaling_channel = []leanruntime.Field{
	{Name: "ChannelARN", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "CurrentVersion", Flag: "current-version", Type: "*string", Required: false},
}

var fields_delete_stream = []leanruntime.Field{
	{Name: "CurrentVersion", Flag: "current-version", Type: "*string", Required: false},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: true},
}

var fields_describe_edge_configuration = []leanruntime.Field{
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
}

var fields_describe_image_generation_configuration = []leanruntime.Field{
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
}

var fields_describe_mapped_resource_configuration = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
}

var fields_describe_media_storage_configuration = []leanruntime.Field{
	{Name: "ChannelARN", Flag: "channel-arn", Type: "*string", Required: false},
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: false},
}

var fields_describe_notification_configuration = []leanruntime.Field{
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
}

var fields_describe_signaling_channel = []leanruntime.Field{
	{Name: "ChannelARN", Flag: "channel-arn", Type: "*string", Required: false},
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: false},
}

var fields_describe_stream = []leanruntime.Field{
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
}

var fields_describe_stream_storage_configuration = []leanruntime.Field{
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
}

var fields_get_data_endpoint = []leanruntime.Field{
	{Name: "APIName", Flag: "api-name", Type: "types.APIName", Required: true},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
}

var fields_get_signaling_channel_endpoint = []leanruntime.Field{
	{Name: "ChannelARN", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "SingleMasterChannelEndpointConfiguration", Flag: "single-master-channel-endpoint-configuration", Type: "*types.SingleMasterChannelEndpointConfiguration", Required: false},
}

var fields_list_edge_agent_configurations = []leanruntime.Field{
	{Name: "HubDeviceArn", Flag: "hub-device-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_signaling_channels = []leanruntime.Field{
	{Name: "ChannelNameCondition", Flag: "channel-name-condition", Type: "*types.ChannelNameCondition", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_streams = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StreamNameCondition", Flag: "stream-name-condition", Type: "*types.StreamNameCondition", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_tags_for_stream = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
}

var fields_start_edge_configuration_update = []leanruntime.Field{
	{Name: "EdgeConfig", Flag: "edge-config", Type: "*types.EdgeConfig", Required: true},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_tag_stream = []leanruntime.Field{
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeyList", Flag: "tag-key-list", Type: "[]string", Required: true},
}

var fields_untag_stream = []leanruntime.Field{
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
	{Name: "TagKeyList", Flag: "tag-key-list", Type: "[]string", Required: true},
}

var fields_update_data_retention = []leanruntime.Field{
	{Name: "CurrentVersion", Flag: "current-version", Type: "*string", Required: true},
	{Name: "DataRetentionChangeInHours", Flag: "data-retention-change-in-hours", Type: "*int32", Required: true},
	{Name: "Operation", Flag: "operation", Type: "types.UpdateDataRetentionOperation", Required: true},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
}

var fields_update_image_generation_configuration = []leanruntime.Field{
	{Name: "ImageGenerationConfiguration", Flag: "image-generation-configuration", Type: "*types.ImageGenerationConfiguration", Required: false},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
}

var fields_update_media_storage_configuration = []leanruntime.Field{
	{Name: "ChannelARN", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "MediaStorageConfiguration", Flag: "media-storage-configuration", Type: "*types.MediaStorageConfiguration", Required: true},
}

var fields_update_notification_configuration = []leanruntime.Field{
	{Name: "NotificationConfiguration", Flag: "notification-configuration", Type: "*types.NotificationConfiguration", Required: false},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
}

var fields_update_signaling_channel = []leanruntime.Field{
	{Name: "ChannelARN", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "CurrentVersion", Flag: "current-version", Type: "*string", Required: true},
	{Name: "SingleMasterConfiguration", Flag: "single-master-configuration", Type: "*types.SingleMasterConfiguration", Required: false},
}

var fields_update_stream = []leanruntime.Field{
	{Name: "CurrentVersion", Flag: "current-version", Type: "*string", Required: true},
	{Name: "DeviceName", Flag: "device-name", Type: "*string", Required: false},
	{Name: "MediaType", Flag: "media-type", Type: "*string", Required: false},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
}

var fields_update_stream_storage_configuration = []leanruntime.Field{
	{Name: "CurrentVersion", Flag: "current-version", Type: "*string", Required: true},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
	{Name: "StreamStorageConfiguration", Flag: "stream-storage-configuration", Type: "*types.StreamStorageConfiguration", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-signaling-channel": {
			Name:   "create-signaling-channel",
			Fields: fields_create_signaling_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSignalingChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_signaling_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSignalingChannel(ctx, input)
			},
		},
		"create-stream": {
			Name:   "create-stream",
			Fields: fields_create_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateStream(ctx, input)
			},
		},
		"delete-edge-configuration": {
			Name:   "delete-edge-configuration",
			Fields: fields_delete_edge_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEdgeConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_edge_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEdgeConfiguration(ctx, input)
			},
		},
		"delete-signaling-channel": {
			Name:   "delete-signaling-channel",
			Fields: fields_delete_signaling_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSignalingChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_signaling_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSignalingChannel(ctx, input)
			},
		},
		"delete-stream": {
			Name:   "delete-stream",
			Fields: fields_delete_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteStream(ctx, input)
			},
		},
		"describe-edge-configuration": {
			Name:   "describe-edge-configuration",
			Fields: fields_describe_edge_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEdgeConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_edge_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEdgeConfiguration(ctx, input)
			},
		},
		"describe-image-generation-configuration": {
			Name:   "describe-image-generation-configuration",
			Fields: fields_describe_image_generation_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeImageGenerationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_image_generation_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeImageGenerationConfiguration(ctx, input)
			},
		},
		"describe-mapped-resource-configuration": {
			Name:   "describe-mapped-resource-configuration",
			Fields: fields_describe_mapped_resource_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMappedResourceConfigurationInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_mapped_resource_configuration, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeMappedResourceConfiguration(ctx, input)
				}
				var results []*svc.DescribeMappedResourceConfigurationOutput
				p := svc.NewDescribeMappedResourceConfigurationPaginator(client, input)
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
		"describe-media-storage-configuration": {
			Name:   "describe-media-storage-configuration",
			Fields: fields_describe_media_storage_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMediaStorageConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_media_storage_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeMediaStorageConfiguration(ctx, input)
			},
		},
		"describe-notification-configuration": {
			Name:   "describe-notification-configuration",
			Fields: fields_describe_notification_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeNotificationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_notification_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeNotificationConfiguration(ctx, input)
			},
		},
		"describe-signaling-channel": {
			Name:   "describe-signaling-channel",
			Fields: fields_describe_signaling_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSignalingChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_signaling_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSignalingChannel(ctx, input)
			},
		},
		"describe-stream": {
			Name:   "describe-stream",
			Fields: fields_describe_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeStream(ctx, input)
			},
		},
		"describe-stream-storage-configuration": {
			Name:   "describe-stream-storage-configuration",
			Fields: fields_describe_stream_storage_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeStreamStorageConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_stream_storage_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeStreamStorageConfiguration(ctx, input)
			},
		},
		"get-data-endpoint": {
			Name:   "get-data-endpoint",
			Fields: fields_get_data_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataEndpoint(ctx, input)
			},
		},
		"get-signaling-channel-endpoint": {
			Name:   "get-signaling-channel-endpoint",
			Fields: fields_get_signaling_channel_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSignalingChannelEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_signaling_channel_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSignalingChannelEndpoint(ctx, input)
			},
		},
		"list-edge-agent-configurations": {
			Name:   "list-edge-agent-configurations",
			Fields: fields_list_edge_agent_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEdgeAgentConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_edge_agent_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEdgeAgentConfigurations(ctx, input)
				}
				var results []*svc.ListEdgeAgentConfigurationsOutput
				p := svc.NewListEdgeAgentConfigurationsPaginator(client, input)
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
		"list-signaling-channels": {
			Name:   "list-signaling-channels",
			Fields: fields_list_signaling_channels,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSignalingChannelsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_signaling_channels, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSignalingChannels(ctx, input)
				}
				var results []*svc.ListSignalingChannelsOutput
				p := svc.NewListSignalingChannelsPaginator(client, input)
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
		"list-streams": {
			Name:   "list-streams",
			Fields: fields_list_streams,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStreamsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_streams, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStreams(ctx, input)
				}
				var results []*svc.ListStreamsOutput
				p := svc.NewListStreamsPaginator(client, input)
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
		"list-tags-for-stream": {
			Name:   "list-tags-for-stream",
			Fields: fields_list_tags_for_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsForStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_tags_for_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTagsForStream(ctx, input)
			},
		},
		"start-edge-configuration-update": {
			Name:   "start-edge-configuration-update",
			Fields: fields_start_edge_configuration_update,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartEdgeConfigurationUpdateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_edge_configuration_update, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartEdgeConfigurationUpdate(ctx, input)
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
		"tag-stream": {
			Name:   "tag-stream",
			Fields: fields_tag_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TagStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_tag_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TagStream(ctx, input)
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
		"untag-stream": {
			Name:   "untag-stream",
			Fields: fields_untag_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UntagStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_untag_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UntagStream(ctx, input)
			},
		},
		"update-data-retention": {
			Name:   "update-data-retention",
			Fields: fields_update_data_retention,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDataRetentionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_data_retention, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDataRetention(ctx, input)
			},
		},
		"update-image-generation-configuration": {
			Name:   "update-image-generation-configuration",
			Fields: fields_update_image_generation_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateImageGenerationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_image_generation_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateImageGenerationConfiguration(ctx, input)
			},
		},
		"update-media-storage-configuration": {
			Name:   "update-media-storage-configuration",
			Fields: fields_update_media_storage_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMediaStorageConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_media_storage_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMediaStorageConfiguration(ctx, input)
			},
		},
		"update-notification-configuration": {
			Name:   "update-notification-configuration",
			Fields: fields_update_notification_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateNotificationConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_notification_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateNotificationConfiguration(ctx, input)
			},
		},
		"update-signaling-channel": {
			Name:   "update-signaling-channel",
			Fields: fields_update_signaling_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSignalingChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_signaling_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSignalingChannel(ctx, input)
			},
		},
		"update-stream": {
			Name:   "update-stream",
			Fields: fields_update_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateStream(ctx, input)
			},
		},
		"update-stream-storage-configuration": {
			Name:   "update-stream-storage-configuration",
			Fields: fields_update_stream_storage_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateStreamStorageConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_stream_storage_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateStreamStorageConfiguration(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("kinesisvideo", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
