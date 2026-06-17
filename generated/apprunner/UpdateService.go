package apprunner

// UpdateService is generated as a reference stub.
// Executable command wiring lives under cmd/apprunner.go.
//
// Update an App Runner service. You can update the source configuration and
// instance configuration of the service. You can also update the ARN of the auto
// scaling configuration resource that's associated with the service. However, you
// can't change the name or the encryption configuration of the service. These can
// be set only when you create the service.
//
// To update the tags applied to your service, use the separate actions TagResource and UntagResource.
//
// This is an asynchronous operation. On a successful call, you can use the
// returned OperationId and the ListOperations call to track the operation's progress.
