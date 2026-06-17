package ec2

// EnableVpcClassicLinkDnsSupport is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// This action is deprecated.
//
// Enables a VPC to support DNS hostname resolution for ClassicLink. If enabled,
// the DNS hostname of a linked EC2-Classic instance resolves to its private IP
// address when addressed from an instance in the VPC to which it's linked.
// Similarly, the DNS hostname of an instance in a VPC resolves to its private IP
// address when addressed from a linked EC2-Classic instance.
//
// You must specify a VPC ID in the request.
