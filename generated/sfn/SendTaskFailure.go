package sfn

// SendTaskFailure is generated as a reference stub.
// Executable command wiring lives under cmd/sfn.go.
//
// Used by activity workers, Task states using the [callback] pattern, and optionally Task
// states using the [job run]pattern to report that the task identified by the taskToken
// failed.
//
// For an execution with encryption enabled, Step Functions will encrypt the error
// and cause fields using the KMS key for the execution role.
//
// A caller can mark a task as fail without using any KMS permissions in the
// execution role if the caller provides a null value for both error and cause
// fields because no data needs to be encrypted.
//
// [callback]: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
// [job run]: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-sync
