package cloudwatchlogs

// DescribeLogStreams is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Lists the log streams for the specified log group. You can list all the log
// streams or filter the results by prefix. You can also control how the results
// are ordered.
//
// You can specify the log group to search by using either logGroupIdentifier or
// logGroupName . You must include one of these two parameters, but you can't
// include both.
//
// This operation has a limit of 25 transactions per second, after which
// transactions are throttled.
//
// If you are using CloudWatch cross-account observability, you can use this
// operation in a monitoring account and view data from the linked source accounts.
// For more information, see [CloudWatch cross-account observability].
//
// [CloudWatch cross-account observability]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html
