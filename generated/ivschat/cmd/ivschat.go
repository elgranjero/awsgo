package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ivschat"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ivschatCmd represents the ivschat command
var _ivschatCmd = &cobra.Command{
	Use:   "ivschat",
	Short: "AWS ivschat CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := ivschat.NewFromConfig(cfg)
		if _ivschatCreateChatToken {
			ivschat_CreateChatToken(cfg, client)
			return
		}
		if _ivschatCreateLoggingConfiguration {
			ivschat_CreateLoggingConfiguration(cfg, client)
			return
		}
		if _ivschatCreateRoom {
			ivschat_CreateRoom(cfg, client)
			return
		}
		if _ivschatDeleteLoggingConfiguration {
			ivschat_DeleteLoggingConfiguration(cfg, client)
			return
		}
		if _ivschatDeleteMessage {
			ivschat_DeleteMessage(cfg, client)
			return
		}
		if _ivschatDeleteRoom {
			ivschat_DeleteRoom(cfg, client)
			return
		}
		if _ivschatDisconnectUser {
			ivschat_DisconnectUser(cfg, client)
			return
		}
		if _ivschatGetLoggingConfiguration {
			ivschat_GetLoggingConfiguration(cfg, client)
			return
		}
		if _ivschatGetRoom {
			ivschat_GetRoom(cfg, client)
			return
		}
		if _ivschatListLoggingConfigurations {
			ivschat_ListLoggingConfigurations(cfg, client)
			return
		}
		if _ivschatListRooms {
			ivschat_ListRooms(cfg, client)
			return
		}
		if _ivschatListTagsForResource {
			ivschat_ListTagsForResource(cfg, client)
			return
		}
		if _ivschatSendEvent {
			ivschat_SendEvent(cfg, client)
			return
		}
		if _ivschatTagResource {
			ivschat_TagResource(cfg, client)
			return
		}
		if _ivschatUntagResource {
			ivschat_UntagResource(cfg, client)
			return
		}
		if _ivschatUpdateLoggingConfiguration {
			ivschat_UpdateLoggingConfiguration(cfg, client)
			return
		}
		if _ivschatUpdateRoom {
			ivschat_UpdateRoom(cfg, client)
			return
		}

	},
}

var (
	_ivschatCreateChatToken            bool
	_ivschatCreateLoggingConfiguration bool
	_ivschatCreateRoom                 bool
	_ivschatDeleteLoggingConfiguration bool
	_ivschatDeleteMessage              bool
	_ivschatDeleteRoom                 bool
	_ivschatDisconnectUser             bool
	_ivschatGetLoggingConfiguration    bool
	_ivschatGetRoom                    bool
	_ivschatListLoggingConfigurations  bool
	_ivschatListRooms                  bool
	_ivschatListTagsForResource        bool
	_ivschatSendEvent                  bool
	_ivschatTagResource                bool
	_ivschatUntagResource              bool
	_ivschatUpdateLoggingConfiguration bool
	_ivschatUpdateRoom                 bool

	_ivschatAttributes                      string
	_ivschatCapabilities                    string
	_ivschatDestinationConfiguration        string
	_ivschatEventName                       string
	_ivschatId                              string
	_ivschatIdentifier                      string
	_ivschatLoggingConfigurationIdentifier  string
	_ivschatLoggingConfigurationIdentifiers []string
	_ivschatMaxResults                      string
	_ivschatMaximumMessageLength            string
	_ivschatMaximumMessageRatePerSecond     string
	_ivschatMessageReviewHandler            string
	_ivschatMessageReviewHandlerUri         string
	_ivschatName                            string
	_ivschatNextToken                       string
	_ivschatReason                          string
	_ivschatResourceArn                     string
	_ivschatRoomIdentifier                  string
	_ivschatSessionDurationInMinutes        string
	_ivschatTagKeys                         []string
	_ivschatTags                            string
	_ivschatUserId                          string
)

// Creates an encrypted token that is used by a chat participant to establish an
// individual WebSocket chat connection to a room. When the token is used to
// connect to chat, the connection is valid for the session duration specified in
// the request. The token becomes invalid at the token-expiration timestamp
// included in the response.
//
// Use the capabilities field to permit an end user to send messages or moderate a
// room.
//
// The attributes field securely attaches structured data to the chat session; the
// data is included within each message sent by the end user and received by other
// participants in the room. Common use cases for attributes include passing
// end-user profile data like an icon, display name, colors, badges, and other
// display features.
//
// Encryption keys are owned by Amazon IVS Chat and never used directly by your
// application.
func ivschat_CreateChatToken(cfg aws.Config, client *ivschat.Client) {
	input := &ivschat.CreateChatTokenInput{
		// RoomIdentifier: *string, // Required
		// UserId: *string, // Required
	}

	if len(_ivschatRoomIdentifier) > 0 {
		input.RoomIdentifier = aws.String(_ivschatRoomIdentifier)
	}
	if len(_ivschatUserId) > 0 {
		input.UserId = aws.String(_ivschatUserId)
	}
	if len(_ivschatAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _ivschatAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_ivschatCapabilities) > 0 {
		if err := assignInputField(input, "Capabilities", _ivschatCapabilities); err != nil {
			log.Errorf("invalid --capabilities: %s", err.Error())
			return
		}
	}
	if len(_ivschatSessionDurationInMinutes) > 0 {
		if err := assignInputField(input, "SessionDurationInMinutes", _ivschatSessionDurationInMinutes); err != nil {
			log.Errorf("invalid --session-duration-in-minutes: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateChatToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a logging configuration that allows clients to store and record sent
// messages.
func ivschat_CreateLoggingConfiguration(cfg aws.Config, client *ivschat.Client) {
	input := &ivschat.CreateLoggingConfigurationInput{
		// DestinationConfiguration: types.DestinationConfiguration, // Required
	}

	if len(_ivschatDestinationConfiguration) > 0 {
		if err := assignInputField(input, "DestinationConfiguration", _ivschatDestinationConfiguration); err != nil {
			log.Errorf("invalid --destination-configuration: %s", err.Error())
			return
		}
	}
	if len(_ivschatName) > 0 {
		input.Name = aws.String(_ivschatName)
	}
	if len(_ivschatTags) > 0 {
		if err := assignInputField(input, "Tags", _ivschatTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLoggingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a room that allows clients to connect and pass messages.
func ivschat_CreateRoom(cfg aws.Config, client *ivschat.Client) {
	input := &ivschat.CreateRoomInput{}

	if len(_ivschatLoggingConfigurationIdentifiers) > 0 {
		input.LoggingConfigurationIdentifiers = append([]string(nil), _ivschatLoggingConfigurationIdentifiers...)
	}
	if len(_ivschatMaximumMessageLength) > 0 {
		if err := assignInputField(input, "MaximumMessageLength", _ivschatMaximumMessageLength); err != nil {
			log.Errorf("invalid --maximum-message-length: %s", err.Error())
			return
		}
	}
	if len(_ivschatMaximumMessageRatePerSecond) > 0 {
		if err := assignInputField(input, "MaximumMessageRatePerSecond", _ivschatMaximumMessageRatePerSecond); err != nil {
			log.Errorf("invalid --maximum-message-rate-per-second: %s", err.Error())
			return
		}
	}
	if len(_ivschatMessageReviewHandler) > 0 {
		if err := assignInputField(input, "MessageReviewHandler", _ivschatMessageReviewHandler); err != nil {
			log.Errorf("invalid --message-review-handler: %s", err.Error())
			return
		}
	}
	if len(_ivschatName) > 0 {
		input.Name = aws.String(_ivschatName)
	}
	if len(_ivschatTags) > 0 {
		if err := assignInputField(input, "Tags", _ivschatTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRoom(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified logging configuration.
func ivschat_DeleteLoggingConfiguration(cfg aws.Config, client *ivschat.Client) {
	input := &ivschat.DeleteLoggingConfigurationInput{
		// Identifier: *string, // Required
	}

	if len(_ivschatIdentifier) > 0 {
		input.Identifier = aws.String(_ivschatIdentifier)
	}

	if resp, err := client.DeleteLoggingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends an event to a specific room which directs clients to delete a specific
// message; that is, unrender it from view and delete it from the client’s chat
// history. This event’s EventName is aws:DELETE_MESSAGE . This replicates the [DeleteMessage]
// WebSocket operation in the Amazon IVS Chat Messaging API.
//
// [DeleteMessage]: https://docs.aws.amazon.com/ivs/latest/chatmsgapireference/actions-deletemessage-publish.html
func ivschat_DeleteMessage(cfg aws.Config, client *ivschat.Client) {
	input := &ivschat.DeleteMessageInput{
		// Id: *string, // Required
		// RoomIdentifier: *string, // Required
	}

	if len(_ivschatId) > 0 {
		input.Id = aws.String(_ivschatId)
	}
	if len(_ivschatRoomIdentifier) > 0 {
		input.RoomIdentifier = aws.String(_ivschatRoomIdentifier)
	}
	if len(_ivschatReason) > 0 {
		input.Reason = aws.String(_ivschatReason)
	}

	if resp, err := client.DeleteMessage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified room.
func ivschat_DeleteRoom(cfg aws.Config, client *ivschat.Client) {
	input := &ivschat.DeleteRoomInput{
		// Identifier: *string, // Required
	}

	if len(_ivschatIdentifier) > 0 {
		input.Identifier = aws.String(_ivschatIdentifier)
	}

	if resp, err := client.DeleteRoom(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disconnects all connections using a specified user ID from a room. This
// replicates the [DisconnectUser]WebSocket operation in the Amazon IVS Chat Messaging API.
//
// [DisconnectUser]: https://docs.aws.amazon.com/ivs/latest/chatmsgapireference/actions-disconnectuser-publish.html
func ivschat_DisconnectUser(cfg aws.Config, client *ivschat.Client) {
	input := &ivschat.DisconnectUserInput{
		// RoomIdentifier: *string, // Required
		// UserId: *string, // Required
	}

	if len(_ivschatRoomIdentifier) > 0 {
		input.RoomIdentifier = aws.String(_ivschatRoomIdentifier)
	}
	if len(_ivschatUserId) > 0 {
		input.UserId = aws.String(_ivschatUserId)
	}
	if len(_ivschatReason) > 0 {
		input.Reason = aws.String(_ivschatReason)
	}

	if resp, err := client.DisconnectUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the specified logging configuration.
func ivschat_GetLoggingConfiguration(cfg aws.Config, client *ivschat.Client) {
	input := &ivschat.GetLoggingConfigurationInput{
		// Identifier: *string, // Required
	}

	if len(_ivschatIdentifier) > 0 {
		input.Identifier = aws.String(_ivschatIdentifier)
	}

	if resp, err := client.GetLoggingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the specified room.
func ivschat_GetRoom(cfg aws.Config, client *ivschat.Client) {
	input := &ivschat.GetRoomInput{
		// Identifier: *string, // Required
	}

	if len(_ivschatIdentifier) > 0 {
		input.Identifier = aws.String(_ivschatIdentifier)
	}

	if resp, err := client.GetRoom(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets summary information about all your logging configurations in the AWS
// region where the API request is processed.
func ivschat_ListLoggingConfigurations(cfg aws.Config, client *ivschat.Client) {
	input := &ivschat.ListLoggingConfigurationsInput{}

	if len(_ivschatMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ivschatMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ivschatNextToken) > 0 {
		input.NextToken = aws.String(_ivschatNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListLoggingConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ivschat.ListLoggingConfigurationsOutput
	p := ivschat.NewListLoggingConfigurationsPaginator(client, input)
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

// Gets summary information about all your rooms in the AWS region where the API
// request is processed. Results are sorted in descending order of updateTime .
func ivschat_ListRooms(cfg aws.Config, client *ivschat.Client) {
	input := &ivschat.ListRoomsInput{}

	if len(_ivschatLoggingConfigurationIdentifier) > 0 {
		input.LoggingConfigurationIdentifier = aws.String(_ivschatLoggingConfigurationIdentifier)
	}
	if len(_ivschatMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ivschatMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ivschatMessageReviewHandlerUri) > 0 {
		input.MessageReviewHandlerUri = aws.String(_ivschatMessageReviewHandlerUri)
	}
	if len(_ivschatName) > 0 {
		input.Name = aws.String(_ivschatName)
	}
	if len(_ivschatNextToken) > 0 {
		input.NextToken = aws.String(_ivschatNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRooms(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ivschat.ListRoomsOutput
	p := ivschat.NewListRoomsPaginator(client, input)
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

// Gets information about AWS tags for the specified ARN.
func ivschat_ListTagsForResource(cfg aws.Config, client *ivschat.Client) {
	input := &ivschat.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_ivschatResourceArn) > 0 {
		input.ResourceArn = aws.String(_ivschatResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends an event to a room. Use this within your application’s business logic to
// send events to clients of a room; e.g., to notify clients to change the way the
// chat UI is rendered.
func ivschat_SendEvent(cfg aws.Config, client *ivschat.Client) {
	input := &ivschat.SendEventInput{
		// EventName: *string, // Required
		// RoomIdentifier: *string, // Required
	}

	if len(_ivschatEventName) > 0 {
		input.EventName = aws.String(_ivschatEventName)
	}
	if len(_ivschatRoomIdentifier) > 0 {
		input.RoomIdentifier = aws.String(_ivschatRoomIdentifier)
	}
	if len(_ivschatAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _ivschatAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}

	if resp, err := client.SendEvent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or updates tags for the AWS resource with the specified ARN.
func ivschat_TagResource(cfg aws.Config, client *ivschat.Client) {
	input := &ivschat.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_ivschatResourceArn) > 0 {
		input.ResourceArn = aws.String(_ivschatResourceArn)
	}
	if len(_ivschatTags) > 0 {
		if err := assignInputField(input, "Tags", _ivschatTags); err != nil {
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
func ivschat_UntagResource(cfg aws.Config, client *ivschat.Client) {
	input := &ivschat.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_ivschatResourceArn) > 0 {
		input.ResourceArn = aws.String(_ivschatResourceArn)
	}
	if len(_ivschatTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _ivschatTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a specified logging configuration.
func ivschat_UpdateLoggingConfiguration(cfg aws.Config, client *ivschat.Client) {
	input := &ivschat.UpdateLoggingConfigurationInput{
		// Identifier: *string, // Required
	}

	if len(_ivschatIdentifier) > 0 {
		input.Identifier = aws.String(_ivschatIdentifier)
	}
	if len(_ivschatDestinationConfiguration) > 0 {
		if err := assignInputField(input, "DestinationConfiguration", _ivschatDestinationConfiguration); err != nil {
			log.Errorf("invalid --destination-configuration: %s", err.Error())
			return
		}
	}
	if len(_ivschatName) > 0 {
		input.Name = aws.String(_ivschatName)
	}

	if resp, err := client.UpdateLoggingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a room’s configuration.
func ivschat_UpdateRoom(cfg aws.Config, client *ivschat.Client) {
	input := &ivschat.UpdateRoomInput{
		// Identifier: *string, // Required
	}

	if len(_ivschatIdentifier) > 0 {
		input.Identifier = aws.String(_ivschatIdentifier)
	}
	if len(_ivschatLoggingConfigurationIdentifiers) > 0 {
		input.LoggingConfigurationIdentifiers = append([]string(nil), _ivschatLoggingConfigurationIdentifiers...)
	}
	if len(_ivschatMaximumMessageLength) > 0 {
		if err := assignInputField(input, "MaximumMessageLength", _ivschatMaximumMessageLength); err != nil {
			log.Errorf("invalid --maximum-message-length: %s", err.Error())
			return
		}
	}
	if len(_ivschatMaximumMessageRatePerSecond) > 0 {
		if err := assignInputField(input, "MaximumMessageRatePerSecond", _ivschatMaximumMessageRatePerSecond); err != nil {
			log.Errorf("invalid --maximum-message-rate-per-second: %s", err.Error())
			return
		}
	}
	if len(_ivschatMessageReviewHandler) > 0 {
		if err := assignInputField(input, "MessageReviewHandler", _ivschatMessageReviewHandler); err != nil {
			log.Errorf("invalid --message-review-handler: %s", err.Error())
			return
		}
	}
	if len(_ivschatName) > 0 {
		input.Name = aws.String(_ivschatName)
	}

	if resp, err := client.UpdateRoom(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_ivschatCmd)
	_ivschatCmd.Flags().SortFlags = false

	_ivschatCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_ivschatCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_ivschatCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_ivschatCmd.Flags().StringVarP(&_ivschatAttributes, "attributes", "", "", "Attributes")
	_ivschatCmd.Flags().StringVarP(&_ivschatCapabilities, "capabilities", "", "", "Capabilities")
	_ivschatCmd.Flags().StringVarP(&_ivschatDestinationConfiguration, "destination-configuration", "", "", "Destination Configuration")
	_ivschatCmd.Flags().StringVarP(&_ivschatEventName, "event-name", "", "", "Event Name")
	_ivschatCmd.Flags().StringVarP(&_ivschatId, "id", "", "", "ID")
	_ivschatCmd.Flags().StringVarP(&_ivschatIdentifier, "identifier", "", "", "Identifier")
	_ivschatCmd.Flags().StringVarP(&_ivschatLoggingConfigurationIdentifier, "logging-configuration-identifier", "", "", "Logging Configuration Identifier")
	_ivschatCmd.Flags().StringSliceVarP(&_ivschatLoggingConfigurationIdentifiers, "logging-configuration-identifiers", "", nil, "Logging Configuration Identifiers")
	_ivschatCmd.Flags().StringVarP(&_ivschatMaxResults, "max-results", "", "", "Max Results")
	_ivschatCmd.Flags().StringVarP(&_ivschatMaximumMessageLength, "maximum-message-length", "", "", "Maximum Message Length")
	_ivschatCmd.Flags().StringVarP(&_ivschatMaximumMessageRatePerSecond, "maximum-message-rate-per-second", "", "", "Maximum Message Rate Per Second")
	_ivschatCmd.Flags().StringVarP(&_ivschatMessageReviewHandler, "message-review-handler", "", "", "Message Review Handler")
	_ivschatCmd.Flags().StringVarP(&_ivschatMessageReviewHandlerUri, "message-review-handler-uri", "", "", "Message Review Handler URI")
	_ivschatCmd.Flags().StringVarP(&_ivschatName, "name", "", "", "Name")
	_ivschatCmd.Flags().StringVarP(&_ivschatNextToken, "next-token", "", "", "Next Token")
	_ivschatCmd.Flags().StringVarP(&_ivschatReason, "reason", "", "", "Reason")
	_ivschatCmd.Flags().StringVarP(&_ivschatResourceArn, "resource-arn", "", "", "Resource ARN")
	_ivschatCmd.Flags().StringVarP(&_ivschatRoomIdentifier, "room-identifier", "", "", "Room Identifier")
	_ivschatCmd.Flags().StringVarP(&_ivschatSessionDurationInMinutes, "session-duration-in-minutes", "", "", "Session Duration In Minutes")
	_ivschatCmd.Flags().StringSliceVarP(&_ivschatTagKeys, "tag-keys", "", nil, "Tag Keys")
	_ivschatCmd.Flags().StringVarP(&_ivschatTags, "tags", "", "", "Tags")
	_ivschatCmd.Flags().StringVarP(&_ivschatUserId, "user-id", "", "", "User ID")

	_ivschatCmd.Flags().BoolVarP(&_ivschatCreateChatToken, "create-chat-token", "", false, "Create Chat Token")
	_ivschatCmd.Flags().BoolVarP(&_ivschatCreateLoggingConfiguration, "create-logging-configuration", "", false, "Create Logging Configuration")
	_ivschatCmd.Flags().BoolVarP(&_ivschatCreateRoom, "create-room", "", false, "Create Room")
	_ivschatCmd.Flags().BoolVarP(&_ivschatDeleteLoggingConfiguration, "delete-logging-configuration", "", false, "Delete Logging Configuration")
	_ivschatCmd.Flags().BoolVarP(&_ivschatDeleteMessage, "delete-message", "", false, "Delete Message")
	_ivschatCmd.Flags().BoolVarP(&_ivschatDeleteRoom, "delete-room", "", false, "Delete Room")
	_ivschatCmd.Flags().BoolVarP(&_ivschatDisconnectUser, "disconnect-user", "", false, "Disconnect User")
	_ivschatCmd.Flags().BoolVarP(&_ivschatGetLoggingConfiguration, "get-logging-configuration", "", false, "Get Logging Configuration")
	_ivschatCmd.Flags().BoolVarP(&_ivschatGetRoom, "get-room", "", false, "Get Room")
	_ivschatCmd.Flags().BoolVarP(&_ivschatListLoggingConfigurations, "list-logging-configurations", "", false, "List Logging Configurations")
	_ivschatCmd.Flags().BoolVarP(&_ivschatListRooms, "list-rooms", "", false, "List Rooms")
	_ivschatCmd.Flags().BoolVarP(&_ivschatListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_ivschatCmd.Flags().BoolVarP(&_ivschatSendEvent, "send-event", "", false, "Send Event")
	_ivschatCmd.Flags().BoolVarP(&_ivschatTagResource, "tag-resource", "", false, "Tag Resource")
	_ivschatCmd.Flags().BoolVarP(&_ivschatUntagResource, "untag-resource", "", false, "Untag Resource")
	_ivschatCmd.Flags().BoolVarP(&_ivschatUpdateLoggingConfiguration, "update-logging-configuration", "", false, "Update Logging Configuration")
	_ivschatCmd.Flags().BoolVarP(&_ivschatUpdateRoom, "update-room", "", false, "Update Room")

}
