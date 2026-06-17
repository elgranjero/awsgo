package batch

// DeleteComputeEnvironment is generated as a reference stub.
// Executable command wiring lives under cmd/batch.go.
//
// Deletes an Batch compute environment.
//
// Before you can delete a compute environment, you must set its state to DISABLED
// with the UpdateComputeEnvironmentAPI operation and disassociate it from any job queues with the UpdateJobQueue API
// operation. Compute environments that use Fargate resources must terminate all
// active jobs on that compute environment before deleting the compute environment.
// If this isn't done, the compute environment enters an invalid state.
