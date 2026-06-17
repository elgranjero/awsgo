package gamelift

// DescribePlayerSessions is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Anywhere, Container
//
// Retrieves properties for one or more player sessions.
//
// This action can be used in the following ways:
//
// - To retrieve a specific player session, provide the player session ID only.
//
// - To retrieve all player sessions in a game session, provide the game session
// ID only.
//
// - To retrieve all player sessions for a specific player, provide a player ID
// only.
//
// To request player sessions, specify either a player session ID, game session
// ID, or player ID. You can filter this request by player session status. If you
// provide a specific PlayerSessionId or PlayerId , Amazon GameLift Servers ignores
// the filter criteria. Use the pagination parameters to retrieve results as a set
// of sequential pages.
//
// If successful, a PlayerSession object is returned for each session that matches
// the request.
//
// # Related actions
//
// [All APIs by task]
//
// [All APIs by task]: https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets
