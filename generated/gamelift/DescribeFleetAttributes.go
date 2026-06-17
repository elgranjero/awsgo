package gamelift

// DescribeFleetAttributes is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Anywhere, Container
//
// Retrieves core fleet-wide properties for fleets in an Amazon Web Services
// Region. Properties include the computing hardware and deployment configuration
// for instances in the fleet.
//
// You can use this operation in the following ways:
//
// - To get attributes for specific fleets, provide a list of fleet IDs or fleet
// ARNs.
//
// - To get attributes for all fleets, do not provide a fleet identifier.
//
// When requesting attributes for multiple fleets, use the pagination parameters
// to retrieve results as a set of sequential pages.
//
// If successful, a FleetAttributes object is returned for each fleet requested,
// unless the fleet identifier is not found.
//
// Some API operations limit the number of fleet IDs that allowed in one request.
// If a request exceeds this limit, the request fails and the error message
// contains the maximum allowed number.
//
// # Learn more
//
// [Setting up Amazon GameLift Servers fleets]
//
// [Setting up Amazon GameLift Servers fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html
