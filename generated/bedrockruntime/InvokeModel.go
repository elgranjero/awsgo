package bedrockruntime

// InvokeModel is generated as a reference stub.
// Executable command wiring lives under cmd/bedrockruntime.go.
//
// Invokes the specified Amazon Bedrock model to run inference using the prompt
// and inference parameters provided in the request body. You use model inference
// to generate text, images, and embeddings.
//
// For example code, see Invoke model code examples in the Amazon Bedrock User
// Guide.
//
// This operation requires permission for the bedrock:InvokeModel action.
//
// To deny all inference access to resources that you specify in the modelId
// field, you need to deny access to the bedrock:InvokeModel and
// bedrock:InvokeModelWithResponseStream actions. Doing this also denies access to
// the resource through the Converse API actions ([Converse] and [ConverseStream]). For more information see [Deny access for inference on specific models]
// .
//
// For troubleshooting some of the common errors you might encounter when using
// the InvokeModel API, see [Troubleshooting Amazon Bedrock API Error Codes] in the Amazon Bedrock User Guide
//
// [Converse]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_runtime_Converse.html
// [ConverseStream]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_runtime_ConverseStream.html
// [Troubleshooting Amazon Bedrock API Error Codes]: https://docs.aws.amazon.com/bedrock/latest/userguide/troubleshooting-api-error-codes.html
// [Deny access for inference on specific models]: https://docs.aws.amazon.com/bedrock/latest/userguide/security_iam_id-based-policy-examples.html#security_iam_id-based-policy-examples-deny-inference
