package ec2

// RevokeSecurityGroupIngress is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Removes the specified inbound (ingress) rules from a security group.
//
// You can specify rules using either rule IDs or security group rule properties.
// If you use rule properties, the values that you specify (for example, ports)
// must match the existing rule's values exactly. Each rule has a protocol, from
// and to ports, and source (CIDR range, security group, or prefix list). For the
// TCP and UDP protocols, you must also specify the destination port or range of
// ports. For the ICMP protocol, you must also specify the ICMP type and code. If
// the security group rule has a description, you do not need to specify the
// description to revoke the rule.
//
// For a default VPC, if the values you specify do not match the existing rule's
// values, no error is returned, and the output describes the security group rules
// that were not revoked.
//
// For a non-default VPC, if the values you specify do not match the existing
// rule's values, an InvalidPermission.NotFound client error is returned, and no
// rules are revoked.
//
// Amazon Web Services recommends that you describe the security group to verify
// that the rules were removed.
//
// Rule changes are propagated to instances within the security group as quickly
// as possible. However, a small delay might occur.
