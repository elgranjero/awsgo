package kendra

// UpdateQuerySuggestionsConfig is generated as a reference stub.
// Executable command wiring lives under cmd/kendra.go.
//
// Updates the settings of query suggestions for an index.
//
// Amazon Kendra supports partial updates, so you only need to provide the fields
// you want to update.
//
// If an update is currently processing, you need to wait for the update to finish
// before making another update.
//
// Updates to query suggestions settings might not take effect right away. The
// time for your updated settings to take effect depends on the updates made and
// the number of search queries in your index.
//
// You can still enable/disable query suggestions at any time.
//
// UpdateQuerySuggestionsConfig is currently not supported in the Amazon Web
// Services GovCloud (US-West) region.
