package eventbridge

// CreateEndpoint is generated as a reference stub.
// Executable command wiring lives under cmd/eventbridge.go.
//
// Creates a global endpoint. Global endpoints improve your application's
// availability by making it regional-fault tolerant. To do this, you define a
// primary and secondary Region with event buses in each Region. You also create a
// Amazon Route 53 health check that will tell EventBridge to route events to the
// secondary Region when an "unhealthy" state is encountered and events will be
// routed back to the primary Region when the health check reports a "healthy"
// state.
