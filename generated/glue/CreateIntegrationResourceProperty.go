package glue

// CreateIntegrationResourceProperty is generated as a reference stub.
// Executable command wiring lives under cmd/glue.go.
//
// This API can be used for setting up the ResourceProperty of the Glue connection
// (for the source) or Glue database ARN (for the target). These properties can
// include the role to access the connection or database. To set both source and
// target properties the same API needs to be invoked with the Glue connection ARN
// as ResourceArn with SourceProcessingProperties and the Glue database ARN as
// ResourceArn with TargetProcessingProperties respectively.
