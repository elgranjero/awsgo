package bedrockagentruntime

// InvokeAgent is generated as a reference stub.
// Executable command wiring lives under cmd/bedrockagentruntime.go.
//
// Sends a prompt for the agent to process and respond to. Note the following
// fields for the request:
//
// - To continue the same conversation with an agent, use the same sessionId
// value in the request.
//
// - To activate trace enablement, turn enableTrace to true . Trace enablement
// helps you follow the agent's reasoning process that led it to the information it
// processed, the actions it took, and the final result it yielded. For more
// information, see [Trace enablement].
//
// - End a conversation by setting endSession to true .
//
// - In the sessionState object, you can include attributes for the session or
// prompt or, if you configured an action group to return control, results from
// invocation of the action group.
//
// The response contains both chunk and trace attributes.
//
// The final response is returned in the bytes field of the chunk object. The
// InvokeAgent returns one chunk for the entire interaction.
//
// - The attribution object contains citations for parts of the response.
//
// - If you set enableTrace to true in the request, you can trace the agent's
// steps and reasoning process that led it to the response.
//
// - If the action predicted was configured to return control, the response
// returns parameters for the action, elicited from the user, in the
// returnControl field.
//
// - Errors are also surfaced in the response.
//
// [Trace enablement]: https://docs.aws.amazon.com/bedrock/latest/userguide/agents-test.html#trace-events
