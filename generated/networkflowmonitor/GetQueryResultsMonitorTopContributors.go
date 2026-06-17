package networkflowmonitor

// GetQueryResultsMonitorTopContributors is generated as a reference stub.
// Executable command wiring lives under cmd/networkflowmonitor.go.
//
// Return the data for a query with the Network Flow Monitor query interface. You
// specify the query that you want to return results for by providing a query ID
// and a monitor name. This query returns the top contributors for a specific
// monitor.
//
// Create a query ID for this call by calling the corresponding API call to start
// the query, StartQueryMonitorTopContributors . Use the scope ID that was returned
// for your account by CreateScope .
//
// Top contributors in Network Flow Monitor are network flows with the highest
// values for a specific metric type. Top contributors can be across all workload
// insights, for a given scope, or for a specific monitor. Use the applicable call
// for the top contributors that you want to be returned.
