package chimesdkmessaging

// RedactChannelMessage is generated as a reference stub.
// Executable command wiring lives under cmd/chimesdkmessaging.go.
//
// Redacts message content and metadata. The message exists in the back end, but
// the action returns null content, and the state shows as redacted.
//
// The x-amz-chime-bearer request header is mandatory. Use the ARN of the
// AppInstanceUser or AppInstanceBot that makes the API call as the value in the
// header.
