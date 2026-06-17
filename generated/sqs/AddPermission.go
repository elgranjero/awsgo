package sqs

// AddPermission is generated as a reference stub.
// Executable command wiring lives under cmd/sqs.go.
//
// Adds a permission to a queue for a specific [principal]. This allows sharing access to the
// queue.
//
// When you create a queue, you have full control access rights for the queue.
// Only you, the owner of the queue, can grant or deny permissions to the queue.
// For more information about these permissions, see [Allow Developers to Write Messages to a Shared Queue]in the Amazon SQS Developer
// Guide.
//
// - AddPermission generates a policy for you. You can use SetQueueAttributesto upload your
// policy. For more information, see [Using Custom Policies with the Amazon SQS Access Policy Language]in the Amazon SQS Developer Guide.
//
// - An Amazon SQS policy can have a maximum of seven actions per statement.
//
// - To remove the ability to change queue permissions, you must deny permission
// to the AddPermission , RemovePermission , and SetQueueAttributes actions in
// your IAM policy.
//
// - Amazon SQS AddPermission does not support adding a non-account principal.
//
// Cross-account permissions don't apply to this action. For more information, see [Grant cross-account permissions to a role and a username]
// in the Amazon SQS Developer Guide.
//
// [principal]: https://docs.aws.amazon.com/general/latest/gr/glos-chap.html#P
// [Allow Developers to Write Messages to a Shared Queue]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-writing-an-sqs-policy.html#write-messages-to-shared-queue
// [Using Custom Policies with the Amazon SQS Access Policy Language]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-creating-custom-policies.html
// [Grant cross-account permissions to a role and a username]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name
