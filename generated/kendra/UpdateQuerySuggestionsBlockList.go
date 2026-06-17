package kendra

// UpdateQuerySuggestionsBlockList is generated as a reference stub.
// Executable command wiring lives under cmd/kendra.go.
//
// Updates a block list used for query suggestions for an index.
//
// Updates to a block list might not take effect right away. Amazon Kendra needs
// to refresh the entire suggestions list to apply any updates to the block list.
// Other changes not related to the block list apply immediately.
//
// If a block list is updating, then you need to wait for the first update to
// finish before submitting another update.
//
// Amazon Kendra supports partial updates, so you only need to provide the fields
// you want to update.
//
// UpdateQuerySuggestionsBlockList is currently not supported in the Amazon Web
// Services GovCloud (US-West) region.
