package fsx

// DescribeDataRepositoryAssociations is generated as a reference stub.
// Executable command wiring lives under cmd/fsx.go.
//
// Returns the description of specific Amazon FSx for Lustre or Amazon File Cache
// data repository associations, if one or more AssociationIds values are provided
// in the request, or if filters are used in the request. Data repository
// associations are supported on Amazon File Cache resources and all FSx for Lustre
// 2.12 and 2,15 file systems, excluding scratch_1 deployment type.
//
// You can use filters to narrow the response to include just data repository
// associations for specific file systems (use the file-system-id filter with the
// ID of the file system) or caches (use the file-cache-id filter with the ID of
// the cache), or data repository associations for a specific repository type (use
// the data-repository-type filter with a value of S3 or NFS ). If you don't use
// filters, the response returns all data repository associations owned by your
// Amazon Web Services account in the Amazon Web Services Region of the endpoint
// that you're calling.
//
// When retrieving all data repository associations, you can paginate the response
// by using the optional MaxResults parameter to limit the number of data
// repository associations returned in a response. If more data repository
// associations remain, a NextToken value is returned in the response. In this
// case, send a later request with the NextToken request parameter set to the
// value of NextToken from the last response.
