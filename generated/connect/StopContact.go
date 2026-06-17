package connect

// StopContact is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Ends the specified contact. Use this API to stop queued callbacks. It does not
// work for voice contacts that use the following initiation methods:
//
// - DISCONNECT
//
// - TRANSFER
//
// - QUEUE_TRANSFER
//
// - EXTERNAL_OUTBOUND
//
// - MONITOR
//
// Chat and task contacts can be terminated in any state, regardless of initiation
// method.
