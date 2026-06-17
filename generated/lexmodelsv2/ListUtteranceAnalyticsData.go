package lexmodelsv2

// ListUtteranceAnalyticsData is generated as a reference stub.
// Executable command wiring lives under cmd/lexmodelsv2.go.
//
// To use this API operation, your IAM role must have permissions to perform the [ListAggregatedUtterances]
// operation, which provides access to utterance-related analytics. See [Viewing utterance statistics]for the
// IAM policy to apply to the IAM role.
//
// Retrieves a list of metadata for individual user utterances to your bot. The
// following fields are required:
//
// - startDateTime and endDateTime – Define a time range for which you want to
// retrieve results.
//
// Of the optional fields, you can organize the results in the following ways:
//
// - Use the filters field to filter the results and the sortBy field to specify
// the values by which to sort the results.
//
// - Use the maxResults field to limit the number of results to return in a
// single response and the nextToken field to return the next batch of results if
// the response does not return the full set of results.
//
// [Viewing utterance statistics]: https://docs.aws.amazon.com/lexv2/latest/dg/monitoring-utterances.html
// [ListAggregatedUtterances]: https://docs.aws.amazon.com/lexv2/latest/APIReference/API_ListAggregatedUtterances.html
