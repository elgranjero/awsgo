package gamelift

// DescribeVpcPeeringConnections is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2
//
// Retrieves information on VPC peering connections. Use this operation to get
// peering information for all fleets or for one specific fleet ID.
//
// To retrieve connection information, call this operation from the Amazon Web
// Services account that is used to manage the Amazon GameLift Servers fleets.
// Specify a fleet ID or leave the parameter empty to retrieve all connection
// records. If successful, the retrieved information includes both active and
// pending connections. Active connections identify the IpV4 CIDR block that the
// VPC uses to connect.
//
// # Related actions
//
// [All APIs by task]
//
// [All APIs by task]: https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets
