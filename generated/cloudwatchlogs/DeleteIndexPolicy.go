package cloudwatchlogs

// DeleteIndexPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Deletes a log-group level field index policy that was applied to a single log
// group. The indexing of the log events that happened before you delete the policy
// will still be used for as many as 30 days to improve CloudWatch Logs Insights
// queries.
//
// If the deleted policy included facet configurations, those facets will no
// longer be available for interactive exploration in the CloudWatch Logs Insights
// console for this log group. However, facet data is retained for up to 30 days.
//
// You can't use this operation to delete an account-level index policy. Instead,
// use [DeleteAccountPolicy].
//
// If you delete a log-group level field index policy and there is an
// account-level field index policy, in a few minutes the log group begins using
// that account-wide policy to index new incoming log events. This operation only
// affects log group-level policies, including any facet configurations, and
// preserves any data source-based account policies that may apply to the log
// group.
//
// [DeleteAccountPolicy]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DeleteAccountPolicy.html
