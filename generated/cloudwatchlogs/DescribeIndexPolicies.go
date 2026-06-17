package cloudwatchlogs

// DescribeIndexPolicies is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Returns the field index policies of the specified log group. For more
// information about field index policies, see [PutIndexPolicy].
//
// If a specified log group has a log-group level index policy, that policy is
// returned by this operation.
//
// If a specified log group doesn't have a log-group level index policy, but an
// account-wide index policy applies to it, that account-wide policy is returned by
// this operation.
//
// To find information about only account-level policies, use [DescribeAccountPolicies] instead.
//
// [PutIndexPolicy]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutIndexPolicy.html
// [DescribeAccountPolicies]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DescribeAccountPolicies.html
