package resourceexplorer2

// Search is generated as a reference stub.
// Executable command wiring lives under cmd/resourceexplorer2.go.
//
// Searches for resources and displays details about all resources that match the
// specified criteria. You must specify a query string.
//
// All search queries must use a view. If you don't explicitly specify a view,
// then Amazon Web Services Resource Explorer uses the default view for the Amazon
// Web Services Region in which you call this operation. The results are the
// logical intersection of the results that match both the QueryString parameter
// supplied to this operation and the SearchFilter parameter attached to the view.
//
// For the complete syntax supported by the QueryString parameter, see [Search query syntax reference for Resource Explorer].
//
// If your search results are empty, or are missing results that you think should
// be there, see [Troubleshooting Resource Explorer search].
//
// [Troubleshooting Resource Explorer search]: https://docs.aws.amazon.com/resource-explorer/latest/userguide/troubleshooting_search.html
// [Search query syntax reference for Resource Explorer]: https://docs.aws.amazon.com/resource-explorer/latest/APIReference/about-query-syntax.html
