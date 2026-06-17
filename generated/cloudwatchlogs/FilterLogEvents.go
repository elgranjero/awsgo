package cloudwatchlogs

// FilterLogEvents is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Lists log events from the specified log group. You can list all the log events
// or filter the results using one or more of the following:
//
// - A filter pattern
//
// - A time range
//
// - The log stream name, or a log stream name prefix that matches multiple log
// streams
//
// You must have the logs:FilterLogEvents permission to perform this operation.
//
// You can specify the log group to search by using either logGroupIdentifier or
// logGroupName . You must include one of these two parameters, but you can't
// include both.
//
// FilterLogEvents is a paginated operation. Each page returned can contain up to
// 1 MB of log events or up to 10,000 log events. A returned page might only be
// partially full, or even empty. For example, if the result of a query would
// return 15,000 log events, the first page isn't guaranteed to have 10,000 log
// events even if they all fit into 1 MB.
//
// Partially full or empty pages don't necessarily mean that pagination is
// finished. If the results include a nextToken , there might be more log events
// available. You can return these additional log events by providing the nextToken
// in a subsequent FilterLogEvents operation. If the results don't include a
// nextToken , then pagination is finished.
//
// Specifying the limit parameter only guarantees that a single page doesn't
// return more log events than the specified limit, but it might return fewer
// events than the limit. This is the expected API behavior.
//
// The returned log events are sorted by event timestamp, the timestamp when the
// event was ingested by CloudWatch Logs, and the ID of the PutLogEvents request.
//
// If you are using CloudWatch cross-account observability, you can use this
// operation in a monitoring account and view data from the linked source accounts.
// For more information, see [CloudWatch cross-account observability].
//
// If you are using [log transformation], the FilterLogEvents operation returns only the original
// versions of log events, before they were transformed. To view the transformed
// versions, you must use a [CloudWatch Logs query.]
//
// [log transformation]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html
// [CloudWatch cross-account observability]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html
// [CloudWatch Logs query.]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AnalyzingLogData.html
