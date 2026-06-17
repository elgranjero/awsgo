package ec2

// CreateVpnConnection is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Creates a VPN connection between an existing virtual private gateway or transit
// gateway and a customer gateway. The supported connection type is ipsec.1 .
//
// The response includes information that you need to give to your network
// administrator to configure your customer gateway.
//
// We strongly recommend that you use HTTPS when calling this operation because
// the response contains sensitive cryptographic information for configuring your
// customer gateway device.
//
// If you decide to shut down your VPN connection for any reason and later create
// a new VPN connection, you must reconfigure your customer gateway with the new
// information returned from this call.
//
// This is an idempotent operation. If you perform the operation more than once,
// Amazon EC2 doesn't return an error.
//
// For more information, see [Amazon Web Services Site-to-Site VPN] in the Amazon Web Services Site-to-Site VPN User
// Guide.
//
// [Amazon Web Services Site-to-Site VPN]: https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html
