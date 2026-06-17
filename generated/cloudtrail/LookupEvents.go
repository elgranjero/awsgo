package cloudtrail

// LookupEvents is generated as a reference stub.
// Executable command wiring lives under cmd/cloudtrail.go.
//
// Looks up [management events] or [CloudTrail Insights events] that are captured by CloudTrail. You can look up events that
// occurred in a Region within the last 90 days.
//
// LookupEvents returns recent Insights events for trails that enable Insights. To
// view Insights events for an event data store, you can run queries on your
// Insights event data store, and you can also view the Lake dashboard for
// Insights.
//
// Lookup supports the following attributes for management events:
//
// - Amazon Web Services access key
//
// - Event ID
//
// - Event name
//
// - Event source
//
// - Read only
//
// - Resource name
//
// - Resource type
//
// - User name
//
// Lookup supports the following attributes for Insights events:
//
// - Event ID
//
// - Event name
//
// - Event source
//
// All attributes are optional. The default number of results returned is 50, with
// a maximum of 50 possible. The response includes a token that you can use to get
// the next page of results.
//
// The rate of lookup requests is limited to two per second, per account, per
// Region. If this limit is exceeded, a throttling error occurs.
//
// [CloudTrail Insights events]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-concepts.html#cloudtrail-concepts-insights-events
// [management events]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-concepts.html#cloudtrail-concepts-management-events
