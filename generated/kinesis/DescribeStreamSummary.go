package kinesis

// DescribeStreamSummary is generated as a reference stub.
// Executable command wiring lives under cmd/kinesis.go.
//
// Provides a summarized description of the specified Kinesis data stream without
// the shard list.
//
// When invoking this API, you must use either the StreamARN or the StreamName
// parameter, or both. It is recommended that you use the StreamARN input
// parameter when you invoke this API.
//
// The information returned includes the stream name, Amazon Resource Name (ARN),
// status, record retention period, approximate creation time, monitoring,
// encryption details, and open shard count.
//
// DescribeStreamSummaryhas a limit of 20 transactions per second per account.
