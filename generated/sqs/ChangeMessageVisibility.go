package sqs

// ChangeMessageVisibility is generated as a reference stub.
// Executable command wiring lives under cmd/sqs.go.
//
// Changes the visibility timeout of a specified message in a queue to a new
// value. The default visibility timeout for a message is 30 seconds. The minimum
// is 0 seconds. The maximum is 12 hours. For more information, see [Visibility Timeout]in the Amazon
// SQS Developer Guide.
//
// For example, if the default timeout for a queue is 60 seconds, 15 seconds have
// elapsed since you received the message, and you send a ChangeMessageVisibility
// call with VisibilityTimeout set to 10 seconds, the 10 seconds begin to count
// from the time that you make the ChangeMessageVisibility call. Thus, any attempt
// to change the visibility timeout or to delete that message 10 seconds after you
// initially change the visibility timeout (a total of 25 seconds) might result in
// an error.
//
// An Amazon SQS message has three basic states:
//
// - Sent to a queue by a producer.
//
// - Received from the queue by a consumer.
//
// - Deleted from the queue.
//
// A message is considered to be stored after it is sent to a queue by a producer,
// but not yet received from the queue by a consumer (that is, between states 1 and
// 2). There is no limit to the number of stored messages. A message is considered
// to be in flight after it is received from a queue by a consumer, but not yet
// deleted from the queue (that is, between states 2 and 3). There is a limit to
// the number of in flight messages.
//
// Limits that apply to in flight messages are unrelated to the unlimited number
// of stored messages.
//
// For most standard queues (depending on queue traffic and message backlog),
// there can be a maximum of approximately 120,000 in flight messages (received
// from a queue by a consumer, but not yet deleted from the queue). If you reach
// this limit, Amazon SQS returns the OverLimit error message. To avoid reaching
// the limit, you should delete messages from the queue after they're processed.
// You can also increase the number of queues you use to process your messages. To
// request a limit increase, [file a support request].
//
// For FIFO queues, there can be a maximum of 120,000 in flight messages (received
// from a queue by a consumer, but not yet deleted from the queue). If you reach
// this limit, Amazon SQS returns no error messages.
//
// If you attempt to set the VisibilityTimeout to a value greater than the maximum
// time left, Amazon SQS returns an error. Amazon SQS doesn't automatically
// recalculate and increase the timeout to the maximum remaining time.
//
// Unlike with a queue, when you change the visibility timeout for a specific
// message the timeout value is applied immediately but isn't saved in memory for
// that message. If you don't delete a message after it is received, the visibility
// timeout for the message reverts to the original timeout value (not to the value
// you set using the ChangeMessageVisibility action) the next time the message is
// received.
//
// [Visibility Timeout]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-visibility-timeout.html
// [file a support request]: https://console.aws.amazon.com/support/home#/case/create?issueType=service-limit-increase&limitType=service-code-sqs
