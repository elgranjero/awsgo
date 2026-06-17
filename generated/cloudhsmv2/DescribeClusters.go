package cloudhsmv2

// DescribeClusters is generated as a reference stub.
// Executable command wiring lives under cmd/cloudhsmv2.go.
//
// Gets information about CloudHSM clusters.
//
// This is a paginated operation, which means that each response might contain
// only a subset of all the clusters. When the response contains only a subset of
// clusters, it includes a NextToken value. Use this value in a subsequent
// DescribeClusters request to get more clusters. When you receive a response with
// no NextToken (or an empty or null value), that means there are no more clusters
// to get.
//
// Cross-account use: No. You cannot perform this operation on CloudHSM clusters
// in a different Amazon Web Services account.
