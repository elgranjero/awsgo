package sfn

// StopExecution is generated as a reference stub.
// Executable command wiring lives under cmd/sfn.go.
//
// Stops an execution.
//
// This API action is not supported by EXPRESS state machines.
//
// For an execution with encryption enabled, Step Functions will encrypt the error
// and cause fields using the KMS key for the execution role.
//
// A caller can stop an execution without using any KMS permissions in the
// execution role if the caller provides a null value for both error and cause
// fields because no data needs to be encrypted.
