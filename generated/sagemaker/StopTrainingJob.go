package sagemaker

// StopTrainingJob is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Stops a training job. To stop a job, SageMaker sends the algorithm the SIGTERM
// signal, which delays job termination for 120 seconds. Algorithms might use this
// 120-second window to save the model artifacts, so the results of the training is
// not lost.
//
// When it receives a StopTrainingJob request, SageMaker changes the status of the
// job to Stopping . After SageMaker stops the job, it sets the status to Stopped .
