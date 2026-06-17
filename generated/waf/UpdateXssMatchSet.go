package waf

// UpdateXssMatchSet is generated as a reference stub.
// Executable command wiring lives under cmd/waf.go.
//
// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Inserts or deletes XssMatchTuple objects (filters) in an XssMatchSet. For each XssMatchTuple object,
// you specify the following values:
//
// - Action : Whether to insert the object into or delete the object from the
// array. To change an XssMatchTuple , you delete the existing object and add a
// new one.
//
// - FieldToMatch : The part of web requests that you want AWS WAF to inspect
// and, if you want AWS WAF to inspect a header or custom query parameter, the name
// of the header or parameter.
//
// - TextTransformation : Which text transformation, if any, to perform on the
// web request before inspecting the request for cross-site scripting attacks.
//
// You can only specify a single type of TextTransformation.
//
// You use XssMatchSet objects to specify which CloudFront requests that you want
// to allow, block, or count. For example, if you're receiving requests that
// contain cross-site scripting attacks in the request body and you want to block
// the requests, you can create an XssMatchSet with the applicable settings, and
// then configure AWS WAF to block the requests.
//
// To create and configure an XssMatchSet , perform the following steps:
//
// - Submit a CreateXssMatchSetrequest.
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// an UpdateIPSetrequest.
//
// - Submit an UpdateXssMatchSet request to specify the parts of web requests
// that you want AWS WAF to inspect for cross-site scripting attacks.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
