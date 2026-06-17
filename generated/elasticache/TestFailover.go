package elasticache

// TestFailover is generated as a reference stub.
// Executable command wiring lives under cmd/elasticache.go.
//
// Represents the input of a TestFailover operation which tests automatic failover
// on a specified node group (called shard in the console) in a replication group
// (called cluster in the console).
//
// This API is designed for testing the behavior of your application in case of
// ElastiCache failover. It is not designed to be an operational tool for
// initiating a failover to overcome a problem you may have with the cluster.
// Moreover, in certain conditions such as large-scale operational events, Amazon
// may block this API.
//
// Note the following
//
// - A customer can use this operation to test automatic failover on up to 15
// shards (called node groups in the ElastiCache API and Amazon CLI) in any rolling
// 24-hour period.
//
// - If calling this operation on shards in different clusters (called
// replication groups in the API and CLI), the calls can be made concurrently.
//
// - If calling this operation multiple times on different shards in the same
// Valkey or Redis OSS (cluster mode enabled) replication group, the first node
// replacement must complete before a subsequent call can be made.
//
// - To determine whether the node replacement is complete you can check Events
// using the Amazon ElastiCache console, the Amazon CLI, or the ElastiCache API.
// Look for the following automatic failover related events, listed here in order
// of occurrance:
//
// - Replication group message: Test Failover API called for node group
//
// - Cache cluster message: Failover from primary node to replica node completed
//
// - Replication group message: Failover from primary node to replica node
// completed
//
// - Cache cluster message: Recovering cache nodes
//
// - Cache cluster message: Finished recovery for cache nodes
//
// For more information see:
//
// [Viewing ElastiCache Events]
// - in the ElastiCache User Guide
//
// [DescribeEvents]
// - in the ElastiCache API Reference
//
// Also see, [Testing Multi-AZ] in the ElastiCache User Guide.
//
// [DescribeEvents]: https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_DescribeEvents.html
// [Testing Multi-AZ]: https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/AutoFailover.html#auto-failover-test
// [Viewing ElastiCache Events]: https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/ECEvents.Viewing.html
