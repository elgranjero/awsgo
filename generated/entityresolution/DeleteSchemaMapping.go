package entityresolution

// DeleteSchemaMapping is generated as a reference stub.
// Executable command wiring lives under cmd/entityresolution.go.
//
// Deletes the SchemaMapping with a given name. This operation will succeed even
// if a schema with the given name does not exist. This operation will fail if
// there is a MatchingWorkflow object that references the SchemaMapping in the
// workflow's InputSourceConfig .
