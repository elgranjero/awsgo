package waf

// UpdateRegexMatchSet is generated as a reference stub.
// Executable command wiring lives under cmd/waf.go.
//
// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Inserts or deletes RegexMatchTuple objects (filters) in a RegexMatchSet. For each RegexMatchSetUpdate
// object, you specify the following values:
//
// - Whether to insert or delete the object from the array. If you want to
// change a RegexMatchSetUpdate object, you delete the existing object and add a
// new one.
//
// - The part of a web request that you want AWS WAF to inspectupdate, such as a
// query string or the value of the User-Agent header.
//
// - The identifier of the pattern (a regular expression) that you want AWS WAF
// to look for. For more information, see RegexPatternSet.
//
// - Whether to perform any conversions on the request, such as converting it to
// lowercase, before inspecting it for the specified string.
//
// For example, you can create a RegexPatternSet that matches any requests with
// User-Agent headers that contain the string B[a(at)]dB[o0]t . You can then configure
// AWS WAF to reject those requests.
//
// To create and configure a RegexMatchSet , perform the following steps:
//
// - Create a RegexMatchSet. For more information, see CreateRegexMatchSet.
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// an UpdateRegexMatchSet request.
//
// - Submit an UpdateRegexMatchSet request to specify the part of the request
// that you want AWS WAF to inspect (for example, the header or the URI) and the
// identifier of the RegexPatternSet that contain the regular expression patters
// you want AWS WAF to watch for.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
