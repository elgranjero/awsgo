package gamelift

// DescribeRuntimeConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2
//
// Retrieves a fleet's runtime configuration settings. The runtime configuration
// determines which server processes run, and how, on computes in the fleet. For
// managed EC2 fleets, the runtime configuration describes server processes that
// run on each fleet instance. You can update a fleet's runtime configuration at
// any time using [UpdateRuntimeConfiguration].
//
// To get the current runtime configuration for a fleet, provide the fleet ID.
//
// If successful, a RuntimeConfiguration object is returned for the requested
// fleet. If the requested fleet has been deleted, the result set is empty.
//
// # Learn more
//
// [Setting up Amazon GameLift Servers fleets]
//
// [Running multiple processes on a fleet]
//
// [UpdateRuntimeConfiguration]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_UpdateRuntimeConfiguration.html
// [Running multiple processes on a fleet]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-multiprocess.html
// [Setting up Amazon GameLift Servers fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html
