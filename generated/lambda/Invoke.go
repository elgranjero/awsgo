package lambda

// Invoke is generated as a reference stub.
// Executable command wiring lives under cmd/lambda.go.
//
// Invokes a Lambda function. You can invoke a function synchronously (and wait
// for the response), or asynchronously. By default, Lambda invokes your function
// synchronously (i.e. the InvocationType is RequestResponse ). To invoke a
// function asynchronously, set InvocationType to Event . Lambda passes the
// ClientContext object to your function for synchronous invocations only.
//
// For synchronous invocations, the maximum payload size is 6 MB. For asynchronous
// invocations, the maximum payload size is 1 MB.
//
// For [synchronous invocation], details about the function response, including errors, are included in
// the response body and headers. For either invocation type, you can find more
// information in the [execution log]and [trace].
//
// When an error occurs, your function may be invoked multiple times. Retry
// behavior varies by error type, client, event source, and invocation type. For
// example, if you invoke a function asynchronously and it returns an error, Lambda
// executes the function up to two more times. For more information, see [Error handling and automatic retries in Lambda].
//
// For [asynchronous invocation], Lambda adds events to a queue before sending them to your function. If
// your function does not have enough capacity to keep up with the queue, events
// may be lost. Occasionally, your function may receive the same event multiple
// times, even if no error occurs. To retain events that were not processed,
// configure your function with a [dead-letter queue].
//
// The status code in the API response doesn't reflect function errors. Error
// codes are reserved for errors that prevent your function from executing, such as
// permissions errors, [quota]errors, or issues with your function's code and
// configuration. For example, Lambda returns TooManyRequestsException if running
// the function would cause you to exceed a concurrency limit at either the account
// level ( ConcurrentInvocationLimitExceeded ) or function level (
// ReservedFunctionConcurrentInvocationLimitExceeded ).
//
// For functions with a long timeout, your client might disconnect during
// synchronous invocation while it waits for a response. Configure your HTTP
// client, SDK, firewall, proxy, or operating system to allow for long connections
// with timeout or keep-alive settings.
//
// This operation requires permission for the [lambda:InvokeFunction] action. For details on how to set
// up permissions for cross-account invocations, see [Granting function access to other accounts].
//
// [execution log]: https://docs.aws.amazon.com/lambda/latest/dg/monitoring-functions.html
// [asynchronous invocation]: https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html
// [trace]: https://docs.aws.amazon.com/lambda/latest/dg/lambda-x-ray.html
// [dead-letter queue]: https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#invocation-dlq
// [Error handling and automatic retries in Lambda]: https://docs.aws.amazon.com/lambda/latest/dg/invocation-retries.html
// [lambda:InvokeFunction]: https://docs.aws.amazon.com/IAM/latest/UserGuide/list_awslambda.html
// [quota]: https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-limits.html
// [synchronous invocation]: https://docs.aws.amazon.com/lambda/latest/dg/invocation-sync.html
// [Granting function access to other accounts]: https://docs.aws.amazon.com/lambda/latest/dg/access-control-resource-based.html#permissions-resource-xaccountinvoke
