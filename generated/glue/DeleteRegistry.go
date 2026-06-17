package glue

// DeleteRegistry is generated as a reference stub.
// Executable command wiring lives under cmd/glue.go.
//
// Delete the entire registry including schema and all of its versions. To get the
// status of the delete operation, you can call the GetRegistry API after the
// asynchronous call. Deleting a registry will deactivate all online operations for
// the registry such as the UpdateRegistry , CreateSchema , UpdateSchema , and
// RegisterSchemaVersion APIs.
