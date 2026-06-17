package sfn

// DescribeExecution is generated as a reference stub.
// Executable command wiring lives under cmd/sfn.go.
//
// Provides information about a state machine execution, such as the state machine
// associated with the execution, the execution input and output, and relevant
// execution metadata. If you've [redriven]an execution, you can use this API action to
// return information about the redrives of that execution. In addition, you can
// use this API action to return the Map Run Amazon Resource Name (ARN) if the
// execution was dispatched by a Map Run.
//
// If you specify a version or alias ARN when you call the StartExecution API action,
// DescribeExecution returns that ARN.
//
// This operation is eventually consistent. The results are best effort and may
// not reflect very recent updates and changes.
//
// Executions of an EXPRESS state machine aren't supported by DescribeExecution
// unless a Map Run dispatched them.
//
// [redriven]: https://docs.aws.amazon.com/step-functions/latest/dg/redrive-executions.html
