package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// bedrockruntimeCmd represents the bedrockruntime command
var _bedrockruntimeCmd = &cobra.Command{
	Use:   "bedrockruntime",
	Short: "AWS bedrockruntime CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := bedrockruntime.NewFromConfig(cfg)
		if _bedrockruntimeApplyGuardrail {
			bedrockruntime_ApplyGuardrail(cfg, client)
			return
		}
		if _bedrockruntimeConverse {
			bedrockruntime_Converse(cfg, client)
			return
		}
		if _bedrockruntimeConverseStream {
			bedrockruntime_ConverseStream(cfg, client)
			return
		}
		if _bedrockruntimeCountTokens {
			bedrockruntime_CountTokens(cfg, client)
			return
		}
		if _bedrockruntimeGetAsyncInvoke {
			bedrockruntime_GetAsyncInvoke(cfg, client)
			return
		}
		if _bedrockruntimeInvokeModel {
			bedrockruntime_InvokeModel(cfg, client)
			return
		}
		if _bedrockruntimeInvokeModelWithBidirectionalStream {
			bedrockruntime_InvokeModelWithBidirectionalStream(cfg, client)
			return
		}
		if _bedrockruntimeInvokeModelWithResponseStream {
			bedrockruntime_InvokeModelWithResponseStream(cfg, client)
			return
		}
		if _bedrockruntimeListAsyncInvokes {
			bedrockruntime_ListAsyncInvokes(cfg, client)
			return
		}
		if _bedrockruntimeStartAsyncInvoke {
			bedrockruntime_StartAsyncInvoke(cfg, client)
			return
		}

	},
}

var (
	_bedrockruntimeApplyGuardrail                     bool
	_bedrockruntimeConverse                           bool
	_bedrockruntimeConverseStream                     bool
	_bedrockruntimeCountTokens                        bool
	_bedrockruntimeGetAsyncInvoke                     bool
	_bedrockruntimeInvokeModel                        bool
	_bedrockruntimeInvokeModelWithBidirectionalStream bool
	_bedrockruntimeInvokeModelWithResponseStream      bool
	_bedrockruntimeListAsyncInvokes                   bool
	_bedrockruntimeStartAsyncInvoke                   bool

	_bedrockruntimeAccept                            string
	_bedrockruntimeAdditionalModelRequestFields      string
	_bedrockruntimeAdditionalModelResponseFieldPaths []string
	_bedrockruntimeBody                              string
	_bedrockruntimeClientRequestToken                string
	_bedrockruntimeContent                           string
	_bedrockruntimeContentType                       string
	_bedrockruntimeGuardrailConfig                   string
	_bedrockruntimeGuardrailIdentifier               string
	_bedrockruntimeGuardrailVersion                  string
	_bedrockruntimeInferenceConfig                   string
	_bedrockruntimeInput                             string
	_bedrockruntimeInvocationArn                     string
	_bedrockruntimeMaxResults                        string
	_bedrockruntimeMessages                          string
	_bedrockruntimeModelId                           string
	_bedrockruntimeModelInput                        string
	_bedrockruntimeNextToken                         string
	_bedrockruntimeOutputConfig                      string
	_bedrockruntimeOutputDataConfig                  string
	_bedrockruntimeOutputScope                       string
	_bedrockruntimePerformanceConfig                 string
	_bedrockruntimePerformanceConfigLatency          string
	_bedrockruntimePromptVariables                   string
	_bedrockruntimeRequestMetadata                   string
	_bedrockruntimeServiceTier                       string
	_bedrockruntimeSortBy                            string
	_bedrockruntimeSortOrder                         string
	_bedrockruntimeSource                            string
	_bedrockruntimeStatusEquals                      string
	_bedrockruntimeSubmitTimeAfter                   string
	_bedrockruntimeSubmitTimeBefore                  string
	_bedrockruntimeSystem                            string
	_bedrockruntimeTags                              string
	_bedrockruntimeToolConfig                        string
	_bedrockruntimeTrace                             string
)

// The action to apply a guardrail.
// For troubleshooting some of the common errors you might encounter when using
// the ApplyGuardrail API, see [Troubleshooting Amazon Bedrock API Error Codes] in the Amazon Bedrock User Guide
//
// [Troubleshooting Amazon Bedrock API Error Codes]: https://docs.aws.amazon.com/bedrock/latest/userguide/troubleshooting-api-error-codes.html
func bedrockruntime_ApplyGuardrail(cfg aws.Config, client *bedrockruntime.Client) {
	input := &bedrockruntime.ApplyGuardrailInput{
		// Content: []types.GuardrailContentBlock, // Required
		// GuardrailIdentifier: *string, // Required
		// GuardrailVersion: *string, // Required
		// Source: types.GuardrailContentSource, // Required
	}

	if len(_bedrockruntimeContent) > 0 {
		if err := assignInputField(input, "Content", _bedrockruntimeContent); err != nil {
			log.Errorf("invalid --content: %s", err.Error())
			return
		}
	}
	if len(_bedrockruntimeGuardrailIdentifier) > 0 {
		input.GuardrailIdentifier = aws.String(_bedrockruntimeGuardrailIdentifier)
	}
	if len(_bedrockruntimeGuardrailVersion) > 0 {
		input.GuardrailVersion = aws.String(_bedrockruntimeGuardrailVersion)
	}
	if len(_bedrockruntimeSource) > 0 {
		if err := assignInputField(input, "Source", _bedrockruntimeSource); err != nil {
			log.Errorf("invalid --source: %s", err.Error())
			return
		}
	}
	if len(_bedrockruntimeOutputScope) > 0 {
		if err := assignInputField(input, "OutputScope", _bedrockruntimeOutputScope); err != nil {
			log.Errorf("invalid --output-scope: %s", err.Error())
			return
		}
	}

	if resp, err := client.ApplyGuardrail(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends messages to the specified Amazon Bedrock model. Converse provides a
// consistent interface that works with all models that support messages. This
// allows you to write code once and use it with different models. If a model has
// unique inference parameters, you can also pass those unique parameters to the
// model.
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
// For example code, see Converse API examples in the Amazon Bedrock User Guide.
//
// This operation requires permission for the bedrock:InvokeModel action.
//
// To deny all inference access to resources that you specify in the modelId
// field, you need to deny access to the bedrock:InvokeModel and
// bedrock:InvokeModelWithResponseStream actions. Doing this also denies access to
// the resource through the base inference actions ([InvokeModel] and [InvokeModelWithResponseStream]). For more information
// see [Deny access for inference on specific models].
//
// For troubleshooting some of the common errors you might encounter when using
// the Converse API, see [Troubleshooting Amazon Bedrock API Error Codes] in the Amazon Bedrock User Guide
//
// [InvokeModelWithResponseStream]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_runtime_InvokeModelWithResponseStream.html
// [Troubleshooting Amazon Bedrock API Error Codes]: https://docs.aws.amazon.com/bedrock/latest/userguide/troubleshooting-api-error-codes.html
// [Deny access for inference on specific models]: https://docs.aws.amazon.com/bedrock/latest/userguide/security_iam_id-based-policy-examples.html#security_iam_id-based-policy-examples-deny-inference
// [InvokeModel]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_runtime_InvokeModel.html
// [Use a prompt from Prompt management]: https://docs.aws.amazon.com/bedrock/latest/userguide/prompt-management-use.html
func bedrockruntime_Converse(cfg aws.Config, client *bedrockruntime.Client) {
	input := &bedrockruntime.ConverseInput{
		// ModelId: *string, // Required
	}

	if len(_bedrockruntimeModelId) > 0 {
		input.ModelId = aws.String(_bedrockruntimeModelId)
	}
	if len(_bedrockruntimeAdditionalModelRequestFields) > 0 {
		if err := assignInputField(input, "AdditionalModelRequestFields", _bedrockruntimeAdditionalModelRequestFields); err != nil {
			log.Errorf("invalid --additional-model-request-fields: %s", err.Error())
			return
		}
	}
	if len(_bedrockruntimeAdditionalModelResponseFieldPaths) > 0 {
		input.AdditionalModelResponseFieldPaths = append([]string(nil), _bedrockruntimeAdditionalModelResponseFieldPaths...)
	}
	if len(_bedrockruntimeGuardrailConfig) > 0 {
		if err := assignInputField(input, "GuardrailConfig", _bedrockruntimeGuardrailConfig); err != nil {
			log.Errorf("invalid --guardrail-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockruntimeInferenceConfig) > 0 {
		if err := assignInputField(input, "InferenceConfig", _bedrockruntimeInferenceConfig); err != nil {
			log.Errorf("invalid --inference-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockruntimeMessages) > 0 {
		if err := assignInputField(input, "Messages", _bedrockruntimeMessages); err != nil {
			log.Errorf("invalid --messages: %s", err.Error())
			return
		}
	}
	if len(_bedrockruntimeOutputConfig) > 0 {
		if err := assignInputField(input, "OutputConfig", _bedrockruntimeOutputConfig); err != nil {
			log.Errorf("invalid --output-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockruntimePerformanceConfig) > 0 {
		if err := assignInputField(input, "PerformanceConfig", _bedrockruntimePerformanceConfig); err != nil {
			log.Errorf("invalid --performance-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockruntimePromptVariables) > 0 {
		if err := assignInputField(input, "PromptVariables", _bedrockruntimePromptVariables); err != nil {
			log.Errorf("invalid --prompt-variables: %s", err.Error())
			return
		}
	}
	if len(_bedrockruntimeRequestMetadata) > 0 {
		if err := assignInputField(input, "RequestMetadata", _bedrockruntimeRequestMetadata); err != nil {
			log.Errorf("invalid --request-metadata: %s", err.Error())
			return
		}
	}
	if len(_bedrockruntimeServiceTier) > 0 {
		if err := assignInputField(input, "ServiceTier", _bedrockruntimeServiceTier); err != nil {
			log.Errorf("invalid --service-tier: %s", err.Error())
			return
		}
	}
	if len(_bedrockruntimeSystem) > 0 {
		if err := assignInputField(input, "System", _bedrockruntimeSystem); err != nil {
			log.Errorf("invalid --system: %s", err.Error())
			return
		}
	}
	if len(_bedrockruntimeToolConfig) > 0 {
		if err := assignInputField(input, "ToolConfig", _bedrockruntimeToolConfig); err != nil {
			log.Errorf("invalid --tool-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.Converse(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

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
func bedrockruntime_ConverseStream(cfg aws.Config, client *bedrockruntime.Client) {
	input := &bedrockruntime.ConverseStreamInput{
		// ModelId: *string, // Required
	}

	if len(_bedrockruntimeModelId) > 0 {
		input.ModelId = aws.String(_bedrockruntimeModelId)
	}
	if len(_bedrockruntimeAdditionalModelRequestFields) > 0 {
		if err := assignInputField(input, "AdditionalModelRequestFields", _bedrockruntimeAdditionalModelRequestFields); err != nil {
			log.Errorf("invalid --additional-model-request-fields: %s", err.Error())
			return
		}
	}
	if len(_bedrockruntimeAdditionalModelResponseFieldPaths) > 0 {
		input.AdditionalModelResponseFieldPaths = append([]string(nil), _bedrockruntimeAdditionalModelResponseFieldPaths...)
	}
	if len(_bedrockruntimeGuardrailConfig) > 0 {
		if err := assignInputField(input, "GuardrailConfig", _bedrockruntimeGuardrailConfig); err != nil {
			log.Errorf("invalid --guardrail-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockruntimeInferenceConfig) > 0 {
		if err := assignInputField(input, "InferenceConfig", _bedrockruntimeInferenceConfig); err != nil {
			log.Errorf("invalid --inference-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockruntimeMessages) > 0 {
		if err := assignInputField(input, "Messages", _bedrockruntimeMessages); err != nil {
			log.Errorf("invalid --messages: %s", err.Error())
			return
		}
	}
	if len(_bedrockruntimeOutputConfig) > 0 {
		if err := assignInputField(input, "OutputConfig", _bedrockruntimeOutputConfig); err != nil {
			log.Errorf("invalid --output-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockruntimePerformanceConfig) > 0 {
		if err := assignInputField(input, "PerformanceConfig", _bedrockruntimePerformanceConfig); err != nil {
			log.Errorf("invalid --performance-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockruntimePromptVariables) > 0 {
		if err := assignInputField(input, "PromptVariables", _bedrockruntimePromptVariables); err != nil {
			log.Errorf("invalid --prompt-variables: %s", err.Error())
			return
		}
	}
	if len(_bedrockruntimeRequestMetadata) > 0 {
		if err := assignInputField(input, "RequestMetadata", _bedrockruntimeRequestMetadata); err != nil {
			log.Errorf("invalid --request-metadata: %s", err.Error())
			return
		}
	}
	if len(_bedrockruntimeServiceTier) > 0 {
		if err := assignInputField(input, "ServiceTier", _bedrockruntimeServiceTier); err != nil {
			log.Errorf("invalid --service-tier: %s", err.Error())
			return
		}
	}
	if len(_bedrockruntimeSystem) > 0 {
		if err := assignInputField(input, "System", _bedrockruntimeSystem); err != nil {
			log.Errorf("invalid --system: %s", err.Error())
			return
		}
	}
	if len(_bedrockruntimeToolConfig) > 0 {
		if err := assignInputField(input, "ToolConfig", _bedrockruntimeToolConfig); err != nil {
			log.Errorf("invalid --tool-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.ConverseStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

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
func bedrockruntime_CountTokens(cfg aws.Config, client *bedrockruntime.Client) {
	input := &bedrockruntime.CountTokensInput{
		// Input: types.CountTokensInput, // Required
		// ModelId: *string, // Required
	}

	if len(_bedrockruntimeInput) > 0 {
		if err := assignInputField(input, "Input", _bedrockruntimeInput); err != nil {
			log.Errorf("invalid --input: %s", err.Error())
			return
		}
	}
	if len(_bedrockruntimeModelId) > 0 {
		input.ModelId = aws.String(_bedrockruntimeModelId)
	}

	if resp, err := client.CountTokens(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve information about an asynchronous invocation.
func bedrockruntime_GetAsyncInvoke(cfg aws.Config, client *bedrockruntime.Client) {
	input := &bedrockruntime.GetAsyncInvokeInput{
		// InvocationArn: *string, // Required
	}

	if len(_bedrockruntimeInvocationArn) > 0 {
		input.InvocationArn = aws.String(_bedrockruntimeInvocationArn)
	}

	if resp, err := client.GetAsyncInvoke(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

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
func bedrockruntime_InvokeModel(cfg aws.Config, client *bedrockruntime.Client) {
	input := &bedrockruntime.InvokeModelInput{
		// ModelId: *string, // Required
	}

	if len(_bedrockruntimeModelId) > 0 {
		input.ModelId = aws.String(_bedrockruntimeModelId)
	}
	if len(_bedrockruntimeAccept) > 0 {
		input.Accept = aws.String(_bedrockruntimeAccept)
	}
	if len(_bedrockruntimeBody) > 0 {
		if err := assignInputField(input, "Body", _bedrockruntimeBody); err != nil {
			log.Errorf("invalid --body: %s", err.Error())
			return
		}
	}
	if len(_bedrockruntimeContentType) > 0 {
		input.ContentType = aws.String(_bedrockruntimeContentType)
	}
	if len(_bedrockruntimeGuardrailIdentifier) > 0 {
		input.GuardrailIdentifier = aws.String(_bedrockruntimeGuardrailIdentifier)
	}
	if len(_bedrockruntimeGuardrailVersion) > 0 {
		input.GuardrailVersion = aws.String(_bedrockruntimeGuardrailVersion)
	}
	if len(_bedrockruntimePerformanceConfigLatency) > 0 {
		if err := assignInputField(input, "PerformanceConfigLatency", _bedrockruntimePerformanceConfigLatency); err != nil {
			log.Errorf("invalid --performance-config-latency: %s", err.Error())
			return
		}
	}
	if len(_bedrockruntimeServiceTier) > 0 {
		if err := assignInputField(input, "ServiceTier", _bedrockruntimeServiceTier); err != nil {
			log.Errorf("invalid --service-tier: %s", err.Error())
			return
		}
	}
	if len(_bedrockruntimeTrace) > 0 {
		if err := assignInputField(input, "Trace", _bedrockruntimeTrace); err != nil {
			log.Errorf("invalid --trace: %s", err.Error())
			return
		}
	}

	if resp, err := client.InvokeModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Invoke the specified Amazon Bedrock model to run inference using the
// bidirectional stream. The response is returned in a stream that remains open for
// 8 minutes. A single session can contain multiple prompts and responses from the
// model. The prompts to the model are provided as audio files and the model's
// responses are spoken back to the user and transcribed.
//
// It is possible for users to interrupt the model's response with a new prompt,
// which will halt the response speech. The model will retain contextual awareness
// of the conversation while pivoting to respond to the new prompt.
func bedrockruntime_InvokeModelWithBidirectionalStream(cfg aws.Config, client *bedrockruntime.Client) {
	input := &bedrockruntime.InvokeModelWithBidirectionalStreamInput{
		// ModelId: *string, // Required
	}

	if len(_bedrockruntimeModelId) > 0 {
		input.ModelId = aws.String(_bedrockruntimeModelId)
	}

	if resp, err := client.InvokeModelWithBidirectionalStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

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
func bedrockruntime_InvokeModelWithResponseStream(cfg aws.Config, client *bedrockruntime.Client) {
	input := &bedrockruntime.InvokeModelWithResponseStreamInput{
		// ModelId: *string, // Required
	}

	if len(_bedrockruntimeModelId) > 0 {
		input.ModelId = aws.String(_bedrockruntimeModelId)
	}
	if len(_bedrockruntimeAccept) > 0 {
		input.Accept = aws.String(_bedrockruntimeAccept)
	}
	if len(_bedrockruntimeBody) > 0 {
		if err := assignInputField(input, "Body", _bedrockruntimeBody); err != nil {
			log.Errorf("invalid --body: %s", err.Error())
			return
		}
	}
	if len(_bedrockruntimeContentType) > 0 {
		input.ContentType = aws.String(_bedrockruntimeContentType)
	}
	if len(_bedrockruntimeGuardrailIdentifier) > 0 {
		input.GuardrailIdentifier = aws.String(_bedrockruntimeGuardrailIdentifier)
	}
	if len(_bedrockruntimeGuardrailVersion) > 0 {
		input.GuardrailVersion = aws.String(_bedrockruntimeGuardrailVersion)
	}
	if len(_bedrockruntimePerformanceConfigLatency) > 0 {
		if err := assignInputField(input, "PerformanceConfigLatency", _bedrockruntimePerformanceConfigLatency); err != nil {
			log.Errorf("invalid --performance-config-latency: %s", err.Error())
			return
		}
	}
	if len(_bedrockruntimeServiceTier) > 0 {
		if err := assignInputField(input, "ServiceTier", _bedrockruntimeServiceTier); err != nil {
			log.Errorf("invalid --service-tier: %s", err.Error())
			return
		}
	}
	if len(_bedrockruntimeTrace) > 0 {
		if err := assignInputField(input, "Trace", _bedrockruntimeTrace); err != nil {
			log.Errorf("invalid --trace: %s", err.Error())
			return
		}
	}

	if resp, err := client.InvokeModelWithResponseStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists asynchronous invocations.
func bedrockruntime_ListAsyncInvokes(cfg aws.Config, client *bedrockruntime.Client) {
	input := &bedrockruntime.ListAsyncInvokesInput{}

	if len(_bedrockruntimeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bedrockruntimeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bedrockruntimeNextToken) > 0 {
		input.NextToken = aws.String(_bedrockruntimeNextToken)
	}
	if len(_bedrockruntimeSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _bedrockruntimeSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_bedrockruntimeSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _bedrockruntimeSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_bedrockruntimeStatusEquals) > 0 {
		if err := assignInputField(input, "StatusEquals", _bedrockruntimeStatusEquals); err != nil {
			log.Errorf("invalid --status-equals: %s", err.Error())
			return
		}
	}
	if len(_bedrockruntimeSubmitTimeAfter) > 0 {
		if err := assignInputField(input, "SubmitTimeAfter", _bedrockruntimeSubmitTimeAfter); err != nil {
			log.Errorf("invalid --submit-time-after: %s", err.Error())
			return
		}
	}
	if len(_bedrockruntimeSubmitTimeBefore) > 0 {
		if err := assignInputField(input, "SubmitTimeBefore", _bedrockruntimeSubmitTimeBefore); err != nil {
			log.Errorf("invalid --submit-time-before: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAsyncInvokes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bedrockruntime.ListAsyncInvokesOutput
	p := bedrockruntime.NewListAsyncInvokesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Starts an asynchronous invocation.
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
func bedrockruntime_StartAsyncInvoke(cfg aws.Config, client *bedrockruntime.Client) {
	input := &bedrockruntime.StartAsyncInvokeInput{
		// ModelId: *string, // Required
		// ModelInput: document.Interface, // Required
		// OutputDataConfig: types.AsyncInvokeOutputDataConfig, // Required
	}

	if len(_bedrockruntimeModelId) > 0 {
		input.ModelId = aws.String(_bedrockruntimeModelId)
	}
	if len(_bedrockruntimeModelInput) > 0 {
		if err := assignInputField(input, "ModelInput", _bedrockruntimeModelInput); err != nil {
			log.Errorf("invalid --model-input: %s", err.Error())
			return
		}
	}
	if len(_bedrockruntimeOutputDataConfig) > 0 {
		if err := assignInputField(input, "OutputDataConfig", _bedrockruntimeOutputDataConfig); err != nil {
			log.Errorf("invalid --output-data-config: %s", err.Error())
			return
		}
	}
	if len(_bedrockruntimeClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_bedrockruntimeClientRequestToken)
	}
	if len(_bedrockruntimeTags) > 0 {
		if err := assignInputField(input, "Tags", _bedrockruntimeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartAsyncInvoke(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_bedrockruntimeCmd)
	_bedrockruntimeCmd.Flags().SortFlags = false

	_bedrockruntimeCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_bedrockruntimeCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_bedrockruntimeCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_bedrockruntimeCmd.Flags().StringVarP(&_bedrockruntimeAccept, "accept", "", "", "Accept")
	_bedrockruntimeCmd.Flags().StringVarP(&_bedrockruntimeAdditionalModelRequestFields, "additional-model-request-fields", "", "", "Additional Model Request Fields")
	_bedrockruntimeCmd.Flags().StringSliceVarP(&_bedrockruntimeAdditionalModelResponseFieldPaths, "additional-model-response-field-paths", "", nil, "Additional Model Response Field Paths")
	_bedrockruntimeCmd.Flags().StringVarP(&_bedrockruntimeBody, "body", "", "", "Body")
	_bedrockruntimeCmd.Flags().StringVarP(&_bedrockruntimeClientRequestToken, "client-request-token", "", "", "Client Request Token")
	_bedrockruntimeCmd.Flags().StringVarP(&_bedrockruntimeContent, "content", "", "", "Content")
	_bedrockruntimeCmd.Flags().StringVarP(&_bedrockruntimeContentType, "content-type", "", "", "Content Type")
	_bedrockruntimeCmd.Flags().StringVarP(&_bedrockruntimeGuardrailConfig, "guardrail-config", "", "", "Guardrail Config")
	_bedrockruntimeCmd.Flags().StringVarP(&_bedrockruntimeGuardrailIdentifier, "guardrail-identifier", "", "", "Guardrail Identifier")
	_bedrockruntimeCmd.Flags().StringVarP(&_bedrockruntimeGuardrailVersion, "guardrail-version", "", "", "Guardrail Version")
	_bedrockruntimeCmd.Flags().StringVarP(&_bedrockruntimeInferenceConfig, "inference-config", "", "", "Inference Config")
	_bedrockruntimeCmd.Flags().StringVarP(&_bedrockruntimeInput, "input", "", "", "Input")
	_bedrockruntimeCmd.Flags().StringVarP(&_bedrockruntimeInvocationArn, "invocation-arn", "", "", "Invocation ARN")
	_bedrockruntimeCmd.Flags().StringVarP(&_bedrockruntimeMaxResults, "max-results", "", "", "Max Results")
	_bedrockruntimeCmd.Flags().StringVarP(&_bedrockruntimeMessages, "messages", "", "", "Messages")
	_bedrockruntimeCmd.Flags().StringVarP(&_bedrockruntimeModelId, "model-id", "", "", "Model ID")
	_bedrockruntimeCmd.Flags().StringVarP(&_bedrockruntimeModelInput, "model-input", "", "", "Model Input")
	_bedrockruntimeCmd.Flags().StringVarP(&_bedrockruntimeNextToken, "next-token", "", "", "Next Token")
	_bedrockruntimeCmd.Flags().StringVarP(&_bedrockruntimeOutputConfig, "output-config", "", "", "Output Config")
	_bedrockruntimeCmd.Flags().StringVarP(&_bedrockruntimeOutputDataConfig, "output-data-config", "", "", "Output Data Config")
	_bedrockruntimeCmd.Flags().StringVarP(&_bedrockruntimeOutputScope, "output-scope", "", "", "Output Scope")
	_bedrockruntimeCmd.Flags().StringVarP(&_bedrockruntimePerformanceConfig, "performance-config", "", "", "Performance Config")
	_bedrockruntimeCmd.Flags().StringVarP(&_bedrockruntimePerformanceConfigLatency, "performance-config-latency", "", "", "Performance Config Latency")
	_bedrockruntimeCmd.Flags().StringVarP(&_bedrockruntimePromptVariables, "prompt-variables", "", "", "Prompt Variables")
	_bedrockruntimeCmd.Flags().StringVarP(&_bedrockruntimeRequestMetadata, "request-metadata", "", "", "Request Metadata")
	_bedrockruntimeCmd.Flags().StringVarP(&_bedrockruntimeServiceTier, "service-tier", "", "", "Service Tier")
	_bedrockruntimeCmd.Flags().StringVarP(&_bedrockruntimeSortBy, "sort-by", "", "", "Sort By")
	_bedrockruntimeCmd.Flags().StringVarP(&_bedrockruntimeSortOrder, "sort-order", "", "", "Sort Order")
	_bedrockruntimeCmd.Flags().StringVarP(&_bedrockruntimeSource, "source", "", "", "Source")
	_bedrockruntimeCmd.Flags().StringVarP(&_bedrockruntimeStatusEquals, "status-equals", "", "", "Status Equals")
	_bedrockruntimeCmd.Flags().StringVarP(&_bedrockruntimeSubmitTimeAfter, "submit-time-after", "", "", "Submit Time After")
	_bedrockruntimeCmd.Flags().StringVarP(&_bedrockruntimeSubmitTimeBefore, "submit-time-before", "", "", "Submit Time Before")
	_bedrockruntimeCmd.Flags().StringVarP(&_bedrockruntimeSystem, "system", "", "", "System")
	_bedrockruntimeCmd.Flags().StringVarP(&_bedrockruntimeTags, "tags", "", "", "Tags")
	_bedrockruntimeCmd.Flags().StringVarP(&_bedrockruntimeToolConfig, "tool-config", "", "", "Tool Config")
	_bedrockruntimeCmd.Flags().StringVarP(&_bedrockruntimeTrace, "trace", "", "", "Trace")

	_bedrockruntimeCmd.Flags().BoolVarP(&_bedrockruntimeApplyGuardrail, "apply-guardrail", "", false, "Apply Guardrail")
	_bedrockruntimeCmd.Flags().BoolVarP(&_bedrockruntimeConverse, "converse", "", false, "Converse")
	_bedrockruntimeCmd.Flags().BoolVarP(&_bedrockruntimeConverseStream, "converse-stream", "", false, "Converse Stream")
	_bedrockruntimeCmd.Flags().BoolVarP(&_bedrockruntimeCountTokens, "count-tokens", "", false, "Count Tokens")
	_bedrockruntimeCmd.Flags().BoolVarP(&_bedrockruntimeGetAsyncInvoke, "get-async-invoke", "", false, "Get Async Invoke")
	_bedrockruntimeCmd.Flags().BoolVarP(&_bedrockruntimeInvokeModel, "invoke-model", "", false, "Invoke Model")
	_bedrockruntimeCmd.Flags().BoolVarP(&_bedrockruntimeInvokeModelWithBidirectionalStream, "invoke-model-with-bidirectional-stream", "", false, "Invoke Model With Bidirectional Stream")
	_bedrockruntimeCmd.Flags().BoolVarP(&_bedrockruntimeInvokeModelWithResponseStream, "invoke-model-with-response-stream", "", false, "Invoke Model With Response Stream")
	_bedrockruntimeCmd.Flags().BoolVarP(&_bedrockruntimeListAsyncInvokes, "list-async-invokes", "", false, "List Async Invokes")
	_bedrockruntimeCmd.Flags().BoolVarP(&_bedrockruntimeStartAsyncInvoke, "start-async-invoke", "", false, "Start Async Invoke")

}
