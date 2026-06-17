package batch

// CreateJobQueue is generated as a reference stub.
// Executable command wiring lives under cmd/batch.go.
//
// Creates an Batch job queue. When you create a job queue, you associate one or
// more compute environments to the queue and assign an order of preference for the
// compute environments.
//
// You also set a priority to the job queue that determines the order that the
// Batch scheduler places jobs onto its associated compute environments. For
// example, if a compute environment is associated with more than one job queue,
// the job queue with a higher priority is given preference for scheduling jobs to
// that compute environment.
