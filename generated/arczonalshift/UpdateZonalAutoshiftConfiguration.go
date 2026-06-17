package arczonalshift

// UpdateZonalAutoshiftConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/arczonalshift.go.
//
// The zonal autoshift configuration for a resource includes the practice run
// configuration and the status for running autoshifts, zonal autoshift status.
// When a resource has a practice run configuration, ARC starts weekly zonal shifts
// for the resource, to shift traffic away from an Availability Zone. Weekly
// practice runs help you to make sure that your application can continue to
// operate normally with the loss of one Availability Zone.
//
// You can update the zonal autoshift status to enable or disable zonal autoshift.
// When zonal autoshift is ENABLED , you authorize Amazon Web Services to shift
// away resource traffic for an application from an Availability Zone during
// events, on your behalf, to help reduce time to recovery. Traffic is also shifted
// away for the required weekly practice runs.
