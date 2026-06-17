package cloudwatchlogs

// AssociateKmsKey is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Associates the specified KMS key with either one log group in the account, or
// with all stored CloudWatch Logs query insights results in the account.
//
// When you use AssociateKmsKey , you specify either the logGroupName parameter or
// the resourceIdentifier parameter. You can't specify both of those parameters in
// the same operation.
//
// - Specify the logGroupName parameter to cause log events ingested into that
// log group to be encrypted with that key. Only the log events ingested after the
// key is associated are encrypted with that key.
//
// Associating a KMS key with a log group overrides any existing associations
//
// between the log group and a KMS key. After a KMS key is associated with a log
// group, all newly ingested data for the log group is encrypted using the KMS key.
// This association is stored as long as the data encrypted with the KMS key is
// still within CloudWatch Logs. This enables CloudWatch Logs to decrypt this data
// whenever it is requested.
//
// Associating a key with a log group does not cause the results of queries of
//
// that log group to be encrypted with that key. To have query results encrypted
// with a KMS key, you must use an AssociateKmsKey operation with the
// resourceIdentifier parameter that specifies a query-result resource.
//
// - Specify the resourceIdentifier parameter with a query-result resource, to
// use that key to encrypt the stored results of all future [StartQuery]operations in the
// account. The response from a [GetQueryResults]operation will still return the query results in
// plain text.
//
// Even if you have not associated a key with your query results, the query
//
// results are encrypted when stored, using the default CloudWatch Logs method.
//
// If you run a query from a monitoring account that queries logs in a source
//
// account, the query results key from the monitoring account, if any, is used.
//
// If you delete the key that is used to encrypt log events or log group query
// results, then all the associated stored log events or query results that were
// encrypted with that key will be unencryptable and unusable.
//
// CloudWatch Logs supports only symmetric KMS keys. Do not associate an
// asymmetric KMS key with your log group or query results. For more information,
// see [Using Symmetric and Asymmetric Keys].
//
// It can take up to 5 minutes for this operation to take effect.
//
// If you attempt to associate a KMS key with a log group but the KMS key does not
// exist or the KMS key is disabled, you receive an InvalidParameterException
// error.
//
// [Using Symmetric and Asymmetric Keys]: https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html
//
// [StartQuery]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_StartQuery.html
// [GetQueryResults]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_GetQueryResults.html
