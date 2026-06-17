package fsx

// DescribeFileCaches is generated as a reference stub.
// Executable command wiring lives under cmd/fsx.go.
//
// Returns the description of a specific Amazon File Cache resource, if a
// FileCacheIds value is provided for that cache. Otherwise, it returns
// descriptions of all caches owned by your Amazon Web Services account in the
// Amazon Web Services Region of the endpoint that you're calling.
//
// When retrieving all cache descriptions, you can optionally specify the
// MaxResults parameter to limit the number of descriptions in a response. If more
// cache descriptions remain, the operation returns a NextToken value in the
// response. In this case, send a later request with the NextToken request
// parameter set to the value of NextToken from the last response.
//
// This operation is used in an iterative process to retrieve a list of your cache
// descriptions. DescribeFileCaches is called first without a NextToken value. Then
// the operation continues to be called with the NextToken parameter set to the
// value of the last NextToken value until a response has no NextToken .
//
// When using this operation, keep the following in mind:
//
// - The implementation might return fewer than MaxResults cache descriptions
// while still including a NextToken value.
//
// - The order of caches returned in the response of one DescribeFileCaches call
// and the order of caches returned across the responses of a multicall iteration
// is unspecified.
