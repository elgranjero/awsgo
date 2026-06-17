package glue

// ListSchemas is generated as a reference stub.
// Executable command wiring lives under cmd/glue.go.
//
// Returns a list of schemas with minimal details. Schemas in Deleting status will
// not be included in the results. Empty results will be returned if there are no
// schemas available.
//
// When the RegistryId is not provided, all the schemas across registries will be
// part of the API response.
