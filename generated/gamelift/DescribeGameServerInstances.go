package gamelift

// DescribeGameServerInstances is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2 (FleetIQ)
//
// Retrieves status information about the Amazon EC2 instances associated with a
// Amazon GameLift Servers FleetIQ game server group. Use this operation to detect
// when instances are active or not available to host new game servers.
//
// To request status for all instances in the game server group, provide a game
// server group ID only. To request status for specific instances, provide the game
// server group ID and one or more instance IDs. Use the pagination parameters to
// retrieve results in sequential segments. If successful, a collection of
// GameServerInstance objects is returned.
//
// This operation is not designed to be called with every game server claim
// request; this practice can cause you to exceed your API limit, which results in
// errors. Instead, as a best practice, cache the results and refresh your cache no
// more than once every 10 seconds.
//
// # Learn more
//
// [Amazon GameLift Servers FleetIQ Guide]
//
// [Amazon GameLift Servers FleetIQ Guide]: https://docs.aws.amazon.com/gamelift/latest/fleetiqguide/gsg-intro.html
