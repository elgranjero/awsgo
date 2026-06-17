package cloudwatchlogs

// ListAggregateLogGroupSummaries is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Returns an aggregate summary of all log groups in the Region grouped by
// specified data source characteristics. Supports optional filtering by log group
// class, name patterns, and data sources. If you perform this action in a
// monitoring account, you can also return aggregated summaries of log groups from
// source accounts that are linked to the monitoring account. For more information
// about using cross-account observability to set up monitoring accounts and source
// accounts, see [CloudWatch cross-account observability].
//
// The operation aggregates log groups by data source name and type and optionally
// format, providing counts of log groups that share these characteristics. The
// operation paginates results. By default, it returns up to 50 results and
// includes a token to retrieve more results.
//
// [CloudWatch cross-account observability]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html
