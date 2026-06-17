package gamelift

// DescribeFleetUtilization is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Container
//
// Retrieves utilization statistics for one or more fleets. Utilization data
// provides a snapshot of how the fleet's hosting resources are currently being
// used. For fleets with remote locations, this operation retrieves data for the
// fleet's home Region only. See [DescribeFleetLocationUtilization]to get utilization statistics for a fleet's
// remote locations.
//
// This operation can be used in the following ways:
//
// - To get utilization data for one or more specific fleets, provide a list of
// fleet IDs or fleet ARNs.
//
// - To get utilization data for all fleets, do not provide a fleet identifier.
//
// When requesting multiple fleets, use the pagination parameters to retrieve
// results as a set of sequential pages.
//
// If successful, a [FleetUtilization] object is returned for each requested fleet ID, unless the
// fleet identifier is not found. Each fleet utilization object includes a Location
// property, which is set to the fleet's home Region.
//
// Some API operations may limit the number of fleet IDs allowed in one request.
// If a request exceeds this limit, the request fails and the error message
// includes the maximum allowed.
//
// # Learn more
//
// [Setting up Amazon GameLift Servers Fleets]
//
// [GameLift Metrics for Fleets]
//
// [FleetUtilization]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_FleetUtilization.html
// [Setting up Amazon GameLift Servers Fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html
// [DescribeFleetLocationUtilization]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeFleetLocationUtilization.html
// [GameLift Metrics for Fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/monitoring-cloudwatch.html#gamelift-metrics-fleet
