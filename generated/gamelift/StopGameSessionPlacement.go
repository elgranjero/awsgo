package gamelift

// StopGameSessionPlacement is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Anywhere, Container
//
// Cancels a game session placement that's in PENDING status. To stop a placement,
// provide the placement ID value.
//
// # Results
//
// If successful, this operation removes the placement request from the queue and
// moves the GameSessionPlacement to CANCELLED status.
//
// This operation results in an InvalidRequestExecption (400) error if a game
// session has already been created for this placement. You can clean up an
// unneeded game session by calling [TerminateGameSession].
//
// [TerminateGameSession]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_TerminateGameSession
