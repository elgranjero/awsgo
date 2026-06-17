package ec2

// CreateNetworkAclEntry is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Creates an entry (a rule) in a network ACL with the specified rule number. Each
// network ACL has a set of numbered ingress rules and a separate set of numbered
// egress rules. When determining whether a packet should be allowed in or out of a
// subnet associated with the ACL, we process the entries in the ACL according to
// the rule numbers, in ascending order. Each network ACL has a set of ingress
// rules and a separate set of egress rules.
//
// We recommend that you leave room between the rule numbers (for example, 100,
// 110, 120, ...), and not number them one right after the other (for example, 101,
// 102, 103, ...). This makes it easier to add a rule between existing ones without
// having to renumber the rules.
//
// After you add an entry, you can't modify it; you must either replace it, or
// create an entry and delete the old one.
//
// For more information about network ACLs, see [Network ACLs] in the Amazon VPC User Guide.
//
// [Network ACLs]: https://docs.aws.amazon.com/vpc/latest/userguide/vpc-network-acls.html
