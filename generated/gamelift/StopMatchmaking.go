package gamelift

// StopMatchmaking is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Anywhere, Container
//
// Cancels a matchmaking ticket or match backfill ticket that is currently being
// processed. To stop the matchmaking operation, specify the ticket ID. If
// successful, work on the ticket is stopped, and the ticket status is changed to
// CANCELLED .
//
// This call is also used to turn off automatic backfill for an individual game
// session. This is for game sessions that are created with a matchmaking
// configuration that has automatic backfill enabled. The ticket ID is included in
// the MatchmakerData of an updated game session object, which is provided to the
// game server.
//
// If the operation is successful, the service sends back an empty JSON struct
// with the HTTP 200 response (not an empty HTTP body).
//
// # Learn more
//
// [Add FlexMatch to a game client]
//
// [Add FlexMatch to a game client]: https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-client.html
