package lambda

// GetDurableExecutionHistory is generated as a reference stub.
// Executable command wiring lives under cmd/lambda.go.
//
// Retrieves the execution history for a [durable execution], showing all the steps, callbacks, and
// events that occurred during the execution. This provides a detailed audit trail
// of the execution's progress over time.
//
// The history is available while the execution is running and for a retention
// period after it completes (1-90 days, default 30 days). You can control whether
// to include execution data such as step results and callback payloads.
//
// [durable execution]: https://docs.aws.amazon.com/lambda/latest/dg/durable-functions.html
