package gamelift

// DescribeInstances is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2
//
// Retrieves information about the EC2 instances in an Amazon GameLift Servers
// managed fleet, including instance ID, connection data, and status. You can use
// this operation with a multi-location fleet to get location-specific instance
// information. As an alternative, use the operations [https://docs.aws.amazon.com/gamelift/latest/apireference/API_ListCompute]and [https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeCompute] to retrieve information
// for compute resources, including EC2 and Anywhere fleets.
//
// You can call this operation in the following ways:
//
// - To get information on all instances in a fleet's home Region, specify the
// fleet ID.
//
// - To get information on all instances in a fleet's remote location, specify
// the fleet ID and location name.
//
// - To get information on a specific instance in a fleet, specify the fleet ID
// and instance ID.
//
// Use the pagination parameters to retrieve results as a set of sequential pages.
//
// If successful, this operation returns Instance objects for each requested
// instance, listed in no particular order. If you call this operation for an
// Anywhere fleet, you receive an InvalidRequestException.
//
// # Learn more
//
// [Remotely connect to fleet instances]
//
// [Debug fleet issues]
//
// # Related actions
//
// [All APIs by task]
//
// [https://docs.aws.amazon.com/gamelift/latest/apireference/API_ListCompute]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_ListCompute
// [Remotely connect to fleet instances]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-remote-access.html
// [Debug fleet issues]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-creating-debug.html
// [https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeCompute]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeCompute
// [All APIs by task]: https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets
