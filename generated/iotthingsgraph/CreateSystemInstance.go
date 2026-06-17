package iotthingsgraph

// CreateSystemInstance is generated as a reference stub.
// Executable command wiring lives under cmd/iotthingsgraph.go.
//
// Creates a system instance.
//
// This action validates the system instance, prepares the deployment-related
// resources. For Greengrass deployments, it updates the Greengrass group that is
// specified by the greengrassGroupName parameter. It also adds a file to the S3
// bucket specified by the s3BucketName parameter. You need to call
// DeploySystemInstance after running this action.
//
// For Greengrass deployments, since this action modifies and adds resources to a
// Greengrass group and an S3 bucket on the caller's behalf, the calling identity
// must have write permissions to both the specified Greengrass group and S3
// bucket. Otherwise, the call will fail with an authorization error.
//
// For cloud deployments, this action requires a flowActionsRoleArn value. This is
// an IAM role that has permissions to access AWS services, such as AWS Lambda and
// AWS IoT, that the flow uses when it executes.
//
// If the definition document doesn't specify a version of the user's namespace,
// the latest version will be used by default.
//
// Deprecated: since: 2022-08-30
