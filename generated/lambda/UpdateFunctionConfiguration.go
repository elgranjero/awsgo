package lambda

// UpdateFunctionConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/lambda.go.
//
// Modify the version-specific settings of a Lambda function.
//
// When you update a function, Lambda provisions an instance of the function and
// its supporting resources. If your function connects to a VPC, this process can
// take a minute. During this time, you can't modify the function, but you can
// still invoke it. The LastUpdateStatus , LastUpdateStatusReason , and
// LastUpdateStatusReasonCode fields in the response from GetFunctionConfiguration indicate when the
// update is complete and the function is processing events with the new
// configuration. For more information, see [Lambda function states].
//
// These settings can vary between versions of a function and are locked when you
// publish a version. You can't modify the configuration of a published version,
// only the unpublished version.
//
// To configure function concurrency, use PutFunctionConcurrency. To grant invoke permissions to an
// Amazon Web Services account or Amazon Web Services service, use AddPermission.
//
// [Lambda function states]: https://docs.aws.amazon.com/lambda/latest/dg/functions-states.html
