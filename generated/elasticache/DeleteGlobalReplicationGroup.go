package elasticache

// DeleteGlobalReplicationGroup is generated as a reference stub.
// Executable command wiring lives under cmd/elasticache.go.
//
// Deleting a Global datastore is a two-step process:
//
// - First, you must DisassociateGlobalReplicationGroupto remove the secondary clusters in the Global datastore.
//
// - Once the Global datastore contains only the primary cluster, you can use
// the DeleteGlobalReplicationGroup API to delete the Global datastore while
// retainining the primary cluster using RetainPrimaryReplicationGroup=true .
//
// Since the Global Datastore has only a primary cluster, you can delete the
// Global Datastore while retaining the primary by setting
// RetainPrimaryReplicationGroup=true . The primary cluster is never deleted when
// deleting a Global Datastore. It can only be deleted when it no longer is
// associated with any Global Datastore.
//
// When you receive a successful response from this operation, Amazon ElastiCache
// immediately begins deleting the selected resources; you cannot cancel or revert
// this operation.
