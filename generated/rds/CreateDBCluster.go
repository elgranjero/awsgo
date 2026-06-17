package rds

// CreateDBCluster is generated as a reference stub.
// Executable command wiring lives under cmd/rds.go.
//
// Creates a new Amazon Aurora DB cluster or Multi-AZ DB cluster.
//
// If you create an Aurora DB cluster, the request creates an empty cluster. You
// must explicitly create the writer instance for your DB cluster using the [CreateDBInstance]
// operation. If you create a Multi-AZ DB cluster, the request creates a writer and
// two reader DB instances for you, each in a different Availability Zone.
//
// You can use the ReplicationSourceIdentifier parameter to create an Amazon
// Aurora DB cluster as a read replica of another DB cluster or Amazon RDS for
// MySQL or PostgreSQL DB instance. For more information about Amazon Aurora, see [What is Amazon Aurora?]
// in the Amazon Aurora User Guide.
//
// You can also use the ReplicationSourceIdentifier parameter to create a Multi-AZ
// DB cluster read replica with an RDS for MySQL or PostgreSQL DB instance as the
// source. For more information about Multi-AZ DB clusters, see [Multi-AZ DB cluster deployments]in the Amazon RDS
// User Guide.
//
// [CreateDBInstance]: https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html
// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
// [Multi-AZ DB cluster deployments]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html
