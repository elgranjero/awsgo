package bedrock

// UpdateGuardrail is generated as a reference stub.
// Executable command wiring lives under cmd/bedrock.go.
//
// Updates a guardrail with the values you specify.
//
// - Specify a name and optional description .
//
// - Specify messages for when the guardrail successfully blocks a prompt or a
// model response in the blockedInputMessaging and blockedOutputsMessaging fields.
//
// - Specify topics for the guardrail to deny in the topicPolicyConfig object.
// Each [GuardrailTopicConfig]object in the topicsConfig list pertains to one topic.
//
// - Give a name and description so that the guardrail can properly identify the
// topic.
//
// - Specify DENY in the type field.
//
// - (Optional) Provide up to five prompts that you would categorize as
// belonging to the topic in the examples list.
//
// - Specify filter strengths for the harmful categories defined in Amazon
// Bedrock in the contentPolicyConfig object. Each [GuardrailContentFilterConfig]object in the filtersConfig
// list pertains to a harmful category. For more information, see [Content filters]. For more
// information about the fields in a content filter, see [GuardrailContentFilterConfig].
//
// - Specify the category in the type field.
//
// - Specify the strength of the filter for prompts in the inputStrength field
// and for model responses in the strength field of the [GuardrailContentFilterConfig].
//
// - (Optional) For security, include the ARN of a KMS key in the kmsKeyId field.
//
// [GuardrailContentFilterConfig]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_GuardrailContentFilterConfig.html
// [Content filters]: https://docs.aws.amazon.com/bedrock/latest/userguide/guardrails-content-filters
// [GuardrailTopicConfig]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_GuardrailTopicConfig.html
