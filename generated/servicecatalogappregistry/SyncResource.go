package servicecatalogappregistry

// SyncResource is generated as a reference stub.
// Executable command wiring lives under cmd/servicecatalogappregistry.go.
//
// Syncs the resource with current AppRegistry records.
//
// Specifically, the resource’s AppRegistry system tags sync with its associated
// application. We remove the resource's AppRegistry system tags if it does not
// associate with the application. The caller must have permissions to read and
// update the resource.
