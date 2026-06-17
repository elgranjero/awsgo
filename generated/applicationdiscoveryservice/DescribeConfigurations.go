package applicationdiscoveryservice

// DescribeConfigurations is generated as a reference stub.
// Executable command wiring lives under cmd/applicationdiscoveryservice.go.
//
// Retrieves attributes for a list of configuration item IDs.
//
// All of the supplied IDs must be for the same asset type from one of the
// following:
//
// - server
//
// - application
//
// - process
//
// - connection
//
// Output fields are specific to the asset type specified. For example, the output
// for a server configuration item includes a list of attributes about the server,
// such as host name, operating system, number of network cards, etc.
//
// For a complete list of outputs for each asset type, see [Using the DescribeConfigurations Action] in the Amazon Web
// Services Application Discovery Service User Guide.
//
// [Using the DescribeConfigurations Action]: https://docs.aws.amazon.com/application-discovery/latest/userguide/discovery-api-queries.html#DescribeConfigurations
