package waf

// UpdateRuleGroup is generated as a reference stub.
// Executable command wiring lives under cmd/waf.go.
//
// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Inserts or deletes ActivatedRule objects in a RuleGroup .
//
// You can only insert REGULAR rules into a rule group.
//
// You can have a maximum of ten rules per rule group.
//
// To create and configure a RuleGroup , perform the following steps:
//
// - Create and update the Rules that you want to include in the RuleGroup . See CreateRule
// .
//
// - Use GetChangeToken to get the change token that you provide in the
// ChangeToken parameter of an UpdateRuleGrouprequest.
//
// - Submit an UpdateRuleGroup request to add Rules to the RuleGroup .
//
// - Create and update a WebACL that contains the RuleGroup . See CreateWebACL.
//
// If you want to replace one Rule with another, you delete the existing one and
// add the new one.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
