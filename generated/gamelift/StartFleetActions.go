package gamelift

// StartFleetActions is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Container
//
// Resumes certain types of activity on fleet instances that were suspended with [StopFleetActions].
// For multi-location fleets, fleet actions are managed separately for each
// location. Currently, this operation is used to restart a fleet's auto-scaling
// activity.
//
// This operation can be used in the following ways:
//
// - To restart actions on instances in the fleet's home Region, provide a fleet
// ID and the type of actions to resume.
//
// - To restart actions on instances in one of the fleet's remote locations,
// provide a fleet ID, a location name, and the type of actions to resume.
//
// If successful, Amazon GameLift Servers once again initiates scaling events as
// triggered by the fleet's scaling policies. If actions on the fleet location were
// never stopped, this operation will have no effect.
//
// # Learn more
//
// [Setting up Amazon GameLift Servers fleets]
//
// [StopFleetActions]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_StopFleetActions.html
// [Setting up Amazon GameLift Servers fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html
