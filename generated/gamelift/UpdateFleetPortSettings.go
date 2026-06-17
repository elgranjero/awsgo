package gamelift

// UpdateFleetPortSettings is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Container
//
// Updates permissions that allow inbound traffic to connect to game sessions in
// the fleet.
//
// To update settings, specify the fleet ID to be updated and specify the changes
// to be made. List the permissions you want to add in
// InboundPermissionAuthorizations , and permissions you want to remove in
// InboundPermissionRevocations . Permissions to be removed must match existing
// fleet permissions.
//
// If successful, the fleet ID for the updated fleet is returned. For fleets with
// remote locations, port setting updates can take time to propagate across all
// locations. You can check the status of updates in each location by calling
// DescribeFleetPortSettings with a location name.
//
// # Learn more
//
// [Setting up Amazon GameLift Servers fleets]
//
// [Setting up Amazon GameLift Servers fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html
