package kendra

// ClearQuerySuggestions is generated as a reference stub.
// Executable command wiring lives under cmd/kendra.go.
//
// Clears existing query suggestions from an index.
//
// This deletes existing suggestions only, not the queries in the query log. After
// you clear suggestions, Amazon Kendra learns new suggestions based on new queries
// added to the query log from the time you cleared suggestions. If you do not see
// any new suggestions, then please allow Amazon Kendra to collect enough queries
// to learn new suggestions.
//
// ClearQuerySuggestions is currently not supported in the Amazon Web Services
// GovCloud (US-West) region.
