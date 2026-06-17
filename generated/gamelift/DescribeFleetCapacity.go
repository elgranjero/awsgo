package gamelift

// DescribeFleetCapacity is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Container
//
// Retrieves the resource capacity settings for one or more fleets. For a
// container fleet, this operation also returns counts for game server container
// groups.
//
// With multi-location fleets, this operation retrieves data for the fleet's home
// Region only. To retrieve capacity for remote locations, see [https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeFleetLocationCapacity.html].
//
// This operation can be used in the following ways:
//
// - To get capacity data for one or more specific fleets, provide a list of
// fleet IDs or fleet ARNs.
//
// - To get capacity data for all fleets, do not provide a fleet identifier.
//
// When requesting multiple fleets, use the pagination parameters to retrieve
// results as a set of sequential pages.
//
// If successful, a FleetCapacity object is returned for each requested fleet ID.
// Each FleetCapacity object includes a Location property, which is set to the
// fleet's home Region. Capacity values are returned only for fleets that currently
// exist.
//
// Some API operations may limit the number of fleet IDs that are allowed in one
// request. If a request exceeds this limit, the request fails and the error
// message includes the maximum allowed.
//
// # Learn more
//
// [Setting up Amazon GameLift Servers fleets]
//
// [GameLift metrics for fleets]
//
// [https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeFleetLocationCapacity.html]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeFleetLocationCapacity.html
// [GameLift metrics for fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/monitoring-cloudwatch.html#gamelift-metrics-fleet
// [Setting up Amazon GameLift Servers fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html
