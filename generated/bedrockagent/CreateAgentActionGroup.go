package bedrockagent

// CreateAgentActionGroup is generated as a reference stub.
// Executable command wiring lives under cmd/bedrockagent.go.
//
// Creates an action group for an agent. An action group represents the actions
// that an agent can carry out for the customer by defining the APIs that an agent
// can call and the logic for calling them.
//
// To allow your agent to request the user for additional information when trying
// to complete a task, add an action group with the parentActionGroupSignature
// field set to AMAZON.UserInput .
//
// To allow your agent to generate, run, and troubleshoot code when trying to
// complete a task, add an action group with the parentActionGroupSignature field
// set to AMAZON.CodeInterpreter .
//
// You must leave the description , apiSchema , and actionGroupExecutor fields
// blank for this action group. During orchestration, if your agent determines that
// it needs to invoke an API in an action group, but doesn't have enough
// information to complete the API request, it will invoke this action group
// instead and return an [Observation]reprompting the user for more information.
//
// [Observation]: https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_Observation.html
