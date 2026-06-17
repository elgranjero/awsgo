package sqs

// PurgeQueue is generated as a reference stub.
// Executable command wiring lives under cmd/sqs.go.
//
// Deletes available messages in a queue (including in-flight messages) specified
// by the QueueURL parameter.
//
// When you use the PurgeQueue action, you can't retrieve any messages deleted
// from a queue.
//
// The message deletion process takes up to 60 seconds. We recommend waiting for
// 60 seconds regardless of your queue's size.
//
// Messages sent to the queue before you call PurgeQueue might be received but are
// deleted within the next minute.
//
// Messages sent to the queue after you call PurgeQueue might be deleted while the
// queue is being purged.
