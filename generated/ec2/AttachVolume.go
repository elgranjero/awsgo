package ec2

// AttachVolume is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Attaches an Amazon EBS volume to a running or stopped instance, and exposes it
// to the instance with the specified device name.
//
// The maximum number of Amazon EBS volumes that you can attach to an instance
// depends on the instance type. If you exceed the volume attachment limit for an
// instance type, the attachment request fails with the AttachmentLimitExceeded
// error. For more information, see [Instance volume limits].
//
// After you attach an EBS volume, you must make it available for use. For more
// information, see [Make an EBS volume available for use].
//
// If a volume has an Amazon Web Services Marketplace product code:
//
// - The volume can be attached only to a stopped instance.
//
// - Amazon Web Services Marketplace product codes are copied from the volume to
// the instance.
//
// - You must be subscribed to the product.
//
// - The instance type and operating system of the instance must support the
// product. For example, you can't detach a volume from a Windows instance and
// attach it to a Linux instance.
//
// For more information, see [Attach an Amazon EBS volume to an instance] in the Amazon EBS User Guide.
//
// [Make an EBS volume available for use]: https://docs.aws.amazon.com/ebs/latest/userguide/ebs-using-volumes.html
// [Attach an Amazon EBS volume to an instance]: https://docs.aws.amazon.com/ebs/latest/userguide/ebs-attaching-volume.html
// [Instance volume limits]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/volume_limits.html
