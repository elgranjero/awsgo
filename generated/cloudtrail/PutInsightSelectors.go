package cloudtrail

// PutInsightSelectors is generated as a reference stub.
// Executable command wiring lives under cmd/cloudtrail.go.
//
// Lets you enable Insights event logging on specific event categories by
// specifying the Insights selectors that you want to enable on an existing trail
// or event data store. You also use PutInsightSelectors to turn off Insights
// event logging, by passing an empty list of Insights types. The valid Insights
// event types are ApiErrorRateInsight and ApiCallRateInsight , and valid
// EventCategories are Management and Data .
//
// Insights on data events are not supported on event data stores. For event data
// stores, you can only enable Insights on management events.
//
// To enable Insights on an event data store, you must specify the ARNs (or ID
// suffix of the ARNs) for the source event data store ( EventDataStore ) and the
// destination event data store ( InsightsDestination ). The source event data
// store logs management events and enables Insights. The destination event data
// store logs Insights events based upon the management event activity of the
// source event data store. The source and destination event data stores must
// belong to the same Amazon Web Services account.
//
// To log Insights events for a trail, you must specify the name ( TrailName ) of
// the CloudTrail trail for which you want to change or add Insights selectors.
//
// - For Management events Insights: To log CloudTrail Insights on the API call
// rate, the trail or event data store must log write management events. To log
// CloudTrail Insights on the API error rate, the trail or event data store must
// log read or write management events.
//
// - For Data events Insights: To log CloudTrail Insights on the API call rate
// or API error rate, the trail must log read or write data events. Data events
// Insights are not supported on event data store.
//
// To log CloudTrail Insights events on API call volume, the trail or event data
// store must log write management events. To log CloudTrail Insights events on
// API error rate, the trail or event data store must log read or write management
// events. You can call GetEventSelectors on a trail to check whether the trail
// logs management events. You can call GetEventDataStore on an event data store
// to check whether the event data store logs management events.
//
// For more information, see [Working with CloudTrail Insights] in the CloudTrail User Guide.
//
// [Working with CloudTrail Insights]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-insights-events-with-cloudtrail.html
