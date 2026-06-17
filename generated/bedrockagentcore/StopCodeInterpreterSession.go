package bedrockagentcore

// StopCodeInterpreterSession is generated as a reference stub.
// Executable command wiring lives under cmd/bedrockagentcore.go.
//
// Terminates an active code interpreter session in Amazon Bedrock AgentCore. This
// operation stops the session, releases associated resources, and makes the
// session unavailable for further use.
//
// To stop a code interpreter session, you must specify both the code interpreter
// identifier and the session ID. Once stopped, a session cannot be restarted; you
// must create a new session using StartCodeInterpreterSession .
//
// The following operations are related to StopCodeInterpreterSession :
//
// [StartCodeInterpreterSession]
//
// [GetCodeInterpreterSession]
//
// [StartCodeInterpreterSession]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_StartCodeInterpreterSession.html
// [GetCodeInterpreterSession]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_GetCodeInterpreterSession.html
