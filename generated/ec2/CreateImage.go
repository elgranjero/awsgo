package ec2

// CreateImage is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Creates an Amazon EBS-backed AMI from an Amazon EBS-backed instance that is
// either running or stopped.
//
// If you customized your instance with instance store volumes or Amazon EBS
// volumes in addition to the root device volume, the new AMI contains block device
// mapping information for those volumes. When you launch an instance from this new
// AMI, the instance automatically launches with those additional volumes.
//
// The location of the source instance determines where you can create the
// snapshots of the AMI:
//
// - If the source instance is in a Region, you must create the snapshots in the
// same Region as the instance.
//
// - If the source instance is in a Local Zone, you can create the snapshots in
// the same Local Zone or in its parent Region.
//
// For more information, see [Create an Amazon EBS-backed AMI] in the Amazon Elastic Compute Cloud User Guide.
//
// [Create an Amazon EBS-backed AMI]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/creating-an-ami-ebs.html
