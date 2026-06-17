package ec2

// CreateTransitGatewayVpcAttachment is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Attaches the specified VPC to the specified transit gateway.
//
// If you attach a VPC with a CIDR range that overlaps the CIDR range of a VPC
// that is already attached, the new VPC CIDR range is not propagated to the
// default propagation route table.
//
// To send VPC traffic to an attached transit gateway, add a route to the VPC
// route table using CreateRoute.
