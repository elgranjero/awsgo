package lambda

// PutFunctionRecursionConfig is generated as a reference stub.
// Executable command wiring lives under cmd/lambda.go.
//
// Sets your function's [recursive loop detection] configuration.
//
// When you configure a Lambda function to output to the same service or resource
// that invokes the function, it's possible to create an infinite recursive loop.
// For example, a Lambda function might write a message to an Amazon Simple Queue
// Service (Amazon SQS) queue, which then invokes the same function. This
// invocation causes the function to write another message to the queue, which in
// turn invokes the function again.
//
// Lambda can detect certain types of recursive loops shortly after they occur.
// When Lambda detects a recursive loop and your function's recursive loop
// detection configuration is set to Terminate , it stops your function being
// invoked and notifies you.
//
// [recursive loop detection]: https://docs.aws.amazon.com/lambda/latest/dg/invocation-recursion.html
