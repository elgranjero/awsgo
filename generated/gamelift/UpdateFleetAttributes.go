package gamelift

// UpdateFleetAttributes is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Anywhere, Container
//
// Updates a fleet's mutable attributes, such as game session protection and
// resource creation limits.
//
// To update fleet attributes, specify the fleet ID and the property values that
// you want to change. If successful, Amazon GameLift Servers returns the
// identifiers for the updated fleet.
//
// A managed fleet's runtime environment, which depends on the fleet's Amazon
// Machine Image {AMI} version, can't be updated. You must create a new fleet. As a
// best practice, we recommend replacing your managed fleets every 30 days to
// maintain a secure and up-to-date runtime environment for your hosted game
// servers. For guidance, see [Security best practices for Amazon GameLift Servers].
//
// # Learn more
//
// [Setting up Amazon GameLift Servers fleets]
//
// [Security best practices for Amazon GameLift Servers]: https://docs.aws.amazon.com/gameliftservers/latest/developerguide/security-best-practices.html
// [Setting up Amazon GameLift Servers fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html
