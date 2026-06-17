package apprunner

// PauseService is generated as a reference stub.
// Executable command wiring lives under cmd/apprunner.go.
//
// Pause an active App Runner service. App Runner reduces compute capacity for the
// service to zero and loses state (for example, ephemeral storage is removed).
//
// This is an asynchronous operation. On a successful call, you can use the
// returned OperationId and the ListOperations call to track the operation's progress.
