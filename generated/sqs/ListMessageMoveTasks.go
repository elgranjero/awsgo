package sqs

// ListMessageMoveTasks is generated as a reference stub.
// Executable command wiring lives under cmd/sqs.go.
//
// Gets the most recent message movement tasks (up to 10) under a specific source
// queue.
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
