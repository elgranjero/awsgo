package waf

// CreateSqlInjectionMatchSet is generated as a reference stub.
// Executable command wiring lives under cmd/waf.go.
//
// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Creates a SqlInjectionMatchSet, which you use to allow, block, or count requests that contain
// snippets of SQL code in a specified part of web requests. AWS WAF searches for
// character sequences that are likely to be malicious strings.
//
// To create and configure a SqlInjectionMatchSet , perform the following steps:
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// a CreateSqlInjectionMatchSet request.
//
// - Submit a CreateSqlInjectionMatchSet request.
//
// - Use GetChangeToken to get the change token that you provide in the
// ChangeToken parameter of an UpdateSqlInjectionMatchSetrequest.
//
// - Submit an UpdateSqlInjectionMatchSetrequest to specify the parts of web requests in which you want to
// allow, block, or count malicious SQL code.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
