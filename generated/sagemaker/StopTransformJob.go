package sagemaker

// StopTransformJob is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Stops a batch transform job.
//
// When Amazon SageMaker receives a StopTransformJob request, the status of the
// job changes to Stopping . After Amazon SageMaker stops the job, the status is
// set to Stopped . When you stop a batch transform job before it is completed,
// Amazon SageMaker doesn't store the job's output in Amazon S3.
