package connect

// ListQueues is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Provides information about the queues for the specified Amazon Connect instance.
//
// If you do not specify a QueueTypes parameter, both standard and agent queues
// are returned. This might cause an unexpected truncation of results if you have
// more than 1000 agents and you limit the number of results of the API call in
// code.
//
// For more information about queues, see [Queues: Standard and Agent] in the Amazon Connect Administrator
// Guide.
//
// [Queues: Standard and Agent]: https://docs.aws.amazon.com/connect/latest/adminguide/concepts-queues-standard-and-agent.html
