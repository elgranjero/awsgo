package gamelift

// UpdateContainerGroupDefinition is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: Container
//
// Updates properties in an existing container group definition. This operation
// doesn't replace the definition. Instead, it creates a new version of the
// definition and saves it separately. You can access all versions that you choose
// to retain.
//
// The only property you can't update is the container group type.
//
// Request options:
//
// - Update based on the latest version of the container group definition.
// Specify the container group definition name only, or use an ARN value without a
// version number. Provide updated values for the properties that you want to
// change only. All other values remain the same as the latest version.
//
// - Update based on a specific version of the container group definition.
// Specify the container group definition name and a source version number, or use
// an ARN value with a version number. Provide updated values for the properties
// that you want to change only. All other values remain the same as the source
// version.
//
// - Change a game server container definition. Provide the updated container
// definition.
//
// - Add or change a support container definition. Provide a complete set of
// container definitions, including the updated definition.
//
// - Remove a support container definition. Provide a complete set of container
// definitions, excluding the definition to remove. If the container group has only
// one support container definition, provide an empty set.
//
// Results:
//
// If successful, this operation returns the complete properties of the new
// container group definition version.
//
// If the container group definition version is used in an active fleets, the
// update automatically initiates a new fleet deployment of the new version. You
// can track a fleet's deployments using [ListFleetDeployments].
//
// [ListFleetDeployments]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_ListFleetDeployments.html
