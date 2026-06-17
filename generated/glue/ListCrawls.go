package glue

// ListCrawls is generated as a reference stub.
// Executable command wiring lives under cmd/glue.go.
//
// Returns all the crawls of a specified crawler. Returns only the crawls that
// have occurred since the launch date of the crawler history feature, and only
// retains up to 12 months of crawls. Older crawls will not be returned.
//
// You may use this API to:
//
// - Retrive all the crawls of a specified crawler.
//
// - Retrieve all the crawls of a specified crawler within a limited count.
//
// - Retrieve all the crawls of a specified crawler in a specific time range.
//
// - Retrieve all the crawls of a specified crawler with a particular state,
// crawl ID, or DPU hour value.
