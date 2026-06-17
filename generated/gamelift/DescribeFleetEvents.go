package gamelift

// DescribeFleetEvents is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Anywhere, Container
//
// Retrieves entries from a fleet's event log. Fleet events are initiated by
// changes in status, such as during fleet creation and termination, changes in
// capacity, etc. If a fleet has multiple locations, events are also initiated by
// changes to status and capacity in remote locations.
//
// You can specify a time range to limit the result set. Use the pagination
// parameters to retrieve results as a set of sequential pages.
//
// If successful, a collection of event log entries matching the request are
// returned.
//
// # Learn more
//
// [Setting up Amazon GameLift Servers fleets]
//
// [Setting up Amazon GameLift Servers fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html
