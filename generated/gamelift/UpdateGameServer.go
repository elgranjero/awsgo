package gamelift

// UpdateGameServer is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2 (FleetIQ)
//
// Updates information about a registered game server to help Amazon GameLift
// Servers FleetIQ track game server availability. This operation is called by a
// game server process that is running on an instance in a game server group.
//
// Use this operation to update the following types of game server information.
// You can make all three types of updates in the same request:
//
// - To update the game server's utilization status from AVAILABLE (when the game
// server is available to be claimed) to UTILIZED (when the game server is
// currently hosting games). Identify the game server and game server group and
// specify the new utilization status. You can't change the status from to
// UTILIZED to AVAILABLE .
//
// - To report health status, identify the game server and game server group and
// set health check to HEALTHY . If a game server does not report health status
// for a certain length of time, the game server is no longer considered healthy.
// As a result, it will be eventually deregistered from the game server group to
// avoid affecting utilization metrics. The best practice is to report health every
// 60 seconds.
//
// - To change game server metadata, provide updated game server data.
//
// Once a game server is successfully updated, the relevant statuses and
// timestamps are updated.
//
// # Learn more
//
// [Amazon GameLift Servers FleetIQ Guide]
//
// [Amazon GameLift Servers FleetIQ Guide]: https://docs.aws.amazon.com/gamelift/latest/fleetiqguide/gsg-intro.html
