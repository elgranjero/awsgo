package greengrassv2

// CreateDeployment is generated as a reference stub.
// Executable command wiring lives under cmd/greengrassv2.go.
//
// Creates a continuous deployment for a target, which is a Greengrass core device
// or group of core devices. When you add a new core device to a group of core
// devices that has a deployment, IoT Greengrass deploys that group's deployment to
// the new device.
//
// You can define one deployment for each target. When you create a new deployment
// for a target that has an existing deployment, you replace the previous
// deployment. IoT Greengrass applies the new deployment to the target devices.
//
// Every deployment has a revision number that indicates how many deployment
// revisions you define for a target. Use this operation to create a new revision
// of an existing deployment.
//
// For more information, see the [Create deployments] in the IoT Greengrass V2 Developer Guide.
//
// [Create deployments]: https://docs.aws.amazon.com/greengrass/v2/developerguide/create-deployments.html
