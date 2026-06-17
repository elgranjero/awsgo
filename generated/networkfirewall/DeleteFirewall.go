package networkfirewall

// DeleteFirewall is generated as a reference stub.
// Executable command wiring lives under cmd/networkfirewall.go.
//
// Deletes the specified Firewall and its FirewallStatus. This operation requires the firewall's
// DeleteProtection flag to be FALSE . You can't revert this operation.
//
// You can check whether a firewall is in use by reviewing the route tables for
// the Availability Zones where you have firewall subnet mappings. Retrieve the
// subnet mappings by calling DescribeFirewall. You define and update the route tables through
// Amazon VPC. As needed, update the route tables for the zones to remove the
// firewall endpoints. When the route tables no longer use the firewall endpoints,
// you can remove the firewall safely.
//
// To delete a firewall, remove the delete protection if you need to using UpdateFirewallDeleteProtection, then
// delete the firewall by calling DeleteFirewall.
