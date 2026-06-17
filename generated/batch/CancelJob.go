package batch

// CancelJob is generated as a reference stub.
// Executable command wiring lives under cmd/batch.go.
//
// Cancels a job in an Batch job queue. Jobs that are in a SUBMITTED , PENDING , or
// RUNNABLE state are cancelled and the job status is updated to FAILED .
//
// A PENDING job is canceled after all dependency jobs are completed. Therefore,
// it may take longer than expected to cancel a job in PENDING status.
//
// When you try to cancel an array parent job in PENDING , Batch attempts to cancel
// all child jobs. The array parent job is canceled when all child jobs are
// completed.
//
// Jobs that progressed to the STARTING or RUNNING state aren't canceled. However,
// the API operation still succeeds, even if no job is canceled. These jobs must be
// terminated with the TerminateJoboperation.
