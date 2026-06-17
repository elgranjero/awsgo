package rds

// EnableHttpEndpoint is generated as a reference stub.
// Executable command wiring lives under cmd/rds.go.
//
// Enables the HTTP endpoint for the DB cluster. By default, the HTTP endpoint
// isn't enabled.
//
// When enabled, this endpoint provides a connectionless web service API (RDS Data
// API) for running SQL queries on the Aurora DB cluster. You can also query your
// database from inside the RDS console with the RDS query editor.
//
// For more information, see [Using RDS Data API] in the Amazon Aurora User Guide.
//
// This operation applies only to Aurora Serverless v2 and provisioned DB
// clusters. To enable the HTTP endpoint for Aurora Serverless v1 DB clusters, use
// the EnableHttpEndpoint parameter of the ModifyDBCluster operation.
//
// [Using RDS Data API]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/data-api.html
