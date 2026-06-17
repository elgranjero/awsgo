package timestreamquery

// Query is generated as a reference stub.
// Executable command wiring lives under cmd/timestreamquery.go.
//
// Query is a synchronous operation that enables you to run a query against your
// Amazon Timestream data.
//
// If you enabled QueryInsights , this API also returns insights and metrics
// related to the query that you executed. QueryInsights helps with performance
// tuning of your query. For more information about QueryInsights , see [Using query insights to optimize queries in Amazon Timestream].
//
// The maximum number of Query API requests you're allowed to make with
// QueryInsights enabled is 1 query per second (QPS). If you exceed this query
// rate, it might result in throttling.
//
// Query will time out after 60 seconds. You must update the default timeout in
// the SDK to support a timeout of 60 seconds. See the [code sample]for details.
//
// Your query request will fail in the following cases:
//
// - If you submit a Query request with the same client token outside of the
// 5-minute idempotency window.
//
// - If you submit a Query request with the same client token, but change other
// parameters, within the 5-minute idempotency window.
//
// - If the size of the row (including the query metadata) exceeds 1 MB, then
// the query will fail with the following error message:
//
// Query aborted as max page response size has been exceeded by the output result
//
// row
//
// - If the IAM principal of the query initiator and the result reader are not
// the same and/or the query initiator and the result reader do not have the same
// query string in the query requests, the query will fail with an Invalid
// pagination token error.
//
// [code sample]: https://docs.aws.amazon.com/timestream/latest/developerguide/code-samples.run-query.html
// [Using query insights to optimize queries in Amazon Timestream]: https://docs.aws.amazon.com/timestream/latest/developerguide/using-query-insights.html
