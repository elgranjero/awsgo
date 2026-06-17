package kinesis

// RegisterStreamConsumer is generated as a reference stub.
// Executable command wiring lives under cmd/kinesis.go.
//
// Registers a consumer with a Kinesis data stream. When you use this operation,
// the consumer you register can then call SubscribeToShardto receive data from the stream using
// enhanced fan-out, at a rate of up to 2 MiB per second for every shard you
// subscribe to. This rate is unaffected by the total number of consumers that read
// from the same stream.
//
// You can add tags to the registered consumer when making a RegisterStreamConsumer
// request by setting the Tags parameter. If you pass the Tags parameter, in
// addition to having the kinesis:RegisterStreamConsumer permission, you must also
// have the kinesis:TagResource permission for the consumer that will be
// registered. Tags will take effect from the CREATING status of the consumer.
//
// With On-demand Advantage streams, you can register up to 50 consumers per
// stream to use Enhanced Fan-out. With On-demand Standard and Provisioned streams,
// you can register up to 20 consumers per stream to use Enhanced Fan-out. A given
// consumer can only be registered with one stream at a time.
//
// For an example of how to use this operation, see [Enhanced Fan-Out Using the Kinesis Data Streams API].
//
// The use of this operation has a limit of five transactions per second per
// account. Also, only 5 consumers can be created simultaneously. In other words,
// you cannot have more than 5 consumers in a CREATING status at the same time.
// Registering a 6th consumer while there are 5 in a CREATING status results in a
// LimitExceededException .
//
// [Enhanced Fan-Out Using the Kinesis Data Streams API]: https://docs.aws.amazon.com/streams/latest/dev/building-enhanced-consumers-api.html
