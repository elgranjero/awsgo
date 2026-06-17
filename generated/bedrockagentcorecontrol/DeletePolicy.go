package bedrockagentcorecontrol

// DeletePolicy is generated as a reference stub.
// Executable command wiring lives under cmd/bedrockagentcorecontrol.go.
//
// Deletes an existing policy from the AgentCore Policy system. Once deleted, the
// policy can no longer be used for agent behavior control and all references to it
// become invalid. This is an asynchronous operation. Use the GetPolicy operation
// to poll the status field to track completion.
