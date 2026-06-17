package wafregional

// DeleteRuleGroup is generated as a reference stub.
// Executable command wiring lives under cmd/wafregional.go.
//
// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Permanently deletes a RuleGroup. You can't delete a RuleGroup if it's still used in any
// WebACL objects or if it still includes any rules.
//
// If you just want to remove a RuleGroup from a WebACL , use UpdateWebACL.
//
// To permanently delete a RuleGroup from AWS WAF, perform the following steps:
//
// - Update the RuleGroup to remove rules, if any. For more information, see UpdateRuleGroup.
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// a DeleteRuleGroup request.
//
// - Submit a DeleteRuleGroup request.
//
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
