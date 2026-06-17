package xray

// BatchGetTraces is generated as a reference stub.
// Executable command wiring lives under cmd/xray.go.
//
// You cannot find traces through this API if Transaction Search is enabled since
// trace is not indexed in X-Ray.
//
// Retrieves a list of traces specified by ID. Each trace is a collection of
// segment documents that originates from a single request. Use GetTraceSummaries
// to get a list of trace IDs.
