package redshiftserverless

// CreateWorkgroup is generated as a reference stub.
// Executable command wiring lives under cmd/redshiftserverless.go.
//
// Creates an workgroup in Amazon Redshift Serverless.
//
// VPC Block Public Access (BPA) enables you to block resources in VPCs and
// subnets that you own in a Region from reaching or being reached from the
// internet through internet gateways and egress-only internet gateways. If a
// workgroup is in an account with VPC BPA turned on, the following capabilities
// are blocked:
//
// - Creating a public access workgroup
//
// - Modifying a private workgroup to public
//
// - Adding a subnet with VPC BPA turned on to the workgroup when the workgroup
// is public
//
// For more information about VPC BPA, see [Block public access to VPCs and subnets] in the Amazon VPC User Guide.
//
// [Block public access to VPCs and subnets]: https://docs.aws.amazon.com/vpc/latest/userguide/security-vpc-bpa.html
