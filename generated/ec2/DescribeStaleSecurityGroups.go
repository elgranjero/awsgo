package ec2

// DescribeStaleSecurityGroups is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Describes the stale security group rules for security groups referenced across
// a VPC peering connection, transit gateway connection, or with a security group
// VPC association. Rules are stale when they reference a deleted security group.
// Rules can also be stale if they reference a security group in a peer VPC for
// which the VPC peering connection has been deleted, across a transit gateway
// where the transit gateway has been deleted (or [the transit gateway security group referencing feature]has been disabled), or if a
// security group VPC association has been disassociated.
//
// [the transit gateway security group referencing feature]: https://docs.aws.amazon.com/vpc/latest/tgw/tgw-vpc-attachments.html#vpc-attachment-security
