package ec2

// DeleteVpcPeeringConnection is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Deletes a VPC peering connection. Either the owner of the requester VPC or the
// owner of the accepter VPC can delete the VPC peering connection if it's in the
// active state. The owner of the requester VPC can delete a VPC peering connection
// in the pending-acceptance state. You cannot delete a VPC peering connection
// that's in the failed or rejected state.
