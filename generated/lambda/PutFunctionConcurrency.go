package lambda

// PutFunctionConcurrency is generated as a reference stub.
// Executable command wiring lives under cmd/lambda.go.
//
// Sets the maximum number of simultaneous executions for a function, and reserves
// capacity for that concurrency level.
//
// Concurrency settings apply to the function as a whole, including all published
// versions and the unpublished version. Reserving concurrency both ensures that
// your function has capacity to process the specified number of events
// simultaneously, and prevents it from scaling beyond that level. Use GetFunctionto see the
// current setting for a function.
//
// Use GetAccountSettings to see your Regional concurrency limit. You can reserve concurrency for as
// many functions as you like, as long as you leave at least 100 simultaneous
// executions unreserved for functions that aren't configured with a per-function
// limit. For more information, see [Lambda function scaling].
//
// [Lambda function scaling]: https://docs.aws.amazon.com/lambda/latest/dg/invocation-scaling.html
