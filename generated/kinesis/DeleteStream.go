package kinesis

// DeleteStream is generated as a reference stub.
// Executable command wiring lives under cmd/kinesis.go.
//
// Deletes a Kinesis data stream and all its shards and data. You must shut down
// any applications that are operating on the stream before you delete the stream.
// If an application attempts to operate on a deleted stream, it receives the
// exception ResourceNotFoundException .
//
// When invoking this API, you must use either the StreamARN or the StreamName
// parameter, or both. It is recommended that you use the StreamARN input
// parameter when you invoke this API.
//
// If the stream is in the ACTIVE state, you can delete it. After a DeleteStream
// request, the specified stream is in the DELETING state until Kinesis Data
// Streams completes the deletion.
//
// Note: Kinesis Data Streams might continue to accept data read and write
// operations, such as PutRecord, PutRecords, and GetRecords, on a stream in the DELETING state until the
// stream deletion is complete.
//
// When you delete a stream, any shards in that stream are also deleted, and any
// tags are dissociated from the stream.
//
// You can use the DescribeStreamSummary operation to check the state of the stream, which is returned
// in StreamStatus .
//
// DeleteStreamhas a limit of five transactions per second per account.
