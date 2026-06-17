package securityhub

// StartConfigurationPolicyDisassociation is generated as a reference stub.
// Executable command wiring lives under cmd/securityhub.go.
//
// Disassociates a target account, organizational unit, or the root from a
//
// specified configuration. When you disassociate a configuration from its target,
// the target inherits the configuration of the closest parent. If there’s no
// configuration to inherit, the target retains its settings but becomes a
// self-managed account. A target can be disassociated from a configuration policy
// or self-managed behavior. Only the Security Hub CSPM delegated administrator can
// invoke this operation from the home Region.
