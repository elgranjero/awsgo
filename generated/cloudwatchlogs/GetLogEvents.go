package cloudwatchlogs

// GetLogEvents is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Lists log events from the specified log stream. You can list all of the log
// events or filter using a time range.
//
// GetLogEvents is a paginated operation. Each page returned can contain up to 1
// MB of log events or up to 10,000 log events. A returned page might only be
// partially full, or even empty. For example, if the result of a query would
// return 15,000 log events, the first page isn't guaranteed to have 10,000 log
// events even if they all fit into 1 MB.
//
// Partially full or empty pages don't necessarily mean that pagination is
// finished. As long as the nextBackwardToken or nextForwardToken returned is NOT
// equal to the nextToken that you passed into the API call, there might be more
// log events available. The token that you use depends on the direction you want
// to move in along the log stream. The returned tokens are never null.
//
// If you set startFromHead to true and you don’t include endTime in your request,
// you can end up in a situation where the pagination doesn't terminate. This can
// happen when the new log events are being added to the target log streams faster
// than they are being read. This situation is a good use case for the CloudWatch
// Logs [Live Tail]feature.
//
// If you are using CloudWatch cross-account observability, you can use this
// operation in a monitoring account and view data from the linked source accounts.
// For more information, see [CloudWatch cross-account observability].
//
// You can specify the log group to search by using either logGroupIdentifier or
// logGroupName . You must include one of these two parameters, but you can't
// include both.
//
// If you are using [log transformation], the GetLogEvents operation returns only the original
// versions of log events, before they were transformed. To view the transformed
// versions, you must use a [CloudWatch Logs query.]
//
// [log transformation]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html
// [CloudWatch cross-account observability]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html
// [CloudWatch Logs query.]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AnalyzingLogData.html
// [Live Tail]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatchLogs_LiveTail.html
