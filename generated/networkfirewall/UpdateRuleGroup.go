package networkfirewall

// UpdateRuleGroup is generated as a reference stub.
// Executable command wiring lives under cmd/networkfirewall.go.
//
// Updates the rule settings for the specified rule group. You use a rule group by
// reference in one or more firewall policies. When you modify a rule group, you
// modify all firewall policies that use the rule group.
//
// To update a rule group, first call DescribeRuleGroup to retrieve the current RuleGroup object, update the
// object as needed, and then provide the updated object to this call.
