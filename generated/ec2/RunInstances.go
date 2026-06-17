package ec2

// RunInstances is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Launches the specified number of instances using an AMI for which you have
// permissions.
//
// You can specify a number of options, or leave the default options. The
// following rules apply:
//
// - If you don't specify a subnet ID, we choose a default subnet from your
// default VPC for you. If you don't have a default VPC, you must specify a subnet
// ID in the request.
//
// - All instances have a network interface with a primary private IPv4 address.
// If you don't specify this address, we choose one from the IPv4 range of your
// subnet.
//
// - Not all instance types support IPv6 addresses. For more information, see [Instance types].
//
// - If you don't specify a security group ID, we use the default security group
// for the VPC. For more information, see [Security groups].
//
// - If any of the AMIs have a product code attached for which the user has not
// subscribed, the request fails.
//
// You can create a [launch template], which is a resource that contains the parameters to launch
// an instance. When you launch an instance using RunInstances, you can specify the launch
// template instead of specifying the launch parameters.
//
// To ensure faster instance launches, break up large requests into smaller
// batches. For example, create five separate launch requests for 100 instances
// each instead of one launch request for 500 instances.
//
// RunInstances is subject to both request rate limiting and resource rate
// limiting. For more information, see [Request throttling].
//
// An instance is ready for you to use when it's in the running state. You can
// check the state of your instance using DescribeInstances. You can tag instances and EBS volumes
// during launch, after launch, or both. For more information, see CreateTagsand [Tagging your Amazon EC2 resources].
//
// Linux instances have access to the public key of the key pair at boot. You can
// use this key to provide secure access to the instance. Amazon EC2 public images
// use this feature to provide secure access without passwords. For more
// information, see [Key pairs].
//
// For troubleshooting, see [What to do if an instance immediately terminates], and [Troubleshooting connecting to your instance].
//
// [Key pairs]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html
// [What to do if an instance immediately terminates]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_InstanceStraightToTerminated.html
// [Tagging your Amazon EC2 resources]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html
// [launch template]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html
// [Security groups]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-network-security.html
// [Request throttling]: https://docs.aws.amazon.com/ec2/latest/devguide/ec2-api-throttling.html
// [Instance types]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html
// [Troubleshooting connecting to your instance]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/TroubleshootingInstancesConnecting.html
