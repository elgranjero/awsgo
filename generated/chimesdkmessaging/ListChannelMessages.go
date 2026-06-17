package chimesdkmessaging

// ListChannelMessages is generated as a reference stub.
// Executable command wiring lives under cmd/chimesdkmessaging.go.
//
// List all the messages in a channel. Returns a paginated list of ChannelMessages
// . By default, sorted by creation timestamp in descending order.
//
// Redacted messages appear in the results as empty, since they are only redacted,
// not deleted. Deleted messages do not appear in the results. This action always
// returns the latest version of an edited message.
//
// Also, the x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
