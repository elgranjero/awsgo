package networkfirewall

// StartFlowCapture is generated as a reference stub.
// Executable command wiring lives under cmd/networkfirewall.go.
//
// Begins capturing the flows in a firewall, according to the filters you define.
// Captures are similar, but not identical to snapshots. Capture operations provide
// visibility into flows that are not closed and are tracked by a firewall's flow
// table. Unlike snapshots, captures are a time-boxed view.
//
// A flow is network traffic that is monitored by a firewall, either by stateful
// or stateless rules. For traffic to be considered part of a flow, it must share
// Destination, DestinationPort, Direction, Protocol, Source, and SourcePort.
//
// To avoid encountering operation limits, you should avoid starting captures with
// broad filters, like wide IP ranges. Instead, we recommend you define more
// specific criteria with FlowFilters , like narrow IP ranges, ports, or protocols.
