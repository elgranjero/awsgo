package sqs

// RemovePermission is generated as a reference stub.
// Executable command wiring lives under cmd/sqs.go.
//
// Revokes any permissions in the queue policy that matches the specified Label
// parameter.
//
// - Only the owner of a queue can remove permissions from it.
//
// - Cross-account permissions don't apply to this action. For more information,
// see [Grant cross-account permissions to a role and a username]in the Amazon SQS Developer Guide.
//
// - To remove the ability to change queue permissions, you must deny permission
// to the AddPermission , RemovePermission , and SetQueueAttributes actions in
// your IAM policy.
//
// [Grant cross-account permissions to a role and a username]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name
