package sqs

// SetQueueAttributes is generated as a reference stub.
// Executable command wiring lives under cmd/sqs.go.
//
// Sets the value of one or more queue attributes, like a policy. When you change
// a queue's attributes, the change can take up to 60 seconds for most of the
// attributes to propagate throughout the Amazon SQS system. Changes made to the
// MessageRetentionPeriod attribute can take up to 15 minutes and will impact
// existing messages in the queue potentially causing them to be expired and
// deleted if the MessageRetentionPeriod is reduced below the age of existing
// messages.
//
// - In the future, new attributes might be added. If you write code that calls
// this action, we recommend that you structure your code so that it can handle new
// attributes gracefully.
//
// - Cross-account permissions don't apply to this action. For more information,
// see [Grant cross-account permissions to a role and a username]in the Amazon SQS Developer Guide.
//
// - To remove the ability to change queue permissions, you must deny permission
// to the AddPermission , RemovePermission , and SetQueueAttributes actions in
// your IAM policy.
//
// [Grant cross-account permissions to a role and a username]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name
