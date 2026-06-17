package securityhub

// DeleteFindingAggregator is generated as a reference stub.
// Executable command wiring lives under cmd/securityhub.go.
//
// The aggregation Region is now called the home Region.
//
// Deletes a finding aggregator. When you delete the finding aggregator, you stop
// cross-Region aggregation. Finding replication stops occurring from the linked
// Regions to the home Region.
//
// When you stop cross-Region aggregation, findings that were already replicated
// and sent to the home Region are still visible from the home Region. However, new
// findings and finding updates are no longer replicated and sent to the home
// Region.
