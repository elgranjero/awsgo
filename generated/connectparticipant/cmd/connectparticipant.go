package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/connectparticipant"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// connectparticipantCmd represents the connectparticipant command
var _connectparticipantCmd = &cobra.Command{
	Use:   "connectparticipant",
	Short: "AWS connectparticipant CLI",
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
		client := connectparticipant.NewFromConfig(cfg)
		if _connectparticipantCancelParticipantAuthentication {
			connectparticipant_CancelParticipantAuthentication(cfg, client)
			return
		}
		if _connectparticipantCompleteAttachmentUpload {
			connectparticipant_CompleteAttachmentUpload(cfg, client)
			return
		}
		if _connectparticipantCreateParticipantConnection {
			connectparticipant_CreateParticipantConnection(cfg, client)
			return
		}
		if _connectparticipantDescribeView {
			connectparticipant_DescribeView(cfg, client)
			return
		}
		if _connectparticipantDisconnectParticipant {
			connectparticipant_DisconnectParticipant(cfg, client)
			return
		}
		if _connectparticipantGetAttachment {
			connectparticipant_GetAttachment(cfg, client)
			return
		}
		if _connectparticipantGetAuthenticationUrl {
			connectparticipant_GetAuthenticationUrl(cfg, client)
			return
		}
		if _connectparticipantGetTranscript {
			connectparticipant_GetTranscript(cfg, client)
			return
		}
		if _connectparticipantSendEvent {
			connectparticipant_SendEvent(cfg, client)
			return
		}
		if _connectparticipantSendMessage {
			connectparticipant_SendMessage(cfg, client)
			return
		}
		if _connectparticipantStartAttachmentUpload {
			connectparticipant_StartAttachmentUpload(cfg, client)
			return
		}

	},
}

var (
	_connectparticipantCancelParticipantAuthentication bool
	_connectparticipantCompleteAttachmentUpload        bool
	_connectparticipantCreateParticipantConnection     bool
	_connectparticipantDescribeView                    bool
	_connectparticipantDisconnectParticipant           bool
	_connectparticipantGetAttachment                   bool
	_connectparticipantGetAuthenticationUrl            bool
	_connectparticipantGetTranscript                   bool
	_connectparticipantSendEvent                       bool
	_connectparticipantSendMessage                     bool
	_connectparticipantStartAttachmentUpload           bool

	_connectparticipantAttachmentId          string
	_connectparticipantAttachmentIds         []string
	_connectparticipantAttachmentName        string
	_connectparticipantAttachmentSizeInBytes string
	_connectparticipantClientToken           string
	_connectparticipantConnectParticipant    string
	_connectparticipantConnectionToken       string
	_connectparticipantContactId             string
	_connectparticipantContent               string
	_connectparticipantContentType           string
	_connectparticipantMaxResults            string
	_connectparticipantNextToken             string
	_connectparticipantParticipantToken      string
	_connectparticipantRedirectUri           string
	_connectparticipantScanDirection         string
	_connectparticipantSessionId             string
	_connectparticipantSortOrder             string
	_connectparticipantStartPosition         string
	_connectparticipantType                  string
	_connectparticipantUrlExpiryInSeconds    string
	_connectparticipantViewToken             string
)

// Cancels the authentication session. The opted out branch of the Authenticate
// Customer flow block will be taken.
//
// The current supported channel is chat. This API is not supported for Apple
// Messages for Business, WhatsApp, or SMS chats.
//
// ConnectionToken is used for invoking this API instead of ParticipantToken .
//
// The Amazon Connect Participant Service APIs do not use [Signature Version 4 authentication].
//
// [Signature Version 4 authentication]: https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html
func connectparticipant_CancelParticipantAuthentication(cfg aws.Config, client *connectparticipant.Client) {
	input := &connectparticipant.CancelParticipantAuthenticationInput{
		// ConnectionToken: *string, // Required
		// SessionId: *string, // Required
	}

	if len(_connectparticipantConnectionToken) > 0 {
		input.ConnectionToken = aws.String(_connectparticipantConnectionToken)
	}
	if len(_connectparticipantSessionId) > 0 {
		input.SessionId = aws.String(_connectparticipantSessionId)
	}

	if resp, err := client.CancelParticipantAuthentication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows you to confirm that the attachment has been uploaded using the
// pre-signed URL provided in StartAttachmentUpload API. A conflict exception is
// thrown when an attachment with that identifier is already being uploaded.
//
// For security recommendations, see [Amazon Connect Chat security best practices].
//
// ConnectionToken is used for invoking this API instead of ParticipantToken .
//
// The Amazon Connect Participant Service APIs do not use [Signature Version 4 authentication].
//
// [Signature Version 4 authentication]: https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html
// [Amazon Connect Chat security best practices]: https://docs.aws.amazon.com/connect/latest/adminguide/security-best-practices.html#bp-security-chat
func connectparticipant_CompleteAttachmentUpload(cfg aws.Config, client *connectparticipant.Client) {
	input := &connectparticipant.CompleteAttachmentUploadInput{
		// AttachmentIds: []string, // Required
		// ClientToken: *string, // Required
		// ConnectionToken: *string, // Required
	}

	if len(_connectparticipantAttachmentIds) > 0 {
		input.AttachmentIds = append([]string(nil), _connectparticipantAttachmentIds...)
	}
	if len(_connectparticipantClientToken) > 0 {
		input.ClientToken = aws.String(_connectparticipantClientToken)
	}
	if len(_connectparticipantConnectionToken) > 0 {
		input.ConnectionToken = aws.String(_connectparticipantConnectionToken)
	}

	if resp, err := client.CompleteAttachmentUpload(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates the participant's connection.
// For security recommendations, see [Amazon Connect Chat security best practices].
//
// For WebRTC security recommendations, see [Amazon Connect WebRTC security best practices].
//
// ParticipantToken is used for invoking this API instead of ConnectionToken .
//
// The participant token is valid for the lifetime of the participant – until they
// are part of a contact. For WebRTC participants, if they leave or are
// disconnected for 60 seconds, a new participant needs to be created using the [CreateParticipant]
// API.
//
// For WEBSOCKET Type:
//
// The response URL for has a connect expiry timeout of 100s. Clients must
// manually connect to the returned websocket URL and subscribe to the desired
// topic.
//
// For chat, you need to publish the following on the established websocket
// connection:
//
// {"topic":"aws/subscribe","content":{"topics":["aws/chat"]}}
//
// Upon websocket URL expiry, as specified in the response ConnectionExpiry
// parameter, clients need to call this API again to obtain a new websocket URL and
// perform the same steps as before.
//
// The expiry time for the connection token is different than the
// ChatDurationInMinutes . Expiry time for the connection token is 1 day.
//
// For WEBRTC_CONNECTION Type:
//
// The response includes connection data required for the client application to
// join the call using the Amazon Chime SDK client libraries. The WebRTCConnection
// response contains Meeting and Attendee information needed to establish the media
// connection.
//
// The attendee join token in WebRTCConnection response is valid for the lifetime
// of the participant in the call. If a participant leaves or is disconnected for
// 60 seconds, their participant credentials will no longer be valid, and a new
// participant will need to be created to rejoin the call.
//
// Message streaming support: This API can also be used together with the [StartContactStreaming] API to
// create a participant connection for chat contacts that are not using a
// websocket. For more information about message streaming, [Enable real-time chat message streaming]in the Amazon Connect
// Administrator Guide.
//
// Multi-user web, in-app, video calling support:
//
// For WebRTC calls, this API is used in conjunction with the CreateParticipant
// API to enable multi-party calling. The StartWebRTCContact API creates the
// initial contact and routes it to an agent, while CreateParticipant adds
// additional participants to the ongoing call. For more information about
// multi-party WebRTC calls, see [Enable multi-user web, in-app, and video calling]in the Amazon Connect Administrator Guide.
//
// Feature specifications: For information about feature specifications, such as
// the allowed number of open websocket connections per participant or maximum
// number of WebRTC participants, see [Feature specifications]in the Amazon Connect Administrator Guide.
//
// The Amazon Connect Participant Service APIs do not use [Signature Version 4 authentication].
//
// [Feature specifications]: https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-service-limits.html#feature-limits
// [StartContactStreaming]: https://docs.aws.amazon.com/connect/latest/APIReference/API_StartContactStreaming.html
// [CreateParticipant]: https://docs.aws.amazon.com/connect/latest/APIReference/API_CreateParticipant.html
// [Amazon Connect WebRTC security best practices]: https://docs.aws.amazon.com/connect/latest/adminguide/security-best-practices.html#bp-webrtc-security
// [Enable real-time chat message streaming]: https://docs.aws.amazon.com/connect/latest/adminguide/chat-message-streaming.html
// [Signature Version 4 authentication]: https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html
// [Enable multi-user web, in-app, and video calling]: https://docs.aws.amazon.com/connect/latest/adminguide/enable-multiuser-inapp.html
// [Amazon Connect Chat security best practices]: https://docs.aws.amazon.com/connect/latest/adminguide/security-best-practices.html#bp-security-chat
func connectparticipant_CreateParticipantConnection(cfg aws.Config, client *connectparticipant.Client) {
	input := &connectparticipant.CreateParticipantConnectionInput{
		// ParticipantToken: *string, // Required
	}

	if len(_connectparticipantParticipantToken) > 0 {
		input.ParticipantToken = aws.String(_connectparticipantParticipantToken)
	}
	if len(_connectparticipantConnectParticipant) > 0 {
		if err := assignInputField(input, "ConnectParticipant", _connectparticipantConnectParticipant); err != nil {
			log.Errorf("invalid --connect-participant: %s", err.Error())
			return
		}
	}
	if len(_connectparticipantType) > 0 {
		if err := assignInputField(input, "Type", _connectparticipantType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateParticipantConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the view for the specified view token.
// For security recommendations, see [Amazon Connect Chat security best practices].
//
// [Amazon Connect Chat security best practices]: https://docs.aws.amazon.com/connect/latest/adminguide/security-best-practices.html#bp-security-chat
func connectparticipant_DescribeView(cfg aws.Config, client *connectparticipant.Client) {
	input := &connectparticipant.DescribeViewInput{
		// ConnectionToken: *string, // Required
		// ViewToken: *string, // Required
	}

	if len(_connectparticipantConnectionToken) > 0 {
		input.ConnectionToken = aws.String(_connectparticipantConnectionToken)
	}
	if len(_connectparticipantViewToken) > 0 {
		input.ViewToken = aws.String(_connectparticipantViewToken)
	}

	if resp, err := client.DescribeView(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disconnects a participant.
// For security recommendations, see [Amazon Connect Chat security best practices].
//
// ConnectionToken is used for invoking this API instead of ParticipantToken .
//
// The Amazon Connect Participant Service APIs do not use [Signature Version 4 authentication].
//
// [Signature Version 4 authentication]: https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html
// [Amazon Connect Chat security best practices]: https://docs.aws.amazon.com/connect/latest/adminguide/security-best-practices.html#bp-security-chat
func connectparticipant_DisconnectParticipant(cfg aws.Config, client *connectparticipant.Client) {
	input := &connectparticipant.DisconnectParticipantInput{
		// ConnectionToken: *string, // Required
	}

	if len(_connectparticipantConnectionToken) > 0 {
		input.ConnectionToken = aws.String(_connectparticipantConnectionToken)
	}
	if len(_connectparticipantClientToken) > 0 {
		input.ClientToken = aws.String(_connectparticipantClientToken)
	}

	if resp, err := client.DisconnectParticipant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides a pre-signed URL for download of a completed attachment. This is an
// asynchronous API for use with active contacts.
//
// For security recommendations, see [Amazon Connect Chat security best practices].
//
// - The participant role CUSTOM_BOT is not permitted to access attachments
// customers may upload. An AccessDeniedException can indicate that the
// participant may be a CUSTOM_BOT, and it doesn't have access to attachments.
//
// - ConnectionToken is used for invoking this API instead of ParticipantToken .
//
// The Amazon Connect Participant Service APIs do not use [Signature Version 4 authentication].
//
// [Signature Version 4 authentication]: https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html
// [Amazon Connect Chat security best practices]: https://docs.aws.amazon.com/connect/latest/adminguide/security-best-practices.html#bp-security-chat
func connectparticipant_GetAttachment(cfg aws.Config, client *connectparticipant.Client) {
	input := &connectparticipant.GetAttachmentInput{
		// AttachmentId: *string, // Required
		// ConnectionToken: *string, // Required
	}

	if len(_connectparticipantAttachmentId) > 0 {
		input.AttachmentId = aws.String(_connectparticipantAttachmentId)
	}
	if len(_connectparticipantConnectionToken) > 0 {
		input.ConnectionToken = aws.String(_connectparticipantConnectionToken)
	}
	if len(_connectparticipantUrlExpiryInSeconds) > 0 {
		if err := assignInputField(input, "UrlExpiryInSeconds", _connectparticipantUrlExpiryInSeconds); err != nil {
			log.Errorf("invalid --url-expiry-in-seconds: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetAttachment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the AuthenticationUrl for the current authentication session for the
// AuthenticateCustomer flow block.
//
// For security recommendations, see [Amazon Connect Chat security best practices].
//
// - This API can only be called within one minute of receiving the
// authenticationInitiated event.
//
// - The current supported channel is chat. This API is not supported for Apple
// Messages for Business, WhatsApp, or SMS chats.
//
// ConnectionToken is used for invoking this API instead of ParticipantToken .
//
// The Amazon Connect Participant Service APIs do not use [Signature Version 4 authentication].
//
// [Signature Version 4 authentication]: https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html
// [Amazon Connect Chat security best practices]: https://docs.aws.amazon.com/connect/latest/adminguide/security-best-practices.html#bp-security-chat
func connectparticipant_GetAuthenticationUrl(cfg aws.Config, client *connectparticipant.Client) {
	input := &connectparticipant.GetAuthenticationUrlInput{
		// ConnectionToken: *string, // Required
		// RedirectUri: *string, // Required
		// SessionId: *string, // Required
	}

	if len(_connectparticipantConnectionToken) > 0 {
		input.ConnectionToken = aws.String(_connectparticipantConnectionToken)
	}
	if len(_connectparticipantRedirectUri) > 0 {
		input.RedirectUri = aws.String(_connectparticipantRedirectUri)
	}
	if len(_connectparticipantSessionId) > 0 {
		input.SessionId = aws.String(_connectparticipantSessionId)
	}

	if resp, err := client.GetAuthenticationUrl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a transcript of the session, including details about any attachments.
// For information about accessing past chat contact transcripts for a persistent
// chat, see [Enable persistent chat].
//
// For security recommendations, see [Amazon Connect Chat security best practices].
//
// If you have a process that consumes events in the transcript of an chat that
// has ended, note that chat transcripts contain the following event content types
// if the event has occurred during the chat session:
//
// - application/vnd.amazonaws.connect.event.participant.invited
//
// - application/vnd.amazonaws.connect.event.participant.joined
//
// - application/vnd.amazonaws.connect.event.participant.left
//
// - application/vnd.amazonaws.connect.event.chat.ended
//
// - application/vnd.amazonaws.connect.event.transfer.succeeded
//
// - application/vnd.amazonaws.connect.event.transfer.failed
//
// ConnectionToken is used for invoking this API instead of ParticipantToken .
//
// The Amazon Connect Participant Service APIs do not use [Signature Version 4 authentication].
//
// [Enable persistent chat]: https://docs.aws.amazon.com/connect/latest/adminguide/chat-persistence.html
// [Signature Version 4 authentication]: https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html
// [Amazon Connect Chat security best practices]: https://docs.aws.amazon.com/connect/latest/adminguide/security-best-practices.html#bp-security-chat
func connectparticipant_GetTranscript(cfg aws.Config, client *connectparticipant.Client) {
	input := &connectparticipant.GetTranscriptInput{
		// ConnectionToken: *string, // Required
	}

	if len(_connectparticipantConnectionToken) > 0 {
		input.ConnectionToken = aws.String(_connectparticipantConnectionToken)
	}
	if len(_connectparticipantContactId) > 0 {
		input.ContactId = aws.String(_connectparticipantContactId)
	}
	if len(_connectparticipantMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _connectparticipantMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_connectparticipantNextToken) > 0 {
		input.NextToken = aws.String(_connectparticipantNextToken)
	}
	if len(_connectparticipantScanDirection) > 0 {
		if err := assignInputField(input, "ScanDirection", _connectparticipantScanDirection); err != nil {
			log.Errorf("invalid --scan-direction: %s", err.Error())
			return
		}
	}
	if len(_connectparticipantSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _connectparticipantSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_connectparticipantStartPosition) > 0 {
		if err := assignInputField(input, "StartPosition", _connectparticipantStartPosition); err != nil {
			log.Errorf("invalid --start-position: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetTranscript(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*connectparticipant.GetTranscriptOutput
	p := connectparticipant.NewGetTranscriptPaginator(client, input)
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

// The application/vnd.amazonaws.connect.event.connection.acknowledged ContentType
// is no longer maintained since December 31, 2024. This event has been migrated to
// the [CreateParticipantConnection]API using the ConnectParticipant field.
//
// Sends an event. Message receipts are not supported when there are more than two
// active participants in the chat. Using the SendEvent API for message receipts
// when a supervisor is barged-in will result in a conflict exception.
//
// For security recommendations, see [Amazon Connect Chat security best practices].
//
// ConnectionToken is used for invoking this API instead of ParticipantToken .
//
// The Amazon Connect Participant Service APIs do not use [Signature Version 4 authentication].
//
// [CreateParticipantConnection]: https://docs.aws.amazon.com/connect-participant/latest/APIReference/API_CreateParticipantConnection.html
// [Signature Version 4 authentication]: https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html
// [Amazon Connect Chat security best practices]: https://docs.aws.amazon.com/connect/latest/adminguide/security-best-practices.html#bp-security-chat
func connectparticipant_SendEvent(cfg aws.Config, client *connectparticipant.Client) {
	input := &connectparticipant.SendEventInput{
		// ConnectionToken: *string, // Required
		// ContentType: *string, // Required
	}

	if len(_connectparticipantConnectionToken) > 0 {
		input.ConnectionToken = aws.String(_connectparticipantConnectionToken)
	}
	if len(_connectparticipantContentType) > 0 {
		input.ContentType = aws.String(_connectparticipantContentType)
	}
	if len(_connectparticipantClientToken) > 0 {
		input.ClientToken = aws.String(_connectparticipantClientToken)
	}
	if len(_connectparticipantContent) > 0 {
		input.Content = aws.String(_connectparticipantContent)
	}

	if resp, err := client.SendEvent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends a message.
// For security recommendations, see [Amazon Connect Chat security best practices].
//
// ConnectionToken is used for invoking this API instead of ParticipantToken .
//
// The Amazon Connect Participant Service APIs do not use [Signature Version 4 authentication].
//
// [Signature Version 4 authentication]: https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html
// [Amazon Connect Chat security best practices]: https://docs.aws.amazon.com/connect/latest/adminguide/security-best-practices.html#bp-security-chat
func connectparticipant_SendMessage(cfg aws.Config, client *connectparticipant.Client) {
	input := &connectparticipant.SendMessageInput{
		// ConnectionToken: *string, // Required
		// Content: *string, // Required
		// ContentType: *string, // Required
	}

	if len(_connectparticipantConnectionToken) > 0 {
		input.ConnectionToken = aws.String(_connectparticipantConnectionToken)
	}
	if len(_connectparticipantContent) > 0 {
		input.Content = aws.String(_connectparticipantContent)
	}
	if len(_connectparticipantContentType) > 0 {
		input.ContentType = aws.String(_connectparticipantContentType)
	}
	if len(_connectparticipantClientToken) > 0 {
		input.ClientToken = aws.String(_connectparticipantClientToken)
	}

	if resp, err := client.SendMessage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides a pre-signed Amazon S3 URL in response for uploading the file directly
// to S3.
//
// For security recommendations, see [Amazon Connect Chat security best practices].
//
// ConnectionToken is used for invoking this API instead of ParticipantToken .
//
// The Amazon Connect Participant Service APIs do not use [Signature Version 4 authentication].
//
// [Signature Version 4 authentication]: https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html
// [Amazon Connect Chat security best practices]: https://docs.aws.amazon.com/connect/latest/adminguide/security-best-practices.html#bp-security-chat
func connectparticipant_StartAttachmentUpload(cfg aws.Config, client *connectparticipant.Client) {
	input := &connectparticipant.StartAttachmentUploadInput{
		// AttachmentName: *string, // Required
		// AttachmentSizeInBytes: int64, // Required
		// ClientToken: *string, // Required
		// ConnectionToken: *string, // Required
		// ContentType: *string, // Required
	}

	if len(_connectparticipantAttachmentName) > 0 {
		input.AttachmentName = aws.String(_connectparticipantAttachmentName)
	}
	if len(_connectparticipantAttachmentSizeInBytes) > 0 {
		if err := assignInputField(input, "AttachmentSizeInBytes", _connectparticipantAttachmentSizeInBytes); err != nil {
			log.Errorf("invalid --attachment-size-in-bytes: %s", err.Error())
			return
		}
	}
	if len(_connectparticipantClientToken) > 0 {
		input.ClientToken = aws.String(_connectparticipantClientToken)
	}
	if len(_connectparticipantConnectionToken) > 0 {
		input.ConnectionToken = aws.String(_connectparticipantConnectionToken)
	}
	if len(_connectparticipantContentType) > 0 {
		input.ContentType = aws.String(_connectparticipantContentType)
	}

	if resp, err := client.StartAttachmentUpload(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_connectparticipantCmd)
	_connectparticipantCmd.Flags().SortFlags = false

	_connectparticipantCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_connectparticipantCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_connectparticipantCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_connectparticipantCmd.Flags().StringVarP(&_connectparticipantAttachmentId, "attachment-id", "", "", "Attachment ID")
	_connectparticipantCmd.Flags().StringSliceVarP(&_connectparticipantAttachmentIds, "attachment-ids", "", nil, "Attachment Ids")
	_connectparticipantCmd.Flags().StringVarP(&_connectparticipantAttachmentName, "attachment-name", "", "", "Attachment Name")
	_connectparticipantCmd.Flags().StringVarP(&_connectparticipantAttachmentSizeInBytes, "attachment-size-in-bytes", "", "", "Attachment Size In Bytes")
	_connectparticipantCmd.Flags().StringVarP(&_connectparticipantClientToken, "client-token", "", "", "Client Token")
	_connectparticipantCmd.Flags().StringVarP(&_connectparticipantConnectParticipant, "connect-participant", "", "", "Connect Participant")
	_connectparticipantCmd.Flags().StringVarP(&_connectparticipantConnectionToken, "connection-token", "", "", "Connection Token")
	_connectparticipantCmd.Flags().StringVarP(&_connectparticipantContactId, "contact-id", "", "", "Contact ID")
	_connectparticipantCmd.Flags().StringVarP(&_connectparticipantContent, "content", "", "", "Content")
	_connectparticipantCmd.Flags().StringVarP(&_connectparticipantContentType, "content-type", "", "", "Content Type")
	_connectparticipantCmd.Flags().StringVarP(&_connectparticipantMaxResults, "max-results", "", "", "Max Results")
	_connectparticipantCmd.Flags().StringVarP(&_connectparticipantNextToken, "next-token", "", "", "Next Token")
	_connectparticipantCmd.Flags().StringVarP(&_connectparticipantParticipantToken, "participant-token", "", "", "Participant Token")
	_connectparticipantCmd.Flags().StringVarP(&_connectparticipantRedirectUri, "redirect-uri", "", "", "Redirect URI")
	_connectparticipantCmd.Flags().StringVarP(&_connectparticipantScanDirection, "scan-direction", "", "", "Scan Direction")
	_connectparticipantCmd.Flags().StringVarP(&_connectparticipantSessionId, "session-id", "", "", "Session ID")
	_connectparticipantCmd.Flags().StringVarP(&_connectparticipantSortOrder, "sort-order", "", "", "Sort Order")
	_connectparticipantCmd.Flags().StringVarP(&_connectparticipantStartPosition, "start-position", "", "", "Start Position")
	_connectparticipantCmd.Flags().StringVarP(&_connectparticipantType, "type", "", "", "Type")
	_connectparticipantCmd.Flags().StringVarP(&_connectparticipantUrlExpiryInSeconds, "url-expiry-in-seconds", "", "", "URL Expiry In Seconds")
	_connectparticipantCmd.Flags().StringVarP(&_connectparticipantViewToken, "view-token", "", "", "View Token")

	_connectparticipantCmd.Flags().BoolVarP(&_connectparticipantCancelParticipantAuthentication, "cancel-participant-authentication", "", false, "Cancel Participant Authentication")
	_connectparticipantCmd.Flags().BoolVarP(&_connectparticipantCompleteAttachmentUpload, "complete-attachment-upload", "", false, "Complete Attachment Upload")
	_connectparticipantCmd.Flags().BoolVarP(&_connectparticipantCreateParticipantConnection, "create-participant-connection", "", false, "Create Participant Connection")
	_connectparticipantCmd.Flags().BoolVarP(&_connectparticipantDescribeView, "describe-view", "", false, "Describe View")
	_connectparticipantCmd.Flags().BoolVarP(&_connectparticipantDisconnectParticipant, "disconnect-participant", "", false, "Disconnect Participant")
	_connectparticipantCmd.Flags().BoolVarP(&_connectparticipantGetAttachment, "get-attachment", "", false, "Get Attachment")
	_connectparticipantCmd.Flags().BoolVarP(&_connectparticipantGetAuthenticationUrl, "get-authentication-url", "", false, "Get Authentication URL")
	_connectparticipantCmd.Flags().BoolVarP(&_connectparticipantGetTranscript, "get-transcript", "", false, "Get Transcript")
	_connectparticipantCmd.Flags().BoolVarP(&_connectparticipantSendEvent, "send-event", "", false, "Send Event")
	_connectparticipantCmd.Flags().BoolVarP(&_connectparticipantSendMessage, "send-message", "", false, "Send Message")
	_connectparticipantCmd.Flags().BoolVarP(&_connectparticipantStartAttachmentUpload, "start-attachment-upload", "", false, "Start Attachment Upload")

}
