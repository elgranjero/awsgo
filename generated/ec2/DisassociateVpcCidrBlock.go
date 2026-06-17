package ec2

// DisassociateVpcCidrBlock is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Disassociates a CIDR block from a VPC. To disassociate the CIDR block, you must
// specify its association ID. You can get the association ID by using DescribeVpcs. You must
// detach or delete all gateways and resources that are associated with the CIDR
// block before you can disassociate it.
//
// You cannot disassociate the CIDR block with which you originally created the
// VPC (the primary CIDR block).
