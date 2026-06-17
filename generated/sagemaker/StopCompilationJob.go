package sagemaker

// StopCompilationJob is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Stops a model compilation job.
//
// To stop a job, Amazon SageMaker AI sends the algorithm the SIGTERM signal. This
// gracefully shuts the job down. If the job hasn't stopped, it sends the SIGKILL
// signal.
//
// When it receives a StopCompilationJob request, Amazon SageMaker AI changes the
// CompilationJobStatus of the job to Stopping . After Amazon SageMaker stops the
// job, it sets the CompilationJobStatus to Stopped .
