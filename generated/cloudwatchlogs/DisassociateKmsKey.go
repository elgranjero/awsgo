package cloudwatchlogs

// DisassociateKmsKey is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Disassociates the specified KMS key from the specified log group or from all
// CloudWatch Logs Insights query results in the account.
//
// When you use DisassociateKmsKey , you specify either the logGroupName parameter
// or the resourceIdentifier parameter. You can't specify both of those parameters
// in the same operation.
//
// - Specify the logGroupName parameter to stop using the KMS key to encrypt
// future log events ingested and stored in the log group. Instead, they will be
// encrypted with the default CloudWatch Logs method. The log events that were
// ingested while the key was associated with the log group are still encrypted
// with that key. Therefore, CloudWatch Logs will need permissions for the key
// whenever that data is accessed.
//
// - Specify the resourceIdentifier parameter with the query-result resource to
// stop using the KMS key to encrypt the results of all future [StartQuery]operations in the
// account. They will instead be encrypted with the default CloudWatch Logs method.
// The results from queries that ran while the key was associated with the account
// are still encrypted with that key. Therefore, CloudWatch Logs will need
// permissions for the key whenever that data is accessed.
//
// It can take up to 5 minutes for this operation to take effect.
//
// [StartQuery]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_StartQuery.html
