package kinesis

// PutResourcePolicy is generated as a reference stub.
// Executable command wiring lives under cmd/kinesis.go.
//
// Attaches a resource-based policy to a data stream or registered consumer. If
// you are using an identity other than the root user of the Amazon Web Services
// account that owns the resource, the calling identity must have the
// PutResourcePolicy permissions on the specified Kinesis Data Streams resource and
// belong to the owner's account in order to use this operation. If you don't have
// PutResourcePolicy permissions, Amazon Kinesis Data Streams returns a 403 Access
// Denied error . If you receive a ResourceNotFoundException , check to see if you
// passed a valid stream or consumer resource.
//
// Request patterns can be one of the following:
//
// - Data stream pattern: arn:aws.*:kinesis:.*:\d{12}:.*stream/\S+
//
// - Consumer pattern:
// ^(arn):aws.*:kinesis:.*:\d{12}:.*stream\/[a-zA-Z0-9_.-]+\/consumer\/[a-zA-Z0-9_.-]+:[0-9]+
//
// For more information, see [Controlling Access to Amazon Kinesis Data Streams Resources Using IAM].
//
// [Controlling Access to Amazon Kinesis Data Streams Resources Using IAM]: https://docs.aws.amazon.com/streams/latest/dev/controlling-access.html
