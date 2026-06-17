package glue

// RegisterSchemaVersion is generated as a reference stub.
// Executable command wiring lives under cmd/glue.go.
//
// Adds a new version to the existing schema. Returns an error if new version of
// schema does not meet the compatibility requirements of the schema set. This API
// will not create a new schema set and will return a 404 error if the schema set
// is not already present in the Schema Registry.
//
// If this is the first schema definition to be registered in the Schema Registry,
// this API will store the schema version and return immediately. Otherwise, this
// call has the potential to run longer than other operations due to compatibility
// modes. You can call the GetSchemaVersion API with the SchemaVersionId to check
// compatibility modes.
//
// If the same schema definition is already stored in Schema Registry as a
// version, the schema ID of the existing schema is returned to the caller.
