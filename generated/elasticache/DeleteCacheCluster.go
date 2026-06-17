package elasticache

// DeleteCacheCluster is generated as a reference stub.
// Executable command wiring lives under cmd/elasticache.go.
//
// Deletes a previously provisioned cluster. DeleteCacheCluster deletes all
// associated cache nodes, node endpoints and the cluster itself. When you receive
// a successful response from this operation, Amazon ElastiCache immediately begins
// deleting the cluster; you cannot cancel or revert this operation.
//
// This operation is not valid for:
//
// - Valkey or Redis OSS (cluster mode enabled) clusters
//
// - Valkey or Redis OSS (cluster mode disabled) clusters
//
// - A cluster that is the last read replica of a replication group
//
// - A cluster that is the primary node of a replication group
//
// - A node group (shard) that has Multi-AZ mode enabled
//
// - A cluster from a Valkey or Redis OSS (cluster mode enabled) replication
// group
//
// - A cluster that is not in the available state
