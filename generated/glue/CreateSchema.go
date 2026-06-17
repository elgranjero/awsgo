package glue

// CreateSchema is generated as a reference stub.
// Executable command wiring lives under cmd/glue.go.
//
// Creates a new schema set and registers the schema definition. Returns an error
// if the schema set already exists without actually registering the version.
//
// When the schema set is created, a version checkpoint will be set to the first
// version. Compatibility mode "DISABLED" restricts any additional schema versions
// from being added after the first schema version. For all other compatibility
// modes, validation of compatibility settings will be applied only from the second
// version onwards when the RegisterSchemaVersion API is used.
//
// When this API is called without a RegistryId , this will create an entry for a
// "default-registry" in the registry database tables, if it is not already
// present.
