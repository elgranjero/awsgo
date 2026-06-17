package storagegateway

// CreateSnapshot is generated as a reference stub.
// Executable command wiring lives under cmd/storagegateway.go.
//
// Initiates a snapshot of a volume.
//
// Storage Gateway provides the ability to back up point-in-time snapshots of your
// data to Amazon Simple Storage (Amazon S3) for durable off-site recovery, and
// also import the data to an Amazon Elastic Block Store (EBS) volume in Amazon
// Elastic Compute Cloud (EC2). You can take snapshots of your gateway volume on a
// scheduled or ad hoc basis. This API enables you to take an ad hoc snapshot. For
// more information, see [Editing a snapshot schedule].
//
// In the CreateSnapshot request, you identify the volume by providing its Amazon
// Resource Name (ARN). You must also provide description for the snapshot. When
// Storage Gateway takes the snapshot of specified volume, the snapshot and
// description appears in the Storage Gateway console. In response, Storage Gateway
// returns you a snapshot ID. You can use this snapshot ID to check the snapshot
// progress or later use it when you want to create a volume from a snapshot. This
// operation is only supported in stored and cached volume gateway type.
//
// To list or delete a snapshot, you must use the Amazon EC2 API. For more
// information, see [DescribeSnapshots]or [DeleteSnapshot] in the Amazon Elastic Compute Cloud API Reference.
//
// Volume and snapshot IDs are changing to a longer length ID format. For more
// information, see the important note on the [Welcome]page.
//
// [Editing a snapshot schedule]: https://docs.aws.amazon.com/storagegateway/latest/userguide/managing-volumes.html#SchedulingSnapshot
// [Welcome]: https://docs.aws.amazon.com/storagegateway/latest/APIReference/Welcome.html
// [DescribeSnapshots]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSnapshots.html
// [DeleteSnapshot]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DeleteSnapshot.html
