package glue

// DeleteSchema is generated as a reference stub.
// Executable command wiring lives under cmd/glue.go.
//
// Deletes the entire schema set, including the schema set and all of its
// versions. To get the status of the delete operation, you can call GetSchema API
// after the asynchronous call. Deleting a registry will deactivate all online
// operations for the schema, such as the GetSchemaByDefinition , and
// RegisterSchemaVersion APIs.
