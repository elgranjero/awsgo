package ec2

// CreateVolume is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Creates an EBS volume that can be attached to an instance in the same
// Availability Zone.
//
// You can create a new empty volume or restore a volume from an EBS snapshot. Any
// Amazon Web Services Marketplace product codes from the snapshot are propagated
// to the volume.
//
// You can create encrypted volumes. Encrypted volumes must be attached to
// instances that support Amazon EBS encryption. Volumes that are created from
// encrypted snapshots are also automatically encrypted. For more information, see [Amazon EBS encryption]
// in the Amazon EBS User Guide.
//
// You can tag your volumes during creation. For more information, see [Tag your Amazon EC2 resources] in the
// Amazon EC2 User Guide.
//
// For more information, see [Create an Amazon EBS volume] in the Amazon EBS User Guide.
//
// [Amazon EBS encryption]: https://docs.aws.amazon.com/ebs/latest/userguide/ebs-encryption.html
// [Create an Amazon EBS volume]: https://docs.aws.amazon.com/ebs/latest/userguide/ebs-creating-volume.html
// [Tag your Amazon EC2 resources]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html
