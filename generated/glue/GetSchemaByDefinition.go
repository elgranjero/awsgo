package glue

// GetSchemaByDefinition is generated as a reference stub.
// Executable command wiring lives under cmd/glue.go.
//
// Retrieves a schema by the SchemaDefinition . The schema definition is sent to
// the Schema Registry, canonicalized, and hashed. If the hash is matched within
// the scope of the SchemaName or ARN (or the default registry, if none is
// supplied), that schema’s metadata is returned. Otherwise, a 404 or NotFound
// error is returned. Schema versions in Deleted statuses will not be included in
// the results.
