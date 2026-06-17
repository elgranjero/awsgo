package ec2

// ResetInstanceAttribute is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Resets an attribute of an instance to its default value. To reset the kernel or
// ramdisk , the instance must be in a stopped state. To reset the sourceDestCheck
// , the instance can be either running or stopped.
//
// The sourceDestCheck attribute controls whether source/destination checking is
// enabled. The default value is true , which means checking is enabled. This value
// must be false for a NAT instance to perform NAT. For more information, see [NAT instances] in
// the Amazon VPC User Guide.
//
// [NAT instances]: https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_NAT_Instance.html
