package sagemaker

// SendPipelineExecutionStepFailure is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Notifies the pipeline that the execution of a callback step failed, along with
// a message describing why. When a callback step is run, the pipeline generates a
// callback token and includes the token in a message sent to Amazon Simple Queue
// Service (Amazon SQS).
