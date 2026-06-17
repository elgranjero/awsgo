package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkmediapipelines"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// chimesdkmediapipelinesCmd represents the chimesdkmediapipelines command
var _chimesdkmediapipelinesCmd = &cobra.Command{
	Use:   "chimesdkmediapipelines",
	Short: "AWS chimesdkmediapipelines CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := chimesdkmediapipelines.NewFromConfig(cfg)
		if _chimesdkmediapipelinesCreateMediaCapturePipeline {
			chimesdkmediapipelines_CreateMediaCapturePipeline(cfg, client)
			return
		}
		if _chimesdkmediapipelinesCreateMediaConcatenationPipeline {
			chimesdkmediapipelines_CreateMediaConcatenationPipeline(cfg, client)
			return
		}
		if _chimesdkmediapipelinesCreateMediaInsightsPipeline {
			chimesdkmediapipelines_CreateMediaInsightsPipeline(cfg, client)
			return
		}
		if _chimesdkmediapipelinesCreateMediaInsightsPipelineConfiguration {
			chimesdkmediapipelines_CreateMediaInsightsPipelineConfiguration(cfg, client)
			return
		}
		if _chimesdkmediapipelinesCreateMediaLiveConnectorPipeline {
			chimesdkmediapipelines_CreateMediaLiveConnectorPipeline(cfg, client)
			return
		}
		if _chimesdkmediapipelinesCreateMediaPipelineKinesisVideoStreamPool {
			chimesdkmediapipelines_CreateMediaPipelineKinesisVideoStreamPool(cfg, client)
			return
		}
		if _chimesdkmediapipelinesCreateMediaStreamPipeline {
			chimesdkmediapipelines_CreateMediaStreamPipeline(cfg, client)
			return
		}
		if _chimesdkmediapipelinesDeleteMediaCapturePipeline {
			chimesdkmediapipelines_DeleteMediaCapturePipeline(cfg, client)
			return
		}
		if _chimesdkmediapipelinesDeleteMediaInsightsPipelineConfiguration {
			chimesdkmediapipelines_DeleteMediaInsightsPipelineConfiguration(cfg, client)
			return
		}
		if _chimesdkmediapipelinesDeleteMediaPipeline {
			chimesdkmediapipelines_DeleteMediaPipeline(cfg, client)
			return
		}
		if _chimesdkmediapipelinesDeleteMediaPipelineKinesisVideoStreamPool {
			chimesdkmediapipelines_DeleteMediaPipelineKinesisVideoStreamPool(cfg, client)
			return
		}
		if _chimesdkmediapipelinesGetMediaCapturePipeline {
			chimesdkmediapipelines_GetMediaCapturePipeline(cfg, client)
			return
		}
		if _chimesdkmediapipelinesGetMediaInsightsPipelineConfiguration {
			chimesdkmediapipelines_GetMediaInsightsPipelineConfiguration(cfg, client)
			return
		}
		if _chimesdkmediapipelinesGetMediaPipeline {
			chimesdkmediapipelines_GetMediaPipeline(cfg, client)
			return
		}
		if _chimesdkmediapipelinesGetMediaPipelineKinesisVideoStreamPool {
			chimesdkmediapipelines_GetMediaPipelineKinesisVideoStreamPool(cfg, client)
			return
		}
		if _chimesdkmediapipelinesGetSpeakerSearchTask {
			chimesdkmediapipelines_GetSpeakerSearchTask(cfg, client)
			return
		}
		if _chimesdkmediapipelinesGetVoiceToneAnalysisTask {
			chimesdkmediapipelines_GetVoiceToneAnalysisTask(cfg, client)
			return
		}
		if _chimesdkmediapipelinesListMediaCapturePipelines {
			chimesdkmediapipelines_ListMediaCapturePipelines(cfg, client)
			return
		}
		if _chimesdkmediapipelinesListMediaInsightsPipelineConfigurations {
			chimesdkmediapipelines_ListMediaInsightsPipelineConfigurations(cfg, client)
			return
		}
		if _chimesdkmediapipelinesListMediaPipelineKinesisVideoStreamPools {
			chimesdkmediapipelines_ListMediaPipelineKinesisVideoStreamPools(cfg, client)
			return
		}
		if _chimesdkmediapipelinesListMediaPipelines {
			chimesdkmediapipelines_ListMediaPipelines(cfg, client)
			return
		}
		if _chimesdkmediapipelinesListTagsForResource {
			chimesdkmediapipelines_ListTagsForResource(cfg, client)
			return
		}
		if _chimesdkmediapipelinesStartSpeakerSearchTask {
			chimesdkmediapipelines_StartSpeakerSearchTask(cfg, client)
			return
		}
		if _chimesdkmediapipelinesStartVoiceToneAnalysisTask {
			chimesdkmediapipelines_StartVoiceToneAnalysisTask(cfg, client)
			return
		}
		if _chimesdkmediapipelinesStopSpeakerSearchTask {
			chimesdkmediapipelines_StopSpeakerSearchTask(cfg, client)
			return
		}
		if _chimesdkmediapipelinesStopVoiceToneAnalysisTask {
			chimesdkmediapipelines_StopVoiceToneAnalysisTask(cfg, client)
			return
		}
		if _chimesdkmediapipelinesTagResource {
			chimesdkmediapipelines_TagResource(cfg, client)
			return
		}
		if _chimesdkmediapipelinesUntagResource {
			chimesdkmediapipelines_UntagResource(cfg, client)
			return
		}
		if _chimesdkmediapipelinesUpdateMediaInsightsPipelineConfiguration {
			chimesdkmediapipelines_UpdateMediaInsightsPipelineConfiguration(cfg, client)
			return
		}
		if _chimesdkmediapipelinesUpdateMediaInsightsPipelineStatus {
			chimesdkmediapipelines_UpdateMediaInsightsPipelineStatus(cfg, client)
			return
		}
		if _chimesdkmediapipelinesUpdateMediaPipelineKinesisVideoStreamPool {
			chimesdkmediapipelines_UpdateMediaPipelineKinesisVideoStreamPool(cfg, client)
			return
		}

	},
}

var (
	_chimesdkmediapipelinesCreateMediaCapturePipeline                bool
	_chimesdkmediapipelinesCreateMediaConcatenationPipeline          bool
	_chimesdkmediapipelinesCreateMediaInsightsPipeline               bool
	_chimesdkmediapipelinesCreateMediaInsightsPipelineConfiguration  bool
	_chimesdkmediapipelinesCreateMediaLiveConnectorPipeline          bool
	_chimesdkmediapipelinesCreateMediaPipelineKinesisVideoStreamPool bool
	_chimesdkmediapipelinesCreateMediaStreamPipeline                 bool
	_chimesdkmediapipelinesDeleteMediaCapturePipeline                bool
	_chimesdkmediapipelinesDeleteMediaInsightsPipelineConfiguration  bool
	_chimesdkmediapipelinesDeleteMediaPipeline                       bool
	_chimesdkmediapipelinesDeleteMediaPipelineKinesisVideoStreamPool bool
	_chimesdkmediapipelinesGetMediaCapturePipeline                   bool
	_chimesdkmediapipelinesGetMediaInsightsPipelineConfiguration     bool
	_chimesdkmediapipelinesGetMediaPipeline                          bool
	_chimesdkmediapipelinesGetMediaPipelineKinesisVideoStreamPool    bool
	_chimesdkmediapipelinesGetSpeakerSearchTask                      bool
	_chimesdkmediapipelinesGetVoiceToneAnalysisTask                  bool
	_chimesdkmediapipelinesListMediaCapturePipelines                 bool
	_chimesdkmediapipelinesListMediaInsightsPipelineConfigurations   bool
	_chimesdkmediapipelinesListMediaPipelineKinesisVideoStreamPools  bool
	_chimesdkmediapipelinesListMediaPipelines                        bool
	_chimesdkmediapipelinesListTagsForResource                       bool
	_chimesdkmediapipelinesStartSpeakerSearchTask                    bool
	_chimesdkmediapipelinesStartVoiceToneAnalysisTask                bool
	_chimesdkmediapipelinesStopSpeakerSearchTask                     bool
	_chimesdkmediapipelinesStopVoiceToneAnalysisTask                 bool
	_chimesdkmediapipelinesTagResource                               bool
	_chimesdkmediapipelinesUntagResource                             bool
	_chimesdkmediapipelinesUpdateMediaInsightsPipelineConfiguration  bool
	_chimesdkmediapipelinesUpdateMediaInsightsPipelineStatus         bool
	_chimesdkmediapipelinesUpdateMediaPipelineKinesisVideoStreamPool bool

	_chimesdkmediapipelinesChimeSdkMeetingConfiguration                          string
	_chimesdkmediapipelinesClientRequestToken                                    string
	_chimesdkmediapipelinesElements                                              string
	_chimesdkmediapipelinesIdentifier                                            string
	_chimesdkmediapipelinesKinesisVideoStreamRecordingSourceRuntimeConfiguration string
	_chimesdkmediapipelinesKinesisVideoStreamSourceRuntimeConfiguration          string
	_chimesdkmediapipelinesKinesisVideoStreamSourceTaskConfiguration             string
	_chimesdkmediapipelinesLanguageCode                                          string
	_chimesdkmediapipelinesMaxResults                                            string
	_chimesdkmediapipelinesMediaInsightsPipelineConfigurationArn                 string
	_chimesdkmediapipelinesMediaInsightsPipelineConfigurationName                string
	_chimesdkmediapipelinesMediaInsightsRuntimeMetadata                          string
	_chimesdkmediapipelinesMediaPipelineId                                       string
	_chimesdkmediapipelinesNextToken                                             string
	_chimesdkmediapipelinesPoolName                                              string
	_chimesdkmediapipelinesRealTimeAlertConfiguration                            string
	_chimesdkmediapipelinesResourceAccessRoleArn                                 string
	_chimesdkmediapipelinesResourceARN                                           string
	_chimesdkmediapipelinesS3RecordingSinkRuntimeConfiguration                   string
	_chimesdkmediapipelinesSinkArn                                               string
	_chimesdkmediapipelinesSinkIamRoleArn                                        string
	_chimesdkmediapipelinesSinkType                                              string
	_chimesdkmediapipelinesSinks                                                 string
	_chimesdkmediapipelinesSourceArn                                             string
	_chimesdkmediapipelinesSourceType                                            string
	_chimesdkmediapipelinesSources                                               string
	_chimesdkmediapipelinesSpeakerSearchTaskId                                   string
	_chimesdkmediapipelinesSseAwsKeyManagementParams                             string
	_chimesdkmediapipelinesStreamConfiguration                                   string
	_chimesdkmediapipelinesTagKeys                                               []string
	_chimesdkmediapipelinesTags                                                  string
	_chimesdkmediapipelinesUpdateStatus                                          string
	_chimesdkmediapipelinesVoiceProfileDomainArn                                 string
	_chimesdkmediapipelinesVoiceToneAnalysisTaskId                               string
)

// Creates a media pipeline.
func chimesdkmediapipelines_CreateMediaCapturePipeline(cfg aws.Config, client *chimesdkmediapipelines.Client) {
	input := &chimesdkmediapipelines.CreateMediaCapturePipelineInput{
		// SinkArn: *string, // Required
		// SinkType: types.MediaPipelineSinkType, // Required
		// SourceArn: *string, // Required
		// SourceType: types.MediaPipelineSourceType, // Required
	}

	if len(_chimesdkmediapipelinesSinkArn) > 0 {
		input.SinkArn = aws.String(_chimesdkmediapipelinesSinkArn)
	}
	if len(_chimesdkmediapipelinesSinkType) > 0 {
		if err := assignInputField(input, "SinkType", _chimesdkmediapipelinesSinkType); err != nil {
			log.Errorf("invalid --sink-type: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmediapipelinesSourceArn) > 0 {
		input.SourceArn = aws.String(_chimesdkmediapipelinesSourceArn)
	}
	if len(_chimesdkmediapipelinesSourceType) > 0 {
		if err := assignInputField(input, "SourceType", _chimesdkmediapipelinesSourceType); err != nil {
			log.Errorf("invalid --source-type: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmediapipelinesChimeSdkMeetingConfiguration) > 0 {
		if err := assignInputField(input, "ChimeSdkMeetingConfiguration", _chimesdkmediapipelinesChimeSdkMeetingConfiguration); err != nil {
			log.Errorf("invalid --chime-sdk-meeting-configuration: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmediapipelinesClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_chimesdkmediapipelinesClientRequestToken)
	}
	if len(_chimesdkmediapipelinesSinkIamRoleArn) > 0 {
		input.SinkIamRoleArn = aws.String(_chimesdkmediapipelinesSinkIamRoleArn)
	}
	if len(_chimesdkmediapipelinesSseAwsKeyManagementParams) > 0 {
		if err := assignInputField(input, "SseAwsKeyManagementParams", _chimesdkmediapipelinesSseAwsKeyManagementParams); err != nil {
			log.Errorf("invalid --sse-aws-key-management-params: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmediapipelinesTags) > 0 {
		if err := assignInputField(input, "Tags", _chimesdkmediapipelinesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateMediaCapturePipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a media concatenation pipeline.
func chimesdkmediapipelines_CreateMediaConcatenationPipeline(cfg aws.Config, client *chimesdkmediapipelines.Client) {
	input := &chimesdkmediapipelines.CreateMediaConcatenationPipelineInput{
		// Sinks: []types.ConcatenationSink, // Required
		// Sources: []types.ConcatenationSource, // Required
	}

	if len(_chimesdkmediapipelinesSinks) > 0 {
		if err := assignInputField(input, "Sinks", _chimesdkmediapipelinesSinks); err != nil {
			log.Errorf("invalid --sinks: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmediapipelinesSources) > 0 {
		if err := assignInputField(input, "Sources", _chimesdkmediapipelinesSources); err != nil {
			log.Errorf("invalid --sources: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmediapipelinesClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_chimesdkmediapipelinesClientRequestToken)
	}
	if len(_chimesdkmediapipelinesTags) > 0 {
		if err := assignInputField(input, "Tags", _chimesdkmediapipelinesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateMediaConcatenationPipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a media insights pipeline.
func chimesdkmediapipelines_CreateMediaInsightsPipeline(cfg aws.Config, client *chimesdkmediapipelines.Client) {
	input := &chimesdkmediapipelines.CreateMediaInsightsPipelineInput{
		// MediaInsightsPipelineConfigurationArn: *string, // Required
	}

	if len(_chimesdkmediapipelinesMediaInsightsPipelineConfigurationArn) > 0 {
		input.MediaInsightsPipelineConfigurationArn = aws.String(_chimesdkmediapipelinesMediaInsightsPipelineConfigurationArn)
	}
	if len(_chimesdkmediapipelinesClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_chimesdkmediapipelinesClientRequestToken)
	}
	if len(_chimesdkmediapipelinesKinesisVideoStreamRecordingSourceRuntimeConfiguration) > 0 {
		if err := assignInputField(input, "KinesisVideoStreamRecordingSourceRuntimeConfiguration", _chimesdkmediapipelinesKinesisVideoStreamRecordingSourceRuntimeConfiguration); err != nil {
			log.Errorf("invalid --kinesis-video-stream-recording-source-runtime-configuration: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmediapipelinesKinesisVideoStreamSourceRuntimeConfiguration) > 0 {
		if err := assignInputField(input, "KinesisVideoStreamSourceRuntimeConfiguration", _chimesdkmediapipelinesKinesisVideoStreamSourceRuntimeConfiguration); err != nil {
			log.Errorf("invalid --kinesis-video-stream-source-runtime-configuration: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmediapipelinesMediaInsightsRuntimeMetadata) > 0 {
		if err := assignInputField(input, "MediaInsightsRuntimeMetadata", _chimesdkmediapipelinesMediaInsightsRuntimeMetadata); err != nil {
			log.Errorf("invalid --media-insights-runtime-metadata: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmediapipelinesS3RecordingSinkRuntimeConfiguration) > 0 {
		if err := assignInputField(input, "S3RecordingSinkRuntimeConfiguration", _chimesdkmediapipelinesS3RecordingSinkRuntimeConfiguration); err != nil {
			log.Errorf("invalid --s3-recording-sink-runtime-configuration: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmediapipelinesTags) > 0 {
		if err := assignInputField(input, "Tags", _chimesdkmediapipelinesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateMediaInsightsPipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// A structure that contains the static configurations for a media insights
// pipeline.
func chimesdkmediapipelines_CreateMediaInsightsPipelineConfiguration(cfg aws.Config, client *chimesdkmediapipelines.Client) {
	input := &chimesdkmediapipelines.CreateMediaInsightsPipelineConfigurationInput{
		// Elements: []types.MediaInsightsPipelineConfigurationElement, // Required
		// MediaInsightsPipelineConfigurationName: *string, // Required
		// ResourceAccessRoleArn: *string, // Required
	}

	if len(_chimesdkmediapipelinesElements) > 0 {
		if err := assignInputField(input, "Elements", _chimesdkmediapipelinesElements); err != nil {
			log.Errorf("invalid --elements: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmediapipelinesMediaInsightsPipelineConfigurationName) > 0 {
		input.MediaInsightsPipelineConfigurationName = aws.String(_chimesdkmediapipelinesMediaInsightsPipelineConfigurationName)
	}
	if len(_chimesdkmediapipelinesResourceAccessRoleArn) > 0 {
		input.ResourceAccessRoleArn = aws.String(_chimesdkmediapipelinesResourceAccessRoleArn)
	}
	if len(_chimesdkmediapipelinesClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_chimesdkmediapipelinesClientRequestToken)
	}
	if len(_chimesdkmediapipelinesRealTimeAlertConfiguration) > 0 {
		if err := assignInputField(input, "RealTimeAlertConfiguration", _chimesdkmediapipelinesRealTimeAlertConfiguration); err != nil {
			log.Errorf("invalid --real-time-alert-configuration: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmediapipelinesTags) > 0 {
		if err := assignInputField(input, "Tags", _chimesdkmediapipelinesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateMediaInsightsPipelineConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a media live connector pipeline in an Amazon Chime SDK meeting.
func chimesdkmediapipelines_CreateMediaLiveConnectorPipeline(cfg aws.Config, client *chimesdkmediapipelines.Client) {
	input := &chimesdkmediapipelines.CreateMediaLiveConnectorPipelineInput{
		// Sinks: []types.LiveConnectorSinkConfiguration, // Required
		// Sources: []types.LiveConnectorSourceConfiguration, // Required
	}

	if len(_chimesdkmediapipelinesSinks) > 0 {
		if err := assignInputField(input, "Sinks", _chimesdkmediapipelinesSinks); err != nil {
			log.Errorf("invalid --sinks: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmediapipelinesSources) > 0 {
		if err := assignInputField(input, "Sources", _chimesdkmediapipelinesSources); err != nil {
			log.Errorf("invalid --sources: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmediapipelinesClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_chimesdkmediapipelinesClientRequestToken)
	}
	if len(_chimesdkmediapipelinesTags) > 0 {
		if err := assignInputField(input, "Tags", _chimesdkmediapipelinesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateMediaLiveConnectorPipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Kinesis Video Stream pool for use with media stream pipelines.
// If a meeting uses an opt-in Region as its [MediaRegion], the KVS stream must be in that same
// Region. For example, if a meeting uses the af-south-1 Region, the KVS stream
// must also be in af-south-1 . However, if the meeting uses a Region that AWS
// turns on by default, the KVS stream can be in any available Region, including an
// opt-in Region. For example, if the meeting uses ca-central-1 , the KVS stream
// can be in eu-west-2 , us-east-1 , af-south-1 , or any other Region that the
// Amazon Chime SDK supports.
//
// To learn which AWS Region a meeting uses, call the [GetMeeting] API and use the [MediaRegion] parameter
// from the response.
//
// For more information about opt-in Regions, refer to [Available Regions] in the Amazon Chime SDK
// Developer Guide, and [Specify which AWS Regions your account can use], in the AWS Account Management Reference Guide.
//
// [GetMeeting]: https://docs.aws.amazon.com/chime-sdk/latest/APIReference/API_meeting-chime_GetMeeting.html
// [Specify which AWS Regions your account can use]: https://docs.aws.amazon.com/accounts/latest/reference/manage-acct-regions.html#rande-manage-enable.html
// [Available Regions]: https://docs.aws.amazon.com/chime-sdk/latest/dg/sdk-available-regions.html
// [MediaRegion]: https://docs.aws.amazon.com/chime-sdk/latest/APIReference/API_meeting-chime_CreateMeeting.html#chimesdk-meeting-chime_CreateMeeting-request-MediaRegion
func chimesdkmediapipelines_CreateMediaPipelineKinesisVideoStreamPool(cfg aws.Config, client *chimesdkmediapipelines.Client) {
	input := &chimesdkmediapipelines.CreateMediaPipelineKinesisVideoStreamPoolInput{
		// PoolName: *string, // Required
		// StreamConfiguration: *types.KinesisVideoStreamConfiguration, // Required
	}

	if len(_chimesdkmediapipelinesPoolName) > 0 {
		input.PoolName = aws.String(_chimesdkmediapipelinesPoolName)
	}
	if len(_chimesdkmediapipelinesStreamConfiguration) > 0 {
		if err := assignInputField(input, "StreamConfiguration", _chimesdkmediapipelinesStreamConfiguration); err != nil {
			log.Errorf("invalid --stream-configuration: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmediapipelinesClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_chimesdkmediapipelinesClientRequestToken)
	}
	if len(_chimesdkmediapipelinesTags) > 0 {
		if err := assignInputField(input, "Tags", _chimesdkmediapipelinesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateMediaPipelineKinesisVideoStreamPool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a streaming media pipeline.
func chimesdkmediapipelines_CreateMediaStreamPipeline(cfg aws.Config, client *chimesdkmediapipelines.Client) {
	input := &chimesdkmediapipelines.CreateMediaStreamPipelineInput{
		// Sinks: []types.MediaStreamSink, // Required
		// Sources: []types.MediaStreamSource, // Required
	}

	if len(_chimesdkmediapipelinesSinks) > 0 {
		if err := assignInputField(input, "Sinks", _chimesdkmediapipelinesSinks); err != nil {
			log.Errorf("invalid --sinks: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmediapipelinesSources) > 0 {
		if err := assignInputField(input, "Sources", _chimesdkmediapipelinesSources); err != nil {
			log.Errorf("invalid --sources: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmediapipelinesClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_chimesdkmediapipelinesClientRequestToken)
	}
	if len(_chimesdkmediapipelinesTags) > 0 {
		if err := assignInputField(input, "Tags", _chimesdkmediapipelinesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateMediaStreamPipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the media pipeline.
func chimesdkmediapipelines_DeleteMediaCapturePipeline(cfg aws.Config, client *chimesdkmediapipelines.Client) {
	input := &chimesdkmediapipelines.DeleteMediaCapturePipelineInput{
		// MediaPipelineId: *string, // Required
	}

	if len(_chimesdkmediapipelinesMediaPipelineId) > 0 {
		input.MediaPipelineId = aws.String(_chimesdkmediapipelinesMediaPipelineId)
	}

	if resp, err := client.DeleteMediaCapturePipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified configuration settings.
func chimesdkmediapipelines_DeleteMediaInsightsPipelineConfiguration(cfg aws.Config, client *chimesdkmediapipelines.Client) {
	input := &chimesdkmediapipelines.DeleteMediaInsightsPipelineConfigurationInput{
		// Identifier: *string, // Required
	}

	if len(_chimesdkmediapipelinesIdentifier) > 0 {
		input.Identifier = aws.String(_chimesdkmediapipelinesIdentifier)
	}

	if resp, err := client.DeleteMediaInsightsPipelineConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the media pipeline.
func chimesdkmediapipelines_DeleteMediaPipeline(cfg aws.Config, client *chimesdkmediapipelines.Client) {
	input := &chimesdkmediapipelines.DeleteMediaPipelineInput{
		// MediaPipelineId: *string, // Required
	}

	if len(_chimesdkmediapipelinesMediaPipelineId) > 0 {
		input.MediaPipelineId = aws.String(_chimesdkmediapipelinesMediaPipelineId)
	}

	if resp, err := client.DeleteMediaPipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Kinesis Video Stream pool.
func chimesdkmediapipelines_DeleteMediaPipelineKinesisVideoStreamPool(cfg aws.Config, client *chimesdkmediapipelines.Client) {
	input := &chimesdkmediapipelines.DeleteMediaPipelineKinesisVideoStreamPoolInput{
		// Identifier: *string, // Required
	}

	if len(_chimesdkmediapipelinesIdentifier) > 0 {
		input.Identifier = aws.String(_chimesdkmediapipelinesIdentifier)
	}

	if resp, err := client.DeleteMediaPipelineKinesisVideoStreamPool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an existing media pipeline.
func chimesdkmediapipelines_GetMediaCapturePipeline(cfg aws.Config, client *chimesdkmediapipelines.Client) {
	input := &chimesdkmediapipelines.GetMediaCapturePipelineInput{
		// MediaPipelineId: *string, // Required
	}

	if len(_chimesdkmediapipelinesMediaPipelineId) > 0 {
		input.MediaPipelineId = aws.String(_chimesdkmediapipelinesMediaPipelineId)
	}

	if resp, err := client.GetMediaCapturePipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the configuration settings for a media insights pipeline.
func chimesdkmediapipelines_GetMediaInsightsPipelineConfiguration(cfg aws.Config, client *chimesdkmediapipelines.Client) {
	input := &chimesdkmediapipelines.GetMediaInsightsPipelineConfigurationInput{
		// Identifier: *string, // Required
	}

	if len(_chimesdkmediapipelinesIdentifier) > 0 {
		input.Identifier = aws.String(_chimesdkmediapipelinesIdentifier)
	}

	if resp, err := client.GetMediaInsightsPipelineConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an existing media pipeline.
func chimesdkmediapipelines_GetMediaPipeline(cfg aws.Config, client *chimesdkmediapipelines.Client) {
	input := &chimesdkmediapipelines.GetMediaPipelineInput{
		// MediaPipelineId: *string, // Required
	}

	if len(_chimesdkmediapipelinesMediaPipelineId) > 0 {
		input.MediaPipelineId = aws.String(_chimesdkmediapipelinesMediaPipelineId)
	}

	if resp, err := client.GetMediaPipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an Kinesis video stream pool.
func chimesdkmediapipelines_GetMediaPipelineKinesisVideoStreamPool(cfg aws.Config, client *chimesdkmediapipelines.Client) {
	input := &chimesdkmediapipelines.GetMediaPipelineKinesisVideoStreamPoolInput{
		// Identifier: *string, // Required
	}

	if len(_chimesdkmediapipelinesIdentifier) > 0 {
		input.Identifier = aws.String(_chimesdkmediapipelinesIdentifier)
	}

	if resp, err := client.GetMediaPipelineKinesisVideoStreamPool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the details of the specified speaker search task.
func chimesdkmediapipelines_GetSpeakerSearchTask(cfg aws.Config, client *chimesdkmediapipelines.Client) {
	input := &chimesdkmediapipelines.GetSpeakerSearchTaskInput{
		// Identifier: *string, // Required
		// SpeakerSearchTaskId: *string, // Required
	}

	if len(_chimesdkmediapipelinesIdentifier) > 0 {
		input.Identifier = aws.String(_chimesdkmediapipelinesIdentifier)
	}
	if len(_chimesdkmediapipelinesSpeakerSearchTaskId) > 0 {
		input.SpeakerSearchTaskId = aws.String(_chimesdkmediapipelinesSpeakerSearchTaskId)
	}

	if resp, err := client.GetSpeakerSearchTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the details of a voice tone analysis task.
func chimesdkmediapipelines_GetVoiceToneAnalysisTask(cfg aws.Config, client *chimesdkmediapipelines.Client) {
	input := &chimesdkmediapipelines.GetVoiceToneAnalysisTaskInput{
		// Identifier: *string, // Required
		// VoiceToneAnalysisTaskId: *string, // Required
	}

	if len(_chimesdkmediapipelinesIdentifier) > 0 {
		input.Identifier = aws.String(_chimesdkmediapipelinesIdentifier)
	}
	if len(_chimesdkmediapipelinesVoiceToneAnalysisTaskId) > 0 {
		input.VoiceToneAnalysisTaskId = aws.String(_chimesdkmediapipelinesVoiceToneAnalysisTaskId)
	}

	if resp, err := client.GetVoiceToneAnalysisTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of media pipelines.
func chimesdkmediapipelines_ListMediaCapturePipelines(cfg aws.Config, client *chimesdkmediapipelines.Client) {
	input := &chimesdkmediapipelines.ListMediaCapturePipelinesInput{}

	if len(_chimesdkmediapipelinesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimesdkmediapipelinesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmediapipelinesNextToken) > 0 {
		input.NextToken = aws.String(_chimesdkmediapipelinesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMediaCapturePipelines(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chimesdkmediapipelines.ListMediaCapturePipelinesOutput
	p := chimesdkmediapipelines.NewListMediaCapturePipelinesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists the available media insights pipeline configurations.
func chimesdkmediapipelines_ListMediaInsightsPipelineConfigurations(cfg aws.Config, client *chimesdkmediapipelines.Client) {
	input := &chimesdkmediapipelines.ListMediaInsightsPipelineConfigurationsInput{}

	if len(_chimesdkmediapipelinesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimesdkmediapipelinesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmediapipelinesNextToken) > 0 {
		input.NextToken = aws.String(_chimesdkmediapipelinesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMediaInsightsPipelineConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chimesdkmediapipelines.ListMediaInsightsPipelineConfigurationsOutput
	p := chimesdkmediapipelines.NewListMediaInsightsPipelineConfigurationsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists the video stream pools in the media pipeline.
func chimesdkmediapipelines_ListMediaPipelineKinesisVideoStreamPools(cfg aws.Config, client *chimesdkmediapipelines.Client) {
	input := &chimesdkmediapipelines.ListMediaPipelineKinesisVideoStreamPoolsInput{}

	if len(_chimesdkmediapipelinesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimesdkmediapipelinesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmediapipelinesNextToken) > 0 {
		input.NextToken = aws.String(_chimesdkmediapipelinesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMediaPipelineKinesisVideoStreamPools(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chimesdkmediapipelines.ListMediaPipelineKinesisVideoStreamPoolsOutput
	p := chimesdkmediapipelines.NewListMediaPipelineKinesisVideoStreamPoolsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns a list of media pipelines.
func chimesdkmediapipelines_ListMediaPipelines(cfg aws.Config, client *chimesdkmediapipelines.Client) {
	input := &chimesdkmediapipelines.ListMediaPipelinesInput{}

	if len(_chimesdkmediapipelinesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimesdkmediapipelinesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmediapipelinesNextToken) > 0 {
		input.NextToken = aws.String(_chimesdkmediapipelinesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMediaPipelines(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chimesdkmediapipelines.ListMediaPipelinesOutput
	p := chimesdkmediapipelines.NewListMediaPipelinesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists the tags available for a media pipeline.
func chimesdkmediapipelines_ListTagsForResource(cfg aws.Config, client *chimesdkmediapipelines.Client) {
	input := &chimesdkmediapipelines.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_chimesdkmediapipelinesResourceARN) > 0 {
		input.ResourceARN = aws.String(_chimesdkmediapipelinesResourceARN)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a speaker search task.
// Before starting any speaker search tasks, you must provide all notices and
// obtain all consents from the speaker as required under applicable privacy and
// biometrics laws, and as required under the [AWS service terms]for the Amazon Chime SDK.
//
// [AWS service terms]: https://aws.amazon.com/service-terms/
func chimesdkmediapipelines_StartSpeakerSearchTask(cfg aws.Config, client *chimesdkmediapipelines.Client) {
	input := &chimesdkmediapipelines.StartSpeakerSearchTaskInput{
		// Identifier: *string, // Required
		// VoiceProfileDomainArn: *string, // Required
	}

	if len(_chimesdkmediapipelinesIdentifier) > 0 {
		input.Identifier = aws.String(_chimesdkmediapipelinesIdentifier)
	}
	if len(_chimesdkmediapipelinesVoiceProfileDomainArn) > 0 {
		input.VoiceProfileDomainArn = aws.String(_chimesdkmediapipelinesVoiceProfileDomainArn)
	}
	if len(_chimesdkmediapipelinesClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_chimesdkmediapipelinesClientRequestToken)
	}
	if len(_chimesdkmediapipelinesKinesisVideoStreamSourceTaskConfiguration) > 0 {
		if err := assignInputField(input, "KinesisVideoStreamSourceTaskConfiguration", _chimesdkmediapipelinesKinesisVideoStreamSourceTaskConfiguration); err != nil {
			log.Errorf("invalid --kinesis-video-stream-source-task-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartSpeakerSearchTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a voice tone analysis task. For more information about voice tone
// analysis, see [Using Amazon Chime SDK voice analytics]in the Amazon Chime SDK Developer Guide.
//
// Before starting any voice tone analysis tasks, you must provide all notices and
// obtain all consents from the speaker as required under applicable privacy and
// biometrics laws, and as required under the [AWS service terms]for the Amazon Chime SDK.
//
// [Using Amazon Chime SDK voice analytics]: https://docs.aws.amazon.com/chime-sdk/latest/dg/voice-analytics.html
// [AWS service terms]: https://aws.amazon.com/service-terms/
func chimesdkmediapipelines_StartVoiceToneAnalysisTask(cfg aws.Config, client *chimesdkmediapipelines.Client) {
	input := &chimesdkmediapipelines.StartVoiceToneAnalysisTaskInput{
		// Identifier: *string, // Required
		// LanguageCode: types.VoiceAnalyticsLanguageCode, // Required
	}

	if len(_chimesdkmediapipelinesIdentifier) > 0 {
		input.Identifier = aws.String(_chimesdkmediapipelinesIdentifier)
	}
	if len(_chimesdkmediapipelinesLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _chimesdkmediapipelinesLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmediapipelinesClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_chimesdkmediapipelinesClientRequestToken)
	}
	if len(_chimesdkmediapipelinesKinesisVideoStreamSourceTaskConfiguration) > 0 {
		if err := assignInputField(input, "KinesisVideoStreamSourceTaskConfiguration", _chimesdkmediapipelinesKinesisVideoStreamSourceTaskConfiguration); err != nil {
			log.Errorf("invalid --kinesis-video-stream-source-task-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartVoiceToneAnalysisTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a speaker search task.
func chimesdkmediapipelines_StopSpeakerSearchTask(cfg aws.Config, client *chimesdkmediapipelines.Client) {
	input := &chimesdkmediapipelines.StopSpeakerSearchTaskInput{
		// Identifier: *string, // Required
		// SpeakerSearchTaskId: *string, // Required
	}

	if len(_chimesdkmediapipelinesIdentifier) > 0 {
		input.Identifier = aws.String(_chimesdkmediapipelinesIdentifier)
	}
	if len(_chimesdkmediapipelinesSpeakerSearchTaskId) > 0 {
		input.SpeakerSearchTaskId = aws.String(_chimesdkmediapipelinesSpeakerSearchTaskId)
	}

	if resp, err := client.StopSpeakerSearchTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a voice tone analysis task.
func chimesdkmediapipelines_StopVoiceToneAnalysisTask(cfg aws.Config, client *chimesdkmediapipelines.Client) {
	input := &chimesdkmediapipelines.StopVoiceToneAnalysisTaskInput{
		// Identifier: *string, // Required
		// VoiceToneAnalysisTaskId: *string, // Required
	}

	if len(_chimesdkmediapipelinesIdentifier) > 0 {
		input.Identifier = aws.String(_chimesdkmediapipelinesIdentifier)
	}
	if len(_chimesdkmediapipelinesVoiceToneAnalysisTaskId) > 0 {
		input.VoiceToneAnalysisTaskId = aws.String(_chimesdkmediapipelinesVoiceToneAnalysisTaskId)
	}

	if resp, err := client.StopVoiceToneAnalysisTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The ARN of the media pipeline that you want to tag. Consists of the pipeline's
// endpoint region, resource ID, and pipeline ID.
func chimesdkmediapipelines_TagResource(cfg aws.Config, client *chimesdkmediapipelines.Client) {
	input := &chimesdkmediapipelines.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_chimesdkmediapipelinesResourceARN) > 0 {
		input.ResourceARN = aws.String(_chimesdkmediapipelinesResourceARN)
	}
	if len(_chimesdkmediapipelinesTags) > 0 {
		if err := assignInputField(input, "Tags", _chimesdkmediapipelinesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.TagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes any tags from a media pipeline.
func chimesdkmediapipelines_UntagResource(cfg aws.Config, client *chimesdkmediapipelines.Client) {
	input := &chimesdkmediapipelines.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_chimesdkmediapipelinesResourceARN) > 0 {
		input.ResourceARN = aws.String(_chimesdkmediapipelinesResourceARN)
	}
	if len(_chimesdkmediapipelinesTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _chimesdkmediapipelinesTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the media insights pipeline's configuration settings.
func chimesdkmediapipelines_UpdateMediaInsightsPipelineConfiguration(cfg aws.Config, client *chimesdkmediapipelines.Client) {
	input := &chimesdkmediapipelines.UpdateMediaInsightsPipelineConfigurationInput{
		// Elements: []types.MediaInsightsPipelineConfigurationElement, // Required
		// Identifier: *string, // Required
		// ResourceAccessRoleArn: *string, // Required
	}

	if len(_chimesdkmediapipelinesElements) > 0 {
		if err := assignInputField(input, "Elements", _chimesdkmediapipelinesElements); err != nil {
			log.Errorf("invalid --elements: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmediapipelinesIdentifier) > 0 {
		input.Identifier = aws.String(_chimesdkmediapipelinesIdentifier)
	}
	if len(_chimesdkmediapipelinesResourceAccessRoleArn) > 0 {
		input.ResourceAccessRoleArn = aws.String(_chimesdkmediapipelinesResourceAccessRoleArn)
	}
	if len(_chimesdkmediapipelinesRealTimeAlertConfiguration) > 0 {
		if err := assignInputField(input, "RealTimeAlertConfiguration", _chimesdkmediapipelinesRealTimeAlertConfiguration); err != nil {
			log.Errorf("invalid --real-time-alert-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateMediaInsightsPipelineConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the status of a media insights pipeline.
func chimesdkmediapipelines_UpdateMediaInsightsPipelineStatus(cfg aws.Config, client *chimesdkmediapipelines.Client) {
	input := &chimesdkmediapipelines.UpdateMediaInsightsPipelineStatusInput{
		// Identifier: *string, // Required
		// UpdateStatus: types.MediaPipelineStatusUpdate, // Required
	}

	if len(_chimesdkmediapipelinesIdentifier) > 0 {
		input.Identifier = aws.String(_chimesdkmediapipelinesIdentifier)
	}
	if len(_chimesdkmediapipelinesUpdateStatus) > 0 {
		if err := assignInputField(input, "UpdateStatus", _chimesdkmediapipelinesUpdateStatus); err != nil {
			log.Errorf("invalid --update-status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateMediaInsightsPipelineStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an Amazon Kinesis Video Stream pool in a media pipeline.
func chimesdkmediapipelines_UpdateMediaPipelineKinesisVideoStreamPool(cfg aws.Config, client *chimesdkmediapipelines.Client) {
	input := &chimesdkmediapipelines.UpdateMediaPipelineKinesisVideoStreamPoolInput{
		// Identifier: *string, // Required
	}

	if len(_chimesdkmediapipelinesIdentifier) > 0 {
		input.Identifier = aws.String(_chimesdkmediapipelinesIdentifier)
	}
	if len(_chimesdkmediapipelinesStreamConfiguration) > 0 {
		if err := assignInputField(input, "StreamConfiguration", _chimesdkmediapipelinesStreamConfiguration); err != nil {
			log.Errorf("invalid --stream-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateMediaPipelineKinesisVideoStreamPool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_chimesdkmediapipelinesCmd)
	_chimesdkmediapipelinesCmd.Flags().SortFlags = false

	_chimesdkmediapipelinesCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_chimesdkmediapipelinesCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_chimesdkmediapipelinesCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_chimesdkmediapipelinesCmd.Flags().StringVarP(&_chimesdkmediapipelinesChimeSdkMeetingConfiguration, "chime-sdk-meeting-configuration", "", "", "Chime Sdk Meeting Configuration")
	_chimesdkmediapipelinesCmd.Flags().StringVarP(&_chimesdkmediapipelinesClientRequestToken, "client-request-token", "", "", "Client Request Token")
	_chimesdkmediapipelinesCmd.Flags().StringVarP(&_chimesdkmediapipelinesElements, "elements", "", "", "Elements")
	_chimesdkmediapipelinesCmd.Flags().StringVarP(&_chimesdkmediapipelinesIdentifier, "identifier", "", "", "Identifier")
	_chimesdkmediapipelinesCmd.Flags().StringVarP(&_chimesdkmediapipelinesKinesisVideoStreamRecordingSourceRuntimeConfiguration, "kinesis-video-stream-recording-source-runtime-configuration", "", "", "Kinesis Video Stream Recording Source Runtime Configuration")
	_chimesdkmediapipelinesCmd.Flags().StringVarP(&_chimesdkmediapipelinesKinesisVideoStreamSourceRuntimeConfiguration, "kinesis-video-stream-source-runtime-configuration", "", "", "Kinesis Video Stream Source Runtime Configuration")
	_chimesdkmediapipelinesCmd.Flags().StringVarP(&_chimesdkmediapipelinesKinesisVideoStreamSourceTaskConfiguration, "kinesis-video-stream-source-task-configuration", "", "", "Kinesis Video Stream Source Task Configuration")
	_chimesdkmediapipelinesCmd.Flags().StringVarP(&_chimesdkmediapipelinesLanguageCode, "language-code", "", "", "Language Code")
	_chimesdkmediapipelinesCmd.Flags().StringVarP(&_chimesdkmediapipelinesMaxResults, "max-results", "", "", "Max Results")
	_chimesdkmediapipelinesCmd.Flags().StringVarP(&_chimesdkmediapipelinesMediaInsightsPipelineConfigurationArn, "media-insights-pipeline-configuration-arn", "", "", "Media Insights Pipeline Configuration ARN")
	_chimesdkmediapipelinesCmd.Flags().StringVarP(&_chimesdkmediapipelinesMediaInsightsPipelineConfigurationName, "media-insights-pipeline-configuration-name", "", "", "Media Insights Pipeline Configuration Name")
	_chimesdkmediapipelinesCmd.Flags().StringVarP(&_chimesdkmediapipelinesMediaInsightsRuntimeMetadata, "media-insights-runtime-metadata", "", "", "Media Insights Runtime Metadata")
	_chimesdkmediapipelinesCmd.Flags().StringVarP(&_chimesdkmediapipelinesMediaPipelineId, "media-pipeline-id", "", "", "Media Pipeline ID")
	_chimesdkmediapipelinesCmd.Flags().StringVarP(&_chimesdkmediapipelinesNextToken, "next-token", "", "", "Next Token")
	_chimesdkmediapipelinesCmd.Flags().StringVarP(&_chimesdkmediapipelinesPoolName, "pool-name", "", "", "Pool Name")
	_chimesdkmediapipelinesCmd.Flags().StringVarP(&_chimesdkmediapipelinesRealTimeAlertConfiguration, "real-time-alert-configuration", "", "", "Real Time Alert Configuration")
	_chimesdkmediapipelinesCmd.Flags().StringVarP(&_chimesdkmediapipelinesResourceAccessRoleArn, "resource-access-role-arn", "", "", "Resource Access Role ARN")
	_chimesdkmediapipelinesCmd.Flags().StringVarP(&_chimesdkmediapipelinesResourceARN, "resource-arn", "", "", "Resource ARN")
	_chimesdkmediapipelinesCmd.Flags().StringVarP(&_chimesdkmediapipelinesS3RecordingSinkRuntimeConfiguration, "s3-recording-sink-runtime-configuration", "", "", "S3 Recording Sink Runtime Configuration")
	_chimesdkmediapipelinesCmd.Flags().StringVarP(&_chimesdkmediapipelinesSinkArn, "sink-arn", "", "", "Sink ARN")
	_chimesdkmediapipelinesCmd.Flags().StringVarP(&_chimesdkmediapipelinesSinkIamRoleArn, "sink-iam-role-arn", "", "", "Sink IAM Role ARN")
	_chimesdkmediapipelinesCmd.Flags().StringVarP(&_chimesdkmediapipelinesSinkType, "sink-type", "", "", "Sink Type")
	_chimesdkmediapipelinesCmd.Flags().StringVarP(&_chimesdkmediapipelinesSinks, "sinks", "", "", "Sinks")
	_chimesdkmediapipelinesCmd.Flags().StringVarP(&_chimesdkmediapipelinesSourceArn, "source-arn", "", "", "Source ARN")
	_chimesdkmediapipelinesCmd.Flags().StringVarP(&_chimesdkmediapipelinesSourceType, "source-type", "", "", "Source Type")
	_chimesdkmediapipelinesCmd.Flags().StringVarP(&_chimesdkmediapipelinesSources, "sources", "", "", "Sources")
	_chimesdkmediapipelinesCmd.Flags().StringVarP(&_chimesdkmediapipelinesSpeakerSearchTaskId, "speaker-search-task-id", "", "", "Speaker Search Task ID")
	_chimesdkmediapipelinesCmd.Flags().StringVarP(&_chimesdkmediapipelinesSseAwsKeyManagementParams, "sse-aws-key-management-params", "", "", "SSE AWS Key Management Params")
	_chimesdkmediapipelinesCmd.Flags().StringVarP(&_chimesdkmediapipelinesStreamConfiguration, "stream-configuration", "", "", "Stream Configuration")
	_chimesdkmediapipelinesCmd.Flags().StringSliceVarP(&_chimesdkmediapipelinesTagKeys, "tag-keys", "", nil, "Tag Keys")
	_chimesdkmediapipelinesCmd.Flags().StringVarP(&_chimesdkmediapipelinesTags, "tags", "", "", "Tags")
	_chimesdkmediapipelinesCmd.Flags().StringVarP(&_chimesdkmediapipelinesUpdateStatus, "update-status", "", "", "Update Status")
	_chimesdkmediapipelinesCmd.Flags().StringVarP(&_chimesdkmediapipelinesVoiceProfileDomainArn, "voice-profile-domain-arn", "", "", "Voice Profile Domain ARN")
	_chimesdkmediapipelinesCmd.Flags().StringVarP(&_chimesdkmediapipelinesVoiceToneAnalysisTaskId, "voice-tone-analysis-task-id", "", "", "Voice Tone Analysis Task ID")

	_chimesdkmediapipelinesCmd.Flags().BoolVarP(&_chimesdkmediapipelinesCreateMediaCapturePipeline, "create-media-capture-pipeline", "", false, "Create Media Capture Pipeline")
	_chimesdkmediapipelinesCmd.Flags().BoolVarP(&_chimesdkmediapipelinesCreateMediaConcatenationPipeline, "create-media-concatenation-pipeline", "", false, "Create Media Concatenation Pipeline")
	_chimesdkmediapipelinesCmd.Flags().BoolVarP(&_chimesdkmediapipelinesCreateMediaInsightsPipeline, "create-media-insights-pipeline", "", false, "Create Media Insights Pipeline")
	_chimesdkmediapipelinesCmd.Flags().BoolVarP(&_chimesdkmediapipelinesCreateMediaInsightsPipelineConfiguration, "create-media-insights-pipeline-configuration", "", false, "Create Media Insights Pipeline Configuration")
	_chimesdkmediapipelinesCmd.Flags().BoolVarP(&_chimesdkmediapipelinesCreateMediaLiveConnectorPipeline, "create-media-live-connector-pipeline", "", false, "Create Media Live Connector Pipeline")
	_chimesdkmediapipelinesCmd.Flags().BoolVarP(&_chimesdkmediapipelinesCreateMediaPipelineKinesisVideoStreamPool, "create-media-pipeline-kinesis-video-stream-pool", "", false, "Create Media Pipeline Kinesis Video Stream Pool")
	_chimesdkmediapipelinesCmd.Flags().BoolVarP(&_chimesdkmediapipelinesCreateMediaStreamPipeline, "create-media-stream-pipeline", "", false, "Create Media Stream Pipeline")
	_chimesdkmediapipelinesCmd.Flags().BoolVarP(&_chimesdkmediapipelinesDeleteMediaCapturePipeline, "delete-media-capture-pipeline", "", false, "Delete Media Capture Pipeline")
	_chimesdkmediapipelinesCmd.Flags().BoolVarP(&_chimesdkmediapipelinesDeleteMediaInsightsPipelineConfiguration, "delete-media-insights-pipeline-configuration", "", false, "Delete Media Insights Pipeline Configuration")
	_chimesdkmediapipelinesCmd.Flags().BoolVarP(&_chimesdkmediapipelinesDeleteMediaPipeline, "delete-media-pipeline", "", false, "Delete Media Pipeline")
	_chimesdkmediapipelinesCmd.Flags().BoolVarP(&_chimesdkmediapipelinesDeleteMediaPipelineKinesisVideoStreamPool, "delete-media-pipeline-kinesis-video-stream-pool", "", false, "Delete Media Pipeline Kinesis Video Stream Pool")
	_chimesdkmediapipelinesCmd.Flags().BoolVarP(&_chimesdkmediapipelinesGetMediaCapturePipeline, "get-media-capture-pipeline", "", false, "Get Media Capture Pipeline")
	_chimesdkmediapipelinesCmd.Flags().BoolVarP(&_chimesdkmediapipelinesGetMediaInsightsPipelineConfiguration, "get-media-insights-pipeline-configuration", "", false, "Get Media Insights Pipeline Configuration")
	_chimesdkmediapipelinesCmd.Flags().BoolVarP(&_chimesdkmediapipelinesGetMediaPipeline, "get-media-pipeline", "", false, "Get Media Pipeline")
	_chimesdkmediapipelinesCmd.Flags().BoolVarP(&_chimesdkmediapipelinesGetMediaPipelineKinesisVideoStreamPool, "get-media-pipeline-kinesis-video-stream-pool", "", false, "Get Media Pipeline Kinesis Video Stream Pool")
	_chimesdkmediapipelinesCmd.Flags().BoolVarP(&_chimesdkmediapipelinesGetSpeakerSearchTask, "get-speaker-search-task", "", false, "Get Speaker Search Task")
	_chimesdkmediapipelinesCmd.Flags().BoolVarP(&_chimesdkmediapipelinesGetVoiceToneAnalysisTask, "get-voice-tone-analysis-task", "", false, "Get Voice Tone Analysis Task")
	_chimesdkmediapipelinesCmd.Flags().BoolVarP(&_chimesdkmediapipelinesListMediaCapturePipelines, "list-media-capture-pipelines", "", false, "List Media Capture Pipelines")
	_chimesdkmediapipelinesCmd.Flags().BoolVarP(&_chimesdkmediapipelinesListMediaInsightsPipelineConfigurations, "list-media-insights-pipeline-configurations", "", false, "List Media Insights Pipeline Configurations")
	_chimesdkmediapipelinesCmd.Flags().BoolVarP(&_chimesdkmediapipelinesListMediaPipelineKinesisVideoStreamPools, "list-media-pipeline-kinesis-video-stream-pools", "", false, "List Media Pipeline Kinesis Video Stream Pools")
	_chimesdkmediapipelinesCmd.Flags().BoolVarP(&_chimesdkmediapipelinesListMediaPipelines, "list-media-pipelines", "", false, "List Media Pipelines")
	_chimesdkmediapipelinesCmd.Flags().BoolVarP(&_chimesdkmediapipelinesListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_chimesdkmediapipelinesCmd.Flags().BoolVarP(&_chimesdkmediapipelinesStartSpeakerSearchTask, "start-speaker-search-task", "", false, "Start Speaker Search Task")
	_chimesdkmediapipelinesCmd.Flags().BoolVarP(&_chimesdkmediapipelinesStartVoiceToneAnalysisTask, "start-voice-tone-analysis-task", "", false, "Start Voice Tone Analysis Task")
	_chimesdkmediapipelinesCmd.Flags().BoolVarP(&_chimesdkmediapipelinesStopSpeakerSearchTask, "stop-speaker-search-task", "", false, "Stop Speaker Search Task")
	_chimesdkmediapipelinesCmd.Flags().BoolVarP(&_chimesdkmediapipelinesStopVoiceToneAnalysisTask, "stop-voice-tone-analysis-task", "", false, "Stop Voice Tone Analysis Task")
	_chimesdkmediapipelinesCmd.Flags().BoolVarP(&_chimesdkmediapipelinesTagResource, "tag-resource", "", false, "Tag Resource")
	_chimesdkmediapipelinesCmd.Flags().BoolVarP(&_chimesdkmediapipelinesUntagResource, "untag-resource", "", false, "Untag Resource")
	_chimesdkmediapipelinesCmd.Flags().BoolVarP(&_chimesdkmediapipelinesUpdateMediaInsightsPipelineConfiguration, "update-media-insights-pipeline-configuration", "", false, "Update Media Insights Pipeline Configuration")
	_chimesdkmediapipelinesCmd.Flags().BoolVarP(&_chimesdkmediapipelinesUpdateMediaInsightsPipelineStatus, "update-media-insights-pipeline-status", "", false, "Update Media Insights Pipeline Status")
	_chimesdkmediapipelinesCmd.Flags().BoolVarP(&_chimesdkmediapipelinesUpdateMediaPipelineKinesisVideoStreamPool, "update-media-pipeline-kinesis-video-stream-pool", "", false, "Update Media Pipeline Kinesis Video Stream Pool")

}
