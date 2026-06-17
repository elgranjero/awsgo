package networkflowmonitor

// GetQueryResultsWorkloadInsightsTopContributorsData is generated as a reference stub.
// Executable command wiring lives under cmd/networkflowmonitor.go.
//
// Return the data for a query with the Network Flow Monitor query interface.
// Specify the query that you want to return results for by providing a query ID
// and a scope ID.
//
// This query returns the data for top contributors for workload insights for a
// specific scope. Workload insights provide a high level view of network flow
// performance data collected by agents for a scope. To return just the top
// contributors, see GetQueryResultsWorkloadInsightsTopContributors .
//
// Create a query ID for this call by calling the corresponding API call to start
// the query, StartQueryWorkloadInsightsTopContributorsData . Use the scope ID that
// was returned for your account by CreateScope .
//
// Top contributors in Network Flow Monitor are network flows with the highest
// values for a specific metric type. Top contributors can be across all workload
// insights, for a given scope, or for a specific monitor. Use the applicable call
// for the top contributors that you want to be returned.
//
// The top contributor network flows overall are for a specific metric type, for
// example, the number of retransmissions.
