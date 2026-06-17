package storagegateway

// CancelCacheReport is generated as a reference stub.
// Executable command wiring lives under cmd/storagegateway.go.
//
// Cancels generation of a specified cache report. You can use this operation to
// manually cancel an IN-PROGRESS report for any reason. This action changes the
// report status from IN-PROGRESS to CANCELLED. You can only cancel in-progress
// reports. If the the report you attempt to cancel is in FAILED, ERROR, or
// COMPLETED state, the cancel operation returns an error.
