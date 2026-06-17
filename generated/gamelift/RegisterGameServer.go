package gamelift

// RegisterGameServer is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2 (FleetIQ)
//
// Creates a new game server resource and notifies Amazon GameLift Servers FleetIQ
// that the game server is ready to host gameplay and players. This operation is
// called by a game server process that is running on an instance in a game server
// group. Registering game servers enables Amazon GameLift Servers FleetIQ to track
// available game servers and enables game clients and services to claim a game
// server for a new game session.
//
// To register a game server, identify the game server group and instance where
// the game server is running, and provide a unique identifier for the game server.
// You can also include connection and game server data.
//
// Once a game server is successfully registered, it is put in status AVAILABLE . A
// request to register a game server may fail if the instance it is running on is
// in the process of shutting down as part of instance balancing or scale-down
// activity.
//
// # Learn more
//
// [Amazon GameLift Servers FleetIQ Guide]
//
// [Amazon GameLift Servers FleetIQ Guide]: https://docs.aws.amazon.com/gamelift/latest/fleetiqguide/gsg-intro.html
