package cloudwatchlogs

// PutSubscriptionFilter is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Creates or updates a subscription filter and associates it with the specified
// log group. With subscription filters, you can subscribe to a real-time stream of
// log events ingested through [PutLogEvents]and have them delivered to a specific destination.
// When log events are sent to the receiving service, they are Base64 encoded and
// compressed with the GZIP format.
//
// The following destinations are supported for subscription filters:
//
// - An Amazon Kinesis data stream belonging to the same account as the
// subscription filter, for same-account delivery.
//
// - A logical destination created with [PutDestination]that belongs to a different account, for
// cross-account delivery. We currently support Kinesis Data Streams and Firehose
// as logical destinations.
//
// - An Amazon Kinesis Data Firehose delivery stream that belongs to the same
// account as the subscription filter, for same-account delivery.
//
// - An Lambda function that belongs to the same account as the subscription
// filter, for same-account delivery.
//
// Each log group can have up to two subscription filters associated with it. If
// you are updating an existing filter, you must specify the correct name in
// filterName .
//
// Using regular expressions in filter patterns is supported. For these filters,
// there is a quotas of quota of two regular expression patterns within a single
// filter pattern. There is also a quota of five regular expression patterns per
// log group. For more information about using regular expressions in filter
// patterns, see [Filter pattern syntax for metric filters, subscription filters, filter log events, and Live Tail].
//
// To perform a PutSubscriptionFilter operation for any destination except a
// Lambda function, you must also have the iam:PassRole permission.
//
// [PutDestination]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDestination.html
// [PutLogEvents]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutLogEvents.html
// [Filter pattern syntax for metric filters, subscription filters, filter log events, and Live Tail]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html
