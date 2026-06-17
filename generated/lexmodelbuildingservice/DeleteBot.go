package lexmodelbuildingservice

// DeleteBot is generated as a reference stub.
// Executable command wiring lives under cmd/lexmodelbuildingservice.go.
//
// Deletes all versions of the bot, including the $LATEST version. To delete a
// specific version of the bot, use the DeleteBotVersionoperation. The DeleteBot operation doesn't
// immediately remove the bot schema. Instead, it is marked for deletion and
// removed later.
//
// Amazon Lex stores utterances indefinitely for improving the ability of your bot
// to respond to user inputs. These utterances are not removed when the bot is
// deleted. To remove the utterances, use the DeleteUtterancesoperation.
//
// If a bot has an alias, you can't delete it. Instead, the DeleteBot operation
// returns a ResourceInUseException exception that includes a reference to the
// alias that refers to the bot. To remove the reference to the bot, delete the
// alias. If you get the same exception again, delete the referring alias until the
// DeleteBot operation is successful.
//
// This operation requires permissions for the lex:DeleteBot action.
