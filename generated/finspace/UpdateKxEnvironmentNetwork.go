package finspace

// UpdateKxEnvironmentNetwork is generated as a reference stub.
// Executable command wiring lives under cmd/finspace.go.
//
// Updates environment network to connect to your internal network by using a
// transit gateway. This API supports request to create a transit gateway
// attachment from FinSpace VPC to your transit gateway ID and create a custom
// Route-53 outbound resolvers.
//
// Once you send a request to update a network, you cannot change it again.
// Network update might require termination of any clusters that are running in the
// existing network.
