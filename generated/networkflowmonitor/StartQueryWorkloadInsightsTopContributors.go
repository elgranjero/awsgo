package networkflowmonitor

// StartQueryWorkloadInsightsTopContributors is generated as a reference stub.
// Executable command wiring lives under cmd/networkflowmonitor.go.
//
// Create a query with the Network Flow Monitor query interface that you can run
// to return workload insights top contributors. Specify the scope that you want to
// create a query for.
//
// The call returns a query ID that you can use with [GetQueryResultsWorkloadInsightsTopContributors] to run the query and return
// the top contributors for the workload insights for a scope.
//
// Top contributors in Network Flow Monitor are network flows with the highest
// values for a specific metric type. Top contributors can be across all workload
// insights, for a given scope, or for a specific monitor. Use the applicable APIs
// for the top contributors that you want to be returned.
//
// [GetQueryResultsWorkloadInsightsTopContributors]: https://docs.aws.amazon.com/networkflowmonitor/2.0/APIReference/API_GetQueryResultsWorkloadInsightsTopContributors.html
