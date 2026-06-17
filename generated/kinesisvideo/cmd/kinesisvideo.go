package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideo"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// kinesisvideoCmd represents the kinesisvideo command
var _kinesisvideoCmd = &cobra.Command{
	Use:   "kinesisvideo",
	Short: "AWS kinesisvideo CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := kinesisvideo.NewFromConfig(cfg)
		if _kinesisvideoCreateSignalingChannel {
			kinesisvideo_CreateSignalingChannel(cfg, client)
			return
		}
		if _kinesisvideoCreateStream {
			kinesisvideo_CreateStream(cfg, client)
			return
		}
		if _kinesisvideoDeleteEdgeConfiguration {
			kinesisvideo_DeleteEdgeConfiguration(cfg, client)
			return
		}
		if _kinesisvideoDeleteSignalingChannel {
			kinesisvideo_DeleteSignalingChannel(cfg, client)
			return
		}
		if _kinesisvideoDeleteStream {
			kinesisvideo_DeleteStream(cfg, client)
			return
		}
		if _kinesisvideoDescribeEdgeConfiguration {
			kinesisvideo_DescribeEdgeConfiguration(cfg, client)
			return
		}
		if _kinesisvideoDescribeImageGenerationConfiguration {
			kinesisvideo_DescribeImageGenerationConfiguration(cfg, client)
			return
		}
		if _kinesisvideoDescribeMappedResourceConfiguration {
			kinesisvideo_DescribeMappedResourceConfiguration(cfg, client)
			return
		}
		if _kinesisvideoDescribeMediaStorageConfiguration {
			kinesisvideo_DescribeMediaStorageConfiguration(cfg, client)
			return
		}
		if _kinesisvideoDescribeNotificationConfiguration {
			kinesisvideo_DescribeNotificationConfiguration(cfg, client)
			return
		}
		if _kinesisvideoDescribeSignalingChannel {
			kinesisvideo_DescribeSignalingChannel(cfg, client)
			return
		}
		if _kinesisvideoDescribeStream {
			kinesisvideo_DescribeStream(cfg, client)
			return
		}
		if _kinesisvideoDescribeStreamStorageConfiguration {
			kinesisvideo_DescribeStreamStorageConfiguration(cfg, client)
			return
		}
		if _kinesisvideoGetDataEndpoint {
			kinesisvideo_GetDataEndpoint(cfg, client)
			return
		}
		if _kinesisvideoGetSignalingChannelEndpoint {
			kinesisvideo_GetSignalingChannelEndpoint(cfg, client)
			return
		}
		if _kinesisvideoListEdgeAgentConfigurations {
			kinesisvideo_ListEdgeAgentConfigurations(cfg, client)
			return
		}
		if _kinesisvideoListSignalingChannels {
			kinesisvideo_ListSignalingChannels(cfg, client)
			return
		}
		if _kinesisvideoListStreams {
			kinesisvideo_ListStreams(cfg, client)
			return
		}
		if _kinesisvideoListTagsForResource {
			kinesisvideo_ListTagsForResource(cfg, client)
			return
		}
		if _kinesisvideoListTagsForStream {
			kinesisvideo_ListTagsForStream(cfg, client)
			return
		}
		if _kinesisvideoStartEdgeConfigurationUpdate {
			kinesisvideo_StartEdgeConfigurationUpdate(cfg, client)
			return
		}
		if _kinesisvideoTagResource {
			kinesisvideo_TagResource(cfg, client)
			return
		}
		if _kinesisvideoTagStream {
			kinesisvideo_TagStream(cfg, client)
			return
		}
		if _kinesisvideoUntagResource {
			kinesisvideo_UntagResource(cfg, client)
			return
		}
		if _kinesisvideoUntagStream {
			kinesisvideo_UntagStream(cfg, client)
			return
		}
		if _kinesisvideoUpdateDataRetention {
			kinesisvideo_UpdateDataRetention(cfg, client)
			return
		}
		if _kinesisvideoUpdateImageGenerationConfiguration {
			kinesisvideo_UpdateImageGenerationConfiguration(cfg, client)
			return
		}
		if _kinesisvideoUpdateMediaStorageConfiguration {
			kinesisvideo_UpdateMediaStorageConfiguration(cfg, client)
			return
		}
		if _kinesisvideoUpdateNotificationConfiguration {
			kinesisvideo_UpdateNotificationConfiguration(cfg, client)
			return
		}
		if _kinesisvideoUpdateSignalingChannel {
			kinesisvideo_UpdateSignalingChannel(cfg, client)
			return
		}
		if _kinesisvideoUpdateStream {
			kinesisvideo_UpdateStream(cfg, client)
			return
		}
		if _kinesisvideoUpdateStreamStorageConfiguration {
			kinesisvideo_UpdateStreamStorageConfiguration(cfg, client)
			return
		}

	},
}

var (
	_kinesisvideoCreateSignalingChannel               bool
	_kinesisvideoCreateStream                         bool
	_kinesisvideoDeleteEdgeConfiguration              bool
	_kinesisvideoDeleteSignalingChannel               bool
	_kinesisvideoDeleteStream                         bool
	_kinesisvideoDescribeEdgeConfiguration            bool
	_kinesisvideoDescribeImageGenerationConfiguration bool
	_kinesisvideoDescribeMappedResourceConfiguration  bool
	_kinesisvideoDescribeMediaStorageConfiguration    bool
	_kinesisvideoDescribeNotificationConfiguration    bool
	_kinesisvideoDescribeSignalingChannel             bool
	_kinesisvideoDescribeStream                       bool
	_kinesisvideoDescribeStreamStorageConfiguration   bool
	_kinesisvideoGetDataEndpoint                      bool
	_kinesisvideoGetSignalingChannelEndpoint          bool
	_kinesisvideoListEdgeAgentConfigurations          bool
	_kinesisvideoListSignalingChannels                bool
	_kinesisvideoListStreams                          bool
	_kinesisvideoListTagsForResource                  bool
	_kinesisvideoListTagsForStream                    bool
	_kinesisvideoStartEdgeConfigurationUpdate         bool
	_kinesisvideoTagResource                          bool
	_kinesisvideoTagStream                            bool
	_kinesisvideoUntagResource                        bool
	_kinesisvideoUntagStream                          bool
	_kinesisvideoUpdateDataRetention                  bool
	_kinesisvideoUpdateImageGenerationConfiguration   bool
	_kinesisvideoUpdateMediaStorageConfiguration      bool
	_kinesisvideoUpdateNotificationConfiguration      bool
	_kinesisvideoUpdateSignalingChannel               bool
	_kinesisvideoUpdateStream                         bool
	_kinesisvideoUpdateStreamStorageConfiguration     bool

	_kinesisvideoAPIName                                  string
	_kinesisvideoChannelARN                               string
	_kinesisvideoChannelName                              string
	_kinesisvideoChannelNameCondition                     string
	_kinesisvideoChannelType                              string
	_kinesisvideoCurrentVersion                           string
	_kinesisvideoDataRetentionChangeInHours               string
	_kinesisvideoDataRetentionInHours                     string
	_kinesisvideoDeviceName                               string
	_kinesisvideoEdgeConfig                               string
	_kinesisvideoHubDeviceArn                             string
	_kinesisvideoImageGenerationConfiguration             string
	_kinesisvideoKmsKeyId                                 string
	_kinesisvideoMaxResults                               string
	_kinesisvideoMediaStorageConfiguration                string
	_kinesisvideoMediaType                                string
	_kinesisvideoNextToken                                string
	_kinesisvideoNotificationConfiguration                string
	_kinesisvideoOperation                                string
	_kinesisvideoResourceARN                              string
	_kinesisvideoSingleMasterChannelEndpointConfiguration string
	_kinesisvideoSingleMasterConfiguration                string
	_kinesisvideoStreamARN                                string
	_kinesisvideoStreamName                               string
	_kinesisvideoStreamNameCondition                      string
	_kinesisvideoStreamStorageConfiguration               string
	_kinesisvideoTagKeyList                               []string
	_kinesisvideoTags                                     string
)

// Creates a signaling channel.
// CreateSignalingChannel is an asynchronous operation.
func kinesisvideo_CreateSignalingChannel(cfg aws.Config, client *kinesisvideo.Client) {
	input := &kinesisvideo.CreateSignalingChannelInput{
		// ChannelName: *string, // Required
	}

	if len(_kinesisvideoChannelName) > 0 {
		input.ChannelName = aws.String(_kinesisvideoChannelName)
	}
	if len(_kinesisvideoChannelType) > 0 {
		if err := assignInputField(input, "ChannelType", _kinesisvideoChannelType); err != nil {
			log.Errorf("invalid --channel-type: %s", err.Error())
			return
		}
	}
	if len(_kinesisvideoSingleMasterConfiguration) > 0 {
		if err := assignInputField(input, "SingleMasterConfiguration", _kinesisvideoSingleMasterConfiguration); err != nil {
			log.Errorf("invalid --single-master-configuration: %s", err.Error())
			return
		}
	}
	if len(_kinesisvideoTags) > 0 {
		if err := assignInputField(input, "Tags", _kinesisvideoTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSignalingChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Kinesis video stream.
// When you create a new stream, Kinesis Video Streams assigns it a version
// number. When you change the stream's metadata, Kinesis Video Streams updates the
// version.
//
// CreateStream is an asynchronous operation.
//
// For information about how the service works, see [How it Works].
//
// You must have permissions for the KinesisVideo:CreateStream action.
//
// [How it Works]: https://docs.aws.amazon.com/kinesisvideostreams/latest/dg/how-it-works.html
func kinesisvideo_CreateStream(cfg aws.Config, client *kinesisvideo.Client) {
	input := &kinesisvideo.CreateStreamInput{
		// StreamName: *string, // Required
	}

	if len(_kinesisvideoStreamName) > 0 {
		input.StreamName = aws.String(_kinesisvideoStreamName)
	}
	if len(_kinesisvideoDataRetentionInHours) > 0 {
		if err := assignInputField(input, "DataRetentionInHours", _kinesisvideoDataRetentionInHours); err != nil {
			log.Errorf("invalid --data-retention-in-hours: %s", err.Error())
			return
		}
	}
	if len(_kinesisvideoDeviceName) > 0 {
		input.DeviceName = aws.String(_kinesisvideoDeviceName)
	}
	if len(_kinesisvideoKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_kinesisvideoKmsKeyId)
	}
	if len(_kinesisvideoMediaType) > 0 {
		input.MediaType = aws.String(_kinesisvideoMediaType)
	}
	if len(_kinesisvideoStreamStorageConfiguration) > 0 {
		if err := assignInputField(input, "StreamStorageConfiguration", _kinesisvideoStreamStorageConfiguration); err != nil {
			log.Errorf("invalid --stream-storage-configuration: %s", err.Error())
			return
		}
	}
	if len(_kinesisvideoTags) > 0 {
		if err := assignInputField(input, "Tags", _kinesisvideoTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// An asynchronous API that deletes a stream’s existing edge configuration, as
// well as the corresponding media from the Edge Agent.
//
// When you invoke this API, the sync status is set to DELETING . A deletion
// process starts, in which active edge jobs are stopped and all media is deleted
// from the edge device. The time to delete varies, depending on the total amount
// of stored media. If the deletion process fails, the sync status changes to
// DELETE_FAILED . You will need to re-try the deletion.
//
// When the deletion process has completed successfully, the edge configuration is
// no longer accessible.
func kinesisvideo_DeleteEdgeConfiguration(cfg aws.Config, client *kinesisvideo.Client) {
	input := &kinesisvideo.DeleteEdgeConfigurationInput{}

	if len(_kinesisvideoStreamARN) > 0 {
		input.StreamARN = aws.String(_kinesisvideoStreamARN)
	}
	if len(_kinesisvideoStreamName) > 0 {
		input.StreamName = aws.String(_kinesisvideoStreamName)
	}

	if resp, err := client.DeleteEdgeConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specified signaling channel. DeleteSignalingChannel is an
// asynchronous operation. If you don't specify the channel's current version, the
// most recent version is deleted.
func kinesisvideo_DeleteSignalingChannel(cfg aws.Config, client *kinesisvideo.Client) {
	input := &kinesisvideo.DeleteSignalingChannelInput{
		// ChannelARN: *string, // Required
	}

	if len(_kinesisvideoChannelARN) > 0 {
		input.ChannelARN = aws.String(_kinesisvideoChannelARN)
	}
	if len(_kinesisvideoCurrentVersion) > 0 {
		input.CurrentVersion = aws.String(_kinesisvideoCurrentVersion)
	}

	if resp, err := client.DeleteSignalingChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Kinesis video stream and the data contained in the stream.
// This method marks the stream for deletion, and makes the data in the stream
// inaccessible immediately.
//
// To ensure that you have the latest version of the stream before deleting it,
// you can specify the stream version. Kinesis Video Streams assigns a version to
// each stream. When you update a stream, Kinesis Video Streams assigns a new
// version number. To get the latest stream version, use the DescribeStream API.
//
// This operation requires permission for the KinesisVideo:DeleteStream action.
func kinesisvideo_DeleteStream(cfg aws.Config, client *kinesisvideo.Client) {
	input := &kinesisvideo.DeleteStreamInput{
		// StreamARN: *string, // Required
	}

	if len(_kinesisvideoStreamARN) > 0 {
		input.StreamARN = aws.String(_kinesisvideoStreamARN)
	}
	if len(_kinesisvideoCurrentVersion) > 0 {
		input.CurrentVersion = aws.String(_kinesisvideoCurrentVersion)
	}

	if resp, err := client.DeleteStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a stream’s edge configuration that was set using the
// StartEdgeConfigurationUpdate API and the latest status of the edge agent's
// recorder and uploader jobs. Use this API to get the status of the configuration
// to determine if the configuration is in sync with the Edge Agent. Use this API
// to evaluate the health of the Edge Agent.
func kinesisvideo_DescribeEdgeConfiguration(cfg aws.Config, client *kinesisvideo.Client) {
	input := &kinesisvideo.DescribeEdgeConfigurationInput{}

	if len(_kinesisvideoStreamARN) > 0 {
		input.StreamARN = aws.String(_kinesisvideoStreamARN)
	}
	if len(_kinesisvideoStreamName) > 0 {
		input.StreamName = aws.String(_kinesisvideoStreamName)
	}

	if resp, err := client.DescribeEdgeConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the ImageGenerationConfiguration for a given Kinesis video stream.
func kinesisvideo_DescribeImageGenerationConfiguration(cfg aws.Config, client *kinesisvideo.Client) {
	input := &kinesisvideo.DescribeImageGenerationConfigurationInput{}

	if len(_kinesisvideoStreamARN) > 0 {
		input.StreamARN = aws.String(_kinesisvideoStreamARN)
	}
	if len(_kinesisvideoStreamName) > 0 {
		input.StreamName = aws.String(_kinesisvideoStreamName)
	}

	if resp, err := client.DescribeImageGenerationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the most current information about the stream. The streamName or
// streamARN should be provided in the input.
func kinesisvideo_DescribeMappedResourceConfiguration(cfg aws.Config, client *kinesisvideo.Client) {
	input := &kinesisvideo.DescribeMappedResourceConfigurationInput{}

	if len(_kinesisvideoMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _kinesisvideoMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_kinesisvideoNextToken) > 0 {
		input.NextToken = aws.String(_kinesisvideoNextToken)
	}
	if len(_kinesisvideoStreamARN) > 0 {
		input.StreamARN = aws.String(_kinesisvideoStreamARN)
	}
	if len(_kinesisvideoStreamName) > 0 {
		input.StreamName = aws.String(_kinesisvideoStreamName)
	}

	if disablePaginator() {
		if resp, err := client.DescribeMappedResourceConfiguration(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kinesisvideo.DescribeMappedResourceConfigurationOutput
	p := kinesisvideo.NewDescribeMappedResourceConfigurationPaginator(client, input)
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

// Returns the most current information about the channel. Specify the ChannelName
// or ChannelARN in the input.
func kinesisvideo_DescribeMediaStorageConfiguration(cfg aws.Config, client *kinesisvideo.Client) {
	input := &kinesisvideo.DescribeMediaStorageConfigurationInput{}

	if len(_kinesisvideoChannelARN) > 0 {
		input.ChannelARN = aws.String(_kinesisvideoChannelARN)
	}
	if len(_kinesisvideoChannelName) > 0 {
		input.ChannelName = aws.String(_kinesisvideoChannelName)
	}

	if resp, err := client.DescribeMediaStorageConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the NotificationConfiguration for a given Kinesis video stream.
func kinesisvideo_DescribeNotificationConfiguration(cfg aws.Config, client *kinesisvideo.Client) {
	input := &kinesisvideo.DescribeNotificationConfigurationInput{}

	if len(_kinesisvideoStreamARN) > 0 {
		input.StreamARN = aws.String(_kinesisvideoStreamARN)
	}
	if len(_kinesisvideoStreamName) > 0 {
		input.StreamName = aws.String(_kinesisvideoStreamName)
	}

	if resp, err := client.DescribeNotificationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the most current information about the signaling channel. You must
// specify either the name or the Amazon Resource Name (ARN) of the channel that
// you want to describe.
func kinesisvideo_DescribeSignalingChannel(cfg aws.Config, client *kinesisvideo.Client) {
	input := &kinesisvideo.DescribeSignalingChannelInput{}

	if len(_kinesisvideoChannelARN) > 0 {
		input.ChannelARN = aws.String(_kinesisvideoChannelARN)
	}
	if len(_kinesisvideoChannelName) > 0 {
		input.ChannelName = aws.String(_kinesisvideoChannelName)
	}

	if resp, err := client.DescribeSignalingChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the most current information about the specified stream. You must
// specify either the StreamName or the StreamARN .
func kinesisvideo_DescribeStream(cfg aws.Config, client *kinesisvideo.Client) {
	input := &kinesisvideo.DescribeStreamInput{}

	if len(_kinesisvideoStreamARN) > 0 {
		input.StreamARN = aws.String(_kinesisvideoStreamARN)
	}
	if len(_kinesisvideoStreamName) > 0 {
		input.StreamName = aws.String(_kinesisvideoStreamName)
	}

	if resp, err := client.DescribeStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the current storage configuration for the specified Kinesis video
// stream.
//
// In the request, you must specify either the StreamName or the StreamARN .
//
// You must have permissions for the
// KinesisVideo:DescribeStreamStorageConfiguration action.
func kinesisvideo_DescribeStreamStorageConfiguration(cfg aws.Config, client *kinesisvideo.Client) {
	input := &kinesisvideo.DescribeStreamStorageConfigurationInput{}

	if len(_kinesisvideoStreamARN) > 0 {
		input.StreamARN = aws.String(_kinesisvideoStreamARN)
	}
	if len(_kinesisvideoStreamName) > 0 {
		input.StreamName = aws.String(_kinesisvideoStreamName)
	}

	if resp, err := client.DescribeStreamStorageConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an endpoint for a specified stream for either reading or writing. Use this
// endpoint in your application to read from the specified stream (using the
// GetMedia or GetMediaForFragmentList operations) or write to it (using the
// PutMedia operation).
//
// The returned endpoint does not have the API name appended. The client needs to
// add the API name to the returned endpoint.
//
// In the request, specify the stream either by StreamName or StreamARN .
func kinesisvideo_GetDataEndpoint(cfg aws.Config, client *kinesisvideo.Client) {
	input := &kinesisvideo.GetDataEndpointInput{
		// APIName: types.APIName, // Required
	}

	if len(_kinesisvideoAPIName) > 0 {
		if err := assignInputField(input, "APIName", _kinesisvideoAPIName); err != nil {
			log.Errorf("invalid --api-name: %s", err.Error())
			return
		}
	}
	if len(_kinesisvideoStreamARN) > 0 {
		input.StreamARN = aws.String(_kinesisvideoStreamARN)
	}
	if len(_kinesisvideoStreamName) > 0 {
		input.StreamName = aws.String(_kinesisvideoStreamName)
	}

	if resp, err := client.GetDataEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides an endpoint for the specified signaling channel to send and receive
// messages. This API uses the SingleMasterChannelEndpointConfiguration input
// parameter, which consists of the Protocols and Role properties.
//
// Protocols is used to determine the communication mechanism. For example, if you
// specify WSS as the protocol, this API produces a secure websocket endpoint. If
// you specify HTTPS as the protocol, this API generates an HTTPS endpoint. If you
// specify WEBRTC as the protocol, but the signaling channel isn't configured for
// ingestion, you will receive the error InvalidArgumentException .
//
// Role determines the messaging permissions. A MASTER role results in this API
// generating an endpoint that a client can use to communicate with any of the
// viewers on the channel. A VIEWER role results in this API generating an
// endpoint that a client can use to communicate only with a MASTER .
func kinesisvideo_GetSignalingChannelEndpoint(cfg aws.Config, client *kinesisvideo.Client) {
	input := &kinesisvideo.GetSignalingChannelEndpointInput{
		// ChannelARN: *string, // Required
	}

	if len(_kinesisvideoChannelARN) > 0 {
		input.ChannelARN = aws.String(_kinesisvideoChannelARN)
	}
	if len(_kinesisvideoSingleMasterChannelEndpointConfiguration) > 0 {
		if err := assignInputField(input, "SingleMasterChannelEndpointConfiguration", _kinesisvideoSingleMasterChannelEndpointConfiguration); err != nil {
			log.Errorf("invalid --single-master-channel-endpoint-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetSignalingChannelEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns an array of edge configurations associated with the specified Edge
// Agent.
//
// In the request, you must specify the Edge Agent HubDeviceArn .
func kinesisvideo_ListEdgeAgentConfigurations(cfg aws.Config, client *kinesisvideo.Client) {
	input := &kinesisvideo.ListEdgeAgentConfigurationsInput{
		// HubDeviceArn: *string, // Required
	}

	if len(_kinesisvideoHubDeviceArn) > 0 {
		input.HubDeviceArn = aws.String(_kinesisvideoHubDeviceArn)
	}
	if len(_kinesisvideoMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _kinesisvideoMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_kinesisvideoNextToken) > 0 {
		input.NextToken = aws.String(_kinesisvideoNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEdgeAgentConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kinesisvideo.ListEdgeAgentConfigurationsOutput
	p := kinesisvideo.NewListEdgeAgentConfigurationsPaginator(client, input)
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

// Returns an array of ChannelInfo objects. Each object describes a signaling
// channel. To retrieve only those channels that satisfy a specific condition, you
// can specify a ChannelNameCondition .
func kinesisvideo_ListSignalingChannels(cfg aws.Config, client *kinesisvideo.Client) {
	input := &kinesisvideo.ListSignalingChannelsInput{}

	if len(_kinesisvideoChannelNameCondition) > 0 {
		if err := assignInputField(input, "ChannelNameCondition", _kinesisvideoChannelNameCondition); err != nil {
			log.Errorf("invalid --channel-name-condition: %s", err.Error())
			return
		}
	}
	if len(_kinesisvideoMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _kinesisvideoMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_kinesisvideoNextToken) > 0 {
		input.NextToken = aws.String(_kinesisvideoNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSignalingChannels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kinesisvideo.ListSignalingChannelsOutput
	p := kinesisvideo.NewListSignalingChannelsPaginator(client, input)
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

// Returns an array of StreamInfo objects. Each object describes a stream. To
// retrieve only streams that satisfy a specific condition, you can specify a
// StreamNameCondition .
func kinesisvideo_ListStreams(cfg aws.Config, client *kinesisvideo.Client) {
	input := &kinesisvideo.ListStreamsInput{}

	if len(_kinesisvideoMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _kinesisvideoMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_kinesisvideoNextToken) > 0 {
		input.NextToken = aws.String(_kinesisvideoNextToken)
	}
	if len(_kinesisvideoStreamNameCondition) > 0 {
		if err := assignInputField(input, "StreamNameCondition", _kinesisvideoStreamNameCondition); err != nil {
			log.Errorf("invalid --stream-name-condition: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListStreams(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*kinesisvideo.ListStreamsOutput
	p := kinesisvideo.NewListStreamsPaginator(client, input)
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

// Returns a list of tags associated with the specified signaling channel.
func kinesisvideo_ListTagsForResource(cfg aws.Config, client *kinesisvideo.Client) {
	input := &kinesisvideo.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_kinesisvideoResourceARN) > 0 {
		input.ResourceARN = aws.String(_kinesisvideoResourceARN)
	}
	if len(_kinesisvideoNextToken) > 0 {
		input.NextToken = aws.String(_kinesisvideoNextToken)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of tags associated with the specified stream.
// In the request, you must specify either the StreamName or the StreamARN .
func kinesisvideo_ListTagsForStream(cfg aws.Config, client *kinesisvideo.Client) {
	input := &kinesisvideo.ListTagsForStreamInput{}

	if len(_kinesisvideoNextToken) > 0 {
		input.NextToken = aws.String(_kinesisvideoNextToken)
	}
	if len(_kinesisvideoStreamARN) > 0 {
		input.StreamARN = aws.String(_kinesisvideoStreamARN)
	}
	if len(_kinesisvideoStreamName) > 0 {
		input.StreamName = aws.String(_kinesisvideoStreamName)
	}

	if resp, err := client.ListTagsForStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// An asynchronous API that updates a stream’s existing edge configuration. The
// Kinesis Video Stream will sync the stream’s edge configuration with the Edge
// Agent IoT Greengrass component that runs on an IoT Hub Device, setup at your
// premise. The time to sync can vary and depends on the connectivity of the Hub
// Device. The SyncStatus will be updated as the edge configuration is
// acknowledged, and synced with the Edge Agent.
//
// If this API is invoked for the first time, a new edge configuration will be
// created for the stream, and the sync status will be set to SYNCING . You will
// have to wait for the sync status to reach a terminal state such as: IN_SYNC , or
// SYNC_FAILED , before using this API again. If you invoke this API during the
// syncing process, a ResourceInUseException will be thrown. The connectivity of
// the stream’s edge configuration and the Edge Agent will be retried for 15
// minutes. After 15 minutes, the status will transition into the SYNC_FAILED
// state.
//
// To move an edge configuration from one device to another, use DeleteEdgeConfiguration to delete the
// current edge configuration. You can then invoke StartEdgeConfigurationUpdate
// with an updated Hub Device ARN.
func kinesisvideo_StartEdgeConfigurationUpdate(cfg aws.Config, client *kinesisvideo.Client) {
	input := &kinesisvideo.StartEdgeConfigurationUpdateInput{
		// EdgeConfig: *types.EdgeConfig, // Required
	}

	if len(_kinesisvideoEdgeConfig) > 0 {
		if err := assignInputField(input, "EdgeConfig", _kinesisvideoEdgeConfig); err != nil {
			log.Errorf("invalid --edge-config: %s", err.Error())
			return
		}
	}
	if len(_kinesisvideoStreamARN) > 0 {
		input.StreamARN = aws.String(_kinesisvideoStreamARN)
	}
	if len(_kinesisvideoStreamName) > 0 {
		input.StreamName = aws.String(_kinesisvideoStreamName)
	}

	if resp, err := client.StartEdgeConfigurationUpdate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more tags to a signaling channel. A tag is a key-value pair (the
// value is optional) that you can define and assign to Amazon Web Services
// resources. If you specify a tag that already exists, the tag value is replaced
// with the value that you specify in the request. For more information, see [Using Cost Allocation Tags]in
// the Billing and Cost Management and Cost Management User Guide.
//
// [Using Cost Allocation Tags]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html
func kinesisvideo_TagResource(cfg aws.Config, client *kinesisvideo.Client) {
	input := &kinesisvideo.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_kinesisvideoResourceARN) > 0 {
		input.ResourceARN = aws.String(_kinesisvideoResourceARN)
	}
	if len(_kinesisvideoTags) > 0 {
		if err := assignInputField(input, "Tags", _kinesisvideoTags); err != nil {
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

// Adds one or more tags to a stream. A tag is a key-value pair (the value is
// optional) that you can define and assign to Amazon Web Services resources. If
// you specify a tag that already exists, the tag value is replaced with the value
// that you specify in the request. For more information, see [Using Cost Allocation Tags]in the Billing and
// Cost Management and Cost Management User Guide.
//
// You must provide either the StreamName or the StreamARN .
//
// This operation requires permission for the KinesisVideo:TagStream action.
//
// A Kinesis video stream can support up to 50 tags.
//
// [Using Cost Allocation Tags]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html
func kinesisvideo_TagStream(cfg aws.Config, client *kinesisvideo.Client) {
	input := &kinesisvideo.TagStreamInput{
		// Tags: map[string]string, // Required
	}

	if len(_kinesisvideoTags) > 0 {
		if err := assignInputField(input, "Tags", _kinesisvideoTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_kinesisvideoStreamARN) > 0 {
		input.StreamARN = aws.String(_kinesisvideoStreamARN)
	}
	if len(_kinesisvideoStreamName) > 0 {
		input.StreamName = aws.String(_kinesisvideoStreamName)
	}

	if resp, err := client.TagStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes one or more tags from a signaling channel. In the request, specify only
// a tag key or keys; don't specify the value. If you specify a tag key that does
// not exist, it's ignored.
func kinesisvideo_UntagResource(cfg aws.Config, client *kinesisvideo.Client) {
	input := &kinesisvideo.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeyList: []string, // Required
	}

	if len(_kinesisvideoResourceARN) > 0 {
		input.ResourceARN = aws.String(_kinesisvideoResourceARN)
	}
	if len(_kinesisvideoTagKeyList) > 0 {
		input.TagKeyList = append([]string(nil), _kinesisvideoTagKeyList...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes one or more tags from a stream. In the request, specify only a tag key
// or keys; don't specify the value. If you specify a tag key that does not exist,
// it's ignored.
//
// In the request, you must provide the StreamName or StreamARN .
func kinesisvideo_UntagStream(cfg aws.Config, client *kinesisvideo.Client) {
	input := &kinesisvideo.UntagStreamInput{
		// TagKeyList: []string, // Required
	}

	if len(_kinesisvideoTagKeyList) > 0 {
		input.TagKeyList = append([]string(nil), _kinesisvideoTagKeyList...)
	}
	if len(_kinesisvideoStreamARN) > 0 {
		input.StreamARN = aws.String(_kinesisvideoStreamARN)
	}
	if len(_kinesisvideoStreamName) > 0 {
		input.StreamName = aws.String(_kinesisvideoStreamName)
	}

	if resp, err := client.UntagStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Increases or decreases the stream's data retention period by the value that you
// specify. To indicate whether you want to increase or decrease the data retention
// period, specify the Operation parameter in the request body. In the request,
// you must specify either the StreamName or the StreamARN .
//
// This operation requires permission for the KinesisVideo:UpdateDataRetention
// action.
//
// Changing the data retention period affects the data in the stream as follows:
//
// - If the data retention period is increased, existing data is retained for
// the new retention period. For example, if the data retention period is increased
// from one hour to seven hours, all existing data is retained for seven hours.
//
// - If the data retention period is decreased, existing data is retained for
// the new retention period. For example, if the data retention period is decreased
// from seven hours to one hour, all existing data is retained for one hour, and
// any data older than one hour is deleted immediately.
func kinesisvideo_UpdateDataRetention(cfg aws.Config, client *kinesisvideo.Client) {
	input := &kinesisvideo.UpdateDataRetentionInput{
		// CurrentVersion: *string, // Required
		// DataRetentionChangeInHours: *int32, // Required
		// Operation: types.UpdateDataRetentionOperation, // Required
	}

	if len(_kinesisvideoCurrentVersion) > 0 {
		input.CurrentVersion = aws.String(_kinesisvideoCurrentVersion)
	}
	if len(_kinesisvideoDataRetentionChangeInHours) > 0 {
		if err := assignInputField(input, "DataRetentionChangeInHours", _kinesisvideoDataRetentionChangeInHours); err != nil {
			log.Errorf("invalid --data-retention-change-in-hours: %s", err.Error())
			return
		}
	}
	if len(_kinesisvideoOperation) > 0 {
		if err := assignInputField(input, "Operation", _kinesisvideoOperation); err != nil {
			log.Errorf("invalid --operation: %s", err.Error())
			return
		}
	}
	if len(_kinesisvideoStreamARN) > 0 {
		input.StreamARN = aws.String(_kinesisvideoStreamARN)
	}
	if len(_kinesisvideoStreamName) > 0 {
		input.StreamName = aws.String(_kinesisvideoStreamName)
	}

	if resp, err := client.UpdateDataRetention(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the StreamInfo and ImageProcessingConfiguration fields.
func kinesisvideo_UpdateImageGenerationConfiguration(cfg aws.Config, client *kinesisvideo.Client) {
	input := &kinesisvideo.UpdateImageGenerationConfigurationInput{}

	if len(_kinesisvideoImageGenerationConfiguration) > 0 {
		if err := assignInputField(input, "ImageGenerationConfiguration", _kinesisvideoImageGenerationConfiguration); err != nil {
			log.Errorf("invalid --image-generation-configuration: %s", err.Error())
			return
		}
	}
	if len(_kinesisvideoStreamARN) > 0 {
		input.StreamARN = aws.String(_kinesisvideoStreamARN)
	}
	if len(_kinesisvideoStreamName) > 0 {
		input.StreamName = aws.String(_kinesisvideoStreamName)
	}

	if resp, err := client.UpdateImageGenerationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a SignalingChannel to a stream to store the media. There are two
// signaling modes that you can specify :
//
// - If StorageStatus is enabled, the data will be stored in the StreamARN
// provided. In order for WebRTC Ingestion to work, the stream must have data
// retention enabled.
//
// - If StorageStatus is disabled, no data will be stored, and the StreamARN
// parameter will not be needed.
//
// If StorageStatus is enabled, direct peer-to-peer (master-viewer) connections no
// longer occur. Peers connect directly to the storage session. You must call the
// JoinStorageSession API to trigger an SDP offer send and establish a connection
// between a peer and the storage session.
func kinesisvideo_UpdateMediaStorageConfiguration(cfg aws.Config, client *kinesisvideo.Client) {
	input := &kinesisvideo.UpdateMediaStorageConfigurationInput{
		// ChannelARN: *string, // Required
		// MediaStorageConfiguration: *types.MediaStorageConfiguration, // Required
	}

	if len(_kinesisvideoChannelARN) > 0 {
		input.ChannelARN = aws.String(_kinesisvideoChannelARN)
	}
	if len(_kinesisvideoMediaStorageConfiguration) > 0 {
		if err := assignInputField(input, "MediaStorageConfiguration", _kinesisvideoMediaStorageConfiguration); err != nil {
			log.Errorf("invalid --media-storage-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateMediaStorageConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the notification information for a stream.
func kinesisvideo_UpdateNotificationConfiguration(cfg aws.Config, client *kinesisvideo.Client) {
	input := &kinesisvideo.UpdateNotificationConfigurationInput{}

	if len(_kinesisvideoNotificationConfiguration) > 0 {
		if err := assignInputField(input, "NotificationConfiguration", _kinesisvideoNotificationConfiguration); err != nil {
			log.Errorf("invalid --notification-configuration: %s", err.Error())
			return
		}
	}
	if len(_kinesisvideoStreamARN) > 0 {
		input.StreamARN = aws.String(_kinesisvideoStreamARN)
	}
	if len(_kinesisvideoStreamName) > 0 {
		input.StreamName = aws.String(_kinesisvideoStreamName)
	}

	if resp, err := client.UpdateNotificationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the existing signaling channel. This is an asynchronous operation and
// takes time to complete.
//
// If the MessageTtlSeconds value is updated (either increased or reduced), it
// only applies to new messages sent via this channel after it's been updated.
// Existing messages are still expired as per the previous MessageTtlSeconds value.
func kinesisvideo_UpdateSignalingChannel(cfg aws.Config, client *kinesisvideo.Client) {
	input := &kinesisvideo.UpdateSignalingChannelInput{
		// ChannelARN: *string, // Required
		// CurrentVersion: *string, // Required
	}

	if len(_kinesisvideoChannelARN) > 0 {
		input.ChannelARN = aws.String(_kinesisvideoChannelARN)
	}
	if len(_kinesisvideoCurrentVersion) > 0 {
		input.CurrentVersion = aws.String(_kinesisvideoCurrentVersion)
	}
	if len(_kinesisvideoSingleMasterConfiguration) > 0 {
		if err := assignInputField(input, "SingleMasterConfiguration", _kinesisvideoSingleMasterConfiguration); err != nil {
			log.Errorf("invalid --single-master-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSignalingChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates stream metadata, such as the device name and media type.
// You must provide the stream name or the Amazon Resource Name (ARN) of the
// stream.
//
// To make sure that you have the latest version of the stream before updating it,
// you can specify the stream version. Kinesis Video Streams assigns a version to
// each stream. When you update a stream, Kinesis Video Streams assigns a new
// version number. To get the latest stream version, use the DescribeStream API.
//
// UpdateStream is an asynchronous operation, and takes time to complete.
func kinesisvideo_UpdateStream(cfg aws.Config, client *kinesisvideo.Client) {
	input := &kinesisvideo.UpdateStreamInput{
		// CurrentVersion: *string, // Required
	}

	if len(_kinesisvideoCurrentVersion) > 0 {
		input.CurrentVersion = aws.String(_kinesisvideoCurrentVersion)
	}
	if len(_kinesisvideoDeviceName) > 0 {
		input.DeviceName = aws.String(_kinesisvideoDeviceName)
	}
	if len(_kinesisvideoMediaType) > 0 {
		input.MediaType = aws.String(_kinesisvideoMediaType)
	}
	if len(_kinesisvideoStreamARN) > 0 {
		input.StreamARN = aws.String(_kinesisvideoStreamARN)
	}
	if len(_kinesisvideoStreamName) > 0 {
		input.StreamName = aws.String(_kinesisvideoStreamName)
	}

	if resp, err := client.UpdateStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the storage configuration for an existing Kinesis video stream.
// This operation allows you to modify the storage tier settings for a stream,
// enabling you to optimize storage costs and performance based on your access
// patterns.
//
// UpdateStreamStorageConfiguration is an asynchronous operation.
//
// You must have permissions for the KinesisVideo:UpdateStreamStorageConfiguration
// action.
func kinesisvideo_UpdateStreamStorageConfiguration(cfg aws.Config, client *kinesisvideo.Client) {
	input := &kinesisvideo.UpdateStreamStorageConfigurationInput{
		// CurrentVersion: *string, // Required
		// StreamStorageConfiguration: *types.StreamStorageConfiguration, // Required
	}

	if len(_kinesisvideoCurrentVersion) > 0 {
		input.CurrentVersion = aws.String(_kinesisvideoCurrentVersion)
	}
	if len(_kinesisvideoStreamStorageConfiguration) > 0 {
		if err := assignInputField(input, "StreamStorageConfiguration", _kinesisvideoStreamStorageConfiguration); err != nil {
			log.Errorf("invalid --stream-storage-configuration: %s", err.Error())
			return
		}
	}
	if len(_kinesisvideoStreamARN) > 0 {
		input.StreamARN = aws.String(_kinesisvideoStreamARN)
	}
	if len(_kinesisvideoStreamName) > 0 {
		input.StreamName = aws.String(_kinesisvideoStreamName)
	}

	if resp, err := client.UpdateStreamStorageConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_kinesisvideoCmd)
	_kinesisvideoCmd.Flags().SortFlags = false

	_kinesisvideoCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_kinesisvideoCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_kinesisvideoCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_kinesisvideoCmd.Flags().StringVarP(&_kinesisvideoAPIName, "api-name", "", "", "API Name")
	_kinesisvideoCmd.Flags().StringVarP(&_kinesisvideoChannelARN, "channel-arn", "", "", "Channel ARN")
	_kinesisvideoCmd.Flags().StringVarP(&_kinesisvideoChannelName, "channel-name", "", "", "Channel Name")
	_kinesisvideoCmd.Flags().StringVarP(&_kinesisvideoChannelNameCondition, "channel-name-condition", "", "", "Channel Name Condition")
	_kinesisvideoCmd.Flags().StringVarP(&_kinesisvideoChannelType, "channel-type", "", "", "Channel Type")
	_kinesisvideoCmd.Flags().StringVarP(&_kinesisvideoCurrentVersion, "current-version", "", "", "Current Version")
	_kinesisvideoCmd.Flags().StringVarP(&_kinesisvideoDataRetentionChangeInHours, "data-retention-change-in-hours", "", "", "Data Retention Change In Hours")
	_kinesisvideoCmd.Flags().StringVarP(&_kinesisvideoDataRetentionInHours, "data-retention-in-hours", "", "", "Data Retention In Hours")
	_kinesisvideoCmd.Flags().StringVarP(&_kinesisvideoDeviceName, "device-name", "", "", "Device Name")
	_kinesisvideoCmd.Flags().StringVarP(&_kinesisvideoEdgeConfig, "edge-config", "", "", "Edge Config")
	_kinesisvideoCmd.Flags().StringVarP(&_kinesisvideoHubDeviceArn, "hub-device-arn", "", "", "Hub Device ARN")
	_kinesisvideoCmd.Flags().StringVarP(&_kinesisvideoImageGenerationConfiguration, "image-generation-configuration", "", "", "Image Generation Configuration")
	_kinesisvideoCmd.Flags().StringVarP(&_kinesisvideoKmsKeyId, "kms-key-id", "", "", "KMS Key ID")
	_kinesisvideoCmd.Flags().StringVarP(&_kinesisvideoMaxResults, "max-results", "", "", "Max Results")
	_kinesisvideoCmd.Flags().StringVarP(&_kinesisvideoMediaStorageConfiguration, "media-storage-configuration", "", "", "Media Storage Configuration")
	_kinesisvideoCmd.Flags().StringVarP(&_kinesisvideoMediaType, "media-type", "", "", "Media Type")
	_kinesisvideoCmd.Flags().StringVarP(&_kinesisvideoNextToken, "next-token", "", "", "Next Token")
	_kinesisvideoCmd.Flags().StringVarP(&_kinesisvideoNotificationConfiguration, "notification-configuration", "", "", "Notification Configuration")
	_kinesisvideoCmd.Flags().StringVarP(&_kinesisvideoOperation, "operation", "", "", "Operation")
	_kinesisvideoCmd.Flags().StringVarP(&_kinesisvideoResourceARN, "resource-arn", "", "", "Resource ARN")
	_kinesisvideoCmd.Flags().StringVarP(&_kinesisvideoSingleMasterChannelEndpointConfiguration, "single-master-channel-endpoint-configuration", "", "", "Single Master Channel Endpoint Configuration")
	_kinesisvideoCmd.Flags().StringVarP(&_kinesisvideoSingleMasterConfiguration, "single-master-configuration", "", "", "Single Master Configuration")
	_kinesisvideoCmd.Flags().StringVarP(&_kinesisvideoStreamARN, "stream-arn", "", "", "Stream ARN")
	_kinesisvideoCmd.Flags().StringVarP(&_kinesisvideoStreamName, "stream-name", "", "", "Stream Name")
	_kinesisvideoCmd.Flags().StringVarP(&_kinesisvideoStreamNameCondition, "stream-name-condition", "", "", "Stream Name Condition")
	_kinesisvideoCmd.Flags().StringVarP(&_kinesisvideoStreamStorageConfiguration, "stream-storage-configuration", "", "", "Stream Storage Configuration")
	_kinesisvideoCmd.Flags().StringSliceVarP(&_kinesisvideoTagKeyList, "tag-key-list", "", nil, "Tag Key List")
	_kinesisvideoCmd.Flags().StringVarP(&_kinesisvideoTags, "tags", "", "", "Tags")

	_kinesisvideoCmd.Flags().BoolVarP(&_kinesisvideoCreateSignalingChannel, "create-signaling-channel", "", false, "Create Signaling Channel")
	_kinesisvideoCmd.Flags().BoolVarP(&_kinesisvideoCreateStream, "create-stream", "", false, "Create Stream")
	_kinesisvideoCmd.Flags().BoolVarP(&_kinesisvideoDeleteEdgeConfiguration, "delete-edge-configuration", "", false, "Delete Edge Configuration")
	_kinesisvideoCmd.Flags().BoolVarP(&_kinesisvideoDeleteSignalingChannel, "delete-signaling-channel", "", false, "Delete Signaling Channel")
	_kinesisvideoCmd.Flags().BoolVarP(&_kinesisvideoDeleteStream, "delete-stream", "", false, "Delete Stream")
	_kinesisvideoCmd.Flags().BoolVarP(&_kinesisvideoDescribeEdgeConfiguration, "describe-edge-configuration", "", false, "Describe Edge Configuration")
	_kinesisvideoCmd.Flags().BoolVarP(&_kinesisvideoDescribeImageGenerationConfiguration, "describe-image-generation-configuration", "", false, "Describe Image Generation Configuration")
	_kinesisvideoCmd.Flags().BoolVarP(&_kinesisvideoDescribeMappedResourceConfiguration, "describe-mapped-resource-configuration", "", false, "Describe Mapped Resource Configuration")
	_kinesisvideoCmd.Flags().BoolVarP(&_kinesisvideoDescribeMediaStorageConfiguration, "describe-media-storage-configuration", "", false, "Describe Media Storage Configuration")
	_kinesisvideoCmd.Flags().BoolVarP(&_kinesisvideoDescribeNotificationConfiguration, "describe-notification-configuration", "", false, "Describe Notification Configuration")
	_kinesisvideoCmd.Flags().BoolVarP(&_kinesisvideoDescribeSignalingChannel, "describe-signaling-channel", "", false, "Describe Signaling Channel")
	_kinesisvideoCmd.Flags().BoolVarP(&_kinesisvideoDescribeStream, "describe-stream", "", false, "Describe Stream")
	_kinesisvideoCmd.Flags().BoolVarP(&_kinesisvideoDescribeStreamStorageConfiguration, "describe-stream-storage-configuration", "", false, "Describe Stream Storage Configuration")
	_kinesisvideoCmd.Flags().BoolVarP(&_kinesisvideoGetDataEndpoint, "get-data-endpoint", "", false, "Get Data Endpoint")
	_kinesisvideoCmd.Flags().BoolVarP(&_kinesisvideoGetSignalingChannelEndpoint, "get-signaling-channel-endpoint", "", false, "Get Signaling Channel Endpoint")
	_kinesisvideoCmd.Flags().BoolVarP(&_kinesisvideoListEdgeAgentConfigurations, "list-edge-agent-configurations", "", false, "List Edge Agent Configurations")
	_kinesisvideoCmd.Flags().BoolVarP(&_kinesisvideoListSignalingChannels, "list-signaling-channels", "", false, "List Signaling Channels")
	_kinesisvideoCmd.Flags().BoolVarP(&_kinesisvideoListStreams, "list-streams", "", false, "List Streams")
	_kinesisvideoCmd.Flags().BoolVarP(&_kinesisvideoListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_kinesisvideoCmd.Flags().BoolVarP(&_kinesisvideoListTagsForStream, "list-tags-for-stream", "", false, "List Tags For Stream")
	_kinesisvideoCmd.Flags().BoolVarP(&_kinesisvideoStartEdgeConfigurationUpdate, "start-edge-configuration-update", "", false, "Start Edge Configuration Update")
	_kinesisvideoCmd.Flags().BoolVarP(&_kinesisvideoTagResource, "tag-resource", "", false, "Tag Resource")
	_kinesisvideoCmd.Flags().BoolVarP(&_kinesisvideoTagStream, "tag-stream", "", false, "Tag Stream")
	_kinesisvideoCmd.Flags().BoolVarP(&_kinesisvideoUntagResource, "untag-resource", "", false, "Untag Resource")
	_kinesisvideoCmd.Flags().BoolVarP(&_kinesisvideoUntagStream, "untag-stream", "", false, "Untag Stream")
	_kinesisvideoCmd.Flags().BoolVarP(&_kinesisvideoUpdateDataRetention, "update-data-retention", "", false, "Update Data Retention")
	_kinesisvideoCmd.Flags().BoolVarP(&_kinesisvideoUpdateImageGenerationConfiguration, "update-image-generation-configuration", "", false, "Update Image Generation Configuration")
	_kinesisvideoCmd.Flags().BoolVarP(&_kinesisvideoUpdateMediaStorageConfiguration, "update-media-storage-configuration", "", false, "Update Media Storage Configuration")
	_kinesisvideoCmd.Flags().BoolVarP(&_kinesisvideoUpdateNotificationConfiguration, "update-notification-configuration", "", false, "Update Notification Configuration")
	_kinesisvideoCmd.Flags().BoolVarP(&_kinesisvideoUpdateSignalingChannel, "update-signaling-channel", "", false, "Update Signaling Channel")
	_kinesisvideoCmd.Flags().BoolVarP(&_kinesisvideoUpdateStream, "update-stream", "", false, "Update Stream")
	_kinesisvideoCmd.Flags().BoolVarP(&_kinesisvideoUpdateStreamStorageConfiguration, "update-stream-storage-configuration", "", false, "Update Stream Storage Configuration")

}
