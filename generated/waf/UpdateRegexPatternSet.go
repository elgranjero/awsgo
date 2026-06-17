package waf

// UpdateRegexPatternSet is generated as a reference stub.
// Executable command wiring lives under cmd/waf.go.
//
// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Inserts or deletes RegexPatternString objects in a RegexPatternSet. For each RegexPatternString
// object, you specify the following values:
//
// - Whether to insert or delete the RegexPatternString .
//
// - The regular expression pattern that you want to insert or delete. For more
// information, see RegexPatternSet.
//
// For example, you can create a RegexPatternString such as B[a(at)]dB[o0]t . AWS WAF
// will match this RegexPatternString to:
//
// - BadBot
//
// - BadB0t
//
// - B(at)dBot
//
// - B(at)dB0t
//
// To create and configure a RegexPatternSet , perform the following steps:
//
// - Create a RegexPatternSet. For more information, see CreateRegexPatternSet.
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// an UpdateRegexPatternSet request.
//
// - Submit an UpdateRegexPatternSet request to specify the regular expression
// pattern that you want AWS WAF to watch for.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
