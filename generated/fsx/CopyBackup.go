package fsx

// CopyBackup is generated as a reference stub.
// Executable command wiring lives under cmd/fsx.go.
//
// Copies an existing backup within the same Amazon Web Services account to
// another Amazon Web Services Region (cross-Region copy) or within the same Amazon
// Web Services Region (in-Region copy). You can have up to five backup copy
// requests in progress to a single destination Region per account.
//
// You can use cross-Region backup copies for cross-Region disaster recovery. You
// can periodically take backups and copy them to another Region so that in the
// event of a disaster in the primary Region, you can restore from backup and
// recover availability quickly in the other Region. You can make cross-Region
// copies only within your Amazon Web Services partition. A partition is a grouping
// of Regions. Amazon Web Services currently has three partitions: aws (Standard
// Regions), aws-cn (China Regions), and aws-us-gov (Amazon Web Services GovCloud
// [US] Regions).
//
// You can also use backup copies to clone your file dataset to another Region or
// within the same Region.
//
// You can use the SourceRegion parameter to specify the Amazon Web Services
// Region from which the backup will be copied. For example, if you make the call
// from the us-west-1 Region and want to copy a backup from the us-east-2 Region,
// you specify us-east-2 in the SourceRegion parameter to make a cross-Region
// copy. If you don't specify a Region, the backup copy is created in the same
// Region where the request is sent from (in-Region copy).
//
// For more information about creating backup copies, see [Copying backups] in the Amazon FSx for
// Windows User Guide, [Copying backups]in the Amazon FSx for Lustre User Guide, and [Copying backups] in the Amazon
// FSx for OpenZFS User Guide.
//
// [Copying backups]: https://docs.aws.amazon.com/fsx/latest/OpenZFSGuide/using-backups.html#copy-backups
