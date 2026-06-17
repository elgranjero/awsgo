package networkfirewall

// RejectNetworkFirewallTransitGatewayAttachment is generated as a reference stub.
// Executable command wiring lives under cmd/networkfirewall.go.
//
// Rejects a transit gateway attachment request for Network Firewall. When you
// reject the attachment request, Network Firewall cancels the creation of routing
// components between the transit gateway and firewall endpoints.
//
// Only the transit gateway owner can reject the attachment. After rejection, no
// traffic will flow through the firewall endpoints for this attachment.
//
// Use DescribeFirewall to monitor the rejection status. To accept the attachment instead of
// rejecting it, use AcceptNetworkFirewallTransitGatewayAttachment.
//
// Once rejected, you cannot reverse this action. To establish connectivity, you
// must create a new transit gateway-attached firewall.
