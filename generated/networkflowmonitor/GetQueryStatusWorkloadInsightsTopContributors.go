package networkflowmonitor

// GetQueryStatusWorkloadInsightsTopContributors is generated as a reference stub.
// Executable command wiring lives under cmd/networkflowmonitor.go.
//
// Return the data for a query with the Network Flow Monitor query interface.
// Specify the query that you want to return results for by providing a query ID
// and a monitor name. This query returns the top contributors for workload
// insights.
//
// When you start a query, use this call to check the status of the query to make
// sure that it has has SUCCEEDED before you review the results. Use the same
// query ID that you used for the corresponding API call to start the query,
// StartQueryWorkloadInsightsTopContributors .
//
// Top contributors in Network Flow Monitor are network flows with the highest
// values for a specific metric type. Top contributors can be across all workload
// insights, for a given scope, or for a specific monitor. Use the applicable call
// for the top contributors that you want to be returned.
