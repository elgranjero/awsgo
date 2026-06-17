package kendra

// Query is generated as a reference stub.
// Executable command wiring lives under cmd/kendra.go.
//
// Searches an index given an input query.
//
// If you are working with large language models (LLMs) or implementing retrieval
// augmented generation (RAG) systems, you can use Amazon Kendra's [Retrieve]API, which can
// return longer semantically relevant passages. We recommend using the Retrieve
// API instead of filing a service limit increase to increase the Query API
// document excerpt length.
//
// You can configure boosting or relevance tuning at the query level to override
// boosting at the index level, filter based on document fields/attributes and
// faceted search, and filter based on the user or their group access to documents.
// You can also include certain fields in the response that might provide useful
// additional information.
//
// A query response contains three types of results.
//
// - Relevant suggested answers. The answers can be either a text excerpt or
// table excerpt. The answer can be highlighted in the excerpt.
//
// - Matching FAQs or questions-answer from your FAQ file.
//
// - Relevant documents. This result type includes an excerpt of the document
// with the document title. The searched terms can be highlighted in the excerpt.
//
// You can specify that the query return only one type of result using the
// QueryResultTypeFilter parameter. Each query returns the 100 most relevant
// results. If you filter result type to only question-answers, a maximum of four
// results are returned. If you filter result type to only answers, a maximum of
// three results are returned.
//
// If you're using an Amazon Kendra Gen AI Enterprise Edition index, you can only
// use ATTRIBUTE_FILTER to filter search results by user context. If you're using
// an Amazon Kendra Gen AI Enterprise Edition index and you try to use USER_TOKEN
// to configure user context policy, Amazon Kendra returns a ValidationException
// error.
//
// [Retrieve]: https://docs.aws.amazon.com/kendra/latest/APIReference/API_Retrieve.html
