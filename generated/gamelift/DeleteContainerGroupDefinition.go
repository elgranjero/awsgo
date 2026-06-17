package gamelift

// DeleteContainerGroupDefinition is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: Container
//
// Request options:
//
// Deletes a container group definition.
//
// - Delete an entire container group definition, including all versions.
// Specify the container group definition name, or use an ARN value without the
// version number.
//
// - Delete a particular version. Specify the container group definition name
// and a version number, or use an ARN value that includes the version number.
//
// - Keep the newest versions and delete all older versions. Specify the
// container group definition name and the number of versions to retain. For
// example, set VersionCountToRetain to 5 to delete all but the five most recent
// versions.
//
// # Result
//
// If successful, Amazon GameLift Servers removes the container group definition
// versions that you request deletion for. This request will fail for any requested
// versions if the following is true:
//
// - If the version is being used in an active fleet
//
// - If the version is being deployed to a fleet in a deployment that's
// currently in progress.
//
// - If the version is designated as a rollback definition in a fleet deployment
// that's currently in progress.
//
// # Learn more
//
// [Manage a container group definition]
//
// [Manage a container group definition]: https://docs.aws.amazon.com/gamelift/latest/developerguide/containers-create-groups.html
