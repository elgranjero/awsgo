package kendra

// CreateQuerySuggestionsBlockList is generated as a reference stub.
// Executable command wiring lives under cmd/kendra.go.
//
// Creates a block list to exlcude certain queries from suggestions.
//
// Any query that contains words or phrases specified in the block list is blocked
// or filtered out from being shown as a suggestion.
//
// You need to provide the file location of your block list text file in your S3
// bucket. In your text file, enter each block word or phrase on a separate line.
//
// For information on the current quota limits for block lists, see [Quotas for Amazon Kendra].
//
// CreateQuerySuggestionsBlockList is currently not supported in the Amazon Web
// Services GovCloud (US-West) region.
//
// For an example of creating a block list for query suggestions using the Python
// SDK, see [Query suggestions block list].
//
// [Quotas for Amazon Kendra]: https://docs.aws.amazon.com/kendra/latest/dg/quotas.html
// [Query suggestions block list]: https://docs.aws.amazon.com/kendra/latest/dg/query-suggestions.html#query-suggestions-blocklist
