package cloudwatchlogs

// StartQuery is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Starts a query of one or more log groups or data sources using CloudWatch Logs
// Insights. You specify the log groups or data sources and time range to query and
// the query string to use. You can query up to 10 data sources in a single query.
//
// For more information, see [CloudWatch Logs Insights Query Syntax].
//
// After you run a query using StartQuery , the query results are stored by
// CloudWatch Logs. You can use [GetQueryResults]to retrieve the results of a query, using the
// queryId that StartQuery returns.
//
// Interactive queries started with StartQuery share concurrency limits with
// automated scheduled query executions. Both types of queries count toward the
// same regional concurrent query quota, so high scheduled query activity may
// affect the availability of concurrent slots for interactive queries.
//
// To specify the log groups to query, a StartQuery operation must include one of
// the following:
//
// - Either exactly one of the following parameters: logGroupName , logGroupNames
// , or logGroupIdentifiers
//
// - Or the queryString must include a SOURCE command to select log groups for
// the query. The SOURCE command can select log groups based on log group name
// prefix, account ID, and log class, or select data sources using dataSource
// syntax in LogsQL, PPL, and SQL.
//
// For more information about the SOURCE command, see [SOURCE].
//
// If you have associated a KMS key with the query results in this account, then [StartQuery]
// uses that key to encrypt the results when it stores them. If no key is
// associated with query results, the query results are encrypted with the default
// CloudWatch Logs encryption method.
//
// Queries time out after 60 minutes of runtime. If your queries are timing out,
// reduce the time range being searched or partition your query into a number of
// queries.
//
// If you are using CloudWatch cross-account observability, you can use this
// operation in a monitoring account to start a query in a linked source account.
// For more information, see [CloudWatch cross-account observability]. For a cross-account StartQuery operation, the query
// definition must be defined in the monitoring account.
//
// You can have up to 30 concurrent CloudWatch Logs insights queries, including
// queries that have been added to dashboards.
//
// [CloudWatch Logs Insights Query Syntax]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax.html
// [CloudWatch cross-account observability]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html
// [SOURCE]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax-Source.html
// [GetQueryResults]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_GetQueryResults.html
// [StartQuery]: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_StartQuery.html
