package greengrassv2

// DeleteDeployment is generated as a reference stub.
// Executable command wiring lives under cmd/greengrassv2.go.
//
// Deletes a deployment. To delete an active deployment, you must first cancel it.
// For more information, see [CancelDeployment].
//
// Deleting a deployment doesn't affect core devices that run that deployment,
// because core devices store the deployment's configuration on the device.
// Additionally, core devices can roll back to a previous deployment that has been
// deleted.
//
// [CancelDeployment]: https://docs.aws.amazon.com/iot/latest/apireference/API_CancelDeployment.html
