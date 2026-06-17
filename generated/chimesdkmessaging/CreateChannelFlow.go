package chimesdkmessaging

// CreateChannelFlow is generated as a reference stub.
// Executable command wiring lives under cmd/chimesdkmessaging.go.
//
// Creates a channel flow, a container for processors. Processors are AWS Lambda
// functions that perform actions on chat messages, such as stripping out
// profanity. You can associate channel flows with channels, and the processors in
// the channel flow then take action on all messages sent to that channel. This is
// a developer API.
//
// Channel flows process the following items:
//
// - New and updated messages
//
// - Persistent and non-persistent messages
//
// - The Standard message type
//
// Channel flows don't process Control or System messages. For more information
// about the message types provided by Chime SDK messaging, refer to [Message types]in the Amazon
// Chime developer guide.
//
// [Message types]: https://docs.aws.amazon.com/chime-sdk/latest/dg/using-the-messaging-sdk.html#msg-types
