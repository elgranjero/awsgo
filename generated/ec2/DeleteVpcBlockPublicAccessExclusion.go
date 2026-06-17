package ec2

// DeleteVpcBlockPublicAccessExclusion is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Delete a VPC Block Public Access (BPA) exclusion. A VPC BPA exclusion is a mode
// that can be applied to a single VPC or subnet that exempts it from the account’s
// BPA mode and will allow bidirectional or egress-only access. You can create BPA
// exclusions for VPCs and subnets even when BPA is not enabled on the account to
// ensure that there is no traffic disruption to the exclusions when VPC BPA is
// turned on. To learn more about VPC BPA, see [Block public access to VPCs and subnets]in the Amazon VPC User Guide.
//
// [Block public access to VPCs and subnets]: https://docs.aws.amazon.com/vpc/latest/userguide/security-vpc-bpa.html
