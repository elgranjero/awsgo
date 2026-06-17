package docdb

// FailoverGlobalCluster is generated as a reference stub.
// Executable command wiring lives under cmd/docdb.go.
//
// Promotes the specified secondary DB cluster to be the primary DB cluster in the
// global cluster when failing over a global cluster occurs.
//
// Use this operation to respond to an unplanned event, such as a regional
// disaster in the primary region. Failing over can result in a loss of write
// transaction data that wasn't replicated to the chosen secondary before the
// failover event occurred. However, the recovery process that promotes a DB
// instance on the chosen seconday DB cluster to be the primary writer DB instance
// guarantees that the data is in a transactionally consistent state.
