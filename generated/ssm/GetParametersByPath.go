package ssm

// GetParametersByPath is generated as a reference stub.
// Executable command wiring lives under cmd/ssm.go.
//
// Retrieve information about one or more parameters under a specified level in a
// hierarchy.
//
// Request results are returned on a best-effort basis. If you specify MaxResults
// in the request, the response includes information up to the limit specified. The
// number of items returned, however, can be between zero and the value of
// MaxResults . If the service reaches an internal limit while processing the
// results, it stops the operation and returns the matching values up to that point
// and a NextToken . You can specify the NextToken in a subsequent call to get the
// next set of results.
//
// Parameter names can't contain spaces. The service removes any spaces specified
// for the beginning or end of a parameter name. If the specified name for a
// parameter contains spaces between characters, the request fails with a
// ValidationException error.
