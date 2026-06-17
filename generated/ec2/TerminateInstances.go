package ec2

// TerminateInstances is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Terminates (deletes) the specified instances. This operation is [idempotent]; if you
// terminate an instance more than once, each call succeeds.
//
// Terminating an instance is permanent and irreversible.
//
// After you terminate an instance, you can no longer connect to it, and it can't
// be recovered. All attached Amazon EBS volumes that are configured to be deleted
// on termination are also permanently deleted and can't be recovered. All data
// stored on instance store volumes is permanently lost. For more information, see [How instance termination works]
// .
//
// Before you terminate an instance, ensure that you have backed up all data that
// you need to retain after the termination to persistent storage.
//
// If you specify multiple instances and the request fails (for example, because
// of a single incorrect instance ID), none of the instances are terminated.
//
// If you terminate multiple instances across multiple Availability Zones, and one
// or more of the specified instances are enabled for termination protection, the
// request fails with the following results:
//
// - The specified instances that are in the same Availability Zone as the
// protected instance are not terminated.
//
// - The specified instances that are in different Availability Zones, where no
// other specified instances are protected, are successfully terminated.
//
// For example, say you have the following instances:
//
// - Instance A: us-east-1a ; Not protected
//
// - Instance B: us-east-1a ; Not protected
//
// - Instance C: us-east-1b ; Protected
//
// - Instance D: us-east-1b ; not protected
//
// If you attempt to terminate all of these instances in the same request, the
// request reports failure with the following results:
//
// - Instance A and Instance B are successfully terminated because none of the
// specified instances in us-east-1a are enabled for termination protection.
//
// - Instance C and Instance D fail to terminate because at least one of the
// specified instances in us-east-1b (Instance C) is enabled for termination
// protection.
//
// Terminated instances remain visible after termination (for approximately one
// hour).
//
// By default, Amazon EC2 deletes all EBS volumes that were attached when the
// instance launched. Volumes attached after instance launch continue running.
//
// By default, the TerminateInstances operation includes a graceful operating
// system (OS) shutdown. To bypass the graceful shutdown, use the skipOsShutdown
// parameter; however, this might risk data integrity.
//
// You can stop, start, and terminate EBS-backed instances. You can only terminate
// instance store-backed instances. What happens to an instance differs if you stop
// or terminate it. For example, when you stop an instance, the root device and any
// other devices attached to the instance persist. When you terminate an instance,
// any attached EBS volumes with the DeleteOnTermination block device mapping
// parameter set to true are automatically deleted. For more information about the
// differences between stopping and terminating instances, see [Amazon EC2 instance state changes]in the Amazon EC2
// User Guide.
//
// When you terminate an instance, we attempt to terminate it forcibly after a
// short while. If your instance appears stuck in the shutting-down state after a
// period of time, there might be an issue with the underlying host computer. For
// more information about terminating and troubleshooting terminating your
// instances, see [Terminate Amazon EC2 instances]and [Troubleshooting terminating your instance] in the Amazon EC2 User Guide.
//
// [idempotent]: https://docs.aws.amazon.com/ec2/latest/devguide/ec2-api-idempotency.html
// [How instance termination works]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/how-ec2-instance-termination-works.html
// [Troubleshooting terminating your instance]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/TroubleshootingInstancesShuttingDown.html
// [Amazon EC2 instance state changes]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-lifecycle.html
// [Terminate Amazon EC2 instances]: https://docs.aws.amazon.com/
