package dynamodb

// ListBackups is generated as a reference stub.
// Executable command wiring lives under cmd/dynamodb.go.
//
// List DynamoDB backups that are associated with an Amazon Web Services account
// and weren't made with Amazon Web Services Backup. To list these backups for a
// given table, specify TableName . ListBackups returns a paginated list of
// results with at most 1 MB worth of items in a page. You can also specify a
// maximum number of entries to be returned in a page.
//
// In the request, start time is inclusive, but end time is exclusive. Note that
// these boundaries are for the time at which the original backup was requested.
//
// You can call ListBackups a maximum of five times per second.
//
// If you want to retrieve the complete list of backups made with Amazon Web
// Services Backup, use the [Amazon Web Services Backup list API.]
//
// [Amazon Web Services Backup list API.]: https://docs.aws.amazon.com/aws-backup/latest/devguide/API_ListBackupJobs.html
