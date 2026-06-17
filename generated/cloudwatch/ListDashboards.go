package cloudwatch

// ListDashboards is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatch.go.
//
// Returns a list of the dashboards for your account. If you include
// DashboardNamePrefix , only those dashboards with names starting with the prefix
// are listed. Otherwise, all dashboards in your account are listed.
//
// ListDashboards returns up to 1000 results on one page. If there are more than
// 1000 dashboards, you can call ListDashboards again and include the value you
// received for NextToken in the first call, to receive the next 1000 results.
