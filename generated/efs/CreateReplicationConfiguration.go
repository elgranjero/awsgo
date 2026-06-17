package efs

// CreateReplicationConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/efs.go.
//
// Creates a replication conﬁguration to either a new or existing EFS file system.
// For more information, see [Amazon EFS replication]in the Amazon EFS User Guide. The replication
// configuration specifies the following:
//
// - Source file system – The EFS file system that you want to replicate.
//
// - Destination file system – The destination file system to which the source
// file system is replicated. There can only be one destination file system in a
// replication configuration.
//
// A file system can be part of only one replication configuration.
//
// The destination parameters for the replication configuration depend on whether
//
// you are replicating to a new file system or to an existing file system, and if
// you are replicating across Amazon Web Services accounts. See DestinationToCreatefor more
// information.
//
// This operation requires permissions for the
// elasticfilesystem:CreateReplicationConfiguration action. Additionally, other
// permissions are required depending on how you are replicating file systems. For
// more information, see [Required permissions for replication]in the Amazon EFS User Guide.
//
// [Required permissions for replication]: https://docs.aws.amazon.com/efs/latest/ug/efs-replication.html#efs-replication-permissions
// [Amazon EFS replication]: https://docs.aws.amazon.com/efs/latest/ug/efs-replication.html
