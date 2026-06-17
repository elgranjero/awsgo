package connect

// CreateContactFlowVersion is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Publishes a new version of the flow provided. Versions are immutable and
// monotonically increasing. If the FlowContentSha256 provided is different from
// the FlowContentSha256 of the $LATEST published flow content, then an error is
// returned. This API only supports creating versions for flows of type Campaign .
