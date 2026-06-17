package cloudwatchlogs

// GetQueryResults is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Returns the results from the specified query.
//
// Only the fields requested in the query are returned, along with a (at)ptr field,
// which is the identifier for the log record. You can use the value of (at)ptr in a [GetLogRecord]
// operation to get the full log record.
//
// GetQueryResults does not start running a query. To run a query, use [StartQuery]. For more
// information about how long results of previous queries are available, see [CloudWatch Logs quotas].
//
// If the value of the Status field in the output is Running , this operation
// returns only partial results. If you see a value of Scheduled or Running for
// the status, you can retry the operation later to see the final results.
//
// This operation is used both for retrieving results from interactive queries and
// from automated scheduled query executions. Scheduled queries use GetQueryResults
// internally to retrieve query results for processing and delivery to configured
// destinations.
//
// If you are using CloudWatch cross-account observability, you can use this
// operation in a monitoring account to start queries in linked source accounts.
// For more information, see [CloudWatch cross-account observability].
//
// [CloudWatch cross-account observability]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html
// [GetLogRecord]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_GetLogRecord.html
// [CloudWatch Logs quotas]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/cloudwatch_limits_cwl.html
// [StartQuery]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_StartQuery.html
