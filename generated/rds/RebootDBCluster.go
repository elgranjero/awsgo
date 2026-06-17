package rds

// RebootDBCluster is generated as a reference stub.
// Executable command wiring lives under cmd/rds.go.
//
// You might need to reboot your DB cluster, usually for maintenance reasons. For
// example, if you make certain modifications, or if you change the DB cluster
// parameter group associated with the DB cluster, reboot the DB cluster for the
// changes to take effect.
//
// Rebooting a DB cluster restarts the database engine service. Rebooting a DB
// cluster results in a momentary outage, during which the DB cluster status is set
// to rebooting.
//
// Use this operation only for a non-Aurora Multi-AZ DB cluster.
//
// For more information on Multi-AZ DB clusters, see [Multi-AZ DB cluster deployments] in the Amazon RDS User
// Guide.
//
// [Multi-AZ DB cluster deployments]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html
