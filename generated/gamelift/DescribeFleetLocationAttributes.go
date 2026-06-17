package gamelift

// DescribeFleetLocationAttributes is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Container
//
// Retrieves information on a fleet's remote locations, including life-cycle
// status and any suspended fleet activity.
//
// This operation can be used in the following ways:
//
// - To get data for specific locations, provide a fleet identifier and a list
// of locations. Location data is returned in the order that it is requested.
//
// - To get data for all locations, provide a fleet identifier only. Location
// data is returned in no particular order.
//
// When requesting attributes for multiple locations, use the pagination
// parameters to retrieve results as a set of sequential pages.
//
// If successful, a LocationAttributes object is returned for each requested
// location. If the fleet does not have a requested location, no information is
// returned. This operation does not return the home Region. To get information on
// a fleet's home Region, call DescribeFleetAttributes .
//
// # Learn more
//
// [Setting up Amazon GameLift Servers fleets]
//
// [Amazon GameLift Servers service locations]for managed hosting
//
// [Amazon GameLift Servers service locations]: https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-regions.html
// [Setting up Amazon GameLift Servers fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html
