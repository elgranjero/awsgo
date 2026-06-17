package docdb

// FailoverDBCluster is generated as a reference stub.
// Executable command wiring lives under cmd/docdb.go.
//
// Forces a failover for a cluster.
//
// A failover for a cluster promotes one of the Amazon DocumentDB replicas
// (read-only instances) in the cluster to be the primary instance (the cluster
// writer).
//
// If the primary instance fails, Amazon DocumentDB automatically fails over to an
// Amazon DocumentDB replica, if one exists. You can force a failover when you want
// to simulate a failure of a primary instance for testing.
