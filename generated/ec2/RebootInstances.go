package ec2

// RebootInstances is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Requests a reboot of the specified instances. This operation is asynchronous;
// it only queues a request to reboot the specified instances. The operation
// succeeds if the instances are valid and belong to you. Requests to reboot
// terminated instances are ignored.
//
// If an instance does not cleanly shut down within a few minutes, Amazon EC2
// performs a hard reboot.
//
// For more information about troubleshooting, see [Troubleshoot an unreachable instance] in the Amazon EC2 User Guide.
//
// [Troubleshoot an unreachable instance]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-console.html
