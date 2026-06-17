package sns

// CreateTopic is generated as a reference stub.
// Executable command wiring lives under cmd/sns.go.
//
// Creates a topic to which notifications can be published. Users can create at
// most 100,000 standard topics (at most 1,000 FIFO topics). For more information,
// see [Creating an Amazon SNS topic]in the Amazon SNS Developer Guide. This action is idempotent, so if the
// requester already owns a topic with the specified name, that topic's ARN is
// returned without creating a new topic.
//
// [Creating an Amazon SNS topic]: https://docs.aws.amazon.com/sns/latest/dg/sns-create-topic.html
