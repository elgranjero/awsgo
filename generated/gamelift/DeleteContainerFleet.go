package gamelift

// DeleteContainerFleet is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: Container
//
// Deletes all resources and information related to a container fleet and shuts
// down currently running fleet instances, including those in remote locations. The
// container fleet must be in ACTIVE status to be deleted.
//
// To delete a fleet, specify the fleet ID to be terminated. During the deletion
// process, the fleet status is changed to DELETING .
//
// # Learn more
//
// [Setting up Amazon GameLift Servers Fleets]
//
// [Setting up Amazon GameLift Servers Fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html
