package lambda

// CheckpointDurableExecution is generated as a reference stub.
// Executable command wiring lives under cmd/lambda.go.
//
// Saves the progress of a [durable function] execution during runtime. This API is used by the
// Lambda durable functions SDK to checkpoint completed steps and schedule
// asynchronous operations. You typically don't need to call this API directly as
// the SDK handles checkpointing automatically.
//
// Each checkpoint operation consumes the current checkpoint token and returns a
// new one for the next checkpoint. This ensures that checkpoints are applied in
// the correct order and prevents duplicate or out-of-order state updates.
//
// [durable function]: https://docs.aws.amazon.com/lambda/latest/dg/durable-functions.html
