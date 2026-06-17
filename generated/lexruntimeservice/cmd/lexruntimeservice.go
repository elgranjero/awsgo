package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lexruntimeservice"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// lexruntimeserviceCmd represents the lexruntimeservice command
var _lexruntimeserviceCmd = &cobra.Command{
	Use:   "lexruntimeservice",
	Short: "AWS lexruntimeservice CLI",
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
		client := lexruntimeservice.NewFromConfig(cfg)
		if _lexruntimeserviceDeleteSession {
			lexruntimeservice_DeleteSession(cfg, client)
			return
		}
		if _lexruntimeserviceGetSession {
			lexruntimeservice_GetSession(cfg, client)
			return
		}
		if _lexruntimeservicePostContent {
			lexruntimeservice_PostContent(cfg, client)
			return
		}
		if _lexruntimeservicePostText {
			lexruntimeservice_PostText(cfg, client)
			return
		}
		if _lexruntimeservicePutSession {
			lexruntimeservice_PutSession(cfg, client)
			return
		}

	},
}

var (
	_lexruntimeserviceDeleteSession bool
	_lexruntimeserviceGetSession    bool
	_lexruntimeservicePostContent   bool
	_lexruntimeservicePostText      bool
	_lexruntimeservicePutSession    bool

	_lexruntimeserviceAccept                  string
	_lexruntimeserviceActiveContexts          string
	_lexruntimeserviceBotAlias                string
	_lexruntimeserviceBotName                 string
	_lexruntimeserviceCheckpointLabelFilter   string
	_lexruntimeserviceContentType             string
	_lexruntimeserviceDialogAction            string
	_lexruntimeserviceInputStream             string
	_lexruntimeserviceInputText               string
	_lexruntimeserviceRecentIntentSummaryView string
	_lexruntimeserviceRequestAttributes       string
	_lexruntimeserviceSessionAttributes       string
	_lexruntimeserviceUserId                  string
)

// Removes session information for a specified bot, alias, and user ID.
func lexruntimeservice_DeleteSession(cfg aws.Config, client *lexruntimeservice.Client) {
	input := &lexruntimeservice.DeleteSessionInput{
		// BotAlias: *string, // Required
		// BotName: *string, // Required
		// UserId: *string, // Required
	}

	if len(_lexruntimeserviceBotAlias) > 0 {
		input.BotAlias = aws.String(_lexruntimeserviceBotAlias)
	}
	if len(_lexruntimeserviceBotName) > 0 {
		input.BotName = aws.String(_lexruntimeserviceBotName)
	}
	if len(_lexruntimeserviceUserId) > 0 {
		input.UserId = aws.String(_lexruntimeserviceUserId)
	}

	if resp, err := client.DeleteSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns session information for a specified bot, alias, and user ID.
func lexruntimeservice_GetSession(cfg aws.Config, client *lexruntimeservice.Client) {
	input := &lexruntimeservice.GetSessionInput{
		// BotAlias: *string, // Required
		// BotName: *string, // Required
		// UserId: *string, // Required
	}

	if len(_lexruntimeserviceBotAlias) > 0 {
		input.BotAlias = aws.String(_lexruntimeserviceBotAlias)
	}
	if len(_lexruntimeserviceBotName) > 0 {
		input.BotName = aws.String(_lexruntimeserviceBotName)
	}
	if len(_lexruntimeserviceUserId) > 0 {
		input.UserId = aws.String(_lexruntimeserviceUserId)
	}
	if len(_lexruntimeserviceCheckpointLabelFilter) > 0 {
		input.CheckpointLabelFilter = aws.String(_lexruntimeserviceCheckpointLabelFilter)
	}

	if resp, err := client.GetSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends user input (text or speech) to Amazon Lex. Clients use this API to send
// text and audio requests to Amazon Lex at runtime. Amazon Lex interprets the user
// input using the machine learning model that it built for the bot.
//
// The PostContent operation supports audio input at 8kHz and 16kHz. You can use
// 8kHz audio to achieve higher speech recognition accuracy in telephone audio
// applications.
//
// In response, Amazon Lex returns the next message to convey to the user.
// Consider the following example messages:
//
// - For a user input "I would like a pizza," Amazon Lex might return a response
// with a message eliciting slot data (for example, PizzaSize ): "What size pizza
// would you like?".
//
// - After the user provides all of the pizza order information, Amazon Lex
// might return a response with a message to get user confirmation: "Order the
// pizza?".
//
// - After the user replies "Yes" to the confirmation prompt, Amazon Lex might
// return a conclusion statement: "Thank you, your cheese pizza has been ordered.".
//
// Not all Amazon Lex messages require a response from the user. For example,
// conclusion statements do not require a response. Some messages require only a
// yes or no response. In addition to the message , Amazon Lex provides additional
// context about the message in the response that you can use to enhance client
// behavior, such as displaying the appropriate client user interface. Consider the
// following examples:
//
// - If the message is to elicit slot data, Amazon Lex returns the following
// context information:
//
// - x-amz-lex-dialog-state header set to ElicitSlot
//
// - x-amz-lex-intent-name header set to the intent name in the current context
//
// - x-amz-lex-slot-to-elicit header set to the slot name for which the message
// is eliciting information
//
// - x-amz-lex-slots header set to a map of slots configured for the intent with
// their current values
//
// - If the message is a confirmation prompt, the x-amz-lex-dialog-state header
// is set to Confirmation and the x-amz-lex-slot-to-elicit header is omitted.
//
// - If the message is a clarification prompt configured for the intent,
// indicating that the user intent is not understood, the x-amz-dialog-state
// header is set to ElicitIntent and the x-amz-slot-to-elicit header is omitted.
//
// In addition, Amazon Lex also returns your application-specific sessionAttributes
// . For more information, see [Managing Conversation Context].
//
// [Managing Conversation Context]: https://docs.aws.amazon.com/lex/latest/dg/context-mgmt.html
func lexruntimeservice_PostContent(cfg aws.Config, client *lexruntimeservice.Client) {
	input := &lexruntimeservice.PostContentInput{
		// BotAlias: *string, // Required
		// BotName: *string, // Required
		// ContentType: *string, // Required
		// InputStream: io.Reader, // Required
		// UserId: *string, // Required
	}

	if len(_lexruntimeserviceBotAlias) > 0 {
		input.BotAlias = aws.String(_lexruntimeserviceBotAlias)
	}
	if len(_lexruntimeserviceBotName) > 0 {
		input.BotName = aws.String(_lexruntimeserviceBotName)
	}
	if len(_lexruntimeserviceContentType) > 0 {
		input.ContentType = aws.String(_lexruntimeserviceContentType)
	}
	if len(_lexruntimeserviceInputStream) > 0 {
		if err := assignInputField(input, "InputStream", _lexruntimeserviceInputStream); err != nil {
			log.Errorf("invalid --input-stream: %s", err.Error())
			return
		}
	}
	if len(_lexruntimeserviceUserId) > 0 {
		input.UserId = aws.String(_lexruntimeserviceUserId)
	}
	if len(_lexruntimeserviceAccept) > 0 {
		input.Accept = aws.String(_lexruntimeserviceAccept)
	}
	if len(_lexruntimeserviceActiveContexts) > 0 {
		input.ActiveContexts = aws.String(_lexruntimeserviceActiveContexts)
	}
	if len(_lexruntimeserviceRequestAttributes) > 0 {
		input.RequestAttributes = aws.String(_lexruntimeserviceRequestAttributes)
	}
	if len(_lexruntimeserviceSessionAttributes) > 0 {
		input.SessionAttributes = aws.String(_lexruntimeserviceSessionAttributes)
	}

	if resp, err := client.PostContent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends user input to Amazon Lex. Client applications can use this API to send
// requests to Amazon Lex at runtime. Amazon Lex then interprets the user input
// using the machine learning model it built for the bot.
//
// In response, Amazon Lex returns the next message to convey to the user an
// optional responseCard to display. Consider the following example messages:
//
// - For a user input "I would like a pizza", Amazon Lex might return a response
// with a message eliciting slot data (for example, PizzaSize): "What size pizza
// would you like?"
//
// - After the user provides all of the pizza order information, Amazon Lex
// might return a response with a message to obtain user confirmation "Proceed with
// the pizza order?".
//
// - After the user replies to a confirmation prompt with a "yes", Amazon Lex
// might return a conclusion statement: "Thank you, your cheese pizza has been
// ordered.".
//
// Not all Amazon Lex messages require a user response. For example, a conclusion
// statement does not require a response. Some messages require only a "yes" or
// "no" user response. In addition to the message , Amazon Lex provides additional
// context about the message in the response that you might use to enhance client
// behavior, for example, to display the appropriate client user interface. These
// are the slotToElicit , dialogState , intentName , and slots fields in the
// response. Consider the following examples:
//
// - If the message is to elicit slot data, Amazon Lex returns the following
// context information:
//
// - dialogState set to ElicitSlot
//
// - intentName set to the intent name in the current context
//
// - slotToElicit set to the slot name for which the message is eliciting
// information
//
// - slots set to a map of slots, configured for the intent, with currently known
// values
//
// - If the message is a confirmation prompt, the dialogState is set to
// ConfirmIntent and SlotToElicit is set to null.
//
// - If the message is a clarification prompt (configured for the intent) that
// indicates that user intent is not understood, the dialogState is set to
// ElicitIntent and slotToElicit is set to null.
//
// In addition, Amazon Lex also returns your application-specific sessionAttributes
// . For more information, see [Managing Conversation Context].
//
// [Managing Conversation Context]: https://docs.aws.amazon.com/lex/latest/dg/context-mgmt.html
func lexruntimeservice_PostText(cfg aws.Config, client *lexruntimeservice.Client) {
	input := &lexruntimeservice.PostTextInput{
		// BotAlias: *string, // Required
		// BotName: *string, // Required
		// InputText: *string, // Required
		// UserId: *string, // Required
	}

	if len(_lexruntimeserviceBotAlias) > 0 {
		input.BotAlias = aws.String(_lexruntimeserviceBotAlias)
	}
	if len(_lexruntimeserviceBotName) > 0 {
		input.BotName = aws.String(_lexruntimeserviceBotName)
	}
	if len(_lexruntimeserviceInputText) > 0 {
		input.InputText = aws.String(_lexruntimeserviceInputText)
	}
	if len(_lexruntimeserviceUserId) > 0 {
		input.UserId = aws.String(_lexruntimeserviceUserId)
	}
	if len(_lexruntimeserviceActiveContexts) > 0 {
		if err := assignInputField(input, "ActiveContexts", _lexruntimeserviceActiveContexts); err != nil {
			log.Errorf("invalid --active-contexts: %s", err.Error())
			return
		}
	}
	if len(_lexruntimeserviceRequestAttributes) > 0 {
		if err := assignInputField(input, "RequestAttributes", _lexruntimeserviceRequestAttributes); err != nil {
			log.Errorf("invalid --request-attributes: %s", err.Error())
			return
		}
	}
	if len(_lexruntimeserviceSessionAttributes) > 0 {
		if err := assignInputField(input, "SessionAttributes", _lexruntimeserviceSessionAttributes); err != nil {
			log.Errorf("invalid --session-attributes: %s", err.Error())
			return
		}
	}

	if resp, err := client.PostText(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new session or modifies an existing session with an Amazon Lex bot.
// Use this operation to enable your application to set the state of the bot.
//
// For more information, see [Managing Sessions].
//
// [Managing Sessions]: https://docs.aws.amazon.com/lex/latest/dg/how-session-api.html
func lexruntimeservice_PutSession(cfg aws.Config, client *lexruntimeservice.Client) {
	input := &lexruntimeservice.PutSessionInput{
		// BotAlias: *string, // Required
		// BotName: *string, // Required
		// UserId: *string, // Required
	}

	if len(_lexruntimeserviceBotAlias) > 0 {
		input.BotAlias = aws.String(_lexruntimeserviceBotAlias)
	}
	if len(_lexruntimeserviceBotName) > 0 {
		input.BotName = aws.String(_lexruntimeserviceBotName)
	}
	if len(_lexruntimeserviceUserId) > 0 {
		input.UserId = aws.String(_lexruntimeserviceUserId)
	}
	if len(_lexruntimeserviceAccept) > 0 {
		input.Accept = aws.String(_lexruntimeserviceAccept)
	}
	if len(_lexruntimeserviceActiveContexts) > 0 {
		if err := assignInputField(input, "ActiveContexts", _lexruntimeserviceActiveContexts); err != nil {
			log.Errorf("invalid --active-contexts: %s", err.Error())
			return
		}
	}
	if len(_lexruntimeserviceDialogAction) > 0 {
		if err := assignInputField(input, "DialogAction", _lexruntimeserviceDialogAction); err != nil {
			log.Errorf("invalid --dialog-action: %s", err.Error())
			return
		}
	}
	if len(_lexruntimeserviceRecentIntentSummaryView) > 0 {
		if err := assignInputField(input, "RecentIntentSummaryView", _lexruntimeserviceRecentIntentSummaryView); err != nil {
			log.Errorf("invalid --recent-intent-summary-view: %s", err.Error())
			return
		}
	}
	if len(_lexruntimeserviceSessionAttributes) > 0 {
		if err := assignInputField(input, "SessionAttributes", _lexruntimeserviceSessionAttributes); err != nil {
			log.Errorf("invalid --session-attributes: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_lexruntimeserviceCmd)
	_lexruntimeserviceCmd.Flags().SortFlags = false

	_lexruntimeserviceCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_lexruntimeserviceCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_lexruntimeserviceCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_lexruntimeserviceCmd.Flags().StringVarP(&_lexruntimeserviceAccept, "accept", "", "", "Accept")
	_lexruntimeserviceCmd.Flags().StringVarP(&_lexruntimeserviceActiveContexts, "active-contexts", "", "", "Active Contexts")
	_lexruntimeserviceCmd.Flags().StringVarP(&_lexruntimeserviceBotAlias, "bot-alias", "", "", "Bot Alias")
	_lexruntimeserviceCmd.Flags().StringVarP(&_lexruntimeserviceBotName, "bot-name", "", "", "Bot Name")
	_lexruntimeserviceCmd.Flags().StringVarP(&_lexruntimeserviceCheckpointLabelFilter, "checkpoint-label-filter", "", "", "Checkpoint Label Filter")
	_lexruntimeserviceCmd.Flags().StringVarP(&_lexruntimeserviceContentType, "content-type", "", "", "Content Type")
	_lexruntimeserviceCmd.Flags().StringVarP(&_lexruntimeserviceDialogAction, "dialog-action", "", "", "Dialog Action")
	_lexruntimeserviceCmd.Flags().StringVarP(&_lexruntimeserviceInputStream, "input-stream", "", "", "Input Stream")
	_lexruntimeserviceCmd.Flags().StringVarP(&_lexruntimeserviceInputText, "input-text", "", "", "Input Text")
	_lexruntimeserviceCmd.Flags().StringVarP(&_lexruntimeserviceRecentIntentSummaryView, "recent-intent-summary-view", "", "", "Recent Intent Summary View")
	_lexruntimeserviceCmd.Flags().StringVarP(&_lexruntimeserviceRequestAttributes, "request-attributes", "", "", "Request Attributes")
	_lexruntimeserviceCmd.Flags().StringVarP(&_lexruntimeserviceSessionAttributes, "session-attributes", "", "", "Session Attributes")
	_lexruntimeserviceCmd.Flags().StringVarP(&_lexruntimeserviceUserId, "user-id", "", "", "User ID")

	_lexruntimeserviceCmd.Flags().BoolVarP(&_lexruntimeserviceDeleteSession, "delete-session", "", false, "Delete Session")
	_lexruntimeserviceCmd.Flags().BoolVarP(&_lexruntimeserviceGetSession, "get-session", "", false, "Get Session")
	_lexruntimeserviceCmd.Flags().BoolVarP(&_lexruntimeservicePostContent, "post-content", "", false, "Post Content")
	_lexruntimeserviceCmd.Flags().BoolVarP(&_lexruntimeservicePostText, "post-text", "", false, "Post Text")
	_lexruntimeserviceCmd.Flags().BoolVarP(&_lexruntimeservicePutSession, "put-session", "", false, "Put Session")

}
