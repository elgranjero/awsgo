package networkmonitor

// UpdateProbe is generated as a reference stub.
// Executable command wiring lives under cmd/networkmonitor.go.
//
// Updates a monitor probe. This action requires both the monitorName and probeId
// parameters. Run ListMonitors to get a list of monitor names. Run GetMonitor to
// get a list of probes and probe IDs.
//
// You can update the following para create a monitor with probes using this
// command. For each probe, you define the following:
//
// - state —The state of the probe.
//
// - destination — The target destination IP address for the probe.
//
// - destinationPort —Required only if the protocol is TCP .
//
// - protocol —The communication protocol between the source and destination.
// This will be either TCP or ICMP .
//
// - packetSize —The size of the packets. This must be a number between 56 and
// 8500 .
//
// - (Optional) tags —Key-value pairs created and assigned to the probe.
