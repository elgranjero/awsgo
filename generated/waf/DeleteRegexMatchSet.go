package waf

// DeleteRegexMatchSet is generated as a reference stub.
// Executable command wiring lives under cmd/waf.go.
//
// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Permanently deletes a RegexMatchSet. You can't delete a RegexMatchSet if it's still used in
// any Rules or if it still includes any RegexMatchTuples objects (any filters).
//
// If you just want to remove a RegexMatchSet from a Rule , use UpdateRule.
//
// To permanently delete a RegexMatchSet , perform the following steps:
//
// - Update the RegexMatchSet to remove filters, if any. For more information,
// see UpdateRegexMatchSet.
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// a DeleteRegexMatchSet request.
//
// - Submit a DeleteRegexMatchSet request.
//
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
