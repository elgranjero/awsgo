package bedrockruntime

// StartAsyncInvoke is generated as a reference stub.
// Executable command wiring lives under cmd/bedrockruntime.go.
//
// Starts an asynchronous invocation.
//
// This operation requires permission for the bedrock:InvokeModel action.
//
// To deny all inference access to resources that you specify in the modelId
// field, you need to deny access to the bedrock:InvokeModel and
// bedrock:InvokeModelWithResponseStream actions. Doing this also denies access to
// the resource through the Converse API actions ([Converse] and [ConverseStream]). For more information see [Deny access for inference on specific models]
// .
//
// [Converse]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_runtime_Converse.html
// [ConverseStream]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_runtime_ConverseStream.html
// [Deny access for inference on specific models]: https://docs.aws.amazon.com/bedrock/latest/userguide/security_iam_id-based-policy-examples.html#security_iam_id-based-policy-examples-deny-inference
