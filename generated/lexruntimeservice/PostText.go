package lexruntimeservice

// PostText is generated as a reference stub.
// Executable command wiring lives under cmd/lexruntimeservice.go.
//
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
