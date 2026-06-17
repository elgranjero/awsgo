package rds

// FailoverGlobalCluster is generated as a reference stub.
// Executable command wiring lives under cmd/rds.go.
//
// Promotes the specified secondary DB cluster to be the primary DB cluster in the
// global database cluster to fail over or switch over a global database.
// Switchover operations were previously called "managed planned failovers."
//
// Although this operation can be used either to fail over or to switch over a
// global database cluster, its intended use is for global database failover. To
// switch over a global database cluster, we recommend that you use the SwitchoverGlobalClusteroperation
// instead.
//
// How you use this operation depends on whether you are failing over or switching
// over your global database cluster:
//
// - Failing over - Specify the AllowDataLoss parameter and don't specify the
// Switchover parameter.
//
// - Switching over - Specify the Switchover parameter or omit it, but don't
// specify the AllowDataLoss parameter.
//
// # About failing over and switching over
//
// While failing over and switching over a global database cluster both change the
// primary DB cluster, you use these operations for different reasons:
//
// - Failing over - Use this operation to respond to an unplanned event, such as
// a Regional disaster in the primary Region. Failing over can result in a loss of
// write transaction data that wasn't replicated to the chosen secondary before the
// failover event occurred. However, the recovery process that promotes a DB
// instance on the chosen seconday DB cluster to be the primary writer DB instance
// guarantees that the data is in a transactionally consistent state.
//
// For more information about failing over an Amazon Aurora global database, see [Performing managed failovers for Aurora global databases]
//
// in the Amazon Aurora User Guide.
//
// - Switching over - Use this operation on a healthy global database cluster
// for planned events, such as Regional rotation or to fail back to the original
// primary DB cluster after a failover operation. With this operation, there is no
// data loss.
//
// For more information about switching over an Amazon Aurora global database, see [Performing switchovers for Aurora global databases]
//
// in the Amazon Aurora User Guide.
//
// [Performing managed failovers for Aurora global databases]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database-disaster-recovery.html#aurora-global-database-failover.managed-unplanned
// [Performing switchovers for Aurora global databases]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database-disaster-recovery.html#aurora-global-database-disaster-recovery.managed-failover
