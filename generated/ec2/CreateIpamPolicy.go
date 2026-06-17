package ec2

// CreateIpamPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Creates an IPAM policy.
//
// An IPAM policy is a set of rules that define how public IPv4 addresses from
// IPAM pools are allocated to Amazon Web Services resources. Each rule maps an
// Amazon Web Services service to IPAM pools that the service will use to get IP
// addresses. A single policy can have multiple rules and be applied to multiple
// Amazon Web Services Regions. If the IPAM pool run out of addresses then the
// services fallback to Amazon-provided IP addresses. A policy can be applied to an
// individual Amazon Web Services account or an entity within Amazon Web Services
// Organizations.
//
// For more information, see [Define public IPv4 allocation strategy with IPAM policies] in the Amazon VPC IPAM User Guide.
//
// [Define public IPv4 allocation strategy with IPAM policies]: https://docs.aws.amazon.com/vpc/latest/ipam/define-public-ipv4-allocation-strategy-with-ipam-policies.html
