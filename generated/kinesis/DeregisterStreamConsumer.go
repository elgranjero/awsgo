package kinesis

// DeregisterStreamConsumer is generated as a reference stub.
// Executable command wiring lives under cmd/kinesis.go.
//
// To deregister a consumer, provide its ARN. Alternatively, you can provide the
// ARN of the data stream and the name you gave the consumer when you registered
// it. You may also provide all three parameters, as long as they don't conflict
// with each other. If you don't know the name or ARN of the consumer that you want
// to deregister, you can use the ListStreamConsumersoperation to get a list of the descriptions of
// all the consumers that are currently registered with a given data stream. The
// description of a consumer contains its name and ARN.
//
// This operation has a limit of five transactions per second per stream.
