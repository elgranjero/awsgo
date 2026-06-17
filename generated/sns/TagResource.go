package sns

// TagResource is generated as a reference stub.
// Executable command wiring lives under cmd/sns.go.
//
// Add tags to the specified Amazon SNS topic. For an overview, see [Amazon SNS Tags] in the Amazon
// SNS Developer Guide.
//
// When you use topic tags, keep the following guidelines in mind:
//
// - Adding more than 50 tags to a topic isn't recommended.
//
// - Tags don't have any semantic meaning. Amazon SNS interprets tags as
// character strings.
//
// - Tags are case-sensitive.
//
// - A new tag with a key identical to that of an existing tag overwrites the
// existing tag.
//
// - Tagging actions are limited to 10 TPS per Amazon Web Services account, per
// Amazon Web Services Region. If your application requires a higher throughput,
// file a [technical support request].
//
// [Amazon SNS Tags]: https://docs.aws.amazon.com/sns/latest/dg/sns-tags.html
// [technical support request]: https://console.aws.amazon.com/support/home#/case/create?issueType=technical
