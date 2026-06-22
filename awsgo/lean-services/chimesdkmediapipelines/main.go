package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/chimesdkmediapipelines"
)

var fields_create_media_capture_pipeline = []leanruntime.Field{
	{Name: "ChimeSdkMeetingConfiguration", Flag: "chime-sdk-meeting-configuration", Type: "*types.ChimeSdkMeetingConfiguration", Required: false},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "SinkArn", Flag: "sink-arn", Type: "*string", Required: true},
	{Name: "SinkIamRoleArn", Flag: "sink-iam-role-arn", Type: "*string", Required: false},
	{Name: "SinkType", Flag: "sink-type", Type: "types.MediaPipelineSinkType", Required: true},
	{Name: "SourceArn", Flag: "source-arn", Type: "*string", Required: true},
	{Name: "SourceType", Flag: "source-type", Type: "types.MediaPipelineSourceType", Required: true},
	{Name: "SseAwsKeyManagementParams", Flag: "sse-aws-key-management-params", Type: "*types.SseAwsKeyManagementParams", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_media_concatenation_pipeline = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Sinks", Flag: "sinks", Type: "[]types.ConcatenationSink", Required: true},
	{Name: "Sources", Flag: "sources", Type: "[]types.ConcatenationSource", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_media_insights_pipeline = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "KinesisVideoStreamRecordingSourceRuntimeConfiguration", Flag: "kinesis-video-stream-recording-source-runtime-configuration", Type: "*types.KinesisVideoStreamRecordingSourceRuntimeConfiguration", Required: false},
	{Name: "KinesisVideoStreamSourceRuntimeConfiguration", Flag: "kinesis-video-stream-source-runtime-configuration", Type: "*types.KinesisVideoStreamSourceRuntimeConfiguration", Required: false},
	{Name: "MediaInsightsPipelineConfigurationArn", Flag: "media-insights-pipeline-configuration-arn", Type: "*string", Required: true},
	{Name: "MediaInsightsRuntimeMetadata", Flag: "media-insights-runtime-metadata", Type: "map[string]string", Required: false},
	{Name: "S3RecordingSinkRuntimeConfiguration", Flag: "s3-recording-sink-runtime-configuration", Type: "*types.S3RecordingSinkRuntimeConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_media_insights_pipeline_configuration = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Elements", Flag: "elements", Type: "[]types.MediaInsightsPipelineConfigurationElement", Required: true},
	{Name: "MediaInsightsPipelineConfigurationName", Flag: "media-insights-pipeline-configuration-name", Type: "*string", Required: true},
	{Name: "RealTimeAlertConfiguration", Flag: "real-time-alert-configuration", Type: "*types.RealTimeAlertConfiguration", Required: false},
	{Name: "ResourceAccessRoleArn", Flag: "resource-access-role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_media_live_connector_pipeline = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Sinks", Flag: "sinks", Type: "[]types.LiveConnectorSinkConfiguration", Required: true},
	{Name: "Sources", Flag: "sources", Type: "[]types.LiveConnectorSourceConfiguration", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_media_pipeline_kinesis_video_stream_pool = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "PoolName", Flag: "pool-name", Type: "*string", Required: true},
	{Name: "StreamConfiguration", Flag: "stream-configuration", Type: "*types.KinesisVideoStreamConfiguration", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_media_stream_pipeline = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Sinks", Flag: "sinks", Type: "[]types.MediaStreamSink", Required: true},
	{Name: "Sources", Flag: "sources", Type: "[]types.MediaStreamSource", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_media_capture_pipeline = []leanruntime.Field{
	{Name: "MediaPipelineId", Flag: "media-pipeline-id", Type: "*string", Required: true},
}

var fields_delete_media_insights_pipeline_configuration = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_media_pipeline = []leanruntime.Field{
	{Name: "MediaPipelineId", Flag: "media-pipeline-id", Type: "*string", Required: true},
}

var fields_delete_media_pipeline_kinesis_video_stream_pool = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_media_capture_pipeline = []leanruntime.Field{
	{Name: "MediaPipelineId", Flag: "media-pipeline-id", Type: "*string", Required: true},
}

var fields_get_media_insights_pipeline_configuration = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_media_pipeline = []leanruntime.Field{
	{Name: "MediaPipelineId", Flag: "media-pipeline-id", Type: "*string", Required: true},
}

var fields_get_media_pipeline_kinesis_video_stream_pool = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_speaker_search_task = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "SpeakerSearchTaskId", Flag: "speaker-search-task-id", Type: "*string", Required: true},
}

var fields_get_voice_tone_analysis_task = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "VoiceToneAnalysisTaskId", Flag: "voice-tone-analysis-task-id", Type: "*string", Required: true},
}

var fields_list_media_capture_pipelines = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_media_insights_pipeline_configurations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_media_pipeline_kinesis_video_stream_pools = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_media_pipelines = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_start_speaker_search_task = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "KinesisVideoStreamSourceTaskConfiguration", Flag: "kinesis-video-stream-source-task-configuration", Type: "*types.KinesisVideoStreamSourceTaskConfiguration", Required: false},
	{Name: "VoiceProfileDomainArn", Flag: "voice-profile-domain-arn", Type: "*string", Required: true},
}

var fields_start_voice_tone_analysis_task = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "KinesisVideoStreamSourceTaskConfiguration", Flag: "kinesis-video-stream-source-task-configuration", Type: "*types.KinesisVideoStreamSourceTaskConfiguration", Required: false},
	{Name: "LanguageCode", Flag: "language-code", Type: "types.VoiceAnalyticsLanguageCode", Required: true},
}

var fields_stop_speaker_search_task = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "SpeakerSearchTaskId", Flag: "speaker-search-task-id", Type: "*string", Required: true},
}

var fields_stop_voice_tone_analysis_task = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "VoiceToneAnalysisTaskId", Flag: "voice-tone-analysis-task-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_media_insights_pipeline_configuration = []leanruntime.Field{
	{Name: "Elements", Flag: "elements", Type: "[]types.MediaInsightsPipelineConfigurationElement", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "RealTimeAlertConfiguration", Flag: "real-time-alert-configuration", Type: "*types.RealTimeAlertConfiguration", Required: false},
	{Name: "ResourceAccessRoleArn", Flag: "resource-access-role-arn", Type: "*string", Required: true},
}

var fields_update_media_insights_pipeline_status = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "UpdateStatus", Flag: "update-status", Type: "types.MediaPipelineStatusUpdate", Required: true},
}

var fields_update_media_pipeline_kinesis_video_stream_pool = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "StreamConfiguration", Flag: "stream-configuration", Type: "*types.KinesisVideoStreamConfigurationUpdate", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-media-capture-pipeline": {
			Name:   "create-media-capture-pipeline",
			Fields: fields_create_media_capture_pipeline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMediaCapturePipelineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_media_capture_pipeline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMediaCapturePipeline(ctx, input)
			},
		},
		"create-media-concatenation-pipeline": {
			Name:   "create-media-concatenation-pipeline",
			Fields: fields_create_media_concatenation_pipeline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMediaConcatenationPipelineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_media_concatenation_pipeline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMediaConcatenationPipeline(ctx, input)
			},
		},
		"create-media-insights-pipeline": {
			Name:   "create-media-insights-pipeline",
			Fields: fields_create_media_insights_pipeline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMediaInsightsPipelineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_media_insights_pipeline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMediaInsightsPipeline(ctx, input)
			},
		},
		"create-media-insights-pipeline-configuration": {
			Name:   "create-media-insights-pipeline-configuration",
			Fields: fields_create_media_insights_pipeline_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMediaInsightsPipelineConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_media_insights_pipeline_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMediaInsightsPipelineConfiguration(ctx, input)
			},
		},
		"create-media-live-connector-pipeline": {
			Name:   "create-media-live-connector-pipeline",
			Fields: fields_create_media_live_connector_pipeline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMediaLiveConnectorPipelineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_media_live_connector_pipeline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMediaLiveConnectorPipeline(ctx, input)
			},
		},
		"create-media-pipeline-kinesis-video-stream-pool": {
			Name:   "create-media-pipeline-kinesis-video-stream-pool",
			Fields: fields_create_media_pipeline_kinesis_video_stream_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMediaPipelineKinesisVideoStreamPoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_media_pipeline_kinesis_video_stream_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMediaPipelineKinesisVideoStreamPool(ctx, input)
			},
		},
		"create-media-stream-pipeline": {
			Name:   "create-media-stream-pipeline",
			Fields: fields_create_media_stream_pipeline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMediaStreamPipelineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_media_stream_pipeline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMediaStreamPipeline(ctx, input)
			},
		},
		"delete-media-capture-pipeline": {
			Name:   "delete-media-capture-pipeline",
			Fields: fields_delete_media_capture_pipeline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMediaCapturePipelineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_media_capture_pipeline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMediaCapturePipeline(ctx, input)
			},
		},
		"delete-media-insights-pipeline-configuration": {
			Name:   "delete-media-insights-pipeline-configuration",
			Fields: fields_delete_media_insights_pipeline_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMediaInsightsPipelineConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_media_insights_pipeline_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMediaInsightsPipelineConfiguration(ctx, input)
			},
		},
		"delete-media-pipeline": {
			Name:   "delete-media-pipeline",
			Fields: fields_delete_media_pipeline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMediaPipelineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_media_pipeline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMediaPipeline(ctx, input)
			},
		},
		"delete-media-pipeline-kinesis-video-stream-pool": {
			Name:   "delete-media-pipeline-kinesis-video-stream-pool",
			Fields: fields_delete_media_pipeline_kinesis_video_stream_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMediaPipelineKinesisVideoStreamPoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_media_pipeline_kinesis_video_stream_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMediaPipelineKinesisVideoStreamPool(ctx, input)
			},
		},
		"get-media-capture-pipeline": {
			Name:   "get-media-capture-pipeline",
			Fields: fields_get_media_capture_pipeline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMediaCapturePipelineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_media_capture_pipeline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMediaCapturePipeline(ctx, input)
			},
		},
		"get-media-insights-pipeline-configuration": {
			Name:   "get-media-insights-pipeline-configuration",
			Fields: fields_get_media_insights_pipeline_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMediaInsightsPipelineConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_media_insights_pipeline_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMediaInsightsPipelineConfiguration(ctx, input)
			},
		},
		"get-media-pipeline": {
			Name:   "get-media-pipeline",
			Fields: fields_get_media_pipeline,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMediaPipelineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_media_pipeline, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMediaPipeline(ctx, input)
			},
		},
		"get-media-pipeline-kinesis-video-stream-pool": {
			Name:   "get-media-pipeline-kinesis-video-stream-pool",
			Fields: fields_get_media_pipeline_kinesis_video_stream_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMediaPipelineKinesisVideoStreamPoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_media_pipeline_kinesis_video_stream_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMediaPipelineKinesisVideoStreamPool(ctx, input)
			},
		},
		"get-speaker-search-task": {
			Name:   "get-speaker-search-task",
			Fields: fields_get_speaker_search_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSpeakerSearchTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_speaker_search_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSpeakerSearchTask(ctx, input)
			},
		},
		"get-voice-tone-analysis-task": {
			Name:   "get-voice-tone-analysis-task",
			Fields: fields_get_voice_tone_analysis_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVoiceToneAnalysisTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_voice_tone_analysis_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVoiceToneAnalysisTask(ctx, input)
			},
		},
		"list-media-capture-pipelines": {
			Name:   "list-media-capture-pipelines",
			Fields: fields_list_media_capture_pipelines,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMediaCapturePipelinesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_media_capture_pipelines, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMediaCapturePipelines(ctx, input)
				}
				var results []*svc.ListMediaCapturePipelinesOutput
				p := svc.NewListMediaCapturePipelinesPaginator(client, input)
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
		"list-media-insights-pipeline-configurations": {
			Name:   "list-media-insights-pipeline-configurations",
			Fields: fields_list_media_insights_pipeline_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMediaInsightsPipelineConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_media_insights_pipeline_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMediaInsightsPipelineConfigurations(ctx, input)
				}
				var results []*svc.ListMediaInsightsPipelineConfigurationsOutput
				p := svc.NewListMediaInsightsPipelineConfigurationsPaginator(client, input)
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
		"list-media-pipeline-kinesis-video-stream-pools": {
			Name:   "list-media-pipeline-kinesis-video-stream-pools",
			Fields: fields_list_media_pipeline_kinesis_video_stream_pools,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMediaPipelineKinesisVideoStreamPoolsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_media_pipeline_kinesis_video_stream_pools, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMediaPipelineKinesisVideoStreamPools(ctx, input)
				}
				var results []*svc.ListMediaPipelineKinesisVideoStreamPoolsOutput
				p := svc.NewListMediaPipelineKinesisVideoStreamPoolsPaginator(client, input)
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
		"list-media-pipelines": {
			Name:   "list-media-pipelines",
			Fields: fields_list_media_pipelines,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMediaPipelinesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_media_pipelines, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMediaPipelines(ctx, input)
				}
				var results []*svc.ListMediaPipelinesOutput
				p := svc.NewListMediaPipelinesPaginator(client, input)
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
		"start-speaker-search-task": {
			Name:   "start-speaker-search-task",
			Fields: fields_start_speaker_search_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartSpeakerSearchTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_speaker_search_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartSpeakerSearchTask(ctx, input)
			},
		},
		"start-voice-tone-analysis-task": {
			Name:   "start-voice-tone-analysis-task",
			Fields: fields_start_voice_tone_analysis_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartVoiceToneAnalysisTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_voice_tone_analysis_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartVoiceToneAnalysisTask(ctx, input)
			},
		},
		"stop-speaker-search-task": {
			Name:   "stop-speaker-search-task",
			Fields: fields_stop_speaker_search_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopSpeakerSearchTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_speaker_search_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopSpeakerSearchTask(ctx, input)
			},
		},
		"stop-voice-tone-analysis-task": {
			Name:   "stop-voice-tone-analysis-task",
			Fields: fields_stop_voice_tone_analysis_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopVoiceToneAnalysisTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_voice_tone_analysis_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopVoiceToneAnalysisTask(ctx, input)
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
		"update-media-insights-pipeline-configuration": {
			Name:   "update-media-insights-pipeline-configuration",
			Fields: fields_update_media_insights_pipeline_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMediaInsightsPipelineConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_media_insights_pipeline_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMediaInsightsPipelineConfiguration(ctx, input)
			},
		},
		"update-media-insights-pipeline-status": {
			Name:   "update-media-insights-pipeline-status",
			Fields: fields_update_media_insights_pipeline_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMediaInsightsPipelineStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_media_insights_pipeline_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMediaInsightsPipelineStatus(ctx, input)
			},
		},
		"update-media-pipeline-kinesis-video-stream-pool": {
			Name:   "update-media-pipeline-kinesis-video-stream-pool",
			Fields: fields_update_media_pipeline_kinesis_video_stream_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMediaPipelineKinesisVideoStreamPoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_media_pipeline_kinesis_video_stream_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMediaPipelineKinesisVideoStreamPool(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("chimesdkmediapipelines", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
