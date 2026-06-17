package gamelift

// CreateContainerFleet is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: Container
//
// Creates a managed fleet of Amazon Elastic Compute Cloud (Amazon EC2) instances
// to host your containerized game servers. Use this operation to define how to
// deploy a container architecture onto each fleet instance and configure fleet
// settings. You can create a container fleet in any Amazon Web Services Regions
// that Amazon GameLift Servers supports for multi-location fleets. A container
// fleet can be deployed to a single location or multiple locations. Container
// fleets are deployed with Amazon Linux 2023 as the instance operating system.
//
// Define the fleet's container architecture using container group definitions.
// Each fleet can have one of the following container group types:
//
// - The game server container group runs your game server build and dependent
// software. Amazon GameLift Servers deploys one or more replicas of this container
// group to each fleet instance. The number of replicas depends on the computing
// capabilities of the fleet instance in use.
//
// - An optional per-instance container group might be used to run other
// software that only needs to run once per instance, such as background services,
// logging, or test processes. One per-instance container group is deployed to each
// fleet instance.
//
// Each container group can include the definition for one or more containers. A
// container definition specifies a container image that is stored in an Amazon
// Elastic Container Registry (Amazon ECR) public or private repository.
//
// # Request options
//
// Use this operation to make the following types of requests. Most fleet settings
// have default values, so you can create a working fleet with a minimal
// configuration and default values, which you can customize later.
//
// - Create a fleet with no container groups. You can configure a container
// fleet and then add container group definitions later. In this scenario, no fleet
// instances are deployed, and the fleet can't host game sessions until you add a
// game server container group definition. Provide the following required parameter
// values:
//
// - FleetRoleArn
//
// - Create a fleet with a game server container group. Provide the following
// required parameter values:
//
// - FleetRoleArn
//
// - GameServerContainerGroupDefinitionName
//
// - Create a fleet with a game server container group and a per-instance
// container group. Provide the following required parameter values:
//
// - FleetRoleArn
//
// - GameServerContainerGroupDefinitionName
//
// - PerInstanceContainerGroupDefinitionName
//
// # Results
//
// If successful, this operation creates a new container fleet resource, places it
// in PENDING status, and initiates the [fleet creation workflow]. For fleets with container groups, this
// workflow starts a fleet deployment and transitions the status to ACTIVE . Fleets
// without a container group are placed in CREATED status.
//
// You can update most of the properties of a fleet, including container group
// definitions, and deploy the update across all fleet instances. Use [UpdateContainerFleet]to deploy a
// new game server version update across the container fleet.
//
// A managed fleet's runtime environment depends on the Amazon Machine Image (AMI)
// version it uses. When a new fleet is created, Amazon GameLift Servers assigns
// the latest available AMI version to the fleet, and all compute instances in that
// fleet are deployed with that version. To update the AMI version, you must create
// a new fleet. As a best practice, we recommend replacing your managed fleets
// every 30 days to maintain a secure and up-to-date runtime environment for your
// hosted game servers. For guidance, see [Security best practices for Amazon GameLift Servers].
//
// [fleet creation workflow]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-creating-all.html#fleets-creation-workflow
// [Security best practices for Amazon GameLift Servers]: https://docs.aws.amazon.com/gameliftservers/latest/developerguide/security-best-practices.html
// [UpdateContainerFleet]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_UpdateContainerFleet.html
