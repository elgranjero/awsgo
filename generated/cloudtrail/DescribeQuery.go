package cloudtrail

// DescribeQuery is generated as a reference stub.
// Executable command wiring lives under cmd/cloudtrail.go.
//
// Returns metadata about a query, including query run time in milliseconds,
// number of events scanned and matched, and query status. If the query results
// were delivered to an S3 bucket, the response also provides the S3 URI and the
// delivery status.
//
// You must specify either QueryId or QueryAlias . Specifying the QueryAlias
// parameter returns information about the last query run for the alias. You can
// provide RefreshId along with QueryAlias to view the query results of a
// dashboard query for the specified RefreshId .
