package wafv2

// CheckCapacity is generated as a reference stub.
// Executable command wiring lives under cmd/wafv2.go.
//
// Returns the web ACL capacity unit (WCU) requirements for a specified scope and
// set of rules. You can use this to check the capacity requirements for the rules
// you want to use in a RuleGroupor WebACL.
//
// WAF uses WCUs to calculate and control the operating resources that are used to
// run your rules, rule groups, and web ACLs. WAF calculates capacity differently
// for each rule type, to reflect the relative cost of each rule. Simple rules that
// cost little to run use fewer WCUs than more complex rules that use more
// processing power. Rule group capacity is fixed at creation, which helps users
// plan their web ACL WCU usage when they use a rule group. For more information,
// see [WAF web ACL capacity units (WCU)]in the WAF Developer Guide.
//
// [WAF web ACL capacity units (WCU)]: https://docs.aws.amazon.com/waf/latest/developerguide/aws-waf-capacity-units.html
