package dynamodbstreams

// DescribeStream is generated as a reference stub.
// Executable command wiring lives under cmd/dynamodbstreams.go.
//
// Returns information about a stream, including the current status of the stream,
// its Amazon Resource Name (ARN), the composition of its shards, and its
// corresponding DynamoDB table.
//
// You can call DescribeStream at a maximum rate of 10 times per second.
//
// Each shard in the stream has a SequenceNumberRange associated with it. If the
// SequenceNumberRange has a StartingSequenceNumber but no EndingSequenceNumber ,
// then the shard is still open (able to receive more stream records). If both
// StartingSequenceNumber and EndingSequenceNumber are present, then that shard is
// closed and can no longer receive more data.
