package rds

// CopyDBClusterSnapshot is generated as a reference stub.
// Executable command wiring lives under cmd/rds.go.
//
// Copies a snapshot of a DB cluster.
//
// To copy a DB cluster snapshot from a shared manual DB cluster snapshot,
// SourceDBClusterSnapshotIdentifier must be the Amazon Resource Name (ARN) of the
// shared DB cluster snapshot.
//
// You can copy an encrypted DB cluster snapshot from another Amazon Web Services
// Region. In that case, the Amazon Web Services Region where you call the
// CopyDBClusterSnapshot operation is the destination Amazon Web Services Region
// for the encrypted DB cluster snapshot to be copied to. To copy an encrypted DB
// cluster snapshot from another Amazon Web Services Region, you must provide the
// following values:
//
// - KmsKeyId - The Amazon Web Services Key Management System (Amazon Web
// Services KMS) key identifier for the key to use to encrypt the copy of the DB
// cluster snapshot in the destination Amazon Web Services Region.
//
// - TargetDBClusterSnapshotIdentifier - The identifier for the new copy of the
// DB cluster snapshot in the destination Amazon Web Services Region.
//
// - SourceDBClusterSnapshotIdentifier - The DB cluster snapshot identifier for
// the encrypted DB cluster snapshot to be copied. This identifier must be in the
// ARN format for the source Amazon Web Services Region and is the same value as
// the SourceDBClusterSnapshotIdentifier in the presigned URL.
//
// To cancel the copy operation once it is in progress, delete the target DB
// cluster snapshot identified by TargetDBClusterSnapshotIdentifier while that DB
// cluster snapshot is in "copying" status.
//
// For more information on copying encrypted Amazon Aurora DB cluster snapshots
// from one Amazon Web Services Region to another, see [Copying a Snapshot]in the Amazon Aurora User
// Guide.
//
// For more information on Amazon Aurora DB clusters, see [What is Amazon Aurora?] in the Amazon Aurora
// User Guide.
//
// For more information on Multi-AZ DB clusters, see [Multi-AZ DB cluster deployments] in the Amazon RDS User Guide.
//
// [Copying a Snapshot]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_CopySnapshot.html
// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
// [Multi-AZ DB cluster deployments]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html
