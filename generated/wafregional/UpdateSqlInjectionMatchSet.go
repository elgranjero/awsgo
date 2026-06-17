package wafregional

// UpdateSqlInjectionMatchSet is generated as a reference stub.
// Executable command wiring lives under cmd/wafregional.go.
//
// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Inserts or deletes SqlInjectionMatchTuple objects (filters) in a SqlInjectionMatchSet. For each SqlInjectionMatchTuple
// object, you specify the following values:
//
// - Action : Whether to insert the object into or delete the object from the
// array. To change a SqlInjectionMatchTuple , you delete the existing object and
// add a new one.
//
// - FieldToMatch : The part of web requests that you want AWS WAF to inspect
// and, if you want AWS WAF to inspect a header or custom query parameter, the name
// of the header or parameter.
//
// - TextTransformation : Which text transformation, if any, to perform on the
// web request before inspecting the request for snippets of malicious SQL code.
//
// You can only specify a single type of TextTransformation.
//
// You use SqlInjectionMatchSet objects to specify which CloudFront requests that
// you want to allow, block, or count. For example, if you're receiving requests
// that contain snippets of SQL code in the query string and you want to block the
// requests, you can create a SqlInjectionMatchSet with the applicable settings,
// and then configure AWS WAF to block the requests.
//
// To create and configure a SqlInjectionMatchSet , perform the following steps:
//
// - Submit a CreateSqlInjectionMatchSetrequest.
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// an UpdateIPSetrequest.
//
// - Submit an UpdateSqlInjectionMatchSet request to specify the parts of web
// requests that you want AWS WAF to inspect for snippets of SQL code.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
