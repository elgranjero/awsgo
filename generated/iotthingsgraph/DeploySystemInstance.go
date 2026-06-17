package iotthingsgraph

// DeploySystemInstance is generated as a reference stub.
// Executable command wiring lives under cmd/iotthingsgraph.go.
//
// Greengrass and Cloud Deployments
//
// Deploys the system instance to the target specified in CreateSystemInstance .
//
// # Greengrass Deployments
//
// If the system or any workflows and entities have been updated before this
// action is called, then the deployment will create a new Amazon Simple Storage
// Service resource file and then deploy it.
//
// Since this action creates a Greengrass deployment on the caller's behalf, the
// calling identity must have write permissions to the specified Greengrass group.
// Otherwise, the call will fail with an authorization error.
//
// For information about the artifacts that get added to your Greengrass core
// device when you use this API, see [AWS IoT Things Graph and AWS IoT Greengrass].
//
// Deprecated: since: 2022-08-30
//
// [AWS IoT Things Graph and AWS IoT Greengrass]: https://docs.aws.amazon.com/thingsgraph/latest/ug/iot-tg-greengrass.html
