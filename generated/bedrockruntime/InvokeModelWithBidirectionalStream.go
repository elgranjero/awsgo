package bedrockruntime

// InvokeModelWithBidirectionalStream is generated as a reference stub.
// Executable command wiring lives under cmd/bedrockruntime.go.
//
// Invoke the specified Amazon Bedrock model to run inference using the
// bidirectional stream. The response is returned in a stream that remains open for
// 8 minutes. A single session can contain multiple prompts and responses from the
// model. The prompts to the model are provided as audio files and the model's
// responses are spoken back to the user and transcribed.
//
// It is possible for users to interrupt the model's response with a new prompt,
// which will halt the response speech. The model will retain contextual awareness
// of the conversation while pivoting to respond to the new prompt.
