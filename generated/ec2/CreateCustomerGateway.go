package ec2

// CreateCustomerGateway is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Provides information to Amazon Web Services about your customer gateway device.
// The customer gateway device is the appliance at your end of the VPN connection.
// You must provide the IP address of the customer gateway device’s external
// interface. The IP address must be static and can be behind a device performing
// network address translation (NAT).
//
// For devices that use Border Gateway Protocol (BGP), you can also provide the
// device's BGP Autonomous System Number (ASN). You can use an existing ASN
// assigned to your network. If you don't have an ASN already, you can use a
// private ASN. For more information, see [Customer gateway options for your Site-to-Site VPN connection]in the Amazon Web Services Site-to-Site
// VPN User Guide.
//
// To create more than one customer gateway with the same VPN type, IP address,
// and BGP ASN, specify a unique device name for each customer gateway. An
// identical request returns information about the existing customer gateway; it
// doesn't create a new customer gateway.
//
// [Customer gateway options for your Site-to-Site VPN connection]: https://docs.aws.amazon.com/vpn/latest/s2svpn/cgw-options.html
