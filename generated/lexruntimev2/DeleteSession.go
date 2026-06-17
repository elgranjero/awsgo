package lexruntimev2

// DeleteSession is generated as a reference stub.
// Executable command wiring lives under cmd/lexruntimev2.go.
//
// Removes session information for a specified bot, alias, and user ID.
//
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
