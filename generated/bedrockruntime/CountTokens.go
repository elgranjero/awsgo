package bedrockruntime

// CountTokens is generated as a reference stub.
// Executable command wiring lives under cmd/bedrockruntime.go.
//
// Returns the token count for a given inference request. This operation helps you
// estimate token usage before sending requests to foundation models by returning
// the token count that would be used if the same input were sent to the model in
// an inference request.
//
// Token counting is model-specific because different models use different
// tokenization strategies. The token count returned by this operation will match
// the token count that would be charged if the same input were sent to the model
// in an InvokeModel or Converse request.
//
// You can use this operation to:
//
// - Estimate costs before sending inference requests.
//
// - Optimize prompts to fit within token limits.
//
// - Plan for token usage in your applications.
//
// This operation accepts the same input formats as InvokeModel and Converse ,
// allowing you to count tokens for both raw text inputs and structured
// conversation formats.
//
// The following operations are related to CountTokens :
//
// [InvokeModel]
// - - Sends inference requests to foundation models
//
// [Converse]
// - - Sends conversation-based inference requests to foundation models
//
// [Converse]: https://docs.aws.amazon.com/bedrock/latest/API/API_runtime_Converse.html
// [InvokeModel]: https://docs.aws.amazon.com/bedrock/latest/API/API_runtime_InvokeModel.html
