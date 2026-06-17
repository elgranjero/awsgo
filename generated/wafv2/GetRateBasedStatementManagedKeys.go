package wafv2

// GetRateBasedStatementManagedKeys is generated as a reference stub.
// Executable command wiring lives under cmd/wafv2.go.
//
// Retrieves the IP addresses that are currently blocked by a rate-based rule
// instance. This is only available for rate-based rules that aggregate solely on
// the IP address or on the forwarded IP address.
//
// The maximum number of addresses that can be blocked for a single rate-based
// rule instance is 10,000. If more than 10,000 addresses exceed the rate limit,
// those with the highest rates are blocked.
//
// For a rate-based rule that you've defined inside a rule group, provide the name
// of the rule group reference statement in your request, in addition to the
// rate-based rule name and the web ACL name.
//
// WAF monitors web requests and manages keys independently for each unique
// combination of web ACL, optional rule group, and rate-based rule. For example,
// if you define a rate-based rule inside a rule group, and then use the rule group
// in a web ACL, WAF monitors web requests and manages keys for that web ACL, rule
// group reference statement, and rate-based rule instance. If you use the same
// rule group in a second web ACL, WAF monitors web requests and manages keys for
// this second usage completely independent of your first.
