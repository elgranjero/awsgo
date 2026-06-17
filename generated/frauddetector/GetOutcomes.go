package frauddetector

// GetOutcomes is generated as a reference stub.
// Executable command wiring lives under cmd/frauddetector.go.
//
// Gets one or more outcomes. This is a paginated API. If you provide a null
// maxResults , this actions retrieves a maximum of 100 records per page. If you
// provide a maxResults , the value must be between 50 and 100. To get the next
// page results, provide the pagination token from the GetOutcomesResult as part
// of your request. A null pagination token fetches the records from the beginning.
