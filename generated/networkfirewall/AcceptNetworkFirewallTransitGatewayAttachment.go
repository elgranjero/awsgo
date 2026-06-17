package networkfirewall

// AcceptNetworkFirewallTransitGatewayAttachment is generated as a reference stub.
// Executable command wiring lives under cmd/networkfirewall.go.
//
// Accepts a transit gateway attachment request for Network Firewall. When you
// accept the attachment request, Network Firewall creates the necessary routing
// components to enable traffic flow between the transit gateway and firewall
// endpoints.
//
// You must accept a transit gateway attachment to complete the creation of a
// transit gateway-attached firewall, unless auto-accept is enabled on the transit
// gateway. After acceptance, use DescribeFirewallto verify the firewall status.
//
// To reject an attachment instead of accepting it, use RejectNetworkFirewallTransitGatewayAttachment.
//
// It can take several minutes for the attachment acceptance to complete and the
// firewall to become available.
