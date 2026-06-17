package connect

// EvaluateDataTableValues is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Evaluates values at the time of the request and returns them. It considers the
// request's timezone or the table's timezone, in that order, when accessing time
// based tables. When a value is accessed, the accessor's identity and the time of
// access are saved alongside the value to help identify values that are actively
// in use. The term "Batch" is not included in the operation name since it does not
// meet all the criteria for a batch operation as specified in Batch Operations:
// Amazon Web Services API Standards.
