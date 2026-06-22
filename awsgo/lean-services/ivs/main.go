package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/ivs"
)

var fields_batch_get_channel = []leanruntime.Field{
	{Name: "Arns", Flag: "arns", Type: "[]string", Required: true},
}

var fields_batch_get_stream_key = []leanruntime.Field{
	{Name: "Arns", Flag: "arns", Type: "[]string", Required: true},
}

var fields_batch_start_viewer_session_revocation = []leanruntime.Field{
	{Name: "ViewerSessions", Flag: "viewer-sessions", Type: "[]types.BatchStartViewerSessionRevocationViewerSession", Required: true},
}

var fields_create_channel = []leanruntime.Field{
	{Name: "Authorized", Flag: "authorized", Type: "bool", Required: false},
	{Name: "ContainerFormat", Flag: "container-format", Type: "types.ContainerFormat", Required: false},
	{Name: "InsecureIngest", Flag: "insecure-ingest", Type: "bool", Required: false},
	{Name: "LatencyMode", Flag: "latency-mode", Type: "types.ChannelLatencyMode", Required: false},
	{Name: "MultitrackInputConfiguration", Flag: "multitrack-input-configuration", Type: "*types.MultitrackInputConfiguration", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "PlaybackRestrictionPolicyArn", Flag: "playback-restriction-policy-arn", Type: "*string", Required: false},
	{Name: "Preset", Flag: "preset", Type: "types.TranscodePreset", Required: false},
	{Name: "RecordingConfigurationArn", Flag: "recording-configuration-arn", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.ChannelType", Required: false},
}

var fields_create_playback_restriction_policy = []leanruntime.Field{
	{Name: "AllowedCountries", Flag: "allowed-countries", Type: "[]string", Required: false},
	{Name: "AllowedOrigins", Flag: "allowed-origins", Type: "[]string", Required: false},
	{Name: "EnableStrictOriginEnforcement", Flag: "enable-strict-origin-enforcement", Type: "*bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_recording_configuration = []leanruntime.Field{
	{Name: "DestinationConfiguration", Flag: "destination-configuration", Type: "*types.DestinationConfiguration", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "RecordingReconnectWindowSeconds", Flag: "recording-reconnect-window-seconds", Type: "int32", Required: false},
	{Name: "RenditionConfiguration", Flag: "rendition-configuration", Type: "*types.RenditionConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "ThumbnailConfiguration", Flag: "thumbnail-configuration", Type: "*types.ThumbnailConfiguration", Required: false},
}

var fields_create_stream_key = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_channel = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_delete_playback_key_pair = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_delete_playback_restriction_policy = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_delete_recording_configuration = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_delete_stream_key = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_channel = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_playback_key_pair = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_playback_restriction_policy = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_recording_configuration = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_stream = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
}

var fields_get_stream_key = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_stream_session = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "StreamId", Flag: "stream-id", Type: "*string", Required: false},
}

var fields_import_playback_key_pair = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "PublicKeyMaterial", Flag: "public-key-material", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_list_channels = []leanruntime.Field{
	{Name: "FilterByName", Flag: "filter-by-name", Type: "*string", Required: false},
	{Name: "FilterByPlaybackRestrictionPolicyArn", Flag: "filter-by-playback-restriction-policy-arn", Type: "*string", Required: false},
	{Name: "FilterByRecordingConfigurationArn", Flag: "filter-by-recording-configuration-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_playback_key_pairs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_playback_restriction_policies = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_recording_configurations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_stream_keys = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_stream_sessions = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_streams = []leanruntime.Field{
	{Name: "FilterBy", Flag: "filter-by", Type: "*types.StreamFilters", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_metadata = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "Metadata", Flag: "metadata", Type: "*string", Required: true},
}

var fields_start_viewer_session_revocation = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
	{Name: "ViewerId", Flag: "viewer-id", Type: "*string", Required: true},
	{Name: "ViewerSessionVersionsLessThanOrEqualTo", Flag: "viewer-session-versions-less-than-or-equal-to", Type: "int32", Required: false},
}

var fields_stop_stream = []leanruntime.Field{
	{Name: "ChannelArn", Flag: "channel-arn", Type: "*string", Required: true},
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
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Authorized", Flag: "authorized", Type: "bool", Required: false},
	{Name: "ContainerFormat", Flag: "container-format", Type: "types.ContainerFormat", Required: false},
	{Name: "InsecureIngest", Flag: "insecure-ingest", Type: "bool", Required: false},
	{Name: "LatencyMode", Flag: "latency-mode", Type: "types.ChannelLatencyMode", Required: false},
	{Name: "MultitrackInputConfiguration", Flag: "multitrack-input-configuration", Type: "*types.MultitrackInputConfiguration", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "PlaybackRestrictionPolicyArn", Flag: "playback-restriction-policy-arn", Type: "*string", Required: false},
	{Name: "Preset", Flag: "preset", Type: "types.TranscodePreset", Required: false},
	{Name: "RecordingConfigurationArn", Flag: "recording-configuration-arn", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.ChannelType", Required: false},
}

var fields_update_playback_restriction_policy = []leanruntime.Field{
	{Name: "AllowedCountries", Flag: "allowed-countries", Type: "[]string", Required: false},
	{Name: "AllowedOrigins", Flag: "allowed-origins", Type: "[]string", Required: false},
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "EnableStrictOriginEnforcement", Flag: "enable-strict-origin-enforcement", Type: "*bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-get-channel": {
			Name:   "batch-get-channel",
			Fields: fields_batch_get_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetChannel(ctx, input)
			},
		},
		"batch-get-stream-key": {
			Name:   "batch-get-stream-key",
			Fields: fields_batch_get_stream_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetStreamKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_stream_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetStreamKey(ctx, input)
			},
		},
		"batch-start-viewer-session-revocation": {
			Name:   "batch-start-viewer-session-revocation",
			Fields: fields_batch_start_viewer_session_revocation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchStartViewerSessionRevocationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_start_viewer_session_revocation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchStartViewerSessionRevocation(ctx, input)
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
		"create-playback-restriction-policy": {
			Name:   "create-playback-restriction-policy",
			Fields: fields_create_playback_restriction_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePlaybackRestrictionPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_playback_restriction_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePlaybackRestrictionPolicy(ctx, input)
			},
		},
		"create-recording-configuration": {
			Name:   "create-recording-configuration",
			Fields: fields_create_recording_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRecordingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_recording_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRecordingConfiguration(ctx, input)
			},
		},
		"create-stream-key": {
			Name:   "create-stream-key",
			Fields: fields_create_stream_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateStreamKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_stream_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateStreamKey(ctx, input)
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
		"delete-playback-key-pair": {
			Name:   "delete-playback-key-pair",
			Fields: fields_delete_playback_key_pair,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePlaybackKeyPairInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_playback_key_pair, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePlaybackKeyPair(ctx, input)
			},
		},
		"delete-playback-restriction-policy": {
			Name:   "delete-playback-restriction-policy",
			Fields: fields_delete_playback_restriction_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePlaybackRestrictionPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_playback_restriction_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePlaybackRestrictionPolicy(ctx, input)
			},
		},
		"delete-recording-configuration": {
			Name:   "delete-recording-configuration",
			Fields: fields_delete_recording_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRecordingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_recording_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRecordingConfiguration(ctx, input)
			},
		},
		"delete-stream-key": {
			Name:   "delete-stream-key",
			Fields: fields_delete_stream_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteStreamKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_stream_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteStreamKey(ctx, input)
			},
		},
		"get-channel": {
			Name:   "get-channel",
			Fields: fields_get_channel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetChannelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_channel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetChannel(ctx, input)
			},
		},
		"get-playback-key-pair": {
			Name:   "get-playback-key-pair",
			Fields: fields_get_playback_key_pair,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPlaybackKeyPairInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_playback_key_pair, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPlaybackKeyPair(ctx, input)
			},
		},
		"get-playback-restriction-policy": {
			Name:   "get-playback-restriction-policy",
			Fields: fields_get_playback_restriction_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPlaybackRestrictionPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_playback_restriction_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPlaybackRestrictionPolicy(ctx, input)
			},
		},
		"get-recording-configuration": {
			Name:   "get-recording-configuration",
			Fields: fields_get_recording_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRecordingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_recording_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRecordingConfiguration(ctx, input)
			},
		},
		"get-stream": {
			Name:   "get-stream",
			Fields: fields_get_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetStream(ctx, input)
			},
		},
		"get-stream-key": {
			Name:   "get-stream-key",
			Fields: fields_get_stream_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetStreamKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_stream_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetStreamKey(ctx, input)
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
		"import-playback-key-pair": {
			Name:   "import-playback-key-pair",
			Fields: fields_import_playback_key_pair,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportPlaybackKeyPairInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_playback_key_pair, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportPlaybackKeyPair(ctx, input)
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
		"list-playback-key-pairs": {
			Name:   "list-playback-key-pairs",
			Fields: fields_list_playback_key_pairs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPlaybackKeyPairsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_playback_key_pairs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPlaybackKeyPairs(ctx, input)
				}
				var results []*svc.ListPlaybackKeyPairsOutput
				p := svc.NewListPlaybackKeyPairsPaginator(client, input)
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
		"list-playback-restriction-policies": {
			Name:   "list-playback-restriction-policies",
			Fields: fields_list_playback_restriction_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPlaybackRestrictionPoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_playback_restriction_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPlaybackRestrictionPolicies(ctx, input)
				}
				var results []*svc.ListPlaybackRestrictionPoliciesOutput
				p := svc.NewListPlaybackRestrictionPoliciesPaginator(client, input)
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
		"list-recording-configurations": {
			Name:   "list-recording-configurations",
			Fields: fields_list_recording_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRecordingConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_recording_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRecordingConfigurations(ctx, input)
				}
				var results []*svc.ListRecordingConfigurationsOutput
				p := svc.NewListRecordingConfigurationsPaginator(client, input)
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
		"list-stream-keys": {
			Name:   "list-stream-keys",
			Fields: fields_list_stream_keys,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStreamKeysInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_stream_keys, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStreamKeys(ctx, input)
				}
				var results []*svc.ListStreamKeysOutput
				p := svc.NewListStreamKeysPaginator(client, input)
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
		"put-metadata": {
			Name:   "put-metadata",
			Fields: fields_put_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutMetadata(ctx, input)
			},
		},
		"start-viewer-session-revocation": {
			Name:   "start-viewer-session-revocation",
			Fields: fields_start_viewer_session_revocation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartViewerSessionRevocationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_viewer_session_revocation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartViewerSessionRevocation(ctx, input)
			},
		},
		"stop-stream": {
			Name:   "stop-stream",
			Fields: fields_stop_stream,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopStreamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_stream, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopStream(ctx, input)
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
		"update-playback-restriction-policy": {
			Name:   "update-playback-restriction-policy",
			Fields: fields_update_playback_restriction_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePlaybackRestrictionPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_playback_restriction_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePlaybackRestrictionPolicy(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("ivs", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
