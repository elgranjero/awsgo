package lambda

// InvokeAsync is generated as a reference stub.
// Executable command wiring lives under cmd/lambda.go.
//
// For asynchronous function invocation, use Invoke.
//
// Invokes a function asynchronously.
//
// The payload limit is 256KB. For larger payloads, for up to 1MB, use Invoke.
//
// If you do use the InvokeAsync action, note that it doesn't support the use of
// X-Ray active tracing. Trace ID is not propagated to the function, even if X-Ray
// active tracing is turned on.
//
// Deprecated: This operation has been deprecated.
