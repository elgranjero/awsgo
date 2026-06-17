package lambda

// DeleteFunction is generated as a reference stub.
// Executable command wiring lives under cmd/lambda.go.
//
// Deletes a Lambda function. To delete a specific function version, use the
// Qualifier parameter. Otherwise, all versions and aliases are deleted. This
// doesn't require the user to have explicit permissions for DeleteAlias.
//
// A deleted Lambda function cannot be recovered. Ensure that you specify the
// correct function name and version before deleting.
//
// To delete Lambda event source mappings that invoke a function, use DeleteEventSourceMapping. For Amazon
// Web Services services and resources that invoke your function directly, delete
// the trigger in the service where you originally configured it.
