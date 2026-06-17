package neptune

// CreateGlobalCluster is generated as a reference stub.
// Executable command wiring lives under cmd/neptune.go.
//
// Creates a Neptune global database spread across multiple Amazon Regions. The
// global database contains a single primary cluster with read-write capability,
// and read-only secondary clusters that receive data from the primary cluster
// through high-speed replication performed by the Neptune storage subsystem.
//
// You can create a global database that is initially empty, and then add a
// primary cluster and secondary clusters to it, or you can specify an existing
// Neptune cluster during the create operation to become the primary cluster of the
// global database.
