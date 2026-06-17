package kinesis

// DescribeStreamConsumer is generated as a reference stub.
// Executable command wiring lives under cmd/kinesis.go.
//
// To get the description of a registered consumer, provide the ARN of the
// consumer. Alternatively, you can provide the ARN of the data stream and the name
// you gave the consumer when you registered it. You may also provide all three
// parameters, as long as they don't conflict with each other. If you don't know
// the name or ARN of the consumer that you want to describe, you can use the ListStreamConsumers
// operation to get a list of the descriptions of all the consumers that are
// currently registered with a given data stream.
//
// This operation has a limit of 20 transactions per second per stream.
//
// When making a cross-account call with DescribeStreamConsumer , make sure to
// provide the ARN of the consumer.
