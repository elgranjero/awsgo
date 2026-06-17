package gamelift

// DeleteScalingPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2
//
// Deletes a fleet scaling policy. Once deleted, the policy is no longer in force
// and Amazon GameLift Servers removes all record of it. To delete a scaling
// policy, specify both the scaling policy name and the fleet ID it is associated
// with.
//
// To temporarily suspend scaling policies, use [StopFleetActions]. This operation suspends all
// policies for the fleet.
//
// [StopFleetActions]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_StopFleetActions.html
