package frauddetector

// GetEventTypes is generated as a reference stub.
// Executable command wiring lives under cmd/frauddetector.go.
//
// Gets all event types or a specific event type if name is provided. This is a
// paginated API. If you provide a null maxResults , this action retrieves a
// maximum of 10 records per page. If you provide a maxResults , the value must be
// between 5 and 10. To get the next page results, provide the pagination token
// from the GetEventTypesResponse as part of your request. A null pagination token
// fetches the records from the beginning.
