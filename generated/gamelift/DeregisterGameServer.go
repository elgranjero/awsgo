package gamelift

// DeregisterGameServer is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2 (FleetIQ)
//
// Removes the game server from a game server group. As a result of this
// operation, the deregistered game server can no longer be claimed and will not be
// returned in a list of active game servers.
//
// To deregister a game server, specify the game server group and game server ID.
// If successful, this operation emits a CloudWatch event with termination
// timestamp and reason.
//
// # Learn more
//
// [Amazon GameLift Servers FleetIQ Guide]
//
// [Amazon GameLift Servers FleetIQ Guide]: https://docs.aws.amazon.com/gamelift/latest/fleetiqguide/gsg-intro.html
