package managedblockchain

// DeleteMember is generated as a reference stub.
// Executable command wiring lives under cmd/managedblockchain.go.
//
// Deletes a member. Deleting a member removes the member and all associated
// resources from the network. DeleteMember can only be called for a specified
// MemberId if the principal performing the action is associated with the Amazon
// Web Services account that owns the member. In all other cases, the DeleteMember
// action is carried out as the result of an approved proposal to remove a member.
// If MemberId is the last member in a network specified by the last Amazon Web
// Services account, the network is deleted also.
//
// Applies only to Hyperledger Fabric.
