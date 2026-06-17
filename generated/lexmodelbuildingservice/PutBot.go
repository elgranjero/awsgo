package lexmodelbuildingservice

// PutBot is generated as a reference stub.
// Executable command wiring lives under cmd/lexmodelbuildingservice.go.
//
// Creates an Amazon Lex conversational bot or replaces an existing bot. When you
// create or update a bot you are only required to specify a name, a locale, and
// whether the bot is directed toward children under age 13. You can use this to
// add intents later, or to remove intents from an existing bot. When you create a
// bot with the minimum information, the bot is created or updated but Amazon Lex
// returns the response FAILED . You can build the bot after you add one or more
// intents. For more information about Amazon Lex bots, see how-it-works.
//
// If you specify the name of an existing bot, the fields in the request replace
// the existing values in the $LATEST version of the bot. Amazon Lex removes any
// fields that you don't provide values for in the request, except for the
// idleTTLInSeconds and privacySettings fields, which are set to their default
// values. If you don't specify values for required fields, Amazon Lex throws an
// exception.
//
// This operation requires permissions for the lex:PutBot action. For more
// information, see security-iam.
