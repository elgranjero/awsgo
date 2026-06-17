package bedrock

// CreateGuardrail is generated as a reference stub.
// Executable command wiring lives under cmd/bedrock.go.
//
// Creates a guardrail to block topics and to implement safeguards for your
// generative AI applications.
//
// You can configure the following policies in a guardrail to avoid undesirable
// and harmful content, filter out denied topics and words, and remove sensitive
// information for privacy protection.
//
// - Content filters - Adjust filter strengths to block input prompts or model
// responses containing harmful content.
//
// - Denied topics - Define a set of topics that are undesirable in the context
// of your application. These topics will be blocked if detected in user queries or
// model responses.
//
// - Word filters - Configure filters to block undesirable words, phrases, and
// profanity. Such words can include offensive terms, competitor names etc.
//
// - Sensitive information filters - Block or mask sensitive information such as
// personally identifiable information (PII) or custom regex in user inputs and
// model responses.
//
// In addition to the above policies, you can also configure the messages to be
// returned to the user if a user input or model response is in violation of the
// policies defined in the guardrail.
//
// For more information, see [Amazon Bedrock Guardrails] in the Amazon Bedrock User Guide.
//
// [Amazon Bedrock Guardrails]: https://docs.aws.amazon.com/bedrock/latest/userguide/guardrails.html
