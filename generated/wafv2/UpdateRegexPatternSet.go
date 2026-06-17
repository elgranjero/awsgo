package wafv2

// UpdateRegexPatternSet is generated as a reference stub.
// Executable command wiring lives under cmd/wafv2.go.
//
// Updates the specified RegexPatternSet.
//
// This operation completely replaces the mutable specifications that you already
// have for the regex pattern set with the ones that you provide to this call.
//
// To modify a regex pattern set, do the following:
//
// - Retrieve it by calling GetRegexPatternSet
//
// - Update its settings as needed
//
// - Provide the complete regex pattern set specification to this call
//
// # Temporary inconsistencies during updates
//
// When you create or change a web ACL or other WAF resources, the changes take a
// small amount of time to propagate to all areas where the resources are stored.
// The propagation time can be from a few seconds to a number of minutes.
//
// The following are examples of the temporary inconsistencies that you might
// notice during change propagation:
//
// - After you create a web ACL, if you try to associate it with a resource, you
// might get an exception indicating that the web ACL is unavailable.
//
// - After you add a rule group to a web ACL, the new rule group rules might be
// in effect in one area where the web ACL is used and not in another.
//
// - After you change a rule action setting, you might see the old action in
// some places and the new action in others.
//
// - After you add an IP address to an IP set that is in use in a blocking rule,
// the new address might be blocked in one area while still allowed in another.
