package gamelift

// UpdateGameServerGroup is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2 (FleetIQ)
//
// Updates Amazon GameLift Servers FleetIQ-specific properties for a game server
// group. Many Auto Scaling group properties are updated on the Auto Scaling group
// directly, including the launch template, Auto Scaling policies, and
// maximum/minimum/desired instance counts.
//
// To update the game server group, specify the game server group ID and provide
// the updated values. Before applying the updates, the new values are validated to
// ensure that Amazon GameLift Servers FleetIQ can continue to perform instance
// balancing activity. If successful, a GameServerGroup object is returned.
//
// # Learn more
//
// [Amazon GameLift Servers FleetIQ Guide]
//
// [Amazon GameLift Servers FleetIQ Guide]: https://docs.aws.amazon.com/gamelift/latest/fleetiqguide/gsg-intro.html
