package securityhub

// GetFindingHistory is generated as a reference stub.
// Executable command wiring lives under cmd/securityhub.go.
//
// Returns the history of a Security Hub CSPM finding. The history includes
//
// changes made to any fields in the Amazon Web Services Security Finding Format
// (ASFF) except top-level timestamp fields, such as the CreatedAt and UpdatedAt
// fields.
//
// This operation might return fewer results than the maximum number of results (
// MaxResults ) specified in a request, even when more results are available. If
// this occurs, the response includes a NextToken value, which you should use to
// retrieve the next set of results in the response. The presence of a NextToken
// value in a response doesn't necessarily indicate that the results are
// incomplete. However, you should continue to specify a NextToken value until you
// receive a response that doesn't include this value.
