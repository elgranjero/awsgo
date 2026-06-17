package timestreamwrite

// DescribeEndpoints is generated as a reference stub.
// Executable command wiring lives under cmd/timestreamwrite.go.
//
// Returns a list of available endpoints to make Timestream API calls against.
// This API operation is available through both the Write and Query APIs.
//
// Because the Timestream SDKs are designed to transparently work with the
// service’s architecture, including the management and mapping of the service
// endpoints, we don't recommend that you use this API operation unless:
//
// - You are using [VPC endpoints (Amazon Web Services PrivateLink) with Timestream]
//
// - Your application uses a programming language that does not yet have SDK
// support
//
// - You require better control over the client-side implementation
//
// For detailed information on how and when to use and implement
// DescribeEndpoints, see [The Endpoint Discovery Pattern].
//
// [The Endpoint Discovery Pattern]: https://docs.aws.amazon.com/timestream/latest/developerguide/Using.API.html#Using-API.endpoint-discovery
// [VPC endpoints (Amazon Web Services PrivateLink) with Timestream]: https://docs.aws.amazon.com/timestream/latest/developerguide/VPCEndpoints
