package ivs

// CreateStreamKey is generated as a reference stub.
// Executable command wiring lives under cmd/ivs.go.
//
// Creates a stream key, used to initiate a stream, for the specified channel ARN.
//
// Note that CreateChannel creates a stream key. If you subsequently use CreateStreamKey on the
// same channel, it will fail because a stream key already exists and there is a
// limit of 1 stream key per channel. To reset the stream key on a channel, use DeleteStreamKey
// and then CreateStreamKey.
