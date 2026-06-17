package cloudwatchlogs

// PutMetricFilter is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Creates or updates a metric filter and associates it with the specified log
// group. With metric filters, you can configure rules to extract metric data from
// log events ingested through [PutLogEvents].
//
// The maximum number of metric filters that can be associated with a log group is
// 100.
//
// Using regular expressions in filter patterns is supported. For these filters,
// there is a quota of two regular expression patterns within a single filter
// pattern. There is also a quota of five regular expression patterns per log
// group. For more information about using regular expressions in filter patterns,
// see [Filter pattern syntax for metric filters, subscription filters, filter log events, and Live Tail].
//
// When you create a metric filter, you can also optionally assign a unit and
// dimensions to the metric that is created.
//
// Metrics extracted from log events are charged as custom metrics. To prevent
// unexpected high charges, do not specify high-cardinality fields such as
// IPAddress or requestID as dimensions. Each different value found for a
// dimension is treated as a separate metric and accrues charges as a separate
// custom metric.
//
// CloudWatch Logs might disable a metric filter if it generates 1,000 different
// name/value pairs for your specified dimensions within one hour.
//
// You can also set up a billing alarm to alert you if your charges are higher
// than expected. For more information, see [Creating a Billing Alarm to Monitor Your Estimated Amazon Web Services Charges].
//
// [Creating a Billing Alarm to Monitor Your Estimated Amazon Web Services Charges]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/monitor_estimated_charges_with_cloudwatch.html
// [PutLogEvents]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutLogEvents.html
// [Filter pattern syntax for metric filters, subscription filters, filter log events, and Live Tail]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html
