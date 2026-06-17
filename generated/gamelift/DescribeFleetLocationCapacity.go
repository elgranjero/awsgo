package gamelift

// DescribeFleetLocationCapacity is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Container
//
// Retrieves the resource capacity settings for a fleet location. The data
// returned includes the current capacity (number of EC2 instances) and some
// scaling settings for the requested fleet location. For a managed container
// fleet, this operation also returns counts for game server container groups.
//
// Use this operation to retrieve capacity information for a fleet's remote
// location or home Region (you can also retrieve home Region capacity by calling
// DescribeFleetCapacity ).
//
// To retrieve capacity data, identify a fleet and location.
//
// If successful, a FleetCapacity object is returned for the requested fleet
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
