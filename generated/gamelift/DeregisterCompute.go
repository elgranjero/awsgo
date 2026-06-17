package gamelift

// DeregisterCompute is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: Anywhere
//
// Removes a compute resource from an Anywhere fleet. Deregistered computes can no
// longer host game sessions through Amazon GameLift Servers. Use this operation
// with an Anywhere fleet that doesn't use the Amazon GameLift Servers Agent For
// Anywhere fleets with the Agent, the Agent handles all compute registry tasks for
// you.
//
// To deregister a compute, call this operation from the compute that's being
// deregistered and specify the compute name and the fleet ID.
