package networkfirewall

// StartFlowFlush is generated as a reference stub.
// Executable command wiring lives under cmd/networkfirewall.go.
//
// Begins the flushing of traffic from the firewall, according to the filters you
// define. When the operation starts, impacted flows are temporarily marked as
// timed out before the Suricata engine prunes, or flushes, the flows from the
// firewall table.
//
// While the flush completes, impacted flows are processed as midstream traffic.
// This may result in a temporary increase in midstream traffic metrics. We
// recommend that you double check your stream exception policy before you perform
// a flush operation.
