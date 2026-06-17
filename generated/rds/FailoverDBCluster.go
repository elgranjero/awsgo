package rds

// FailoverDBCluster is generated as a reference stub.
// Executable command wiring lives under cmd/rds.go.
//
// Forces a failover for a DB cluster.
//
// For an Aurora DB cluster, failover for a DB cluster promotes one of the Aurora
// Replicas (read-only instances) in the DB cluster to be the primary DB instance
// (the cluster writer).
//
// For a Multi-AZ DB cluster, after RDS terminates the primary DB instance, the
// internal monitoring system detects that the primary DB instance is unhealthy and
// promotes a readable standby (read-only instances) in the DB cluster to be the
// primary DB instance (the cluster writer). Failover times are typically less than
// 35 seconds.
//
// An Amazon Aurora DB cluster automatically fails over to an Aurora Replica, if
// one exists, when the primary DB instance fails. A Multi-AZ DB cluster
// automatically fails over to a readable standby DB instance when the primary DB
// instance fails.
//
// To simulate a failure of a primary instance for testing, you can force a
// failover. Because each instance in a DB cluster has its own endpoint address,
// make sure to clean up and re-establish any existing connections that use those
// endpoint addresses when the failover is complete.
//
// For more information on Amazon Aurora DB clusters, see [What is Amazon Aurora?] in the Amazon Aurora
// User Guide.
//
// For more information on Multi-AZ DB clusters, see [Multi-AZ DB cluster deployments] in the Amazon RDS User Guide.
//
// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
// [Multi-AZ DB cluster deployments]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html
