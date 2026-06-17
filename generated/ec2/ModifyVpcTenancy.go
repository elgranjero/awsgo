package ec2

// ModifyVpcTenancy is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Modifies the instance tenancy attribute of the specified VPC. You can change
// the instance tenancy attribute of a VPC to default only. You cannot change the
// instance tenancy attribute to dedicated .
//
// After you modify the tenancy of the VPC, any new instances that you launch into
// the VPC have a tenancy of default , unless you specify otherwise during launch.
// The tenancy of any existing instances in the VPC is not affected.
//
// For more information, see [Dedicated Instances] in the Amazon EC2 User Guide.
//
// [Dedicated Instances]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-instance.html
