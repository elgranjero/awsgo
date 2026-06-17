package gamelift

// UpdateContainerFleet is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: Container
//
// Updates the properties of a managed container fleet. Depending on the
// properties being updated, this operation might initiate a fleet deployment. You
// can track deployments for a fleet using [https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeFleetDeployment.html].
//
// A managed fleet's runtime environment, which depends on the fleet's Amazon
// Machine Image {AMI} version, can't be updated. You must create a new fleet. As a
// best practice, we recommend replacing your managed fleets every 30 days to
// maintain a secure and up-to-date runtime environment for your hosted game
// servers. For guidance, see [Security best practices for Amazon GameLift Servers].
//
// # Request options
//
// As with CreateContainerFleet, many fleet properties use common defaults or are
// calculated based on the fleet's container group definitions.
//
// - Update fleet properties that result in a fleet deployment. Include only
// those properties that you want to change. Specify deployment configuration
// settings.
//
// - Update fleet properties that don't result in a fleet deployment. Include
// only those properties that you want to change.
//
// Changes to the following properties initiate a fleet deployment:
//
// - GameServerContainerGroupDefinition
//
// - PerInstanceContainerGroupDefinition
//
// - GameServerContainerGroupsPerInstance
//
// - InstanceInboundPermissions
//
// - InstanceConnectionPortRange
//
// - LogConfiguration
//
// # Results
//
// If successful, this operation updates the container fleet resource, and might
// initiate a new deployment of fleet resources using the deployment configuration
// provided. A deployment replaces existing fleet instances with new instances that
// are deployed with the updated fleet properties. The fleet is placed in UPDATING
// status until the deployment is complete, then return to ACTIVE .
//
// You can have only one update deployment active at a time for a fleet. If a
// second update request initiates a deployment while another deployment is in
// progress, the first deployment is cancelled.
//
// [https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeFleetDeployment.html]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeFleetDeployment.html
// [Security best practices for Amazon GameLift Servers]: https://docs.aws.amazon.com/gameliftservers/latest/developerguide/security-best-practices.html
