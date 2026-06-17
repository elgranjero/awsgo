package cloudtrail

// ListInsightsData is generated as a reference stub.
// Executable command wiring lives under cmd/cloudtrail.go.
//
// Returns Insights events generated on a trail that logs data events. You can
// list Insights events that occurred in a Region within the last 90 days.
//
// ListInsightsData supports the following Dimensions for Insights events:
//
// - Event ID
//
// - Event name
//
// - Event source
//
// All dimensions are optional. The default number of results returned is 50, with
// a maximum of 50 possible. The response includes a token that you can use to get
// the next page of results.
//
// The rate of ListInsightsData requests is limited to two per second, per
// account, per Region. If this limit is exceeded, a throttling error occurs.
