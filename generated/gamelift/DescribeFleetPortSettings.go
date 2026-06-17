package gamelift

// DescribeFleetPortSettings is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Container
//
// Retrieves a fleet's inbound connection permissions. Connection permissions
// specify IP addresses and port settings that incoming traffic can use to access
// server processes in the fleet. Game server processes that are running in the
// fleet must use a port that falls within this range.
//
// Use this operation in the following ways:
//
// - To retrieve the port settings for a fleet, identify the fleet's unique
// identifier.
//
// - To check the status of recent updates to a fleet remote location, specify
// the fleet ID and a location. Port setting updates can take time to propagate
// across all locations.
//
// If successful, a set of IpPermission objects is returned for the requested
// fleet ID. When specifying a location, this operation returns a pending status.
// If the requested fleet has been deleted, the result set is empty.
//
// # Learn more
//
// [Setting up Amazon GameLift Servers fleets]
//
// [Setting up Amazon GameLift Servers fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html
