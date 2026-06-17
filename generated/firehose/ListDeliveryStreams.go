package firehose

// ListDeliveryStreams is generated as a reference stub.
// Executable command wiring lives under cmd/firehose.go.
//
// Lists your Firehose streams in alphabetical order of their names.
//
// The number of Firehose streams might be too large to return using a single call
// to ListDeliveryStreams . You can limit the number of Firehose streams returned,
// using the Limit parameter. To determine whether there are more delivery streams
// to list, check the value of HasMoreDeliveryStreams in the output. If there are
// more Firehose streams to list, you can request them by calling this operation
// again and setting the ExclusiveStartDeliveryStreamName parameter to the name of
// the last Firehose stream returned in the last call.
