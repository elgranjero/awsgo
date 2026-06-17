package wafv2

// UpdateWebACL is generated as a reference stub.
// Executable command wiring lives under cmd/wafv2.go.
//
// Updates the specified WebACL. While updating a web ACL, WAF provides continuous
// coverage to the resources that you have associated with the web ACL.
//
// This operation completely replaces the mutable specifications that you already
// have for the web ACL with the ones that you provide to this call.
//
// To modify a web ACL, do the following:
//
// - Retrieve it by calling GetWebACL
//
// - Update its settings as needed
//
// - Provide the complete web ACL specification to this call
//
// A web ACL defines a collection of rules to use to inspect and control web
// requests. Each rule has a statement that defines what to look for in web
// requests and an action that WAF applies to requests that match the statement. In
// the web ACL, you assign a default action to take (allow, block) for any request
// that does not match any of the rules. The rules in a web ACL can be a
// combination of the types Rule, RuleGroup, and managed rule group. You can associate a web
// ACL with one or more Amazon Web Services resources to protect. The resource
// types include Amazon CloudFront distribution, Amazon API Gateway REST API,
// Application Load Balancer, AppSync GraphQL API, Amazon Cognito user pool, App
// Runner service, Amplify application, and Amazon Web Services Verified Access
// instance.
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
