package gamelift

// CreateGameSession is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Anywhere, Container
//
// Creates a multiplayer game session for players in a specific fleet location.
// This operation prompts an available server process to start a game session and
// retrieves connection information for the new game session. As an alternative,
// consider using the Amazon GameLift Servers game session placement feature with [StartGameSessionPlacement]
// , which uses the FleetIQ algorithm and queues to optimize the placement process.
//
// When creating a game session, you specify exactly where you want to place it
// and provide a set of game session configuration settings. The target fleet must
// be in ACTIVE status.
//
// You can use this operation in the following ways:
//
// - To create a game session on an instance in a fleet's home Region, provide a
// fleet or alias ID along with your game session configuration.
//
// - To create a game session on an instance in a fleet's remote location,
// provide a fleet or alias ID and a location name, along with your game session
// configuration.
//
// - To create a game session on an instance in an Anywhere fleet, specify the
// fleet's custom location.
//
// If successful, Amazon GameLift Servers initiates a workflow to start a new game
// session and returns a GameSession object containing the game session
// configuration and status. When the game session status is ACTIVE , it is updated
// with connection information and you can create player sessions for the game
// session. By default, newly created game sessions are open to new players. You
// can restrict new player access by using [UpdateGameSession]to change the game session's player
// session creation policy.
//
// Amazon GameLift Servers retains logs for active for 14 days. To access the
// logs, call [GetGameSessionLogUrl]to download the log files.
//
// Available in Amazon GameLift Servers Local.
//
// # Learn more
//
// [Start a game session]
//
// [All APIs by task]
//
// [Start a game session]: https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api.html#gamelift-sdk-server-startsession
// [GetGameSessionLogUrl]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_GetGameSessionLogUrl.html
// [StartGameSessionPlacement]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_StartGameSessionPlacement.html
// [UpdateGameSession]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_UpdateGameSession.html
// [All APIs by task]: https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets
