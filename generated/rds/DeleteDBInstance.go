package rds

// DeleteDBInstance is generated as a reference stub.
// Executable command wiring lives under cmd/rds.go.
//
// Deletes a previously provisioned DB instance. When you delete a DB instance,
// all automated backups for that instance are deleted and can't be recovered.
// However, manual DB snapshots of the DB instance aren't deleted.
//
// If you request a final DB snapshot, the status of the Amazon RDS DB instance is
// deleting until the DB snapshot is created. This operation can't be canceled or
// reverted after it begins. To monitor the status of this operation, use
// DescribeDBInstance .
//
// When a DB instance is in a failure state and has a status of failed ,
// incompatible-restore , or incompatible-network , you can only delete it when you
// skip creation of the final snapshot with the SkipFinalSnapshot parameter.
//
// If the specified DB instance is part of an Amazon Aurora DB cluster, you can't
// delete the DB instance if both of the following conditions are true:
//
// - The DB cluster is a read replica of another Amazon Aurora DB cluster.
//
// - The DB instance is the only instance in the DB cluster.
//
// To delete a DB instance in this case, first use the PromoteReadReplicaDBCluster
// operation to promote the DB cluster so that it's no longer a read replica. After
// the promotion completes, use the DeleteDBInstance operation to delete the final
// instance in the DB cluster.
//
// For RDS Custom DB instances, deleting the DB instance permanently deletes the
// EC2 instance and the associated EBS volumes. Make sure that you don't terminate
// or delete these resources before you delete the DB instance. Otherwise, deleting
// the DB instance and creation of the final snapshot might fail.
