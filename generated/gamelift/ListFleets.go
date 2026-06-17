package gamelift

// ListFleets is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Anywhere, Container
//
// Retrieves a collection of fleet resources in an Amazon Web Services Region. You
// can filter the result set to find only those fleets that are deployed with a
// specific build or script. For fleets that have multiple locations, this
// operation retrieves fleets based on their home Region only.
//
// You can use operation in the following ways:
//
// - To get a list of all fleets in a Region, don't provide a build or script
// identifier.
//
// - To get a list of all fleets where a specific game build is deployed,
// provide the build ID.
//
// - To get a list of all Amazon GameLift Servers Realtime fleets with a
// specific configuration script, provide the script ID.
//
// Use the pagination parameters to retrieve results as a set of sequential pages.
//
// If successful, this operation returns a list of fleet IDs that match the
// request parameters. A NextToken value is also returned if there are more result
// pages to retrieve.
//
// Fleet IDs are returned in no particular order.
