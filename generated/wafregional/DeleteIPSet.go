package wafregional

// DeleteIPSet is generated as a reference stub.
// Executable command wiring lives under cmd/wafregional.go.
//
// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Permanently deletes an IPSet. You can't delete an IPSet if it's still used in any
// Rules or if it still includes any IP addresses.
//
// If you just want to remove an IPSet from a Rule , use UpdateRule.
//
// To permanently delete an IPSet from AWS WAF, perform the following steps:
//
// - Update the IPSet to remove IP address ranges, if any. For more information,
// see UpdateIPSet.
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// a DeleteIPSet request.
//
// - Submit a DeleteIPSet request.
//
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
