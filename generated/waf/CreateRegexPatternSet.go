package waf

// CreateRegexPatternSet is generated as a reference stub.
// Executable command wiring lives under cmd/waf.go.
//
// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Creates a RegexPatternSet . You then use UpdateRegexPatternSet to specify the regular expression
// (regex) pattern that you want AWS WAF to search for, such as B[a(at)]dB[o0]t . You
// can then configure AWS WAF to reject those requests.
//
// To create and configure a RegexPatternSet , perform the following steps:
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// a CreateRegexPatternSet request.
//
// - Submit a CreateRegexPatternSet request.
//
// - Use GetChangeToken to get the change token that you provide in the
// ChangeToken parameter of an UpdateRegexPatternSet request.
//
// - Submit an UpdateRegexPatternSetrequest to specify the string that you want AWS WAF to watch for.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
