package ec2

// CreateSecurityGroup is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Creates a security group.
//
// A security group acts as a virtual firewall for your instance to control
// inbound and outbound traffic. For more information, see [Amazon EC2 security groups]in the Amazon EC2 User
// Guide and [Security groups for your VPC]in the Amazon VPC User Guide.
//
// When you create a security group, you specify a friendly name of your choice.
// You can't have two security groups for the same VPC with the same name.
//
// You have a default security group for use in your VPC. If you don't specify a
// security group when you launch an instance, the instance is launched into the
// appropriate default security group. A default security group includes a default
// rule that grants instances unrestricted network access to each other.
//
// You can add or remove rules from your security groups using AuthorizeSecurityGroupIngress, AuthorizeSecurityGroupEgress, RevokeSecurityGroupIngress, and RevokeSecurityGroupEgress.
//
// For more information about VPC security group limits, see [Amazon VPC Limits].
//
// [Amazon VPC Limits]: https://docs.aws.amazon.com/vpc/latest/userguide/amazon-vpc-limits.html
// [Amazon EC2 security groups]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-network-security.html
// [Security groups for your VPC]: https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_SecurityGroups.html
