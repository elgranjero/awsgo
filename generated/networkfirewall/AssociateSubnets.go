package networkfirewall

// AssociateSubnets is generated as a reference stub.
// Executable command wiring lives under cmd/networkfirewall.go.
//
// Associates the specified subnets in the Amazon VPC to the firewall. You can
// specify one subnet for each of the Availability Zones that the VPC spans.
//
// This request creates an Network Firewall firewall endpoint in each of the
// subnets. To enable the firewall's protections, you must also modify the VPC's
// route tables for each subnet's Availability Zone, to redirect the traffic that's
// coming into and going out of the zone through the firewall endpoint.
