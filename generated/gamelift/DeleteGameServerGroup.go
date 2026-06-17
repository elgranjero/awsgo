package gamelift

// DeleteGameServerGroup is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2 (FleetIQ)
//
// Terminates a game server group and permanently deletes the game server group
// record. You have several options for how these resources are impacted when
// deleting the game server group. Depending on the type of delete operation
// selected, this operation might affect these resources:
//
// - The game server group
//
// - The corresponding Auto Scaling group
//
// - All game servers that are currently running in the group
//
// To delete a game server group, identify the game server group to delete and
// specify the type of delete operation to initiate. Game server groups can only be
// deleted if they are in ACTIVE or ERROR status.
//
// If the delete request is successful, a series of operations are kicked off. The
// game server group status is changed to DELETE_SCHEDULED , which prevents new
// game servers from being registered and stops automatic scaling activity. Once
// all game servers in the game server group are deregistered, Amazon GameLift
// Servers FleetIQ can begin deleting resources. If any of the delete operations
// fail, the game server group is placed in ERROR status.
//
// Amazon GameLift Servers FleetIQ emits delete events to Amazon CloudWatch.
//
// # Learn more
//
// [Amazon GameLift Servers FleetIQ Guide]
//
// [Amazon GameLift Servers FleetIQ Guide]: https://docs.aws.amazon.com/gamelift/latest/fleetiqguide/gsg-intro.html
