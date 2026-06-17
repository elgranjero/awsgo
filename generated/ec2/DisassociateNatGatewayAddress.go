package ec2

// DisassociateNatGatewayAddress is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Disassociates secondary Elastic IP addresses (EIPs) from a public NAT gateway.
// You cannot disassociate your primary EIP. For more information, see [Edit secondary IP address associations]in the
// Amazon VPC User Guide.
//
// While disassociating is in progress, you cannot associate/disassociate
// additional EIPs while the connections are being drained. You are, however,
// allowed to delete the NAT gateway.
//
// An EIP is released only at the end of MaxDrainDurationSeconds. It stays
// associated and supports the existing connections but does not support any new
// connections (new connections are distributed across the remaining associated
// EIPs). As the existing connections drain out, the EIPs (and the corresponding
// private IP addresses mapped to them) are released.
//
// [Edit secondary IP address associations]: https://docs.aws.amazon.com/vpc/latest/userguide/nat-gateway-working-with.html#nat-gateway-edit-secondary
