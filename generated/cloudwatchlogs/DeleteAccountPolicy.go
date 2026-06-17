package cloudwatchlogs

// DeleteAccountPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Deletes a CloudWatch Logs account policy. This stops the account-wide policy
// from applying to log groups or data sources in the account. If you delete a data
// protection policy or subscription filter policy, any log-group level policies of
// those types remain in effect. This operation supports deletion of data
// source-based field index policies, including facet configurations, in addition
// to log group-based policies.
//
// To use this operation, you must be signed on with the correct permissions
// depending on the type of policy that you are deleting.
//
// - To delete a data protection policy, you must have the
// logs:DeleteDataProtectionPolicy and logs:DeleteAccountPolicy permissions.
//
// - To delete a subscription filter policy, you must have the
// logs:DeleteSubscriptionFilter and logs:DeleteAccountPolicy permissions.
//
// - To delete a transformer policy, you must have the logs:DeleteTransformer and
// logs:DeleteAccountPolicy permissions.
//
// - To delete a field index policy, you must have the logs:DeleteIndexPolicy and
// logs:DeleteAccountPolicy permissions.
//
// If you delete a field index policy that included facet configurations, those
//
// facets will no longer be available for interactive exploration in the CloudWatch
// Logs Insights console. However, facet data is retained for up to 30 days.
//
// If you delete a field index policy, the indexing of the log events that
// happened before you deleted the policy will still be used for up to 30 days to
// improve CloudWatch Logs Insights queries.
