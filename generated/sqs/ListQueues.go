package sqs

// ListQueues is generated as a reference stub.
// Executable command wiring lives under cmd/sqs.go.
//
// Returns a list of your queues in the current region. The response includes a
// maximum of 1,000 results. If you specify a value for the optional
// QueueNamePrefix parameter, only queues with a name that begins with the
// specified value are returned.
//
// The listQueues methods supports pagination. Set parameter MaxResults in the
// request to specify the maximum number of results to be returned in the response.
// If you do not set MaxResults , the response includes a maximum of 1,000 results.
// If you set MaxResults and there are additional results to display, the response
// includes a value for NextToken . Use NextToken as a parameter in your next
// request to listQueues to receive the next page of results.
//
// Cross-account permissions don't apply to this action. For more information, see [Grant cross-account permissions to a role and a username]
// in the Amazon SQS Developer Guide.
//
// [Grant cross-account permissions to a role and a username]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name
