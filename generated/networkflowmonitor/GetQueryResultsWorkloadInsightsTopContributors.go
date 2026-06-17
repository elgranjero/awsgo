package networkflowmonitor

// GetQueryResultsWorkloadInsightsTopContributors is generated as a reference stub.
// Executable command wiring lives under cmd/networkflowmonitor.go.
//
// Return the data for a query with the Network Flow Monitor query interface. You
// specify the query that you want to return results for by providing a query ID
// and a monitor name.
//
// This query returns the top contributors for a scope for workload insights.
// Workload insights provide a high level view of network flow performance data
// collected by agents. To return the data for the top contributors, see
// GetQueryResultsWorkloadInsightsTopContributorsData .
//
// Create a query ID for this call by calling the corresponding API call to start
// the query, StartQueryWorkloadInsightsTopContributors . Use the scope ID that was
// returned for your account by CreateScope .
//
// Top contributors in Network Flow Monitor are network flows with the highest
// values for a specific metric type. Top contributors can be across all workload
// insights, for a given scope, or for a specific monitor. Use the applicable call
// for the top contributors that you want to be returned.
