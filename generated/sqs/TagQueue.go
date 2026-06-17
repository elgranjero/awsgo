package sqs

// TagQueue is generated as a reference stub.
// Executable command wiring lives under cmd/sqs.go.
//
// Add cost allocation tags to the specified Amazon SQS queue. For an overview,
// see [Tagging Your Amazon SQS Queues]in the Amazon SQS Developer Guide.
//
// When you use queue tags, keep the following guidelines in mind:
//
// - Adding more than 50 tags to a queue isn't recommended.
//
// - Tags don't have any semantic meaning. Amazon SQS interprets tags as
// character strings.
//
// - Tags are case-sensitive.
//
// - A new tag with a key identical to that of an existing tag overwrites the
// existing tag.
//
// For a full list of tag restrictions, see [Quotas related to queues] in the Amazon SQS Developer Guide.
//
// Cross-account permissions don't apply to this action. For more information, see [Grant cross-account permissions to a role and a username]
// in the Amazon SQS Developer Guide.
//
// [Tagging Your Amazon SQS Queues]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-queue-tags.html
// [Quotas related to queues]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-limits.html#limits-queues
// [Grant cross-account permissions to a role and a username]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name
