package sqs

// DeleteQueue is generated as a reference stub.
// Executable command wiring lives under cmd/sqs.go.
//
// Deletes the queue specified by the QueueUrl , regardless of the queue's contents.
//
// Be careful with the DeleteQueue action: When you delete a queue, any messages
// in the queue are no longer available.
//
// When you delete a queue, the deletion process takes up to 60 seconds. Requests
// you send involving that queue during the 60 seconds might succeed. For example,
// a SendMessagerequest might succeed, but after 60 seconds the queue and the message you
// sent no longer exist.
//
// When you delete a queue, you must wait at least 60 seconds before creating a
// queue with the same name.
//
// Cross-account permissions don't apply to this action. For more information, see [Grant cross-account permissions to a role and a username]
// in the Amazon SQS Developer Guide.
//
// The delete operation uses the HTTP GET verb.
//
// [Grant cross-account permissions to a role and a username]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name
