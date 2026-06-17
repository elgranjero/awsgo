package databasemigrationservice

// CreateReplicationSubnetGroup is generated as a reference stub.
// Executable command wiring lives under cmd/databasemigrationservice.go.
//
// Creates a replication subnet group given a list of the subnet IDs in a VPC.
//
// The VPC needs to have at least one subnet in at least two availability zones in
// the Amazon Web Services Region, otherwise the service will throw a
// ReplicationSubnetGroupDoesNotCoverEnoughAZs exception.
//
// If a replication subnet group exists in your Amazon Web Services account, the
// CreateReplicationSubnetGroup action returns the following error message: The
// Replication Subnet Group already exists. In this case, delete the existing
// replication subnet group. To do so, use the [DeleteReplicationSubnetGroup]action. Optionally, choose Subnet
// groups in the DMS console, then choose your subnet group. Next, choose Delete
// from Actions.
//
// [DeleteReplicationSubnetGroup]: https://docs.aws.amazon.com/en_us/dms/latest/APIReference/API_DeleteReplicationSubnetGroup.html
