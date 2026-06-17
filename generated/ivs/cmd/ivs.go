package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ivs"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ivsCmd represents the ivs command
var _ivsCmd = &cobra.Command{
	Use:   "ivs",
	Short: "AWS ivs CLI",
	Run: func(cmd *cobra.Command, args []string) {
		_awsOutput = resolveAWSOutput(_awsProfile, cmd.Flags().Changed("output"))
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := ivs.NewFromConfig(cfg)
		if _ivsBatchGetChannel {
			ivs_BatchGetChannel(cfg, client)
			return
		}
		if _ivsBatchGetStreamKey {
			ivs_BatchGetStreamKey(cfg, client)
			return
		}
		if _ivsBatchStartViewerSessionRevocation {
			ivs_BatchStartViewerSessionRevocation(cfg, client)
			return
		}
		if _ivsCreateChannel {
			ivs_CreateChannel(cfg, client)
			return
		}
		if _ivsCreatePlaybackRestrictionPolicy {
			ivs_CreatePlaybackRestrictionPolicy(cfg, client)
			return
		}
		if _ivsCreateRecordingConfiguration {
			ivs_CreateRecordingConfiguration(cfg, client)
			return
		}
		if _ivsCreateStreamKey {
			ivs_CreateStreamKey(cfg, client)
			return
		}
		if _ivsDeleteChannel {
			ivs_DeleteChannel(cfg, client)
			return
		}
		if _ivsDeletePlaybackKeyPair {
			ivs_DeletePlaybackKeyPair(cfg, client)
			return
		}
		if _ivsDeletePlaybackRestrictionPolicy {
			ivs_DeletePlaybackRestrictionPolicy(cfg, client)
			return
		}
		if _ivsDeleteRecordingConfiguration {
			ivs_DeleteRecordingConfiguration(cfg, client)
			return
		}
		if _ivsDeleteStreamKey {
			ivs_DeleteStreamKey(cfg, client)
			return
		}
		if _ivsGetChannel {
			ivs_GetChannel(cfg, client)
			return
		}
		if _ivsGetPlaybackKeyPair {
			ivs_GetPlaybackKeyPair(cfg, client)
			return
		}
		if _ivsGetPlaybackRestrictionPolicy {
			ivs_GetPlaybackRestrictionPolicy(cfg, client)
			return
		}
		if _ivsGetRecordingConfiguration {
			ivs_GetRecordingConfiguration(cfg, client)
			return
		}
		if _ivsGetStream {
			ivs_GetStream(cfg, client)
			return
		}
		if _ivsGetStreamKey {
			ivs_GetStreamKey(cfg, client)
			return
		}
		if _ivsGetStreamSession {
			ivs_GetStreamSession(cfg, client)
			return
		}
		if _ivsImportPlaybackKeyPair {
			ivs_ImportPlaybackKeyPair(cfg, client)
			return
		}
		if _ivsListChannels {
			ivs_ListChannels(cfg, client)
			return
		}
		if _ivsListPlaybackKeyPairs {
			ivs_ListPlaybackKeyPairs(cfg, client)
			return
		}
		if _ivsListPlaybackRestrictionPolicies {
			ivs_ListPlaybackRestrictionPolicies(cfg, client)
			return
		}
		if _ivsListRecordingConfigurations {
			ivs_ListRecordingConfigurations(cfg, client)
			return
		}
		if _ivsListStreamKeys {
			ivs_ListStreamKeys(cfg, client)
			return
		}
		if _ivsListStreamSessions {
			ivs_ListStreamSessions(cfg, client)
			return
		}
		if _ivsListStreams {
			ivs_ListStreams(cfg, client)
			return
		}
		if _ivsListTagsForResource {
			ivs_ListTagsForResource(cfg, client)
			return
		}
		if _ivsPutMetadata {
			ivs_PutMetadata(cfg, client)
			return
		}
		if _ivsStartViewerSessionRevocation {
			ivs_StartViewerSessionRevocation(cfg, client)
			return
		}
		if _ivsStopStream {
			ivs_StopStream(cfg, client)
			return
		}
		if _ivsTagResource {
			ivs_TagResource(cfg, client)
			return
		}
		if _ivsUntagResource {
			ivs_UntagResource(cfg, client)
			return
		}
		if _ivsUpdateChannel {
			ivs_UpdateChannel(cfg, client)
			return
		}
		if _ivsUpdatePlaybackRestrictionPolicy {
			ivs_UpdatePlaybackRestrictionPolicy(cfg, client)
			return
		}

	},
}

var (
	_ivsBatchGetChannel                   bool
	_ivsBatchGetStreamKey                 bool
	_ivsBatchStartViewerSessionRevocation bool
	_ivsCreateChannel                     bool
	_ivsCreatePlaybackRestrictionPolicy   bool
	_ivsCreateRecordingConfiguration      bool
	_ivsCreateStreamKey                   bool
	_ivsDeleteChannel                     bool
	_ivsDeletePlaybackKeyPair             bool
	_ivsDeletePlaybackRestrictionPolicy   bool
	_ivsDeleteRecordingConfiguration      bool
	_ivsDeleteStreamKey                   bool
	_ivsGetChannel                        bool
	_ivsGetPlaybackKeyPair                bool
	_ivsGetPlaybackRestrictionPolicy      bool
	_ivsGetRecordingConfiguration         bool
	_ivsGetStream                         bool
	_ivsGetStreamKey                      bool
	_ivsGetStreamSession                  bool
	_ivsImportPlaybackKeyPair             bool
	_ivsListChannels                      bool
	_ivsListPlaybackKeyPairs              bool
	_ivsListPlaybackRestrictionPolicies   bool
	_ivsListRecordingConfigurations       bool
	_ivsListStreamKeys                    bool
	_ivsListStreamSessions                bool
	_ivsListStreams                       bool
	_ivsListTagsForResource               bool
	_ivsPutMetadata                       bool
	_ivsStartViewerSessionRevocation      bool
	_ivsStopStream                        bool
	_ivsTagResource                       bool
	_ivsUntagResource                     bool
	_ivsUpdateChannel                     bool
	_ivsUpdatePlaybackRestrictionPolicy   bool

	_ivsAllowedCountries                       []string
	_ivsAllowedOrigins                         []string
	_ivsArn                                    string
	_ivsArns                                   []string
	_ivsAuthorized                             string
	_ivsChannelArn                             string
	_ivsContainerFormat                        string
	_ivsDestinationConfiguration               string
	_ivsEnableStrictOriginEnforcement          string
	_ivsFilterBy                               string
	_ivsFilterByName                           string
	_ivsFilterByPlaybackRestrictionPolicyArn   string
	_ivsFilterByRecordingConfigurationArn      string
	_ivsInsecureIngest                         string
	_ivsLatencyMode                            string
	_ivsMaxResults                             string
	_ivsMetadata                               string
	_ivsMultitrackInputConfiguration           string
	_ivsName                                   string
	_ivsNextToken                              string
	_ivsPlaybackRestrictionPolicyArn           string
	_ivsPreset                                 string
	_ivsPublicKeyMaterial                      string
	_ivsRecordingConfigurationArn              string
	_ivsRecordingReconnectWindowSeconds        string
	_ivsRenditionConfiguration                 string
	_ivsResourceArn                            string
	_ivsStreamId                               string
	_ivsTagKeys                                []string
	_ivsTags                                   string
	_ivsThumbnailConfiguration                 string
	_ivsType                                   string
	_ivsViewerId                               string
	_ivsViewerSessionVersionsLessThanOrEqualTo string
	_ivsViewerSessions                         string
)

// Performs GetChannel on multiple ARNs simultaneously.
func ivs_BatchGetChannel(cfg aws.Config, client *ivs.Client) {
	input := &ivs.BatchGetChannelInput{
		// Arns: []string, // Required
	}

	if len(_ivsArns) > 0 {
		input.Arns = append([]string(nil), _ivsArns...)
	}

	if resp, err := client.BatchGetChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Performs GetStreamKey on multiple ARNs simultaneously.
func ivs_BatchGetStreamKey(cfg aws.Config, client *ivs.Client) {
	input := &ivs.BatchGetStreamKeyInput{
		// Arns: []string, // Required
	}

	if len(_ivsArns) > 0 {
		input.Arns = append([]string(nil), _ivsArns...)
	}

	if resp, err := client.BatchGetStreamKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Performs StartViewerSessionRevocation on multiple channel ARN and viewer ID pairs simultaneously.
func ivs_BatchStartViewerSessionRevocation(cfg aws.Config, client *ivs.Client) {
	input := &ivs.BatchStartViewerSessionRevocationInput{
		// ViewerSessions: []types.BatchStartViewerSessionRevocationViewerSession, // Required
	}

	if len(_ivsViewerSessions) > 0 {
		if err := assignInputField(input, "ViewerSessions", _ivsViewerSessions); err != nil {
			log.Errorf("invalid --viewer-sessions: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchStartViewerSessionRevocation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new channel and an associated stream key to start streaming.
func ivs_CreateChannel(cfg aws.Config, client *ivs.Client) {
	input := &ivs.CreateChannelInput{}

	if len(_ivsAuthorized) > 0 {
		if err := assignInputField(input, "Authorized", _ivsAuthorized); err != nil {
			log.Errorf("invalid --authorized: %s", err.Error())
			return
		}
	}
	if len(_ivsContainerFormat) > 0 {
		if err := assignInputField(input, "ContainerFormat", _ivsContainerFormat); err != nil {
			log.Errorf("invalid --container-format: %s", err.Error())
			return
		}
	}
	if len(_ivsInsecureIngest) > 0 {
		if err := assignInputField(input, "InsecureIngest", _ivsInsecureIngest); err != nil {
			log.Errorf("invalid --insecure-ingest: %s", err.Error())
			return
		}
	}
	if len(_ivsLatencyMode) > 0 {
		if err := assignInputField(input, "LatencyMode", _ivsLatencyMode); err != nil {
			log.Errorf("invalid --latency-mode: %s", err.Error())
			return
		}
	}
	if len(_ivsMultitrackInputConfiguration) > 0 {
		if err := assignInputField(input, "MultitrackInputConfiguration", _ivsMultitrackInputConfiguration); err != nil {
			log.Errorf("invalid --multitrack-input-configuration: %s", err.Error())
			return
		}
	}
	if len(_ivsName) > 0 {
		input.Name = aws.String(_ivsName)
	}
	if len(_ivsPlaybackRestrictionPolicyArn) > 0 {
		input.PlaybackRestrictionPolicyArn = aws.String(_ivsPlaybackRestrictionPolicyArn)
	}
	if len(_ivsPreset) > 0 {
		if err := assignInputField(input, "Preset", _ivsPreset); err != nil {
			log.Errorf("invalid --preset: %s", err.Error())
			return
		}
	}
	if len(_ivsRecordingConfigurationArn) > 0 {
		input.RecordingConfigurationArn = aws.String(_ivsRecordingConfigurationArn)
	}
	if len(_ivsTags) > 0 {
		if err := assignInputField(input, "Tags", _ivsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_ivsType) > 0 {
		if err := assignInputField(input, "Type", _ivsType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new playback restriction policy, for constraining playback by
// countries and/or origins.
func ivs_CreatePlaybackRestrictionPolicy(cfg aws.Config, client *ivs.Client) {
	input := &ivs.CreatePlaybackRestrictionPolicyInput{}

	if len(_ivsAllowedCountries) > 0 {
		input.AllowedCountries = append([]string(nil), _ivsAllowedCountries...)
	}
	if len(_ivsAllowedOrigins) > 0 {
		input.AllowedOrigins = append([]string(nil), _ivsAllowedOrigins...)
	}
	if len(_ivsEnableStrictOriginEnforcement) > 0 {
		if err := assignInputField(input, "EnableStrictOriginEnforcement", _ivsEnableStrictOriginEnforcement); err != nil {
			log.Errorf("invalid --enable-strict-origin-enforcement: %s", err.Error())
			return
		}
	}
	if len(_ivsName) > 0 {
		input.Name = aws.String(_ivsName)
	}
	if len(_ivsTags) > 0 {
		if err := assignInputField(input, "Tags", _ivsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePlaybackRestrictionPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new recording configuration, used to enable recording to Amazon S3.
// Known issue: In the us-east-1 region, if you use the Amazon Web Services CLI to
// create a recording configuration, it returns success even if the S3 bucket is in
// a different region. In this case, the state of the recording configuration is
// CREATE_FAILED (instead of ACTIVE ). (In other regions, the CLI correctly returns
// failure if the bucket is in a different region.)
//
// Workaround: Ensure that your S3 bucket is in the same region as the recording
// configuration. If you create a recording configuration in a different region as
// your S3 bucket, delete that recording configuration and create a new one with an
// S3 bucket from the correct region.
func ivs_CreateRecordingConfiguration(cfg aws.Config, client *ivs.Client) {
	input := &ivs.CreateRecordingConfigurationInput{
		// DestinationConfiguration: *types.DestinationConfiguration, // Required
	}

	if len(_ivsDestinationConfiguration) > 0 {
		if err := assignInputField(input, "DestinationConfiguration", _ivsDestinationConfiguration); err != nil {
			log.Errorf("invalid --destination-configuration: %s", err.Error())
			return
		}
	}
	if len(_ivsName) > 0 {
		input.Name = aws.String(_ivsName)
	}
	if len(_ivsRecordingReconnectWindowSeconds) > 0 {
		if err := assignInputField(input, "RecordingReconnectWindowSeconds", _ivsRecordingReconnectWindowSeconds); err != nil {
			log.Errorf("invalid --recording-reconnect-window-seconds: %s", err.Error())
			return
		}
	}
	if len(_ivsRenditionConfiguration) > 0 {
		if err := assignInputField(input, "RenditionConfiguration", _ivsRenditionConfiguration); err != nil {
			log.Errorf("invalid --rendition-configuration: %s", err.Error())
			return
		}
	}
	if len(_ivsTags) > 0 {
		if err := assignInputField(input, "Tags", _ivsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_ivsThumbnailConfiguration) > 0 {
		if err := assignInputField(input, "ThumbnailConfiguration", _ivsThumbnailConfiguration); err != nil {
			log.Errorf("invalid --thumbnail-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRecordingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a stream key, used to initiate a stream, for the specified channel ARN.
// Note that CreateChannel creates a stream key. If you subsequently use CreateStreamKey on the
// same channel, it will fail because a stream key already exists and there is a
// limit of 1 stream key per channel. To reset the stream key on a channel, use DeleteStreamKey
// and then CreateStreamKey.
func ivs_CreateStreamKey(cfg aws.Config, client *ivs.Client) {
	input := &ivs.CreateStreamKeyInput{
		// ChannelArn: *string, // Required
	}

	if len(_ivsChannelArn) > 0 {
		input.ChannelArn = aws.String(_ivsChannelArn)
	}
	if len(_ivsTags) > 0 {
		if err := assignInputField(input, "Tags", _ivsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateStreamKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified channel and its associated stream keys.
// If you try to delete a live channel, you will get an error (409
// ConflictException). To delete a channel that is live, call StopStream, wait for the
// Amazon EventBridge "Stream End" event (to verify that the stream's state is no
// longer Live), then call DeleteChannel. (See [Using EventBridge with Amazon IVS].)
//
// [Using EventBridge with Amazon IVS]: https://docs.aws.amazon.com/ivs/latest/userguide/eventbridge.html
func ivs_DeleteChannel(cfg aws.Config, client *ivs.Client) {
	input := &ivs.DeleteChannelInput{
		// Arn: *string, // Required
	}

	if len(_ivsArn) > 0 {
		input.Arn = aws.String(_ivsArn)
	}

	if resp, err := client.DeleteChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specified authorization key pair. This invalidates future viewer
// tokens generated using the key pair’s privateKey . For more information, see [Setting Up Private Channels]
// in the Amazon IVS User Guide.
//
// [Setting Up Private Channels]: https://docs.aws.amazon.com/ivs/latest/userguide/private-channels.html
func ivs_DeletePlaybackKeyPair(cfg aws.Config, client *ivs.Client) {
	input := &ivs.DeletePlaybackKeyPairInput{
		// Arn: *string, // Required
	}

	if len(_ivsArn) > 0 {
		input.Arn = aws.String(_ivsArn)
	}

	if resp, err := client.DeletePlaybackKeyPair(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified playback restriction policy.
func ivs_DeletePlaybackRestrictionPolicy(cfg aws.Config, client *ivs.Client) {
	input := &ivs.DeletePlaybackRestrictionPolicyInput{
		// Arn: *string, // Required
	}

	if len(_ivsArn) > 0 {
		input.Arn = aws.String(_ivsArn)
	}

	if resp, err := client.DeletePlaybackRestrictionPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the recording configuration for the specified ARN.
// If you try to delete a recording configuration that is associated with a
// channel, you will get an error (409 ConflictException). To avoid this, for all
// channels that reference the recording configuration, first use UpdateChannelto set the
// recordingConfigurationArn field to an empty string, then use
// DeleteRecordingConfiguration.
func ivs_DeleteRecordingConfiguration(cfg aws.Config, client *ivs.Client) {
	input := &ivs.DeleteRecordingConfigurationInput{
		// Arn: *string, // Required
	}

	if len(_ivsArn) > 0 {
		input.Arn = aws.String(_ivsArn)
	}

	if resp, err := client.DeleteRecordingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the stream key for the specified ARN, so it can no longer be used to
// stream.
func ivs_DeleteStreamKey(cfg aws.Config, client *ivs.Client) {
	input := &ivs.DeleteStreamKeyInput{
		// Arn: *string, // Required
	}

	if len(_ivsArn) > 0 {
		input.Arn = aws.String(_ivsArn)
	}

	if resp, err := client.DeleteStreamKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the channel configuration for the specified channel ARN. See also BatchGetChannel.
func ivs_GetChannel(cfg aws.Config, client *ivs.Client) {
	input := &ivs.GetChannelInput{
		// Arn: *string, // Required
	}

	if len(_ivsArn) > 0 {
		input.Arn = aws.String(_ivsArn)
	}

	if resp, err := client.GetChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a specified playback authorization key pair and returns the arn and
// fingerprint . The privateKey held by the caller can be used to generate viewer
// authorization tokens, to grant viewers access to private channels. For more
// information, see [Setting Up Private Channels]in the Amazon IVS User Guide.
//
// [Setting Up Private Channels]: https://docs.aws.amazon.com/ivs/latest/userguide/private-channels.html
func ivs_GetPlaybackKeyPair(cfg aws.Config, client *ivs.Client) {
	input := &ivs.GetPlaybackKeyPairInput{
		// Arn: *string, // Required
	}

	if len(_ivsArn) > 0 {
		input.Arn = aws.String(_ivsArn)
	}

	if resp, err := client.GetPlaybackKeyPair(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the specified playback restriction policy.
func ivs_GetPlaybackRestrictionPolicy(cfg aws.Config, client *ivs.Client) {
	input := &ivs.GetPlaybackRestrictionPolicyInput{
		// Arn: *string, // Required
	}

	if len(_ivsArn) > 0 {
		input.Arn = aws.String(_ivsArn)
	}

	if resp, err := client.GetPlaybackRestrictionPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the recording configuration for the specified ARN.
func ivs_GetRecordingConfiguration(cfg aws.Config, client *ivs.Client) {
	input := &ivs.GetRecordingConfigurationInput{
		// Arn: *string, // Required
	}

	if len(_ivsArn) > 0 {
		input.Arn = aws.String(_ivsArn)
	}

	if resp, err := client.GetRecordingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the active (live) stream on a specified channel.
func ivs_GetStream(cfg aws.Config, client *ivs.Client) {
	input := &ivs.GetStreamInput{
		// ChannelArn: *string, // Required
	}

	if len(_ivsChannelArn) > 0 {
		input.ChannelArn = aws.String(_ivsChannelArn)
	}

	if resp, err := client.GetStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets stream-key information for a specified ARN.
func ivs_GetStreamKey(cfg aws.Config, client *ivs.Client) {
	input := &ivs.GetStreamKeyInput{
		// Arn: *string, // Required
	}

	if len(_ivsArn) > 0 {
		input.Arn = aws.String(_ivsArn)
	}

	if resp, err := client.GetStreamKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets metadata on a specified stream.
func ivs_GetStreamSession(cfg aws.Config, client *ivs.Client) {
	input := &ivs.GetStreamSessionInput{
		// ChannelArn: *string, // Required
	}

	if len(_ivsChannelArn) > 0 {
		input.ChannelArn = aws.String(_ivsChannelArn)
	}
	if len(_ivsStreamId) > 0 {
		input.StreamId = aws.String(_ivsStreamId)
	}

	if resp, err := client.GetStreamSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Imports the public portion of a new key pair and returns its arn and fingerprint
// . The privateKey can then be used to generate viewer authorization tokens, to
// grant viewers access to private channels. For more information, see [Setting Up Private Channels]in the
// Amazon IVS User Guide.
//
// [Setting Up Private Channels]: https://docs.aws.amazon.com/ivs/latest/userguide/private-channels.html
func ivs_ImportPlaybackKeyPair(cfg aws.Config, client *ivs.Client) {
	input := &ivs.ImportPlaybackKeyPairInput{
		// PublicKeyMaterial: *string, // Required
	}

	if len(_ivsPublicKeyMaterial) > 0 {
		input.PublicKeyMaterial = aws.String(_ivsPublicKeyMaterial)
	}
	if len(_ivsName) > 0 {
		input.Name = aws.String(_ivsName)
	}
	if len(_ivsTags) > 0 {
		if err := assignInputField(input, "Tags", _ivsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.ImportPlaybackKeyPair(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets summary information about all channels in your account, in the Amazon Web
// Services region where the API request is processed. This list can be filtered to
// match a specified name or recording-configuration ARN. Filters are mutually
// exclusive and cannot be used together. If you try to use both filters, you will
// get an error (409 ConflictException).
func ivs_ListChannels(cfg aws.Config, client *ivs.Client) {
	input := &ivs.ListChannelsInput{}

	if len(_ivsFilterByName) > 0 {
		input.FilterByName = aws.String(_ivsFilterByName)
	}
	if len(_ivsFilterByPlaybackRestrictionPolicyArn) > 0 {
		input.FilterByPlaybackRestrictionPolicyArn = aws.String(_ivsFilterByPlaybackRestrictionPolicyArn)
	}
	if len(_ivsFilterByRecordingConfigurationArn) > 0 {
		input.FilterByRecordingConfigurationArn = aws.String(_ivsFilterByRecordingConfigurationArn)
	}
	if len(_ivsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ivsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ivsNextToken) > 0 {
		input.NextToken = aws.String(_ivsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListChannels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ivs.ListChannelsOutput
	p := ivs.NewListChannelsPaginator(client, input)
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

// Gets summary information about playback key pairs. For more information, see [Setting Up Private Channels]
// in the Amazon IVS User Guide.
//
// [Setting Up Private Channels]: https://docs.aws.amazon.com/ivs/latest/userguide/private-channels.html
func ivs_ListPlaybackKeyPairs(cfg aws.Config, client *ivs.Client) {
	input := &ivs.ListPlaybackKeyPairsInput{}

	if len(_ivsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ivsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ivsNextToken) > 0 {
		input.NextToken = aws.String(_ivsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPlaybackKeyPairs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ivs.ListPlaybackKeyPairsOutput
	p := ivs.NewListPlaybackKeyPairsPaginator(client, input)
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

// Gets summary information about playback restriction policies.
func ivs_ListPlaybackRestrictionPolicies(cfg aws.Config, client *ivs.Client) {
	input := &ivs.ListPlaybackRestrictionPoliciesInput{}

	if len(_ivsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ivsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ivsNextToken) > 0 {
		input.NextToken = aws.String(_ivsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPlaybackRestrictionPolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ivs.ListPlaybackRestrictionPoliciesOutput
	p := ivs.NewListPlaybackRestrictionPoliciesPaginator(client, input)
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

// Gets summary information about all recording configurations in your account, in
// the Amazon Web Services region where the API request is processed.
func ivs_ListRecordingConfigurations(cfg aws.Config, client *ivs.Client) {
	input := &ivs.ListRecordingConfigurationsInput{}

	if len(_ivsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ivsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ivsNextToken) > 0 {
		input.NextToken = aws.String(_ivsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRecordingConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ivs.ListRecordingConfigurationsOutput
	p := ivs.NewListRecordingConfigurationsPaginator(client, input)
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

// Gets summary information about stream keys for the specified channel.
func ivs_ListStreamKeys(cfg aws.Config, client *ivs.Client) {
	input := &ivs.ListStreamKeysInput{
		// ChannelArn: *string, // Required
	}

	if len(_ivsChannelArn) > 0 {
		input.ChannelArn = aws.String(_ivsChannelArn)
	}
	if len(_ivsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ivsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ivsNextToken) > 0 {
		input.NextToken = aws.String(_ivsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListStreamKeys(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ivs.ListStreamKeysOutput
	p := ivs.NewListStreamKeysPaginator(client, input)
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

// Gets a summary of current and previous streams for a specified channel in your
// account, in the AWS region where the API request is processed.
func ivs_ListStreamSessions(cfg aws.Config, client *ivs.Client) {
	input := &ivs.ListStreamSessionsInput{
		// ChannelArn: *string, // Required
	}

	if len(_ivsChannelArn) > 0 {
		input.ChannelArn = aws.String(_ivsChannelArn)
	}
	if len(_ivsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ivsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ivsNextToken) > 0 {
		input.NextToken = aws.String(_ivsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListStreamSessions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ivs.ListStreamSessionsOutput
	p := ivs.NewListStreamSessionsPaginator(client, input)
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

// Gets summary information about live streams in your account, in the Amazon Web
// Services region where the API request is processed.
func ivs_ListStreams(cfg aws.Config, client *ivs.Client) {
	input := &ivs.ListStreamsInput{}

	if len(_ivsFilterBy) > 0 {
		if err := assignInputField(input, "FilterBy", _ivsFilterBy); err != nil {
			log.Errorf("invalid --filter-by: %s", err.Error())
			return
		}
	}
	if len(_ivsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ivsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ivsNextToken) > 0 {
		input.NextToken = aws.String(_ivsNextToken)
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

	var results []*ivs.ListStreamsOutput
	p := ivs.NewListStreamsPaginator(client, input)
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

// Gets information about Amazon Web Services tags for the specified ARN.
func ivs_ListTagsForResource(cfg aws.Config, client *ivs.Client) {
	input := &ivs.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_ivsResourceArn) > 0 {
		input.ResourceArn = aws.String(_ivsResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Inserts metadata into the active stream of the specified channel. At most 5
// requests per second per channel are allowed, each with a maximum 1 KB payload.
// (If 5 TPS is not sufficient for your needs, we recommend batching your data into
// a single PutMetadata call.) At most 155 requests per second per account are
// allowed. Also see [Embedding Metadata within a Video Stream]in the Amazon IVS User Guide.
//
// [Embedding Metadata within a Video Stream]: https://docs.aws.amazon.com/ivs/latest/userguide/metadata.html
func ivs_PutMetadata(cfg aws.Config, client *ivs.Client) {
	input := &ivs.PutMetadataInput{
		// ChannelArn: *string, // Required
		// Metadata: *string, // Required
	}

	if len(_ivsChannelArn) > 0 {
		input.ChannelArn = aws.String(_ivsChannelArn)
	}
	if len(_ivsMetadata) > 0 {
		input.Metadata = aws.String(_ivsMetadata)
	}

	if resp, err := client.PutMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the process of revoking the viewer session associated with a specified
// channel ARN and viewer ID. Optionally, you can provide a version to revoke
// viewer sessions less than and including that version. For instructions on
// associating a viewer ID with a viewer session, see [Setting Up Private Channels].
//
// [Setting Up Private Channels]: https://docs.aws.amazon.com/ivs/latest/userguide/private-channels.html
func ivs_StartViewerSessionRevocation(cfg aws.Config, client *ivs.Client) {
	input := &ivs.StartViewerSessionRevocationInput{
		// ChannelArn: *string, // Required
		// ViewerId: *string, // Required
	}

	if len(_ivsChannelArn) > 0 {
		input.ChannelArn = aws.String(_ivsChannelArn)
	}
	if len(_ivsViewerId) > 0 {
		input.ViewerId = aws.String(_ivsViewerId)
	}
	if len(_ivsViewerSessionVersionsLessThanOrEqualTo) > 0 {
		if err := assignInputField(input, "ViewerSessionVersionsLessThanOrEqualTo", _ivsViewerSessionVersionsLessThanOrEqualTo); err != nil {
			log.Errorf("invalid --viewer-session-versions-less-than-or-equal-to: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartViewerSessionRevocation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disconnects the incoming RTMPS stream for the specified channel. Can be used in
// conjunction with DeleteStreamKeyto prevent further streaming to a channel.
//
// Many streaming client-software libraries automatically reconnect a dropped
// RTMPS session, so to stop the stream permanently, you may want to first revoke
// the streamKey attached to the channel.
func ivs_StopStream(cfg aws.Config, client *ivs.Client) {
	input := &ivs.StopStreamInput{
		// ChannelArn: *string, // Required
	}

	if len(_ivsChannelArn) > 0 {
		input.ChannelArn = aws.String(_ivsChannelArn)
	}

	if resp, err := client.StopStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or updates tags for the Amazon Web Services resource with the specified
// ARN.
func ivs_TagResource(cfg aws.Config, client *ivs.Client) {
	input := &ivs.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_ivsResourceArn) > 0 {
		input.ResourceArn = aws.String(_ivsResourceArn)
	}
	if len(_ivsTags) > 0 {
		if err := assignInputField(input, "Tags", _ivsTags); err != nil {
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

// Removes tags from the resource with the specified ARN.
func ivs_UntagResource(cfg aws.Config, client *ivs.Client) {
	input := &ivs.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_ivsResourceArn) > 0 {
		input.ResourceArn = aws.String(_ivsResourceArn)
	}
	if len(_ivsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _ivsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a channel's configuration. Live channels cannot be updated. You must
// stop the ongoing stream, update the channel, and restart the stream for the
// changes to take effect.
func ivs_UpdateChannel(cfg aws.Config, client *ivs.Client) {
	input := &ivs.UpdateChannelInput{
		// Arn: *string, // Required
	}

	if len(_ivsArn) > 0 {
		input.Arn = aws.String(_ivsArn)
	}
	if len(_ivsAuthorized) > 0 {
		if err := assignInputField(input, "Authorized", _ivsAuthorized); err != nil {
			log.Errorf("invalid --authorized: %s", err.Error())
			return
		}
	}
	if len(_ivsContainerFormat) > 0 {
		if err := assignInputField(input, "ContainerFormat", _ivsContainerFormat); err != nil {
			log.Errorf("invalid --container-format: %s", err.Error())
			return
		}
	}
	if len(_ivsInsecureIngest) > 0 {
		if err := assignInputField(input, "InsecureIngest", _ivsInsecureIngest); err != nil {
			log.Errorf("invalid --insecure-ingest: %s", err.Error())
			return
		}
	}
	if len(_ivsLatencyMode) > 0 {
		if err := assignInputField(input, "LatencyMode", _ivsLatencyMode); err != nil {
			log.Errorf("invalid --latency-mode: %s", err.Error())
			return
		}
	}
	if len(_ivsMultitrackInputConfiguration) > 0 {
		if err := assignInputField(input, "MultitrackInputConfiguration", _ivsMultitrackInputConfiguration); err != nil {
			log.Errorf("invalid --multitrack-input-configuration: %s", err.Error())
			return
		}
	}
	if len(_ivsName) > 0 {
		input.Name = aws.String(_ivsName)
	}
	if len(_ivsPlaybackRestrictionPolicyArn) > 0 {
		input.PlaybackRestrictionPolicyArn = aws.String(_ivsPlaybackRestrictionPolicyArn)
	}
	if len(_ivsPreset) > 0 {
		if err := assignInputField(input, "Preset", _ivsPreset); err != nil {
			log.Errorf("invalid --preset: %s", err.Error())
			return
		}
	}
	if len(_ivsRecordingConfigurationArn) > 0 {
		input.RecordingConfigurationArn = aws.String(_ivsRecordingConfigurationArn)
	}
	if len(_ivsType) > 0 {
		if err := assignInputField(input, "Type", _ivsType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a specified playback restriction policy.
func ivs_UpdatePlaybackRestrictionPolicy(cfg aws.Config, client *ivs.Client) {
	input := &ivs.UpdatePlaybackRestrictionPolicyInput{
		// Arn: *string, // Required
	}

	if len(_ivsArn) > 0 {
		input.Arn = aws.String(_ivsArn)
	}
	if len(_ivsAllowedCountries) > 0 {
		input.AllowedCountries = append([]string(nil), _ivsAllowedCountries...)
	}
	if len(_ivsAllowedOrigins) > 0 {
		input.AllowedOrigins = append([]string(nil), _ivsAllowedOrigins...)
	}
	if len(_ivsEnableStrictOriginEnforcement) > 0 {
		if err := assignInputField(input, "EnableStrictOriginEnforcement", _ivsEnableStrictOriginEnforcement); err != nil {
			log.Errorf("invalid --enable-strict-origin-enforcement: %s", err.Error())
			return
		}
	}
	if len(_ivsName) > 0 {
		input.Name = aws.String(_ivsName)
	}

	if resp, err := client.UpdatePlaybackRestrictionPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_ivsCmd)
	_ivsCmd.Flags().SortFlags = false

	_ivsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_ivsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_ivsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_ivsCmd.Flags().StringSliceVarP(&_ivsAllowedCountries, "allowed-countries", "", nil, "Allowed Countries")
	_ivsCmd.Flags().StringSliceVarP(&_ivsAllowedOrigins, "allowed-origins", "", nil, "Allowed Origins")
	_ivsCmd.Flags().StringVarP(&_ivsArn, "arn", "", "", "ARN")
	_ivsCmd.Flags().StringSliceVarP(&_ivsArns, "arns", "", nil, "Arns")
	_ivsCmd.Flags().StringVarP(&_ivsAuthorized, "authorized", "", "", "Authorized")
	_ivsCmd.Flags().StringVarP(&_ivsChannelArn, "channel-arn", "", "", "Channel ARN")
	_ivsCmd.Flags().StringVarP(&_ivsContainerFormat, "container-format", "", "", "Container Format")
	_ivsCmd.Flags().StringVarP(&_ivsDestinationConfiguration, "destination-configuration", "", "", "Destination Configuration")
	_ivsCmd.Flags().StringVarP(&_ivsEnableStrictOriginEnforcement, "enable-strict-origin-enforcement", "", "", "Enable Strict Origin Enforcement")
	_ivsCmd.Flags().StringVarP(&_ivsFilterBy, "filter-by", "", "", "Filter By")
	_ivsCmd.Flags().StringVarP(&_ivsFilterByName, "filter-by-name", "", "", "Filter By Name")
	_ivsCmd.Flags().StringVarP(&_ivsFilterByPlaybackRestrictionPolicyArn, "filter-by-playback-restriction-policy-arn", "", "", "Filter By Playback Restriction Policy ARN")
	_ivsCmd.Flags().StringVarP(&_ivsFilterByRecordingConfigurationArn, "filter-by-recording-configuration-arn", "", "", "Filter By Recording Configuration ARN")
	_ivsCmd.Flags().StringVarP(&_ivsInsecureIngest, "insecure-ingest", "", "", "Insecure Ingest")
	_ivsCmd.Flags().StringVarP(&_ivsLatencyMode, "latency-mode", "", "", "Latency Mode")
	_ivsCmd.Flags().StringVarP(&_ivsMaxResults, "max-results", "", "", "Max Results")
	_ivsCmd.Flags().StringVarP(&_ivsMetadata, "metadata", "", "", "Metadata")
	_ivsCmd.Flags().StringVarP(&_ivsMultitrackInputConfiguration, "multitrack-input-configuration", "", "", "Multitrack Input Configuration")
	_ivsCmd.Flags().StringVarP(&_ivsName, "name", "", "", "Name")
	_ivsCmd.Flags().StringVarP(&_ivsNextToken, "next-token", "", "", "Next Token")
	_ivsCmd.Flags().StringVarP(&_ivsPlaybackRestrictionPolicyArn, "playback-restriction-policy-arn", "", "", "Playback Restriction Policy ARN")
	_ivsCmd.Flags().StringVarP(&_ivsPreset, "preset", "", "", "Preset")
	_ivsCmd.Flags().StringVarP(&_ivsPublicKeyMaterial, "public-key-material", "", "", "Public Key Material")
	_ivsCmd.Flags().StringVarP(&_ivsRecordingConfigurationArn, "recording-configuration-arn", "", "", "Recording Configuration ARN")
	_ivsCmd.Flags().StringVarP(&_ivsRecordingReconnectWindowSeconds, "recording-reconnect-window-seconds", "", "", "Recording Reconnect Window Seconds")
	_ivsCmd.Flags().StringVarP(&_ivsRenditionConfiguration, "rendition-configuration", "", "", "Rendition Configuration")
	_ivsCmd.Flags().StringVarP(&_ivsResourceArn, "resource-arn", "", "", "Resource ARN")
	_ivsCmd.Flags().StringVarP(&_ivsStreamId, "stream-id", "", "", "Stream ID")
	_ivsCmd.Flags().StringSliceVarP(&_ivsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_ivsCmd.Flags().StringVarP(&_ivsTags, "tags", "", "", "Tags")
	_ivsCmd.Flags().StringVarP(&_ivsThumbnailConfiguration, "thumbnail-configuration", "", "", "Thumbnail Configuration")
	_ivsCmd.Flags().StringVarP(&_ivsType, "type", "", "", "Type")
	_ivsCmd.Flags().StringVarP(&_ivsViewerId, "viewer-id", "", "", "Viewer ID")
	_ivsCmd.Flags().StringVarP(&_ivsViewerSessionVersionsLessThanOrEqualTo, "viewer-session-versions-less-than-or-equal-to", "", "", "Viewer Session Versions Less Than Or Equal To")
	_ivsCmd.Flags().StringVarP(&_ivsViewerSessions, "viewer-sessions", "", "", "Viewer Sessions")

	_ivsCmd.Flags().BoolVarP(&_ivsBatchGetChannel, "batch-get-channel", "", false, "Batch Get Channel")
	_ivsCmd.Flags().BoolVarP(&_ivsBatchGetStreamKey, "batch-get-stream-key", "", false, "Batch Get Stream Key")
	_ivsCmd.Flags().BoolVarP(&_ivsBatchStartViewerSessionRevocation, "batch-start-viewer-session-revocation", "", false, "Batch Start Viewer Session Revocation")
	_ivsCmd.Flags().BoolVarP(&_ivsCreateChannel, "create-channel", "", false, "Create Channel")
	_ivsCmd.Flags().BoolVarP(&_ivsCreatePlaybackRestrictionPolicy, "create-playback-restriction-policy", "", false, "Create Playback Restriction Policy")
	_ivsCmd.Flags().BoolVarP(&_ivsCreateRecordingConfiguration, "create-recording-configuration", "", false, "Create Recording Configuration")
	_ivsCmd.Flags().BoolVarP(&_ivsCreateStreamKey, "create-stream-key", "", false, "Create Stream Key")
	_ivsCmd.Flags().BoolVarP(&_ivsDeleteChannel, "delete-channel", "", false, "Delete Channel")
	_ivsCmd.Flags().BoolVarP(&_ivsDeletePlaybackKeyPair, "delete-playback-key-pair", "", false, "Delete Playback Key Pair")
	_ivsCmd.Flags().BoolVarP(&_ivsDeletePlaybackRestrictionPolicy, "delete-playback-restriction-policy", "", false, "Delete Playback Restriction Policy")
	_ivsCmd.Flags().BoolVarP(&_ivsDeleteRecordingConfiguration, "delete-recording-configuration", "", false, "Delete Recording Configuration")
	_ivsCmd.Flags().BoolVarP(&_ivsDeleteStreamKey, "delete-stream-key", "", false, "Delete Stream Key")
	_ivsCmd.Flags().BoolVarP(&_ivsGetChannel, "get-channel", "", false, "Get Channel")
	_ivsCmd.Flags().BoolVarP(&_ivsGetPlaybackKeyPair, "get-playback-key-pair", "", false, "Get Playback Key Pair")
	_ivsCmd.Flags().BoolVarP(&_ivsGetPlaybackRestrictionPolicy, "get-playback-restriction-policy", "", false, "Get Playback Restriction Policy")
	_ivsCmd.Flags().BoolVarP(&_ivsGetRecordingConfiguration, "get-recording-configuration", "", false, "Get Recording Configuration")
	_ivsCmd.Flags().BoolVarP(&_ivsGetStream, "get-stream", "", false, "Get Stream")
	_ivsCmd.Flags().BoolVarP(&_ivsGetStreamKey, "get-stream-key", "", false, "Get Stream Key")
	_ivsCmd.Flags().BoolVarP(&_ivsGetStreamSession, "get-stream-session", "", false, "Get Stream Session")
	_ivsCmd.Flags().BoolVarP(&_ivsImportPlaybackKeyPair, "import-playback-key-pair", "", false, "Import Playback Key Pair")
	_ivsCmd.Flags().BoolVarP(&_ivsListChannels, "list-channels", "", false, "List Channels")
	_ivsCmd.Flags().BoolVarP(&_ivsListPlaybackKeyPairs, "list-playback-key-pairs", "", false, "List Playback Key Pairs")
	_ivsCmd.Flags().BoolVarP(&_ivsListPlaybackRestrictionPolicies, "list-playback-restriction-policies", "", false, "List Playback Restriction Policies")
	_ivsCmd.Flags().BoolVarP(&_ivsListRecordingConfigurations, "list-recording-configurations", "", false, "List Recording Configurations")
	_ivsCmd.Flags().BoolVarP(&_ivsListStreamKeys, "list-stream-keys", "", false, "List Stream Keys")
	_ivsCmd.Flags().BoolVarP(&_ivsListStreamSessions, "list-stream-sessions", "", false, "List Stream Sessions")
	_ivsCmd.Flags().BoolVarP(&_ivsListStreams, "list-streams", "", false, "List Streams")
	_ivsCmd.Flags().BoolVarP(&_ivsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_ivsCmd.Flags().BoolVarP(&_ivsPutMetadata, "put-metadata", "", false, "Put Metadata")
	_ivsCmd.Flags().BoolVarP(&_ivsStartViewerSessionRevocation, "start-viewer-session-revocation", "", false, "Start Viewer Session Revocation")
	_ivsCmd.Flags().BoolVarP(&_ivsStopStream, "stop-stream", "", false, "Stop Stream")
	_ivsCmd.Flags().BoolVarP(&_ivsTagResource, "tag-resource", "", false, "Tag Resource")
	_ivsCmd.Flags().BoolVarP(&_ivsUntagResource, "untag-resource", "", false, "Untag Resource")
	_ivsCmd.Flags().BoolVarP(&_ivsUpdateChannel, "update-channel", "", false, "Update Channel")
	_ivsCmd.Flags().BoolVarP(&_ivsUpdatePlaybackRestrictionPolicy, "update-playback-restriction-policy", "", false, "Update Playback Restriction Policy")

}
