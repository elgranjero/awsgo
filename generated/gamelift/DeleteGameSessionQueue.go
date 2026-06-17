package gamelift

// DeleteGameSessionQueue is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Anywhere, Container
//
// Deletes a game session queue. Once a queue is successfully deleted, unfulfilled [StartGameSessionPlacement]
// requests that reference the queue will fail. To delete a queue, specify the
// queue name.
//
// [StartGameSessionPlacement]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_StartGameSessionPlacement.html
