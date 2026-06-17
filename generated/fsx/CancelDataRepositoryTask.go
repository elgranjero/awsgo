package fsx

// CancelDataRepositoryTask is generated as a reference stub.
// Executable command wiring lives under cmd/fsx.go.
//
// Cancels an existing Amazon FSx for Lustre data repository task if that task is
// in either the PENDING or EXECUTING state. When you cancel an export task,
// Amazon FSx does the following.
//
// - Any files that FSx has already exported are not reverted.
//
// - FSx continues to export any files that are in-flight when the cancel
// operation is received.
//
// - FSx does not export any files that have not yet been exported.
//
// For a release task, Amazon FSx will stop releasing files upon cancellation. Any
// files that have already been released will remain in the released state.
