package configservice

// BatchGetResourceConfig is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// Returns the BaseConfigurationItem for one or more requested resources. The
// operation also returns a list of resources that are not processed in the current
// request. If there are no unprocessed resources, the operation returns an empty
// unprocessedResourceKeys list.
//
// - The API does not return results for deleted resources.
//
// - The API does not return any tags for the requested resources. This
// information is filtered out of the supplementaryConfiguration section of the API
// response.
