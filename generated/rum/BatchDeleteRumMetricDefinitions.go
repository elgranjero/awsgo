package rum

// BatchDeleteRumMetricDefinitions is generated as a reference stub.
// Executable command wiring lives under cmd/rum.go.
//
// Removes the specified metrics from being sent to an extended metrics
// destination.
//
// If some metric definition IDs specified in a BatchDeleteRumMetricDefinitions
// operations are not valid, those metric definitions fail and return errors, but
// all valid metric definition IDs in the same operation are still deleted.
//
// The maximum number of metric definitions that you can specify in one
// BatchDeleteRumMetricDefinitions operation is 200.
