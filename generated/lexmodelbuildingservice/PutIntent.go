package lexmodelbuildingservice

// PutIntent is generated as a reference stub.
// Executable command wiring lives under cmd/lexmodelbuildingservice.go.
//
// Creates an intent or replaces an existing intent.
//
// To define the interaction between the user and your bot, you use one or more
// intents. For a pizza ordering bot, for example, you would create an OrderPizza
// intent.
//
// To create an intent or replace an existing intent, you must provide the
// following:
//
// - Intent name. For example, OrderPizza .
//
// - Sample utterances. For example, "Can I order a pizza, please." and "I want
// to order a pizza."
//
// - Information to be gathered. You specify slot types for the information that
// your bot will request from the user. You can specify standard slot types, such
// as a date or a time, or custom slot types such as the size and crust of a pizza.
//
// - How the intent will be fulfilled. You can provide a Lambda function or
// configure the intent to return the intent information to the client application.
// If you use a Lambda function, when all of the intent information is available,
// Amazon Lex invokes your Lambda function. If you configure your intent to return
// the intent information to the client application.
//
// You can specify other optional information in the request, such as:
//
// - A confirmation prompt to ask the user to confirm an intent. For example,
// "Shall I order your pizza?"
//
// - A conclusion statement to send to the user after the intent has been
// fulfilled. For example, "I placed your pizza order."
//
// - A follow-up prompt that asks the user for additional activity. For example,
// asking "Do you want to order a drink with your pizza?"
//
// If you specify an existing intent name to update the intent, Amazon Lex
// replaces the values in the $LATEST version of the intent with the values in the
// request. Amazon Lex removes fields that you don't provide in the request. If you
// don't specify the required fields, Amazon Lex throws an exception. When you
// update the $LATEST version of an intent, the status field of any bot that uses
// the $LATEST version of the intent is set to NOT_BUILT .
//
// For more information, see how-it-works.
//
// This operation requires permissions for the lex:PutIntent action.
