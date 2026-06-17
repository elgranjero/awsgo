package glue

// DeleteSchemaVersions is generated as a reference stub.
// Executable command wiring lives under cmd/glue.go.
//
// Remove versions from the specified schema. A version number or range may be
// supplied. If the compatibility mode forbids deleting of a version that is
// necessary, such as BACKWARDS_FULL, an error is returned. Calling the
// GetSchemaVersions API after this call will list the status of the deleted
// versions.
//
// When the range of version numbers contain check pointed version, the API will
// return a 409 conflict and will not proceed with the deletion. You have to remove
// the checkpoint first using the DeleteSchemaCheckpoint API before using this API.
//
// You cannot use the DeleteSchemaVersions API to delete the first schema version
// in the schema set. The first schema version can only be deleted by the
// DeleteSchema API. This operation will also delete the attached
// SchemaVersionMetadata under the schema versions. Hard deletes will be enforced
// on the database.
//
// If the compatibility mode forbids deleting of a version that is necessary, such
// as BACKWARDS_FULL, an error is returned.
