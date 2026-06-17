package sfn

// StartSyncExecution is generated as a reference stub.
// Executable command wiring lives under cmd/sfn.go.
//
// Starts a Synchronous Express state machine execution. StartSyncExecution is not
// available for STANDARD workflows.
//
// StartSyncExecution will return a 200 OK response, even if your execution fails,
// because the status code in the API response doesn't reflect function errors.
// Error codes are reserved for errors that prevent your execution from running,
// such as permissions errors, limit errors, or issues with your state machine code
// and configuration.
//
// This API action isn't logged in CloudTrail.
