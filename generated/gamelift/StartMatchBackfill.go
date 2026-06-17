package gamelift

// StartMatchBackfill is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Anywhere, Container
//
// Finds new players to fill open slots in currently running game sessions. The
// backfill match process is essentially identical to the process of forming new
// matches. Backfill requests use the same matchmaker that was used to make the
// original match, and they provide matchmaking data for all players currently in
// the game session. FlexMatch uses this information to select new players so that
// backfilled match continues to meet the original match requirements.
//
// When using FlexMatch with Amazon GameLift Servers managed hosting, you can
// request a backfill match from a client service by calling this operation with a
// GameSessions ID. You also have the option of making backfill requests directly
// from your game server. In response to a request, FlexMatch creates player
// sessions for the new players, updates the GameSession resource, and sends
// updated matchmaking data to the game server. You can request a backfill match at
// any point after a game session is started. Each game session can have only one
// active backfill request at a time; a subsequent request automatically replaces
// the earlier request.
//
// When using FlexMatch as a standalone component, request a backfill match by
// calling this operation without a game session identifier. As with newly formed
// matches, matchmaking results are returned in a matchmaking event so that your
// game can update the game session that is being backfilled.
//
// To request a backfill match, specify a unique ticket ID, the original
// matchmaking configuration, and matchmaking data for all current players in the
// game session being backfilled. Optionally, specify the GameSession ARN. If
// successful, a match backfill ticket is created and returned with status set to
// QUEUED. Track the status of backfill tickets using the same method for tracking
// tickets for new matches.
//
// Only game sessions created by FlexMatch are supported for match backfill.
//
// # Learn more
//
// [Backfill existing games with FlexMatch]
//
// [Matchmaking events](reference)
//
// [How Amazon GameLift Servers FlexMatch works]
//
// [Matchmaking events]: https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-events.html
// [Backfill existing games with FlexMatch]: https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-backfill.html
// [How Amazon GameLift Servers FlexMatch works]: https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/gamelift-match.html
