package lexmodelsv2

// DeleteBot is generated as a reference stub.
// Executable command wiring lives under cmd/lexmodelsv2.go.
//
// Deletes all versions of a bot, including the Draft version. To delete a
// specific version, use the DeleteBotVersion operation.
//
// When you delete a bot, all of the resources contained in the bot are also
// deleted. Deleting a bot removes all locales, intents, slot, and slot types
// defined for the bot.
//
// If a bot has an alias, the DeleteBot operation returns a ResourceInUseException
// exception. If you want to delete the bot and the alias, set the
// skipResourceInUseCheck parameter to true .
