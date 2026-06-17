package waf

// PutPermissionPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/waf.go.
//
// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Attaches an IAM policy to the specified resource. The only supported use for
// this action is to share a RuleGroup across accounts.
//
// The PutPermissionPolicy is subject to the following restrictions:
//
// - You can attach only one policy with each PutPermissionPolicy request.
//
// - The policy must include an Effect , Action and Principal .
//
// - Effect must specify Allow .
//
// - The Action in the policy must be waf:UpdateWebACL ,
// waf-regional:UpdateWebACL , waf:GetRuleGroup and waf-regional:GetRuleGroup .
// Any extra or wildcard actions in the policy will be rejected.
//
// - The policy cannot include a Resource parameter.
//
// - The ARN in the request must be a valid WAF RuleGroup ARN and the RuleGroup
// must exist in the same region.
//
// - The user making the request must be the owner of the RuleGroup.
//
// - Your policy must be composed using IAM Policy version 2012-10-17.
//
// For more information, see [IAM Policies].
//
// An example of a valid policy parameter is shown in the Examples section below.
//
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [IAM Policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html
