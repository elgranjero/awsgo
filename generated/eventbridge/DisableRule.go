package eventbridge

// DisableRule is generated as a reference stub.
// Executable command wiring lives under cmd/eventbridge.go.
//
// Disables the specified rule. A disabled rule won't match any events, and won't
// self-trigger if it has a schedule expression.
//
// When you disable a rule, incoming events might continue to match to the
// disabled rule. Allow a short period of time for changes to take effect.
