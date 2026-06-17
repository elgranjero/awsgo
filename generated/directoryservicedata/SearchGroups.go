package directoryservicedata

// SearchGroups is generated as a reference stub.
// Executable command wiring lives under cmd/directoryservicedata.go.
//
// Searches the specified directory for a group. You can find groups that match
//
// the SearchString parameter with the value of their attributes included in the
// SearchString parameter.
//
// This operation supports pagination with the use of the NextToken request and
// response parameters. If more results are available, the SearchGroups.NextToken
// member contains a token that you pass in the next call to SearchGroups . This
// retrieves the next set of items.
//
// You can also specify a maximum number of return results with the MaxResults
// parameter.
