package omics

// DeleteRun is generated as a reference stub.
// Executable command wiring lives under cmd/omics.go.
//
// Deletes a run and returns a response with no body if the operation is
// successful. You can only delete a run that has reached a COMPLETED , FAILED , or
// CANCELLED stage. A completed run has delivered an output, or was cancelled and
// resulted in no output. When you delete a run, only the metadata associated with
// the run is deleted. The run outputs remain in Amazon S3 and logs remain in
// CloudWatch.
//
// To verify that the workflow is deleted:
//
// - Use ListRuns to confirm the workflow no longer appears in the list.
//
// - Use GetRun to verify the workflow cannot be found.
