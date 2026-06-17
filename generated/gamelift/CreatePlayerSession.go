package gamelift

// CreatePlayerSession is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Anywhere, Container
//
// Reserves an open player slot in a game session for a player. New player
// sessions can be created in any game session with an open slot that is in ACTIVE
// status and has a player creation policy of ACCEPT_ALL . You can add a group of
// players to a game session with [CreatePlayerSessions].
//
// To create a player session, specify a game session ID, player ID, and
// optionally a set of player data.
//
// If successful, a slot is reserved in the game session for the player and a new
// PlayerSessions object is returned with a player session ID. The player
// references the player session ID when sending a connection request to the game
// session, and the game server can use it to validate the player reservation with
// the Amazon GameLift Servers service. Player sessions cannot be updated.
//
// The maximum number of players per game session is 200. It is not adjustable.
//
// # Related actions
//
// [All APIs by task]
//
// [CreatePlayerSessions]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_CreatePlayerSessions.html
// [All APIs by task]: https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets
