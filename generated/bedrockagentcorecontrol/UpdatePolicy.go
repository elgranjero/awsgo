package bedrockagentcorecontrol

// UpdatePolicy is generated as a reference stub.
// Executable command wiring lives under cmd/bedrockagentcorecontrol.go.
//
// Updates an existing policy within the AgentCore Policy system. This operation
// allows modification of the policy description and definition while maintaining
// the policy's identity. The updated policy is validated against the Cedar schema
// before being applied. This is an asynchronous operation. Use the GetPolicy
// operation to poll the status field to track completion.
