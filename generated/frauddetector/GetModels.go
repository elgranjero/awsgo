package frauddetector

// GetModels is generated as a reference stub.
// Executable command wiring lives under cmd/frauddetector.go.
//
// Gets one or more models. Gets all models for the Amazon Web Services account if
// no model type and no model id provided. Gets all models for the Amazon Web
// Services account and model type, if the model type is specified but model id is
// not provided. Gets a specific model if (model type, model id) tuple is
// specified.
//
// This is a paginated API. If you provide a null maxResults , this action
// retrieves a maximum of 10 records per page. If you provide a maxResults , the
// value must be between 1 and 10. To get the next page results, provide the
// pagination token from the response as part of your request. A null pagination
// token fetches the records from the beginning.
