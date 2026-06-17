package wafregional

// CreateWebACL is generated as a reference stub.
// Executable command wiring lives under cmd/wafregional.go.
//
// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Creates a WebACL , which contains the Rules that identify the CloudFront web
// requests that you want to allow, block, or count. AWS WAF evaluates Rules in
// order based on the value of Priority for each Rule .
//
// You also specify a default action, either ALLOW or BLOCK . If a web request
// doesn't match any of the Rules in a WebACL , AWS WAF responds to the request
// with the default action.
//
// To create and configure a WebACL , perform the following steps:
//
// - Create and update the ByteMatchSet objects and other predicates that you
// want to include in Rules . For more information, see CreateByteMatchSet, UpdateByteMatchSet, CreateIPSet, UpdateIPSet, CreateSqlInjectionMatchSet, and UpdateSqlInjectionMatchSet.
//
// - Create and update the Rules that you want to include in the WebACL . For
// more information, see CreateRuleand UpdateRule.
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// a CreateWebACL request.
//
// - Submit a CreateWebACL request.
//
// - Use GetChangeToken to get the change token that you provide in the
// ChangeToken parameter of an UpdateWebACLrequest.
//
// - Submit an UpdateWebACLrequest to specify the Rules that you want to include in the
// WebACL , to specify the default action, and to associate the WebACL with a
// CloudFront distribution.
//
// For more information about how to use the AWS WAF API, see the [AWS WAF Developer Guide].
//
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
