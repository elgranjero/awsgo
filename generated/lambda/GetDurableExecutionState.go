package lambda

// GetDurableExecutionState is generated as a reference stub.
// Executable command wiring lives under cmd/lambda.go.
//
// Retrieves the current execution state required for the replay process during [durable function]
// execution. This API is used by the Lambda durable functions SDK to get state
// information needed for replay. You typically don't need to call this API
// directly as the SDK handles state management automatically.
//
// The response contains operations ordered by start sequence number in ascending
// order. Completed operations with children don't include child operation details
// since they don't need to be replayed.
//
// [durable function]: https://docs.aws.amazon.com/lambda/latest/dg/durable-functions.html
