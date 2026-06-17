package ec2

// CreateTransitGateway is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Creates a transit gateway.
//
// You can use a transit gateway to interconnect your virtual private clouds (VPC)
// and on-premises networks. After the transit gateway enters the available state,
// you can attach your VPCs and VPN connections to the transit gateway.
//
// To attach your VPCs, use CreateTransitGatewayVpcAttachment.
//
// To attach a VPN connection, use CreateCustomerGateway to create a customer gateway and specify the
// ID of the customer gateway and the ID of the transit gateway in a call to CreateVpnConnection.
//
// When you create a transit gateway, we create a default transit gateway route
// table and use it as the default association route table and the default
// propagation route table. You can use CreateTransitGatewayRouteTableto create additional transit gateway route
// tables. If you disable automatic route propagation, we do not create a default
// transit gateway route table. You can use EnableTransitGatewayRouteTablePropagationto propagate routes from a resource
// attachment to a transit gateway route table. If you disable automatic
// associations, you can use AssociateTransitGatewayRouteTableto associate a resource attachment with a transit
// gateway route table.
