package dynamodbstreams

// GetRecords is generated as a reference stub.
// Executable command wiring lives under cmd/dynamodbstreams.go.
//
// Retrieves the stream records from a given shard.
//
// Specify a shard iterator using the ShardIterator parameter. The shard iterator
// specifies the position in the shard from which you want to start reading stream
// records sequentially. If there are no stream records available in the portion of
// the shard that the iterator points to, GetRecords returns an empty list. Note
// that it might take multiple calls to get to a portion of the shard that contains
// stream records.
//
// GetRecords can retrieve a maximum of 1 MB of data or 1000 stream records,
// whichever comes first.
