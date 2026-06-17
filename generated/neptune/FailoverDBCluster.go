package neptune

// FailoverDBCluster is generated as a reference stub.
// Executable command wiring lives under cmd/neptune.go.
//
// Forces a failover for a DB cluster.
//
// A failover for a DB cluster promotes one of the Read Replicas (read-only
// instances) in the DB cluster to be the primary instance (the cluster writer).
//
// Amazon Neptune will automatically fail over to a Read Replica, if one exists,
// when the primary instance fails. You can force a failover when you want to
// simulate a failure of a primary instance for testing. Because each instance in a
// DB cluster has its own endpoint address, you will need to clean up and
// re-establish any existing connections that use those endpoint addresses when the
// failover is complete.
