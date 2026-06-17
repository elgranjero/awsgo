package rds

// SwitchoverGlobalCluster is generated as a reference stub.
// Executable command wiring lives under cmd/rds.go.
//
// Switches over the specified secondary DB cluster to be the new primary DB
// cluster in the global database cluster. Switchover operations were previously
// called "managed planned failovers."
//
// Aurora promotes the specified secondary cluster to assume full read/write
// capabilities and demotes the current primary cluster to a secondary (read-only)
// cluster, maintaining the orginal replication topology. All secondary clusters
// are synchronized with the primary at the beginning of the process so the new
// primary continues operations for the Aurora global database without losing any
// data. Your database is unavailable for a short time while the primary and
// selected secondary clusters are assuming their new roles. For more information
// about switching over an Aurora global database, see [Performing switchovers for Amazon Aurora global databases]in the Amazon Aurora User
// Guide.
//
// This operation is intended for controlled environments, for operations such as
// "regional rotation" or to fall back to the original primary after a global
// database failover.
//
// [Performing switchovers for Amazon Aurora global databases]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database-disaster-recovery.html#aurora-global-database-disaster-recovery.managed-failover
