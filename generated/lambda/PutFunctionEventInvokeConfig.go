package lambda

// PutFunctionEventInvokeConfig is generated as a reference stub.
// Executable command wiring lives under cmd/lambda.go.
//
// Configures options for [asynchronous invocation] on a function, version, or alias. If a configuration
// already exists for a function, version, or alias, this operation overwrites it.
// If you exclude any settings, they are removed. To set one option without
// affecting existing settings for other options, use UpdateFunctionEventInvokeConfig.
//
// By default, Lambda retries an asynchronous invocation twice if the function
// returns an error. It retains events in a queue for up to six hours. When an
// event fails all processing attempts or stays in the asynchronous invocation
// queue for too long, Lambda discards it. To retain discarded events, configure a
// dead-letter queue with UpdateFunctionConfiguration.
//
// To send an invocation record to a queue, topic, S3 bucket, function, or event
// bus, specify a [destination]. You can configure separate destinations for successful
// invocations (on-success) and events that fail all processing attempts
// (on-failure). You can configure destinations in addition to or instead of a
// dead-letter queue.
//
// S3 buckets are supported only for on-failure destinations. To retain records of
// successful invocations, use another destination type.
//
// [asynchronous invocation]: https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html
// [destination]: https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#invocation-async-destinations
