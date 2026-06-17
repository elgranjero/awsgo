package bedrockruntime

// InvokeModelWithResponseStream is generated as a reference stub.
// Executable command wiring lives under cmd/bedrockruntime.go.
//
// Invoke the specified Amazon Bedrock model to run inference using the prompt and
// inference parameters provided in the request body. The response is returned in a
// stream.
//
// To see if a model supports streaming, call [GetFoundationModel] and check the
// responseStreamingSupported field in the response.
//
// The CLI doesn't support streaming operations in Amazon Bedrock, including
// InvokeModelWithResponseStream .
//
// For example code, see Invoke model with streaming code example in the Amazon
// Bedrock User Guide.
//
// This operation requires permissions to perform the
// bedrock:InvokeModelWithResponseStream action.
//
// To deny all inference access to resources that you specify in the modelId
// field, you need to deny access to the bedrock:InvokeModel and
// bedrock:InvokeModelWithResponseStream actions. Doing this also denies access to
// the resource through the Converse API actions ([Converse] and [ConverseStream]). For more information see [Deny access for inference on specific models]
// .
//
// For troubleshooting some of the common errors you might encounter when using
// the InvokeModelWithResponseStream API, see [Troubleshooting Amazon Bedrock API Error Codes] in the Amazon Bedrock User Guide
//
// [GetFoundationModel]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_GetFoundationModel.html
// [Converse]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_runtime_Converse.html
// [ConverseStream]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_runtime_ConverseStream.html
// [Troubleshooting Amazon Bedrock API Error Codes]: https://docs.aws.amazon.com/bedrock/latest/userguide/troubleshooting-api-error-codes.html
// [Deny access for inference on specific models]: https://docs.aws.amazon.com/bedrock/latest/userguide/security_iam_id-based-policy-examples.html#security_iam_id-based-policy-examples-deny-inference
