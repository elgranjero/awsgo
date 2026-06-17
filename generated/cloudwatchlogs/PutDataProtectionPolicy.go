package cloudwatchlogs

// PutDataProtectionPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Creates a data protection policy for the specified log group. A data protection
// policy can help safeguard sensitive data that's ingested by the log group by
// auditing and masking the sensitive log data.
//
// Sensitive data is detected and masked when it is ingested into the log group.
// When you set a data protection policy, log events ingested into the log group
// before that time are not masked.
//
// By default, when a user views a log event that includes masked data, the
// sensitive data is replaced by asterisks. A user who has the logs:Unmask
// permission can use a [GetLogEvents]or [FilterLogEvents] operation with the unmask parameter set to true to
// view the unmasked log events. Users with the logs:Unmask can also view unmasked
// data in the CloudWatch Logs console by running a CloudWatch Logs Insights query
// with the unmask query command.
//
// For more information, including a list of types of data that can be audited and
// masked, see [Protect sensitive log data with masking].
//
// The PutDataProtectionPolicy operation applies to only the specified log group.
// You can also use [PutAccountPolicy]to create an account-level data protection policy that applies
// to all log groups in the account, including both existing log groups and log
// groups that are created level. If a log group has its own data protection policy
// and the account also has an account-level data protection policy, then the two
// policies are cumulative. Any sensitive term specified in either policy is
// masked.
//
// [Protect sensitive log data with masking]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/mask-sensitive-log-data.html
// [FilterLogEvents]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_FilterLogEvents.html
// [PutAccountPolicy]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutAccountPolicy.html
// [GetLogEvents]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_GetLogEvents.html
