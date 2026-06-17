package networkflowmonitor

// StartQueryWorkloadInsightsTopContributorsData is generated as a reference stub.
// Executable command wiring lives under cmd/networkflowmonitor.go.
//
// Create a query with the Network Flow Monitor query interface that you can run
// to return data for workload insights top contributors. Specify the scope that
// you want to create a query for.
//
// The call returns a query ID that you can use with [GetQueryResultsWorkloadInsightsTopContributorsData] to run the query and return
// the data for the top contributors for the workload insights for a scope.
//
// Top contributors in Network Flow Monitor are network flows with the highest
// values for a specific metric type. Top contributors can be across all workload
// insights, for a given scope, or for a specific monitor. Use the applicable call
// for the top contributors that you want to be returned.
//
// [GetQueryResultsWorkloadInsightsTopContributorsData]: https://docs.aws.amazon.com/networkflowmonitor/2.0/APIReference/API_GetQueryResultsWorkloadInsightsTopContributorsData.html
