package cloudwatchlogs

// ListLogGroups is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Returns a list of log groups in the Region in your account. If you are
// performing this action in a monitoring account, you can choose to also return
// log groups from source accounts that are linked to the monitoring account. For
// more information about using cross-account observability to set up monitoring
// accounts and source accounts, see [CloudWatch cross-account observability].
//
// You can optionally filter the list by log group class, by using regular
// expressions in your request to match strings in the log group names, by using
// the fieldIndexes parameter to filter log groups based on which field indexes are
// configured, by using the dataSources parameter to filter log groups by data
// source types, and by using the fieldIndexNames parameter to filter by specific
// field index names.
//
// This operation is paginated. By default, your first use of this operation
// returns 50 results, and includes a token to use in a subsequent operation to
// return more results.
//
// [CloudWatch cross-account observability]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html
