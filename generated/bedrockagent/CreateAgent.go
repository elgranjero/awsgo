package bedrockagent

// CreateAgent is generated as a reference stub.
// Executable command wiring lives under cmd/bedrockagent.go.
//
// Creates an agent that orchestrates interactions between foundation models, data
// sources, software applications, user conversations, and APIs to carry out tasks
// to help customers.
//
// - Specify the following fields for security purposes.
//
// - agentResourceRoleArn – The Amazon Resource Name (ARN) of the role with
// permissions to invoke API operations on an agent.
//
// - (Optional) customerEncryptionKeyArn – The Amazon Resource Name (ARN) of a
// KMS key to encrypt the creation of the agent.
//
// - (Optional) idleSessionTTLinSeconds – Specify the number of seconds for which
// the agent should maintain session information. After this time expires, the
// subsequent InvokeAgent request begins a new session.
//
// - To enable your agent to retain conversational context across multiple
// sessions, include a memoryConfiguration object. For more information, see [Configure memory].
//
// - To override the default prompt behavior for agent orchestration and to use
// advanced prompts, include a promptOverrideConfiguration object. For more
// information, see [Advanced prompts].
//
// - If your agent fails to be created, the response returns a list of
// failureReasons alongside a list of recommendedActions for you to troubleshoot.
//
// - The agent instructions will not be honored if your agent has only one
// knowledge base, uses default prompts, has no action group, and user input is
// disabled.
//
// [Advanced prompts]: https://docs.aws.amazon.com/bedrock/latest/userguide/advanced-prompts.html
// [Configure memory]: https://docs.aws.amazon.com/bedrock/latest/userguide/agents-configure-memory.html
