package ec2

// DeregisterImage is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Deregisters the specified AMI. A deregistered AMI can't be used to launch new
// instances.
//
// If a deregistered EBS-backed AMI matches a Recycle Bin retention rule, it moves
// to the Recycle Bin for the specified retention period. It can be restored before
// its retention period expires, after which it is permanently deleted. If the
// deregistered AMI doesn't match a retention rule, it is permanently deleted
// immediately. For more information, see [Recover deleted Amazon EBS snapshots and EBS-backed AMIs with Recycle Bin]in the Amazon EBS User Guide.
//
// When deregistering an EBS-backed AMI, you can optionally delete its associated
// snapshots at the same time. However, if a snapshot is associated with multiple
// AMIs, it won't be deleted even if specified for deletion, although the AMI will
// still be deregistered.
//
// Deregistering an AMI does not delete the following:
//
// - Instances already launched from the AMI. You'll continue to incur usage
// costs for the instances until you terminate them.
//
// - For EBS-backed AMIs: Snapshots that are associated with multiple AMIs.
// You'll continue to incur snapshot storage costs.
//
// - For instance store-backed AMIs: The files uploaded to Amazon S3 during AMI
// creation. You'll continue to incur S3 storage costs.
//
// For more information, see [Deregister an Amazon EC2 AMI] in the Amazon EC2 User Guide.
//
// [Deregister an Amazon EC2 AMI]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/deregister-ami.html
// [Recover deleted Amazon EBS snapshots and EBS-backed AMIs with Recycle Bin]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/recycle-bin.html
