package ec2

// DescribeSnapshots is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Describes the specified EBS snapshots available to you or all of the EBS
// snapshots available to you.
//
// The snapshots available to you include public snapshots, private snapshots that
// you own, and private snapshots owned by other Amazon Web Services accounts for
// which you have explicit create volume permissions.
//
// The create volume permissions fall into the following categories:
//
// - public: The owner of the snapshot granted create volume permissions for the
// snapshot to the all group. All Amazon Web Services accounts have create volume
// permissions for these snapshots.
//
// - explicit: The owner of the snapshot granted create volume permissions to a
// specific Amazon Web Services account.
//
// - implicit: An Amazon Web Services account has implicit create volume
// permissions for all snapshots it owns.
//
// The list of snapshots returned can be filtered by specifying snapshot IDs,
// snapshot owners, or Amazon Web Services accounts with create volume permissions.
// If no options are specified, Amazon EC2 returns all snapshots for which you have
// create volume permissions.
//
// If you specify one or more snapshot IDs, only snapshots that have the specified
// IDs are returned. If you specify an invalid snapshot ID, an error is returned.
// If you specify a snapshot ID for which you do not have access, it is not
// included in the returned results.
//
// If you specify one or more snapshot owners using the OwnerIds option, only
// snapshots from the specified owners and for which you have access are returned.
// The results can include the Amazon Web Services account IDs of the specified
// owners, amazon for snapshots owned by Amazon, or self for snapshots that you
// own.
//
// If you specify a list of restorable users, only snapshots with create snapshot
// permissions for those users are returned. You can specify Amazon Web Services
// account IDs (if you own the snapshots), self for snapshots for which you own or
// have explicit permissions, or all for public snapshots.
//
// If you are describing a long list of snapshots, we recommend that you paginate
// the output to make the list more manageable. For more information, see [Pagination].
//
// For more information about EBS snapshots, see [Amazon EBS snapshots] in the Amazon EBS User Guide.
//
// We strongly recommend using only paginated requests. Unpaginated requests are
// susceptible to throttling and timeouts.
//
// [Pagination]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination
// [Amazon EBS snapshots]: https://docs.aws.amazon.com/ebs/latest/userguide/ebs-snapshots.html
