package kinesis

// ListShards is generated as a reference stub.
// Executable command wiring lives under cmd/kinesis.go.
//
// Lists the shards in a stream and provides information about each shard. This
// operation has a limit of 1000 transactions per second per data stream.
//
// When invoking this API, you must use either the StreamARN or the StreamName
// parameter, or both. It is recommended that you use the StreamARN input
// parameter when you invoke this API.
//
// This action does not list expired shards. For information about expired shards,
// see [Data Routing, Data Persistence, and Shard State after a Reshard].
//
// This API is a new operation that is used by the Amazon Kinesis Client Library
// (KCL). If you have a fine-grained IAM policy that only allows specific
// operations, you must update your policy to allow calls to this API. For more
// information, see [Controlling Access to Amazon Kinesis Data Streams Resources Using IAM].
//
// [Data Routing, Data Persistence, and Shard State after a Reshard]: https://docs.aws.amazon.com/streams/latest/dev/kinesis-using-sdk-java-after-resharding.html#kinesis-using-sdk-java-resharding-data-routing
// [Controlling Access to Amazon Kinesis Data Streams Resources Using IAM]: https://docs.aws.amazon.com/streams/latest/dev/controlling-access.html
