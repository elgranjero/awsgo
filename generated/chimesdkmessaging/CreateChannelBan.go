package chimesdkmessaging

// CreateChannelBan is generated as a reference stub.
// Executable command wiring lives under cmd/chimesdkmessaging.go.
//
// Permanently bans a member from a channel. Moderators can't add banned members
// to a channel. To undo a ban, you first have to DeleteChannelBan , and then
// CreateChannelMembership . Bans are cleaned up when you delete users or channels.
//
// If you ban a user who is already part of a channel, that user is automatically
// kicked from the channel.
//
// The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
