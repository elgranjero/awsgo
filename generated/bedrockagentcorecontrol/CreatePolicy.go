package bedrockagentcorecontrol

// CreatePolicy is generated as a reference stub.
// Executable command wiring lives under cmd/bedrockagentcorecontrol.go.
//
// Creates a policy within the AgentCore Policy system. Policies provide
// real-time, deterministic control over agentic interactions with AgentCore
// Gateway. Using the Cedar policy language, you can define fine-grained policies
// that specify which interactions with Gateway tools are permitted based on input
// parameters and OAuth claims, ensuring agents operate within defined boundaries
// and business rules. The policy is validated during creation against the Cedar
// schema generated from the Gateway's tools' input schemas, which defines the
// available tools, their parameters, and expected data types. This is an
// asynchronous operation. Use the [GetPolicy]operation to poll the status field to track
// completion.
//
// [GetPolicy]: https://docs.aws.amazon.com/bedrock-agentcore-control/latest/APIReference/API_GetPolicy.html
