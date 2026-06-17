package bedrockagentcorecontrol

// DeletePolicyEngine is generated as a reference stub.
// Executable command wiring lives under cmd/bedrockagentcorecontrol.go.
//
// Deletes an existing policy engine from the AgentCore Policy system. The policy
// engine must not have any associated policies before deletion. Once deleted, the
// policy engine and all its configurations become unavailable for policy
// management and evaluation. This is an asynchronous operation. Use the
// GetPolicyEngine operation to poll the status field to track completion.
