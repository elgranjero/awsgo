package dynamodb

// TagResource is generated as a reference stub.
// Executable command wiring lives under cmd/dynamodb.go.
//
// Associate a set of tags with an Amazon DynamoDB resource. You can then activate
// these user-defined tags so that they appear on the Billing and Cost Management
// console for cost allocation tracking. You can call TagResource up to five times
// per second, per account.
//
// - TagResource is an asynchronous operation. If you issue a ListTagsOfResourcerequest
// immediately after a TagResource request, DynamoDB might return your previous
// tag set, if there was one, or an empty tag set. This is because
// ListTagsOfResource uses an eventually consistent query, and the metadata for
// your tags or table might not be available at that moment. Wait for a few
// seconds, and then try the ListTagsOfResource request again.
//
// - The application or removal of tags using TagResource and UntagResource APIs
// is eventually consistent. ListTagsOfResource API will only reflect the changes
// after a few seconds.
//
// For an overview on tagging DynamoDB resources, see [Tagging for DynamoDB] in the Amazon DynamoDB
// Developer Guide.
//
// [Tagging for DynamoDB]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Tagging.html
