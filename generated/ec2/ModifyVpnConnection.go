package ec2

// ModifyVpnConnection is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Modifies the customer gateway or the target gateway of an Amazon Web Services
// Site-to-Site VPN connection. To modify the target gateway, the following
// migration options are available:
//
// - An existing virtual private gateway to a new virtual private gateway
//
// - An existing virtual private gateway to a transit gateway
//
// - An existing transit gateway to a new transit gateway
//
// - An existing transit gateway to a virtual private gateway
//
// Before you perform the migration to the new gateway, you must configure the new
// gateway. Use CreateVpnGatewayto create a virtual private gateway, or CreateTransitGateway to create a transit
// gateway.
//
// This step is required when you migrate from a virtual private gateway with
// static routes to a transit gateway.
//
// You must delete the static routes before you migrate to the new gateway.
//
// Keep a copy of the static route before you delete it. You will need to add back
// these routes to the transit gateway after the VPN connection migration is
// complete.
//
// After you migrate to the new gateway, you might need to modify your VPC route
// table. Use CreateRouteand DeleteRoute to make the changes described in [Update VPC route tables] in the Amazon Web Services
// Site-to-Site VPN User Guide.
//
// When the new gateway is a transit gateway, modify the transit gateway route
// table to allow traffic between the VPC and the Amazon Web Services Site-to-Site
// VPN connection. Use CreateTransitGatewayRouteto add the routes.
//
// If you deleted VPN static routes, you must add the static routes to the transit
// gateway route table.
//
// After you perform this operation, the VPN endpoint's IP addresses on the Amazon
// Web Services side and the tunnel options remain intact. Your Amazon Web Services
// Site-to-Site VPN connection will be temporarily unavailable for a brief period
// while we provision the new endpoints.
//
// [Update VPC route tables]: https://docs.aws.amazon.com/vpn/latest/s2svpn/modify-vpn-target.html#step-update-routing
