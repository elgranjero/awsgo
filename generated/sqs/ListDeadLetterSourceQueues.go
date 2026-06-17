package sqs

// ListDeadLetterSourceQueues is generated as a reference stub.
// Executable command wiring lives under cmd/sqs.go.
//
// Returns a list of your queues that have the RedrivePolicy queue attribute
// configured with a dead-letter queue.
//
// The ListDeadLetterSourceQueues methods supports pagination. Set parameter
// MaxResults in the request to specify the maximum number of results to be
// returned in the response. If you do not set MaxResults , the response includes a
// maximum of 1,000 results. If you set MaxResults and there are additional
// results to display, the response includes a value for NextToken . Use NextToken
// as a parameter in your next request to ListDeadLetterSourceQueues to receive
// the next page of results.
//
// For more information about using dead-letter queues, see [Using Amazon SQS Dead-Letter Queues] in the Amazon SQS
// Developer Guide.
//
// [Using Amazon SQS Dead-Letter Queues]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html
