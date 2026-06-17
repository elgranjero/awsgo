package bedrockagentruntime

// InvokeInlineAgent is generated as a reference stub.
// Executable command wiring lives under cmd/bedrockagentruntime.go.
//
// Invokes an inline Amazon Bedrock agent using the configurations you provide
//
// with the request.
//
// - Specify the following fields for security purposes.
//
// - (Optional) customerEncryptionKeyArn – The Amazon Resource Name (ARN) of a
// KMS key to encrypt the creation of the agent.
//
// - (Optional) idleSessionTTLinSeconds – Specify the number of seconds for which
// the agent should maintain session information. After this time expires, the
// subsequent InvokeInlineAgent request begins a new session.
//
// - To override the default prompt behavior for agent orchestration and to use
// advanced prompts, include a promptOverrideConfiguration object. For more
// information, see [Advanced prompts].
//
// - The agent instructions will not be honored if your agent has only one
// knowledge base, uses default prompts, has no action group, and user input is
// disabled.
//
// [Advanced prompts]: https://docs.aws.amazon.com/bedrock/latest/userguide/advanced-prompts.html
