package ec2

// CreateDefaultVpc is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Creates a default VPC with a size /16 IPv4 CIDR block and a default subnet in
// each Availability Zone. For more information about the components of a default
// VPC, see [Default VPCs]in the Amazon VPC User Guide. You cannot specify the components of the
// default VPC yourself.
//
// If you deleted your previous default VPC, you can create a default VPC. You
// cannot have more than one default VPC per Region.
//
// [Default VPCs]: https://docs.aws.amazon.com/vpc/latest/userguide/default-vpc.html
