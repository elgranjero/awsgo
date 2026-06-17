package sqs

// SendMessageBatch is generated as a reference stub.
// Executable command wiring lives under cmd/sqs.go.
//
// You can use SendMessageBatch to send up to 10 messages to the specified queue
// by assigning either identical or different values to each message (or by not
// assigning values at all). This is a batch version of SendMessage. For a FIFO queue,
// multiple messages within a single batch are enqueued in the order they are sent.
//
// The result of sending each message is reported individually in the response.
// Because the batch request can result in a combination of successful and
// unsuccessful actions, you should check for batch errors even when the call
// returns an HTTP status code of 200 .
//
// The maximum allowed individual message size and the maximum total payload size
// (the sum of the individual lengths of all of the batched messages) are both 1
// MiB 1,048,576 bytes.
//
// A message can include only XML, JSON, and unformatted text. The following
// Unicode characters are allowed. For more information, see the [W3C specification for characters].
//
// #x9 | #xA | #xD | #x20 to #xD7FF | #xE000 to #xFFFD | #x10000 to #x10FFFF
//
// If a message contains characters outside the allowed set, Amazon SQS rejects
// the message and returns an InvalidMessageContents error. Ensure that your
// message body includes only valid characters to avoid this exception.
//
// If you don't specify the DelaySeconds parameter for an entry, Amazon SQS uses
// the default value for the queue.
//
// [W3C specification for characters]: http://www.w3.org/TR/REC-xml/#charsets
