package lexmodelsv2

// ListIntentStageMetrics is generated as a reference stub.
// Executable command wiring lives under cmd/lexmodelsv2.go.
//
// Retrieves summary metrics for the stages within intents in your bot. The
// following fields are required:
//
// - metrics – A list of [AnalyticsIntentStageMetric]objects. In each object, use the name field to specify
// the metric to calculate, the statistic field to specify whether to calculate
// the Sum , Average , or Max number, and the order field to specify whether to
// sort the results in Ascending or Descending order.
//
// - startDateTime and endDateTime – Define a time range for which you want to
// retrieve results.
//
// Of the optional fields, you can organize the results in the following ways:
//
// - Use the filters field to filter the results, the groupBy field to specify
// categories by which to group the results, and the binBy field to specify time
// intervals by which to group the results.
//
// - Use the maxResults field to limit the number of results to return in a
// single response and the nextToken field to return the next batch of results if
// the response does not return the full set of results.
//
// Note that an order field exists in both binBy and metrics . You can only specify
// one order in a given request.
//
// [AnalyticsIntentStageMetric]: https://docs.aws.amazon.com/lexv2/latest/APIReference/API_AnalyticsIntentStageMetric.html
