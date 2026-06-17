package cloudwatchlogs

// DescribeLogGroups is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Returns information about log groups, including data sources that ingest into
// each log group. You can return all your log groups or filter the results by
// prefix. The results are ASCII-sorted by log group name.
//
// CloudWatch Logs doesn't support IAM policies that control access to the
// DescribeLogGroups action by using the aws:ResourceTag/key-name  condition key.
// Other CloudWatch Logs actions do support the use of the
// aws:ResourceTag/key-name condition key to control access. For more information
// about using tags to control access, see [Controlling access to Amazon Web Services resources using tags].
//
// If you are using CloudWatch cross-account observability, you can use this
// operation in a monitoring account and view data from the linked source accounts.
// For more information, see [CloudWatch cross-account observability].
//
// [CloudWatch cross-account observability]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html
// [Controlling access to Amazon Web Services resources using tags]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html
