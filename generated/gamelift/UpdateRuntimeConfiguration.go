package gamelift

// UpdateRuntimeConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2
//
// Updates the runtime configuration for the specified fleet. The runtime
// configuration tells Amazon GameLift Servers how to launch server processes on
// computes in managed EC2 and Anywhere fleets. You can update a fleet's runtime
// configuration at any time after the fleet is created; it does not need to be in
// ACTIVE status.
//
// To update runtime configuration, specify the fleet ID and provide a
// RuntimeConfiguration with an updated set of server process configurations.
//
// If successful, the fleet's runtime configuration settings are updated. Fleet
// computes that run game server processes regularly check for and receive updated
// runtime configurations. The computes immediately take action to comply with the
// new configuration by launching new server processes or by not replacing existing
// processes when they shut down. Updating a fleet's runtime configuration never
// affects existing server processes.
//
// # Learn more
//
// [Setting up Amazon GameLift Servers fleets]
//
// [Setting up Amazon GameLift Servers fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html
