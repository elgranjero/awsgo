package wafv2

// PutPermissionPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/wafv2.go.
//
// Use this to share a rule group with other accounts.
//
// This action attaches an IAM policy to the specified resource. You must be the
// owner of the rule group to perform this operation.
//
// This action is subject to the following restrictions:
//
// - You can attach only one policy with each PutPermissionPolicy request.
//
// - The ARN in the request must be a valid WAF RuleGroupARN and the rule group must
// exist in the same Region.
//
// - The user making the request must be the owner of the rule group.
//
// If a rule group has been shared with your account, you can access it through
// the call GetRuleGroup , and you can reference it in CreateWebACL and
// UpdateWebACL . Rule groups that are shared with you don't appear in your WAF
// console rule groups listing.
