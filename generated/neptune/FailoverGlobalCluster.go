package neptune

// FailoverGlobalCluster is generated as a reference stub.
// Executable command wiring lives under cmd/neptune.go.
//
// Initiates the failover process for a Neptune global database.
//
// A failover for a Neptune global database promotes one of secondary read-only DB
// clusters to be the primary DB cluster and demotes the primary DB cluster to
// being a secondary (read-only) DB cluster. In other words, the role of the
// current primary DB cluster and the selected target secondary DB cluster are
// switched. The selected secondary DB cluster assumes full read/write capabilities
// for the Neptune global database.
//
// This action applies only to Neptune global databases. This action is only
// intended for use on healthy Neptune global databases with healthy Neptune DB
// clusters and no region-wide outages, to test disaster recovery scenarios or to
// reconfigure the global database topology.
