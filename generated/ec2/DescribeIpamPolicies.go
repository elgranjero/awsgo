package ec2

// DescribeIpamPolicies is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Describes one or more IPAM policies.
//
// An IPAM policy is a set of rules that define how public IPv4 addresses from
// IPAM pools are allocated to Amazon Web Services resources. Each rule maps an
// Amazon Web Services service to IPAM pools that the service will use to get IP
// addresses. A single policy can have multiple rules and be applied to multiple
// Amazon Web Services Regions. If the IPAM pool run out of addresses then the
// services fallback to Amazon-provided IP addresses. A policy can be applied to an
// individual Amazon Web Services account or an entity within Amazon Web Services
// Organizations.
