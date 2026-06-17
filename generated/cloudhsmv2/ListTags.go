package cloudhsmv2

// ListTags is generated as a reference stub.
// Executable command wiring lives under cmd/cloudhsmv2.go.
//
// Gets a list of tags for the specified CloudHSM cluster.
//
// This is a paginated operation, which means that each response might contain
// only a subset of all the tags. When the response contains only a subset of tags,
// it includes a NextToken value. Use this value in a subsequent ListTags request
// to get more tags. When you receive a response with no NextToken (or an empty or
// null value), that means there are no more tags to get.
//
// Cross-account use: No. You cannot perform this operation on an CloudHSM
// resource in a different Amazon Web Services account.
