package networkflowmonitor

// StartQueryMonitorTopContributors is generated as a reference stub.
// Executable command wiring lives under cmd/networkflowmonitor.go.
//
// Create a query that you can use with the Network Flow Monitor query interface
// to return the top contributors for a monitor. Specify the monitor that you want
// to create the query for.
//
// The call returns a query ID that you can use with [GetQueryResultsMonitorTopContributors] to run the query and return
// the top contributors for a specific monitor.
//
// Top contributors in Network Flow Monitor are network flows with the highest
// values for a specific metric type. Top contributors can be across all workload
// insights, for a given scope, or for a specific monitor. Use the applicable APIs
// for the top contributors that you want to be returned.
//
// [GetQueryResultsMonitorTopContributors]: https://docs.aws.amazon.com/networkflowmonitor/2.0/APIReference/API_GetQueryResultsMonitorTopContributors.html
