package gamelift

// ListContainerGroupDefinitionVersions is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: Container
//
// Retrieves all versions of a container group definition. Use the pagination
// parameters to retrieve results in a set of sequential pages.
//
// Request options:
//
// - Get all versions of a specified container group definition. Specify the
// container group definition name or ARN value. (If the ARN value has a version
// number, it's ignored.)
//
// Results:
//
// If successful, this operation returns the complete properties of a set of
// container group definition versions that match the request.
//
// This operation returns the list of container group definitions in descending
// version order (latest first).
//
// # Learn more
//
// [Manage a container group definition]
//
// [Manage a container group definition]: https://docs.aws.amazon.com/gamelift/latest/developerguide/containers-create-groups.html
