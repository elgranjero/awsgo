package fsx

// DescribeDataRepositoryTasks is generated as a reference stub.
// Executable command wiring lives under cmd/fsx.go.
//
// Returns the description of specific Amazon FSx for Lustre or Amazon File Cache
// data repository tasks, if one or more TaskIds values are provided in the
// request, or if filters are used in the request. You can use filters to narrow
// the response to include just tasks for specific file systems or caches, or tasks
// in a specific lifecycle state. Otherwise, it returns all data repository tasks
// owned by your Amazon Web Services account in the Amazon Web Services Region of
// the endpoint that you're calling.
//
// When retrieving all tasks, you can paginate the response by using the optional
// MaxResults parameter to limit the number of tasks returned in a response. If
// more tasks remain, a NextToken value is returned in the response. In this case,
// send a later request with the NextToken request parameter set to the value of
// NextToken from the last response.
