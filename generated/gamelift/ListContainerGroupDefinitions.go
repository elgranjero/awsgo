package gamelift

// ListContainerGroupDefinitions is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: Container
//
// Retrieves container group definitions for the Amazon Web Services account and
// Amazon Web Services Region. Use the pagination parameters to retrieve results in
// a set of sequential pages.
//
// This operation returns only the latest version of each definition. To retrieve
// all versions of a container group definition, use [ListContainerGroupDefinitionVersions].
//
// Request options:
//
// - Retrieve the most recent versions of all container group definitions.
//
// - Retrieve the most recent versions of all container group definitions,
// filtered by type. Specify the container group type to filter on.
//
// Results:
//
// If successful, this operation returns the complete properties of a set of
// container group definition versions that match the request.
//
// This operation returns the list of container group definitions in no particular
// order.
//
// [ListContainerGroupDefinitionVersions]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_ListContainerGroupDefinitionVersions.html
