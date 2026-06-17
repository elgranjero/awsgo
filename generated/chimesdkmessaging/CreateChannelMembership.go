package chimesdkmessaging

// CreateChannelMembership is generated as a reference stub.
// Executable command wiring lives under cmd/chimesdkmessaging.go.
//
// Adds a member to a channel. The InvitedBy field in ChannelMembership is derived
// from the request header. A channel member can:
//
// - List messages
//
// - Send messages
//
// - Receive messages
//
// - Edit their own messages
//
// - Leave the channel
//
// Privacy settings impact this action as follows:
//
// - Public Channels: You do not need to be a member to list messages, but you
// must be a member to send messages.
//
// - Private Channels: You must be a member to list or send messages.
//
// The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUserArn or AppInstanceBot that makes the API call as the value in
// the header.
