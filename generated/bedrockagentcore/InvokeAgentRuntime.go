package bedrockagentcore

// InvokeAgentRuntime is generated as a reference stub.
// Executable command wiring lives under cmd/bedrockagentcore.go.
//
// Sends a request to an agent or tool hosted in an Amazon Bedrock AgentCore
// Runtime and receives responses in real-time.
//
// To invoke an agent you must specify the AgentCore Runtime ARN and provide a
// payload containing your request. You can optionally specify a qualifier to
// target a specific version or endpoint of the agent.
//
// This operation supports streaming responses, allowing you to receive partial
// responses as they become available. We recommend using pagination to ensure that
// the operation returns quickly and successfully when processing large responses.
//
// For example code, see [Invoke an AgentCore Runtime agent].
//
// If you're integrating your agent with OAuth, you can't use the Amazon Web
// Services SDK to call InvokeAgentRuntime . Instead, make a HTTPS request to
// InvokeAgentRuntime . For an example, see [Authenticate and authorize with Inbound Auth and Outbound Auth].
//
// To use this operation, you must have the bedrock-agentcore:InvokeAgentRuntime
// permission. If you are making a call to InvokeAgentRuntime on behalf of a user
// ID with the X-Amzn-Bedrock-AgentCore-Runtime-User-Id header, You require
// permissions to both actions ( bedrock-agentcore:InvokeAgentRuntime and
// bedrock-agentcore:InvokeAgentRuntimeForUser ).
//
// [Invoke an AgentCore Runtime agent]: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/runtime-invoke-agent.html
// [Authenticate and authorize with Inbound Auth and Outbound Auth]: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/runtime-oauth.html
