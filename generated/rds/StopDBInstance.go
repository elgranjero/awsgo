package rds

// StopDBInstance is generated as a reference stub.
// Executable command wiring lives under cmd/rds.go.
//
// Stops an Amazon RDS DB instance temporarily. When you stop a DB instance,
// Amazon RDS retains the DB instance's metadata, including its endpoint, DB
// parameter group, and option group membership. Amazon RDS also retains the
// transaction logs so you can do a point-in-time restore if necessary. The
// instance restarts automatically after 7 days.
//
// For more information, see [Stopping an Amazon RDS DB Instance Temporarily] in the Amazon RDS User Guide.
//
// This command doesn't apply to RDS Custom, Aurora MySQL, and Aurora PostgreSQL.
// For Aurora clusters, use StopDBCluster instead.
//
// [Stopping an Amazon RDS DB Instance Temporarily]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_StopInstance.html
