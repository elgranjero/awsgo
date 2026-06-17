package gamelift

// DescribeCompute is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Anywhere, Container
//
// Retrieves properties for a specific compute resource in an Amazon GameLift
// Servers fleet. You can list all computes in a fleet by calling [ListCompute].
//
// # Request options
//
// Provide the fleet ID and compute name. The compute name varies depending on the
// type of fleet.
//
// - For a compute in a managed EC2 fleet, provide an instance ID. Each instance
// in the fleet is a compute.
//
// - For a compute in a managed container fleet, provide a compute name. In a
// container fleet, each game server container group on a fleet instance is
// assigned a compute name.
//
// - For a compute in an Anywhere fleet, provide a registered compute name.
// Anywhere fleet computes are created when you register a hosting resource with
// the fleet.
//
// # Results
//
// If successful, this operation returns details for the requested compute
// resource. Depending on the fleet's compute type, the result includes the
// following information:
//
// - For a managed EC2 fleet, this operation returns information about the EC2
// instance.
//
// - For an Anywhere fleet, this operation returns information about the
// registered compute.
//
// [ListCompute]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_ListCompute.html
