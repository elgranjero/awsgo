package gamelift

// DeleteFleet is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Anywhere, Container
//
// Deletes all resources and information related to a fleet and shuts down any
// currently running fleet instances, including those in remote locations.
//
// If the fleet being deleted has a VPC peering connection, you first need to get
// a valid authorization (good for 24 hours) by calling [CreateVpcPeeringAuthorization]. You don't need to
// explicitly delete the VPC peering connection.
//
// To delete a fleet, specify the fleet ID to be terminated. During the deletion
// process, the fleet status is changed to DELETING . When completed, the status
// switches to TERMINATED and the fleet event FLEET_DELETED is emitted.
//
// # Learn more
//
// [Setting up Amazon GameLift Servers Fleets]
//
// [CreateVpcPeeringAuthorization]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_CreateVpcPeeringAuthorization.html
// [Setting up Amazon GameLift Servers Fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html
