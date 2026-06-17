package lambda

// PublishVersion is generated as a reference stub.
// Executable command wiring lives under cmd/lambda.go.
//
// Creates a [version] from the current code and configuration of a function. Use versions
// to create a snapshot of your function code and configuration that doesn't
// change.
//
// Lambda doesn't publish a version if the function's configuration and code
// haven't changed since the last version. Use UpdateFunctionCodeor UpdateFunctionConfiguration to update the function before
// publishing a version.
//
// Clients can invoke versions directly or with an alias. To create an alias, use CreateAlias.
//
// [version]: https://docs.aws.amazon.com/lambda/latest/dg/versioning-aliases.html
