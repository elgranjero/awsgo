package cloudwatch

// GetInsightRuleReport is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatch.go.
//
// This operation returns the time series data collected by a Contributor Insights
// rule. The data includes the identity and number of contributors to the log
// group.
//
// You can also optionally return one or more statistics about each data point in
// the time series. These statistics can include the following:
//
// - UniqueContributors -- the number of unique contributors for each data point.
//
// - MaxContributorValue -- the value of the top contributor for each data point.
// The identity of the contributor might change for each data point in the graph.
//
// If this rule aggregates by COUNT, the top contributor for each data point is
//
// the contributor with the most occurrences in that period. If the rule aggregates
// by SUM, the top contributor is the contributor with the highest sum in the log
// field specified by the rule's Value , during that period.
//
// - SampleCount -- the number of data points matched by the rule.
//
// - Sum -- the sum of the values from all contributors during the time period
// represented by that data point.
//
// - Minimum -- the minimum value from a single observation during the time
// period represented by that data point.
//
// - Maximum -- the maximum value from a single observation during the time
// period represented by that data point.
//
// - Average -- the average value from all contributors during the time period
// represented by that data point.
