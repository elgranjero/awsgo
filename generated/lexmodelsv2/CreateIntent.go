package lexmodelsv2

// CreateIntent is generated as a reference stub.
// Executable command wiring lives under cmd/lexmodelsv2.go.
//
// Creates an intent.
//
// To define the interaction between the user and your bot, you define one or more
// intents. For example, for a pizza ordering bot you would create an OrderPizza
// intent.
//
// When you create an intent, you must provide a name. You can optionally provide
// the following:
//
// - Sample utterances. For example, "I want to order a pizza" and "Can I order
// a pizza." You can't provide utterances for built-in intents.
//
// - Information to be gathered. You specify slots for the information that you
// bot requests from the user. You can specify standard slot types, such as date
// and time, or custom slot types for your application.
//
// - How the intent is fulfilled. You can provide a Lambda function or configure
// the intent to return the intent information to your client application. If you
// use a Lambda function, Amazon Lex invokes the function when all of the intent
// information is available.
//
// - A confirmation prompt to send to the user to confirm an intent. For
// example, "Shall I order your pizza?"
//
// - A conclusion statement to send to the user after the intent is fulfilled.
// For example, "I ordered your pizza."
//
// - A follow-up prompt that asks the user for additional activity. For example,
// "Do you want a drink with your pizza?"
