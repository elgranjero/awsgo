package kinesis

// DescribeStream is generated as a reference stub.
// Executable command wiring lives under cmd/kinesis.go.
//
// Describes the specified Kinesis data stream.
//
// This API has been revised. It's highly recommended that you use the DescribeStreamSummary API to get
// a summarized description of the specified Kinesis data stream and the ListShardsAPI to
// list the shards in a specified data stream and obtain information about each
// shard.
//
// When invoking this API, you must use either the StreamARN or the StreamName
// parameter, or both. It is recommended that you use the StreamARN input
// parameter when you invoke this API.
//
// The information returned includes the stream name, Amazon Resource Name (ARN),
// creation time, enhanced metric configuration, and shard map. The shard map is an
// array of shard objects. For each shard object, there is the hash key and
// sequence number ranges that the shard spans, and the IDs of any earlier shards
// that played in a role in creating the shard. Every record ingested in the stream
// is identified by a sequence number, which is assigned when the record is put
// into the stream.
//
// You can limit the number of shards returned by each call. For more information,
// see [Retrieving Shards from a Stream]in the Amazon Kinesis Data Streams Developer Guide.
//
// There are no guarantees about the chronological order shards returned. To
// process shards in chronological order, use the ID of the parent shard to track
// the lineage to the oldest shard.
//
// This operation has a limit of 10 transactions per second per account.
//
// [Retrieving Shards from a Stream]: https://docs.aws.amazon.com/kinesis/latest/dev/kinesis-using-sdk-java-retrieve-shards.html
