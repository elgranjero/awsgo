package gamelift

// RegisterCompute is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: Anywhere, Container
//
// Registers a compute resource in an Amazon GameLift Servers Anywhere fleet.
//
// For an Anywhere fleet that's running the Amazon GameLift Servers Agent, the
// Agent handles all compute registry tasks for you. For an Anywhere fleet that
// doesn't use the Agent, call this operation to register fleet computes.
//
// To register a compute, give the compute a name (must be unique within the
// fleet) and specify the compute resource's DNS name or IP address. Provide a
// fleet ID and a fleet location to associate with the compute being registered.
// You can optionally include the path to a TLS certificate on the compute
// resource.
//
// If successful, this operation returns compute details, including an Amazon
// GameLift Servers SDK endpoint or Agent endpoint. Game server processes running
// on the compute can use this endpoint to communicate with the Amazon GameLift
// Servers service. Each server process includes the SDK endpoint in its call to
// the Amazon GameLift Servers server SDK action InitSDK() .
//
// To view compute details, call [DescribeCompute] with the compute name.
//
// # Learn more
//
// [Create an Anywhere fleet]
//
// [Test your integration]
//
// [Server SDK reference guides]
// - (for version 5.x)
//
// [Test your integration]: https://docs.aws.amazon.com/gamelift/latest/developerguide/integration-testing.html
// [Server SDK reference guides]: https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-serversdk.html
// [DescribeCompute]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeCompute.html
// [Create an Anywhere fleet]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-creating-anywhere.html
