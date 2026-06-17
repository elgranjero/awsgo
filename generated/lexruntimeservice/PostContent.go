package lexruntimeservice

// PostContent is generated as a reference stub.
// Executable command wiring lives under cmd/lexruntimeservice.go.
//
// Sends user input (text or speech) to Amazon Lex. Clients use this API to send
//
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
