package gamelift

// UpdateFleetCapacity is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Container
//
// Updates capacity settings for a managed EC2 fleet or managed container fleet.
// For these fleets, you adjust capacity by changing the number of instances in the
// fleet. Fleet capacity determines the number of game sessions and players that
// the fleet can host based on its configuration. For fleets with multiple
// locations, use this operation to manage capacity settings in each location
// individually.
//
// - Minimum/maximum size: Set hard limits on the number of Amazon EC2 instances
// allowed. If Amazon GameLift Servers receives a request--either through manual
// update or automatic scaling--it won't change the capacity to a value outside of
// this range.
//
// - Desired capacity: As an alternative to automatic scaling, manually set the
// number of Amazon EC2 instances to be maintained. Before changing a fleet's
// desired capacity, check the maximum capacity of the fleet's Amazon EC2 instance
// type by calling [DescribeEC2InstanceLimits].
//
// To update capacity for a fleet's home Region, or if the fleet has no remote
// locations, omit the Location parameter. The fleet must be in ACTIVE status.
//
// To update capacity for a fleet's remote location, set the Location parameter to
// the location to update. The location must be in ACTIVE status.
//
// If successful, Amazon GameLift Servers updates the capacity settings and
// returns the identifiers for the updated fleet and/or location. If a requested
// change to desired capacity exceeds the instance type's limit, the LimitExceeded
// exception occurs.
//
// Updates often prompt an immediate change in fleet capacity, such as when
// current capacity is different than the new desired capacity or outside the new
// limits. In this scenario, Amazon GameLift Servers automatically initiates steps
// to add or remove instances in the fleet location. You can track a fleet's
// current capacity by calling [DescribeFleetCapacity]or [DescribeFleetLocationCapacity].
//
// Use ManagedCapacityConfiguration with the "SCALE_TO_AND_FROM_ZERO"
// ZeroCapacityStrategy to enable Amazon GameLift Servers to fully manage the
// MinSize value, switching between 0 and 1 based on game session activity. This is
// ideal for eliminating compute costs during periods of no game activity. It is
// particularly beneficial during development when you're away from your desk,
// iterating on builds for extended periods, in production environments serving
// low-traffic locations, or for games with long, predictable downtime windows. By
// automatically managing capacity between 0 and 1 instances, you avoid paying for
// idle instances while maintaining the ability to serve game sessions when demand
// arrives. Note that while scale-out is triggered immediately upon receiving a
// game session request, actual game session availability depends on your server
// process startup time, so this approach works best with multi-location Fleets
// where cold-start latency is tolerable. With a "MANUAL" ZeroCapacityStrategy
// Amazon GameLift Servers will not modify Fleet MinSize values automatically and
// will not scale out from zero instances in response to game sessions. This is
// configurable per-location.
//
// # Learn more
//
// [Scaling fleet capacity]
//
// [Scaling fleet capacity]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-manage-capacity.html
// [DescribeEC2InstanceLimits]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeEC2InstanceLimits.html
// [DescribeFleetLocationCapacity]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeFleetLocationCapacity.html
// [DescribeFleetCapacity]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeFleetCapacity.html
