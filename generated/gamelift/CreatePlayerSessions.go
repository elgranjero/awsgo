package gamelift

// CreatePlayerSessions is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Anywhere, Container
//
// Reserves open slots in a game session for a group of players. New player
// sessions can be created in any game session with an open slot that is in ACTIVE
// status and has a player creation policy of ACCEPT_ALL . To add a single player
// to a game session, use [CreatePlayerSession]
//
// To create player sessions, specify a game session ID and a list of player IDs.
// Optionally, provide a set of player data for each player ID.
//
// If successful, a slot is reserved in the game session for each player, and new
// PlayerSession objects are returned with player session IDs. Each player
// references their player session ID when sending a connection request to the game
// session, and the game server can use it to validate the player reservation with
// the Amazon GameLift Servers service. Player sessions cannot be updated.
//
// The maximum number of players per game session is 200. It is not adjustable.
//
// # Related actions
//
// [All APIs by task]
//
// [CreatePlayerSession]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_CreatePlayerSession.html
// [All APIs by task]: https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets
