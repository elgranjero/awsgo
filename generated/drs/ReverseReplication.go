package drs

// ReverseReplication is generated as a reference stub.
// Executable command wiring lives under cmd/drs.go.
//
// Start replication to origin / target region - applies only to protected
// instances that originated in EC2. For recovery instances on target region -
// starts replication back to origin region. For failback instances on origin
// region - starts replication to target region to re-protect them.
