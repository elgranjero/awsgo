package networkfirewall

// AssociateAvailabilityZones is generated as a reference stub.
// Executable command wiring lives under cmd/networkfirewall.go.
//
// Associates the specified Availability Zones with a transit gateway-attached
// firewall. For each Availability Zone, Network Firewall creates a firewall
// endpoint to process traffic. You can specify one or more Availability Zones
// where you want to deploy the firewall.
//
// After adding Availability Zones, you must update your transit gateway route
// tables to direct traffic through the new firewall endpoints. Use DescribeFirewallto monitor the
// status of the new endpoints.
