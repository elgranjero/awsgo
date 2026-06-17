package athena

// ImportNotebook is generated as a reference stub.
// Executable command wiring lives under cmd/athena.go.
//
// Imports a single ipynb file to a Spark enabled workgroup. To import the
// notebook, the request must specify a value for either Payload or
// NoteBookS3LocationUri . If neither is specified or both are specified, an
// InvalidRequestException occurs. The maximum file size that can be imported is 10
// megabytes. If an ipynb file with the same name already exists in the workgroup,
// throws an error.
