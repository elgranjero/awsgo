package ec2

// ModifyVpcPeeringConnectionOptions is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Modifies the VPC peering connection options on one side of a VPC peering
// connection.
//
// If the peered VPCs are in the same Amazon Web Services account, you can enable
// DNS resolution for queries from the local VPC. This ensures that queries from
// the local VPC resolve to private IP addresses in the peer VPC. This option is
// not available if the peered VPCs are in different Amazon Web Services accounts
// or different Regions. For peered VPCs in different Amazon Web Services accounts,
// each Amazon Web Services account owner must initiate a separate request to
// modify the peering connection options. For inter-region peering connections, you
// must use the Region for the requester VPC to modify the requester VPC peering
// options and the Region for the accepter VPC to modify the accepter VPC peering
// options. To verify which VPCs are the accepter and the requester for a VPC
// peering connection, use the DescribeVpcPeeringConnectionscommand.
