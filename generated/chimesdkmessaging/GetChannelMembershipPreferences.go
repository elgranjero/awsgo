package chimesdkmessaging

// GetChannelMembershipPreferences is generated as a reference stub.
// Executable command wiring lives under cmd/chimesdkmessaging.go.
//
// Gets the membership preferences of an AppInstanceUser or AppInstanceBot for the
// specified channel. A user or a bot must be a member of the channel and own the
// membership in order to retrieve membership preferences. Users or bots in the
// AppInstanceAdmin and channel moderator roles can't retrieve preferences for
// other users or bots. Banned users or bots can't retrieve membership preferences
// for the channel from which they are banned.
//
// The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
