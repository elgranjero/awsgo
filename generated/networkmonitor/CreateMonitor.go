package networkmonitor

// CreateMonitor is generated as a reference stub.
// Executable command wiring lives under cmd/networkmonitor.go.
//
// Creates a monitor between a source subnet and destination IP address. Within a
// monitor you'll create one or more probes that monitor network traffic between
// your source Amazon Web Services VPC subnets and your destination IP addresses.
// Each probe then aggregates and sends metrics to Amazon CloudWatch.
//
// You can also create a monitor with probes using this command. For each probe,
// you define the following:
//
// - source —The subnet IDs where the probes will be created.
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
