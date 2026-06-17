package gamelift

// ResumeGameServerGroup is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2 (FleetIQ)
//
// Reinstates activity on a game server group after it has been suspended. A game
// server group might be suspended by the SuspendGameServerGroupoperation, or it might be suspended
// involuntarily due to a configuration problem. In the second case, you can
// manually resume activity on the group once the configuration problem has been
// resolved. Refer to the game server group status and status reason for more
// information on why group activity is suspended.
//
// To resume activity, specify a game server group ARN and the type of activity to
// be resumed. If successful, a GameServerGroup object is returned showing that
// the resumed activity is no longer listed in SuspendedActions .
//
// # Learn more
//
// [Amazon GameLift Servers FleetIQ Guide]
//
// [Amazon GameLift Servers FleetIQ Guide]: https://docs.aws.amazon.com/gamelift/latest/fleetiqguide/gsg-intro.html
