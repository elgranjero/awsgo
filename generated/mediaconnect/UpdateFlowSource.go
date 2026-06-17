package mediaconnect

// UpdateFlowSource is generated as a reference stub.
// Executable command wiring lives under cmd/mediaconnect.go.
//
// Updates the source of a flow.
//
// Because UpdateFlowSources and UpdateFlow are separate operations, you can't
// change both the source type AND the flow size in a single request.
//
// - If you have a MEDIUM flow and you want to change the flow source to NDI®:
//
// - First, use the UpdateFlow operation to upgrade the flow size to LARGE .
//
// - After that, you can then use the UpdateFlowSource operation to configure the
// NDI source.
//
// - If you're switching from an NDI source to a transport stream (TS) source
// and want to downgrade the flow size:
//
// - First, use the UpdateFlowSource operation to change the flow source type.
//
// - After that, you can then use the UpdateFlow operation to downgrade the flow
// size to MEDIUM .
