package gamelift

// ListCompute is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Anywhere, Container
//
// Retrieves information on the compute resources in an Amazon GameLift Servers
// fleet. Use the pagination parameters to retrieve results in a set of sequential
// pages.
//
// Request options
//
// - Retrieve a list of all computes in a fleet. Specify a fleet ID.
//
// - Retrieve a list of all computes in a specific fleet location. Specify a
// fleet ID and location.
//
// # Results
//
// If successful, this operation returns information on a set of computes.
// Depending on the type of fleet, the result includes the following information:
//
// - For a managed EC2 fleet (compute type EC2 ), this operation returns
// information about the EC2 instance. Compute names are EC2 instance IDs.
//
// - For an Anywhere fleet (compute type ANYWHERE ), this operation returns
// compute names and details from when the compute was registered with
// RegisterCompute . This includes GameLiftServiceSdkEndpoint or
// GameLiftAgentEndpoint .
