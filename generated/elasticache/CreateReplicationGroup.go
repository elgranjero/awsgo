package elasticache

// CreateReplicationGroup is generated as a reference stub.
// Executable command wiring lives under cmd/elasticache.go.
//
// Creates a Valkey or Redis OSS (cluster mode disabled) or a Valkey or Redis OSS
// (cluster mode enabled) replication group.
//
// This API can be used to create a standalone regional replication group or a
// secondary replication group associated with a Global datastore.
//
// A Valkey or Redis OSS (cluster mode disabled) replication group is a collection
// of nodes, where one of the nodes is a read/write primary and the others are
// read-only replicas. Writes to the primary are asynchronously propagated to the
// replicas.
//
// A Valkey or Redis OSS cluster-mode enabled cluster is comprised of from 1 to 90
// shards (API/CLI: node groups). Each shard has a primary node and up to 5
// read-only replica nodes. The configuration can range from 90 shards and 0
// replicas to 15 shards and 5 replicas, which is the maximum number or replicas
// allowed.
//
// The node or shard limit can be increased to a maximum of 500 per cluster if the
// Valkey or Redis OSS engine version is 5.0.6 or higher. For example, you can
// choose to configure a 500 node cluster that ranges between 83 shards (one
// primary and 5 replicas per shard) and 500 shards (single primary and no
// replicas). Make sure there are enough available IP addresses to accommodate the
// increase. Common pitfalls include the subnets in the subnet group have too small
// a CIDR range or the subnets are shared and heavily used by other clusters. For
// more information, see [Creating a Subnet Group]. For versions below 5.0.6, the limit is 250 per cluster.
//
// To request a limit increase, see [Amazon Service Limits] and choose the limit type Nodes per cluster
// per instance type.
//
// When a Valkey or Redis OSS (cluster mode disabled) replication group has been
// successfully created, you can add one or more read replicas to it, up to a total
// of 5 read replicas. If you need to increase or decrease the number of node
// groups (console: shards), you can use scaling. For more information, see [Scaling self-designed clusters]in the
// ElastiCache User Guide.
//
// This operation is valid for Valkey and Redis OSS only.
//
// [Scaling self-designed clusters]: https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/Scaling.html
// [Creating a Subnet Group]: https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/SubnetGroups.Creating.html
// [Amazon Service Limits]: https://docs.aws.amazon.com/general/latest/gr/aws_service_limits.html
