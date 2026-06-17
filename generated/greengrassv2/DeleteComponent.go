package greengrassv2

// DeleteComponent is generated as a reference stub.
// Executable command wiring lives under cmd/greengrassv2.go.
//
// Deletes a version of a component from IoT Greengrass.
//
// This operation deletes the component's recipe and artifacts. As a result,
// deployments that refer to this component version will fail. If you have
// deployments that use this component version, you can remove the component from
// the deployment or update the deployment to use a valid version.
