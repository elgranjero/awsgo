package datazone

// CancelMetadataGenerationRun is generated as a reference stub.
// Executable command wiring lives under cmd/datazone.go.
//
// Cancels the metadata generation run.
//
// Prerequisites:
//
// - The run must exist and be in a cancelable status (e.g., SUBMITTED,
// IN_PROGRESS).
//
// - Runs in SUCCEEDED status cannot be cancelled.
//
// - User must have access to the run and cancel permissions.
