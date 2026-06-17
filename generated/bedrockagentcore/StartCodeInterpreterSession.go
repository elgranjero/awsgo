package bedrockagentcore

// StartCodeInterpreterSession is generated as a reference stub.
// Executable command wiring lives under cmd/bedrockagentcore.go.
//
// Creates and initializes a code interpreter session in Amazon Bedrock AgentCore.
// The session enables agents to execute code as part of their response generation,
// supporting programming languages such as Python for data analysis,
// visualization, and computation tasks.
//
// To create a session, you must specify a code interpreter identifier and a name.
// The session remains active until it times out or you explicitly stop it using
// the StopCodeInterpreterSession operation.
//
// The following operations are related to StartCodeInterpreterSession :
//
// [InvokeCodeInterpreter]
//
// [GetCodeInterpreterSession]
//
// [StopCodeInterpreterSession]
//
// [InvokeCodeInterpreter]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_InvokeCodeInterpreter.html
// [StopCodeInterpreterSession]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_StopCodeInterpreterSession.html
// [GetCodeInterpreterSession]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_GetCodeInterpreterSession.html
