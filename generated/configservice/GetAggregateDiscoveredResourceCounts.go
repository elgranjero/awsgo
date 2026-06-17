package configservice

// GetAggregateDiscoveredResourceCounts is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// Returns the resource counts across accounts and regions that are present in
// your Config aggregator. You can request the resource counts by providing filters
// and GroupByKey.
//
// For example, if the input contains accountID 12345678910 and region us-east-1
// in filters, the API returns the count of resources in account ID 12345678910 and
// region us-east-1. If the input contains ACCOUNT_ID as a GroupByKey, the API
// returns resource counts for all source accounts that are present in your
// aggregator.
