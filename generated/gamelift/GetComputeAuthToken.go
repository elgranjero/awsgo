package gamelift

// GetComputeAuthToken is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Anywhere, Container
//
// Requests an authentication token from Amazon GameLift Servers for a compute
// resource in an Amazon GameLift Servers fleet. Game servers that are running on
// the compute use this token to communicate with the Amazon GameLift Servers
// service, such as when calling the Amazon GameLift Servers server SDK action
// InitSDK() . Authentication tokens are valid for a limited time span, so you need
// to request a fresh token before the current token expires.
//
// Request options
//
// - For managed EC2 fleets (compute type EC2 ), auth token retrieval and refresh
// is handled automatically. All game servers that are running on all fleet
// instances have access to a valid auth token.
//
// - For Anywhere fleets (compute type ANYWHERE ), if you're using the Amazon
// GameLift Servers Agent, auth token retrieval and refresh is handled
// automatically for any compute where the Agent is running. If you're not using
// the Agent, create a mechanism to retrieve and refresh auth tokens for computes
// that are running game server processes.
//
// # Learn more
//
// [Create an Anywhere fleet]
//
// [Test your integration]
//
// [Server SDK reference guides]
// - (for version 5.x)
//
// [Test your integration]: https://docs.aws.amazon.com/gamelift/latest/developerguide/integration-testing.html
// [Server SDK reference guides]: https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-serversdk.html
// [Create an Anywhere fleet]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-creating-anywhere.html
