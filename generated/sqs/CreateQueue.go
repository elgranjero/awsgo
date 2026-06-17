package sqs

// CreateQueue is generated as a reference stub.
// Executable command wiring lives under cmd/sqs.go.
//
// Creates a new standard or FIFO queue. You can pass one or more attributes in
// the request. Keep the following in mind:
//
// - If you don't specify the FifoQueue attribute, Amazon SQS creates a standard
// queue.
//
// You can't change the queue type after you create it and you can't convert an
//
// existing standard queue into a FIFO queue. You must either create a new FIFO
// queue for your application or delete your existing standard queue and recreate
// it as a FIFO queue. For more information, see [Moving From a standard queue to a FIFO queue]in the Amazon SQS Developer
// Guide.
//
// - If you don't provide a value for an attribute, the queue is created with
// the default value for the attribute.
//
// - If you delete a queue, you must wait at least 60 seconds before creating a
// queue with the same name.
//
// To successfully create a new queue, you must provide a queue name that adheres
// to the [limits related to queues]and is unique within the scope of your queues.
//
// After you create a queue, you must wait at least one second after the queue is
// created to be able to use the queue.
//
// To retrieve the URL of a queue, use the [GetQueueUrl]GetQueueUrl action. This action only
// requires the [QueueName]QueueName parameter.
//
// When creating queues, keep the following points in mind:
//
// - If you specify the name of an existing queue and provide the exact same
// names and values for all its attributes, the [CreateQueue]CreateQueue action will return
// the URL of the existing queue instead of creating a new one.
//
// - If you attempt to create a queue with a name that already exists but with
// different attribute names or values, the CreateQueue action will return an
// error. This ensures that existing queues are not inadvertently altered.
//
// Cross-account permissions don't apply to this action. For more information, see [Grant cross-account permissions to a role and a username]
// in the Amazon SQS Developer Guide.
//
// [limits related to queues]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/limits-queues.html
// [CreateQueue]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_CreateQueue.html
// [Grant cross-account permissions to a role and a username]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name
// [QueueName]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_CreateQueue.html#API_CreateQueue_RequestSyntax
// [GetQueueUrl]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_GetQueueUrl.html
//
// [Moving From a standard queue to a FIFO queue]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html#FIFO-queues-moving
