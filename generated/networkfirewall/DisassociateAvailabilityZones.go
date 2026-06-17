package networkfirewall

// DisassociateAvailabilityZones is generated as a reference stub.
// Executable command wiring lives under cmd/networkfirewall.go.
//
// Removes the specified Availability Zone associations from a transit
// gateway-attached firewall. This removes the firewall endpoints from these
// Availability Zones and stops traffic filtering in those zones. Before removing
// an Availability Zone, ensure you've updated your transit gateway route tables to
// redirect traffic appropriately.
//
// If AvailabilityZoneChangeProtection is enabled, you must first disable it using UpdateAvailabilityZoneChangeProtection
// .
//
// To verify the status of your Availability Zone changes, use DescribeFirewall.
