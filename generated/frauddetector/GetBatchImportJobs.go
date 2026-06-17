package frauddetector

// GetBatchImportJobs is generated as a reference stub.
// Executable command wiring lives under cmd/frauddetector.go.
//
// Gets all batch import jobs or a specific job of the specified ID. This is a
// paginated API. If you provide a null maxResults , this action retrieves a
// maximum of 50 records per page. If you provide a maxResults , the value must be
// between 1 and 50. To get the next page results, provide the pagination token
// from the GetBatchImportJobsResponse as part of your request. A null pagination
// token fetches the records from the beginning.
