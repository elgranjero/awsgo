package gamelift

// ListFleetDeployments is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: Container
//
// Retrieves a collection of container fleet deployments in an Amazon Web Services
// Region. Use the pagination parameters to retrieve results as a set of sequential
// pages.
//
// Request options
//
// - Get a list of all deployments. Call this operation without specifying a
// fleet ID.
//
// - Get a list of all deployments for a fleet. Specify the container fleet ID
// or ARN value.
//
// # Results
//
// If successful, this operation returns a list of deployments that match the
// request parameters. A NextToken value is also returned if there are more result
// pages to retrieve.
//
// Deployments are returned starting with the latest.
