package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkmessaging"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// chimesdkmessagingCmd represents the chimesdkmessaging command
var _chimesdkmessagingCmd = &cobra.Command{
	Use:   "chimesdkmessaging",
	Short: "AWS chimesdkmessaging CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := chimesdkmessaging.NewFromConfig(cfg)
		if _chimesdkmessagingAssociateChannelFlow {
			chimesdkmessaging_AssociateChannelFlow(cfg, client)
			return
		}
		if _chimesdkmessagingBatchCreateChannelMembership {
			chimesdkmessaging_BatchCreateChannelMembership(cfg, client)
			return
		}
		if _chimesdkmessagingChannelFlowCallback {
			chimesdkmessaging_ChannelFlowCallback(cfg, client)
			return
		}
		if _chimesdkmessagingCreateChannel {
			chimesdkmessaging_CreateChannel(cfg, client)
			return
		}
		if _chimesdkmessagingCreateChannelBan {
			chimesdkmessaging_CreateChannelBan(cfg, client)
			return
		}
		if _chimesdkmessagingCreateChannelFlow {
			chimesdkmessaging_CreateChannelFlow(cfg, client)
			return
		}
		if _chimesdkmessagingCreateChannelMembership {
			chimesdkmessaging_CreateChannelMembership(cfg, client)
			return
		}
		if _chimesdkmessagingCreateChannelModerator {
			chimesdkmessaging_CreateChannelModerator(cfg, client)
			return
		}
		if _chimesdkmessagingDeleteChannel {
			chimesdkmessaging_DeleteChannel(cfg, client)
			return
		}
		if _chimesdkmessagingDeleteChannelBan {
			chimesdkmessaging_DeleteChannelBan(cfg, client)
			return
		}
		if _chimesdkmessagingDeleteChannelFlow {
			chimesdkmessaging_DeleteChannelFlow(cfg, client)
			return
		}
		if _chimesdkmessagingDeleteChannelMembership {
			chimesdkmessaging_DeleteChannelMembership(cfg, client)
			return
		}
		if _chimesdkmessagingDeleteChannelMessage {
			chimesdkmessaging_DeleteChannelMessage(cfg, client)
			return
		}
		if _chimesdkmessagingDeleteChannelModerator {
			chimesdkmessaging_DeleteChannelModerator(cfg, client)
			return
		}
		if _chimesdkmessagingDeleteMessagingStreamingConfigurations {
			chimesdkmessaging_DeleteMessagingStreamingConfigurations(cfg, client)
			return
		}
		if _chimesdkmessagingDescribeChannel {
			chimesdkmessaging_DescribeChannel(cfg, client)
			return
		}
		if _chimesdkmessagingDescribeChannelBan {
			chimesdkmessaging_DescribeChannelBan(cfg, client)
			return
		}
		if _chimesdkmessagingDescribeChannelFlow {
			chimesdkmessaging_DescribeChannelFlow(cfg, client)
			return
		}
		if _chimesdkmessagingDescribeChannelMembership {
			chimesdkmessaging_DescribeChannelMembership(cfg, client)
			return
		}
		if _chimesdkmessagingDescribeChannelMembershipForAppInstanceUser {
			chimesdkmessaging_DescribeChannelMembershipForAppInstanceUser(cfg, client)
			return
		}
		if _chimesdkmessagingDescribeChannelModeratedByAppInstanceUser {
			chimesdkmessaging_DescribeChannelModeratedByAppInstanceUser(cfg, client)
			return
		}
		if _chimesdkmessagingDescribeChannelModerator {
			chimesdkmessaging_DescribeChannelModerator(cfg, client)
			return
		}
		if _chimesdkmessagingDisassociateChannelFlow {
			chimesdkmessaging_DisassociateChannelFlow(cfg, client)
			return
		}
		if _chimesdkmessagingGetChannelMembershipPreferences {
			chimesdkmessaging_GetChannelMembershipPreferences(cfg, client)
			return
		}
		if _chimesdkmessagingGetChannelMessage {
			chimesdkmessaging_GetChannelMessage(cfg, client)
			return
		}
		if _chimesdkmessagingGetChannelMessageStatus {
			chimesdkmessaging_GetChannelMessageStatus(cfg, client)
			return
		}
		if _chimesdkmessagingGetMessagingSessionEndpoint {
			chimesdkmessaging_GetMessagingSessionEndpoint(cfg, client)
			return
		}
		if _chimesdkmessagingGetMessagingStreamingConfigurations {
			chimesdkmessaging_GetMessagingStreamingConfigurations(cfg, client)
			return
		}
		if _chimesdkmessagingListChannelBans {
			chimesdkmessaging_ListChannelBans(cfg, client)
			return
		}
		if _chimesdkmessagingListChannelFlows {
			chimesdkmessaging_ListChannelFlows(cfg, client)
			return
		}
		if _chimesdkmessagingListChannelMemberships {
			chimesdkmessaging_ListChannelMemberships(cfg, client)
			return
		}
		if _chimesdkmessagingListChannelMembershipsForAppInstanceUser {
			chimesdkmessaging_ListChannelMembershipsForAppInstanceUser(cfg, client)
			return
		}
		if _chimesdkmessagingListChannelMessages {
			chimesdkmessaging_ListChannelMessages(cfg, client)
			return
		}
		if _chimesdkmessagingListChannelModerators {
			chimesdkmessaging_ListChannelModerators(cfg, client)
			return
		}
		if _chimesdkmessagingListChannels {
			chimesdkmessaging_ListChannels(cfg, client)
			return
		}
		if _chimesdkmessagingListChannelsAssociatedWithChannelFlow {
			chimesdkmessaging_ListChannelsAssociatedWithChannelFlow(cfg, client)
			return
		}
		if _chimesdkmessagingListChannelsModeratedByAppInstanceUser {
			chimesdkmessaging_ListChannelsModeratedByAppInstanceUser(cfg, client)
			return
		}
		if _chimesdkmessagingListSubChannels {
			chimesdkmessaging_ListSubChannels(cfg, client)
			return
		}
		if _chimesdkmessagingListTagsForResource {
			chimesdkmessaging_ListTagsForResource(cfg, client)
			return
		}
		if _chimesdkmessagingPutChannelExpirationSettings {
			chimesdkmessaging_PutChannelExpirationSettings(cfg, client)
			return
		}
		if _chimesdkmessagingPutChannelMembershipPreferences {
			chimesdkmessaging_PutChannelMembershipPreferences(cfg, client)
			return
		}
		if _chimesdkmessagingPutMessagingStreamingConfigurations {
			chimesdkmessaging_PutMessagingStreamingConfigurations(cfg, client)
			return
		}
		if _chimesdkmessagingRedactChannelMessage {
			chimesdkmessaging_RedactChannelMessage(cfg, client)
			return
		}
		if _chimesdkmessagingSearchChannels {
			chimesdkmessaging_SearchChannels(cfg, client)
			return
		}
		if _chimesdkmessagingSendChannelMessage {
			chimesdkmessaging_SendChannelMessage(cfg, client)
			return
		}
		if _chimesdkmessagingTagResource {
			chimesdkmessaging_TagResource(cfg, client)
			return
		}
		if _chimesdkmessagingUntagResource {
			chimesdkmessaging_UntagResource(cfg, client)
			return
		}
		if _chimesdkmessagingUpdateChannel {
			chimesdkmessaging_UpdateChannel(cfg, client)
			return
		}
		if _chimesdkmessagingUpdateChannelFlow {
			chimesdkmessaging_UpdateChannelFlow(cfg, client)
			return
		}
		if _chimesdkmessagingUpdateChannelMessage {
			chimesdkmessaging_UpdateChannelMessage(cfg, client)
			return
		}
		if _chimesdkmessagingUpdateChannelReadMarker {
			chimesdkmessaging_UpdateChannelReadMarker(cfg, client)
			return
		}

	},
}

var (
	_chimesdkmessagingAssociateChannelFlow                        bool
	_chimesdkmessagingBatchCreateChannelMembership                bool
	_chimesdkmessagingChannelFlowCallback                         bool
	_chimesdkmessagingCreateChannel                               bool
	_chimesdkmessagingCreateChannelBan                            bool
	_chimesdkmessagingCreateChannelFlow                           bool
	_chimesdkmessagingCreateChannelMembership                     bool
	_chimesdkmessagingCreateChannelModerator                      bool
	_chimesdkmessagingDeleteChannel                               bool
	_chimesdkmessagingDeleteChannelBan                            bool
	_chimesdkmessagingDeleteChannelFlow                           bool
	_chimesdkmessagingDeleteChannelMembership                     bool
	_chimesdkmessagingDeleteChannelMessage                        bool
	_chimesdkmessagingDeleteChannelModerator                      bool
	_chimesdkmessagingDeleteMessagingStreamingConfigurations      bool
	_chimesdkmessagingDescribeChannel                             bool
	_chimesdkmessagingDescribeChannelBan                          bool
	_chimesdkmessagingDescribeChannelFlow                         bool
	_chimesdkmessagingDescribeChannelMembership                   bool
	_chimesdkmessagingDescribeChannelMembershipForAppInstanceUser bool
	_chimesdkmessagingDescribeChannelModeratedByAppInstanceUser   bool
	_chimesdkmessagingDescribeChannelModerator                    bool
	_chimesdkmessagingDisassociateChannelFlow                     bool
	_chimesdkmessagingGetChannelMembershipPreferences             bool
	_chimesdkmessagingGetChannelMessage                           bool
	_chimesdkmessagingGetChannelMessageStatus                     bool
	_chimesdkmessagingGetMessagingSessionEndpoint                 bool
	_chimesdkmessagingGetMessagingStreamingConfigurations         bool
	_chimesdkmessagingListChannelBans                             bool
	_chimesdkmessagingListChannelFlows                            bool
	_chimesdkmessagingListChannelMemberships                      bool
	_chimesdkmessagingListChannelMembershipsForAppInstanceUser    bool
	_chimesdkmessagingListChannelMessages                         bool
	_chimesdkmessagingListChannelModerators                       bool
	_chimesdkmessagingListChannels                                bool
	_chimesdkmessagingListChannelsAssociatedWithChannelFlow       bool
	_chimesdkmessagingListChannelsModeratedByAppInstanceUser      bool
	_chimesdkmessagingListSubChannels                             bool
	_chimesdkmessagingListTagsForResource                         bool
	_chimesdkmessagingPutChannelExpirationSettings                bool
	_chimesdkmessagingPutChannelMembershipPreferences             bool
	_chimesdkmessagingPutMessagingStreamingConfigurations         bool
	_chimesdkmessagingRedactChannelMessage                        bool
	_chimesdkmessagingSearchChannels                              bool
	_chimesdkmessagingSendChannelMessage                          bool
	_chimesdkmessagingTagResource                                 bool
	_chimesdkmessagingUntagResource                               bool
	_chimesdkmessagingUpdateChannel                               bool
	_chimesdkmessagingUpdateChannelFlow                           bool
	_chimesdkmessagingUpdateChannelMessage                        bool
	_chimesdkmessagingUpdateChannelReadMarker                     bool

	_chimesdkmessagingAppInstanceArn              string
	_chimesdkmessagingAppInstanceUserArn          string
	_chimesdkmessagingCallbackId                  string
	_chimesdkmessagingChannelArn                  string
	_chimesdkmessagingChannelFlowArn              string
	_chimesdkmessagingChannelId                   string
	_chimesdkmessagingChannelMessage              string
	_chimesdkmessagingChannelModeratorArn         string
	_chimesdkmessagingChimeBearer                 string
	_chimesdkmessagingClientRequestToken          string
	_chimesdkmessagingContent                     string
	_chimesdkmessagingContentType                 string
	_chimesdkmessagingDeleteResource              string
	_chimesdkmessagingElasticChannelConfiguration string
	_chimesdkmessagingExpirationSettings          string
	_chimesdkmessagingFields                      string
	_chimesdkmessagingMaxResults                  string
	_chimesdkmessagingMemberArn                   string
	_chimesdkmessagingMemberArns                  []string
	_chimesdkmessagingMessageAttributes           string
	_chimesdkmessagingMessageId                   string
	_chimesdkmessagingMetadata                    string
	_chimesdkmessagingMode                        string
	_chimesdkmessagingModeratorArns               []string
	_chimesdkmessagingName                        string
	_chimesdkmessagingNetworkType                 string
	_chimesdkmessagingNextToken                   string
	_chimesdkmessagingNotAfter                    string
	_chimesdkmessagingNotBefore                   string
	_chimesdkmessagingPersistence                 string
	_chimesdkmessagingPreferences                 string
	_chimesdkmessagingPrivacy                     string
	_chimesdkmessagingProcessors                  string
	_chimesdkmessagingPushNotification            string
	_chimesdkmessagingResourceARN                 string
	_chimesdkmessagingSortOrder                   string
	_chimesdkmessagingStreamingConfigurations     string
	_chimesdkmessagingSubChannelId                string
	_chimesdkmessagingTagKeys                     []string
	_chimesdkmessagingTags                        string
	_chimesdkmessagingTarget                      string
	_chimesdkmessagingType                        string
)

// Associates a channel flow with a channel. Once associated, all messages to that
// channel go through channel flow processors. To stop processing, use the
// DisassociateChannelFlow API.
//
// Only administrators or channel moderators can associate a channel flow. The
// x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
func chimesdkmessaging_AssociateChannelFlow(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.AssociateChannelFlowInput{
		// ChannelArn: *string, // Required
		// ChannelFlowArn: *string, // Required
		// ChimeBearer: *string, // Required
	}

	if len(_chimesdkmessagingChannelArn) > 0 {
		input.ChannelArn = aws.String(_chimesdkmessagingChannelArn)
	}
	if len(_chimesdkmessagingChannelFlowArn) > 0 {
		input.ChannelFlowArn = aws.String(_chimesdkmessagingChannelFlowArn)
	}
	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}

	if resp, err := client.AssociateChannelFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a specified number of users and bots to a channel.
func chimesdkmessaging_BatchCreateChannelMembership(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.BatchCreateChannelMembershipInput{
		// ChannelArn: *string, // Required
		// ChimeBearer: *string, // Required
		// MemberArns: []string, // Required
	}

	if len(_chimesdkmessagingChannelArn) > 0 {
		input.ChannelArn = aws.String(_chimesdkmessagingChannelArn)
	}
	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}
	if len(_chimesdkmessagingMemberArns) > 0 {
		input.MemberArns = append([]string(nil), _chimesdkmessagingMemberArns...)
	}
	if len(_chimesdkmessagingSubChannelId) > 0 {
		input.SubChannelId = aws.String(_chimesdkmessagingSubChannelId)
	}
	if len(_chimesdkmessagingType) > 0 {
		if err := assignInputField(input, "Type", _chimesdkmessagingType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchCreateChannelMembership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Calls back Amazon Chime SDK messaging with a processing response message. This
// should be invoked from the processor Lambda. This is a developer API.
//
// You can return one of the following processing responses:
//
// - Update message content or metadata
//
// - Deny a message
//
// - Make no changes to the message
func chimesdkmessaging_ChannelFlowCallback(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.ChannelFlowCallbackInput{
		// CallbackId: *string, // Required
		// ChannelArn: *string, // Required
		// ChannelMessage: *types.ChannelMessageCallback, // Required
	}

	if len(_chimesdkmessagingCallbackId) > 0 {
		input.CallbackId = aws.String(_chimesdkmessagingCallbackId)
	}
	if len(_chimesdkmessagingChannelArn) > 0 {
		input.ChannelArn = aws.String(_chimesdkmessagingChannelArn)
	}
	if len(_chimesdkmessagingChannelMessage) > 0 {
		if err := assignInputField(input, "ChannelMessage", _chimesdkmessagingChannelMessage); err != nil {
			log.Errorf("invalid --channel-message: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmessagingDeleteResource) > 0 {
		if err := assignInputField(input, "DeleteResource", _chimesdkmessagingDeleteResource); err != nil {
			log.Errorf("invalid --delete-resource: %s", err.Error())
			return
		}
	}

	if resp, err := client.ChannelFlowCallback(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a channel to which you can add users and send messages.
// Restriction: You can't change a channel's privacy.
//
// The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
func chimesdkmessaging_CreateChannel(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.CreateChannelInput{
		// AppInstanceArn: *string, // Required
		// ChimeBearer: *string, // Required
		// ClientRequestToken: *string, // Required
		// Name: *string, // Required
	}

	if len(_chimesdkmessagingAppInstanceArn) > 0 {
		input.AppInstanceArn = aws.String(_chimesdkmessagingAppInstanceArn)
	}
	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}
	if len(_chimesdkmessagingClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_chimesdkmessagingClientRequestToken)
	}
	if len(_chimesdkmessagingName) > 0 {
		input.Name = aws.String(_chimesdkmessagingName)
	}
	if len(_chimesdkmessagingChannelId) > 0 {
		input.ChannelId = aws.String(_chimesdkmessagingChannelId)
	}
	if len(_chimesdkmessagingElasticChannelConfiguration) > 0 {
		if err := assignInputField(input, "ElasticChannelConfiguration", _chimesdkmessagingElasticChannelConfiguration); err != nil {
			log.Errorf("invalid --elastic-channel-configuration: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmessagingExpirationSettings) > 0 {
		if err := assignInputField(input, "ExpirationSettings", _chimesdkmessagingExpirationSettings); err != nil {
			log.Errorf("invalid --expiration-settings: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmessagingMemberArns) > 0 {
		input.MemberArns = append([]string(nil), _chimesdkmessagingMemberArns...)
	}
	if len(_chimesdkmessagingMetadata) > 0 {
		input.Metadata = aws.String(_chimesdkmessagingMetadata)
	}
	if len(_chimesdkmessagingMode) > 0 {
		if err := assignInputField(input, "Mode", _chimesdkmessagingMode); err != nil {
			log.Errorf("invalid --mode: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmessagingModeratorArns) > 0 {
		input.ModeratorArns = append([]string(nil), _chimesdkmessagingModeratorArns...)
	}
	if len(_chimesdkmessagingPrivacy) > 0 {
		if err := assignInputField(input, "Privacy", _chimesdkmessagingPrivacy); err != nil {
			log.Errorf("invalid --privacy: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmessagingTags) > 0 {
		if err := assignInputField(input, "Tags", _chimesdkmessagingTags); err != nil {
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

// Permanently bans a member from a channel. Moderators can't add banned members
// to a channel. To undo a ban, you first have to DeleteChannelBan , and then
// CreateChannelMembership . Bans are cleaned up when you delete users or channels.
//
// If you ban a user who is already part of a channel, that user is automatically
// kicked from the channel.
//
// The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
func chimesdkmessaging_CreateChannelBan(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.CreateChannelBanInput{
		// ChannelArn: *string, // Required
		// ChimeBearer: *string, // Required
		// MemberArn: *string, // Required
	}

	if len(_chimesdkmessagingChannelArn) > 0 {
		input.ChannelArn = aws.String(_chimesdkmessagingChannelArn)
	}
	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}
	if len(_chimesdkmessagingMemberArn) > 0 {
		input.MemberArn = aws.String(_chimesdkmessagingMemberArn)
	}

	if resp, err := client.CreateChannelBan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a channel flow, a container for processors. Processors are AWS Lambda
// functions that perform actions on chat messages, such as stripping out
// profanity. You can associate channel flows with channels, and the processors in
// the channel flow then take action on all messages sent to that channel. This is
// a developer API.
//
// Channel flows process the following items:
//
// - New and updated messages
//
// - Persistent and non-persistent messages
//
// - The Standard message type
//
// Channel flows don't process Control or System messages. For more information
// about the message types provided by Chime SDK messaging, refer to [Message types]in the Amazon
// Chime developer guide.
//
// [Message types]: https://docs.aws.amazon.com/chime-sdk/latest/dg/using-the-messaging-sdk.html#msg-types
func chimesdkmessaging_CreateChannelFlow(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.CreateChannelFlowInput{
		// AppInstanceArn: *string, // Required
		// ClientRequestToken: *string, // Required
		// Name: *string, // Required
		// Processors: []types.Processor, // Required
	}

	if len(_chimesdkmessagingAppInstanceArn) > 0 {
		input.AppInstanceArn = aws.String(_chimesdkmessagingAppInstanceArn)
	}
	if len(_chimesdkmessagingClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_chimesdkmessagingClientRequestToken)
	}
	if len(_chimesdkmessagingName) > 0 {
		input.Name = aws.String(_chimesdkmessagingName)
	}
	if len(_chimesdkmessagingProcessors) > 0 {
		if err := assignInputField(input, "Processors", _chimesdkmessagingProcessors); err != nil {
			log.Errorf("invalid --processors: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmessagingTags) > 0 {
		if err := assignInputField(input, "Tags", _chimesdkmessagingTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateChannelFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a member to a channel. The InvitedBy field in ChannelMembership is derived
// from the request header. A channel member can:
//
// - List messages
//
// - Send messages
//
// - Receive messages
//
// - Edit their own messages
//
// - Leave the channel
//
// Privacy settings impact this action as follows:
//
// - Public Channels: You do not need to be a member to list messages, but you
// must be a member to send messages.
//
// - Private Channels: You must be a member to list or send messages.
//
// The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUserArn or AppInstanceBot that makes the API call as the value in
// the header.
func chimesdkmessaging_CreateChannelMembership(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.CreateChannelMembershipInput{
		// ChannelArn: *string, // Required
		// ChimeBearer: *string, // Required
		// MemberArn: *string, // Required
		// Type: types.ChannelMembershipType, // Required
	}

	if len(_chimesdkmessagingChannelArn) > 0 {
		input.ChannelArn = aws.String(_chimesdkmessagingChannelArn)
	}
	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}
	if len(_chimesdkmessagingMemberArn) > 0 {
		input.MemberArn = aws.String(_chimesdkmessagingMemberArn)
	}
	if len(_chimesdkmessagingType) > 0 {
		if err := assignInputField(input, "Type", _chimesdkmessagingType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmessagingSubChannelId) > 0 {
		input.SubChannelId = aws.String(_chimesdkmessagingSubChannelId)
	}

	if resp, err := client.CreateChannelMembership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new ChannelModerator . A channel moderator can:
// - Add and remove other members of the channel.
//
// - Add and remove other moderators of the channel.
//
// - Add and remove user bans for the channel.
//
// - Redact messages in the channel.
//
// - List messages in the channel.
//
// The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot of the user that makes the API call as the
// value in the header.
func chimesdkmessaging_CreateChannelModerator(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.CreateChannelModeratorInput{
		// ChannelArn: *string, // Required
		// ChannelModeratorArn: *string, // Required
		// ChimeBearer: *string, // Required
	}

	if len(_chimesdkmessagingChannelArn) > 0 {
		input.ChannelArn = aws.String(_chimesdkmessagingChannelArn)
	}
	if len(_chimesdkmessagingChannelModeratorArn) > 0 {
		input.ChannelModeratorArn = aws.String(_chimesdkmessagingChannelModeratorArn)
	}
	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}

	if resp, err := client.CreateChannelModerator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Immediately makes a channel and its memberships inaccessible and marks them for
// deletion. This is an irreversible process.
//
// The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUserArn or AppInstanceBot that makes the API call as the value in
// the header.
func chimesdkmessaging_DeleteChannel(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.DeleteChannelInput{
		// ChannelArn: *string, // Required
		// ChimeBearer: *string, // Required
	}

	if len(_chimesdkmessagingChannelArn) > 0 {
		input.ChannelArn = aws.String(_chimesdkmessagingChannelArn)
	}
	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}

	if resp, err := client.DeleteChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a member from a channel's ban list.
// The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
func chimesdkmessaging_DeleteChannelBan(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.DeleteChannelBanInput{
		// ChannelArn: *string, // Required
		// ChimeBearer: *string, // Required
		// MemberArn: *string, // Required
	}

	if len(_chimesdkmessagingChannelArn) > 0 {
		input.ChannelArn = aws.String(_chimesdkmessagingChannelArn)
	}
	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}
	if len(_chimesdkmessagingMemberArn) > 0 {
		input.MemberArn = aws.String(_chimesdkmessagingMemberArn)
	}

	if resp, err := client.DeleteChannelBan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a channel flow, an irreversible process. This is a developer API.
// This API works only when the channel flow is not associated with any channel.
// To get a list of all channels that a channel flow is associated with, use the
// ListChannelsAssociatedWithChannelFlow API. Use the DisassociateChannelFlow API
// to disassociate a channel flow from all channels.
func chimesdkmessaging_DeleteChannelFlow(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.DeleteChannelFlowInput{
		// ChannelFlowArn: *string, // Required
	}

	if len(_chimesdkmessagingChannelFlowArn) > 0 {
		input.ChannelFlowArn = aws.String(_chimesdkmessagingChannelFlowArn)
	}

	if resp, err := client.DeleteChannelFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a member from a channel.
// The x-amz-chime-bearer request header is mandatory. Use the AppInstanceUserArn
// of the user that makes the API call as the value in the header.
func chimesdkmessaging_DeleteChannelMembership(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.DeleteChannelMembershipInput{
		// ChannelArn: *string, // Required
		// ChimeBearer: *string, // Required
		// MemberArn: *string, // Required
	}

	if len(_chimesdkmessagingChannelArn) > 0 {
		input.ChannelArn = aws.String(_chimesdkmessagingChannelArn)
	}
	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}
	if len(_chimesdkmessagingMemberArn) > 0 {
		input.MemberArn = aws.String(_chimesdkmessagingMemberArn)
	}
	if len(_chimesdkmessagingSubChannelId) > 0 {
		input.SubChannelId = aws.String(_chimesdkmessagingSubChannelId)
	}

	if resp, err := client.DeleteChannelMembership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a channel message. Only admins can perform this action. Deletion makes
// messages inaccessible immediately. A background process deletes any revisions
// created by UpdateChannelMessage .
//
// The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
func chimesdkmessaging_DeleteChannelMessage(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.DeleteChannelMessageInput{
		// ChannelArn: *string, // Required
		// ChimeBearer: *string, // Required
		// MessageId: *string, // Required
	}

	if len(_chimesdkmessagingChannelArn) > 0 {
		input.ChannelArn = aws.String(_chimesdkmessagingChannelArn)
	}
	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}
	if len(_chimesdkmessagingMessageId) > 0 {
		input.MessageId = aws.String(_chimesdkmessagingMessageId)
	}
	if len(_chimesdkmessagingSubChannelId) > 0 {
		input.SubChannelId = aws.String(_chimesdkmessagingSubChannelId)
	}

	if resp, err := client.DeleteChannelMessage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a channel moderator.
// The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
func chimesdkmessaging_DeleteChannelModerator(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.DeleteChannelModeratorInput{
		// ChannelArn: *string, // Required
		// ChannelModeratorArn: *string, // Required
		// ChimeBearer: *string, // Required
	}

	if len(_chimesdkmessagingChannelArn) > 0 {
		input.ChannelArn = aws.String(_chimesdkmessagingChannelArn)
	}
	if len(_chimesdkmessagingChannelModeratorArn) > 0 {
		input.ChannelModeratorArn = aws.String(_chimesdkmessagingChannelModeratorArn)
	}
	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}

	if resp, err := client.DeleteChannelModerator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the streaming configurations for an AppInstance . For more information,
// see [Streaming messaging data]in the Amazon Chime SDK Developer Guide.
//
// [Streaming messaging data]: https://docs.aws.amazon.com/chime-sdk/latest/dg/streaming-export.html
func chimesdkmessaging_DeleteMessagingStreamingConfigurations(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.DeleteMessagingStreamingConfigurationsInput{
		// AppInstanceArn: *string, // Required
	}

	if len(_chimesdkmessagingAppInstanceArn) > 0 {
		input.AppInstanceArn = aws.String(_chimesdkmessagingAppInstanceArn)
	}

	if resp, err := client.DeleteMessagingStreamingConfigurations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the full details of a channel in an Amazon Chime AppInstance .
// The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
func chimesdkmessaging_DescribeChannel(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.DescribeChannelInput{
		// ChannelArn: *string, // Required
		// ChimeBearer: *string, // Required
	}

	if len(_chimesdkmessagingChannelArn) > 0 {
		input.ChannelArn = aws.String(_chimesdkmessagingChannelArn)
	}
	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}

	if resp, err := client.DescribeChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the full details of a channel ban.
// The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
func chimesdkmessaging_DescribeChannelBan(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.DescribeChannelBanInput{
		// ChannelArn: *string, // Required
		// ChimeBearer: *string, // Required
		// MemberArn: *string, // Required
	}

	if len(_chimesdkmessagingChannelArn) > 0 {
		input.ChannelArn = aws.String(_chimesdkmessagingChannelArn)
	}
	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}
	if len(_chimesdkmessagingMemberArn) > 0 {
		input.MemberArn = aws.String(_chimesdkmessagingMemberArn)
	}

	if resp, err := client.DescribeChannelBan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the full details of a channel flow in an Amazon Chime AppInstance . This
// is a developer API.
func chimesdkmessaging_DescribeChannelFlow(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.DescribeChannelFlowInput{
		// ChannelFlowArn: *string, // Required
	}

	if len(_chimesdkmessagingChannelFlowArn) > 0 {
		input.ChannelFlowArn = aws.String(_chimesdkmessagingChannelFlowArn)
	}

	if resp, err := client.DescribeChannelFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the full details of a user's channel membership.
// The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
func chimesdkmessaging_DescribeChannelMembership(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.DescribeChannelMembershipInput{
		// ChannelArn: *string, // Required
		// ChimeBearer: *string, // Required
		// MemberArn: *string, // Required
	}

	if len(_chimesdkmessagingChannelArn) > 0 {
		input.ChannelArn = aws.String(_chimesdkmessagingChannelArn)
	}
	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}
	if len(_chimesdkmessagingMemberArn) > 0 {
		input.MemberArn = aws.String(_chimesdkmessagingMemberArn)
	}
	if len(_chimesdkmessagingSubChannelId) > 0 {
		input.SubChannelId = aws.String(_chimesdkmessagingSubChannelId)
	}

	if resp, err := client.DescribeChannelMembership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the details of a channel based on the membership of the specified
// AppInstanceUser or AppInstanceBot .
//
// The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
func chimesdkmessaging_DescribeChannelMembershipForAppInstanceUser(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.DescribeChannelMembershipForAppInstanceUserInput{
		// AppInstanceUserArn: *string, // Required
		// ChannelArn: *string, // Required
		// ChimeBearer: *string, // Required
	}

	if len(_chimesdkmessagingAppInstanceUserArn) > 0 {
		input.AppInstanceUserArn = aws.String(_chimesdkmessagingAppInstanceUserArn)
	}
	if len(_chimesdkmessagingChannelArn) > 0 {
		input.ChannelArn = aws.String(_chimesdkmessagingChannelArn)
	}
	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}

	if resp, err := client.DescribeChannelMembershipForAppInstanceUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the full details of a channel moderated by the specified AppInstanceUser
// or AppInstanceBot .
//
// The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
func chimesdkmessaging_DescribeChannelModeratedByAppInstanceUser(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.DescribeChannelModeratedByAppInstanceUserInput{
		// AppInstanceUserArn: *string, // Required
		// ChannelArn: *string, // Required
		// ChimeBearer: *string, // Required
	}

	if len(_chimesdkmessagingAppInstanceUserArn) > 0 {
		input.AppInstanceUserArn = aws.String(_chimesdkmessagingAppInstanceUserArn)
	}
	if len(_chimesdkmessagingChannelArn) > 0 {
		input.ChannelArn = aws.String(_chimesdkmessagingChannelArn)
	}
	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}

	if resp, err := client.DescribeChannelModeratedByAppInstanceUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the full details of a single ChannelModerator.
// The x-amz-chime-bearer request header is mandatory. Use the AppInstanceUserArn
// of the user that makes the API call as the value in the header.
func chimesdkmessaging_DescribeChannelModerator(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.DescribeChannelModeratorInput{
		// ChannelArn: *string, // Required
		// ChannelModeratorArn: *string, // Required
		// ChimeBearer: *string, // Required
	}

	if len(_chimesdkmessagingChannelArn) > 0 {
		input.ChannelArn = aws.String(_chimesdkmessagingChannelArn)
	}
	if len(_chimesdkmessagingChannelModeratorArn) > 0 {
		input.ChannelModeratorArn = aws.String(_chimesdkmessagingChannelModeratorArn)
	}
	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}

	if resp, err := client.DescribeChannelModerator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a channel flow from all its channels. Once disassociated, all
// messages to that channel stop going through the channel flow processor.
//
// Only administrators or channel moderators can disassociate a channel flow.
//
// The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
func chimesdkmessaging_DisassociateChannelFlow(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.DisassociateChannelFlowInput{
		// ChannelArn: *string, // Required
		// ChannelFlowArn: *string, // Required
		// ChimeBearer: *string, // Required
	}

	if len(_chimesdkmessagingChannelArn) > 0 {
		input.ChannelArn = aws.String(_chimesdkmessagingChannelArn)
	}
	if len(_chimesdkmessagingChannelFlowArn) > 0 {
		input.ChannelFlowArn = aws.String(_chimesdkmessagingChannelFlowArn)
	}
	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}

	if resp, err := client.DisassociateChannelFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the membership preferences of an AppInstanceUser or AppInstanceBot for the
// specified channel. A user or a bot must be a member of the channel and own the
// membership in order to retrieve membership preferences. Users or bots in the
// AppInstanceAdmin and channel moderator roles can't retrieve preferences for
// other users or bots. Banned users or bots can't retrieve membership preferences
// for the channel from which they are banned.
//
// The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
func chimesdkmessaging_GetChannelMembershipPreferences(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.GetChannelMembershipPreferencesInput{
		// ChannelArn: *string, // Required
		// ChimeBearer: *string, // Required
		// MemberArn: *string, // Required
	}

	if len(_chimesdkmessagingChannelArn) > 0 {
		input.ChannelArn = aws.String(_chimesdkmessagingChannelArn)
	}
	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}
	if len(_chimesdkmessagingMemberArn) > 0 {
		input.MemberArn = aws.String(_chimesdkmessagingMemberArn)
	}

	if resp, err := client.GetChannelMembershipPreferences(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the full details of a channel message.
// The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
func chimesdkmessaging_GetChannelMessage(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.GetChannelMessageInput{
		// ChannelArn: *string, // Required
		// ChimeBearer: *string, // Required
		// MessageId: *string, // Required
	}

	if len(_chimesdkmessagingChannelArn) > 0 {
		input.ChannelArn = aws.String(_chimesdkmessagingChannelArn)
	}
	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}
	if len(_chimesdkmessagingMessageId) > 0 {
		input.MessageId = aws.String(_chimesdkmessagingMessageId)
	}
	if len(_chimesdkmessagingSubChannelId) > 0 {
		input.SubChannelId = aws.String(_chimesdkmessagingSubChannelId)
	}

	if resp, err := client.GetChannelMessage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets message status for a specified messageId . Use this API to determine the
// intermediate status of messages going through channel flow processing. The API
// provides an alternative to retrieving message status if the event was not
// received because a client wasn't connected to a websocket.
//
// Messages can have any one of these statuses.
//
// # SENT Message processed successfully
//
// # PENDING Ongoing processing
//
// # FAILED Processing failed
//
// # DENIED Message denied by the processor
//
// - This API does not return statuses for denied messages, because we don't
// store them once the processor denies them.
//
// - Only the message sender can invoke this API.
//
// - The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
func chimesdkmessaging_GetChannelMessageStatus(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.GetChannelMessageStatusInput{
		// ChannelArn: *string, // Required
		// ChimeBearer: *string, // Required
		// MessageId: *string, // Required
	}

	if len(_chimesdkmessagingChannelArn) > 0 {
		input.ChannelArn = aws.String(_chimesdkmessagingChannelArn)
	}
	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}
	if len(_chimesdkmessagingMessageId) > 0 {
		input.MessageId = aws.String(_chimesdkmessagingMessageId)
	}
	if len(_chimesdkmessagingSubChannelId) > 0 {
		input.SubChannelId = aws.String(_chimesdkmessagingSubChannelId)
	}

	if resp, err := client.GetChannelMessageStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The details of the endpoint for the messaging session.
func chimesdkmessaging_GetMessagingSessionEndpoint(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.GetMessagingSessionEndpointInput{}

	if len(_chimesdkmessagingNetworkType) > 0 {
		if err := assignInputField(input, "NetworkType", _chimesdkmessagingNetworkType); err != nil {
			log.Errorf("invalid --network-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetMessagingSessionEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the data streaming configuration for an AppInstance . For more
// information, see [Streaming messaging data]in the Amazon Chime SDK Developer Guide.
//
// [Streaming messaging data]: https://docs.aws.amazon.com/chime-sdk/latest/dg/streaming-export.html
func chimesdkmessaging_GetMessagingStreamingConfigurations(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.GetMessagingStreamingConfigurationsInput{
		// AppInstanceArn: *string, // Required
	}

	if len(_chimesdkmessagingAppInstanceArn) > 0 {
		input.AppInstanceArn = aws.String(_chimesdkmessagingAppInstanceArn)
	}

	if resp, err := client.GetMessagingStreamingConfigurations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all the users and bots banned from a particular channel.
// The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
func chimesdkmessaging_ListChannelBans(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.ListChannelBansInput{
		// ChannelArn: *string, // Required
		// ChimeBearer: *string, // Required
	}

	if len(_chimesdkmessagingChannelArn) > 0 {
		input.ChannelArn = aws.String(_chimesdkmessagingChannelArn)
	}
	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}
	if len(_chimesdkmessagingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimesdkmessagingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmessagingNextToken) > 0 {
		input.NextToken = aws.String(_chimesdkmessagingNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListChannelBans(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chimesdkmessaging.ListChannelBansOutput
	p := chimesdkmessaging.NewListChannelBansPaginator(client, input)
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

// Returns a paginated lists of all the channel flows created under a single
// Chime. This is a developer API.
func chimesdkmessaging_ListChannelFlows(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.ListChannelFlowsInput{
		// AppInstanceArn: *string, // Required
	}

	if len(_chimesdkmessagingAppInstanceArn) > 0 {
		input.AppInstanceArn = aws.String(_chimesdkmessagingAppInstanceArn)
	}
	if len(_chimesdkmessagingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimesdkmessagingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmessagingNextToken) > 0 {
		input.NextToken = aws.String(_chimesdkmessagingNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListChannelFlows(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chimesdkmessaging.ListChannelFlowsOutput
	p := chimesdkmessaging.NewListChannelFlowsPaginator(client, input)
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

// Lists all channel memberships in a channel.
// The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
//
// If you want to list the channels to which a specific app instance user belongs,
// see the [ListChannelMembershipsForAppInstanceUser]API.
//
// [ListChannelMembershipsForAppInstanceUser]: https://docs.aws.amazon.com/chime-sdk/latest/APIReference/API_messaging-chime_ListChannelMembershipsForAppInstanceUser.html
func chimesdkmessaging_ListChannelMemberships(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.ListChannelMembershipsInput{
		// ChannelArn: *string, // Required
		// ChimeBearer: *string, // Required
	}

	if len(_chimesdkmessagingChannelArn) > 0 {
		input.ChannelArn = aws.String(_chimesdkmessagingChannelArn)
	}
	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}
	if len(_chimesdkmessagingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimesdkmessagingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmessagingNextToken) > 0 {
		input.NextToken = aws.String(_chimesdkmessagingNextToken)
	}
	if len(_chimesdkmessagingSubChannelId) > 0 {
		input.SubChannelId = aws.String(_chimesdkmessagingSubChannelId)
	}
	if len(_chimesdkmessagingType) > 0 {
		if err := assignInputField(input, "Type", _chimesdkmessagingType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListChannelMemberships(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chimesdkmessaging.ListChannelMembershipsOutput
	p := chimesdkmessaging.NewListChannelMembershipsPaginator(client, input)
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

// Lists all channels that an AppInstanceUser or AppInstanceBot is a part of.
// Only an AppInstanceAdmin can call the API with a user ARN that is not their
// own.
//
// The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
func chimesdkmessaging_ListChannelMembershipsForAppInstanceUser(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.ListChannelMembershipsForAppInstanceUserInput{
		// ChimeBearer: *string, // Required
	}

	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}
	if len(_chimesdkmessagingAppInstanceUserArn) > 0 {
		input.AppInstanceUserArn = aws.String(_chimesdkmessagingAppInstanceUserArn)
	}
	if len(_chimesdkmessagingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimesdkmessagingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmessagingNextToken) > 0 {
		input.NextToken = aws.String(_chimesdkmessagingNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListChannelMembershipsForAppInstanceUser(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chimesdkmessaging.ListChannelMembershipsForAppInstanceUserOutput
	p := chimesdkmessaging.NewListChannelMembershipsForAppInstanceUserPaginator(client, input)
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

// List all the messages in a channel. Returns a paginated list of ChannelMessages
// . By default, sorted by creation timestamp in descending order.
//
// Redacted messages appear in the results as empty, since they are only redacted,
// not deleted. Deleted messages do not appear in the results. This action always
// returns the latest version of an edited message.
//
// Also, the x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
func chimesdkmessaging_ListChannelMessages(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.ListChannelMessagesInput{
		// ChannelArn: *string, // Required
		// ChimeBearer: *string, // Required
	}

	if len(_chimesdkmessagingChannelArn) > 0 {
		input.ChannelArn = aws.String(_chimesdkmessagingChannelArn)
	}
	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}
	if len(_chimesdkmessagingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimesdkmessagingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmessagingNextToken) > 0 {
		input.NextToken = aws.String(_chimesdkmessagingNextToken)
	}
	if len(_chimesdkmessagingNotAfter) > 0 {
		if err := assignInputField(input, "NotAfter", _chimesdkmessagingNotAfter); err != nil {
			log.Errorf("invalid --not-after: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmessagingNotBefore) > 0 {
		if err := assignInputField(input, "NotBefore", _chimesdkmessagingNotBefore); err != nil {
			log.Errorf("invalid --not-before: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmessagingSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _chimesdkmessagingSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmessagingSubChannelId) > 0 {
		input.SubChannelId = aws.String(_chimesdkmessagingSubChannelId)
	}

	if disablePaginator() {
		if resp, err := client.ListChannelMessages(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chimesdkmessaging.ListChannelMessagesOutput
	p := chimesdkmessaging.NewListChannelMessagesPaginator(client, input)
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

// Lists all the moderators for a channel.
// The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
func chimesdkmessaging_ListChannelModerators(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.ListChannelModeratorsInput{
		// ChannelArn: *string, // Required
		// ChimeBearer: *string, // Required
	}

	if len(_chimesdkmessagingChannelArn) > 0 {
		input.ChannelArn = aws.String(_chimesdkmessagingChannelArn)
	}
	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}
	if len(_chimesdkmessagingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimesdkmessagingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmessagingNextToken) > 0 {
		input.NextToken = aws.String(_chimesdkmessagingNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListChannelModerators(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chimesdkmessaging.ListChannelModeratorsOutput
	p := chimesdkmessaging.NewListChannelModeratorsPaginator(client, input)
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

// Lists all Channels created under a single Chime App as a paginated list. You
// can specify filters to narrow results.
//
// Functionality & restrictions
//
// - Use privacy = PUBLIC to retrieve all public channels in the account.
//
// - Only an AppInstanceAdmin can set privacy = PRIVATE to list the private
// channels in an account.
//
// The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
func chimesdkmessaging_ListChannels(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.ListChannelsInput{
		// AppInstanceArn: *string, // Required
		// ChimeBearer: *string, // Required
	}

	if len(_chimesdkmessagingAppInstanceArn) > 0 {
		input.AppInstanceArn = aws.String(_chimesdkmessagingAppInstanceArn)
	}
	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}
	if len(_chimesdkmessagingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimesdkmessagingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmessagingNextToken) > 0 {
		input.NextToken = aws.String(_chimesdkmessagingNextToken)
	}
	if len(_chimesdkmessagingPrivacy) > 0 {
		if err := assignInputField(input, "Privacy", _chimesdkmessagingPrivacy); err != nil {
			log.Errorf("invalid --privacy: %s", err.Error())
			return
		}
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

	var results []*chimesdkmessaging.ListChannelsOutput
	p := chimesdkmessaging.NewListChannelsPaginator(client, input)
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

// Lists all channels associated with a specified channel flow. You can associate
// a channel flow with multiple channels, but you can only associate a channel with
// one channel flow. This is a developer API.
func chimesdkmessaging_ListChannelsAssociatedWithChannelFlow(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.ListChannelsAssociatedWithChannelFlowInput{
		// ChannelFlowArn: *string, // Required
	}

	if len(_chimesdkmessagingChannelFlowArn) > 0 {
		input.ChannelFlowArn = aws.String(_chimesdkmessagingChannelFlowArn)
	}
	if len(_chimesdkmessagingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimesdkmessagingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmessagingNextToken) > 0 {
		input.NextToken = aws.String(_chimesdkmessagingNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListChannelsAssociatedWithChannelFlow(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chimesdkmessaging.ListChannelsAssociatedWithChannelFlowOutput
	p := chimesdkmessaging.NewListChannelsAssociatedWithChannelFlowPaginator(client, input)
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

// A list of the channels moderated by an AppInstanceUser .
// The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
func chimesdkmessaging_ListChannelsModeratedByAppInstanceUser(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.ListChannelsModeratedByAppInstanceUserInput{
		// ChimeBearer: *string, // Required
	}

	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}
	if len(_chimesdkmessagingAppInstanceUserArn) > 0 {
		input.AppInstanceUserArn = aws.String(_chimesdkmessagingAppInstanceUserArn)
	}
	if len(_chimesdkmessagingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimesdkmessagingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmessagingNextToken) > 0 {
		input.NextToken = aws.String(_chimesdkmessagingNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListChannelsModeratedByAppInstanceUser(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chimesdkmessaging.ListChannelsModeratedByAppInstanceUserOutput
	p := chimesdkmessaging.NewListChannelsModeratedByAppInstanceUserPaginator(client, input)
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

// Lists all the SubChannels in an elastic channel when given a channel ID.
// Available only to the app instance admins and channel moderators of elastic
// channels.
func chimesdkmessaging_ListSubChannels(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.ListSubChannelsInput{
		// ChannelArn: *string, // Required
		// ChimeBearer: *string, // Required
	}

	if len(_chimesdkmessagingChannelArn) > 0 {
		input.ChannelArn = aws.String(_chimesdkmessagingChannelArn)
	}
	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}
	if len(_chimesdkmessagingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimesdkmessagingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmessagingNextToken) > 0 {
		input.NextToken = aws.String(_chimesdkmessagingNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSubChannels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chimesdkmessaging.ListSubChannelsOutput
	p := chimesdkmessaging.NewListSubChannelsPaginator(client, input)
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

// Lists the tags applied to an Amazon Chime SDK messaging resource.
func chimesdkmessaging_ListTagsForResource(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_chimesdkmessagingResourceARN) > 0 {
		input.ResourceARN = aws.String(_chimesdkmessagingResourceARN)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the number of days before the channel is automatically deleted.
// - A background process deletes expired channels within 6 hours of expiration.
// Actual deletion times may vary.
//
// - Expired channels that have not yet been deleted appear as active, and you
// can update their expiration settings. The system honors the new settings.
//
// - The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
func chimesdkmessaging_PutChannelExpirationSettings(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.PutChannelExpirationSettingsInput{
		// ChannelArn: *string, // Required
	}

	if len(_chimesdkmessagingChannelArn) > 0 {
		input.ChannelArn = aws.String(_chimesdkmessagingChannelArn)
	}
	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}
	if len(_chimesdkmessagingExpirationSettings) > 0 {
		if err := assignInputField(input, "ExpirationSettings", _chimesdkmessagingExpirationSettings); err != nil {
			log.Errorf("invalid --expiration-settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutChannelExpirationSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the membership preferences of an AppInstanceUser or AppInstanceBot for the
// specified channel. The user or bot must be a member of the channel. Only the
// user or bot who owns the membership can set preferences. Users or bots in the
// AppInstanceAdmin and channel moderator roles can't set preferences for other
// users. Banned users or bots can't set membership preferences for the channel
// from which they are banned.
//
// The x-amz-chime-bearer request header is mandatory. Use the ARN of an
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
func chimesdkmessaging_PutChannelMembershipPreferences(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.PutChannelMembershipPreferencesInput{
		// ChannelArn: *string, // Required
		// ChimeBearer: *string, // Required
		// MemberArn: *string, // Required
		// Preferences: *types.ChannelMembershipPreferences, // Required
	}

	if len(_chimesdkmessagingChannelArn) > 0 {
		input.ChannelArn = aws.String(_chimesdkmessagingChannelArn)
	}
	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}
	if len(_chimesdkmessagingMemberArn) > 0 {
		input.MemberArn = aws.String(_chimesdkmessagingMemberArn)
	}
	if len(_chimesdkmessagingPreferences) > 0 {
		if err := assignInputField(input, "Preferences", _chimesdkmessagingPreferences); err != nil {
			log.Errorf("invalid --preferences: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutChannelMembershipPreferences(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the data streaming configuration for an AppInstance . For more information,
// see [Streaming messaging data]in the Amazon Chime SDK Developer Guide.
//
// [Streaming messaging data]: https://docs.aws.amazon.com/chime-sdk/latest/dg/streaming-export.html
func chimesdkmessaging_PutMessagingStreamingConfigurations(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.PutMessagingStreamingConfigurationsInput{
		// AppInstanceArn: *string, // Required
		// StreamingConfigurations: []types.StreamingConfiguration, // Required
	}

	if len(_chimesdkmessagingAppInstanceArn) > 0 {
		input.AppInstanceArn = aws.String(_chimesdkmessagingAppInstanceArn)
	}
	if len(_chimesdkmessagingStreamingConfigurations) > 0 {
		if err := assignInputField(input, "StreamingConfigurations", _chimesdkmessagingStreamingConfigurations); err != nil {
			log.Errorf("invalid --streaming-configurations: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutMessagingStreamingConfigurations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Redacts message content and metadata. The message exists in the back end, but
// the action returns null content, and the state shows as redacted.
//
// The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
func chimesdkmessaging_RedactChannelMessage(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.RedactChannelMessageInput{
		// ChannelArn: *string, // Required
		// ChimeBearer: *string, // Required
		// MessageId: *string, // Required
	}

	if len(_chimesdkmessagingChannelArn) > 0 {
		input.ChannelArn = aws.String(_chimesdkmessagingChannelArn)
	}
	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}
	if len(_chimesdkmessagingMessageId) > 0 {
		input.MessageId = aws.String(_chimesdkmessagingMessageId)
	}
	if len(_chimesdkmessagingSubChannelId) > 0 {
		input.SubChannelId = aws.String(_chimesdkmessagingSubChannelId)
	}

	if resp, err := client.RedactChannelMessage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows the ChimeBearer to search channels by channel members. Users or bots can
// search across the channels that they belong to. Users in the AppInstanceAdmin
// role can search across all channels.
//
// The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
//
// This operation isn't supported for AppInstanceUsers with a large number of
// memberships.
func chimesdkmessaging_SearchChannels(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.SearchChannelsInput{
		// Fields: []types.SearchField, // Required
	}

	if len(_chimesdkmessagingFields) > 0 {
		if err := assignInputField(input, "Fields", _chimesdkmessagingFields); err != nil {
			log.Errorf("invalid --fields: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}
	if len(_chimesdkmessagingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimesdkmessagingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmessagingNextToken) > 0 {
		input.NextToken = aws.String(_chimesdkmessagingNextToken)
	}

	if disablePaginator() {
		if resp, err := client.SearchChannels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chimesdkmessaging.SearchChannelsOutput
	p := chimesdkmessaging.NewSearchChannelsPaginator(client, input)
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

// Sends a message to a particular channel that the member is a part of.
// The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
//
// Also, STANDARD messages can be up to 4KB in size and contain metadata. Metadata
// is arbitrary, and you can use it in a variety of ways, such as containing a link
// to an attachment.
//
// CONTROL messages are limited to 30 bytes and do not contain metadata.
func chimesdkmessaging_SendChannelMessage(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.SendChannelMessageInput{
		// ChannelArn: *string, // Required
		// ChimeBearer: *string, // Required
		// ClientRequestToken: *string, // Required
		// Content: *string, // Required
		// Persistence: types.ChannelMessagePersistenceType, // Required
		// Type: types.ChannelMessageType, // Required
	}

	if len(_chimesdkmessagingChannelArn) > 0 {
		input.ChannelArn = aws.String(_chimesdkmessagingChannelArn)
	}
	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}
	if len(_chimesdkmessagingClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_chimesdkmessagingClientRequestToken)
	}
	if len(_chimesdkmessagingContent) > 0 {
		input.Content = aws.String(_chimesdkmessagingContent)
	}
	if len(_chimesdkmessagingPersistence) > 0 {
		if err := assignInputField(input, "Persistence", _chimesdkmessagingPersistence); err != nil {
			log.Errorf("invalid --persistence: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmessagingType) > 0 {
		if err := assignInputField(input, "Type", _chimesdkmessagingType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmessagingContentType) > 0 {
		input.ContentType = aws.String(_chimesdkmessagingContentType)
	}
	if len(_chimesdkmessagingMessageAttributes) > 0 {
		if err := assignInputField(input, "MessageAttributes", _chimesdkmessagingMessageAttributes); err != nil {
			log.Errorf("invalid --message-attributes: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmessagingMetadata) > 0 {
		input.Metadata = aws.String(_chimesdkmessagingMetadata)
	}
	if len(_chimesdkmessagingPushNotification) > 0 {
		if err := assignInputField(input, "PushNotification", _chimesdkmessagingPushNotification); err != nil {
			log.Errorf("invalid --push-notification: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmessagingSubChannelId) > 0 {
		input.SubChannelId = aws.String(_chimesdkmessagingSubChannelId)
	}
	if len(_chimesdkmessagingTarget) > 0 {
		if err := assignInputField(input, "Target", _chimesdkmessagingTarget); err != nil {
			log.Errorf("invalid --target: %s", err.Error())
			return
		}
	}

	if resp, err := client.SendChannelMessage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Applies the specified tags to the specified Amazon Chime SDK messaging resource.
func chimesdkmessaging_TagResource(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_chimesdkmessagingResourceARN) > 0 {
		input.ResourceARN = aws.String(_chimesdkmessagingResourceARN)
	}
	if len(_chimesdkmessagingTags) > 0 {
		if err := assignInputField(input, "Tags", _chimesdkmessagingTags); err != nil {
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

// Removes the specified tags from the specified Amazon Chime SDK messaging
// resource.
func chimesdkmessaging_UntagResource(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_chimesdkmessagingResourceARN) > 0 {
		input.ResourceARN = aws.String(_chimesdkmessagingResourceARN)
	}
	if len(_chimesdkmessagingTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _chimesdkmessagingTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a channel's attributes.
// Restriction: You can't change a channel's privacy.
//
// The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
func chimesdkmessaging_UpdateChannel(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.UpdateChannelInput{
		// ChannelArn: *string, // Required
		// ChimeBearer: *string, // Required
	}

	if len(_chimesdkmessagingChannelArn) > 0 {
		input.ChannelArn = aws.String(_chimesdkmessagingChannelArn)
	}
	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}
	if len(_chimesdkmessagingMetadata) > 0 {
		input.Metadata = aws.String(_chimesdkmessagingMetadata)
	}
	if len(_chimesdkmessagingMode) > 0 {
		if err := assignInputField(input, "Mode", _chimesdkmessagingMode); err != nil {
			log.Errorf("invalid --mode: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmessagingName) > 0 {
		input.Name = aws.String(_chimesdkmessagingName)
	}

	if resp, err := client.UpdateChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates channel flow attributes. This is a developer API.
func chimesdkmessaging_UpdateChannelFlow(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.UpdateChannelFlowInput{
		// ChannelFlowArn: *string, // Required
		// Name: *string, // Required
		// Processors: []types.Processor, // Required
	}

	if len(_chimesdkmessagingChannelFlowArn) > 0 {
		input.ChannelFlowArn = aws.String(_chimesdkmessagingChannelFlowArn)
	}
	if len(_chimesdkmessagingName) > 0 {
		input.Name = aws.String(_chimesdkmessagingName)
	}
	if len(_chimesdkmessagingProcessors) > 0 {
		if err := assignInputField(input, "Processors", _chimesdkmessagingProcessors); err != nil {
			log.Errorf("invalid --processors: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateChannelFlow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the content of a message.
// The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
func chimesdkmessaging_UpdateChannelMessage(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.UpdateChannelMessageInput{
		// ChannelArn: *string, // Required
		// ChimeBearer: *string, // Required
		// Content: *string, // Required
		// MessageId: *string, // Required
	}

	if len(_chimesdkmessagingChannelArn) > 0 {
		input.ChannelArn = aws.String(_chimesdkmessagingChannelArn)
	}
	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}
	if len(_chimesdkmessagingContent) > 0 {
		input.Content = aws.String(_chimesdkmessagingContent)
	}
	if len(_chimesdkmessagingMessageId) > 0 {
		input.MessageId = aws.String(_chimesdkmessagingMessageId)
	}
	if len(_chimesdkmessagingContentType) > 0 {
		input.ContentType = aws.String(_chimesdkmessagingContentType)
	}
	if len(_chimesdkmessagingMetadata) > 0 {
		input.Metadata = aws.String(_chimesdkmessagingMetadata)
	}
	if len(_chimesdkmessagingSubChannelId) > 0 {
		input.SubChannelId = aws.String(_chimesdkmessagingSubChannelId)
	}

	if resp, err := client.UpdateChannelMessage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The details of the time when a user last read messages in a channel.
// The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
func chimesdkmessaging_UpdateChannelReadMarker(cfg aws.Config, client *chimesdkmessaging.Client) {
	input := &chimesdkmessaging.UpdateChannelReadMarkerInput{
		// ChannelArn: *string, // Required
		// ChimeBearer: *string, // Required
	}

	if len(_chimesdkmessagingChannelArn) > 0 {
		input.ChannelArn = aws.String(_chimesdkmessagingChannelArn)
	}
	if len(_chimesdkmessagingChimeBearer) > 0 {
		input.ChimeBearer = aws.String(_chimesdkmessagingChimeBearer)
	}

	if resp, err := client.UpdateChannelReadMarker(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_chimesdkmessagingCmd)
	_chimesdkmessagingCmd.Flags().SortFlags = false

	_chimesdkmessagingCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_chimesdkmessagingCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_chimesdkmessagingCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingAppInstanceArn, "app-instance-arn", "", "", "App Instance ARN")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingAppInstanceUserArn, "app-instance-user-arn", "", "", "App Instance User ARN")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingCallbackId, "callback-id", "", "", "Callback ID")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingChannelArn, "channel-arn", "", "", "Channel ARN")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingChannelFlowArn, "channel-flow-arn", "", "", "Channel Flow ARN")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingChannelId, "channel-id", "", "", "Channel ID")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingChannelMessage, "channel-message", "", "", "Channel Message")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingChannelModeratorArn, "channel-moderator-arn", "", "", "Channel Moderator ARN")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingChimeBearer, "chime-bearer", "", "", "Chime Bearer")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingClientRequestToken, "client-request-token", "", "", "Client Request Token")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingContent, "content", "", "", "Content")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingContentType, "content-type", "", "", "Content Type")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingDeleteResource, "delete-resource", "", "", "Delete Resource")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingElasticChannelConfiguration, "elastic-channel-configuration", "", "", "Elastic Channel Configuration")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingExpirationSettings, "expiration-settings", "", "", "Expiration Settings")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingFields, "fields", "", "", "Fields")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingMaxResults, "max-results", "", "", "Max Results")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingMemberArn, "member-arn", "", "", "Member ARN")
	_chimesdkmessagingCmd.Flags().StringSliceVarP(&_chimesdkmessagingMemberArns, "member-arns", "", nil, "Member Arns")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingMessageAttributes, "message-attributes", "", "", "Message Attributes")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingMessageId, "message-id", "", "", "Message ID")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingMetadata, "metadata", "", "", "Metadata")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingMode, "mode", "", "", "Mode")
	_chimesdkmessagingCmd.Flags().StringSliceVarP(&_chimesdkmessagingModeratorArns, "moderator-arns", "", nil, "Moderator Arns")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingName, "name", "", "", "Name")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingNetworkType, "network-type", "", "", "Network Type")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingNextToken, "next-token", "", "", "Next Token")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingNotAfter, "not-after", "", "", "Not After")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingNotBefore, "not-before", "", "", "Not Before")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingPersistence, "persistence", "", "", "Persistence")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingPreferences, "preferences", "", "", "Preferences")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingPrivacy, "privacy", "", "", "Privacy")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingProcessors, "processors", "", "", "Processors")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingPushNotification, "push-notification", "", "", "Push Notification")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingResourceARN, "resource-arn", "", "", "Resource ARN")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingSortOrder, "sort-order", "", "", "Sort Order")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingStreamingConfigurations, "streaming-configurations", "", "", "Streaming Configurations")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingSubChannelId, "sub-channel-id", "", "", "Sub Channel ID")
	_chimesdkmessagingCmd.Flags().StringSliceVarP(&_chimesdkmessagingTagKeys, "tag-keys", "", nil, "Tag Keys")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingTags, "tags", "", "", "Tags")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingTarget, "target", "", "", "Target")
	_chimesdkmessagingCmd.Flags().StringVarP(&_chimesdkmessagingType, "type", "", "", "Type")

	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingAssociateChannelFlow, "associate-channel-flow", "", false, "Associate Channel Flow")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingBatchCreateChannelMembership, "batch-create-channel-membership", "", false, "Batch Create Channel Membership")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingChannelFlowCallback, "channel-flow-callback", "", false, "Channel Flow Callback")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingCreateChannel, "create-channel", "", false, "Create Channel")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingCreateChannelBan, "create-channel-ban", "", false, "Create Channel Ban")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingCreateChannelFlow, "create-channel-flow", "", false, "Create Channel Flow")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingCreateChannelMembership, "create-channel-membership", "", false, "Create Channel Membership")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingCreateChannelModerator, "create-channel-moderator", "", false, "Create Channel Moderator")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingDeleteChannel, "delete-channel", "", false, "Delete Channel")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingDeleteChannelBan, "delete-channel-ban", "", false, "Delete Channel Ban")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingDeleteChannelFlow, "delete-channel-flow", "", false, "Delete Channel Flow")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingDeleteChannelMembership, "delete-channel-membership", "", false, "Delete Channel Membership")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingDeleteChannelMessage, "delete-channel-message", "", false, "Delete Channel Message")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingDeleteChannelModerator, "delete-channel-moderator", "", false, "Delete Channel Moderator")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingDeleteMessagingStreamingConfigurations, "delete-messaging-streaming-configurations", "", false, "Delete Messaging Streaming Configurations")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingDescribeChannel, "describe-channel", "", false, "Describe Channel")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingDescribeChannelBan, "describe-channel-ban", "", false, "Describe Channel Ban")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingDescribeChannelFlow, "describe-channel-flow", "", false, "Describe Channel Flow")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingDescribeChannelMembership, "describe-channel-membership", "", false, "Describe Channel Membership")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingDescribeChannelMembershipForAppInstanceUser, "describe-channel-membership-for-app-instance-user", "", false, "Describe Channel Membership For App Instance User")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingDescribeChannelModeratedByAppInstanceUser, "describe-channel-moderated-by-app-instance-user", "", false, "Describe Channel Moderated By App Instance User")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingDescribeChannelModerator, "describe-channel-moderator", "", false, "Describe Channel Moderator")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingDisassociateChannelFlow, "disassociate-channel-flow", "", false, "Disassociate Channel Flow")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingGetChannelMembershipPreferences, "get-channel-membership-preferences", "", false, "Get Channel Membership Preferences")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingGetChannelMessage, "get-channel-message", "", false, "Get Channel Message")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingGetChannelMessageStatus, "get-channel-message-status", "", false, "Get Channel Message Status")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingGetMessagingSessionEndpoint, "get-messaging-session-endpoint", "", false, "Get Messaging Session Endpoint")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingGetMessagingStreamingConfigurations, "get-messaging-streaming-configurations", "", false, "Get Messaging Streaming Configurations")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingListChannelBans, "list-channel-bans", "", false, "List Channel Bans")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingListChannelFlows, "list-channel-flows", "", false, "List Channel Flows")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingListChannelMemberships, "list-channel-memberships", "", false, "List Channel Memberships")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingListChannelMembershipsForAppInstanceUser, "list-channel-memberships-for-app-instance-user", "", false, "List Channel Memberships For App Instance User")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingListChannelMessages, "list-channel-messages", "", false, "List Channel Messages")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingListChannelModerators, "list-channel-moderators", "", false, "List Channel Moderators")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingListChannels, "list-channels", "", false, "List Channels")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingListChannelsAssociatedWithChannelFlow, "list-channels-associated-with-channel-flow", "", false, "List Channels Associated With Channel Flow")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingListChannelsModeratedByAppInstanceUser, "list-channels-moderated-by-app-instance-user", "", false, "List Channels Moderated By App Instance User")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingListSubChannels, "list-sub-channels", "", false, "List Sub Channels")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingPutChannelExpirationSettings, "put-channel-expiration-settings", "", false, "Put Channel Expiration Settings")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingPutChannelMembershipPreferences, "put-channel-membership-preferences", "", false, "Put Channel Membership Preferences")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingPutMessagingStreamingConfigurations, "put-messaging-streaming-configurations", "", false, "Put Messaging Streaming Configurations")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingRedactChannelMessage, "redact-channel-message", "", false, "Redact Channel Message")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingSearchChannels, "search-channels", "", false, "Search Channels")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingSendChannelMessage, "send-channel-message", "", false, "Send Channel Message")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingTagResource, "tag-resource", "", false, "Tag Resource")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingUntagResource, "untag-resource", "", false, "Untag Resource")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingUpdateChannel, "update-channel", "", false, "Update Channel")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingUpdateChannelFlow, "update-channel-flow", "", false, "Update Channel Flow")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingUpdateChannelMessage, "update-channel-message", "", false, "Update Channel Message")
	_chimesdkmessagingCmd.Flags().BoolVarP(&_chimesdkmessagingUpdateChannelReadMarker, "update-channel-read-marker", "", false, "Update Channel Read Marker")

}
