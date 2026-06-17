package gamelift

// CreateFleet is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Anywhere, Container
//
// Creates a fleet of compute resources to host your game servers. Use this
// operation to set up a fleet for the following compute types:
//
// # Managed EC2 fleet
//
// An EC2 fleet is a set of Amazon Elastic Compute Cloud (Amazon EC2) instances.
// Your game server build is deployed to each fleet instance. Amazon GameLift
// Servers manages the fleet's instances and controls the lifecycle of game server
// processes, which host game sessions for players. EC2 fleets can have instances
// in multiple locations. Each instance in the fleet is designated a Compute .
//
// To create an EC2 fleet, provide these required parameters:
//
// - Either BuildId or ScriptId
//
// - ComputeType set to EC2 (the default value)
//
// - EC2InboundPermissions
//
// - EC2InstanceType
//
// - FleetType
//
// - Name
//
// - RuntimeConfiguration with at least one ServerProcesses configuration
//
// If successful, this operation creates a new fleet resource and places it in NEW
// status while Amazon GameLift Servers initiates the [fleet creation workflow]. To debug your fleet, fetch
// logs, view performance metrics or other actions on the fleet, create a
// development fleet with port 22/3389 open. As a best practice, we recommend
// opening ports for remote access only when you need them and closing them when
// you're finished.
//
// When the fleet status is ACTIVE, you can adjust capacity settings and turn
// autoscaling on/off for each location.
//
// A managed fleet's runtime environment depends on the Amazon Machine Image (AMI)
// version it uses. When a new fleet is created, Amazon GameLift Servers assigns
// the latest available AMI version to the fleet, and all compute instances in that
// fleet are deployed with that version. To update the AMI version, you must create
// a new fleet. As a best practice, we recommend replacing your managed fleets
// every 30 days to maintain a secure and up-to-date runtime environment for your
// hosted game servers. For guidance, see [Security best practices for Amazon GameLift Servers].
//
// # Anywhere fleet
//
// An Anywhere fleet represents compute resources that are not owned or managed by
// Amazon GameLift Servers. You might create an Anywhere fleet with your local
// machine for testing, or use one to host game servers with on-premises hardware
// or other game hosting solutions.
//
// To create an Anywhere fleet, provide these required parameters:
//
// - ComputeType set to ANYWHERE
//
// - Locations specifying a custom location
//
// - Name
//
// If successful, this operation creates a new fleet resource and places it in
// ACTIVE status. You can register computes with a fleet in ACTIVE status.
//
// # Learn more
//
// [Setting up fleets]
//
// [Debug fleet creation issues]
//
// [Multi-location fleets]
//
// [fleet creation workflow]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-creating-all.html#fleets-creation-workflow
// [Multi-location fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html
// [Debug fleet creation issues]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-creating-debug.html#fleets-creating-debug-creation
// [Security best practices for Amazon GameLift Servers]: https://docs.aws.amazon.com/gameliftservers/latest/developerguide/security-best-practices.html
// [Setting up fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html
