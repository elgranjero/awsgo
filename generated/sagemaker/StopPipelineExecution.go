package sagemaker

// StopPipelineExecution is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Stops a pipeline execution.
//
// # Callback Step
//
// A pipeline execution won't stop while a callback step is running. When you call
// StopPipelineExecution on a pipeline execution with a running callback step,
// SageMaker Pipelines sends an additional Amazon SQS message to the specified SQS
// queue. The body of the SQS message contains a "Status" field which is set to
// "Stopping".
//
// You should add logic to your Amazon SQS message consumer to take any needed
// action (for example, resource cleanup) upon receipt of the message followed by a
// call to SendPipelineExecutionStepSuccess or SendPipelineExecutionStepFailure .
//
// Only when SageMaker Pipelines receives one of these calls will it stop the
// pipeline execution.
//
// # Lambda Step
//
// A pipeline execution can't be stopped while a lambda step is running because
// the Lambda function invoked by the lambda step can't be stopped. If you attempt
// to stop the execution while the Lambda function is running, the pipeline waits
// for the Lambda function to finish or until the timeout is hit, whichever occurs
// first, and then stops. If the Lambda function finishes, the pipeline execution
// status is Stopped . If the timeout is hit the pipeline execution status is
// Failed .
