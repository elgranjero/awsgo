package storagegateway

// CreateSnapshotFromVolumeRecoveryPoint is generated as a reference stub.
// Executable command wiring lives under cmd/storagegateway.go.
//
// Initiates a snapshot of a gateway from a volume recovery point. This operation
// is only supported in the cached volume gateway type.
//
// A volume recovery point is a point in time at which all data of the volume is
// consistent and from which you can create a snapshot. To get a list of volume
// recovery point for cached volume gateway, use ListVolumeRecoveryPoints.
//
// In the CreateSnapshotFromVolumeRecoveryPoint request, you identify the volume
// by providing its Amazon Resource Name (ARN). You must also provide a description
// for the snapshot. When the gateway takes a snapshot of the specified volume, the
// snapshot and its description appear in the Storage Gateway console. In response,
// the gateway returns you a snapshot ID. You can use this snapshot ID to check the
// snapshot progress or later use it when you want to create a volume from a
// snapshot.
//
// To list or delete a snapshot, you must use the Amazon EC2 API. For more
// information, see [DescribeSnapshots]or [DeleteSnapshot] in the Amazon Elastic Compute Cloud API Reference.
//
// [DescribeSnapshots]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSnapshots.html
// [DeleteSnapshot]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DeleteSnapshot.html
