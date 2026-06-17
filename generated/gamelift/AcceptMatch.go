package gamelift

// AcceptMatch is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Anywhere, Container
//
// Registers a player's acceptance or rejection of a proposed FlexMatch match. A
// matchmaking configuration may require player acceptance; if so, then matches
// built with that configuration cannot be completed unless all players accept the
// proposed match within a specified time limit.
//
// When FlexMatch builds a match, all the matchmaking tickets involved in the
// proposed match are placed into status REQUIRES_ACCEPTANCE . This is a trigger
// for your game to get acceptance from all players in each ticket. Calls to this
// action are only valid for tickets that are in this status; calls for tickets not
// in this status result in an error.
//
// To register acceptance, specify the ticket ID, one or more players, and an
// acceptance response. When all players have accepted, Amazon GameLift Servers
// advances the matchmaking tickets to status PLACING , and attempts to create a
// new game session for the match.
//
// If any player rejects the match, or if acceptances are not received before a
// specified timeout, the proposed match is dropped. Each matchmaking ticket in the
// failed match is handled as follows:
//
// - If the ticket has one or more players who rejected the match or failed to
// respond, the ticket status is set CANCELLED and processing is terminated.
//
// - If all players in the ticket accepted the match, the ticket status is
// returned to SEARCHING to find a new match.
//
// # Learn more
//
// [Add FlexMatch to a game client]
//
// [FlexMatch events](reference)
//
// [FlexMatch events]: https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-events.html
// [Add FlexMatch to a game client]: https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-client.html
