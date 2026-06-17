package neptune

// SwitchoverGlobalCluster is generated as a reference stub.
// Executable command wiring lives under cmd/neptune.go.
//
// Switches over the specified secondary DB cluster to be the new primary DB
// cluster in the global database cluster. Switchover operations were previously
// called "managed planned failovers."
//
// Promotes the specified secondary cluster to assume full read/write capabilities
// and demotes the current primary cluster to a secondary (read-only) cluster,
// maintaining the original replication topology. All secondary clusters are
// synchronized with the primary at the beginning of the process so the new primary
// continues operations for the global database without losing any data. Your
// database is unavailable for a short time while the primary and selected
// secondary clusters are assuming their new roles.
//
// This operation is intended for controlled environments, for operations such as
// "regional rotation" or to fall back to the original primary after a global
// database failover.
