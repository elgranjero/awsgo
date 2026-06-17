package elasticache

// DeleteReplicationGroup is generated as a reference stub.
// Executable command wiring lives under cmd/elasticache.go.
//
// Deletes an existing replication group. By default, this operation deletes the
// entire replication group, including the primary/primaries and all of the read
// replicas. If the replication group has only one primary, you can optionally
// delete only the read replicas, while retaining the primary by setting
// RetainPrimaryCluster=true .
//
// When you receive a successful response from this operation, Amazon ElastiCache
// immediately begins deleting the selected resources; you cannot cancel or revert
// this operation.
//
// - CreateSnapshot permission is required to create a final snapshot. Without
// this permission, the API call will fail with an Access Denied exception.
//
// - This operation is valid for Redis OSS only.
