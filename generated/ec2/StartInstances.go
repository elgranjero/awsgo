package ec2

// StartInstances is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Starts an Amazon EBS-backed instance that you've previously stopped.
//
// Instances that use Amazon EBS volumes as their root devices can be quickly
// stopped and started. When an instance is stopped, the compute resources are
// released and you are not billed for instance usage. However, your root partition
// Amazon EBS volume remains and continues to persist your data, and you are
// charged for Amazon EBS volume usage. You can restart your instance at any time.
// Every time you start your instance, Amazon EC2 charges a one-minute minimum for
// instance usage, and thereafter charges per second for instance usage.
//
// Before stopping an instance, make sure it is in a state from which it can be
// restarted. Stopping an instance does not preserve data stored in RAM.
//
// Performing this operation on an instance that uses an instance store as its
// root device returns an error.
//
// If you attempt to start a T3 instance with host tenancy and the unlimited CPU
// credit option, the request fails. The unlimited CPU credit option is not
// supported on Dedicated Hosts. Before you start the instance, either change its
// CPU credit option to standard , or change its tenancy to default or dedicated .
//
// For more information, see [Stop and start Amazon EC2 instances] in the Amazon EC2 User Guide.
//
// [Stop and start Amazon EC2 instances]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Stop_Start.html
