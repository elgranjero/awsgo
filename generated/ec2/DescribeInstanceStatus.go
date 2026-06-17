package ec2

// DescribeInstanceStatus is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Describes the status of the specified instances or all of your instances. By
// default, only running instances are described, unless you specifically indicate
// to return the status of all instances.
//
// Instance status includes the following components:
//
// - Status checks - Amazon EC2 performs status checks on running EC2 instances
// to identify hardware and software issues. For more information, see [Status checks for your instances]and [Troubleshoot instances with failed status checks]in
// the Amazon EC2 User Guide.
//
// - Scheduled events - Amazon EC2 can schedule events (such as reboot, stop, or
// terminate) for your instances related to hardware issues, software updates, or
// system maintenance. For more information, see [Scheduled events for your instances]in the Amazon EC2 User Guide.
//
// - Instance state - You can manage your instances from the moment you launch
// them through their termination. For more information, see [Instance lifecycle]in the Amazon EC2
// User Guide.
//
// The Amazon EC2 API follows an eventual consistency model. This means that the
// result of an API command you run that creates or modifies resources might not be
// immediately available to all subsequent commands you run. For guidance on how to
// manage eventual consistency, see [Eventual consistency in the Amazon EC2 API]in the Amazon EC2 Developer Guide.
//
// The order of the elements in the response, including those within nested
// structures, might vary. Applications should not assume the elements appear in a
// particular order.
//
// [Troubleshoot instances with failed status checks]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/TroubleshootingInstances.html
// [Instance lifecycle]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-lifecycle.html
// [Status checks for your instances]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/monitoring-system-instance-status-check.html
// [Eventual consistency in the Amazon EC2 API]: https://docs.aws.amazon.com/ec2/latest/devguide/eventual-consistency.html
// [Scheduled events for your instances]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/monitoring-instances-status-check_sched.html
