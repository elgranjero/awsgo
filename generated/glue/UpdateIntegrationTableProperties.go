package glue

// UpdateIntegrationTableProperties is generated as a reference stub.
// Executable command wiring lives under cmd/glue.go.
//
// This API is used to provide optional override properties for the tables that
// need to be replicated. These properties can include properties for filtering and
// partitioning for the source and target tables. To set both source and target
// properties the same API need to be invoked with the Glue connection ARN as
// ResourceArn with SourceTableConfig , and the Glue database ARN as ResourceArn
// with TargetTableConfig respectively.
//
// The override will be reflected across all the integrations using same
// ResourceArn and source table.
