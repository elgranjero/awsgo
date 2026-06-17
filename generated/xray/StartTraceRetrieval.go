package xray

// StartTraceRetrieval is generated as a reference stub.
// Executable command wiring lives under cmd/xray.go.
//
// Initiates a trace retrieval process using the specified time range and for the
//
// given trace IDs in the Transaction Search generated CloudWatch log group. For
// more information, see [Transaction Search].
//
// API returns a RetrievalToken , which can be used with ListRetrievedTraces or
// GetRetrievedTracesGraph to fetch results. Retrievals will time out after 60
// minutes. To execute long time ranges, consider segmenting into multiple
// retrievals.
//
// If you are using [CloudWatch cross-account observability], you can use this operation in a monitoring account to
// retrieve data from a linked source account, as long as both accounts have
// transaction search enabled.
//
// For retrieving data from X-Ray directly as opposed to the Transaction-Search
// Log group, see [BatchGetTraces].
//
// [BatchGetTraces]: https://docs.aws.amazon.com/xray/latest/api/API_BatchGetTraces.html
// [Transaction Search]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Transaction-Search.html
// [CloudWatch cross-account observability]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html
