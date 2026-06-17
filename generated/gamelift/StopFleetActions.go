package gamelift

// StopFleetActions is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Container
//
// Suspends certain types of activity in a fleet location. Currently, this
// operation is used to stop auto-scaling activity. For multi-location fleets,
// fleet actions are managed separately for each location.
//
// Stopping fleet actions has several potential purposes. It allows you to
// temporarily stop auto-scaling activity but retain your scaling policies for use
// in the future. For multi-location fleets, you can set up fleet-wide
// auto-scaling, and then opt out of it for certain locations.
//
// This operation can be used in the following ways:
//
// - To stop actions on instances in the fleet's home Region, provide a fleet ID
// and the type of actions to suspend.
//
// - To stop actions on instances in one of the fleet's remote locations,
// provide a fleet ID, a location name, and the type of actions to suspend.
//
// If successful, Amazon GameLift Servers no longer initiates scaling events
// except in response to manual changes using [UpdateFleetCapacity]. To restart fleet actions again,
// call [StartFleetActions].
//
// # Learn more
//
// [Setting up Amazon GameLift Servers Fleets]
//
// [UpdateFleetCapacity]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_UpdateFleetCapacity.html
// [Setting up Amazon GameLift Servers Fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html
// [StartFleetActions]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_StartFleetActions.html
