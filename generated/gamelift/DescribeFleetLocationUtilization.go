package gamelift

// DescribeFleetLocationUtilization is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Anywhere, Container
//
// Retrieves current usage data for a fleet location. Utilization data provides a
// snapshot of current game hosting activity at the requested location. Use this
// operation to retrieve utilization information for a fleet's remote location or
// home Region (you can also retrieve home Region utilization by calling
// DescribeFleetUtilization ).
//
// To retrieve utilization data, identify a fleet and location.
//
// If successful, a FleetUtilization object is returned for the requested fleet
// location.
//
// # Learn more
//
// [Setting up Amazon GameLift Servers fleets]
//
// [Amazon GameLift Servers service locations]for managed hosting
//
// [GameLift metrics for fleets]
//
// [Amazon GameLift Servers service locations]: https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-regions.html
// [GameLift metrics for fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/monitoring-cloudwatch.html#gamelift-metrics-fleet
// [Setting up Amazon GameLift Servers fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html
