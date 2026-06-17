package bedrockagentcore

// InvokeCodeInterpreter is generated as a reference stub.
// Executable command wiring lives under cmd/bedrockagentcore.go.
//
// Executes code within an active code interpreter session in Amazon Bedrock
// AgentCore. This operation processes the provided code, runs it in a secure
// environment, and returns the execution results including output, errors, and
// generated visualizations.
//
// To execute code, you must specify the code interpreter identifier, session ID,
// and the code to run in the arguments parameter. The operation returns a stream
// containing the execution results, which can include text output, error messages,
// and data visualizations.
//
// This operation is subject to request rate limiting based on your account's
// service quotas.
//
// The following operations are related to InvokeCodeInterpreter :
//
// [StartCodeInterpreterSession]
//
// [GetCodeInterpreterSession]
//
// [StartCodeInterpreterSession]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_StartCodeInterpreterSession.html
// [GetCodeInterpreterSession]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_GetCodeInterpreterSession.html
