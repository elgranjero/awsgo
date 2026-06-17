package lambda

// ListFunctions is generated as a reference stub.
// Executable command wiring lives under cmd/lambda.go.
//
// Returns a list of Lambda functions, with the version-specific configuration of
// each. Lambda returns up to 50 functions per call.
//
// Set FunctionVersion to ALL to include all published versions of each function
// in addition to the unpublished version.
//
// The ListFunctions operation returns a subset of the FunctionConfiguration fields. To get the
// additional fields (State, StateReasonCode, StateReason, LastUpdateStatus,
// LastUpdateStatusReason, LastUpdateStatusReasonCode, RuntimeVersionConfig) for a
// function or version, use GetFunction.
