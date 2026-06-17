package batch

// DeleteJobQueue is generated as a reference stub.
// Executable command wiring lives under cmd/batch.go.
//
// Deletes the specified job queue. You must first disable submissions for a queue
// with the UpdateJobQueueoperation. All jobs in the queue are eventually terminated when you
// delete a job queue. The jobs are terminated at a rate of about 16 jobs each
// second.
//
// It's not necessary to disassociate compute environments from a queue before
// submitting a DeleteJobQueue request.
