package ec2

// AssociateClientVpnTargetNetwork is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Associates a target network with a Client VPN endpoint. A target network is a
// subnet in a VPC. You can associate multiple subnets from the same VPC with a
// Client VPN endpoint. You can associate only one subnet in each Availability
// Zone. We recommend that you associate at least two subnets to provide
// Availability Zone redundancy.
//
// If you specified a VPC when you created the Client VPN endpoint or if you have
// previous subnet associations, the specified subnet must be in the same VPC. To
// specify a subnet that's in a different VPC, you must first modify the Client VPN
// endpoint (ModifyClientVpnEndpoint ) and change the VPC that's associated with it.
