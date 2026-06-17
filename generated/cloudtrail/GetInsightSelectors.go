package cloudtrail

// GetInsightSelectors is generated as a reference stub.
// Executable command wiring lives under cmd/cloudtrail.go.
//
// Describes the settings for the Insights event selectors that you configured for
// your trail or event data store. GetInsightSelectors shows if CloudTrail
// Insights logging is enabled and which Insights types are configured with
// corresponding event categories. If you run GetInsightSelectors on a trail or
// event data store that does not have Insights events enabled, the operation
// throws the exception InsightNotEnabledException
//
// Specify either the EventDataStore parameter to get Insights event selectors for
// an event data store, or the TrailName parameter to the get Insights event
// selectors for a trail. You cannot specify these parameters together.
//
// For more information, see [Working with CloudTrail Insights] in the CloudTrail User Guide.
//
// [Working with CloudTrail Insights]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-insights-events-with-cloudtrail.html
