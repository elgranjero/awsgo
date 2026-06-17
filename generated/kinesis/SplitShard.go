package kinesis

// SplitShard is generated as a reference stub.
// Executable command wiring lives under cmd/kinesis.go.
//
// Splits a shard into two new shards in the Kinesis data stream, to increase the
// stream's capacity to ingest and transport data. SplitShard is called when there
// is a need to increase the overall capacity of a stream because of an expected
// increase in the volume of data records being ingested. This API is only
// supported for the data streams with the provisioned capacity mode.
//
// When invoking this API, you must use either the StreamARN or the StreamName
// parameter, or both. It is recommended that you use the StreamARN input
// parameter when you invoke this API.
//
// You can also use SplitShard when a shard appears to be approaching its maximum
// utilization; for example, the producers sending data into the specific shard are
// suddenly sending more than previously anticipated. You can also call SplitShard
// to increase stream capacity, so that more Kinesis Data Streams applications can
// simultaneously read data from the stream for real-time processing.
//
// You must specify the shard to be split and the new hash key, which is the
// position in the shard where the shard gets split in two. In many cases, the new
// hash key might be the average of the beginning and ending hash key, but it can
// be any hash key value in the range being mapped into the shard. For more
// information, see [Split a Shard]in the Amazon Kinesis Data Streams Developer Guide.
//
// You can use DescribeStreamSummary and the ListShards APIs to determine the shard ID and hash key values for
// the ShardToSplit and NewStartingHashKey parameters that are specified in the
// SplitShard request.
//
// SplitShard is an asynchronous operation. Upon receiving a SplitShard request,
// Kinesis Data Streams immediately returns a response and sets the stream status
// to UPDATING . After the operation is completed, Kinesis Data Streams sets the
// stream status to ACTIVE . Read and write operations continue to work while the
// stream is in the UPDATING state.
//
// You can use DescribeStreamSummary to check the status of the stream, which is returned in
// StreamStatus . If the stream is in the ACTIVE state, you can call SplitShard .
//
// If the specified stream does not exist, DescribeStreamSummary returns a ResourceNotFoundException .
// If you try to create more shards than are authorized for your account, you
// receive a LimitExceededException .
//
// For the default shard limit for an Amazon Web Services account, see [Kinesis Data Streams Limits] in the
// Amazon Kinesis Data Streams Developer Guide. To increase this limit, [contact Amazon Web Services Support].
//
// If you try to operate on too many streams simultaneously using CreateStream, DeleteStream, MergeShards, and/or SplitShard,
// you receive a LimitExceededException .
//
// SplitShard has a limit of five transactions per second per account.
//
// [Kinesis Data Streams Limits]: https://docs.aws.amazon.com/kinesis/latest/dev/service-sizes-and-limits.html
// [contact Amazon Web Services Support]: https://docs.aws.amazon.com/general/latest/gr/aws_service_limits.html
// [Split a Shard]: https://docs.aws.amazon.com/kinesis/latest/dev/kinesis-using-sdk-java-resharding-split.html
