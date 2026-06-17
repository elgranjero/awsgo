package ssm

// GetCommandInvocation is generated as a reference stub.
// Executable command wiring lives under cmd/ssm.go.
//
// Returns detailed information about command execution for an invocation or
// plugin. The Run Command API follows an eventual consistency model, due to the
// distributed nature of the system supporting the API. This means that the result
// of an API command you run that affects your resources might not be immediately
// visible to all subsequent commands you run. You should keep this in mind when
// you carry out an API command that immediately follows a previous API command.
//
// GetCommandInvocation only gives the execution status of a plugin in a document.
// To get the command execution status on a specific managed node, use ListCommandInvocations. To get
// the command execution status across managed nodes, use ListCommands.
