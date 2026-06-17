package fsx

// DescribeFileSystems is generated as a reference stub.
// Executable command wiring lives under cmd/fsx.go.
//
// Returns the description of specific Amazon FSx file systems, if a FileSystemIds
// value is provided for that file system. Otherwise, it returns descriptions of
// all file systems owned by your Amazon Web Services account in the Amazon Web
// Services Region of the endpoint that you're calling.
//
// When retrieving all file system descriptions, you can optionally specify the
// MaxResults parameter to limit the number of descriptions in a response. If more
// file system descriptions remain, Amazon FSx returns a NextToken value in the
// response. In this case, send a later request with the NextToken request
// parameter set to the value of NextToken from the last response.
//
// This operation is used in an iterative process to retrieve a list of your file
// system descriptions. DescribeFileSystems is called first without a NextToken
// value. Then the operation continues to be called with the NextToken parameter
// set to the value of the last NextToken value until a response has no NextToken .
//
// When using this operation, keep the following in mind:
//
// - The implementation might return fewer than MaxResults file system
// descriptions while still including a NextToken value.
//
// - The order of file systems returned in the response of one
// DescribeFileSystems call and the order of file systems returned across the
// responses of a multicall iteration is unspecified.
