package networkfirewall

// UpdateAvailabilityZoneChangeProtection is generated as a reference stub.
// Executable command wiring lives under cmd/networkfirewall.go.
//
// Modifies the AvailabilityZoneChangeProtection setting for a transit
// gateway-attached firewall. When enabled, this setting prevents accidental
// changes to the firewall's Availability Zone configuration. This helps protect
// against disrupting traffic flow in production environments.
//
// When enabled, you must disable this protection before using AssociateAvailabilityZones or DisassociateAvailabilityZones to modify the
// firewall's Availability Zone configuration.
