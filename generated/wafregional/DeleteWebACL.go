package wafregional

// DeleteWebACL is generated as a reference stub.
// Executable command wiring lives under cmd/wafregional.go.
//
// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Permanently deletes a WebACL. You can't delete a WebACL if it still contains any Rules
// .
//
// To delete a WebACL , perform the following steps:
//
// - Update the WebACL to remove Rules , if any. For more information, see UpdateWebACL.
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// a DeleteWebACL request.
//
// - Submit a DeleteWebACL request.
//
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
