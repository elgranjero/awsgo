package bedrockagentcore

// StartMemoryExtractionJob is generated as a reference stub.
// Executable command wiring lives under cmd/bedrockagentcore.go.
//
// Starts a memory extraction job that processes events that failed extraction
//
// previously in an AgentCore Memory resource and produces structured memory
// records. When earlier extraction attempts have left events unprocessed, this job
// will pick up and extract those as well.
//
// To use this operation, you must have the
// bedrock-agentcore:StartMemoryExtractionJob permission.
