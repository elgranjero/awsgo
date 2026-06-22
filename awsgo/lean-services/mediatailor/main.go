package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/mediatailor"
)

var fields_configure_logs_for_channel = []leanruntime.Field{
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
	{Name: "LogTypes", Flag: "log-types", Type: "[]types.LogType", Required: true},
}

var fields_configure_logs_for_playback_configuration = []leanruntime.Field{
	{Name: "AdsInteractionLog", Flag: "ads-interaction-log", Type: "*types.AdsInteractionLog", Required: false},
	{Name: "EnabledLoggingStrategies", Flag: "enabled-logging-strategies", Type: "[]types.LoggingStrategy", Required: false},
	{Name: "ManifestServiceInteractionLog", Flag: "manifest-service-interaction-log", Type: "*types.ManifestServiceInteractionLog", Required: false},
	{Name: "PercentEnabled", Flag: "percent-enabled", Type: "int32", Required: true},
	{Name: "PlaybackConfigurationName", Flag: "playback-configuration-name", Type: "*string", Required: true},
}

var fields_create_channel = []leanruntime.Field{
	{Name: "Audiences", Flag: "audiences", Type: "[]string", Required: false},
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
	{Name: "FillerSlate", Flag: "filler-slate", Type: "*types.SlateSource", Required: false},
	{Name: "Outputs", Flag: "outputs", Type: "[]types.RequestOutputItem", Required: true},
	{Name: "PlaybackMode", Flag: "playback-mode", Type: "types.PlaybackMode", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Tier", Flag: "tier", Type: "types.Tier", Required: false},
	{Name: "TimeShiftConfiguration", Flag: "time-shift-configuration", Type: "*types.TimeShiftConfiguration", Required: false},
}

var fields_create_live_source = []leanruntime.Field{
	{Name: "HttpPackageConfigurations", Flag: "http-package-configurations", Type: "[]types.HttpPackageConfiguration", Required: true},
	{Name: "LiveSourceName", Flag: "live-source-name", Type: "*string", Required: true},
	{Name: "SourceLocationName", Flag: "source-location-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_prefetch_schedule = []leanruntime.Field{
	{Name: "Consumption", Flag: "consumption", Type: "*types.PrefetchConsumption", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PlaybackConfigurationName", Flag: "playback-configuration-name", Type: "*string", Required: true},
	{Name: "RecurringPrefetchConfiguration", Flag: "recurring-prefetch-configuration", Type: "*types.RecurringPrefetchConfiguration", Required: false},
	{Name: "Retrieval", Flag: "retrieval", Type: "*types.PrefetchRetrieval", Required: false},
	{Name: "ScheduleType", Flag: "schedule-type", Type: "types.PrefetchScheduleType", Required: false},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
}

var fields_create_program = []leanruntime.Field{
	{Name: "AdBreaks", Flag: "ad-breaks", Type: "[]types.AdBreak", Required: false},
	{Name: "AudienceMedia", Flag: "audience-media", Type: "[]types.AudienceMedia", Required: false},
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
	{Name: "LiveSourceName", Flag: "live-source-name", Type: "*string", Required: false},
	{Name: "ProgramName", Flag: "program-name", Type: "*string", Required: true},
	{Name: "ScheduleConfiguration", Flag: "schedule-configuration", Type: "*types.ScheduleConfiguration", Required: true},
	{Name: "SourceLocationName", Flag: "source-location-name", Type: "*string", Required: true},
	{Name: "VodSourceName", Flag: "vod-source-name", Type: "*string", Required: false},
}

var fields_create_source_location = []leanruntime.Field{
	{Name: "AccessConfiguration", Flag: "access-configuration", Type: "*types.AccessConfiguration", Required: false},
	{Name: "DefaultSegmentDeliveryConfiguration", Flag: "default-segment-delivery-configuration", Type: "*types.DefaultSegmentDeliveryConfiguration", Required: false},
	{Name: "HttpConfiguration", Flag: "http-configuration", Type: "*types.HttpConfiguration", Required: true},
	{Name: "SegmentDeliveryConfigurations", Flag: "segment-delivery-configurations", Type: "[]types.SegmentDeliveryConfiguration", Required: false},
	{Name: "SourceLocationName", Flag: "source-location-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_vod_source = []leanruntime.Field{
	{Name: "HttpPackageConfigurations", Flag: "http-package-configurations", Type: "[]types.HttpPackageConfiguration", Required: true},
	{Name: "SourceLocationName", Flag: "source-location-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "VodSourceName", Flag: "vod-source-name", Type: "*string", Required: true},
}

var fields_delete_channel = []leanruntime.Field{
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
}

var fields_delete_channel_policy = []leanruntime.Field{
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
}

var fields_delete_live_source = []leanruntime.Field{
	{Name: "LiveSourceName", Flag: "live-source-name", Type: "*string", Required: true},
	{Name: "SourceLocationName", Flag: "source-location-name", Type: "*string", Required: true},
}

var fields_delete_playback_configuration = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_prefetch_schedule = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PlaybackConfigurationName", Flag: "playback-configuration-name", Type: "*string", Required: true},
}

var fields_delete_program = []leanruntime.Field{
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
	{Name: "ProgramName", Flag: "program-name", Type: "*string", Required: true},
}

var fields_delete_source_location = []leanruntime.Field{
	{Name: "SourceLocationName", Flag: "source-location-name", Type: "*string", Required: true},
}

var fields_delete_vod_source = []leanruntime.Field{
	{Name: "SourceLocationName", Flag: "source-location-name", Type: "*string", Required: true},
	{Name: "VodSourceName", Flag: "vod-source-name", Type: "*string", Required: true},
}

var fields_describe_channel = []leanruntime.Field{
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
}

var fields_describe_live_source = []leanruntime.Field{
	{Name: "LiveSourceName", Flag: "live-source-name", Type: "*string", Required: true},
	{Name: "SourceLocationName", Flag: "source-location-name", Type: "*string", Required: true},
}

var fields_describe_program = []leanruntime.Field{
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
	{Name: "ProgramName", Flag: "program-name", Type: "*string", Required: true},
}

var fields_describe_source_location = []leanruntime.Field{
	{Name: "SourceLocationName", Flag: "source-location-name", Type: "*string", Required: true},
}

var fields_describe_vod_source = []leanruntime.Field{
	{Name: "SourceLocationName", Flag: "source-location-name", Type: "*string", Required: true},
	{Name: "VodSourceName", Flag: "vod-source-name", Type: "*string", Required: true},
}

var fields_get_channel_policy = []leanruntime.Field{
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
}

var fields_get_channel_schedule = []leanruntime.Field{
	{Name: "Audience", Flag: "audience", Type: "*string", Required: false},
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
	{Name: "DurationMinutes", Flag: "duration-minutes", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_playback_configuration = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_prefetch_schedule = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PlaybackConfigurationName", Flag: "playback-configuration-name", Type: "*string", Required: true},
}

var fields_list_alerts = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_channels = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_live_sources = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SourceLocationName", Flag: "source-location-name", Type: "*string", Required: true},
}

var fields_list_playback_configurations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_prefetch_schedules = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PlaybackConfigurationName", Flag: "playback-configuration-name", Type: "*string", Required: true},
	{Name: "ScheduleType", Flag: "schedule-type", Type: "types.ListPrefetchScheduleType", Required: false},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
}

var fields_list_source_locations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_vod_sources = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SourceLocationName", Flag: "source-location-name", Type: "*string", Required: true},
}

var fields_put_channel_policy = []leanruntime.Field{
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
}

var fields_put_playback_configuration = []leanruntime.Field{
	{Name: "AdConditioningConfiguration", Flag: "ad-conditioning-configuration", Type: "*types.AdConditioningConfiguration", Required: false},
	{Name: "AdDecisionServerConfiguration", Flag: "ad-decision-server-configuration", Type: "*types.AdDecisionServerConfiguration", Required: false},
	{Name: "AdDecisionServerUrl", Flag: "ad-decision-server-url", Type: "*string", Required: false},
	{Name: "AvailSuppression", Flag: "avail-suppression", Type: "*types.AvailSuppression", Required: false},
	{Name: "Bumper", Flag: "bumper", Type: "*types.Bumper", Required: false},
	{Name: "CdnConfiguration", Flag: "cdn-configuration", Type: "*types.CdnConfiguration", Required: false},
	{Name: "ConfigurationAliases", Flag: "configuration-aliases", Type: "map[string]map[string]string", Required: false},
	{Name: "DashConfiguration", Flag: "dash-configuration", Type: "*types.DashConfigurationForPut", Required: false},
	{Name: "InsertionMode", Flag: "insertion-mode", Type: "types.InsertionMode", Required: false},
	{Name: "LivePreRollConfiguration", Flag: "live-pre-roll-configuration", Type: "*types.LivePreRollConfiguration", Required: false},
	{Name: "ManifestProcessingRules", Flag: "manifest-processing-rules", Type: "*types.ManifestProcessingRules", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PersonalizationThresholdSeconds", Flag: "personalization-threshold-seconds", Type: "*int32", Required: false},
	{Name: "SlateAdUrl", Flag: "slate-ad-url", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TranscodeProfileName", Flag: "transcode-profile-name", Type: "*string", Required: false},
	{Name: "VideoContentSourceUrl", Flag: "video-content-source-url", Type: "*string", Required: false},
}

var fields_start_channel = []leanruntime.Field{
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
}

var fields_stop_channel = []leanruntime.Field{
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_channel = []leanruntime.Field{
	{Name: "Audiences", Flag: "audiences", Type: "[]string", Required: false},
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
	{Name: "FillerSlate", Flag: "filler-slate", Type: "*types.SlateSource", Required: false},
	{Name: "Outputs", Flag: "outputs", Type: "[]types.RequestOutputItem", Required: true},
	{Name: "TimeShiftConfiguration", Flag: "time-shift-configuration", Type: "*types.TimeShiftConfiguration", Required: false},
}

var fields_update_live_source = []leanruntime.Field{
	{Name: "HttpPackageConfigurations", Flag: "http-package-configurations", Type: "[]types.HttpPackageConfiguration", Required: true},
	{Name: "LiveSourceName", Flag: "live-source-name", Type: "*string", Required: true},
	{Name: "SourceLocationName", Flag: "source-location-name", Type: "*string", Required: true},
}

var fields_update_program = []leanruntime.Field{
	{Name: "AdBreaks", Flag: "ad-breaks", Type: "[]types.AdBreak", Required: false},
	{Name: "AudienceMedia", Flag: "audience-media", Type: "[]types.AudienceMedia", Required: false},
	{Name: "ChannelName", Flag: "channel-name", Type: "*string", Required: true},
	{Name: "ProgramName", Flag: "program-name", Type: "*string", Required: true},
	{Name: "ScheduleConfiguration", Flag: "schedule-configuration", Type: "*types.UpdateProgramScheduleConfiguration", Required: true},
}

var fields_update_source_location = []leanruntime.Field{
	{Name: "AccessConfiguration", Flag: "access-configuration", Type: "*types.AccessConfiguration", Required: false},
	{Name: "DefaultSegmentDeliveryConfiguration", Flag: "default-segment-delivery-configuration", Type: "*types.DefaultSegmentDeliveryConfiguration", Required: false},
	{Name: "HttpConfiguration", Flag: "http-configuration", Type: "*types.HttpConfiguration", Required: true},
	{Name: "SegmentDeliveryConfigurations", Flag: "segment-delivery-configurations", Type: "[]types.SegmentDeliveryConfiguration", Required: false},
	{Name: "SourceLocationName", Flag: "source-location-name", Type: "*string", Required: true},
}

var fields_update_vod_source = []leanruntime.Field{
	{Name: "HttpPackageConfigurations", Flag: "http-package-configurations", Type: "[]types.HttpPackageConfiguration", Required: true},
	{Name: "SourceLocationName", Flag: "source-location-name", Type: "*string", Required: true},
	{Name: "VodSourceName", Flag: "vod-source-name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"configure-logs-for-channel": {
			Name:   "configure-logs-for-channel",
			Fields: fields_configure_logs_for_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ConfigureLogsForChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_configure_logs_for_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ConfigureLogsForChannel(ctx, input)
			},
		},
		"configure-logs-for-playback-configuration": {
			Name:   "configure-logs-for-playback-configuration",
			Fields: fields_configure_logs_for_playback_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ConfigureLogsForPlaybackConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_configure_logs_for_playback_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ConfigureLogsForPlaybackConfiguration(ctx, input)
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
		"create-live-source": {
			Name:   "create-live-source",
			Fields: fields_create_live_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLiveSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_live_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLiveSource(ctx, input)
			},
		},
		"create-prefetch-schedule": {
			Name:   "create-prefetch-schedule",
			Fields: fields_create_prefetch_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePrefetchScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_prefetch_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePrefetchSchedule(ctx, input)
			},
		},
		"create-program": {
			Name:   "create-program",
			Fields: fields_create_program,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProgramInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_program, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProgram(ctx, input)
			},
		},
		"create-source-location": {
			Name:   "create-source-location",
			Fields: fields_create_source_location,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSourceLocationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_source_location, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSourceLocation(ctx, input)
			},
		},
		"create-vod-source": {
			Name:   "create-vod-source",
			Fields: fields_create_vod_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVodSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_vod_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVodSource(ctx, input)
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
		"delete-channel-policy": {
			Name:   "delete-channel-policy",
			Fields: fields_delete_channel_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteChannelPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_channel_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteChannelPolicy(ctx, input)
			},
		},
		"delete-live-source": {
			Name:   "delete-live-source",
			Fields: fields_delete_live_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLiveSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_live_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLiveSource(ctx, input)
			},
		},
		"delete-playback-configuration": {
			Name:   "delete-playback-configuration",
			Fields: fields_delete_playback_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePlaybackConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_playback_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePlaybackConfiguration(ctx, input)
			},
		},
		"delete-prefetch-schedule": {
			Name:   "delete-prefetch-schedule",
			Fields: fields_delete_prefetch_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePrefetchScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_prefetch_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePrefetchSchedule(ctx, input)
			},
		},
		"delete-program": {
			Name:   "delete-program",
			Fields: fields_delete_program,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProgramInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_program, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProgram(ctx, input)
			},
		},
		"delete-source-location": {
			Name:   "delete-source-location",
			Fields: fields_delete_source_location,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSourceLocationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_source_location, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSourceLocation(ctx, input)
			},
		},
		"delete-vod-source": {
			Name:   "delete-vod-source",
			Fields: fields_delete_vod_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVodSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vod_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVodSource(ctx, input)
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
		"describe-live-source": {
			Name:   "describe-live-source",
			Fields: fields_describe_live_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLiveSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_live_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLiveSource(ctx, input)
			},
		},
		"describe-program": {
			Name:   "describe-program",
			Fields: fields_describe_program,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeProgramInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_program, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeProgram(ctx, input)
			},
		},
		"describe-source-location": {
			Name:   "describe-source-location",
			Fields: fields_describe_source_location,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSourceLocationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_source_location, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSourceLocation(ctx, input)
			},
		},
		"describe-vod-source": {
			Name:   "describe-vod-source",
			Fields: fields_describe_vod_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVodSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_vod_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeVodSource(ctx, input)
			},
		},
		"get-channel-policy": {
			Name:   "get-channel-policy",
			Fields: fields_get_channel_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetChannelPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_channel_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetChannelPolicy(ctx, input)
			},
		},
		"get-channel-schedule": {
			Name:   "get-channel-schedule",
			Fields: fields_get_channel_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetChannelScheduleInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_channel_schedule, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetChannelSchedule(ctx, input)
				}
				var results []*svc.GetChannelScheduleOutput
				p := svc.NewGetChannelSchedulePaginator(client, input)
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
		"get-playback-configuration": {
			Name:   "get-playback-configuration",
			Fields: fields_get_playback_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPlaybackConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_playback_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPlaybackConfiguration(ctx, input)
			},
		},
		"get-prefetch-schedule": {
			Name:   "get-prefetch-schedule",
			Fields: fields_get_prefetch_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPrefetchScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_prefetch_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPrefetchSchedule(ctx, input)
			},
		},
		"list-alerts": {
			Name:   "list-alerts",
			Fields: fields_list_alerts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAlertsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_alerts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAlerts(ctx, input)
				}
				var results []*svc.ListAlertsOutput
				p := svc.NewListAlertsPaginator(client, input)
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
		"list-live-sources": {
			Name:   "list-live-sources",
			Fields: fields_list_live_sources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLiveSourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_live_sources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLiveSources(ctx, input)
				}
				var results []*svc.ListLiveSourcesOutput
				p := svc.NewListLiveSourcesPaginator(client, input)
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
		"list-playback-configurations": {
			Name:   "list-playback-configurations",
			Fields: fields_list_playback_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPlaybackConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_playback_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPlaybackConfigurations(ctx, input)
				}
				var results []*svc.ListPlaybackConfigurationsOutput
				p := svc.NewListPlaybackConfigurationsPaginator(client, input)
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
		"list-prefetch-schedules": {
			Name:   "list-prefetch-schedules",
			Fields: fields_list_prefetch_schedules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPrefetchSchedulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_prefetch_schedules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPrefetchSchedules(ctx, input)
				}
				var results []*svc.ListPrefetchSchedulesOutput
				p := svc.NewListPrefetchSchedulesPaginator(client, input)
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
		"list-source-locations": {
			Name:   "list-source-locations",
			Fields: fields_list_source_locations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSourceLocationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_source_locations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSourceLocations(ctx, input)
				}
				var results []*svc.ListSourceLocationsOutput
				p := svc.NewListSourceLocationsPaginator(client, input)
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
		"list-vod-sources": {
			Name:   "list-vod-sources",
			Fields: fields_list_vod_sources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVodSourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_vod_sources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListVodSources(ctx, input)
				}
				var results []*svc.ListVodSourcesOutput
				p := svc.NewListVodSourcesPaginator(client, input)
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
		"put-channel-policy": {
			Name:   "put-channel-policy",
			Fields: fields_put_channel_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutChannelPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_channel_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutChannelPolicy(ctx, input)
			},
		},
		"put-playback-configuration": {
			Name:   "put-playback-configuration",
			Fields: fields_put_playback_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutPlaybackConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_playback_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutPlaybackConfiguration(ctx, input)
			},
		},
		"start-channel": {
			Name:   "start-channel",
			Fields: fields_start_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartChannel(ctx, input)
			},
		},
		"stop-channel": {
			Name:   "stop-channel",
			Fields: fields_stop_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopChannel(ctx, input)
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
		"update-live-source": {
			Name:   "update-live-source",
			Fields: fields_update_live_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLiveSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_live_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLiveSource(ctx, input)
			},
		},
		"update-program": {
			Name:   "update-program",
			Fields: fields_update_program,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProgramInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_program, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProgram(ctx, input)
			},
		},
		"update-source-location": {
			Name:   "update-source-location",
			Fields: fields_update_source_location,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSourceLocationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_source_location, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSourceLocation(ctx, input)
			},
		},
		"update-vod-source": {
			Name:   "update-vod-source",
			Fields: fields_update_vod_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateVodSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_vod_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateVodSource(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("mediatailor", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
