package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediapackagev2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// mediapackagev2Cmd represents the mediapackagev2 command
var _mediapackagev2Cmd = &cobra.Command{
	Use:   "mediapackagev2",
	Short: "AWS mediapackagev2 CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := mediapackagev2.NewFromConfig(cfg)
		if _mediapackagev2CancelHarvestJob {
			mediapackagev2_CancelHarvestJob(cfg, client)
			return
		}
		if _mediapackagev2CreateChannel {
			mediapackagev2_CreateChannel(cfg, client)
			return
		}
		if _mediapackagev2CreateChannelGroup {
			mediapackagev2_CreateChannelGroup(cfg, client)
			return
		}
		if _mediapackagev2CreateHarvestJob {
			mediapackagev2_CreateHarvestJob(cfg, client)
			return
		}
		if _mediapackagev2CreateOriginEndpoint {
			mediapackagev2_CreateOriginEndpoint(cfg, client)
			return
		}
		if _mediapackagev2DeleteChannel {
			mediapackagev2_DeleteChannel(cfg, client)
			return
		}
		if _mediapackagev2DeleteChannelGroup {
			mediapackagev2_DeleteChannelGroup(cfg, client)
			return
		}
		if _mediapackagev2DeleteChannelPolicy {
			mediapackagev2_DeleteChannelPolicy(cfg, client)
			return
		}
		if _mediapackagev2DeleteOriginEndpoint {
			mediapackagev2_DeleteOriginEndpoint(cfg, client)
			return
		}
		if _mediapackagev2DeleteOriginEndpointPolicy {
			mediapackagev2_DeleteOriginEndpointPolicy(cfg, client)
			return
		}
		if _mediapackagev2GetChannel {
			mediapackagev2_GetChannel(cfg, client)
			return
		}
		if _mediapackagev2GetChannelGroup {
			mediapackagev2_GetChannelGroup(cfg, client)
			return
		}
		if _mediapackagev2GetChannelPolicy {
			mediapackagev2_GetChannelPolicy(cfg, client)
			return
		}
		if _mediapackagev2GetHarvestJob {
			mediapackagev2_GetHarvestJob(cfg, client)
			return
		}
		if _mediapackagev2GetOriginEndpoint {
			mediapackagev2_GetOriginEndpoint(cfg, client)
			return
		}
		if _mediapackagev2GetOriginEndpointPolicy {
			mediapackagev2_GetOriginEndpointPolicy(cfg, client)
			return
		}
		if _mediapackagev2ListChannelGroups {
			mediapackagev2_ListChannelGroups(cfg, client)
			return
		}
		if _mediapackagev2ListChannels {
			mediapackagev2_ListChannels(cfg, client)
			return
		}
		if _mediapackagev2ListHarvestJobs {
			mediapackagev2_ListHarvestJobs(cfg, client)
			return
		}
		if _mediapackagev2ListOriginEndpoints {
			mediapackagev2_ListOriginEndpoints(cfg, client)
			return
		}
		if _mediapackagev2ListTagsForResource {
			mediapackagev2_ListTagsForResource(cfg, client)
			return
		}
		if _mediapackagev2PutChannelPolicy {
			mediapackagev2_PutChannelPolicy(cfg, client)
			return
		}
		if _mediapackagev2PutOriginEndpointPolicy {
			mediapackagev2_PutOriginEndpointPolicy(cfg, client)
			return
		}
		if _mediapackagev2ResetChannelState {
			mediapackagev2_ResetChannelState(cfg, client)
			return
		}
		if _mediapackagev2ResetOriginEndpointState {
			mediapackagev2_ResetOriginEndpointState(cfg, client)
			return
		}
		if _mediapackagev2TagResource {
			mediapackagev2_TagResource(cfg, client)
			return
		}
		if _mediapackagev2UntagResource {
			mediapackagev2_UntagResource(cfg, client)
			return
		}
		if _mediapackagev2UpdateChannel {
			mediapackagev2_UpdateChannel(cfg, client)
			return
		}
		if _mediapackagev2UpdateChannelGroup {
			mediapackagev2_UpdateChannelGroup(cfg, client)
			return
		}
		if _mediapackagev2UpdateOriginEndpoint {
			mediapackagev2_UpdateOriginEndpoint(cfg, client)
			return
		}

	},
}

var (
	_mediapackagev2CancelHarvestJob           bool
	_mediapackagev2CreateChannel              bool
	_mediapackagev2CreateChannelGroup         bool
	_mediapackagev2CreateHarvestJob           bool
	_mediapackagev2CreateOriginEndpoint       bool
	_mediapackagev2DeleteChannel              bool
	_mediapackagev2DeleteChannelGroup         bool
	_mediapackagev2DeleteChannelPolicy        bool
	_mediapackagev2DeleteOriginEndpoint       bool
	_mediapackagev2DeleteOriginEndpointPolicy bool
	_mediapackagev2GetChannel                 bool
	_mediapackagev2GetChannelGroup            bool
	_mediapackagev2GetChannelPolicy           bool
	_mediapackagev2GetHarvestJob              bool
	_mediapackagev2GetOriginEndpoint          bool
	_mediapackagev2GetOriginEndpointPolicy    bool
	_mediapackagev2ListChannelGroups          bool
	_mediapackagev2ListChannels               bool
	_mediapackagev2ListHarvestJobs            bool
	_mediapackagev2ListOriginEndpoints        bool
	_mediapackagev2ListTagsForResource        bool
	_mediapackagev2PutChannelPolicy           bool
	_mediapackagev2PutOriginEndpointPolicy    bool
	_mediapackagev2ResetChannelState          bool
	_mediapackagev2ResetOriginEndpointState   bool
	_mediapackagev2TagResource                bool
	_mediapackagev2UntagResource              bool
	_mediapackagev2UpdateChannel              bool
	_mediapackagev2UpdateChannelGroup         bool
	_mediapackagev2UpdateOriginEndpoint       bool

	_mediapackagev2CdnAuthConfiguration            string
	_mediapackagev2ChannelGroupName                string
	_mediapackagev2ChannelName                     string
	_mediapackagev2ClientToken                     string
	_mediapackagev2ContainerType                   string
	_mediapackagev2DashManifests                   string
	_mediapackagev2Description                     string
	_mediapackagev2Destination                     string
	_mediapackagev2ETag                            string
	_mediapackagev2ForceEndpointErrorConfiguration string
	_mediapackagev2HarvestJobName                  string
	_mediapackagev2HarvestedManifests              string
	_mediapackagev2HlsManifests                    string
	_mediapackagev2InputSwitchConfiguration        string
	_mediapackagev2InputType                       string
	_mediapackagev2LowLatencyHlsManifests          string
	_mediapackagev2MaxResults                      string
	_mediapackagev2MssManifests                    string
	_mediapackagev2NextToken                       string
	_mediapackagev2OriginEndpointName              string
	_mediapackagev2OutputHeaderConfiguration       string
	_mediapackagev2Policy                          string
	_mediapackagev2ResourceArn                     string
	_mediapackagev2ScheduleConfiguration           string
	_mediapackagev2Segment                         string
	_mediapackagev2StartoverWindowSeconds          string
	_mediapackagev2Status                          string
	_mediapackagev2TagKeys                         []string
	_mediapackagev2Tags                            string
)

// Cancels an in-progress harvest job.
func mediapackagev2_CancelHarvestJob(cfg aws.Config, client *mediapackagev2.Client) {
	input := &mediapackagev2.CancelHarvestJobInput{
		// ChannelGroupName: *string, // Required
		// ChannelName: *string, // Required
		// HarvestJobName: *string, // Required
		// OriginEndpointName: *string, // Required
	}

	if len(_mediapackagev2ChannelGroupName) > 0 {
		input.ChannelGroupName = aws.String(_mediapackagev2ChannelGroupName)
	}
	if len(_mediapackagev2ChannelName) > 0 {
		input.ChannelName = aws.String(_mediapackagev2ChannelName)
	}
	if len(_mediapackagev2HarvestJobName) > 0 {
		input.HarvestJobName = aws.String(_mediapackagev2HarvestJobName)
	}
	if len(_mediapackagev2OriginEndpointName) > 0 {
		input.OriginEndpointName = aws.String(_mediapackagev2OriginEndpointName)
	}
	if len(_mediapackagev2ETag) > 0 {
		input.ETag = aws.String(_mediapackagev2ETag)
	}

	if resp, err := client.CancelHarvestJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a channel to start receiving content streams. The channel represents the
// input to MediaPackage for incoming live content from an encoder such as AWS
// Elemental MediaLive. The channel receives content, and after packaging it,
// outputs it through an origin endpoint to downstream devices (such as video
// players or CDNs) that request the content. You can create only one channel with
// each request. We recommend that you spread out channels between channel groups,
// such as putting redundant channels in the same AWS Region in different channel
// groups.
func mediapackagev2_CreateChannel(cfg aws.Config, client *mediapackagev2.Client) {
	input := &mediapackagev2.CreateChannelInput{
		// ChannelGroupName: *string, // Required
		// ChannelName: *string, // Required
	}

	if len(_mediapackagev2ChannelGroupName) > 0 {
		input.ChannelGroupName = aws.String(_mediapackagev2ChannelGroupName)
	}
	if len(_mediapackagev2ChannelName) > 0 {
		input.ChannelName = aws.String(_mediapackagev2ChannelName)
	}
	if len(_mediapackagev2ClientToken) > 0 {
		input.ClientToken = aws.String(_mediapackagev2ClientToken)
	}
	if len(_mediapackagev2Description) > 0 {
		input.Description = aws.String(_mediapackagev2Description)
	}
	if len(_mediapackagev2InputSwitchConfiguration) > 0 {
		if err := assignInputField(input, "InputSwitchConfiguration", _mediapackagev2InputSwitchConfiguration); err != nil {
			log.Errorf("invalid --input-switch-configuration: %s", err.Error())
			return
		}
	}
	if len(_mediapackagev2InputType) > 0 {
		if err := assignInputField(input, "InputType", _mediapackagev2InputType); err != nil {
			log.Errorf("invalid --input-type: %s", err.Error())
			return
		}
	}
	if len(_mediapackagev2OutputHeaderConfiguration) > 0 {
		if err := assignInputField(input, "OutputHeaderConfiguration", _mediapackagev2OutputHeaderConfiguration); err != nil {
			log.Errorf("invalid --output-header-configuration: %s", err.Error())
			return
		}
	}
	if len(_mediapackagev2Tags) > 0 {
		if err := assignInputField(input, "Tags", _mediapackagev2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
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

// Create a channel group to group your channels and origin endpoints. A channel
// group is the top-level resource that consists of channels and origin endpoints
// that are associated with it and that provides predictable URLs for stream
// delivery. All channels and origin endpoints within the channel group are
// guaranteed to share the DNS. You can create only one channel group with each
// request.
func mediapackagev2_CreateChannelGroup(cfg aws.Config, client *mediapackagev2.Client) {
	input := &mediapackagev2.CreateChannelGroupInput{
		// ChannelGroupName: *string, // Required
	}

	if len(_mediapackagev2ChannelGroupName) > 0 {
		input.ChannelGroupName = aws.String(_mediapackagev2ChannelGroupName)
	}
	if len(_mediapackagev2ClientToken) > 0 {
		input.ClientToken = aws.String(_mediapackagev2ClientToken)
	}
	if len(_mediapackagev2Description) > 0 {
		input.Description = aws.String(_mediapackagev2Description)
	}
	if len(_mediapackagev2Tags) > 0 {
		if err := assignInputField(input, "Tags", _mediapackagev2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateChannelGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new harvest job to export content from a MediaPackage v2 channel to
// an S3 bucket.
func mediapackagev2_CreateHarvestJob(cfg aws.Config, client *mediapackagev2.Client) {
	input := &mediapackagev2.CreateHarvestJobInput{
		// ChannelGroupName: *string, // Required
		// ChannelName: *string, // Required
		// Destination: *types.Destination, // Required
		// HarvestedManifests: *types.HarvestedManifests, // Required
		// OriginEndpointName: *string, // Required
		// ScheduleConfiguration: *types.HarvesterScheduleConfiguration, // Required
	}

	if len(_mediapackagev2ChannelGroupName) > 0 {
		input.ChannelGroupName = aws.String(_mediapackagev2ChannelGroupName)
	}
	if len(_mediapackagev2ChannelName) > 0 {
		input.ChannelName = aws.String(_mediapackagev2ChannelName)
	}
	if len(_mediapackagev2Destination) > 0 {
		if err := assignInputField(input, "Destination", _mediapackagev2Destination); err != nil {
			log.Errorf("invalid --destination: %s", err.Error())
			return
		}
	}
	if len(_mediapackagev2HarvestedManifests) > 0 {
		if err := assignInputField(input, "HarvestedManifests", _mediapackagev2HarvestedManifests); err != nil {
			log.Errorf("invalid --harvested-manifests: %s", err.Error())
			return
		}
	}
	if len(_mediapackagev2OriginEndpointName) > 0 {
		input.OriginEndpointName = aws.String(_mediapackagev2OriginEndpointName)
	}
	if len(_mediapackagev2ScheduleConfiguration) > 0 {
		if err := assignInputField(input, "ScheduleConfiguration", _mediapackagev2ScheduleConfiguration); err != nil {
			log.Errorf("invalid --schedule-configuration: %s", err.Error())
			return
		}
	}
	if len(_mediapackagev2ClientToken) > 0 {
		input.ClientToken = aws.String(_mediapackagev2ClientToken)
	}
	if len(_mediapackagev2Description) > 0 {
		input.Description = aws.String(_mediapackagev2Description)
	}
	if len(_mediapackagev2HarvestJobName) > 0 {
		input.HarvestJobName = aws.String(_mediapackagev2HarvestJobName)
	}
	if len(_mediapackagev2Tags) > 0 {
		if err := assignInputField(input, "Tags", _mediapackagev2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateHarvestJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The endpoint is attached to a channel, and represents the output of the live
// content. You can associate multiple endpoints to a single channel. Each endpoint
// gives players and downstream CDNs (such as Amazon CloudFront) access to the
// content for playback. Content can't be served from a channel until it has an
// endpoint. You can create only one endpoint with each request.
func mediapackagev2_CreateOriginEndpoint(cfg aws.Config, client *mediapackagev2.Client) {
	input := &mediapackagev2.CreateOriginEndpointInput{
		// ChannelGroupName: *string, // Required
		// ChannelName: *string, // Required
		// ContainerType: types.ContainerType, // Required
		// OriginEndpointName: *string, // Required
	}

	if len(_mediapackagev2ChannelGroupName) > 0 {
		input.ChannelGroupName = aws.String(_mediapackagev2ChannelGroupName)
	}
	if len(_mediapackagev2ChannelName) > 0 {
		input.ChannelName = aws.String(_mediapackagev2ChannelName)
	}
	if len(_mediapackagev2ContainerType) > 0 {
		if err := assignInputField(input, "ContainerType", _mediapackagev2ContainerType); err != nil {
			log.Errorf("invalid --container-type: %s", err.Error())
			return
		}
	}
	if len(_mediapackagev2OriginEndpointName) > 0 {
		input.OriginEndpointName = aws.String(_mediapackagev2OriginEndpointName)
	}
	if len(_mediapackagev2ClientToken) > 0 {
		input.ClientToken = aws.String(_mediapackagev2ClientToken)
	}
	if len(_mediapackagev2DashManifests) > 0 {
		if err := assignInputField(input, "DashManifests", _mediapackagev2DashManifests); err != nil {
			log.Errorf("invalid --dash-manifests: %s", err.Error())
			return
		}
	}
	if len(_mediapackagev2Description) > 0 {
		input.Description = aws.String(_mediapackagev2Description)
	}
	if len(_mediapackagev2ForceEndpointErrorConfiguration) > 0 {
		if err := assignInputField(input, "ForceEndpointErrorConfiguration", _mediapackagev2ForceEndpointErrorConfiguration); err != nil {
			log.Errorf("invalid --force-endpoint-error-configuration: %s", err.Error())
			return
		}
	}
	if len(_mediapackagev2HlsManifests) > 0 {
		if err := assignInputField(input, "HlsManifests", _mediapackagev2HlsManifests); err != nil {
			log.Errorf("invalid --hls-manifests: %s", err.Error())
			return
		}
	}
	if len(_mediapackagev2LowLatencyHlsManifests) > 0 {
		if err := assignInputField(input, "LowLatencyHlsManifests", _mediapackagev2LowLatencyHlsManifests); err != nil {
			log.Errorf("invalid --low-latency-hls-manifests: %s", err.Error())
			return
		}
	}
	if len(_mediapackagev2MssManifests) > 0 {
		if err := assignInputField(input, "MssManifests", _mediapackagev2MssManifests); err != nil {
			log.Errorf("invalid --mss-manifests: %s", err.Error())
			return
		}
	}
	if len(_mediapackagev2Segment) > 0 {
		if err := assignInputField(input, "Segment", _mediapackagev2Segment); err != nil {
			log.Errorf("invalid --segment: %s", err.Error())
			return
		}
	}
	if len(_mediapackagev2StartoverWindowSeconds) > 0 {
		if err := assignInputField(input, "StartoverWindowSeconds", _mediapackagev2StartoverWindowSeconds); err != nil {
			log.Errorf("invalid --startover-window-seconds: %s", err.Error())
			return
		}
	}
	if len(_mediapackagev2Tags) > 0 {
		if err := assignInputField(input, "Tags", _mediapackagev2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateOriginEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a channel to stop AWS Elemental MediaPackage from receiving further
// content. You must delete the channel's origin endpoints before you can delete
// the channel.
func mediapackagev2_DeleteChannel(cfg aws.Config, client *mediapackagev2.Client) {
	input := &mediapackagev2.DeleteChannelInput{
		// ChannelGroupName: *string, // Required
		// ChannelName: *string, // Required
	}

	if len(_mediapackagev2ChannelGroupName) > 0 {
		input.ChannelGroupName = aws.String(_mediapackagev2ChannelGroupName)
	}
	if len(_mediapackagev2ChannelName) > 0 {
		input.ChannelName = aws.String(_mediapackagev2ChannelName)
	}

	if resp, err := client.DeleteChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a channel group. You must delete the channel group's channels and origin
// endpoints before you can delete the channel group. If you delete a channel
// group, you'll lose access to the egress domain and will have to create a new
// channel group to replace it.
func mediapackagev2_DeleteChannelGroup(cfg aws.Config, client *mediapackagev2.Client) {
	input := &mediapackagev2.DeleteChannelGroupInput{
		// ChannelGroupName: *string, // Required
	}

	if len(_mediapackagev2ChannelGroupName) > 0 {
		input.ChannelGroupName = aws.String(_mediapackagev2ChannelGroupName)
	}

	if resp, err := client.DeleteChannelGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a channel policy.
func mediapackagev2_DeleteChannelPolicy(cfg aws.Config, client *mediapackagev2.Client) {
	input := &mediapackagev2.DeleteChannelPolicyInput{
		// ChannelGroupName: *string, // Required
		// ChannelName: *string, // Required
	}

	if len(_mediapackagev2ChannelGroupName) > 0 {
		input.ChannelGroupName = aws.String(_mediapackagev2ChannelGroupName)
	}
	if len(_mediapackagev2ChannelName) > 0 {
		input.ChannelName = aws.String(_mediapackagev2ChannelName)
	}

	if resp, err := client.DeleteChannelPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Origin endpoints can serve content until they're deleted. Delete the endpoint
// if it should no longer respond to playback requests. You must delete all
// endpoints from a channel before you can delete the channel.
func mediapackagev2_DeleteOriginEndpoint(cfg aws.Config, client *mediapackagev2.Client) {
	input := &mediapackagev2.DeleteOriginEndpointInput{
		// ChannelGroupName: *string, // Required
		// ChannelName: *string, // Required
		// OriginEndpointName: *string, // Required
	}

	if len(_mediapackagev2ChannelGroupName) > 0 {
		input.ChannelGroupName = aws.String(_mediapackagev2ChannelGroupName)
	}
	if len(_mediapackagev2ChannelName) > 0 {
		input.ChannelName = aws.String(_mediapackagev2ChannelName)
	}
	if len(_mediapackagev2OriginEndpointName) > 0 {
		input.OriginEndpointName = aws.String(_mediapackagev2OriginEndpointName)
	}

	if resp, err := client.DeleteOriginEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an origin endpoint policy.
func mediapackagev2_DeleteOriginEndpointPolicy(cfg aws.Config, client *mediapackagev2.Client) {
	input := &mediapackagev2.DeleteOriginEndpointPolicyInput{
		// ChannelGroupName: *string, // Required
		// ChannelName: *string, // Required
		// OriginEndpointName: *string, // Required
	}

	if len(_mediapackagev2ChannelGroupName) > 0 {
		input.ChannelGroupName = aws.String(_mediapackagev2ChannelGroupName)
	}
	if len(_mediapackagev2ChannelName) > 0 {
		input.ChannelName = aws.String(_mediapackagev2ChannelName)
	}
	if len(_mediapackagev2OriginEndpointName) > 0 {
		input.OriginEndpointName = aws.String(_mediapackagev2OriginEndpointName)
	}

	if resp, err := client.DeleteOriginEndpointPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the specified channel that's configured in AWS Elemental MediaPackage.
func mediapackagev2_GetChannel(cfg aws.Config, client *mediapackagev2.Client) {
	input := &mediapackagev2.GetChannelInput{
		// ChannelGroupName: *string, // Required
		// ChannelName: *string, // Required
	}

	if len(_mediapackagev2ChannelGroupName) > 0 {
		input.ChannelGroupName = aws.String(_mediapackagev2ChannelGroupName)
	}
	if len(_mediapackagev2ChannelName) > 0 {
		input.ChannelName = aws.String(_mediapackagev2ChannelName)
	}

	if resp, err := client.GetChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the specified channel group that's configured in AWS Elemental
// MediaPackage.
func mediapackagev2_GetChannelGroup(cfg aws.Config, client *mediapackagev2.Client) {
	input := &mediapackagev2.GetChannelGroupInput{
		// ChannelGroupName: *string, // Required
	}

	if len(_mediapackagev2ChannelGroupName) > 0 {
		input.ChannelGroupName = aws.String(_mediapackagev2ChannelGroupName)
	}

	if resp, err := client.GetChannelGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the specified channel policy that's configured in AWS Elemental
// MediaPackage. With policies, you can specify who has access to AWS resources and
// what actions they can perform on those resources.
func mediapackagev2_GetChannelPolicy(cfg aws.Config, client *mediapackagev2.Client) {
	input := &mediapackagev2.GetChannelPolicyInput{
		// ChannelGroupName: *string, // Required
		// ChannelName: *string, // Required
	}

	if len(_mediapackagev2ChannelGroupName) > 0 {
		input.ChannelGroupName = aws.String(_mediapackagev2ChannelGroupName)
	}
	if len(_mediapackagev2ChannelName) > 0 {
		input.ChannelName = aws.String(_mediapackagev2ChannelName)
	}

	if resp, err := client.GetChannelPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the details of a specific harvest job.
func mediapackagev2_GetHarvestJob(cfg aws.Config, client *mediapackagev2.Client) {
	input := &mediapackagev2.GetHarvestJobInput{
		// ChannelGroupName: *string, // Required
		// ChannelName: *string, // Required
		// HarvestJobName: *string, // Required
		// OriginEndpointName: *string, // Required
	}

	if len(_mediapackagev2ChannelGroupName) > 0 {
		input.ChannelGroupName = aws.String(_mediapackagev2ChannelGroupName)
	}
	if len(_mediapackagev2ChannelName) > 0 {
		input.ChannelName = aws.String(_mediapackagev2ChannelName)
	}
	if len(_mediapackagev2HarvestJobName) > 0 {
		input.HarvestJobName = aws.String(_mediapackagev2HarvestJobName)
	}
	if len(_mediapackagev2OriginEndpointName) > 0 {
		input.OriginEndpointName = aws.String(_mediapackagev2OriginEndpointName)
	}

	if resp, err := client.GetHarvestJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the specified origin endpoint that's configured in AWS Elemental
// MediaPackage to obtain its playback URL and to view the packaging settings that
// it's currently using.
func mediapackagev2_GetOriginEndpoint(cfg aws.Config, client *mediapackagev2.Client) {
	input := &mediapackagev2.GetOriginEndpointInput{
		// ChannelGroupName: *string, // Required
		// ChannelName: *string, // Required
		// OriginEndpointName: *string, // Required
	}

	if len(_mediapackagev2ChannelGroupName) > 0 {
		input.ChannelGroupName = aws.String(_mediapackagev2ChannelGroupName)
	}
	if len(_mediapackagev2ChannelName) > 0 {
		input.ChannelName = aws.String(_mediapackagev2ChannelName)
	}
	if len(_mediapackagev2OriginEndpointName) > 0 {
		input.OriginEndpointName = aws.String(_mediapackagev2OriginEndpointName)
	}

	if resp, err := client.GetOriginEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the specified origin endpoint policy that's configured in AWS
// Elemental MediaPackage.
func mediapackagev2_GetOriginEndpointPolicy(cfg aws.Config, client *mediapackagev2.Client) {
	input := &mediapackagev2.GetOriginEndpointPolicyInput{
		// ChannelGroupName: *string, // Required
		// ChannelName: *string, // Required
		// OriginEndpointName: *string, // Required
	}

	if len(_mediapackagev2ChannelGroupName) > 0 {
		input.ChannelGroupName = aws.String(_mediapackagev2ChannelGroupName)
	}
	if len(_mediapackagev2ChannelName) > 0 {
		input.ChannelName = aws.String(_mediapackagev2ChannelName)
	}
	if len(_mediapackagev2OriginEndpointName) > 0 {
		input.OriginEndpointName = aws.String(_mediapackagev2OriginEndpointName)
	}

	if resp, err := client.GetOriginEndpointPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves all channel groups that are configured in Elemental MediaPackage.
func mediapackagev2_ListChannelGroups(cfg aws.Config, client *mediapackagev2.Client) {
	input := &mediapackagev2.ListChannelGroupsInput{}

	if len(_mediapackagev2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediapackagev2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediapackagev2NextToken) > 0 {
		input.NextToken = aws.String(_mediapackagev2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListChannelGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mediapackagev2.ListChannelGroupsOutput
	p := mediapackagev2.NewListChannelGroupsPaginator(client, input)
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

// Retrieves all channels in a specific channel group that are configured in AWS
// Elemental MediaPackage.
func mediapackagev2_ListChannels(cfg aws.Config, client *mediapackagev2.Client) {
	input := &mediapackagev2.ListChannelsInput{
		// ChannelGroupName: *string, // Required
	}

	if len(_mediapackagev2ChannelGroupName) > 0 {
		input.ChannelGroupName = aws.String(_mediapackagev2ChannelGroupName)
	}
	if len(_mediapackagev2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediapackagev2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediapackagev2NextToken) > 0 {
		input.NextToken = aws.String(_mediapackagev2NextToken)
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

	var results []*mediapackagev2.ListChannelsOutput
	p := mediapackagev2.NewListChannelsPaginator(client, input)
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

// Retrieves a list of harvest jobs that match the specified criteria.
func mediapackagev2_ListHarvestJobs(cfg aws.Config, client *mediapackagev2.Client) {
	input := &mediapackagev2.ListHarvestJobsInput{
		// ChannelGroupName: *string, // Required
	}

	if len(_mediapackagev2ChannelGroupName) > 0 {
		input.ChannelGroupName = aws.String(_mediapackagev2ChannelGroupName)
	}
	if len(_mediapackagev2ChannelName) > 0 {
		input.ChannelName = aws.String(_mediapackagev2ChannelName)
	}
	if len(_mediapackagev2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediapackagev2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediapackagev2NextToken) > 0 {
		input.NextToken = aws.String(_mediapackagev2NextToken)
	}
	if len(_mediapackagev2OriginEndpointName) > 0 {
		input.OriginEndpointName = aws.String(_mediapackagev2OriginEndpointName)
	}
	if len(_mediapackagev2Status) > 0 {
		if err := assignInputField(input, "Status", _mediapackagev2Status); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListHarvestJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mediapackagev2.ListHarvestJobsOutput
	p := mediapackagev2.NewListHarvestJobsPaginator(client, input)
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

// Retrieves all origin endpoints in a specific channel that are configured in AWS
// Elemental MediaPackage.
func mediapackagev2_ListOriginEndpoints(cfg aws.Config, client *mediapackagev2.Client) {
	input := &mediapackagev2.ListOriginEndpointsInput{
		// ChannelGroupName: *string, // Required
		// ChannelName: *string, // Required
	}

	if len(_mediapackagev2ChannelGroupName) > 0 {
		input.ChannelGroupName = aws.String(_mediapackagev2ChannelGroupName)
	}
	if len(_mediapackagev2ChannelName) > 0 {
		input.ChannelName = aws.String(_mediapackagev2ChannelName)
	}
	if len(_mediapackagev2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediapackagev2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediapackagev2NextToken) > 0 {
		input.NextToken = aws.String(_mediapackagev2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListOriginEndpoints(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mediapackagev2.ListOriginEndpointsOutput
	p := mediapackagev2.NewListOriginEndpointsPaginator(client, input)
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

// Lists the tags assigned to a resource.
func mediapackagev2_ListTagsForResource(cfg aws.Config, client *mediapackagev2.Client) {
	input := &mediapackagev2.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_mediapackagev2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_mediapackagev2ResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches an IAM policy to the specified channel. With policies, you can specify
// who has access to AWS resources and what actions they can perform on those
// resources. You can attach only one policy with each request.
func mediapackagev2_PutChannelPolicy(cfg aws.Config, client *mediapackagev2.Client) {
	input := &mediapackagev2.PutChannelPolicyInput{
		// ChannelGroupName: *string, // Required
		// ChannelName: *string, // Required
		// Policy: *string, // Required
	}

	if len(_mediapackagev2ChannelGroupName) > 0 {
		input.ChannelGroupName = aws.String(_mediapackagev2ChannelGroupName)
	}
	if len(_mediapackagev2ChannelName) > 0 {
		input.ChannelName = aws.String(_mediapackagev2ChannelName)
	}
	if len(_mediapackagev2Policy) > 0 {
		input.Policy = aws.String(_mediapackagev2Policy)
	}

	if resp, err := client.PutChannelPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches an IAM policy to the specified origin endpoint. You can attach only
// one policy with each request.
func mediapackagev2_PutOriginEndpointPolicy(cfg aws.Config, client *mediapackagev2.Client) {
	input := &mediapackagev2.PutOriginEndpointPolicyInput{
		// ChannelGroupName: *string, // Required
		// ChannelName: *string, // Required
		// OriginEndpointName: *string, // Required
		// Policy: *string, // Required
	}

	if len(_mediapackagev2ChannelGroupName) > 0 {
		input.ChannelGroupName = aws.String(_mediapackagev2ChannelGroupName)
	}
	if len(_mediapackagev2ChannelName) > 0 {
		input.ChannelName = aws.String(_mediapackagev2ChannelName)
	}
	if len(_mediapackagev2OriginEndpointName) > 0 {
		input.OriginEndpointName = aws.String(_mediapackagev2OriginEndpointName)
	}
	if len(_mediapackagev2Policy) > 0 {
		input.Policy = aws.String(_mediapackagev2Policy)
	}
	if len(_mediapackagev2CdnAuthConfiguration) > 0 {
		if err := assignInputField(input, "CdnAuthConfiguration", _mediapackagev2CdnAuthConfiguration); err != nil {
			log.Errorf("invalid --cdn-auth-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutOriginEndpointPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Resetting the channel can help to clear errors from misconfigurations in the
// encoder. A reset refreshes the ingest stream and removes previous content.
//
// Be sure to stop the encoder before you reset the channel, and wait at least 30
// seconds before you restart the encoder.
func mediapackagev2_ResetChannelState(cfg aws.Config, client *mediapackagev2.Client) {
	input := &mediapackagev2.ResetChannelStateInput{
		// ChannelGroupName: *string, // Required
		// ChannelName: *string, // Required
	}

	if len(_mediapackagev2ChannelGroupName) > 0 {
		input.ChannelGroupName = aws.String(_mediapackagev2ChannelGroupName)
	}
	if len(_mediapackagev2ChannelName) > 0 {
		input.ChannelName = aws.String(_mediapackagev2ChannelName)
	}

	if resp, err := client.ResetChannelState(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Resetting the origin endpoint can help to resolve unexpected behavior and other
// content packaging issues. It also helps to preserve special events when you
// don't want the previous content to be available for viewing. A reset clears out
// all previous content from the origin endpoint.
//
// MediaPackage might return old content from this endpoint in the first 30
// seconds after the endpoint reset. For best results, when possible, wait 30
// seconds from endpoint reset to send playback requests to this endpoint.
func mediapackagev2_ResetOriginEndpointState(cfg aws.Config, client *mediapackagev2.Client) {
	input := &mediapackagev2.ResetOriginEndpointStateInput{
		// ChannelGroupName: *string, // Required
		// ChannelName: *string, // Required
		// OriginEndpointName: *string, // Required
	}

	if len(_mediapackagev2ChannelGroupName) > 0 {
		input.ChannelGroupName = aws.String(_mediapackagev2ChannelGroupName)
	}
	if len(_mediapackagev2ChannelName) > 0 {
		input.ChannelName = aws.String(_mediapackagev2ChannelName)
	}
	if len(_mediapackagev2OriginEndpointName) > 0 {
		input.OriginEndpointName = aws.String(_mediapackagev2OriginEndpointName)
	}

	if resp, err := client.ResetOriginEndpointState(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns one of more tags (key-value pairs) to the specified MediaPackage
// resource.
//
// Tags can help you organize and categorize your resources. You can also use them
// to scope user permissions, by granting a user permission to access or change
// only resources with certain tag values. You can use the TagResource operation
// with a resource that already has tags. If you specify a new tag key for the
// resource, this tag is appended to the list of tags associated with the resource.
// If you specify a tag key that is already associated with the resource, the new
// tag value that you specify replaces the previous value for that tag.
func mediapackagev2_TagResource(cfg aws.Config, client *mediapackagev2.Client) {
	input := &mediapackagev2.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_mediapackagev2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_mediapackagev2ResourceArn)
	}
	if len(_mediapackagev2Tags) > 0 {
		if err := assignInputField(input, "Tags", _mediapackagev2Tags); err != nil {
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

// Removes one or more tags from the specified resource.
func mediapackagev2_UntagResource(cfg aws.Config, client *mediapackagev2.Client) {
	input := &mediapackagev2.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_mediapackagev2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_mediapackagev2ResourceArn)
	}
	if len(_mediapackagev2TagKeys) > 0 {
		input.TagKeys = append([]string(nil), _mediapackagev2TagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the specified channel. You can edit if MediaPackage sends ingest or
// egress access logs to the CloudWatch log group, if content will be encrypted,
// the description on a channel, and your channel's policy settings. You can't edit
// the name of the channel or CloudFront distribution details.
//
// Any edits you make that impact the video output may not be reflected for a few
// minutes.
func mediapackagev2_UpdateChannel(cfg aws.Config, client *mediapackagev2.Client) {
	input := &mediapackagev2.UpdateChannelInput{
		// ChannelGroupName: *string, // Required
		// ChannelName: *string, // Required
	}

	if len(_mediapackagev2ChannelGroupName) > 0 {
		input.ChannelGroupName = aws.String(_mediapackagev2ChannelGroupName)
	}
	if len(_mediapackagev2ChannelName) > 0 {
		input.ChannelName = aws.String(_mediapackagev2ChannelName)
	}
	if len(_mediapackagev2Description) > 0 {
		input.Description = aws.String(_mediapackagev2Description)
	}
	if len(_mediapackagev2ETag) > 0 {
		input.ETag = aws.String(_mediapackagev2ETag)
	}
	if len(_mediapackagev2InputSwitchConfiguration) > 0 {
		if err := assignInputField(input, "InputSwitchConfiguration", _mediapackagev2InputSwitchConfiguration); err != nil {
			log.Errorf("invalid --input-switch-configuration: %s", err.Error())
			return
		}
	}
	if len(_mediapackagev2OutputHeaderConfiguration) > 0 {
		if err := assignInputField(input, "OutputHeaderConfiguration", _mediapackagev2OutputHeaderConfiguration); err != nil {
			log.Errorf("invalid --output-header-configuration: %s", err.Error())
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

// Update the specified channel group. You can edit the description on a channel
// group for easier identification later from the AWS Elemental MediaPackage
// console. You can't edit the name of the channel group.
//
// Any edits you make that impact the video output may not be reflected for a few
// minutes.
func mediapackagev2_UpdateChannelGroup(cfg aws.Config, client *mediapackagev2.Client) {
	input := &mediapackagev2.UpdateChannelGroupInput{
		// ChannelGroupName: *string, // Required
	}

	if len(_mediapackagev2ChannelGroupName) > 0 {
		input.ChannelGroupName = aws.String(_mediapackagev2ChannelGroupName)
	}
	if len(_mediapackagev2Description) > 0 {
		input.Description = aws.String(_mediapackagev2Description)
	}
	if len(_mediapackagev2ETag) > 0 {
		input.ETag = aws.String(_mediapackagev2ETag)
	}

	if resp, err := client.UpdateChannelGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the specified origin endpoint. Edit the packaging preferences on an
// endpoint to optimize the viewing experience. You can't edit the name of the
// endpoint.
//
// Any edits you make that impact the video output may not be reflected for a few
// minutes.
func mediapackagev2_UpdateOriginEndpoint(cfg aws.Config, client *mediapackagev2.Client) {
	input := &mediapackagev2.UpdateOriginEndpointInput{
		// ChannelGroupName: *string, // Required
		// ChannelName: *string, // Required
		// ContainerType: types.ContainerType, // Required
		// OriginEndpointName: *string, // Required
	}

	if len(_mediapackagev2ChannelGroupName) > 0 {
		input.ChannelGroupName = aws.String(_mediapackagev2ChannelGroupName)
	}
	if len(_mediapackagev2ChannelName) > 0 {
		input.ChannelName = aws.String(_mediapackagev2ChannelName)
	}
	if len(_mediapackagev2ContainerType) > 0 {
		if err := assignInputField(input, "ContainerType", _mediapackagev2ContainerType); err != nil {
			log.Errorf("invalid --container-type: %s", err.Error())
			return
		}
	}
	if len(_mediapackagev2OriginEndpointName) > 0 {
		input.OriginEndpointName = aws.String(_mediapackagev2OriginEndpointName)
	}
	if len(_mediapackagev2DashManifests) > 0 {
		if err := assignInputField(input, "DashManifests", _mediapackagev2DashManifests); err != nil {
			log.Errorf("invalid --dash-manifests: %s", err.Error())
			return
		}
	}
	if len(_mediapackagev2Description) > 0 {
		input.Description = aws.String(_mediapackagev2Description)
	}
	if len(_mediapackagev2ETag) > 0 {
		input.ETag = aws.String(_mediapackagev2ETag)
	}
	if len(_mediapackagev2ForceEndpointErrorConfiguration) > 0 {
		if err := assignInputField(input, "ForceEndpointErrorConfiguration", _mediapackagev2ForceEndpointErrorConfiguration); err != nil {
			log.Errorf("invalid --force-endpoint-error-configuration: %s", err.Error())
			return
		}
	}
	if len(_mediapackagev2HlsManifests) > 0 {
		if err := assignInputField(input, "HlsManifests", _mediapackagev2HlsManifests); err != nil {
			log.Errorf("invalid --hls-manifests: %s", err.Error())
			return
		}
	}
	if len(_mediapackagev2LowLatencyHlsManifests) > 0 {
		if err := assignInputField(input, "LowLatencyHlsManifests", _mediapackagev2LowLatencyHlsManifests); err != nil {
			log.Errorf("invalid --low-latency-hls-manifests: %s", err.Error())
			return
		}
	}
	if len(_mediapackagev2MssManifests) > 0 {
		if err := assignInputField(input, "MssManifests", _mediapackagev2MssManifests); err != nil {
			log.Errorf("invalid --mss-manifests: %s", err.Error())
			return
		}
	}
	if len(_mediapackagev2Segment) > 0 {
		if err := assignInputField(input, "Segment", _mediapackagev2Segment); err != nil {
			log.Errorf("invalid --segment: %s", err.Error())
			return
		}
	}
	if len(_mediapackagev2StartoverWindowSeconds) > 0 {
		if err := assignInputField(input, "StartoverWindowSeconds", _mediapackagev2StartoverWindowSeconds); err != nil {
			log.Errorf("invalid --startover-window-seconds: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateOriginEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_mediapackagev2Cmd)
	_mediapackagev2Cmd.Flags().SortFlags = false

	_mediapackagev2Cmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_mediapackagev2Cmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_mediapackagev2Cmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_mediapackagev2Cmd.Flags().StringVarP(&_mediapackagev2CdnAuthConfiguration, "cdn-auth-configuration", "", "", "Cdn Auth Configuration")
	_mediapackagev2Cmd.Flags().StringVarP(&_mediapackagev2ChannelGroupName, "channel-group-name", "", "", "Channel Group Name")
	_mediapackagev2Cmd.Flags().StringVarP(&_mediapackagev2ChannelName, "channel-name", "", "", "Channel Name")
	_mediapackagev2Cmd.Flags().StringVarP(&_mediapackagev2ClientToken, "client-token", "", "", "Client Token")
	_mediapackagev2Cmd.Flags().StringVarP(&_mediapackagev2ContainerType, "container-type", "", "", "Container Type")
	_mediapackagev2Cmd.Flags().StringVarP(&_mediapackagev2DashManifests, "dash-manifests", "", "", "Dash Manifests")
	_mediapackagev2Cmd.Flags().StringVarP(&_mediapackagev2Description, "description", "", "", "Description")
	_mediapackagev2Cmd.Flags().StringVarP(&_mediapackagev2Destination, "destination", "", "", "Destination")
	_mediapackagev2Cmd.Flags().StringVarP(&_mediapackagev2ETag, "etag", "", "", "Etag")
	_mediapackagev2Cmd.Flags().StringVarP(&_mediapackagev2ForceEndpointErrorConfiguration, "force-endpoint-error-configuration", "", "", "Force Endpoint Error Configuration")
	_mediapackagev2Cmd.Flags().StringVarP(&_mediapackagev2HarvestJobName, "harvest-job-name", "", "", "Harvest Job Name")
	_mediapackagev2Cmd.Flags().StringVarP(&_mediapackagev2HarvestedManifests, "harvested-manifests", "", "", "Harvested Manifests")
	_mediapackagev2Cmd.Flags().StringVarP(&_mediapackagev2HlsManifests, "hls-manifests", "", "", "Hls Manifests")
	_mediapackagev2Cmd.Flags().StringVarP(&_mediapackagev2InputSwitchConfiguration, "input-switch-configuration", "", "", "Input Switch Configuration")
	_mediapackagev2Cmd.Flags().StringVarP(&_mediapackagev2InputType, "input-type", "", "", "Input Type")
	_mediapackagev2Cmd.Flags().StringVarP(&_mediapackagev2LowLatencyHlsManifests, "low-latency-hls-manifests", "", "", "Low Latency Hls Manifests")
	_mediapackagev2Cmd.Flags().StringVarP(&_mediapackagev2MaxResults, "max-results", "", "", "Max Results")
	_mediapackagev2Cmd.Flags().StringVarP(&_mediapackagev2MssManifests, "mss-manifests", "", "", "Mss Manifests")
	_mediapackagev2Cmd.Flags().StringVarP(&_mediapackagev2NextToken, "next-token", "", "", "Next Token")
	_mediapackagev2Cmd.Flags().StringVarP(&_mediapackagev2OriginEndpointName, "origin-endpoint-name", "", "", "Origin Endpoint Name")
	_mediapackagev2Cmd.Flags().StringVarP(&_mediapackagev2OutputHeaderConfiguration, "output-header-configuration", "", "", "Output Header Configuration")
	_mediapackagev2Cmd.Flags().StringVarP(&_mediapackagev2Policy, "policy", "", "", "Policy")
	_mediapackagev2Cmd.Flags().StringVarP(&_mediapackagev2ResourceArn, "resource-arn", "", "", "Resource ARN")
	_mediapackagev2Cmd.Flags().StringVarP(&_mediapackagev2ScheduleConfiguration, "schedule-configuration", "", "", "Schedule Configuration")
	_mediapackagev2Cmd.Flags().StringVarP(&_mediapackagev2Segment, "segment", "", "", "Segment")
	_mediapackagev2Cmd.Flags().StringVarP(&_mediapackagev2StartoverWindowSeconds, "startover-window-seconds", "", "", "Startover Window Seconds")
	_mediapackagev2Cmd.Flags().StringVarP(&_mediapackagev2Status, "status", "", "", "Status")
	_mediapackagev2Cmd.Flags().StringSliceVarP(&_mediapackagev2TagKeys, "tag-keys", "", nil, "Tag Keys")
	_mediapackagev2Cmd.Flags().StringVarP(&_mediapackagev2Tags, "tags", "", "", "Tags")

	_mediapackagev2Cmd.Flags().BoolVarP(&_mediapackagev2CancelHarvestJob, "cancel-harvest-job", "", false, "Cancel Harvest Job")
	_mediapackagev2Cmd.Flags().BoolVarP(&_mediapackagev2CreateChannel, "create-channel", "", false, "Create Channel")
	_mediapackagev2Cmd.Flags().BoolVarP(&_mediapackagev2CreateChannelGroup, "create-channel-group", "", false, "Create Channel Group")
	_mediapackagev2Cmd.Flags().BoolVarP(&_mediapackagev2CreateHarvestJob, "create-harvest-job", "", false, "Create Harvest Job")
	_mediapackagev2Cmd.Flags().BoolVarP(&_mediapackagev2CreateOriginEndpoint, "create-origin-endpoint", "", false, "Create Origin Endpoint")
	_mediapackagev2Cmd.Flags().BoolVarP(&_mediapackagev2DeleteChannel, "delete-channel", "", false, "Delete Channel")
	_mediapackagev2Cmd.Flags().BoolVarP(&_mediapackagev2DeleteChannelGroup, "delete-channel-group", "", false, "Delete Channel Group")
	_mediapackagev2Cmd.Flags().BoolVarP(&_mediapackagev2DeleteChannelPolicy, "delete-channel-policy", "", false, "Delete Channel Policy")
	_mediapackagev2Cmd.Flags().BoolVarP(&_mediapackagev2DeleteOriginEndpoint, "delete-origin-endpoint", "", false, "Delete Origin Endpoint")
	_mediapackagev2Cmd.Flags().BoolVarP(&_mediapackagev2DeleteOriginEndpointPolicy, "delete-origin-endpoint-policy", "", false, "Delete Origin Endpoint Policy")
	_mediapackagev2Cmd.Flags().BoolVarP(&_mediapackagev2GetChannel, "get-channel", "", false, "Get Channel")
	_mediapackagev2Cmd.Flags().BoolVarP(&_mediapackagev2GetChannelGroup, "get-channel-group", "", false, "Get Channel Group")
	_mediapackagev2Cmd.Flags().BoolVarP(&_mediapackagev2GetChannelPolicy, "get-channel-policy", "", false, "Get Channel Policy")
	_mediapackagev2Cmd.Flags().BoolVarP(&_mediapackagev2GetHarvestJob, "get-harvest-job", "", false, "Get Harvest Job")
	_mediapackagev2Cmd.Flags().BoolVarP(&_mediapackagev2GetOriginEndpoint, "get-origin-endpoint", "", false, "Get Origin Endpoint")
	_mediapackagev2Cmd.Flags().BoolVarP(&_mediapackagev2GetOriginEndpointPolicy, "get-origin-endpoint-policy", "", false, "Get Origin Endpoint Policy")
	_mediapackagev2Cmd.Flags().BoolVarP(&_mediapackagev2ListChannelGroups, "list-channel-groups", "", false, "List Channel Groups")
	_mediapackagev2Cmd.Flags().BoolVarP(&_mediapackagev2ListChannels, "list-channels", "", false, "List Channels")
	_mediapackagev2Cmd.Flags().BoolVarP(&_mediapackagev2ListHarvestJobs, "list-harvest-jobs", "", false, "List Harvest Jobs")
	_mediapackagev2Cmd.Flags().BoolVarP(&_mediapackagev2ListOriginEndpoints, "list-origin-endpoints", "", false, "List Origin Endpoints")
	_mediapackagev2Cmd.Flags().BoolVarP(&_mediapackagev2ListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_mediapackagev2Cmd.Flags().BoolVarP(&_mediapackagev2PutChannelPolicy, "put-channel-policy", "", false, "Put Channel Policy")
	_mediapackagev2Cmd.Flags().BoolVarP(&_mediapackagev2PutOriginEndpointPolicy, "put-origin-endpoint-policy", "", false, "Put Origin Endpoint Policy")
	_mediapackagev2Cmd.Flags().BoolVarP(&_mediapackagev2ResetChannelState, "reset-channel-state", "", false, "Reset Channel State")
	_mediapackagev2Cmd.Flags().BoolVarP(&_mediapackagev2ResetOriginEndpointState, "reset-origin-endpoint-state", "", false, "Reset Origin Endpoint State")
	_mediapackagev2Cmd.Flags().BoolVarP(&_mediapackagev2TagResource, "tag-resource", "", false, "Tag Resource")
	_mediapackagev2Cmd.Flags().BoolVarP(&_mediapackagev2UntagResource, "untag-resource", "", false, "Untag Resource")
	_mediapackagev2Cmd.Flags().BoolVarP(&_mediapackagev2UpdateChannel, "update-channel", "", false, "Update Channel")
	_mediapackagev2Cmd.Flags().BoolVarP(&_mediapackagev2UpdateChannelGroup, "update-channel-group", "", false, "Update Channel Group")
	_mediapackagev2Cmd.Flags().BoolVarP(&_mediapackagev2UpdateOriginEndpoint, "update-origin-endpoint", "", false, "Update Origin Endpoint")

}
