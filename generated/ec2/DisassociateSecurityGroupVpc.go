package ec2

// DisassociateSecurityGroupVpc is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Disassociates a security group from a VPC. You cannot disassociate the security
// group if any Elastic network interfaces in the associated VPC are still
// associated with the security group.
//
// Note that the disassociation is asynchronous and you can check the status of
// the request with [DescribeSecurityGroupVpcAssociations].
//
// [DescribeSecurityGroupVpcAssociations]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSecurityGroupVpcAssociations.html
