package bedrockagentcorecontrol

// CreatePolicyEngine is generated as a reference stub.
// Executable command wiring lives under cmd/bedrockagentcorecontrol.go.
//
// Creates a new policy engine within the AgentCore Policy system. A policy engine
// is a collection of policies that evaluates and authorizes agent tool calls. When
// associated with Gateways (each Gateway can be associated with at most one policy
// engine, but multiple Gateways can be associated with the same engine), the
// policy engine intercepts all agent requests and determines whether to allow or
// deny each action based on the defined policies. This is an asynchronous
// operation. Use the [GetPolicyEngine]operation to poll the status field to track completion.
//
// [GetPolicyEngine]: https://docs.aws.amazon.com/bedrock-agentcore-control/latest/APIReference/API_GetPolicyEngine.html
