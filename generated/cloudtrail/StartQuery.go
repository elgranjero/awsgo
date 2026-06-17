package cloudtrail

// StartQuery is generated as a reference stub.
// Executable command wiring lives under cmd/cloudtrail.go.
//
// Starts a CloudTrail Lake query. Use the QueryStatement parameter to provide
// your SQL query, enclosed in single quotation marks. Use the optional
// DeliveryS3Uri parameter to deliver the query results to an S3 bucket.
//
// StartQuery requires you specify either the QueryStatement parameter, or a
// QueryAlias and any QueryParameters . In the current release, the QueryAlias and
// QueryParameters parameters are used only for the queries that populate the
// CloudTrail Lake dashboards.
