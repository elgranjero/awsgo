package gamelift

// StartMatchmaking is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Anywhere, Container
//
// Uses FlexMatch to create a game match for a group of players based on custom
// matchmaking rules. With games that use Amazon GameLift Servers managed hosting,
// this operation also triggers Amazon GameLift Servers to find hosting resources
// and start a new game session for the new match. Each matchmaking request
// includes information on one or more players and specifies the FlexMatch
// matchmaker to use. When a request is for multiple players, FlexMatch attempts to
// build a match that includes all players in the request, placing them in the same
// team and finding additional players as needed to fill the match.
//
// To start matchmaking, provide a unique ticket ID, specify a matchmaking
// configuration, and include the players to be matched. You must also include any
// player attributes that are required by the matchmaking configuration's rule set.
// If successful, a matchmaking ticket is returned with status set to QUEUED .
//
// Track matchmaking events to respond as needed and acquire game session
// connection information for successfully completed matches. Ticket status updates
// are tracked using event notification through Amazon Simple Notification Service,
// which is defined in the matchmaking configuration.
//
// # Learn more
//
// [Add FlexMatch to a game client]
//
// [Set Up FlexMatch event notification]
//
// [How Amazon GameLift Servers FlexMatch works]
//
// [Set Up FlexMatch event notification]: https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-notification.html
// [Add FlexMatch to a game client]: https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-client.html
// [How Amazon GameLift Servers FlexMatch works]: https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/gamelift-match.html
