package gamelift

// TerminateGameSession is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Anywhere, Container
//
// Ends a game session that's currently in progress. Use this action to terminate
// any game session that isn't in ERROR status. Terminating a game session is the
// most efficient way to free up a server process when it's hosting a game session
// that's in a bad state or not ending properly. You can use this action to
// terminate a game session that's being hosted on any type of Amazon GameLift
// Servers fleet compute, including computes for managed EC2, managed container,
// and Anywhere fleets. The game server must be integrated with Amazon GameLift
// Servers server SDK 5.x or greater.
//
// # Request options
//
// Request termination for a single game session. Provide the game session ID and
// the termination mode. There are two potential methods for terminating a game
// session:
//
// - Initiate a graceful termination using the normal game session shutdown
// sequence. With this mode, the Amazon GameLift Servers service prompts the server
// process that's hosting the game session by calling the server SDK callback
// method OnProcessTerminate() . The callback implementation is part of the
// custom game server code. It might involve a variety of actions to gracefully end
// a game session, such as notifying players, before stopping the server process.
//
// - Force an immediate game session termination. With this mode, the Amazon
// GameLift Servers service takes action to stop the server process, which ends the
// game session without the normal game session shutdown sequence.
//
// # Results
//
// If successful, game session termination is initiated. During this activity, the
// game session status is changed to TERMINATING . When completed, the server
// process that was hosting the game session has been stopped and replaced with a
// new server process that's ready to host a new game session. The old game
// session's status is changed to TERMINATED with a status reason that indicates
// the termination method used.
//
// # Learn more
//
// [Add Amazon GameLift Servers to your game server]
//
// Amazon GameLift Servers server SDK 5 reference guide for OnProcessTerminate() ([C++]
// ) ([C#] ) ([Unreal] ) ([Go] )
//
// [C#]: https://docs.aws.amazon.com/gamelift/latest/developerguide/integration-server-sdk5-csharp-initsdk.html
// [C++]: https://docs.aws.amazon.com/gamelift/latest/developerguide/integration-server-sdk5-cpp-initsdk.html
// [Add Amazon GameLift Servers to your game server]: https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api.html
// [Unreal]: https://docs.aws.amazon.com/gamelift/latest/developerguide/integration-server-sdk5-unreal-initsdk.html
// [Go]: https://docs.aws.amazon.com/gamelift/latest/developerguide/integration-server-sdk-go-initsdk.html
