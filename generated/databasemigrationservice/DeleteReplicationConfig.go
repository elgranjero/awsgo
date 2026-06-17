package databasemigrationservice

// DeleteReplicationConfig is generated as a reference stub.
// Executable command wiring lives under cmd/databasemigrationservice.go.
//
// Deletes an DMS Serverless replication configuration. This effectively
// deprovisions any and all replications that use this configuration. You can't
// delete the configuration for an DMS Serverless replication that is ongoing. You
// can delete the configuration when the replication is in a non-RUNNING and
// non-STARTING state.
