package backup

// ListTags is generated as a reference stub.
// Executable command wiring lives under cmd/backup.go.
//
// Returns the tags assigned to the resource, such as a target recovery point,
// backup plan, or backup vault.
//
// This operation returns results depending on the resource type used in the value
// for resourceArn . For example, recovery points of Amazon DynamoDB with Advanced
// Settings have an ARN (Amazon Resource Name) that begins with arn:aws:backup .
// Recovery points (backups) of DynamoDB without Advanced Settings enabled have an
// ARN that begins with arn:aws:dynamodb .
//
// When this operation is called and when you include values of resourceArn that
// have an ARN other than arn:aws:backup , it may return one of the exceptions
// listed below. To prevent this exception, include only values representing
// resource types that are fully managed by Backup. These have an ARN that begins
// arn:aws:backup and they are noted in the [Feature availability by resource] table.
//
// [Feature availability by resource]: https://docs.aws.amazon.com/aws-backup/latest/devguide/backup-feature-availability.html#features-by-resource
