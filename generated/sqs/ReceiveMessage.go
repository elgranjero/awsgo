package sqs

// ReceiveMessage is generated as a reference stub.
// Executable command wiring lives under cmd/sqs.go.
//
// Retrieves one or more messages (up to 10), from the specified queue. Using the
// WaitTimeSeconds parameter enables long-poll support. For more information, see [Amazon SQS Long Polling]
// in the Amazon SQS Developer Guide.
//
// Short poll is the default behavior where a weighted random set of machines is
// sampled on a ReceiveMessage call. Therefore, only the messages on the sampled
// machines are returned. If the number of messages in the queue is small (fewer
// than 1,000), you most likely get fewer messages than you requested per
// ReceiveMessage call. If the number of messages in the queue is extremely small,
// you might not receive any messages in a particular ReceiveMessage response. If
// this happens, repeat the request.
//
// For each message returned, the response includes the following:
//
// - The message body.
//
// - An MD5 digest of the message body. For information about MD5, see [RFC1321].
//
// - The MessageId you received when you sent the message to the queue.
//
// - The receipt handle.
//
// - The message attributes.
//
// - An MD5 digest of the message attributes.
//
// The receipt handle is the identifier you must provide when deleting the
// message. For more information, see [Queue and Message Identifiers]in the Amazon SQS Developer Guide.
//
// You can provide the VisibilityTimeout parameter in your request. The parameter
// is applied to the messages that Amazon SQS returns in the response. If you don't
// include the parameter, the overall visibility timeout for the queue is used for
// the returned messages. The default visibility timeout for a queue is 30 seconds.
//
// In the future, new attributes might be added. If you write code that calls this
// action, we recommend that you structure your code so that it can handle new
// attributes gracefully.
//
// [Queue and Message Identifiers]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-queue-message-identifiers.html
// [Amazon SQS Long Polling]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-long-polling.html
// [RFC1321]: https://www.ietf.org/rfc/rfc1321.txt
