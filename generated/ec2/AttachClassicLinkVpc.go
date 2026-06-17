package ec2

// AttachClassicLinkVpc is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// This action is deprecated.
//
// Links an EC2-Classic instance to a ClassicLink-enabled VPC through one or more
// of the VPC security groups. You cannot link an EC2-Classic instance to more than
// one VPC at a time. You can only link an instance that's in the running state.
// An instance is automatically unlinked from a VPC when it's stopped - you can
// link it to the VPC again when you restart it.
//
// After you've linked an instance, you cannot change the VPC security groups that
// are associated with it. To change the security groups, you must first unlink the
// instance, and then link it again.
//
// Linking your instance to a VPC is sometimes referred to as attaching your
// instance.
