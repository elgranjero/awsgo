package gamelift

// CreateVpcPeeringAuthorization is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2
//
// Requests authorization to create or delete a peer connection between the VPC
// for your Amazon GameLift Servers fleet and a virtual private cloud (VPC) in your
// Amazon Web Services account. VPC peering enables the game servers on your fleet
// to communicate directly with other Amazon Web Services resources. After you've
// received authorization, use [CreateVpcPeeringConnection]to establish the peering connection. For more
// information, see [VPC Peering with Amazon GameLift Servers Fleets].
//
// You can peer with VPCs that are owned by any Amazon Web Services account you
// have access to, including the account that you use to manage your Amazon
// GameLift Servers fleets. You cannot peer with VPCs that are in different
// Regions.
//
// To request authorization to create a connection, call this operation from the
// Amazon Web Services account with the VPC that you want to peer to your Amazon
// GameLift Servers fleet. For example, to enable your game servers to retrieve
// data from a DynamoDB table, use the account that manages that DynamoDB resource.
// Identify the following values: (1) The ID of the VPC that you want to peer with,
// and (2) the ID of the Amazon Web Services account that you use to manage Amazon
// GameLift Servers. If successful, VPC peering is authorized for the specified
// VPC.
//
// To request authorization to delete a connection, call this operation from the
// Amazon Web Services account with the VPC that is peered with your Amazon
// GameLift Servers fleet. Identify the following values: (1) VPC ID that you want
// to delete the peering connection for, and (2) ID of the Amazon Web Services
// account that you use to manage Amazon GameLift Servers.
//
// The authorization remains valid for 24 hours unless it is canceled. You must
// create or delete the peering connection while the authorization is valid.
//
// # Related actions
//
// [All APIs by task]
//
// [CreateVpcPeeringConnection]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_CreateVpcPeeringConnection.html
// [VPC Peering with Amazon GameLift Servers Fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/vpc-peering.html
// [All APIs by task]: https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets
