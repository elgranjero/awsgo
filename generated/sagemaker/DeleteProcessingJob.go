package sagemaker

// DeleteProcessingJob is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Deletes a processing job. After Amazon SageMaker deletes a processing job, all
// of the metadata for the processing job is lost. You can delete only processing
// jobs that are in a terminal state ( Stopped , Failed , or Completed ). You
// cannot delete a job that is in the InProgress or Stopping state. After deleting
// the job, you can reuse its name to create another processing job.
