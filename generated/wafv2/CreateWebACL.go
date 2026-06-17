package wafv2

// CreateWebACL is generated as a reference stub.
// Executable command wiring lives under cmd/wafv2.go.
//
// Creates a WebACL per the specifications provided.
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
