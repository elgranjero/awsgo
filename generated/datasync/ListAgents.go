package datasync

// ListAgents is generated as a reference stub.
// Executable command wiring lives under cmd/datasync.go.
//
// Returns a list of DataSync agents that belong to an Amazon Web Services account
// in the Amazon Web Services Region specified in the request.
//
// With pagination, you can reduce the number of agents returned in a response. If
// you get a truncated list of agents in a response, the response contains a marker
// that you can specify in your next request to fetch the next page of agents.
//
// ListAgents is eventually consistent. This means the result of running the
// operation might not reflect that you just created or deleted an agent. For
// example, if you create an agent with [CreateAgent]and then immediately run ListAgents , that
// agent might not show up in the list right away. In situations like this, you can
// always confirm whether an agent has been created (or deleted) by using [DescribeAgent].
//
// [DescribeAgent]: https://docs.aws.amazon.com/datasync/latest/userguide/API_DescribeAgent.html
// [CreateAgent]: https://docs.aws.amazon.com/datasync/latest/userguide/API_CreateAgent.html
