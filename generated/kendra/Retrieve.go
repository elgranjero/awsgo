package kendra

// Retrieve is generated as a reference stub.
// Executable command wiring lives under cmd/kendra.go.
//
// Retrieves relevant passages or text excerpts given an input query.
//
// This API is similar to the [Query] API. However, by default, the Query API only
// returns excerpt passages of up to 100 token words. With the Retrieve API, you
// can retrieve longer passages of up to 200 token words and up to 100 semantically
// relevant passages. This doesn't include question-answer or FAQ type responses
// from your index. The passages are text excerpts that can be semantically
// extracted from multiple documents and multiple parts of the same document. If in
// extreme cases your documents produce zero passages using the Retrieve API, you
// can alternatively use the Query API and its types of responses.
//
// You can also do the following:
//
// - Override boosting at the index level
//
// - Filter based on document fields or attributes
//
// - Filter based on the user or their group access to documents
//
// - View the confidence score bucket for a retrieved passage result. The
// confidence bucket provides a relative ranking that indicates how confident
// Amazon Kendra is that the response is relevant to the query.
//
// Confidence score buckets are currently available only for English.
//
// You can also include certain fields in the response that might provide useful
// additional information.
//
// The Retrieve API shares the number of [query capacity units] that you set for your index. For more
// information on what's included in a single capacity unit and the default base
// capacity for an index, see [Adjusting capacity].
//
// If you're using an Amazon Kendra Gen AI Enterprise Edition index, you can only
// use ATTRIBUTE_FILTER to filter search results by user context. If you're using
// an Amazon Kendra Gen AI Enterprise Edition index and you try to use USER_TOKEN
// to configure user context policy, Amazon Kendra returns a ValidationException
// error.
//
// [Adjusting capacity]: https://docs.aws.amazon.com/kendra/latest/dg/adjusting-capacity.html
// [Query]: https://docs.aws.amazon.com/kendra/latest/APIReference/API_Query.html
// [query capacity units]: https://docs.aws.amazon.com/kendra/latest/APIReference/API_CapacityUnitsConfiguration.html
