package gamelift

// CreateVpcPeeringConnection is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2
//
// Establishes a VPC peering connection between a virtual private cloud (VPC) in
// an Amazon Web Services account with the VPC for your Amazon GameLift Servers
// fleet. VPC peering enables the game servers on your fleet to communicate
// directly with other Amazon Web Services resources. You can peer with VPCs in any
// Amazon Web Services account that you have access to, including the account that
// you use to manage your Amazon GameLift Servers fleets. You cannot peer with VPCs
// that are in different Regions. For more information, see [VPC Peering with Amazon GameLift Servers Fleets].
//
// Before calling this operation to establish the peering connection, you first
// need to use [CreateVpcPeeringAuthorization]and identify the VPC you want to peer with. Once the authorization
// for the specified VPC is issued, you have 24 hours to establish the connection.
// These two operations handle all tasks necessary to peer the two VPCs, including
// acceptance, updating routing tables, etc.
//
// To establish the connection, call this operation from the Amazon Web Services
// account that is used to manage the Amazon GameLift Servers fleets. Identify the
// following values: (1) The ID of the fleet you want to be enable a VPC peering
// connection for; (2) The Amazon Web Services account with the VPC that you want
// to peer with; and (3) The ID of the VPC you want to peer with. This operation is
// asynchronous. If successful, a connection request is created. You can use
// continuous polling to track the request's status using [DescribeVpcPeeringConnections], or by monitoring fleet
// events for success or failure using [DescribeFleetEvents].
//
// # Related actions
//
// [All APIs by task]
//
// [DescribeFleetEvents]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeFleetEvents.html
// [CreateVpcPeeringAuthorization]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_CreateVpcPeeringAuthorization.html
// [VPC Peering with Amazon GameLift Servers Fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/vpc-peering.html
// [DescribeVpcPeeringConnections]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeVpcPeeringConnections.html
// [All APIs by task]: https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets
