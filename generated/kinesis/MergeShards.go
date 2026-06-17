package kinesis

// MergeShards is generated as a reference stub.
// Executable command wiring lives under cmd/kinesis.go.
//
// Merges two adjacent shards in a Kinesis data stream and combines them into a
// single shard to reduce the stream's capacity to ingest and transport data. This
// API is only supported for the data streams with the provisioned capacity mode.
// Two shards are considered adjacent if the union of the hash key ranges for the
// two shards form a contiguous set with no gaps. For example, if you have two
// shards, one with a hash key range of 276...381 and the other with a hash key
// range of 382...454, then you could merge these two shards into a single shard
// that would have a hash key range of 276...454. After the merge, the single child
// shard receives data for all hash key values covered by the two parent shards.
//
// When invoking this API, you must use either the StreamARN or the StreamName
// parameter, or both. It is recommended that you use the StreamARN input
// parameter when you invoke this API.
//
// MergeShards is called when there is a need to reduce the overall capacity of a
// stream because of excess capacity that is not being used. You must specify the
// shard to be merged and the adjacent shard for a stream. For more information
// about merging shards, see [Merge Two Shards]in the Amazon Kinesis Data Streams Developer Guide.
//
// If the stream is in the ACTIVE state, you can call MergeShards . If a stream is
// in the CREATING , UPDATING , or DELETING state, MergeShards returns a
// ResourceInUseException . If the specified stream does not exist, MergeShards
// returns a ResourceNotFoundException .
//
// You can use DescribeStreamSummary to check the state of the stream, which is returned in StreamStatus
// .
//
// MergeShards is an asynchronous operation. Upon receiving a MergeShards request,
// Amazon Kinesis Data Streams immediately returns a response and sets the
// StreamStatus to UPDATING . After the operation is completed, Kinesis Data
// Streams sets the StreamStatus to ACTIVE . Read and write operations continue to
// work while the stream is in the UPDATING state.
//
// You use DescribeStreamSummary and the ListShards APIs to determine the shard IDs that are specified in the
// MergeShards request.
//
// If you try to operate on too many streams in parallel using CreateStream, DeleteStream, MergeShards , or SplitShard
// , you receive a LimitExceededException .
//
// MergeShards has a limit of five transactions per second per account.
//
// [Merge Two Shards]: https://docs.aws.amazon.com/kinesis/latest/dev/kinesis-using-sdk-java-resharding-merge.html
