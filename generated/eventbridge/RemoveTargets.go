package eventbridge

// RemoveTargets is generated as a reference stub.
// Executable command wiring lives under cmd/eventbridge.go.
//
// Removes the specified targets from the specified rule. When the rule is
// triggered, those targets are no longer be invoked.
//
// A successful execution of RemoveTargets doesn't guarantee all targets are
// removed from the rule, it means that the target(s) listed in the request are
// removed.
//
// When you remove a target, when the associated rule triggers, removed targets
// might continue to be invoked. Allow a short period of time for changes to take
// effect.
//
// This action can partially fail if too many requests are made at the same time.
// If that happens, FailedEntryCount is non-zero in the response and each entry in
// FailedEntries provides the ID of the failed target and the error code.
//
// The maximum number of entries per request is 10.
