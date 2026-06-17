package rds

// DeleteDBCluster is generated as a reference stub.
// Executable command wiring lives under cmd/rds.go.
//
// The DeleteDBCluster action deletes a previously provisioned DB cluster. When
// you delete a DB cluster, all automated backups for that DB cluster are deleted
// and can't be recovered. Manual DB cluster snapshots of the specified DB cluster
// are not deleted.
//
// If you're deleting a Multi-AZ DB cluster with read replicas, all cluster
// members are terminated and read replicas are promoted to standalone instances.
//
// For more information on Amazon Aurora, see [What is Amazon Aurora?] in the Amazon Aurora User Guide.
//
// For more information on Multi-AZ DB clusters, see [Multi-AZ DB cluster deployments] in the Amazon RDS User Guide.
//
// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
// [Multi-AZ DB cluster deployments]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html
