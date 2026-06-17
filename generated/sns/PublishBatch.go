package sns

// PublishBatch is generated as a reference stub.
// Executable command wiring lives under cmd/sns.go.
//
// Publishes up to 10 messages to the specified topic in a single batch. This is a
// batch version of the Publish API. If you try to send more than 10 messages in a
// single batch request, you will receive a TooManyEntriesInBatchRequest exception.
//
// For FIFO topics, multiple messages within a single batch are published in the
// order they are sent, and messages are deduplicated within the batch and across
// batches for five minutes.
//
// The result of publishing each message is reported individually in the response.
// Because the batch request can result in a combination of successful and
// unsuccessful actions, you should check for batch errors even when the call
// returns an HTTP status code of 200.
//
// The maximum allowed individual message size and the maximum total payload size
// (the sum of the individual lengths of all of the batched messages) are both 256
// KB (262,144 bytes).
//
// The PublishBatch API can send up to 10 messages at a time. If you attempt to
// send more than 10 messages in one request, you will encounter a
// TooManyEntriesInBatchRequest exception. In such cases, split your messages into
// multiple requests, each containing no more than 10 messages.
//
// Some actions take lists of parameters. These lists are specified using the
// param.n notation. Values of n are integers starting from 1. For example, a
// parameter list with two elements looks like this:
//
// &AttributeName.1=first
//
// &AttributeName.2=second
//
// If you send a batch message to a topic, Amazon SNS publishes the batch message
// to each endpoint that is subscribed to the topic. The format of the batch
// message depends on the notification protocol for each subscribed endpoint.
//
// When a messageId is returned, the batch message is saved, and Amazon SNS
// immediately delivers the message to subscribers.
