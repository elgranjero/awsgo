package networkflowmonitor

// GetQueryStatusWorkloadInsightsTopContributorsData is generated as a reference stub.
// Executable command wiring lives under cmd/networkflowmonitor.go.
//
// Returns the current status of a query for the Network Flow Monitor query
// interface, for a specified query ID and monitor. This call returns the query
// status for the top contributors data for workload insights.
//
// When you start a query, use this call to check the status of the query to make
// sure that it has has SUCCEEDED before you review the results. Use the same
// query ID that you used for the corresponding API call to start the query,
// StartQueryWorkloadInsightsTopContributorsData .
//
// Top contributors in Network Flow Monitor are network flows with the highest
// values for a specific metric type. Top contributors can be across all workload
// insights, for a given scope, or for a specific monitor. Use the applicable call
// for the top contributors that you want to be returned.
//
// The top contributor network flows overall are for a specific metric type, for
// example, the number of retransmissions.
