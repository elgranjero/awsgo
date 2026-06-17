package lexmodelbuildingservice

// DeleteBotAlias is generated as a reference stub.
// Executable command wiring lives under cmd/lexmodelbuildingservice.go.
//
// Deletes an alias for the specified bot.
//
// You can't delete an alias that is used in the association between a bot and a
// messaging channel. If an alias is used in a channel association, the DeleteBot
// operation returns a ResourceInUseException exception that includes a reference
// to the channel association that refers to the bot. You can remove the reference
// to the alias by deleting the channel association. If you get the same exception
// again, delete the referring association until the DeleteBotAlias operation is
// successful.
