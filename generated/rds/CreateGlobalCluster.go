package rds

// CreateGlobalCluster is generated as a reference stub.
// Executable command wiring lives under cmd/rds.go.
//
// Creates an Aurora global database spread across multiple Amazon Web Services
// Regions. The global database contains a single primary cluster with read-write
// capability, and a read-only secondary cluster that receives data from the
// primary cluster through high-speed replication performed by the Aurora storage
// subsystem.
//
// You can create a global database that is initially empty, and then create the
// primary and secondary DB clusters in the global database. Or you can specify an
// existing Aurora cluster during the create operation, and this cluster becomes
// the primary cluster of the global database.
//
// This operation applies only to Aurora DB clusters.
