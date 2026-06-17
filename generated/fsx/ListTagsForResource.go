package fsx

// ListTagsForResource is generated as a reference stub.
// Executable command wiring lives under cmd/fsx.go.
//
// Lists tags for Amazon FSx resources.
//
// When retrieving all tags, you can optionally specify the MaxResults parameter
// to limit the number of tags in a response. If more tags remain, Amazon FSx
// returns a NextToken value in the response. In this case, send a later request
// with the NextToken request parameter set to the value of NextToken from the
// last response.
//
// This action is used in an iterative process to retrieve a list of your tags.
// ListTagsForResource is called first without a NextToken value. Then the action
// continues to be called with the NextToken parameter set to the value of the
// last NextToken value until a response has no NextToken .
//
// When using this action, keep the following in mind:
//
// - The implementation might return fewer than MaxResults file system
// descriptions while still including a NextToken value.
//
// - The order of tags returned in the response of one ListTagsForResource call
// and the order of tags returned across the responses of a multi-call iteration is
// unspecified.
