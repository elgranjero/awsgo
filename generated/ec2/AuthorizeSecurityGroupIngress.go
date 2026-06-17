package ec2

// AuthorizeSecurityGroupIngress is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Adds the specified inbound (ingress) rules to a security group.
//
// An inbound rule permits instances to receive traffic from the specified IPv4 or
// IPv6 address range, the IP address ranges that are specified by a prefix list,
// or the instances that are associated with a destination security group. For more
// information, see [Security group rules].
//
// You must specify exactly one of the following sources: an IPv4 or IPv6 address
// range, a prefix list, or a security group. You must specify a protocol for each
// rule (for example, TCP). If the protocol is TCP or UDP, you must also specify a
// port or port range. If the protocol is ICMP or ICMPv6, you must also specify the
// ICMP/ICMPv6 type and code.
//
// Rule changes are propagated to instances associated with the security group as
// quickly as possible. However, a small delay might occur.
//
// For examples of rules that you can add to security groups for specific access
// scenarios, see [Security group rules for different use cases]in the Amazon EC2 User Guide.
//
// For more information about security group quotas, see [Amazon VPC quotas] in the Amazon VPC User
// Guide.
//
// [Amazon VPC quotas]: https://docs.aws.amazon.com/vpc/latest/userguide/amazon-vpc-limits.html
// [Security group rules]: https://docs.aws.amazon.com/vpc/latest/userguide/security-group-rules.html
// [Security group rules for different use cases]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/security-group-rules-reference.html
