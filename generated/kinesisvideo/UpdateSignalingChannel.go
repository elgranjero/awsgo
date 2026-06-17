package kinesisvideo

// UpdateSignalingChannel is generated as a reference stub.
// Executable command wiring lives under cmd/kinesisvideo.go.
//
// Updates the existing signaling channel. This is an asynchronous operation and
// takes time to complete.
//
// If the MessageTtlSeconds value is updated (either increased or reduced), it
// only applies to new messages sent via this channel after it's been updated.
// Existing messages are still expired as per the previous MessageTtlSeconds value.
