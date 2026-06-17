package appintegrations

// DeleteDataIntegration is generated as a reference stub.
// Executable command wiring lives under cmd/appintegrations.go.
//
// Deletes the DataIntegration. Only DataIntegrations that don't have any
// DataIntegrationAssociations can be deleted. Deleting a DataIntegration also
// deletes the underlying Amazon AppFlow flow and service linked role.
//
// You cannot create a DataIntegration association for a DataIntegration that has
// been previously associated. Use a different DataIntegration, or recreate the
// DataIntegration using the [CreateDataIntegration]API.
//
// [CreateDataIntegration]: https://docs.aws.amazon.com/appintegrations/latest/APIReference/API_CreateDataIntegration.html
