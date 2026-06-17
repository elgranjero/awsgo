package configservice

// BatchGetAggregateResourceConfig is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// Returns the current configuration items for resources that are present in your
// Config aggregator. The operation also returns a list of resources that are not
// processed in the current request. If there are no unprocessed resources, the
// operation returns an empty unprocessedResourceIdentifiers list.
//
// - The API does not return results for deleted resources.
//
// - The API does not return tags and relationships.
