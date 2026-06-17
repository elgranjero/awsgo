package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lexruntimev2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// lexruntimev2Cmd represents the lexruntimev2 command
var _lexruntimev2Cmd = &cobra.Command{
	Use:   "lexruntimev2",
	Short: "AWS lexruntimev2 CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := lexruntimev2.NewFromConfig(cfg)
		if _lexruntimev2DeleteSession {
			lexruntimev2_DeleteSession(cfg, client)
			return
		}
		if _lexruntimev2GetSession {
			lexruntimev2_GetSession(cfg, client)
			return
		}
		if _lexruntimev2PutSession {
			lexruntimev2_PutSession(cfg, client)
			return
		}
		if _lexruntimev2RecognizeText {
			lexruntimev2_RecognizeText(cfg, client)
			return
		}
		if _lexruntimev2RecognizeUtterance {
			lexruntimev2_RecognizeUtterance(cfg, client)
			return
		}
		if _lexruntimev2StartConversation {
			lexruntimev2_StartConversation(cfg, client)
			return
		}

	},
}

var (
	_lexruntimev2DeleteSession      bool
	_lexruntimev2GetSession         bool
	_lexruntimev2PutSession         bool
	_lexruntimev2RecognizeText      bool
	_lexruntimev2RecognizeUtterance bool
	_lexruntimev2StartConversation  bool

	_lexruntimev2BotAliasId          string
	_lexruntimev2BotId               string
	_lexruntimev2ConversationMode    string
	_lexruntimev2InputStream         string
	_lexruntimev2LocaleId            string
	_lexruntimev2Messages            string
	_lexruntimev2RequestAttributes   string
	_lexruntimev2RequestContentType  string
	_lexruntimev2ResponseContentType string
	_lexruntimev2SessionId           string
	_lexruntimev2SessionState        string
	_lexruntimev2Text                string
)

// Removes session information for a specified bot, alias, and user ID.
// You can use this operation to restart a conversation with a bot. When you
// remove a session, the entire history of the session is removed so that you can
// start again.
//
// You don't need to delete a session. Sessions have a time limit and will expire.
// Set the session time limit when you create the bot. The default is 5 minutes,
// but you can specify anything between 1 minute and 24 hours.
//
// If you specify a bot or alias ID that doesn't exist, you receive a
// BadRequestException.
//
// If the locale doesn't exist in the bot, or if the locale hasn't been enables
// for the alias, you receive a BadRequestException .
func lexruntimev2_DeleteSession(cfg aws.Config, client *lexruntimev2.Client) {
	input := &lexruntimev2.DeleteSessionInput{
		// BotAliasId: *string, // Required
		// BotId: *string, // Required
		// LocaleId: *string, // Required
		// SessionId: *string, // Required
	}

	if len(_lexruntimev2BotAliasId) > 0 {
		input.BotAliasId = aws.String(_lexruntimev2BotAliasId)
	}
	if len(_lexruntimev2BotId) > 0 {
		input.BotId = aws.String(_lexruntimev2BotId)
	}
	if len(_lexruntimev2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexruntimev2LocaleId)
	}
	if len(_lexruntimev2SessionId) > 0 {
		input.SessionId = aws.String(_lexruntimev2SessionId)
	}

	if resp, err := client.DeleteSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns session information for a specified bot, alias, and user.
// For example, you can use this operation to retrieve session information for a
// user that has left a long-running session in use.
//
// If the bot, alias, or session identifier doesn't exist, Amazon Lex V2 returns a
// BadRequestException . If the locale doesn't exist or is not enabled for the
// alias, you receive a BadRequestException .
func lexruntimev2_GetSession(cfg aws.Config, client *lexruntimev2.Client) {
	input := &lexruntimev2.GetSessionInput{
		// BotAliasId: *string, // Required
		// BotId: *string, // Required
		// LocaleId: *string, // Required
		// SessionId: *string, // Required
	}

	if len(_lexruntimev2BotAliasId) > 0 {
		input.BotAliasId = aws.String(_lexruntimev2BotAliasId)
	}
	if len(_lexruntimev2BotId) > 0 {
		input.BotId = aws.String(_lexruntimev2BotId)
	}
	if len(_lexruntimev2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexruntimev2LocaleId)
	}
	if len(_lexruntimev2SessionId) > 0 {
		input.SessionId = aws.String(_lexruntimev2SessionId)
	}

	if resp, err := client.GetSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new session or modifies an existing session with an Amazon Lex V2
// bot. Use this operation to enable your application to set the state of the bot.
func lexruntimev2_PutSession(cfg aws.Config, client *lexruntimev2.Client) {
	input := &lexruntimev2.PutSessionInput{
		// BotAliasId: *string, // Required
		// BotId: *string, // Required
		// LocaleId: *string, // Required
		// SessionId: *string, // Required
		// SessionState: *types.SessionState, // Required
	}

	if len(_lexruntimev2BotAliasId) > 0 {
		input.BotAliasId = aws.String(_lexruntimev2BotAliasId)
	}
	if len(_lexruntimev2BotId) > 0 {
		input.BotId = aws.String(_lexruntimev2BotId)
	}
	if len(_lexruntimev2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexruntimev2LocaleId)
	}
	if len(_lexruntimev2SessionId) > 0 {
		input.SessionId = aws.String(_lexruntimev2SessionId)
	}
	if len(_lexruntimev2SessionState) > 0 {
		if err := assignInputField(input, "SessionState", _lexruntimev2SessionState); err != nil {
			log.Errorf("invalid --session-state: %s", err.Error())
			return
		}
	}
	if len(_lexruntimev2Messages) > 0 {
		if err := assignInputField(input, "Messages", _lexruntimev2Messages); err != nil {
			log.Errorf("invalid --messages: %s", err.Error())
			return
		}
	}
	if len(_lexruntimev2RequestAttributes) > 0 {
		if err := assignInputField(input, "RequestAttributes", _lexruntimev2RequestAttributes); err != nil {
			log.Errorf("invalid --request-attributes: %s", err.Error())
			return
		}
	}
	if len(_lexruntimev2ResponseContentType) > 0 {
		input.ResponseContentType = aws.String(_lexruntimev2ResponseContentType)
	}

	if resp, err := client.PutSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends user input to Amazon Lex V2. Client applications use this API to send
// requests to Amazon Lex V2 at runtime. Amazon Lex V2 then interprets the user
// input using the machine learning model that it build for the bot.
//
// In response, Amazon Lex V2 returns the next message to convey to the user and
// an optional response card to display.
//
// If the optional post-fulfillment response is specified, the messages are
// returned as follows. For more information, see [PostFulfillmentStatusSpecification].
//
// - Success message - Returned if the Lambda function completes successfully
// and the intent state is fulfilled or ready fulfillment if the message is
// present.
//
// - Failed message - The failed message is returned if the Lambda function
// throws an exception or if the Lambda function returns a failed intent state
// without a message.
//
// - Timeout message - If you don't configure a timeout message and a timeout,
// and the Lambda function doesn't return within 30 seconds, the timeout message is
// returned. If you configure a timeout, the timeout message is returned when the
// period times out.
//
// For more information, see [Completion message].
//
// [PostFulfillmentStatusSpecification]: https://docs.aws.amazon.com/lexv2/latest/dg/API_PostFulfillmentStatusSpecification.html
// [Completion message]: https://docs.aws.amazon.com/lexv2/latest/dg/streaming-progress.html#progress-complete.html
func lexruntimev2_RecognizeText(cfg aws.Config, client *lexruntimev2.Client) {
	input := &lexruntimev2.RecognizeTextInput{
		// BotAliasId: *string, // Required
		// BotId: *string, // Required
		// LocaleId: *string, // Required
		// SessionId: *string, // Required
		// Text: *string, // Required
	}

	if len(_lexruntimev2BotAliasId) > 0 {
		input.BotAliasId = aws.String(_lexruntimev2BotAliasId)
	}
	if len(_lexruntimev2BotId) > 0 {
		input.BotId = aws.String(_lexruntimev2BotId)
	}
	if len(_lexruntimev2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexruntimev2LocaleId)
	}
	if len(_lexruntimev2SessionId) > 0 {
		input.SessionId = aws.String(_lexruntimev2SessionId)
	}
	if len(_lexruntimev2Text) > 0 {
		input.Text = aws.String(_lexruntimev2Text)
	}
	if len(_lexruntimev2RequestAttributes) > 0 {
		if err := assignInputField(input, "RequestAttributes", _lexruntimev2RequestAttributes); err != nil {
			log.Errorf("invalid --request-attributes: %s", err.Error())
			return
		}
	}
	if len(_lexruntimev2SessionState) > 0 {
		if err := assignInputField(input, "SessionState", _lexruntimev2SessionState); err != nil {
			log.Errorf("invalid --session-state: %s", err.Error())
			return
		}
	}

	if resp, err := client.RecognizeText(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends user input to Amazon Lex V2. You can send text or speech. Clients use
// this API to send text and audio requests to Amazon Lex V2 at runtime. Amazon Lex
// V2 interprets the user input using the machine learning model built for the bot.
//
// The following request fields must be compressed with gzip and then base64
// encoded before you send them to Amazon Lex V2.
//
// - requestAttributes
//
// - sessionState
//
// The following response fields are compressed using gzip and then base64 encoded
// by Amazon Lex V2. Before you can use these fields, you must decode and
// decompress them.
//
// - inputTranscript
//
// - interpretations
//
// - messages
//
// - requestAttributes
//
// - sessionState
//
// The example contains a Java application that compresses and encodes a Java
// object to send to Amazon Lex V2, and a second that decodes and decompresses a
// response from Amazon Lex V2.
//
// If the optional post-fulfillment response is specified, the messages are
// returned as follows. For more information, see [PostFulfillmentStatusSpecification].
//
// - Success message - Returned if the Lambda function completes successfully
// and the intent state is fulfilled or ready fulfillment if the message is
// present.
//
// - Failed message - The failed message is returned if the Lambda function
// throws an exception or if the Lambda function returns a failed intent state
// without a message.
//
// - Timeout message - If you don't configure a timeout message and a timeout,
// and the Lambda function doesn't return within 30 seconds, the timeout message is
// returned. If you configure a timeout, the timeout message is returned when the
// period times out.
//
// For more information, see [Completion message].
//
// [PostFulfillmentStatusSpecification]: https://docs.aws.amazon.com/lexv2/latest/dg/API_PostFulfillmentStatusSpecification.html
// [Completion message]: https://docs.aws.amazon.com/lexv2/latest/dg/streaming-progress.html#progress-complete.html
func lexruntimev2_RecognizeUtterance(cfg aws.Config, client *lexruntimev2.Client) {
	input := &lexruntimev2.RecognizeUtteranceInput{
		// BotAliasId: *string, // Required
		// BotId: *string, // Required
		// LocaleId: *string, // Required
		// RequestContentType: *string, // Required
		// SessionId: *string, // Required
	}

	if len(_lexruntimev2BotAliasId) > 0 {
		input.BotAliasId = aws.String(_lexruntimev2BotAliasId)
	}
	if len(_lexruntimev2BotId) > 0 {
		input.BotId = aws.String(_lexruntimev2BotId)
	}
	if len(_lexruntimev2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexruntimev2LocaleId)
	}
	if len(_lexruntimev2RequestContentType) > 0 {
		input.RequestContentType = aws.String(_lexruntimev2RequestContentType)
	}
	if len(_lexruntimev2SessionId) > 0 {
		input.SessionId = aws.String(_lexruntimev2SessionId)
	}
	if len(_lexruntimev2InputStream) > 0 {
		if err := assignInputField(input, "InputStream", _lexruntimev2InputStream); err != nil {
			log.Errorf("invalid --input-stream: %s", err.Error())
			return
		}
	}
	if len(_lexruntimev2RequestAttributes) > 0 {
		input.RequestAttributes = aws.String(_lexruntimev2RequestAttributes)
	}
	if len(_lexruntimev2ResponseContentType) > 0 {
		input.ResponseContentType = aws.String(_lexruntimev2ResponseContentType)
	}
	if len(_lexruntimev2SessionState) > 0 {
		input.SessionState = aws.String(_lexruntimev2SessionState)
	}

	if resp, err := client.RecognizeUtterance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an HTTP/2 bidirectional event stream that enables you to send audio,
// text, or DTMF input in real time. After your application starts a conversation,
// users send input to Amazon Lex V2 as a stream of events. Amazon Lex V2 processes
// the incoming events and responds with streaming text or audio events.
//
// Audio input must be in the following format: audio/lpcm sample-rate=8000
// sample-size-bits=16 channel-count=1; is-big-endian=false .
//
// If the optional post-fulfillment response is specified, the messages are
// returned as follows. For more information, see [PostFulfillmentStatusSpecification].
//
// - Success message - Returned if the Lambda function completes successfully
// and the intent state is fulfilled or ready fulfillment if the message is
// present.
//
// - Failed message - The failed message is returned if the Lambda function
// throws an exception or if the Lambda function returns a failed intent state
// without a message.
//
// - Timeout message - If you don't configure a timeout message and a timeout,
// and the Lambda function doesn't return within 30 seconds, the timeout message is
// returned. If you configure a timeout, the timeout message is returned when the
// period times out.
//
// For more information, see [Completion message].
//
// If the optional update message is configured, it is played at the specified
// frequency while the Lambda function is running and the update message state is
// active. If the fulfillment update message is not active, the Lambda function
// runs with a 30 second timeout.
//
// For more information, see [Update message]
//
// The StartConversation operation is supported only in the following SDKs:
//
// [AWS SDK for C++]
//
// [AWS SDK for Java V2]
//
// [AWS SDK for Ruby V3]
//
// [AWS SDK for Ruby V3]: https://docs.aws.amazon.com/goto/SdkForRubyV3/runtime.lex.v2-2020-08-07/StartConversation
// [PostFulfillmentStatusSpecification]: https://docs.aws.amazon.com/lexv2/latest/dg/API_PostFulfillmentStatusSpecification.html
// [AWS SDK for Java V2]: https://docs.aws.amazon.com/goto/SdkForJavaV2/runtime.lex.v2-2020-08-07/StartConversation
// [AWS SDK for C++]: https://docs.aws.amazon.com/goto/SdkForCpp/runtime.lex.v2-2020-08-07/StartConversation
// [Update message]: https://docs.aws.amazon.com/lexv2/latest/dg/streaming-progress.html#progress-update.html
// [Completion message]: https://docs.aws.amazon.com/lexv2/latest/dg/streaming-progress.html#progress-complete.html
func lexruntimev2_StartConversation(cfg aws.Config, client *lexruntimev2.Client) {
	input := &lexruntimev2.StartConversationInput{
		// BotAliasId: *string, // Required
		// BotId: *string, // Required
		// LocaleId: *string, // Required
		// SessionId: *string, // Required
	}

	if len(_lexruntimev2BotAliasId) > 0 {
		input.BotAliasId = aws.String(_lexruntimev2BotAliasId)
	}
	if len(_lexruntimev2BotId) > 0 {
		input.BotId = aws.String(_lexruntimev2BotId)
	}
	if len(_lexruntimev2LocaleId) > 0 {
		input.LocaleId = aws.String(_lexruntimev2LocaleId)
	}
	if len(_lexruntimev2SessionId) > 0 {
		input.SessionId = aws.String(_lexruntimev2SessionId)
	}
	if len(_lexruntimev2ConversationMode) > 0 {
		if err := assignInputField(input, "ConversationMode", _lexruntimev2ConversationMode); err != nil {
			log.Errorf("invalid --conversation-mode: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartConversation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_lexruntimev2Cmd)
	_lexruntimev2Cmd.Flags().SortFlags = false

	_lexruntimev2Cmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_lexruntimev2Cmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_lexruntimev2Cmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_lexruntimev2Cmd.Flags().StringVarP(&_lexruntimev2BotAliasId, "bot-alias-id", "", "", "Bot Alias ID")
	_lexruntimev2Cmd.Flags().StringVarP(&_lexruntimev2BotId, "bot-id", "", "", "Bot ID")
	_lexruntimev2Cmd.Flags().StringVarP(&_lexruntimev2ConversationMode, "conversation-mode", "", "", "Conversation Mode")
	_lexruntimev2Cmd.Flags().StringVarP(&_lexruntimev2InputStream, "input-stream", "", "", "Input Stream")
	_lexruntimev2Cmd.Flags().StringVarP(&_lexruntimev2LocaleId, "locale-id", "", "", "Locale ID")
	_lexruntimev2Cmd.Flags().StringVarP(&_lexruntimev2Messages, "messages", "", "", "Messages")
	_lexruntimev2Cmd.Flags().StringVarP(&_lexruntimev2RequestAttributes, "request-attributes", "", "", "Request Attributes")
	_lexruntimev2Cmd.Flags().StringVarP(&_lexruntimev2RequestContentType, "request-content-type", "", "", "Request Content Type")
	_lexruntimev2Cmd.Flags().StringVarP(&_lexruntimev2ResponseContentType, "response-content-type", "", "", "Response Content Type")
	_lexruntimev2Cmd.Flags().StringVarP(&_lexruntimev2SessionId, "session-id", "", "", "Session ID")
	_lexruntimev2Cmd.Flags().StringVarP(&_lexruntimev2SessionState, "session-state", "", "", "Session State")
	_lexruntimev2Cmd.Flags().StringVarP(&_lexruntimev2Text, "text", "", "", "Text")

	_lexruntimev2Cmd.Flags().BoolVarP(&_lexruntimev2DeleteSession, "delete-session", "", false, "Delete Session")
	_lexruntimev2Cmd.Flags().BoolVarP(&_lexruntimev2GetSession, "get-session", "", false, "Get Session")
	_lexruntimev2Cmd.Flags().BoolVarP(&_lexruntimev2PutSession, "put-session", "", false, "Put Session")
	_lexruntimev2Cmd.Flags().BoolVarP(&_lexruntimev2RecognizeText, "recognize-text", "", false, "Recognize Text")
	_lexruntimev2Cmd.Flags().BoolVarP(&_lexruntimev2RecognizeUtterance, "recognize-utterance", "", false, "Recognize Utterance")
	_lexruntimev2Cmd.Flags().BoolVarP(&_lexruntimev2StartConversation, "start-conversation", "", false, "Start Conversation")

}
