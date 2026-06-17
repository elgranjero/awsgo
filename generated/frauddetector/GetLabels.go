package frauddetector

// GetLabels is generated as a reference stub.
// Executable command wiring lives under cmd/frauddetector.go.
//
// Gets all labels or a specific label if name is provided. This is a paginated
// API. If you provide a null maxResults , this action retrieves a maximum of 50
// records per page. If you provide a maxResults , the value must be between 10 and
// 50. To get the next page results, provide the pagination token from the
// GetGetLabelsResponse as part of your request. A null pagination token fetches
// the records from the beginning.
