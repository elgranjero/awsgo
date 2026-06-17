package athena

// GetQueryResults is generated as a reference stub.
// Executable command wiring lives under cmd/athena.go.
//
// Streams the results of a single query execution specified by QueryExecutionId
// from the Athena query results location in Amazon S3. For more information, see [Working with query results, recent queries, and output files]
// in the Amazon Athena User Guide. This request does not execute the query but
// returns results. Use StartQueryExecutionto run a query.
//
// To stream query results successfully, the IAM principal with permission to call
// GetQueryResults also must have permissions to the Amazon S3 GetObject action
// for the Athena query results location.
//
// IAM principals with permission to the Amazon S3 GetObject action for the query
// results location are able to retrieve query results from Amazon S3 even if
// permission to the GetQueryResults action is denied. To restrict user or role
// access, ensure that Amazon S3 permissions to the Athena query location are
// denied.
//
// [Working with query results, recent queries, and output files]: https://docs.aws.amazon.com/athena/latest/ug/querying.html
