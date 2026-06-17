package gamelift

// DescribeContainerFleet is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: Container
//
// Retrieves the properties for a container fleet. When requesting attributes for
// multiple fleets, use the pagination parameters to retrieve results as a set of
// sequential pages.
//
// Request options
//
// - Get container fleet properties for a single fleet. Provide either the fleet
// ID or ARN value.
//
// # Results
//
// If successful, a ContainerFleet object is returned. This object includes the
// fleet properties, including information about the most recent deployment.
//
// Some API operations limit the number of fleet IDs that allowed in one request.
// If a request exceeds this limit, the request fails and the error message
// contains the maximum allowed number.
