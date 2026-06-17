package kinesis

// DeleteResourcePolicy is generated as a reference stub.
// Executable command wiring lives under cmd/kinesis.go.
//
// Delete a policy for the specified data stream or consumer. Request patterns can
// be one of the following:
//
// - Data stream pattern: arn:aws.*:kinesis:.*:\d{12}:.*stream/\S+
//
// - Consumer pattern:
// ^(arn):aws.*:kinesis:.*:\d{12}:.*stream\/[a-zA-Z0-9_.-]+\/consumer\/[a-zA-Z0-9_.-]+:[0-9]+
