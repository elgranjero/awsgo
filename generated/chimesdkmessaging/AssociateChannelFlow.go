package chimesdkmessaging

// AssociateChannelFlow is generated as a reference stub.
// Executable command wiring lives under cmd/chimesdkmessaging.go.
//
// Associates a channel flow with a channel. Once associated, all messages to that
// channel go through channel flow processors. To stop processing, use the
// DisassociateChannelFlow API.
//
// Only administrators or channel moderators can associate a channel flow. The
// x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
