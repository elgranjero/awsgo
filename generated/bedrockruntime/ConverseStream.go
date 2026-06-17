package bedrockruntime

// ConverseStream is generated as a reference stub.
// Executable command wiring lives under cmd/bedrockruntime.go.
//
// Sends messages to the specified Amazon Bedrock model and returns the response
// in a stream. ConverseStream provides a consistent API that works with all
// Amazon Bedrock models that support messages. This allows you to write code once
// and use it with different models. Should a model have unique inference
// parameters, you can also pass those unique parameters to the model.
//
// To find out if a model supports streaming, call [GetFoundationModel] and check the
// responseStreamingSupported field in the response.
//
// The CLI doesn't support streaming operations in Amazon Bedrock, including
// ConverseStream .
//
// Amazon Bedrock doesn't store any text, images, or documents that you provide as
// content. The data is only used to generate the response.
//
// You can submit a prompt by including it in the messages field, specifying the
// modelId of a foundation model or inference profile to run inference on it, and
// including any other fields that are relevant to your use case.
//
// You can also submit a prompt from Prompt management by specifying the ARN of
// the prompt version and including a map of variables to values in the
// promptVariables field. You can append more messages to the prompt by using the
// messages field. If you use a prompt from Prompt management, you can't include
// the following fields in the request: additionalModelRequestFields ,
// inferenceConfig , system , or toolConfig . Instead, these fields must be defined
// through Prompt management. For more information, see [Use a prompt from Prompt management].
//
// For information about the Converse API, see Use the Converse API in the Amazon
// Bedrock User Guide. To use a guardrail, see Use a guardrail with the Converse
// API in the Amazon Bedrock User Guide. To use a tool with a model, see Tool use
// (Function calling) in the Amazon Bedrock User Guide
//
// For example code, see Conversation streaming example in the Amazon Bedrock User
// Guide.
//
// This operation requires permission for the bedrock:InvokeModelWithResponseStream
// action.
//
// To deny all inference access to resources that you specify in the modelId
// field, you need to deny access to the bedrock:InvokeModel and
// bedrock:InvokeModelWithResponseStream actions. Doing this also denies access to
// the resource through the base inference actions ([InvokeModel] and [InvokeModelWithResponseStream]). For more information
// see [Deny access for inference on specific models].
//
// For troubleshooting some of the common errors you might encounter when using
// the ConverseStream API, see [Troubleshooting Amazon Bedrock API Error Codes] in the Amazon Bedrock User Guide
//
// [InvokeModelWithResponseStream]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_runtime_InvokeModelWithResponseStream.html
// [GetFoundationModel]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_GetFoundationModel.html
// [Troubleshooting Amazon Bedrock API Error Codes]: https://docs.aws.amazon.com/bedrock/latest/userguide/troubleshooting-api-error-codes.html
// [Deny access for inference on specific models]: https://docs.aws.amazon.com/bedrock/latest/userguide/security_iam_id-based-policy-examples.html#security_iam_id-based-policy-examples-deny-inference
// [InvokeModel]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_runtime_InvokeModel.html
// [Use a prompt from Prompt management]: https://docs.aws.amazon.com/bedrock/latest/userguide/prompt-management-use.html
