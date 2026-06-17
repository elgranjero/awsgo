package bedrockagent

// UpdateKnowledgeBase is generated as a reference stub.
// Executable command wiring lives under cmd/bedrockagent.go.
//
// Updates the configuration of a knowledge base with the fields that you specify.
// Because all fields will be overwritten, you must include the same values for
// fields that you want to keep the same.
//
// You can change the following fields:
//
// - name
//
// - description
//
// - roleArn
//
// You can't change the knowledgeBaseConfiguration or storageConfiguration fields,
// so you must specify the same configurations as when you created the knowledge
// base. You can send a [GetKnowledgeBase]request and copy the same configurations.
//
// [GetKnowledgeBase]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent_GetKnowledgeBase.html
