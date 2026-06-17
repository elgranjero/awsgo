package amp

// PutRuleGroupsNamespace is generated as a reference stub.
// Executable command wiring lives under cmd/amp.go.
//
// Updates an existing rule groups namespace within a workspace. A rule groups
// namespace is associated with exactly one rules file. A workspace can have
// multiple rule groups namespaces.
//
// The combined length of a rule group namespace and a rule group name cannot
// exceed 721 UTF-8 bytes.
//
// Use this operation only to update existing rule groups namespaces. To create a
// new rule groups namespace, use CreateRuleGroupsNamespace .
//
// You can't use this operation to add tags to an existing rule groups namespace.
// Instead, use TagResource .
