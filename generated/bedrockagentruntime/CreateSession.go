package bedrockagentruntime

// CreateSession is generated as a reference stub.
// Executable command wiring lives under cmd/bedrockagentruntime.go.
//
// Creates a session to temporarily store conversations for generative AI (GenAI)
// applications built with open-source frameworks such as LangGraph and LlamaIndex.
// Sessions enable you to save the state of conversations at checkpoints, with the
// added security and infrastructure of Amazon Web Services. For more information,
// see [Store and retrieve conversation history and context with Amazon Bedrock sessions].
//
// By default, Amazon Bedrock uses Amazon Web Services-managed keys for session
// encryption, including session metadata, or you can use your own KMS key. For
// more information, see [Amazon Bedrock session encryption].
//
// You use a session to store state and conversation history for generative AI
// applications built with open-source frameworks. For Amazon Bedrock Agents, the
// service automatically manages conversation context and associates them with the
// agent-specific sessionId you specify in the [InvokeAgent]API operation.
//
// Related APIs:
//
// [ListSessions]
//
// [GetSession]
//
// [EndSession]
//
// [DeleteSession]
//
// [Store and retrieve conversation history and context with Amazon Bedrock sessions]: https://docs.aws.amazon.com/bedrock/latest/userguide/sessions.html
// [DeleteSession]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_DeleteSession.html
// [Amazon Bedrock session encryption]: https://docs.aws.amazon.com/bedrock/latest/userguide/session-encryption.html
// [InvokeAgent]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_InvokeAgent.html
// [EndSession]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_EndSession.html
// [ListSessions]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_ListSessions.html
// [GetSession]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_GetSession.html
