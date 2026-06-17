package dynamodbstreams

// GetShardIterator is generated as a reference stub.
// Executable command wiring lives under cmd/dynamodbstreams.go.
//
// Returns a shard iterator. A shard iterator provides information about how to
// retrieve the stream records from within a shard. Use the shard iterator in a
// subsequent GetRecords request to read the stream records from the shard.
//
// A shard iterator expires 15 minutes after it is returned to the requester.
