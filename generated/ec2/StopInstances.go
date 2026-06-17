package ec2

// StopInstances is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Stops an Amazon EBS-backed instance. You can restart your instance at any time
// using the [StartInstances]API. For more information, see [Stop and start Amazon EC2 instances] in the Amazon EC2 User Guide.
//
// When you stop or hibernate an instance, we shut it down. By default, this
// includes a graceful operating system (OS) shutdown. To bypass the graceful
// shutdown, use the skipOsShutdown parameter; however, this might risk data
// integrity.
//
// You can use the StopInstances operation together with the Hibernate parameter
// to hibernate an instance if the instance is [enabled for hibernation]and meets the [hibernation prerequisites]. Stopping an
// instance doesn't preserve data stored in RAM, while hibernation does. If
// hibernation fails, a normal shutdown occurs. For more information, see [Hibernate your Amazon EC2 instance]in the
// Amazon EC2 User Guide.
//
// If your instance appears stuck in the stopping state, there might be an issue
// with the underlying host computer. You can use the StopInstances operation
// together with the Force parameter to force stop your instance. For more
// information, see [Troubleshoot Amazon EC2 instance stop issues]in the Amazon EC2 User Guide.
//
// Stopping and hibernating an instance differs from rebooting or terminating it.
// For example, a stopped or hibernated instance retains its root volume and any
// data volumes, unlike terminated instances where these volumes are automatically
// deleted. For more information about the differences between stopping,
// hibernating, rebooting, and terminating instances, see [Amazon EC2 instance state changes]in the Amazon EC2 User
// Guide.
//
// We don't charge for instance usage or data transfer fees when an instance is
// stopped. However, the root volume and any data volumes remain and continue to
// persist your data, and you're charged for volume usage. Every time you start
// your instance, Amazon EC2 charges a one-minute minimum for instance usage,
// followed by per-second billing.
//
// You can't stop or hibernate instance store-backed instances.
//
// [Troubleshoot Amazon EC2 instance stop issues]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/TroubleshootingInstancesStopping.html
// [Stop and start Amazon EC2 instances]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Stop_Start.html
// [Hibernate your Amazon EC2 instance]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Hibernate.html
// [Amazon EC2 instance state changes]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-lifecycle.html
// [enabled for hibernation]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/enabling-hibernation.html
// [StartInstances]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_StartInstances.html
// [hibernation prerequisites]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/hibernating-prerequisites.html
