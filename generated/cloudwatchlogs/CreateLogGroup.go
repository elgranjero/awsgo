package cloudwatchlogs

// CreateLogGroup is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Creates a log group with the specified name. You can create up to 1,000,000 log
// groups per Region per account.
//
// You must use the following guidelines when naming a log group:
//
// - Log group names must be unique within a Region for an Amazon Web Services
// account.
//
// - Log group names can be between 1 and 512 characters long.
//
// - Log group names consist of the following characters: a-z, A-Z, 0-9, '_'
// (underscore), '-' (hyphen), '/' (forward slash), '.' (period), and '#' (number
// sign)
//
// - Log group names can't start with the string aws/
//
// When you create a log group, by default the log events in the log group do not
// expire. To set a retention policy so that events expire and are deleted after a
// specified time, use [PutRetentionPolicy].
//
// If you associate an KMS key with the log group, ingested data is encrypted
// using the KMS key. This association is stored as long as the data encrypted with
// the KMS key is still within CloudWatch Logs. This enables CloudWatch Logs to
// decrypt this data whenever it is requested.
//
// If you attempt to associate a KMS key with the log group but the KMS key does
// not exist or the KMS key is disabled, you receive an InvalidParameterException
// error.
//
// CloudWatch Logs supports only symmetric KMS keys. Do not associate an
// asymmetric KMS key with your log group. For more information, see [Using Symmetric and Asymmetric Keys].
//
// [Using Symmetric and Asymmetric Keys]: https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html
// [PutRetentionPolicy]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutRetentionPolicy.html
