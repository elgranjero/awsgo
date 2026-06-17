package apprunner

// StartDeployment is generated as a reference stub.
// Executable command wiring lives under cmd/apprunner.go.
//
// Initiate a manual deployment of the latest commit in a source code repository
// or the latest image in a source image repository to an App Runner service.
//
// For a source code repository, App Runner retrieves the commit and builds a
// Docker image. For a source image repository, App Runner retrieves the latest
// Docker image. In both cases, App Runner then deploys the new image to your
// service and starts a new container instance.
//
// This is an asynchronous operation. On a successful call, you can use the
// returned OperationId and the ListOperations call to track the operation's progress.
