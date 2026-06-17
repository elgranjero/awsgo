package kinesis

// UpdateStreamWarmThroughput is generated as a reference stub.
// Executable command wiring lives under cmd/kinesis.go.
//
// Updates the warm throughput configuration for the specified Amazon Kinesis Data
// Streams on-demand data stream. This operation allows you to proactively scale
// your on-demand data stream to a specified throughput level, enabling better
// performance for sudden traffic spikes.
//
// When invoking this API, you must use either the StreamARN or the StreamName
// parameter, or both. It is recommended that you use the StreamARN input
// parameter when you invoke this API.
//
// Updating the warm throughput is an asynchronous operation. Upon receiving the
// request, Kinesis Data Streams returns immediately and sets the status of the
// stream to UPDATING . After the update is complete, Kinesis Data Streams sets the
// status of the stream back to ACTIVE . Depending on the size of the stream, the
// scaling action could take a few minutes to complete. You can continue to read
// and write data to your stream while its status is UPDATING .
//
// This operation is only supported for data streams with the on-demand capacity
// mode in accounts that have MinimumThroughputBillingCommitment enabled.
// Provisioned capacity mode streams do not support warm throughput configuration.
//
// This operation has the following default limits. By default, you cannot do the
// following:
//
// - Scale to more than 10 GiBps for an on-demand stream.
//
// - This API has a call limit of 5 transactions per second (TPS) for each
// Amazon Web Services account. TPS over 5 will initiate the
// LimitExceededException .
//
// For the default limits for an Amazon Web Services account, see [Streams Limits] in the Amazon
// Kinesis Data Streams Developer Guide. To request an increase in the call rate
// limit, the shard limit for this API, or your overall shard limit, use the [limits form].
//
// [limits form]: https://console.aws.amazon.com/support/v1#/case/create?issueType=service-limit-increase&limitType=service-code-kinesis
// [Streams Limits]: https://docs.aws.amazon.com/kinesis/latest/dev/service-sizes-and-limits.html
