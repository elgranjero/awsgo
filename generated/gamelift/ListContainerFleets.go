package gamelift

// ListContainerFleets is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: Container
//
// Retrieves a collection of container fleet resources in an Amazon Web Services
// Region. For fleets that have multiple locations, this operation retrieves fleets
// based on their home Region only.
//
// Request options
//
// - Get a list of all fleets. Call this operation without specifying a
// container group definition.
//
// - Get a list of fleets filtered by container group definition. Provide the
// container group definition name or ARN value.
//
// - To get a list of all Amazon GameLift Servers Realtime fleets with a
// specific configuration script, provide the script ID.
//
// Use the pagination parameters to retrieve results as a set of sequential pages.
//
// If successful, this operation returns a collection of container fleets that
// match the request parameters. A NextToken value is also returned if there are
// more result pages to retrieve.
//
// Fleet IDs are returned in no particular order.
