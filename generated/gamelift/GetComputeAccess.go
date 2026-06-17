package gamelift

// GetComputeAccess is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Container
//
// Requests authorization to remotely connect to a hosting resource in a Amazon
// GameLift Servers managed fleet. This operation is not used with Amazon GameLift
// Servers Anywhere fleets.
//
// # Request options
//
// Provide the fleet ID and compute name. The compute name varies depending on the
// type of fleet.
//
// - For a compute in a managed EC2 fleet, provide an instance ID. Each instance
// in the fleet is a compute.
//
// - For a compute in a managed container fleet, provide a compute name. In a
// container fleet, each game server container group on a fleet instance is
// assigned a compute name.
//
// # Results
//
// If successful, this operation returns a set of temporary Amazon Web Services
// credentials, including a two-part access key and a session token.
//
// - With a managed EC2 fleet (where compute type is EC2 ), use these credentials
// with Amazon EC2 Systems Manager (SSM) to start a session with the compute. For
// more details, see [Starting a session (CLI)]in the Amazon EC2 Systems Manager User Guide.
//
// [Starting a session (CLI)]: https://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-working-with-sessions-start.html#sessions-start-cli
