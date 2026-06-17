package storagegateway

// DeleteSnapshotSchedule is generated as a reference stub.
// Executable command wiring lives under cmd/storagegateway.go.
//
// Deletes a snapshot of a volume.
//
// You can take snapshots of your gateway volumes on a scheduled or ad hoc basis.
// This API action enables you to delete a snapshot schedule for a volume. For more
// information, see [Backing up your volumes]. In the DeleteSnapshotSchedule request, you identify the
// volume by providing its Amazon Resource Name (ARN). This operation is only
// supported for cached volume gateway types.
//
// To list or delete a snapshot, you must use the Amazon EC2 API. For more
// information, go to [DescribeSnapshots]in the Amazon Elastic Compute Cloud API Reference.
//
// [Backing up your volumes]: https://docs.aws.amazon.com/storagegateway/latest/userguide/backing-up-volumes.html
// [DescribeSnapshots]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSnapshots.html
