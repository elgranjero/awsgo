package cloudwatchlogs

// PutQueryDefinition is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchlogs.go.
//
// Creates or updates a query definition for CloudWatch Logs Insights. For more
// information, see [Analyzing Log Data with CloudWatch Logs Insights].
//
// To update a query definition, specify its queryDefinitionId in your request.
// The values of name , queryString , and logGroupNames are changed to the values
// that you specify in your update operation. No current values are retained from
// the current query definition. For example, imagine updating a current query
// definition that includes log groups. If you don't specify the logGroupNames
// parameter in your update operation, the query definition changes to contain no
// log groups.
//
// You must have the logs:PutQueryDefinition permission to be able to perform this
// operation.
//
// [Analyzing Log Data with CloudWatch Logs Insights]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AnalyzingLogData.html
