package frauddetector

// ListEventPredictions is generated as a reference stub.
// Executable command wiring lives under cmd/frauddetector.go.
//
// Gets a list of past predictions. The list can be filtered by detector ID,
// detector version ID, event ID, event type, or by specifying a time period. If
// filter is not specified, the most recent prediction is returned.
//
// For example, the following filter lists all past predictions for xyz event type
// - { "eventType":{ "value": "xyz" }” }
//
// This is a paginated API. If you provide a null maxResults , this action will
// retrieve a maximum of 10 records per page. If you provide a maxResults , the
// value must be between 50 and 100. To get the next page results, provide the
// nextToken from the response as part of your request. A null nextToken fetches
// the records from the beginning.
