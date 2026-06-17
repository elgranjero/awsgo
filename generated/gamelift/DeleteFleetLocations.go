package gamelift

// DeleteFleetLocations is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Container
//
// Removes locations from a multi-location fleet. When deleting a location, all
// game server process and all instances that are still active in the location are
// shut down.
//
// To delete fleet locations, identify the fleet ID and provide a list of the
// locations to be deleted.
//
// If successful, GameLift sets the location status to DELETING , and begins to
// shut down existing server processes and terminate instances in each location
// being deleted. When completed, the location status changes to TERMINATED .
//
// # Learn more
//
// [Setting up Amazon GameLift Servers fleets]
//
// [Setting up Amazon GameLift Servers fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html
