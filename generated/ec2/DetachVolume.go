package ec2

// DetachVolume is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Detaches an EBS volume from an instance. Make sure to unmount any file systems
// on the device within your operating system before detaching the volume. Failure
// to do so can result in the volume becoming stuck in the busy state while
// detaching. If this happens, detachment can be delayed indefinitely until you
// unmount the volume, force detachment, reboot the instance, or all three. If an
// EBS volume is the root device of an instance, it can't be detached while the
// instance is running. To detach the root volume, stop the instance first.
//
// When a volume with an Amazon Web Services Marketplace product code is detached
// from an instance, the product code is no longer associated with the instance.
//
// You can't detach or force detach volumes that are attached to Amazon Web
// Services-managed resources. Attempting to do this results in the
// UnsupportedOperationException exception.
//
// For more information, see [Detach an Amazon EBS volume] in the Amazon EBS User Guide.
//
// [Detach an Amazon EBS volume]: https://docs.aws.amazon.com/ebs/latest/userguide/ebs-detaching-volume.html
