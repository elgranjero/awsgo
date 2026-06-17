package frauddetector

// GetExternalModels is generated as a reference stub.
// Executable command wiring lives under cmd/frauddetector.go.
//
// Gets the details for one or more Amazon SageMaker models that have been
// imported into the service. This is a paginated API. If you provide a null
// maxResults , this actions retrieves a maximum of 10 records per page. If you
// provide a maxResults , the value must be between 5 and 10. To get the next page
// results, provide the pagination token from the GetExternalModelsResult as part
// of your request. A null pagination token fetches the records from the beginning.
