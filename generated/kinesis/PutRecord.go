package kinesis

// PutRecord is generated as a reference stub.
// Executable command wiring lives under cmd/kinesis.go.
//
// Writes a single data record into an Amazon Kinesis data stream. Call PutRecord
// to send data into the stream for real-time ingestion and subsequent processing,
// one record at a time. Each shard can support writes up to 1,000 records per
// second, up to a maximum data write total of 10 MiB per second.
//
// When invoking this API, you must use either the StreamARN or the StreamName
// parameter, or both. It is recommended that you use the StreamARN input
// parameter when you invoke this API.
//
// You must specify the name of the stream that captures, stores, and transports
// the data; a partition key; and the data blob itself.
//
// The data blob can be any type of data; for example, a segment from a log file,
// geographic/location data, website clickstream data, and so on.
//
// The partition key is used by Kinesis Data Streams to distribute data across
// shards. Kinesis Data Streams segregates the data records that belong to a stream
// into multiple shards, using the partition key associated with each data record
// to determine the shard to which a given data record belongs.
//
// Partition keys are Unicode strings, with a maximum length limit of 256
// characters for each key. An MD5 hash function is used to map partition keys to
// 128-bit integer values and to map associated data records to shards using the
// hash key ranges of the shards. You can override hashing the partition key to
// determine the shard by explicitly specifying a hash value using the
// ExplicitHashKey parameter. For more information, see [Adding Data to a Stream] in the Amazon Kinesis
// Data Streams Developer Guide.
//
// PutRecord returns the shard ID of where the data record was placed and the
// sequence number that was assigned to the data record.
//
// Sequence numbers increase over time and are specific to a shard within a
// stream, not across all shards within a stream. To guarantee strictly increasing
// ordering, write serially to a shard and use the SequenceNumberForOrdering
// parameter. For more information, see [Adding Data to a Stream]in the Amazon Kinesis Data Streams
// Developer Guide.
//
// After you write a record to a stream, you cannot modify that record or its
// order within the stream.
//
// If a PutRecord request cannot be processed because of insufficient provisioned
// throughput on the shard involved in the request, PutRecord throws
// ProvisionedThroughputExceededException .
//
// By default, data records are accessible for 24 hours from the time that they
// are added to a stream. You can use IncreaseStreamRetentionPeriodor DecreaseStreamRetentionPeriod to modify this retention period.
//
// [Adding Data to a Stream]: https://docs.aws.amazon.com/kinesis/latest/dev/developing-producers-with-sdk.html#kinesis-using-sdk-java-add-data-to-stream
