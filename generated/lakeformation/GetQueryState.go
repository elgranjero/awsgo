package lakeformation

// GetQueryState is generated as a reference stub.
// Executable command wiring lives under cmd/lakeformation.go.
//
// Returns the state of a query previously submitted. Clients are expected to poll
// GetQueryState to monitor the current state of the planning before retrieving the
// work units. A query state is only visible to the principal that made the initial
// call to StartQueryPlanning .
