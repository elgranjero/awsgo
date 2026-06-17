package kinesis

// SubscribeToShard is generated as a reference stub.
// Executable command wiring lives under cmd/kinesis.go.
//
// This operation establishes an HTTP/2 connection between the consumer you
// specify in the ConsumerARN parameter and the shard you specify in the ShardId
// parameter. After the connection is successfully established, Kinesis Data
// Streams pushes records from the shard to the consumer over this connection.
// Before you call this operation, call RegisterStreamConsumerto register the consumer with Kinesis Data
// Streams.
//
// When the SubscribeToShard call succeeds, your consumer starts receiving events
// of type SubscribeToShardEventover the HTTP/2 connection for up to 5 minutes, after which time you
// need to call SubscribeToShard again to renew the subscription if you want to
// continue to receive records.
//
// You can make one call to SubscribeToShard per second per registered consumer
// per shard. For example, if you have a 4000 shard stream and two registered
// stream consumers, you can make one SubscribeToShard request per second for each
// combination of shard and registered consumer, allowing you to subscribe both
// consumers to all 4000 shards in one second.
//
// If you call SubscribeToShard again with the same ConsumerARN and ShardId within
// 5 seconds of a successful call, you'll get a ResourceInUseException . If you
// call SubscribeToShard 5 seconds or more after a successful call, the second
// call takes over the subscription and the previous connection expires or fails
// with a ResourceInUseException .
//
// For an example of how to use this operation, see [Enhanced Fan-Out Using the Kinesis Data Streams API].
//
// [Enhanced Fan-Out Using the Kinesis Data Streams API]: https://docs.aws.amazon.com/streams/latest/dev/building-enhanced-consumers-api.html
