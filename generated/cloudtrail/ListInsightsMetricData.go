package cloudtrail

// ListInsightsMetricData is generated as a reference stub.
// Executable command wiring lives under cmd/cloudtrail.go.
//
// Returns Insights metrics data for trails that have enabled Insights. The
// request must include the EventSource , EventName , and InsightType parameters.
//
// If the InsightType is set to ApiErrorRateInsight , the request must also include
// the ErrorCode parameter.
//
// The following are the available time periods for ListInsightsMetricData . Each
// cutoff is inclusive.
//
// - Data points with a period of 60 seconds (1-minute) are available for 15
// days.
//
// - Data points with a period of 300 seconds (5-minute) are available for 63
// days.
//
// - Data points with a period of 3600 seconds (1 hour) are available for 90
// days.
//
// To use ListInsightsMetricData operation, you must have the following
// permissions:
//
// - If ListInsightsMetricData is invoked with TrailName parameter, access to the
// ListInsightsMetricData API operation is linked to the cloudtrail:LookupEvents
// action and cloudtrail:ListInsightsData . To use this operation, you must have
// permissions to perform the cloudtrail:LookupEvents and
// cloudtrail:ListInsightsData action on the specific trail.
//
// - If ListInsightsMetricData is invoked without TrailName parameter, access to
// the ListInsightsMetricData API operation is linked to the
// cloudtrail:LookupEvents action only. To use this operation, you must have
// permissions to perform the cloudtrail:LookupEvents action.
