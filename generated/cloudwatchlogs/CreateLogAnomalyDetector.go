package cloudwatchlogs

// CreateLogAnomalyDetector is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Creates an anomaly detector that regularly scans one or more log groups and
// look for patterns and anomalies in the logs.
//
// An anomaly detector can help surface issues by automatically discovering
// anomalies in your log event traffic. An anomaly detector uses machine learning
// algorithms to scan log events and find patterns. A pattern is a shared text
// structure that recurs among your log fields. Patterns provide a useful tool for
// analyzing large sets of logs because a large number of log events can often be
// compressed into a few patterns.
//
// The anomaly detector uses pattern recognition to find anomalies , which are
// unusual log events. It uses the evaluationFrequency to compare current log
// events and patterns with trained baselines.
//
// Fields within a pattern are called tokens. Fields that vary within a pattern,
// such as a request ID or timestamp, are referred to as dynamic tokens and
// represented by <*> .
//
// The following is an example of a pattern:
//
// [INFO] Request time: <*> ms
//
// This pattern represents log events like [INFO] Request time: 327 ms and other
// similar log events that differ only by the number, in this csse 327. When the
// pattern is displayed, the different numbers are replaced by <*>
//
// Any parts of log events that are masked as sensitive data are not scanned for
// anomalies. For more information about masking sensitive data, see [Help protect sensitive log data with masking].
//
// [Help protect sensitive log data with masking]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/mask-sensitive-log-data.html
