package sqs

// CancelMessageMoveTask is generated as a reference stub.
// Executable command wiring lives under cmd/sqs.go.
//
// Cancels a specified message movement task. A message movement can only be
// cancelled when the current status is RUNNING. Cancelling a message movement task
// does not revert the messages that have already been moved. It can only stop the
// messages that have not been moved yet.
//
// - This action is currently limited to supporting message redrive from [dead-letter queues (DLQs)]only.
// In this context, the source queue is the dead-letter queue (DLQ), while the
// destination queue can be the original source queue (from which the messages were
// driven to the dead-letter-queue), or a custom destination queue.
//
// - Only one active message movement task is supported per queue at any given
// time.
//
// [dead-letter queues (DLQs)]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html
