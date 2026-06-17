package sqs

// StartMessageMoveTask is generated as a reference stub.
// Executable command wiring lives under cmd/sqs.go.
//
// Starts an asynchronous task to move messages from a specified source queue to a
// specified destination queue.
//
// - This action is currently limited to supporting message redrive from queues
// that are configured as [dead-letter queues (DLQs)]of other Amazon SQS queues only. Non-SQS queue sources
// of dead-letter queues, such as Lambda or Amazon SNS topics, are currently not
// supported.
//
// - In dead-letter queues redrive context, the StartMessageMoveTask the source
// queue is the DLQ, while the destination queue can be the original source queue
// (from which the messages were driven to the dead-letter-queue), or a custom
// destination queue.
//
// - Only one active message movement task is supported per queue at any given
// time.
//
// [dead-letter queues (DLQs)]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html
