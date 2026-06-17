package directoryservice

// DescribeDirectories is generated as a reference stub.
// Executable command wiring lives under cmd/directoryservice.go.
//
// Obtains information about the directories that belong to this account.
//
// You can retrieve information about specific directories by passing the
// directory identifiers in the DirectoryIds parameter. Otherwise, all directories
// that belong to the current account are returned.
//
// This operation supports pagination with the use of the NextToken request and
// response parameters. If more results are available, the
// DescribeDirectoriesResult.NextToken member contains a token that you pass in the
// next call to DescribeDirectoriesto retrieve the next set of items.
//
// You can also specify a maximum number of return results with the Limit
// parameter.
