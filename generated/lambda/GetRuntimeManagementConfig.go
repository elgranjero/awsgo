package lambda

// GetRuntimeManagementConfig is generated as a reference stub.
// Executable command wiring lives under cmd/lambda.go.
//
// Retrieves the runtime management configuration for a function's version. If the
// runtime update mode is Manual, this includes the ARN of the runtime version and
// the runtime update mode. If the runtime update mode is Auto or Function update,
// this includes the runtime update mode and null is returned for the ARN. For
// more information, see [Runtime updates].
//
// [Runtime updates]: https://docs.aws.amazon.com/lambda/latest/dg/runtimes-update.html
